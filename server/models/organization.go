package models

import (
	"time"

	"gorm.io/datatypes"
)

var OrganizationSearchFields = []string{"name", "enterpriseCode", "domainName"}
var ProxySearchFields = []string{"name", "ipAddress"}

type AuthConfig struct {
	ClientId     string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
}

type Organization struct {
	BaseDbModel
	Name           string                          `gorm:"column:name;not null"`
	EnterpriseCode string                          `gorm:"column:enterpriseCode;unique"`
	DomainName     string                          `gorm:"column:domainName;unique"`
	Active         bool                            `gorm:"column:active;type:bool;default:true"`
	LicenseCount   uint32                          `gorm:"column:licenseCount;type:int;default:0"`
	AuthType       uint8                           `gorm:"column:authType;type:int;default:0"`
	AuthConfig     *datatypes.JSONType[AuthConfig] `gorm:"column:authConfig;type:json;default:null"`
}

type Proxy struct {
	BaseDbModel
	Name           string       `gorm:"column:name;uniqueIndex:idx_name_organization_id;not null"`
	Active         bool         `gorm:"column:active;type:bool;default:true"`
	SecretKey      string       `gorm:"column:secretKey;not null"`
	IpAddress      string       `gorm:"column:ipAddress;uniqueIndex:idx_ip_address_organization_id;not null"`
	ProxyId        *string      `gorm:"column:proxyId"`
	LastSeen       *time.Time   `gorm:"column:lastSeen;default:null"`
	OrganizationId string       `gorm:"column:organizationId;uniqueIndex:idx_ip_address_organization_id;uniqueIndex:idx_name_organization_id;not null"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}
