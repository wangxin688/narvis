package schemas

import (
	"time"

	"github.com/wangxin688/narvis/server/tools/schemas"
)

type ScanDeviceCreate struct {
	Range          []string `json:"range" binding:"required,list_cidr"`
	Community      *string  `json:"community" binding:"omitempty"`
	Port           *uint16  `json:"port" binding:"omitempty,gt=0,lte=65535"`
	Timeout        *uint8   `json:"timeout" binding:"omitempty,gt=1,lte=60"`
	MaxRepetitions *uint8   `json:"maxRepetitions" binding:"omitempty,gt=10,lte=100"`
}

func (s *ScanDeviceCreate) SetDefaultValue() {
	if s.Port == nil {
		s.Port = new(uint16)
		*s.Port = 161
	}
	if s.MaxRepetitions == nil {
		s.MaxRepetitions = new(uint8)
		*s.MaxRepetitions = 10
	}
	if s.Timeout == nil {
		s.Timeout = new(uint8)
		*s.Timeout = 3
	}
	if s.Community == nil {
		s.Community = new(string)
		*s.Community = "public"
	}
}

type ScanDeviceCreateResponse struct {
	TaskIds []string `json:"taskIds"`
}

type ScanDevice struct {
	Id           string    `json:"id"`
	CreatedAt    time.Time `json:"createdAt"`
	Name         string    `json:"name"`
	ManagementIp string    `json:"managementIp"`
	Platform     string    `json:"platform"`
	DeviceModel  string    `json:"deviceModel"`
	Manufacturer string    `json:"manufacturer"`
	ChassisId    string    `json:"chassisId"`
	Description  string    `json:"description"`
}

type ScanDeviceQuery struct {
	schemas.PageInfo
	Platform     *[]string `form:"platform" binding:"omitempty"`
	Manufacturer *[]string `form:"manufacturer" binding:"omitempty"`
	ManagementIp *[]string `form:"managementIp" binding:"omitempty"`
	DeviceModel  *[]string `form:"deviceModel" binding:"omitempty"`
	ChassisId    *[]string `form:"chassisId" binding:"omitempty"`
}

type ScanDeviceUpdate struct {
	SiteId       string   `json:"siteId" binding:"required,uuid"`
	Status       string   `json:"status" binding:"required,oneof= Active Inactive"`
	DeviceRole   string   `json:"deviceRole" binding:"required"`
	Floor        *string  `json:"floor" binding:"omitempty"`
	RackId       *string  `json:"rackId" binding:"omitempty,uuid"`
	RackPosition *[]uint8 `json:"rackPosition" binding:"omitempty"`
}

type ScanDeviceBatchUpdate struct {
	Ids        []string `json:"ids" binding:"required,list_uuid"`
	SiteId     string   `json:"siteId" binding:"required,uuid"`
	Status     string   `json:"status" binding:"required,oneof= Active Inactive"`
	DeviceRole string   `json:"deviceRole" binding:"required"`
	Floor      *string  `json:"floor" binding:"omitempty"`
}

type ScanApCreate struct {
	SiteId string `json:"siteId" binding:"required,uuid"`
}

type ScanDeviceDetailTask struct {
	SiteId string `json:"siteId" binding:"required,uuid"`
}
