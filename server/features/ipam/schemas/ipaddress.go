package schemas

import "time"

type IpAddress struct {
	Id          string    `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Address     string    `json:"address"`
	Status      string    `json:"status"` // Active, Reserved, Inactive, Gateway
	MacAddress  *string   `json:"macAddress"`
	Type        string    `json:"type"` // Dynamic, Static, Gateway
	Description *string   `json:"description"`
	SiteId      string    `json:"siteId"`
}

type IpAddressCreate struct {
	Address     string  `json:"address" binding:"required,ip_addr"`
	Status      string  `json:"status" binding:"required,oneof=Active Reserved Inactive Gateway"`
	MacAddress  *string `json:"macAddress" binding:"omitempty,mac"`
	Type        string  `json:"type" binding:"required,oneof=Dynamic Static Gateway"`
	Description *string `json:"description"`
	SiteId      string  `json:"siteId" binding:"required,uuid"`
}

type IpAddressUpdate struct {
	Address     *string `json:"address" binding:"omitempty,ip_addr"`
	Status      *string `json:"status" binding:"omitempty,oneof=Active Reserved Inactive Gateway"`
	MacAddress  *string `json:"macAddress" binding:"omitempty,mac"`
	Type        *string `json:"type" binding:"omitempty,oneof=Dynamic Static Gateway"`
	Description *string `json:"description" binding:"omitempty"`
	SiteId      *string `json:"siteId" binding:"omitempty,uuid"`
}

type IpAddressQuery struct {
	Address *[]string `json:"address" binding:"omitempty,ip_addr"`
	Status  *[]string `json:"status" binding:"omitempty"` // Active, Reserved, Inactive, Gateway
	Type    *[]string `json:"type" binding:"omitempty"`   // Dynamic, Static, Gateway
	Range   *string   `json:"range" binding:"omitempty,ip_range,cidr"`
	SiteId  *string   `json:"siteId" binding:"omitempty,uuid"`
}
