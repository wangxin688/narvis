package schemas

import "github.com/wangxin688/narvis/server/schemas"

type AuthConfig struct {
	ClientId     *string `json:"client_id" binding:"omitempty"`
	ClientSecret *string `json:"client_secret" binding:"omitempty"`
	Password     *string `json:"password" binding:"omitempty"`
}

type OrganizationCreate struct {
	Name           string      `json:"name" binding:"required"`
	EnterpriseCode string      `json:"enterprise_code" binding:"required"`
	DomainName     string      `json:"domain_name" binding:"required"`
	Active         bool        `json:"active" binding:"required,bool"`
	LicenseCount   uint32      `json:"license_count" binding:"required,gte=0,lte=10000"`
	AuthType       uint8       `json:"auth_type" binding:"required,gte=0,lte=4"` // 0: local 1: slack 2: google 3: teams 4: github
	AuthConfig     *AuthConfig `json:"auth_config" binding:"omitempty"`
}

type OrganizationUpdate struct {
	Name           *string     `json:"name" binding:"omitempty"`
	EnterpriseCode *string     `json:"enterprise_code" binding:"omitempty"`
	DomainName     *string     `json:"domain_name" binding:"omitempty"`
	Active         *bool       `json:"active" binding:"omitempty"`
	LicenseCount   *uint32     `json:"license_count" binding:"omitempty,gte=0,lte=10000"`
	AuthType       *uint8      `json:"auth_type" binding:"gte=0,lte=4,omitempty"` // 0: local 1: slack 2: google 3: teams 4: github
	AuthConfig     *AuthConfig `json:"auth_config" binding:"omitempty"`
}

type Organization struct {
	schemas.BaseResponse
	OrganizationCreate
}

type OrganizationList []Organization

type OrganizationShort struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type OrganizationShortList []OrganizationShort

type OrganizationQuery struct {
	schemas.PageInfo
	Id             *[]string `json:"id" binding:"omitempty,list_uuid"`
	Name           *[]string `json:"name" binding:"omitempty"`
	EnterpriseCode *[]string `json:"enterprise_code" binding:"omitempty"`
	DomainName     *[]string `json:"domain_name" binding:"omitempty"`
	Active         *bool     `json:"active" binding:"omitempty"`
	AuthType       *uint8    `json:"auth_type" binding:"gte=0,lte=4,omitempty"`
}
