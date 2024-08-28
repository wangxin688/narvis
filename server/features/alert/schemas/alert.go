package schemas

import "time"

type Label struct {
	Tag   string `json:"tag" binding:"required"`
	Value string `json:"value" binding:"required"`
}

type AlertCreate struct {
	AlertName string     `json:"alert_name" binding:"required"`
	HostID    string     `json:"host_id" binding:"required,uuid"`
	Labels    []*Label   `json:"tags" binding:"omitempty"`
	EventID   string     `json:"event_id" binding:"required"`
	TriggerID string     `json:"trigger_id" binding:"required"`
	Status    string     `json:"status" binding:"required,oneof: Problem OK"`
	StartedAt *time.Time `json:"started_at" binding:"omitempty,datetime"`
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
	SiteID string 
	DeviceID *string 
	ApID *string
	InterfaceID *string
	
}
