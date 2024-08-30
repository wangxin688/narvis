package schemas

import "time"

type Label struct {
	Tag   string `json:"tag" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type AlertCreate struct {
	AlertName string     `json:"alertName" binding:"required"`
	HostID    string     `json:"hostId" binding:"required,uuid"`
	Labels    []*Label   `json:"tags" binding:"omitempty"`
	EventID   string     `json:"eventId" binding:"required"`
	TriggerID string     `json:"triggerId" binding:"required"`
	Status    string     `json:"status" binding:"required,oneof: Problem OK"`
	StartedAt *time.Time `json:"startedAt" binding:"omitempty,datetime"`
	Severity  string     `json:"severity" binding:"required,oneof: P1 P2 P3 P4"`
}

func (a *AlertCreate) ValidateStartTime() bool {
	if a.StartedAt != nil {
		return !a.StartedAt.After(time.Now())
	}
	return true
}

func GetStatus(status string) uint8 {
	statusMap := map[string]uint8{
		"firing":   0,
		"Problem":  0,
		"OK":       1,
		"resolved": 1,
	}
	return statusMap[status]
}

type AlertConcrete struct {
	AlertCreate
	SiteID      string
	DeviceID    *string
	ApID        *string
	InterfaceID *string
}
