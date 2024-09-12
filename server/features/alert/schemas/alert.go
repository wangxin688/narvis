package schemas

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/wangxin688/narvis/intend/alerts"
	"github.com/wangxin688/narvis/server/global/constants"
	"github.com/wangxin688/narvis/server/tools/errors"
)

type Label struct {
	Tag   string `json:"tag" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type AlertCreate struct {
	AlertName string     `json:"alertName" binding:"required"`
	HostId    string     `json:"hostId" binding:"required,string"`
	Labels    []*Label   `json:"tags" binding:"omitempty"`
	EventId   string     `json:"eventId" binding:"required"`
	TriggerId string     `json:"triggerId" binding:"required"`
	Status    string     `json:"status" binding:"required,oneof: Problem OK"`
	StartedAt *time.Time `json:"startedAt" binding:"omitempty,datetime"`
	Severity  string     `json:"severity" binding:"required,oneof: P1 P2 P3 P4"`
}

func (a *AlertCreate) validateStartTime() error {
	if a.StartedAt != nil {
		if a.StartedAt.After(time.Now()) {
			errors.NewError(errors.CodeAlertStartTimeInFuture, errors.MsgAlertStartTimeInFuture)
		}
	}
	return nil
}

func (a *AlertCreate) validateAlertName() error {
	alertNames := alerts.GetAlertEnumNames()
	if !lo.Contains(alertNames, alerts.AlertNameEnum(a.AlertName)) {
		return errors.NewError(errors.CodeAlertNameNotDefined, errors.MsgAlertNameNotDefined, a.AlertName)
	}
	return nil
}

func (a *AlertCreate) ValidateHostId() error {
	if !strings.Contains(a.HostId, "d_") || !strings.Contains(a.HostId, "c_") || !strings.Contains(a.HostId, "cd_") {
		return errors.NewError(errors.CodeAlertHostIdInvalid, errors.MsgAlertHostIdInvalid, a.HostId)
	}
	splitString := strings.Split(a.HostId, "_")[1]
	if _, err := uuid.Parse(splitString); err != nil {
		return errors.NewError(errors.CodeAlertHostIdInvalid, errors.MsgAlertHostIdInvalid, a.HostId)
	}
	return nil
}

func (a *AlertCreate) GetUuidHostId() string {
	return strings.Split(a.HostId, "_")[1]
}

func (a *AlertCreate) Validate() error {
	if err := a.validateAlertName(); err != nil {
		return err
	}
	if err := a.validateStartTime(); err != nil {
		return err
	}
	a.updateLabels()
	return nil
}

// remove duplicate and no-need labels
func (a *AlertCreate) updateLabels() {
	excludeKeys := []string{"siteCode", "name", "scope", "hostname"}
	existedKeys := make([]string, 0)
	newLabels := make([]*Label, 0)
	for _, label := range a.Labels {
		if !lo.Contains(existedKeys, label.Tag) && !lo.Contains(excludeKeys, label.Tag) {
			existedKeys = append(existedKeys, label.Tag)
			newLabels = append(newLabels, label)
		}
	}
	a.Labels = newLabels
}

// get status integer from status string
func GetStatus(status string) uint8 {
	statusMap := map[string]uint8{
		"firing":   constants.AlertFiringStatus,
		"Problem":  constants.AlertFiringStatus,
		"OK":       constants.AlertResolvedStatus,
		"resolved": constants.AlertResolvedStatus,
	}
	return statusMap[status]
}

type AlertConcrete struct {
	AlertName      string
	Labels         []*Label
	EventId        string
	TriggerId      string
	Status         uint8
	StartedAt      time.Time
	Severity       string
	SiteId         string
	DeviceId       *string
	CircuitId      *string
	ApId           *string
	InterfaceId    *string
	DeviceRole     *string
	OrganizationId string
}
