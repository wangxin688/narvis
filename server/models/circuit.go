package models

import "github.com/wangxin688/narvis/server/global"

var CircuitSearchFields = []string{"name", "ip_address"}

type Circuit struct {
	global.BaseDbModel

	Name           string          `gorm:"uniqueIndex:idx_name_organization_id;not null"`
	CId            string          `gorm:"column:c_id;uniqueIndex:idx_circuit_id_organization_id;not null"`
	Status         string          `gorm:"not null;default:Active"`   // Active/Inactive
	CircuitType    string          `gorm:"not null;default:Internet"` // Internet/Intranet
	BandWidth      uint32          `gorm:"not null"`                  //Mbps
	IpAddress      string          `gorm:"not null"`
	Description    *string         `gorm:"column:description;default:null"`
	ProviderId     string          `gorm:"type:uuid;not null"`
	Provider       Provider        `gorm:"constraint:Ondelete:RESTRICT"`
	ASiteId        string          `gorm:"type:uuid;not null"`
	ASite          Site            `gorm:"constraint:Ondelete:RESTRICT;foreignKey:ASiteId"`
	ADeviceId      string          `gorm:"type:uuid;not null"`
	ADevice        Device          `gorm:"constraint:Ondelete:RESTRICT;foreignKey:ADeviceId"`
	AInterfaceId   string          `gorm:"type:uuid;not null"`
	AInterface     DeviceInterface `gorm:"constraint:Ondelete:RESTRICT;foreignKey:AInterfaceId"`
	ZSiteId        string          `gorm:"type:uuid;not null"`
	ZSite          Site            `gorm:"constraint:Ondelete:RESTRICT;foreignKey:ZSiteId"`
	ZDeviceId      string          `gorm:"type:uuid;not null"`
	ZDevice        Device          `gorm:"constraint:Ondelete:RESTRICT;foreignKey:ZDeviceId"`
	ZInterfaceId   string          `gorm:"type:uuid;not null"`
	ZInterface     DeviceInterface `gorm:"constraint:Ondelete:RESTRICT;foreignKey:ZInterfaceId"`
	OrganizationId string          `gorm:"type:uuid;not null"`
	Organization   Organization    `gorm:"constraint:Ondelete:CASCADE"`
}

var ProviderSearchFields = []string{"name"}

type Provider struct {
	global.BaseDbModel
	Name        string  `gorm:"unique;not null"`
	Icon        *string `json:"icon"`
	Description *string `json:"description"`
}
