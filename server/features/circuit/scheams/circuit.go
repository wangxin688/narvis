package schemas

import (
	"time"

	"github.com/wangxin688/narvis/server/tools/schemas"
)

type CircuitCreate struct {
	Name         string  `json:"name" binding:"required"`
	CId          string  `json:"cid" binding:"required"`
	Status       string  `json:"status" binding:"required,oneof: Active Inactive"`
	CircuitType  string  `json:"circuitType" binding:"required,oneof: Internet, Intranet"`
	BandWidth    uint32  `json:"bandWidth" binding:"required,gt=1, lt=800000"`
	IpAddress    *string `json:"ipAddress" binding:"ip_addr"`
	Description  *string `json:"description" binding:"omitempty"`
	ProviderId   string  `json:"providerId" binding:"required,uuid"`
	AInterfaceId string  `json:"aInterfaceId" binding:"required,uuid"`
	ZInterfaceId *string `json:"zInterfaceId" binding:"required,uuid"`
}

type CircuitUpdate struct {
	Name         *string `json:"name" binding:"omitempty"`
	CId          *string `json:"cid" binding:"omitempty"`
	Status       *string `json:"status" binding:"omitempty,oneof: Active Inactive"`
	CircuitType  *string `json:"circuitType" binding:"omitempty,oneof: Internet, Intranet"`
	BandWidth    *uint32 `json:"bandWidth" binding:"omitempty,gt=1, lt=800000"`
	IpAddress    *string `json:"ipAddress" binding:"omitempty,ip_addr"`
	Description  *string `json:"description" binding:"omitempty"`
	ProviderId   *string `json:"providerId" binding:"omitempty,uuid"`
	AInterfaceId *string `json:"aInterfaceId" binding:"omitempty,uuid"`
	ZInterfaceId *string `json:"zInterfaceId" binding:"omitempty,uuid"`
}

type CircuitQuery struct {
	schemas.PageInfo
	Name        *[]string `form:"name" binding:"omitempty"`
	CId         *[]string `form:"cid" binding:"omitempty"`
	Status      *string   `form:"status" binding:"omitempty,oneof: Active Inactive"`
	BandWidth   *uint32   `form:"bandWidth" binding:"omitempty"`
	IpAddress   *string   `form:"ipAddress" binding:"omitempty,ip_addr"`
	CircuitType *string   `form:"circuitType" binding:"omitempty,oneof: Internet, Intranet"`
	ProviderId  *[]string `form:"providerId" binding:"omitempty,list_uuid"`
	SiteId      *[]string `form:"siteId" binding:"omitempty,list_uuid"`
	DeviceId    *[]string `form:"deviceId" binding:"omitempty,list_uuid"`
	InterfaceId *[]string `form:"interfaceId" binding:"omitempty,list_uuid"`
	MonitorId   *[]string `form:"monitorId" binding:"omitempty"`
}

type Circuit struct {
	Id          string        `json:"id"`
	CreatedAt   time.Time     `json:"createdAt"`
	UpdatedAt   time.Time     `json:"updatedAt"`
	Name        string        `json:"name"`
	CId         string        `json:"cid"`
	Status      string        `json:"status"`
	CircuitType string        `json:"circuitType"`
	BandWidth   uint32        `json:"bandWidth"`
	IpAddress   *string       `json:"ipAddress"`
	Description *string       `json:"description"`
	MonitorId   *string       `json:"monitorId"`
	Provider    ProviderShort `json:"provider"`
	// ASite ds.SiteShort `json:"a_site"`
	// ADevice ds.DeviceShort `json:"a_device"`
	// AInterface ds.InterfaceShort `json:"a_interface"`
	// ZSite ds.SiteShort `json:"z_site"`
	// ZDevice ds.DeviceShort `json:"z_device"`
	// ZInterface ds.InterfaceShort `json:"z_interface"`
}

type CircuitList []Circuit

type CircuitShort struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	BandWidth uint32 `json:"bandWidth"`
}

type CircuitShortList []CircuitShort
