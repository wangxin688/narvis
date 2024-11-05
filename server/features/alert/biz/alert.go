package alert_biz

import (
	"fmt"
	"slices"
	"strings"
	"time"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/intend/alerts"
	"github.com/wangxin688/narvis/intend/devicerole"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/dal/gen"
	admin_sc "github.com/wangxin688/narvis/server/features/admin/schemas"
	"github.com/wangxin688/narvis/server/features/alert/schemas"
	infra_biz "github.com/wangxin688/narvis/server/features/infra/biz"
	infra_sc "github.com/wangxin688/narvis/server/features/infra/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/global/constants"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/pkg/am"
	"github.com/wangxin688/narvis/server/tools"
	"github.com/wangxin688/narvis/server/tools/errors"
	"github.com/wangxin688/narvis/server/tools/helpers"
	"go.uber.org/zap"
	"gorm.io/datatypes"
)

type AlertService struct{}

func NewAlertService() *AlertService {
	return &AlertService{}
}

func (a *AlertService) CreateAlert(alert *schemas.AlertCreate) (*models.Alert, error) {
	dbAlert, err := a.GetAlertByEventId(alert.EventId)
	if err != nil {
		return nil, err
	}
	if dbAlert != nil {
		global.OrganizationId.Set(dbAlert.OrganizationId)
		alertStatus := schemas.GetStatus(alert.Status)
		dbAlert.Status = alertStatus
		if alertStatus == constants.AlertResolvedStatus {
			resolvedAt := time.Now().UTC()
			dbAlert.ResolvedAt = &resolvedAt
		}
	} else {
		alertPre, err := a.alertPreProcess(alert)
		if err != nil {
			return nil, err
		}
		dbAlert = &models.Alert{
			AlertName:         alertPre.AlertName,
			Labels:            datatypes.NewJSONSlice(a.labelToDbLabel(alert.Labels)),
			EventId:           alertPre.EventId,
			TriggerId:         alertPre.TriggerId,
			Status:            schemas.GetStatus(alert.Status),
			StartedAt:         alertPre.StartedAt,
			Severity:          alertPre.Severity,
			SiteId:            alertPre.SiteId,
			DeviceId:          alertPre.DeviceId,
			ApId:              alertPre.ApId,
			DeviceInterfaceId: alertPre.InterfaceId,
			DeviceRole:        alertPre.DeviceRole,
			OrganizationId:    alertPre.OrganizationId,
		}
		if dbAlert.Status == constants.AlertResolvedStatus {
			resolvedAt := time.Now().UTC()
			dbAlert.ResolvedAt = &resolvedAt
			core.Logger.Info("[createAlert]: received resolved event", zap.Any("alert", alert))
		}
	}
	dbAlert = a.silenceAlert(dbAlert)
	if err = gen.Alert.UnderlyingDB().Save(dbAlert).Error; err != nil {
		return nil, err
	}
	if !dbAlert.Suppressed {
		tools.BackgroundTask(func() {
			postAlert := a.AlertManagerMessage(dbAlert)
			err = am.NewAlertManager().CreateAlerts(postAlert)
			if err != nil {
				core.Logger.Error("[createAlert]post alert to alertmanager error", zap.Error(err))
			}
		})
	}
	return dbAlert, nil
}

func (a *AlertService) GetAlertByEventId(eventId string) (*models.Alert, error) {
	alerts, err := gen.Alert.Where(gen.Alert.EventId.Eq(eventId)).Find()
	if err != nil {
		return nil, err
	}
	if len(alerts) == 0 {
		return nil, nil
	}
	return alerts[0], nil
}

