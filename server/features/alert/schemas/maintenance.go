package schemas

import "time"

type MaintenanceCreate struct {
	Name            string      `json:"name" binding:"required"`
	StartedAt       time.Time   `json:"startedAt" binding:"required,datetime"`
	EndedAt         time.Time   `json:"endedAt" binding:"required,datetime"`
	MaintenanceType string      `json:"maintenanceType" binding:"required"`
	Description     string      `json:"description" binding:"required"`
	Conditions      []Condition `json:"conditions" binding:"required"`
}

type MaintenanceUpdate struct {
	Name            *string      `json:"name" binding:"omitempty"`
	StartedAt       *time.Time   `json:"startedAt" binding:"omitempty,datetime"`
	EndedAt         *time.Time   `json:"endedAt" binding:"omitempty,datetime"`
	MaintenanceType *string      `json:"maintenanceType" binding:"omitempty"`
	Description     *string      `json:"description" binding:"omitempty"`
	Conditions      *[]Condition `json:"conditions" binding:"omitempty"`
}
