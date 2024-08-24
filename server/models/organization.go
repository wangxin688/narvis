package models

import (
	"time"

	"gorm.io/datatypes"
)

var OrganizationSearchFields = []string{"name", "enterprise_code", "domain_name"}
var ProxySearchFields = []string{"name", "ip_address"}

type AuthConfig struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type Organization struct {
	BaseDbModel
	Name           string                          `gorm:"not null"`
	EnterpriseCode string                          `gorm:"unique"`
	DomainName     string                          `gorm:"unique"`
	Active         bool                            `gorm:"type:bool;default:true"`
	LicenseCount   uint32                          `gorm:"type:int;default:0"`
	AuthType       uint8                           `gorm:"type:int;default:0"`
	AuthConfig     *datatypes.JSONType[AuthConfig] `gorm:"type:json;default:null"`
}

type Proxy struct {
	BaseDbModel
	Name           string       `gorm:"uniqueIndex:idx_name_organization_id;not null"`
	Active         bool         `gorm:"default:true"`
	SecretKey      string       `gorm:"column:secret_key;not null"`
	IpAddress      string       `gorm:"uniqueIndex:idx_ip_address_organization_id;not null"`
	ProxyID        *string      `gorm:"column:proxy_id;"`
	LastSeen       *time.Time   `gorm:"default:null"`
	OrganizationID string       `gorm:"uniqueIndex:idx_ip_address_organization_id;uniqueIndex:idx_name_organization_id;not null"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}
