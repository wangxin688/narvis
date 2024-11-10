package schemas

import (
	"time"

	"github.com/wangxin688/narvis/server/tools/schemas"
)

type ServerCreate struct {
	Name         string   `json:"name" binding:"required"`
	ManagementIp string   `json:"managementIp" binding:"required,ip_addr"`
	Status       string   `json:"status" binding:"required,oneof=Active Inactive"`
	Manufacturer *string  `json:"manufacturer" binding:"omitempty"`
	Description  *string  `json:"description" binding:"omitempty"`
	OsVersion    *string  `json:"osVersion" binding:"omitempty"`
	RackId       *string  `json:"rackId" binding:"omitempty,uuid"`
	RackPosition *[]uint8 `json:"rackPosition" binding:"omitempty"`
	SiteId       string   `json:"siteId" binding:"required,uuid"`
	Cpu          uint8    `json:"cpu" binding:"required"`
	Memory       uint64   `json:"memory" binding:"required"` // MB
	Disk         uint64   `json:"disk" binding:"required"`   // MB
}

type ServerUpdate struct {
	Name         *string  `json:"name" binding:"omitempty"`
	ManagementIp *string  `json:"managementIp" binding:"omitempty,ip_addr"`
	Status       *string  `json:"status" binding:"omitempty,oneof=Active Inactive"`
	Manufacturer *string  `json:"manufacturer" binding:"omitempty"`
	SerialNumber *string  `json:"serialNumber" binding:"omitempty"`
	Description  *string  `json:"description" binding:"omitempty"`
	OsVersion    *string  `json:"osVersion" binding:"omitempty"`
	RackId       *string  `json:"rackId" binding:"omitempty,uuid"`
	RackPosition *[]uint8 `json:"rackPosition" binding:"omitempty"`
	Cpu          *uint8   `json:"cpu" binding:"omitempty"`
	Memory       *uint64  `json:"memory" binding:"omitempty"` // MB
	Disk         *uint64  `json:"disk" binding:"omitempty"`   // MB
}

type ServerQuery struct {
	schemas.PageInfo
	Id           *[]string `form:"id" binding:"omitempty,list_uuid"`
	Name         *[]string `form:"name" binding:"omitempty"`
	ManagementIp *[]string `form:"managementIp" binding:"omitempty,list_ip"`
	OsVersion    *[]string `form:"osVersion" binding:"omitempty"`
	Manufacturer *[]string `form:"manufacturer" binding:"omitempty"`
	Status       *string   `form:"status" binding:"omitempty,oneof=Active Inactive"`
	SiteId       *string   `form:"siteId" binding:"omitempty,uuid"`
	RackId       *string   `form:"rackId" binding:"omitempty,uuid"`
}

type ServerShort struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	ManagementIp string `json:"managementIp"`
	Status       string `json:"status"`
}

type ServerShortList []*ServerShort

type Server struct {
	Id           string    `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
	Name         string    `json:"name"`
	ManagementIp string    `json:"managementIp"`
	Status       string    `json:"status"`
	OperStatus   string    `json:"operStatus"`
	HealthStatus string    `json:"healthStatus"`
	SerialNumber *string   `json:"serialNumber"`
	Description  *string   `json:"description"`
	Manufacturer string    `json:"manufacturer"`
	OsVersion    string    `json:"osVersion"`
	RackId       *string   `json:"rackId"`
	RackPosition *[]uint8  `json:"rackPosition"`
	Cpu          uint8     `json:"cpu"`
	Memory       uint64    `json:"memory"` // MB
	Disk         uint64    `json:"disk"`   // MB
	MonitorId    *string   `json:"monitorId"`
	TemplateId   *string   `json:"templateId"`
	SiteId       string    `json:"siteId"`
}
