package schemas

import "github.com/wangxin688/narvis/server/schemas"

type ProxyCreate struct {
	Name           string `json:"name" binding:"required"`
	Active         bool   `json:"active" binding:"required"`
	SecretKey      string `json:"secret_key" binding:"required;len=32"`
	OrganizationId string `json:"organization_id" binding:"required;uuid"`
}

type ProxyUpdate struct {
	Name           *string `json:"name" binding:"omitempty"`
	Active         *bool   `json:"active" binding:"omitempty,bool"`
	OrganizationId *string `json:"organization_id" binding:"omitempty,uuid"`
}

type ProxyQuery struct {
	schemas.PageInfo
	Id             *[]string `json:"id" binding:"omitempty,list_uuid"`
	Name           *[]string `json:"name" binding:"omitempty"`
	Active         *bool     `json:"active" binding:"omitempty,bool"`
	OrganizationId *string   `json:"organization_id" binding:"omitempty,uuid"`
}

type Proxy struct {
	schemas.BaseResponse
	ProxyCreate
	Organization OrganizationShort `json:"organization"`
}

type ProxyList []Proxy

type ProxyShort struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type ProxyShortList []ProxyShort
