package alert_biz

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/alert/dispatcher"
	"github.com/wangxin688/narvis/server/features/alert/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/global/constants"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/pkg/am"
	"github.com/wangxin688/narvis/server/tools"
	ts "github.com/wangxin688/narvis/server/tools/errors"
	"github.com/wangxin688/narvis/server/tools/helpers"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AlertGroupService struct{}

func NewAlertGroupService() *AlertGroupService {
	return &AlertGroupService{}
}

func (ag *AlertGroupService) CreateAlertGroup(group *schemas.AlertGroupCreate) (*models.AlertGroup, error) {
	groupKey := helpers.StringToMd5(group.GroupKey)
	contentByte, _ := json.Marshal(group.GetAlertLabels())
	contentHash := helpers.ByteToMd5(contentByte)
	orgId, ok := group.GroupLabels["organization_id"]
	if !ok {
		return nil, ts.NewError(ts.CodeAlertGroupMissingOrganizationId, ts.MsgAlertGroupMissingOrganizationId)
	}
	global.OrganizationId.Set(orgId)
	dbGroup, err := ag.GetAlertGroupByGroupKey(groupKey)
	if err != nil {
		return nil, err
	}
	var sendFlag = false
	events, groupEventIds, err := ag.GetGroupedAlerts(group)
	if err != nil {
		return nil, err
	}
	if dbGroup == nil {
		if len(groupEventIds) > 0 || (len(groupEventIds) == 0 && group.Status == "firing") {
			siteId := group.GroupLabels["siteId"]
			sendFlag = true
			core.Logger.Info(fmt.Sprintf("[alertGroup]: receive new alert group with hash: %s with %d alerts", groupKey, len(groupEventIds)))
			dbGroup = &models.AlertGroup{
				GroupKey:       groupKey,
				Status:         schemas.GetStatus(group.Status),
				StartedAt:      time.Now().UTC(),
				Acknowledged:   false,
				AlertName:      group.GroupLabels["alertName"],
				HashKey:        contentHash,
				SiteId:         siteId,
				OrganizationId: orgId,
				Severity:       ag.GetSeverityFromEvents(events),
			}
			err = gen.AlertGroup.Create(dbGroup)
			if err != nil {
				core.Logger.Error(fmt.Sprintf("[alertGroup]: create alert group error: %s", err.Error()))
				return nil, err
			}
			ag.UpdateRelatedAlertByEvents(events, dbGroup.Id)
		} else {
			core.Logger.Warn("[alertGroup]: ignore alert group", zap.Any("alertGroup", group))
			core.Logger.Warn(fmt.Sprintf("[alertGroup]: ignore alert group with hash: %s with %d alerts", groupKey, len(groupEventIds)))
			core.Logger.Warn("[alertGroup]: ignore alert group because of flapping resolved alerts")
		}
	} else {
		if dbGroup.HashKey != contentHash {
			severity := ag.GetSeverityFromEvents(events)
			ag.UpdateRelatedAlertByEvents(events, dbGroup.Id)
			sendFlag = true
			dbGroup.Severity = severity
		}
		dbGroup.HashKey = contentHash
		if group.Status == "resolved" {
			activeEventIds, err := ag.GetActiveEventIds(dbGroup.Id)
			if err != nil {
				core.Logger.Error(fmt.Sprintf("[alertGroup]: get active event ids error: %s", err.Error()))
				return nil, err
			}
			if len(activeEventIds) > 0 {
				dbGroup.Status = constants.AlertFiringStatus
				core.Logger.Warn("[alertGroup]: un-compatible alert group and alert status")
				core.Logger.Warn("[alertGroup]: received resolved alert group but conflict with active alerts", zap.Any("alertGroup", group))
			} else {
				dbGroup.Status = constants.AlertResolvedStatus
				recoveryTime := time.Now().UTC()
				dbGroup.ResolvedAt = &recoveryTime
				core.Logger.Info("[alertGroup]: receive resolved alert group and all related alerts are resolved", zap.Any("alertGroup", group))
			}

		}
		err := gen.AlertGroup.UnderlyingDB().Save(dbGroup).Error
		if err != nil {
			core.Logger.Error(fmt.Sprintf("[alertGroup]: update alert group error: %s", err.Error()))
			return nil, err
		}
	}
	if !dbGroup.Suppressed && !dbGroup.Acknowledged {
		tools.BackgroundTask(func() {
			disp := dispatcher.NewDispatcher(dbGroup.Id, sendFlag)
			disp.Dispatch()
		})
	}
	return dbGroup, nil
}