func (a *AlertService) alertPreProcess(alert *schemas.AlertCreate) (*schemas.AlertConcrete, error) {
	startedAt := alert.StartedAt
	acc := &schemas.AlertConcrete{
		AlertName: alert.AlertName,
		Labels:    alert.Labels,
		EventId:   alert.EventId,
		TriggerId: alert.TriggerId,
		Status:    schemas.GetStatus(alert.Status),
		StartedAt: *startedAt,
	}
	hostId := alert.GetUuidHostId()
	if strings.Contains(alert.HostId, "d_") {
		host, err := gen.Device.Where(gen.Device.Id.Eq(hostId)).First()
		if err != nil {
			return nil, errors.NewError(errors.CodeNotFound, errors.MsgNotFound, gen.Device.TableName(), "Id", hostId)
		}
		acc.OrganizationId = host.OrganizationId
		acc.SiteId = host.SiteId
		if lo.Contains(alerts.GetApAlertEnumNames(), alerts.AlertNameEnum(alert.AlertName)) {
			apId, err := a.getApInfo(alert, host.SiteId, host.OrganizationId)
			if err != nil {
				return nil, err
			}
			acc.ApId = &apId
			acc.Severity = a.severity(alert.AlertName, devicerole.WlanAP)
			*acc.DeviceRole = string(devicerole.WlanAP)
		} else if lo.Contains(alerts.GetInterfaceAlertEnumNames(), alerts.AlertNameEnum(alert.AlertName)) {
			interfaceId, err := a.getInterfaceInfo(alert, host.Id)
			if err != nil {
				return nil, err
			}
			acc.DeviceId = &host.Id
			acc.InterfaceId = &interfaceId
			acc.Severity = a.severity(alert.AlertName, devicerole.DeviceRoleEnum(host.DeviceRole))
			acc.DeviceRole = &host.DeviceRole
		} else {
			acc.DeviceId = &host.Id
			acc.Severity = a.severity(alert.AlertName, devicerole.DeviceRoleEnum(host.DeviceRole))
			acc.DeviceRole = &host.DeviceRole
		}
	} else if strings.Contains(alert.HostId, "c_") || strings.Contains(alert.HostId, "cd_") {
		circuit, err := gen.Circuit.Where(gen.Circuit.Id.Eq(hostId)).First()
		if err != nil {
			return nil, errors.NewError(errors.CodeNotFound, errors.MsgNotFound, gen.Circuit.TableName(), "Id", hostId)
		}
		acc.CircuitId = &circuit.Id
		acc.SiteId = circuit.SiteId
		acc.OrganizationId = circuit.OrganizationId
		if alert.AlertName == string(alerts.NodePingTimeout) {
			acc.Severity = string(alerts.SeverityDisaster)
		} else {
			acc.Severity = string(alerts.SeverityWarning)
		}
	}
	return acc, nil
}

func (a *AlertService) getApInfo(alert *schemas.AlertCreate, siteId string, OrgId string) (apId string, err error) {
	apName := ""
	for _, label := range alert.Labels {
		if label.Tag == "apName" {
			apName = label.Value
			break
		}
	}
	if apName != "" {
		ap, err := gen.AP.Select(gen.AP.Id, gen.AP.SiteId, gen.AP.OrganizationId).Where(
			gen.AP.Name.Eq(apName),
			gen.AP.SiteId.Eq(siteId),
			gen.AP.OrganizationId.Eq(OrgId),
		).First()
		if err != nil {
			return "", err
		}
		return ap.Id, nil
	}
	return "", errors.NewError(errors.CodeApNameTagMissing, errors.MsgApNameTagMissing)
}

func (a *AlertService) getInterfaceInfo(alert *schemas.AlertCreate, deviceId string) (string, error) {
	interfaceName := ""
	for _, label := range alert.Labels {
		if label.Tag == "interface" {
			interfaceName = label.Value
			break
		}
	}
	if interfaceName != "" {
		interface_, err := gen.DeviceInterface.Select(gen.DeviceInterface.Id).Where(
			gen.DeviceInterface.IfName.Eq(interfaceName),
			gen.DeviceInterface.DeviceId.Eq(deviceId),
		).First()
		if err != nil {
			return "", err
		}
		return interface_.Id, nil
	}
	return "", errors.NewError(errors.CodeInterfaceTagMissing, errors.MsgInterfaceTagMissing)
}

func (a *AlertService) severity(alertName string, deviceRole devicerole.DeviceRoleEnum) string {
	sev := alerts.GetAlertName(alertName).Severity
	if deviceRole == devicerole.CoreSwitch || deviceRole == devicerole.Firewall || deviceRole == devicerole.WanRouter {
		if sev == alerts.SeverityInfo {
			return string(alerts.SeverityWarning)
		}
		if sev == alerts.SeverityWarning {
			return string(alerts.SeverityCritical)
		}
		return string(sev)
	}
	return string(sev)
}

func (a *AlertService) labelToDbLabel(labels []*schemas.Label) datatypes.JSONSlice[models.Label] {

	dbLabels := make([]models.Label, 0)
	for _, label := range labels {
		dbLabels = append(dbLabels, models.Label{
			Tag:   label.Tag,
			Value: label.Value,
		})
	}
	return datatypes.NewJSONSlice(dbLabels)
}
func (a *AlertService) labelToSchemaLabel(labels datatypes.JSONSlice[models.Label]) []schemas.Label {

	schemaLabels := make([]schemas.Label, 0)
	for _, label := range labels {
		schemaLabels = append(schemaLabels, schemas.Label{
			Tag:   label.Tag,
			Value: label.Value,
		})
	}
	return schemaLabels
}

