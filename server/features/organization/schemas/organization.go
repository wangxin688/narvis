package schemas

import (
	"time"

	"github.com/wangxin688/narvis/server/tools/schemas"
)

type AuthConfig struct {
	ClientID     string `json:"client_id" binding:"omitempty"`
	ClientSecret string `json:"client_secret" binding:"omitempty"`
}

type OrganizationCreate struct {
	Name           string      `json:"name" binding:"required"`
	EnterpriseCode string      `json:"enterprise_code" binding:"required"`
	DomainName     string      `json:"domain_name" binding:"required"`
	Active         bool        `json:"active" binding:"required,bool"`
	LicenseCount   uint32      `json:"license_count" binding:"required,gte=0,lte=10000"`
	AuthType       uint8       `json:"auth_type" binding:"required,gte=0,lte=4"` // 0: local 1: slack 2: google 3: teams 4: github
	AdminPassword  string      `json:"admin_password" binding:"required"`
	AuthConfig     *AuthConfig `json:"auth_config" binding:"omitempty"`
}

type OrganizationUpdate struct {
	Name         *string     `json:"name" binding:"omitempty"`
	DomainName   *string     `json:"domain_name" binding:"omitempty"`
	Active       *bool       `json:"active" binding:"omitempty"`
	LicenseCount *uint32     `json:"license_count" binding:"omitempty,gte=0,lte=10000"`
	AuthType     *uint8      `json:"auth_type" binding:"gte=0,lte=4,omitempty"` // 0: local 1: slack 2: google 3: teams 4: github
	AuthConfig   *AuthConfig `json:"auth_config" binding:"omitempty"`
}

type Organization struct {
	ID             string      `json:"id" binding:"required,uuid"`
	CreatedAt      time.Time   `json:"created_at"`
	UpdatedAt      time.Time   `json:"updated_at"`
	Name           string      `json:"name" binding:"required"`
	EnterpriseCode string      `json:"enterprise_code" binding:"required"`
	DomainName     string      `json:"domain_name" binding:"required"`
	Active         bool        `json:"active" binding:"required,bool"`
	LicenseCount   uint32      `json:"license_count" binding:"required,gte=0,lte=10000"`
	AuthType       uint8       `json:"auth_type" binding:"required,gte=0,lte=4"` // 0: local 1: slack 2: google 3: teams 4: github
	AuthConfig     *AuthConfig `json:"auth_config" binding:"omitempty"`
}

type OrganizationList []Organization

type OrganizationShort struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type OrganizationShortList []OrganizationShort

type OrganizationQuery struct {
	schemas.PageInfo
	ID             *[]string `form:"id" binding:"omitempty,list_uuid"`
	Name           *[]string `form:"name" binding:"omitempty"`
	EnterpriseCode *[]string `form:"enterprise_code" binding:"omitempty"`
	DomainName     *[]string `form:"domain_name" binding:"omitempty"`
	Active         *bool     `form:"active" binding:"omitempty"`
	AuthType       *uint8    `form:"auth_type" binding:"gte=0,lte=4,omitempty"`
}