func (ag *AlertGroupService) GetAlertGroupByGroupKey(groupKey string) (*models.AlertGroup, error) {
	result, err := gen.AlertGroup.Where(gen.AlertGroup.GroupKey.Eq(groupKey)).First()
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return result, err
}

func (ag *AlertGroupService) GetGroupedAlerts(group *schemas.AlertGroupCreate) (events []*schemas.Event, groupEventIds []string, err error) {
	filters := make([]string, 0)
	for key, value := range group.GroupLabels {
		filters = append(filters, fmt.Sprintf("%s = '%s'", key, value))
	}
	query := &am.AlertRequest{
		Filter: filters,
	}
	amGroupAlerts, err := am.NewAlertManager().GetAlertGroups(query)
	if err != nil {
		return nil, nil, err
	}
	if len(amGroupAlerts) == 0 {
		return nil, nil, nil
	}
	for _, ag := range amGroupAlerts {
		for _, al := range ag.Alerts {
			event := &schemas.Event{
				EventId:   al.Labels["eventId"],
				IsActive:  false,
				Inhibited: false,
				Severity:  al.Labels["severity"],
			}
			if al.Status.InhibitedBy != nil{
				event.Inhibited = true
			}
			if al.Status.State == "active" {
				event.IsActive = true
				groupEventIds = append(groupEventIds, al.Labels["eventId"])
			}
			events = append(events, event)
		}
	}
	return events, groupEventIds, nil
}

func (ag *AlertGroupService) GetSeverityFromEvents(events []*schemas.Event) string {
	list := make([]int, 0)
	for _, event := range events {
		level, _ := strconv.Atoi(strings.TrimPrefix(event.Severity, "P"))
		list = append(list, level)
	}
	minLevel := lo.Min(list)
	return "P" + strconv.Itoa(minLevel)
}

func (ag *AlertGroupService) UpdateRelatedAlertByEvents(events []*schemas.Event, alertGroupId string) {
	eventIds := make([]string, 0)
	for _, event := range events {
		eventIds = append(eventIds, event.EventId)
	}
	_events := make(map[string]*schemas.Event)
	for _, event := range events {
		_events[event.EventId] = event
	}
	gen.Alert.UnderlyingDB().Begin()
	alerts, err := gen.Alert.Where(gen.Alert.EventId.In(eventIds...)).Find()
	if err != nil {
		return
	}
	for _, alert := range alerts {
		if _events[alert.EventId].IsActive {
			alert.AlertGroupId = &alertGroupId
		}
		if _events[alert.EventId].Inhibited {
			alert.Inhibited = true
		} else {
			alert.Inhibited = false
		}
		gen.Alert.UnderlyingDB().Save(alert)
	}
	gen.Alert.UnderlyingDB().Commit()
}

func (ag *AlertGroupService) GetActiveEventIds(alertGroupId string) ([]string, error) {
	alerts, err := gen.Alert.Select(gen.Alert.EventId).Where(
		gen.Alert.AlertGroupId.Eq(alertGroupId),
		gen.Alert.Status.Eq(constants.AlertFiringStatus)).Find()
	if err != nil {
		return nil, err
	}
	eventIds := make([]string, 0)
	for _, alert := range alerts {
		eventIds = append(eventIds, alert.EventId)
	}
	return eventIds, nil
}
