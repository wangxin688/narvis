package schemas

import (
	"time"

	"github.com/wangxin688/narvis/server/tools/schemas"
)

type IpAddress struct {
	Id          string    `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Address     string    `json:"address"`
	Status      string    `json:"status"` // Active, Reserved, Deprecated
	MacAddress  *string   `json:"macAddress"`
	Type        string    `json:"type"` // Dynamic, Static, Gateway, Broadcast, NetworkId
	Vlan        *uint32   `json:"vlan"`
	Range       *string   `json:"range"`
	Description *string   `json:"description"`
	SiteId      string    `json:"siteId"`
}

type IpAddressCreate struct {
	Address     string  `json:"address" binding:"required,ip_addr"`
	Status      string  `json:"status" binding:"required,oneof=Active Reserved Deprecated"`
	MacAddress  *string `json:"macAddress" binding:"omitempty,mac"`
	Type        string  `json:"type" binding:"required,oneof=Dynamic Static Gateway Broadcast NetworkId"`
	Vlan        *uint32 `json:"vlan" binding:"omitempty"`
	Range       *string `json:"range" binding:"omitempty,ip_range,cidr"`
	Description *string `json:"description"`
	SiteId      string  `json:"siteId" binding:"required,uuid"`
}

type IpAddressUpdate struct {
	Address     *string `json:"address" binding:"omitempty,ip_addr"`
	Status      *string `json:"status" binding:"omitempty,oneof=Active Reserved Deprecated"`
	MacAddress  *string `json:"macAddress" binding:"omitempty,mac"`
	Type        *string `json:"type" binding:"omitempty,oneof=Dynamic Static Gateway Broadcast NetworkId"`
	Vlan        *uint32 `json:"vlan" binding:"omitempty"`
	Range       *string `json:"range" binding:"omitempty,ip_range,cidr"`
	Description *string `json:"description" binding:"omitempty"`
	SiteId      *string `json:"siteId" binding:"omitempty,uuid"`
}

type IpAddressQuery struct {
	schemas.PageInfo
	Address *[]string `json:"address" binding:"omitempty,ip_addr"`
	Status  *[]string `json:"status" binding:"omitempty"` // Active, Reserved, Deprecated
	Vlan    *[]uint32 `json:"vlan" binding:"omitempty"`
	Type    *[]string `json:"type" binding:"omitempty"` // Dynamic, Static, Gateway, Broadcast, NetworkId
	Range   *string   `json:"range" binding:"omitempty,ip_range,cidr"`
	SiteId  *string   `json:"siteId" binding:"omitempty,uuid"`
}