func (a *AlertService) silenceAlert(alert *models.Alert) *models.Alert {
	alert.Suppressed = false
	alert.MaintenanceId = nil

	activeMaintenances, err := NewMaintenanceService().GetActiveMaintenance()
	if err != nil || len(activeMaintenances) == 0 {
		return alert
	}

	for _, maintenance := range activeMaintenances {
		if a.conditionMatch(maintenance.Conditions, alert) {
			alert.Suppressed = true
			alert.MaintenanceId = &maintenance.Id
			break
		}
	}
	return alert
}

func (a *AlertService) conditionMatch(conditions []models.Condition, alert *models.Alert) bool {
	if len(conditions) == 0 {
		return false
	}
	matchAll := make([]bool, 0)
	for _, cond := range conditions {
		if helpers.HasStructField(alert, cond.Item) {
			if slices.Equal(cond.Value, []string{"*"}) {
				matchAll = append(matchAll, true)
			} else if value, ok := helpers.StructGetFieldValue(alert, cond.Item); ok {
				if slices.Contains(cond.Value, fmt.Sprintf("%v", value)) {
					matchAll = append(matchAll, true)
				}
			} else {
				matchAll = append(matchAll, false)
			}
		}
	}
	if len(matchAll) == 0 || slices.Contains(matchAll, false) {
		return false
	}

	return true
}

func (a *AlertService) AlertManagerMessage(alert *models.Alert) []*am.Alert {
	annotations := make(map[string]string, 0)
	for _, label := range alert.Labels {
		if label.Value != "" {
			annotations[label.Tag] = label.Value
		}
	}
	postAlert := &am.Alert{
		Labels: map[string]string{
			"alertName":      alert.AlertName,
			"siteId":         alert.SiteId,
			"organizationId": alert.OrganizationId,
			"eventId":        alert.EventId,
			"severity":       alert.Severity,
		},
		Annotations:  annotations,
		Status:       schemas.GetReverseStatus(alert.Status),
		GeneratorURL: fmt.Sprintf("%s/api/v2/alerts/%s", core.Settings.System.BaseUrl, alert.Id),
		StartsAt:     alert.StartedAt,
	}
	if alert.DeviceId != nil {
		postAlert.Labels["deviceId"] = *alert.DeviceId
	}
	if alert.DeviceInterfaceId != nil {
		postAlert.Labels["deviceInterfaceId"] = *alert.DeviceInterfaceId
	}
	if alert.CircuitId != nil {
		postAlert.Labels["circuitId"] = *alert.CircuitId
	}
	if alert.DeviceRole != nil {
		postAlert.Labels["deviceRole"] = string(*alert.DeviceRole)
	}
	return []*am.Alert{postAlert}
}

