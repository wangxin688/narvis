package schemas

import (
	"time"

	"github.com/wangxin688/narvis/server/tools/schemas"
)

type ProxyCreate struct {
	Name           string `json:"name" binding:"required"`
	Active         bool   `json:"active" binding:"required"`
	OrganizationId string `json:"organizationId" binding:"required;uuid"`
}

type ProxyUpdate struct {
	Name           *string `json:"name" binding:"omitempty"`
	Active         *bool   `json:"active" binding:"omitempty,bool"`
	OrganizationId *string `json:"organizationId" binding:"omitempty,uuid"`
}

type ProxyQuery struct {
	schemas.PageInfo
	Id             *[]string `form:"id" binding:"omitempty,list_uuid"`
	Name           *[]string `form:"name" binding:"omitempty"`
	Active         *bool     `form:"active" binding:"omitempty,bool"`
	OrganizationId *string   `form:"organizationId" binding:"omitempty,uuid"`
}

type Proxy struct {
	Id           string            `json:"id" binding:"required,uuid"`
	CreatedAt    time.Time         `json:"CreatedAt"`
	UpdatedAt    time.Time         `json:"UpdatedAt"`
	Name         string            `json:"name" binding:"required"`
	Active       bool              `json:"active" binding:"required"`
	Organization OrganizationShort `json:"organization"`
}

type ProxyList []Proxy

type ProxyShort struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ProxyShortList []ProxyShort
