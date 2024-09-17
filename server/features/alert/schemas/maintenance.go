package schemas

import (
	"time"

	"github.com/wangxin688/narvis/server/features/admin/schemas"
	ts "github.com/wangxin688/narvis/server/tools/schemas"
)

type MaintenanceCreate struct {
	Name            string      `json:"name" binding:"required"`
	StartedAt       time.Time   `json:"startedAt" binding:"required,datetime"`
	EndedAt         time.Time   `json:"endedAt" binding:"required,datetime"`
	MaintenanceType string      `json:"maintenanceType" binding:"required"`
	Description     *string     `json:"description" binding:"omitempty"`
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

type Maintenance struct {
	Id              string             `json:"id"`
	CreatedAt       time.Time          `json:"createdAt"`
	UpdatedAt       *time.Time         `json:"updatedAt"`
	Name            string             `json:"name"`
	StartedAt       time.Time          `json:"startedAt"`
	EndedAt         time.Time          `json:"endedAt"`
	MaintenanceType string             `json:"maintenanceType"`
	Description     *string            `json:"description"`
	Conditions      []Condition        `json:"conditions"`
	CreatedBy       schemas.UserShort  `json:"createdBy"`
	UpdatedBy       *schemas.UserShort `json:"updatedBy"`
}

type MaintenanceQuery struct {
	ts.PageInfo
	Id              *[]string `json:"id" binding:"omitempty,list_uuid"`
	MaintenanceType *string   `json:"maintenanceType" binding:"omitempty"`
	Status          *string   `json:"status" binding:"omitempty,oneof=Approaching Active Expired"`
	SiteId          *[]string `json:"siteId" binding:"omitempty,list_uuid"`
}
