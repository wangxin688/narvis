package alert_biz

import (
	"strings"
	"time"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/intend/alerts"
	"github.com/wangxin688/narvis/intend/devicerole"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/alert/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/global/constants"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/tools/errors"
	"go.uber.org/zap"
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
			Labels:            a.labelToDbLabel(alert.Labels),
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
			OrganizationId:    global.OrganizationId.Get(),
		}
		if dbAlert.Status == constants.AlertResolvedStatus {
			resolvedAt := time.Now().UTC()
			dbAlert.ResolvedAt = &resolvedAt
			core.Logger.Info("[createAlert]: received resolved event", zap.Any("alert", dbAlert))
		}
	}
	// TODO: complete maintenance status checking here

	if err = gen.Alert.Create(dbAlert); err != nil {
		return nil, err
	}
	return dbAlert, nil
}

func (a *AlertService) GetAlertByEventId(eventId string) (*models.Alert, error) {
	return gen.Alert.Where(gen.Alert.EventId.Eq(eventId)).First()
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
		if lo.Contains(alerts.GetApAlertEnumNames(), alerts.AlertNameEnum(alert.AlertName)) {
			apId, err := a.getApInfo(alert, host.SiteId, host.OrganizationId)
			if err != nil {
				return nil, err
			}
			acc.ApId = &apId
			acc.SiteId = host.SiteId
			acc.Severity = a.severity(alert.AlertName, devicerole.WlanAP)
		} else if lo.Contains(alerts.GetInterfaceAlertEnumNames(), alerts.AlertNameEnum(alert.AlertName)) {
			interfaceId, err := a.getInterfaceInfo(alert, host.Id)
			if err != nil {
				return nil, err
			}
			acc.SiteId = host.SiteId
			acc.DeviceId = &host.Id
			acc.InterfaceId = &interfaceId
			acc.Severity = a.severity(alert.AlertName, devicerole.DeviceRoleEnum(host.DeviceRole))
		}
	} else if strings.Contains(alert.HostId, "c_") || strings.Contains(alert.HostId, "cd_") {
		circuit, err := gen.Circuit.Where(gen.Circuit.Id.Eq(hostId)).First()
		if err != nil {
			return nil, errors.NewError(errors.CodeNotFound, errors.MsgNotFound, gen.Circuit.TableName(), "Id", hostId)
		}
		acc.CircuitId = &circuit.Id
		acc.SiteId = circuit.SiteId
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
		ap, err := gen.AP.Select(gen.AP.Id, gen.AP.SiteId).Where(
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

func (a *AlertService) labelToDbLabel(labels []*schemas.Label) []models.Label {

	dbLabels := make([]models.Label, 0)
	for _, label := range labels {
		dbLabels = append(dbLabels, models.Label{
			Tag:   label.Tag,
			Value: label.Value,
		})
	}
	return dbLabels
}

// func (a *AlertService) silenceAlert(alert *models.Alert) *models.Alert {
// 	ms, err := NewMaintenanceService().GetActiveMaintenance()
// 	if err != nil || len(ms) == 0 {
// 		return alert
// 	}

// 	for _, m := range ms {

// 	}
// }

// func (a *AlertService) conditionMatch(conditions []models.Condition, alert *models.Alert) bool {
// 	if len(conditions) == 0 {
// 		return false
// 	}
// 	matchAll := make([]bool, 0)
// 	for _, cond := range conditions {

// 	}
// }