func (a *AlertService) ListAlerts(query schemas.AlertQuery) (int64, []*schemas.Alert, error) {
	res := make([]*schemas.Alert, 0)
	orgId := global.OrganizationId.Get()
	stmt := gen.Alert.Where(gen.Alert.OrganizationId.Eq(orgId))
	if query.SiteId != nil {
		stmt = stmt.Where(gen.Alert.SiteId.In(*query.SiteId...))
	}
	if query.AlertName != nil {
		stmt = stmt.Where(gen.Alert.AlertName.In(*query.AlertName...))
	}
	if query.DeviceId != nil {
		stmt = stmt.Where(gen.Alert.DeviceId.In(*query.DeviceId...))
	}
	if query.ApId != nil {
		stmt = stmt.Where(gen.Alert.ApId.In(*query.ApId...))
	}
	if query.CircuitId != nil {
		stmt = stmt.Where(gen.Alert.CircuitId.In(*query.CircuitId...))
	}
	if query.DeviceInterfaceId != nil {
		stmt = stmt.Where(gen.Alert.DeviceInterfaceId.In(*query.DeviceInterfaceId...))
	}
	if query.DeviceRole != nil {
		stmt = stmt.Where(gen.Alert.DeviceRole.In(*query.DeviceRole...))
	}
	if query.Status != nil {
		stmt = stmt.Where(gen.Alert.Status.Eq(*query.Status))
	}
	if query.Severity != nil {
		stmt = stmt.Where(gen.Alert.Severity.In(*query.Severity...))
	}
	if query.Acknowledged != nil {
		stmt = stmt.Where(gen.Alert.Acknowledged.Is(*query.Acknowledged))
	}
	if query.Suppressed != nil {
		stmt = stmt.Where(gen.Alert.Suppressed.Is(*query.Suppressed))
	}
	if query.StartedAtGte != nil {
		stmt = stmt.Where(gen.Alert.StartedAt.Gte(*query.StartedAtGte))
	}
	if query.StartedAtLte != nil {
		stmt = stmt.Where(gen.Alert.StartedAt.Lte(*query.StartedAtLte))
	}
	if query.ResolvedAtGte != nil {
		stmt = stmt.Where(gen.Alert.ResolvedAt.Gte(*query.ResolvedAtGte))
	}
	if query.ResolvedAtLte != nil {
		stmt = stmt.Where(gen.Alert.ResolvedAt.Lte(*query.ResolvedAtLte))
	}
	if query.Keyword != nil {
		deviceIds, apIds, err := a.GetEntityIdBySearch(*query.Keyword, orgId)
		if err != nil {
			return 0, res, err
		}
		if len(deviceIds) == 0 && len(apIds) == 0 {
			return 0, res, nil
		}
		if len(deviceIds) > 0 && len(apIds) == 0 {
			stmt = stmt.Where(gen.Alert.DeviceId.In(deviceIds...))
		} else if len(deviceIds) == 0 && len(apIds) > 0 {
			stmt = stmt.Where(gen.Alert.ApId.In(apIds...))
		} else {
			stmt = stmt.Where(gen.Alert.DeviceId.In(deviceIds...)).Or(gen.Alert.ApId.In(apIds...))
		}
	}
	total, err := stmt.Count()
	if err != nil {
		return 0, res, err
	}
	defaultOrderBy := "startedAt"
	replacedOrderBy := "createdAt"
	if query.PageInfo.OrderBy == nil || *query.PageInfo.OrderBy == replacedOrderBy {
		query.PageInfo.OrderBy = &defaultOrderBy	
	}
	stmt.UnderlyingDB().Scopes(query.OrderByField())
	stmt.UnderlyingDB().Scopes(query.Pagination())
	list, err := stmt.Find()
	if err != nil {
		return 0, res, err
	}
	res, err = a.formatAlert(list)
	if err != nil {
		return 0, res, err
	}
	return total, res, err

}

func (a *AlertService) formatAlert(alert []*models.Alert) ([]*schemas.Alert, error) {
	siteIds := make([]string, 0)
	deviceIds := make([]string, 0)
	circuitIds := make([]string, 0)
	apIds := make([]string, 0)
	results := make([]*schemas.Alert, 0)
	alertIds := make([]string, 0)
	for _, item := range alert {
		siteIds = append(siteIds, item.SiteId)
		alertIds = append(alertIds, item.Id)
		if item.DeviceId != nil {
			deviceIds = append(deviceIds, *item.DeviceId)
		}
		if item.CircuitId != nil {
			circuitIds = append(circuitIds, *item.CircuitId)
		}
		if item.ApId != nil {
			apIds = append(apIds, *item.ApId)
		}
	}
	siteMap, err := infra_biz.NewSiteService().GetSiteShortMap(lo.Uniq(siteIds))
	deviceMap := make(map[string]*infra_sc.DeviceShort)
	circuitMap := make(map[string]*infra_sc.CircuitShort)
	apMap := make(map[string]*infra_sc.APShort)
	if err != nil {
		core.Logger.Error("[alertQuery]: get site short map error", zap.Error(err))
		return results, err
	}
	if len(deviceIds) > 0 {
		deviceMap, err = infra_biz.NewDeviceService().GetDeviceShortMap(lo.Uniq(deviceIds))
		if err != nil {
			core.Logger.Error("[alertQuery]: get device short map error", zap.Error(err))
			return results, err
		}
	}
	if len(circuitIds) > 0 {
		circuitMap, err = infra_biz.NewCircuitService().GetCircuitShortMap(lo.Uniq(circuitIds))
		if err != nil {
			core.Logger.Error("[alertQuery]: get circuit short map error", zap.Error(err))
			return results, err
		}
	}
	if len(apIds) > 0 {
		apMap, err = infra_biz.NewApService().GetApShortMap(apIds)
		if err != nil {
			core.Logger.Error("[alertQuery]: get ap short map error", zap.Error(err))
			return results, err
		}
	}
	actionLogCount, err := a.GetActionLogCountByAlertIds(alertIds)
	if err != nil {
		return results, err
	}
	for _, item := range alert {
		as := &schemas.Alert{
			Id:         item.Id,
			AlertName:  alerts.GetAlertName(item.AlertName),
			Labels:     a.labelToSchemaLabel(item.Labels),
			Status:     item.Status,
			StartedAt:  item.StartedAt,
			ResolvedAt: item.ResolvedAt,
			Severity:   item.Severity,
			Site: infra_sc.SiteShort{
				Id:       item.SiteId,
				Name:     siteMap[item.SiteId].Name,
				SiteCode: siteMap[item.SiteId].SiteCode,
			},
			DeviceRole:     item.DeviceRole,
			ActionLogCount: actionLogCount[item.Id],
		}
		as.Duration = as.GetDuration()
		if item.DeviceId != nil {
			as.Entity = schemas.Entity{
				Id:   *item.DeviceId,
				Name: deviceMap[*item.DeviceId].Name,
				Type: "device",
			}
		}
		if item.CircuitId != nil {
			as.Entity = schemas.Entity{
				Id:   *item.CircuitId,
				Name: circuitMap[*item.CircuitId].Name,
				Type: "circuit",
			}
		}
		if item.ApId != nil {
			as.Entity = schemas.Entity{
				Id:   *item.ApId,
				Name: apMap[*item.ApId].Name,
				Type: "ap",
			}
		}
		results = append(results, as)
	}
	return results, nil
}

