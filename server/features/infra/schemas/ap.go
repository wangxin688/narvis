package schemas

import (
	"time"

	"github.com/wangxin688/narvis/server/tools/schemas"
)

type ApCoordinate struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

type AP struct {
	Id              string    `json:"id"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
	Name            string    `json:"name"`
	Status          string    `json:"status"`
	OperStatus      string    `json:"operStatus"`
	HealthStatus    string    `json:"healthStatus"`
	MacAddress      *string   `json:"macAddress"`
	SerialNumber    *string   `json:"serialNumber"`
	ManagementIp    string    `json:"managementIp"`
	DeviceModel     string    `json:"deviceType"`
	Manufacturer    string    `json:"manufacturer"`
	DeviceRole      string    `json:"deviceRole"`
	OsVersion       *string   `json:"osVersion"`
	Floor           *string   `json:"floor"`
	GroupName       *string   `json:"groupName"`
	CoordinateX     *float32  `json:"coordinate"`
	CoordinateY     *float32  `json:"coordinate"`
	CoordinateZ     *float32  `json:"coordinate"`
	WlanACIpAddress *string   `json:"wlanACIpAddress"`
	SiteId          string    `json:"siteId"`
}

type APList []AP

type APShort struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	ManagementIP string `json:"managementIp"`
}

type APShortList []APShort

type ApQuery struct {
	schemas.PageInfo
	Name         *[]string `form:"name" binding:"omitempty"`
	ManagementIp *[]string `form:"managementIp" binding:"omitempty,list_ip"`
	DeviceModel  *[]string `form:"deviceType" binding:"omitempty"`
	Manufacturer *[]string `form:"manufacturer" binding:"omitempty"`
	Floor        *string   `form:"location" binding:"omitempty"`
	Status       *string   `form:"status" binding:"omitempty,oneof=Active Inactive"`
	SiteId       *string   `form:"siteId" binding:"omitempty,uuid"`
	SerialNumber *string   `form:"serialNumber" binding:"omitempty"`
}
