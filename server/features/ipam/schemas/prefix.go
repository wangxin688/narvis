package schemas

import (
	"time"

	"github.com/wangxin688/narvis/server/tools/schemas"
)

type Prefix struct {
	Id          string    `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Range       string    `json:"range"`
	Version     string    `json:"version"`
	VlanId      *uint32   `json:"vlanId"`
	VlanName    *string   `json:"vlanName"`
	Type        string    `json:"type"`
	SiteId      string    `json:"siteId"`
	Utilization float64   `json:"utilization"` // percentage
}

type PrefixCreate struct {
	Range    string  `json:"range" binding:"required,cidr_v_any"`
	VlanId   *uint32 `json:"vlanId" binding:"omitempty,gte=1,lte=4094"` // not support VxLAN now
	VlanName *string `json:"vlanName" binding:"omitempty"`
	Type     string  `json:"type" binding:"required,oneof=Dynamic Static"`
	SiteId   string  `json:"siteId" binding:"required,uuid"`
}

type PrefixUpdate struct {
	Range    *string `json:"range" binding:"omitempty,cidr"`
	VlanId   *uint32 `json:"vlanId" binding:"omitempty,gt=1,lte=4094"`
	VlanName *string `json:"vlanName" binding:"omitempty"`
	Type     *string `json:"type" binding:"omitempty,oneof=Dynamic Static"`
	SiteId   *string `json:"siteId" binding:"omitempty,uuid"`
}

type PrefixQuery struct {
	schemas.PageInfo
	SiteId   *string   `form:"siteId" binding:"omitempty,uuid"`
	Range    *[]string `form:"range" binding:"omitempty,list_cidr"`
	Type     *string   `form:"type" binding:"omitempty,oneof=Dynamic Static"`
	Version  string    `form:"version" binding:"omitempty,oneof=IPv4 IPv6"`
	VlanId   *[]uint32 `form:"vlanId" binding:"omitempty,gte=1,lte=4094"`
	VlanName *[]string `form:"vlanName" binding:"omitempty,"`
}