func (a *AlertService) GetAlertCountByInterfaceIds(interfaceIds []string) (map[string]int, error) {
	var alerts []struct {
		InterfaceId string
		Count       int
	}
	err := gen.Alert.Select(gen.Alert.DeviceInterfaceId, gen.Alert.Id.Count().As("count")).
		Where(gen.Alert.DeviceInterfaceId.In(interfaceIds...), gen.Alert.Status.Eq(constants.AlertFiringStatus)).
		Group(gen.Alert.DeviceInterfaceId).Scan(&alerts)
	if err != nil {
		return nil, err
	}
	res := make(map[string]int)
	for _, item := range alerts {
		res[item.InterfaceId] = item.Count
	}
	return res, nil
}

func (a *AlertService) GetAlertCountByDeviceIds(deviceIds []string) (map[string]int, error) {
	var alerts []struct {
		DeviceId string
		Count    int
	}
	err := gen.Alert.Select(gen.Alert.DeviceId, gen.Alert.Id.Count().As("count")).
		Where(gen.Alert.DeviceId.In(deviceIds...), gen.Alert.Status.Eq(constants.AlertFiringStatus)).
		Group(gen.Alert.DeviceId).Scan(&alerts)
	if err != nil {
		return nil, err
	}
	res := make(map[string]int)
	for _, item := range alerts {
		res[item.DeviceId] = item.Count
	}
	return res, nil
}

func (a *AlertService) GetAlertCountByCircuitIds(circuitIds []string) (map[string]int, error) {
	var alerts []struct {
		CircuitId string
		Count     int
	}
	err := gen.Alert.Select(gen.Alert.CircuitId, gen.Alert.Id.Count().As("count")).
		Where(gen.Alert.CircuitId.In(circuitIds...), gen.Alert.Status.Eq(constants.AlertFiringStatus)).
		Group(gen.Alert.CircuitId).Scan(&alerts)
	if err != nil {
		return nil, err
	}
	res := make(map[string]int)
	for _, item := range alerts {
		res[item.CircuitId] = item.Count
	}
	return res, nil
}

func (a *AlertService) GetAlertCountByApIds(apIds []string) (map[string]int, error) {
	var alerts []struct {
		ApId  string
		Count int
	}
	err := gen.Alert.Select(gen.Alert.ApId, gen.Alert.Id.Count().As("count")).
		Where(gen.Alert.ApId.In(apIds...), gen.Alert.Status.Eq(constants.AlertFiringStatus)).
		Group(gen.Alert.ApId).Scan(&alerts)
	if err != nil {
		return nil, err
	}
	res := make(map[string]int)
	for _, item := range alerts {
		res[item.ApId] = item.Count
	}
	return res, nil
}

