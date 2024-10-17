package schemas

import (
	"time"

	ifs "github.com/wangxin688/narvis/server/features/infra/schemas"
	"github.com/wangxin688/narvis/server/tools/schemas"
)

type AuthConfig struct {
	ClientId     string `json:"clientId" binding:"omitempty"`
	ClientSecret string `json:"clientSecret" binding:"omitempty"`
}

type OrganizationCreate struct {
	Name           string      `json:"name" binding:"required"`
	EnterpriseCode string      `json:"enterpriseCode" binding:"required"`
	DomainName     string      `json:"domainName" binding:"required"`
	Active         bool        `json:"active" binding:"required,bool"`
	LicenseCount   uint32      `json:"licenseCount" binding:"required,gte=0,lte=10000"`
	AuthType       uint8       `json:"authType" binding:"required,gte=0,lte=4"` // 0: local 1: slack 2: google 3: teams 4: github
	AdminPassword  string      `json:"adminPassword" binding:"required"`
	AuthConfig     *AuthConfig `json:"authConfig" binding:"omitempty"`
}

type OrganizationUpdate struct {
	Name         *string     `json:"name" binding:"omitempty"`
	DomainName   *string     `json:"domainName" binding:"omitempty"`
	Active       *bool       `json:"active" binding:"omitempty"`
	LicenseCount *uint32     `json:"licenseCount" binding:"omitempty,gte=0,lte=10000"`
	AuthType     *uint8      `json:"authType" binding:"gte=0,lte=4,omitempty"` // 0: local 1: slack 2: google 3: teams 4: github
	AuthConfig   *AuthConfig `json:"authConfig" binding:"omitempty"`
}

type Organization struct {
	Id             string      `json:"id" binding:"required,uuid"`
	CreatedAt      time.Time   `json:"createdAt"`
	UpdatedAt      time.Time   `json:"updatedAt"`
	Name           string      `json:"name" binding:"required"`
	EnterpriseCode string      `json:"enterpriseCode" binding:"required"`
	DomainName     string      `json:"domainName" binding:"required"`
	Active         bool        `json:"active" binding:"required,bool"`
	LicenseCount   uint32      `json:"licenseCount" binding:"required,gte=0,lte=10000"`
	AuthType       uint8       `json:"authType" binding:"required,gte=0,lte=4"` // 0: local 1: slack 2: google 3: teams 4: github
	AuthConfig     *AuthConfig `json:"authConfig" binding:"omitempty"`
}

type OrganizationList []Organization

type OrganizationShort struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type OrganizationShortList []OrganizationShort

type OrganizationQuery struct {
	schemas.PageInfo
	Id             *[]string `form:"id" binding:"omitempty,list_uuid"`
	Name           *[]string `form:"name" binding:"omitempty"`
	EnterpriseCode *[]string `form:"enterpriseCode" binding:"omitempty"`
	DomainName     *[]string `form:"domainName" binding:"omitempty"`
	Active         *bool     `form:"active" binding:"omitempty"`
	AuthType       *uint8    `form:"authType" binding:"gte=0,lte=4,omitempty"`
}

type OrganizationSettings struct {
	SnmpCredential     *ifs.SnmpV2CredentialUpdate   `json:"snmpCredential"`
	CliCredential      *ifs.CliCredentialUpdate      `json:"cliCredential"`
	RestconfCredential *ifs.RestconfCredentialUpdate `json:"restconfCredential"`
}
