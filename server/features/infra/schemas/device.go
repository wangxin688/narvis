package schemas

import (
	"time"

	"github.com/wangxin688/narvis/server/tools/schemas"
)

type DeviceCreate struct {
	Name         string  `json:"name" binding:"required"`
	ManagementIp string  `json:"managementIp" binding:"required,ip_addr"`
	Status       string  `json:"status" binding:"required,oneof=Active Inactive"`
	DeviceRole   string  `json:"deviceRole" binding:"required"`
	DeviceModel  *string `json:"deviceModel" binding:"omitempty"`
	Manufacturer *string `json:"manufacturer" binding:"omitempty"`
	SerialNumber *string `json:"serialNumber" binding:"omitempty"`
	Description  *string `json:"description" binding:"omitempty"`
	OsVersion    *string `json:"osVersion" binding:"omitempty"`
	Floor        *string `json:"location" binding:"omitempty"`
	RackId       *string `json:"rackId" binding:"omitempty,uuid"`
	RackPosition []uint8 `json:"rackPosition" binding:"omitempty"`
	SiteId       string  `json:"siteId" binding:"required,uuid"`
}

type DeviceUpdate struct {
	Name         *string  `json:"name" binding:"omitempty"`
	ManagementIp *string  `json:"managementIp" binding:"omitempty,ip_addr"`
	Status       *string  `json:"status" binding:"omitempty,oneof=Active Inactive"`
	DeviceRole   *string  `json:"deviceRole" binding:"omitempty"`
	DeviceModel  *string  `json:"deviceModel" binding:"omitempty"`
	Manufacturer *string  `json:"manufacturer" binding:"omitempty"`
	SerialNumber *string  `json:"serialNumber" binding:"omitempty"`
	Description  *string  `json:"description" binding:"omitempty"`
	OsVersion    *string  `json:"osVersion" binding:"omitempty"`
	Floor        *string  `json:"location" binding:"omitempty"`
	RackId       *string  `json:"rackId" binding:"omitempty,uuid"`
	RackPosition *[]uint8 `json:"rackPosition" binding:"omitempty"`
}

type DeviceQuery struct {
	schemas.PageInfo
	Id            *[]string `form:"id" binding:"omitempty,list_uuid"`
	Name          *[]string `form:"name" binding:"omitempty"`
	ManagementIp  *[]string `form:"managementIp" binding:"omitempty,list_ip"`
	DeviceRole    *[]string `form:"deviceRole" binding:"omitempty"`
	DeviceModel   *[]string `form:"deviceModel" binding:"omitempty"`
	ProductFamily *string   `form:"productFamily" binding:"omitempty,oneof=Gateway Switching WlanAP"`
	Manufacturer  *[]string `form:"manufacturer" binding:"omitempty"`
	Status        *string   `form:"status" binding:"omitempty,oneof=Active Inactive"`
	SiteId        *string   `form:"siteId" binding:"omitempty,uuid"`
	RackId        *string   `form:"rackId" binding:"omitempty,uuid"`
	Floor         *string   `form:"location" binding:"omitempty"`
	SerialNumber  *[]string `form:"serialNumber" binding:"omitempty"`
}

type DeviceShort struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	ManagementIp string `json:"managementIp"`
	Status       string `json:"status"`
}

type DeviceShortList []DeviceShort

type Device struct {
	Id            string    `json:"id"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	Name          string    `json:"name"`
	ManagementIp  string    `json:"managementIp"`
	Platform      string    `json:"platform"`
	ProductFamily string    `json:"productFamily"`
	Status        string    `json:"status"`
	OperStatus    string    `json:"operStatus"`
	HealthStatus  string    `json:"healthStatus"`
	SerialNumber  *string   `json:"serialNumber"`
	Description   *string   `json:"description"`
	DeviceModel   string    `json:"deviceModel"`
	DeviceRole    string    `json:"deviceRole"`
	Floor         *string   `json:"location"`
	IsRegistered  bool      `json:"isRegistered"`
	OsVersion     *string   `json:"osVersion"`
	OsPatch       *string   `json:"osPatch"`
	RackId        *string   `json:"rackId"`
	PackPosition  *[]uint8  `json:"packPosition"`
	MonitorId     *string   `json:"monitorId"`
	TemplateId    *string   `json:"templateId"`
	SiteId        string    `json:"siteId"`
}

type DeviceList []Device
