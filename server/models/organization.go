package models

import (
	"time"

	"github.com/wangxin688/narvis/server/global"
	"gorm.io/datatypes"
)

var OrganizationSearchFields = []string{"name", "enterprise_code", "domain_name"}
var ProxySearchFields = []string{"name", "ip_address"}

type AuthConfig struct {
	ClientId     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

type Organization struct {
	global.BaseDbModel
	Name           string                          `gorm:"unique"`
	EnterpriseCode string                          `gorm:"unique"`
	DomainName     string                          `gorm:"unique"`
	Active         bool                            `gorm:"type:bool;default:true"`
	LicenseCount   uint32                          `gorm:"type:int;default:0"`
	AuthType       uint8                           `gorm:"type:int;default:0"`
	AuthConfig     *datatypes.JSONType[AuthConfig] `gorm:"type:json;default:null"`
}

type Proxy struct {
	global.BaseDbModel
	Name           string       `gorm:"uniqueIndex:idx_name_organization_id;not null"`
	Active         bool         `gorm:"default:true"`
	SecretKey      string       `gorm:"column:secret_key;not null"`
	IpAddress      string       `gorm:"uniqueIndex:idx_ip_address_organization_id;not null"`
	ProxyId        *string      `gorm:"column:proxy_id;"`
	LastSeen       *time.Time   `gorm:"default:null"`
	OrganizationId string       `gorm:"uniqueIndex:idx_ip_address_organization_id;uniqueIndex:idx_name_organization_id;not null"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}
