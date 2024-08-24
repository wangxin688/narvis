package schemas

import (
	"time"

	"github.com/wangxin688/narvis/server/tools/schemas"
)

type ProxyCreate struct {
	Name           string `json:"name" binding:"required"`
	Active         bool   `json:"active" binding:"required"`
	SecretKey      string `json:"secret_key" binding:"required;len=32"`
	OrganizationID string `json:"organization_id" binding:"required;uuid"`
}

type ProxyUpdate struct {
	Name           *string `json:"name" binding:"omitempty"`
	Active         *bool   `json:"active" binding:"omitempty,bool"`
	OrganizationID *string `json:"organization_id" binding:"omitempty,uuid"`
}

type ProxyQuery struct {
	schemas.PageInfo
	ID             *[]string `form:"id" binding:"omitempty,list_uuid"`
	Name           *[]string `form:"name" binding:"omitempty"`
	Active         *bool     `form:"active" binding:"omitempty,bool"`
	OrganizationID *string   `form:"organization_id" binding:"omitempty,uuid"`
}

type Proxy struct {
	ID             string            `json:"id" binding:"required,uuid"`
	CreatedAt      time.Time         `json:"created_at"`
	UpdatedAt      time.Time         `json:"updated_at"`
	Name           string            `json:"name" binding:"required"`
	Active         bool              `json:"active" binding:"required"`
	SecretKey      string            `json:"secret_key" binding:"required;len=32"`
	Organization   OrganizationShort `json:"organization"`
}

type ProxyList []Proxy

type ProxyShort struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ProxyShortList []ProxyShort
