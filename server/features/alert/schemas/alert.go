package schemas

import (
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/samber/lo"
	"github.com/wangxin688/narvis/intend/alerts"
	as "github.com/wangxin688/narvis/server/features/admin/schemas"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	"github.com/wangxin688/narvis/server/global/constants"
	"github.com/wangxin688/narvis/server/tools/errors"
	"github.com/wangxin688/narvis/server/tools/helpers"
	ts "github.com/wangxin688/narvis/server/tools/schemas"
)

type Label struct {
	Tag   string `json:"tag" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type AlertCreate struct {
	AlertName string     `json:"alertName" binding:"required"`
	HostId    string     `json:"hostId" binding:"required"`
	Labels    []*Label   `json:"tags" binding:"omitempty"`
	EventId   string     `json:"eventId" binding:"required"`
	TriggerId string     `json:"triggerId" binding:"required"`
	Status    string     `json:"status" binding:"required,oneof=PROBLEM OK"`
	StartedAt *time.Time `json:"startedAt" binding:"omitempty,datetime"`
	Severity  string     `json:"severity" binding:"omitempty,oneof=P1 P2 P3 P4"`
}

func (a *AlertCreate) validateStartTime() error {
	if a.StartedAt != nil {
		if a.StartedAt.After(time.Now()) {
			return errors.NewError(errors.CodeAlertStartTimeInFuture, errors.MsgAlertStartTimeInFuture)
		}
	}
	if a.StartedAt == nil {
		timeNow := time.Now()
		a.StartedAt = &timeNow
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
	excludeKeys := []string{"siteCode", "name", "scope", "hostname", "organizationId"}
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
		"PROBLEM":  constants.AlertFiringStatus,
		"OK":       constants.AlertResolvedStatus,
		"resolved": constants.AlertResolvedStatus,
	}
	return statusMap[status]
}

func GetReverseStatus(status uint8) string {
	statusMap := map[uint8]string{
		constants.AlertFiringStatus:   "firing",
		constants.AlertResolvedStatus: "resolved",
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
	ServerId       *string
	InterfaceId    *string
	DeviceRole     *string
	OrganizationId string
}

type Alert struct {
	Status         uint8             `json:"status"`
	StartedAt      time.Time         `json:"startedAt"`
	ResolvedAt     *time.Time        `json:"resolvedAt"`
	Duration       string            `json:"duration"`
	Acknowledged   bool              `json:"acknowledged"`
	Suppressed     bool              `json:"suppressed"`
	Severity       string            `json:"severity"`
	Id             string            `json:"id"`
	AlertName      alerts.AlertName  `json:"alertName"`
	Site           schemas.SiteShort `json:"site"`
	Entity         Entity            `json:"entity"`
	Labels         []Label           `json:"labels"`
	DeviceRole     *string           `json:"deviceRole"`
	User           *as.UserShort     `json:"user"`
	ActionLogCount int               `json:"actionLogCount"`
}

func (a *Alert) GetDuration() string {
	var duration string
	if a.ResolvedAt != nil {
		duration = helpers.HumanReadableDuration(int64(a.ResolvedAt.Sub(a.StartedAt).Seconds()))
	} else {
		duration = helpers.HumanReadableDuration(int64(time.Since(a.StartedAt).Seconds()))
	}
	return duration
}

type AlertDetail struct {
	Status       uint8             `json:"status"`
	StartedAt    time.Time         `json:"startedAt"`
	ResolvedAt   *time.Time        `json:"resolvedAt"`
	Duration     string            `json:"duration"`
	Acknowledged bool              `json:"acknowledged"`
	Suppressed   bool              `json:"suppressed"`
	Severity     string            `json:"severity"`
	Id           string            `json:"id"`
	AlertName    alerts.AlertName  `json:"alertName"`
	Site         schemas.SiteShort `json:"site"`
	Entity       Entity            `json:"entity"`
	Labels       []Label           `json:"labels"`
	DeviceRole   *string           `json:"deviceRole"`
	RootCause    *RootCause        `json:"rootCause"`
	User         *as.UserShort     `json:"user"`
	ActionLog    []*ActionLog      `json:"actionLog"`
}

func (a *AlertDetail) GetDuration() string {
	var duration string
	if a.ResolvedAt != nil {
		duration = helpers.HumanReadableDuration(int64(a.ResolvedAt.Sub(a.StartedAt).Seconds()))
	} else {
		duration = helpers.HumanReadableDuration(int64(time.Since(a.StartedAt).Seconds()))
	}
	return duration
}

type Entity struct {
	Id   string `json:"Id"`
	Name string `json:"Name"`
	Type string `json:"Type"`
}

type AlertQuery struct {
	ts.PageInfo
	SiteId            *[]string  `form:"siteId" binding:"omitempty,list_uuid"`
	AlertName         *[]string  `form:"alertName" binding:"omitempty"`
	DeviceId          *[]string  `form:"deviceId" binding:"omitempty,list_uuid"`
	ApId              *[]string  `form:"apId" binding:"omitempty,list_uuid"`
	ServerId          *[]string  `form:"serverId" binding:"omitempty,list_uuid"`
	CircuitId         *[]string  `form:"circuitId" binding:"omitempty,list_uuid"`
	DeviceInterfaceId *[]string  `form:"deviceInterfaceId" binding:"omitempty,list_uuid"`
	DeviceRole        *[]string  `form:"deviceRole" binding:"omitempty"`
	Severity          *[]string  `form:"severity" binding:"omitempty"`
	Status            *uint8     `form:"status" binding:"omitempty"`
	Acknowledged      *bool      `form:"acknowledged" binding:"omitempty"`
	Suppressed        *bool      `form:"suppressed" binding:"omitempty"`
	StartedAtGte      *time.Time `form:"startedAtGte" binding:"omitempty"`
	StartedAtLte      *time.Time `form:"startedAtLte" binding:"omitempty"`
	ResolvedAtGte     *time.Time `form:"endsAtGte" binding:"omitempty"`
	ResolvedAtLte     *time.Time `form:"endsAtLte" binding:"omitempty"`
}
