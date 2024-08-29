package schemas

import (
	"time"

	"github.com/wangxin688/narvis/server/tools/schemas"
)

type CircuitCreate struct {
	Name         string  `json:"name" binding:"required"`
	CID          string  `json:"cid" binding:"required"`
	Status       string  `json:"status" binding:"required,oneof: Active Inactive"`
	CircuitType  string  `json:"circuit_type" binding:"required,oneof: Internet, Intranet"`
	BandWidth    uint32  `json:"band_width" binding:"required,gt=1, lt=800000"`
	IpAddress    *string `json:"ip_address" binding:"ip_addr"`
	Description  *string `json:"description" binding:"omitempty"`
	ProviderID   string  `json:"provider_id" binding:"required,uuid"`
	AInterfaceID string  `json:"a_interface_id" binding:"required,uuid"`
	ZInterfaceID *string `json:"z_interface_id" binding:"required,uuid"`
}

type CircuitUpdate struct {
	Name         *string `json:"name" binding:"omitempty"`
	CID          *string `json:"cid" binding:"omitempty"`
	Status       *string `json:"status" binding:"omitempty,oneof: Active Inactive"`
	CircuitType  *string `json:"circuit_type" binding:"omitempty,oneof: Internet, Intranet"`
	BandWidth    *uint32 `json:"band_width" binding:"omitempty,gt=1, lt=800000"`
	IpAddress    *string `json:"ip_address" binding:"omitempty,ip_addr"`
	Description  *string `json:"description" binding:"omitempty"`
	ProviderID   *string `json:"provider_id" binding:"omitempty,uuid"`
	AInterfaceID *string `json:"a_interface_id" binding:"omitempty,uuid"`
	ZInterfaceID *string `json:"z_interface_id" binding:"omitempty,uuid"`
}

type CircuitQuery struct {
	schemas.PageInfo
	Name        *[]string `form:"name" binding:"omitempty"`
	CID         *[]string `form:"cid" binding:"omitempty"`
	Status      *string   `form:"status" binding:"omitempty,oneof: Active Inactive"`
	BandWidth   *uint32   `json:"band_width" binding:"omitempty"`
	IpAddress   *string   `json:"ip_address" binding:"omitempty,ip_addr"`
	CircuitType *string   `form:"circuit_type" binding:"omitempty,oneof: Internet, Intranet"`
	ProviderID  *[]string `form:"provider_id" binding:"omitempty,list_uuid"`
	SiteID      *[]string `form:"site_id" binding:"omitempty,list_uuid"`
	DeviceID    *[]string `form:"device_id" binding:"omitempty,list_uuid"`
	InterfaceID *[]string `form:"interface_id" binding:"omitempty,list_uuid"`
	MonitorID   *[]string `form:"monitor_id" binding:"omitempty"`
}

type Circuit struct {
	ID          string        `json:"id"`
	CreatedAt   time.Time     `json:"created_at"`
	UpdatedAt   time.Time     `json:"updated_at"`
	Name        string        `json:"name"`
	CID         string        `json:"cid"`
	Status      string        `json:"status"`
	CircuitType string        `json:"circuit_type"`
	BandWidth   uint32        `json:"band_width"`
	IpAddress   *string       `json:"ip_address"`
	Description *string       `json:"description"`
	MonitorID   *string       `json:"monitor_id"`
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
	ID        string `json:"id"`
	Name      string `json:"name"`
	BandWidth uint32 `json:"band_width"`
}

type CircuitShortList []CircuitShort
