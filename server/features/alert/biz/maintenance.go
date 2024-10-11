package alert_biz

import (
	"fmt"
	"reflect"
	"slices"
	"time"

	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/alert/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/global/constants"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/tools/helpers"
	"gorm.io/datatypes"
)

type MaintenanceService struct{}

func NewMaintenanceService() *MaintenanceService {
	return &MaintenanceService{}
}

func (m *MaintenanceService) GetActiveMaintenance() ([]*models.Maintenance, error) {
	timeNow := time.Now().UTC()
	return gen.Maintenance.Where(
		gen.Maintenance.StartedAt.Lte(timeNow),
		gen.Maintenance.EndedAt.Gt(timeNow),
	).Find()
}

func (m *MaintenanceService) CreateMaintenance(mt *schemas.MaintenanceCreate) (string, error) {
	newMt := &models.Maintenance{
		OrganizationId:  global.OrganizationId.Get(),
		StartedAt:       mt.StartedAt,
		EndedAt:         mt.EndedAt,
		MaintenanceType: mt.MaintenanceType,
		Description:     mt.Description,
		Conditions:      condToDbCond(mt.Conditions),
	}

	err := gen.Maintenance.Create(newMt)
	if err != nil {
		return "", err
	}
	timeNow := time.Now().UTC()
	if mt.StartedAt.Before(timeNow) && mt.EndedAt.After(timeNow) {
		err = m.updateAlertByConditions(newMt.Id, newMt.Conditions)
	}
	return newMt.Id, err
}

func (m *MaintenanceService) UpdateMaintenance(mtID string, mt *schemas.MaintenanceUpdate) error {

	dbMt, err := gen.Maintenance.Where(gen.Maintenance.Id.Eq(mtID), gen.Maintenance.OrganizationId.Eq(global.OrganizationId.Get())).First()
	if err != nil {
		return err
	}
	if mt.StartedAt != nil {
		dbMt.StartedAt = *mt.StartedAt
	}
	if mt.EndedAt != nil {
		dbMt.EndedAt = *mt.EndedAt
	}
	if mt.MaintenanceType != nil {
		dbMt.MaintenanceType = *mt.MaintenanceType
	}
	if mt.Description != nil {
		dbMt.Description = mt.Description
	}
	if mt.Conditions != nil {
		dbMt.Conditions = condToDbCond(*mt.Conditions)
	}
	err = gen.Maintenance.UnderlyingDB().Save(dbMt).Error
	if err != nil {
		return err
	}
	timeNow := time.Now().UTC()
	if timeNow.Before(dbMt.StartedAt) || timeNow.After(dbMt.EndedAt) {
		err = m.removeSilencedAlerts(mtID)
		if err != nil {
			return err
		}
	} else if timeNow.After(dbMt.StartedAt) && timeNow.Before(dbMt.EndedAt) {
		err = m.updateAlertByConditions(mtID, dbMt.Conditions)
		if err != nil {
			return err
		}
	}
	return err
}

func (m *MaintenanceService) DeleteMaintenance(mtID string) error {

	_, err := gen.Maintenance.Where(gen.Maintenance.Id.Eq(mtID), gen.Maintenance.OrganizationId.Eq(global.OrganizationId.Get())).Delete()

	if err != nil {
		return err
	}
	return m.removeSilencedAlerts(mtID)
}

func (m *MaintenanceService) GetById(mtID string) (*schemas.Maintenance, error) {
	var result schemas.Maintenance
	err := gen.Maintenance.Where(
		gen.Maintenance.Id.Eq(mtID),
		gen.Maintenance.OrganizationId.Eq(global.OrganizationId.Get())).Preload(
		gen.Maintenance.CreatedBy,
	).Preload(gen.Maintenance.UpdatedBy).Scan(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (m *MaintenanceService) ListMaintenances(query *schemas.MaintenanceQuery) (int64, []*schemas.Maintenance, error) {
	result := make([]*schemas.Maintenance, 0)
	count := int64(0)
	stmt := gen.Maintenance.UnderlyingDB().Where("organization_id = ?", global.OrganizationId.Get())
	if query.Id != nil {
		stmt = stmt.Where("id IN ?", *query.Id)
	}
	if query.MaintenanceType != nil {
		stmt = stmt.Where("maintenanceType = ?", *query.MaintenanceType)
	}
	if query.Status != nil {
		timeNow := time.Now().UTC()
		if *query.Status == "Approaching" {
			stmt = stmt.Where("startedAt < ?", timeNow)
		} else if *query.Status == "Active" {
			stmt = stmt.Where("startedAt > ? AND endedAt < ?", timeNow, timeNow)
		} else if *query.Status == "Ended" {
			stmt = stmt.Where("endedAt < ?", timeNow)
		}
	}
	if query.IsSearchable() {
		keyword := "%" + *query.Keyword + "%"
		stmt = stmt.Where("name ILIKE ?", keyword)
	}
	if query.SiteId != nil {
		stmt = stmt.Where(
			`
			EXISTS (
            SELECT *
            FROM jsonb_array_elements(alert_maintenance.conditions) AS elements
            WHERE elements->>'item' = 'siteId'  AND EXISTS (
            SELECT 1
            FROM jsonb_array_elements_text(elements->'value') AS val
            WHERE val.value IN ?
            )
        )`, *query.SiteId,
		)
	}
	err := stmt.Count(&count).Error
	if err != nil {
		return 0, nil, err
	}
	stmt.Scopes(query.OrderByField())
	stmt.Scopes(query.Pagination())
	err = stmt.Preload("createdBy").Preload("updatedBy").Scan(&result).Error
	if err != nil {
		return 0, nil, err
	}
	return count, result, nil

}

func condToDbCond(cond []schemas.Condition) datatypes.JSONSlice[models.Condition] {
	dbConditions := make([]models.Condition, 0)

	for _, c := range cond {
		dbConditions = append(dbConditions, models.Condition{
			Item:  c.Item,
			Value: c.Value,
		})
	}
	return datatypes.NewJSONSlice(dbConditions)
}

func (m *MaintenanceService) updateAlertByConditions(mtID string, conditions []models.Condition) error {
	stmt := gen.Alert.UnderlyingDB().Where("status = ?", constants.AlertFiringStatus)
	for _, cond := range conditions {
		t := reflect.TypeOf(models.Alert{})
		if helpers.HasStructTypeField(t, cond.Item) {
			if !slices.Equal(cond.Value, []string{"*"}) {
				stmt.Where(fmt.Sprintf("%s IN ?", cond.Item), cond.Value)
			} else {
				stmt.Where(fmt.Sprintf("%s IS NOT NULL", cond.Item))
			}

		}
	}
	stmt.Updates(map[string]any{
		"suppressed":    true,
		"maintenanceId": &mtID,
	})
	return stmt.Error

}

func (m *MaintenanceService) removeSilencedAlerts(mtID string) error {
	_, err := gen.Alert.Where(
		gen.Alert.MaintenanceId.Eq(mtID),
		gen.Alert.Status.Eq(constants.AlertFiringStatus)).
		UpdateColumns(
			map[string]any{"suppressed": false, "maintenanceId": nil})
	return err

}
