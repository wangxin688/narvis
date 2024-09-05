package schemas

import "github.com/wangxin688/narvis/server/tools/schemas"

type Prefix struct {
	Id          string  `json:"id"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
	Range       string  `json:"range"`
	Version     string  `json:"version"`
	VlanId      *uint32 `json:"vlanId"`
	VlanName    *string `json:"vlanName"`
	Type        string  `json:"type"`
	SiteId      string  `json:"siteId"`
	Utilization float32 `json:"utilization"`
}

type PrefixCreate struct {
	Range    string  `json:"range" binding:"required,cidr"`
	VlanId   *uint32 `json:"vlanId" binding:"omitempty,gt=1,lte=4094"` // not support VxLAN now
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
	SiteId   *string   `json:"siteId" binding:"omitempty,uuid"`
	Range    *[]string `json:"range" binding:"omitempty,list_cidr"`
	Type     *string   `json:"type" binding:"omitempty,oneof=Dynamic Static"`
	VlanId   *[]uint32 `json:"vlanId" binding:"omitempty,list_gte=1,lte=4094"`
	VlanName *[]string `json:"vlanName" binding:"omitempty,list"`
}
