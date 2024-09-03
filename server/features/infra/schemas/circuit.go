package schemas

import (
	"time"

	"github.com/wangxin688/narvis/server/tools/schemas"
)

type CircuitCreate struct {
	Name        string  `json:"name" binding:"required"`
	CId         *string `json:"cid" binding:"required"`
	Status      string  `json:"status" binding:"required,oneof: Active Inactive"`
	CircuitType string  `json:"circuitType" binding:"required,oneof: Internet, MPLS, IEPL, DPLC, DarkFiber, ADSL"`
	RxBandWidth uint32  `json:"rxBandWidth" binding:"required,gt=1, lt=800000"`
	TxBandWidth uint32  `json:"txBandWidth" binding:"required,gt=1, lt=800000"`
	Ipv4Address *string `json:"ipv4Address" binding:"ipv4"`
	Ipv6Address *string `json:"ipv6Address" binding:"ipv6"`
	Description *string `json:"description" binding:"omitempty"`
	Provider    string  `json:"provider" binding:"required"`
	InterfaceId string  `json:"interfaceId" binding:"required,uuid"`
}

type CircuitUpdate struct {
	Name        *string `json:"name" binding:"omitempty"`
	CId         *string `json:"cid" binding:"omitempty"`
	Status      *string `json:"status" binding:"omitempty,oneof: Active Inactive"`
	CircuitType *string `json:"circuitType" binding:"omitempty,oneof: Internet, Intranet"`
	RxBandWidth *uint32 `json:"rxBandWidth" binding:"omitempty,gt=1, lt=800000"`
	TxBandWidth *uint32 `json:"txBandWidth" binding:"omitempty,gt=1, lt=800000"`
	Ipv4Address *string `json:"ipv4Address" binding:"omitempty,ipv4"`
	Ipv6Address *string `json:"ipv6Address" binding:"omitempty,ipv6"`
	Description *string `json:"description" binding:"omitempty"`
	Provider    *string `json:"provider" binding:"omitempty"`
	InterfaceId *string `json:"interfaceId" binding:"omitempty,uuid"`
}

type CircuitQuery struct {
	schemas.PageInfo
	Name        *[]string `form:"name" binding:"omitempty"`
	CId         *[]string `form:"cid" binding:"omitempty"`
	Status      *string   `form:"status" binding:"omitempty,oneof: Active Inactive"`
	Ipv4Address *[]string   `form:"ipv4Address" binding:"omitempty,list_ip"`
	Ipv6Address *[]string   `form:"ipv6Address" binding:"omitempty,list_ip"`
	CircuitType *[]string `form:"circuitType" binding:"omitempty,oneof: Internet, MPLS, IEPL, DPLC, DarkFiber, ADSL"`
	Provider    *[]string `form:"provider" binding:"omitempty"`
	SiteId      *[]string `form:"siteId" binding:"omitempty,list_uuid"`
	DeviceId    *[]string `form:"deviceId" binding:"omitempty,list_uuid"`
	InterfaceId *[]string `form:"interfaceId" binding:"omitempty,list_uuid"`
	MonitorId   *[]string `form:"monitorId" binding:"omitempty"`
}

type Circuit struct {
	Id          string    `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Name        string    `json:"name"`
	CId         string    `json:"cid"`
	Status      string    `json:"status"`
	CircuitType string    `json:"circuitType"`
	RxBandWidth uint32    `json:"rxBandWidth"`
	TxBandWidth uint32    `json:"txBandWidth"`
	Ipv4Address *string   `json:"ipv4Address"`
	Ipv6Address *string   `json:"ipv6Address"`
	Description *string   `json:"description"`
	MonitorId   *string   `json:"monitorId"`
	Provider    string    `json:"provider"`
	// ASite ds.SiteShort `json:"a_site"`
	// ADevice ds.DeviceShort `json:"a_device"`
	// AInterface ds.InterfaceShort `json:"a_interface"`
	// ZSite ds.SiteShort `json:"z_site"`
	// ZDevice ds.DeviceShort `json:"z_device"`
	// ZInterface ds.InterfaceShort `json:"z_interface"`
}

type CircuitList []Circuit

type CircuitShort struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Provider    string `json:"provider"`
	RxBandWidth uint32 `json:"RxBandWidth"`
	TxBandWidth uint32 `json:"TxBandWidth"`
}

type CircuitShortList []CircuitShort