func (a *AlertService) GetActionLogCountByAlertIds(alertIds []string) (map[string]int, error) {
	var results []struct {
		AlertId string
		Count   int
	}
	err := gen.AlertActionLog.Select(gen.AlertActionLog.AlertId, gen.AlertActionLog.Id.Count().As("count")).
		Where(gen.AlertActionLog.AlertId.In(alertIds...)).
		Group(gen.AlertActionLog.AlertId).Scan(&results)
	if err != nil {
		return nil, err
	}
	res := make(map[string]int)
	for _, item := range results {
		res[item.AlertId] = item.Count
	}
	return res, nil

}

func (a *AlertService) GetEntityIdBySearch(keyword string, orgId string) (deviceIds []string, apIds []string, err error) {
	deviceIds, err = infra_biz.NewDeviceService().SearchDeviceByKeyword(keyword, orgId)
	if err != nil {
		return nil, nil, err
	}
	apIds, err = infra_biz.NewApService().SearchApByKeyword(keyword, orgId)
	if err != nil {
		return nil, nil, err
	}
	return deviceIds, apIds, nil
}

func (a *AlertService) GetById(id string) (*schemas.AlertDetail, error) {
	alert, err := gen.Alert.Where(
		gen.Alert.Id.Eq(id),
		gen.Alert.OrganizationId.Eq(global.OrganizationId.Get()),
	).Preload(gen.Alert.Device).Preload(gen.Alert.Circuit).
		Preload(gen.Alert.Ap).Preload(gen.Alert.Site).Preload(gen.Alert.RootCause).
		Preload(gen.Alert.ActionLog.AssignUser).First()

	if err != nil {
		return nil, err
	}
	actionLogs, err := gen.AlertActionLog.Where(gen.AlertActionLog.AlertId.Eq(id)).
		Preload(gen.AlertActionLog.RootCause).
		Preload(gen.AlertActionLog.CreatedBy).Preload(gen.AlertActionLog.AssignUser).Find()

	if err != nil {
		return nil, err
	}
	am := &schemas.AlertDetail{
		Status:       alert.Status,
		StartedAt:    alert.StartedAt,
		ResolvedAt:   alert.ResolvedAt,
		Acknowledged: alert.Acknowledged,
		Suppressed:   alert.Suppressed,
		Severity:     alert.Severity,
		Id:           alert.Id,
		AlertName:    alerts.GetAlertName(alert.AlertName),
		Labels:       a.labelToSchemaLabel(alert.Labels),
		Site: infra_sc.SiteShort{
			Id:       alert.Site.Id,
			Name:     alert.Site.Name,
			SiteCode: alert.Site.SiteCode},
		DeviceRole: alert.DeviceRole,
	}
	am.Duration = am.GetDuration()
	if alert.RootCauseId != nil {
		am.RootCause = &schemas.RootCause{
			Id:          alert.RootCause.Id,
			Name:        alert.RootCause.Name,
			Description: alert.RootCause.Description,
			Category:    alert.RootCause.Category,
		}
	}
	if alert.UserId != nil {
		am.User = &admin_sc.UserShort{
			Id:       alert.User.Id,
			Username: alert.User.Username,
			Email:    alert.User.Email,
			Avatar:   alert.User.Avatar,
		}
	}
	if alert.ActionLog != nil {
		am.ActionLog = make([]*schemas.ActionLog, 0)
		for _, log := range actionLogs {
			acLog := &schemas.ActionLog{
				Id:           log.Id,
				Resolved:     log.Resolved,
				CreatedAt:    log.CreatedAt,
				Acknowledged: log.Acknowledged,
				Suppressed:   log.Suppressed,
				Comment:      log.Comment,
				CreatedBy: &admin_sc.UserShort{
					Id:       log.CreatedBy.Id,
					Username: log.CreatedBy.Username,
					Email:    log.CreatedBy.Email,
					Avatar:   log.CreatedBy.Avatar,
				},
			}
			if log.RootCauseId != nil {
				acLog.RootCause = &schemas.RootCause{
					Id:          log.RootCause.Id,
					Name:        log.RootCause.Name,
					Description: log.RootCause.Description,
					Category:    log.RootCause.Category,
				}
			}
			if log.AssignUserId != nil {
				acLog.AssignUser = &admin_sc.UserShort{
					Id:       *log.AssignUserId,
					Username: log.AssignUser.Username,
					Email:    log.AssignUser.Email,
					Avatar:   log.AssignUser.Avatar,
				}
			}
			acLog.GenerateActions()
			am.ActionLog = append(am.ActionLog, acLog)
		}
	}

	return am, nil
}
