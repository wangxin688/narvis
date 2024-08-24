package models

var CircuitSearchFields = []string{"name", "ip_address"}
var ProviderSearchFields = []string{"name"}

var CircuitTableName = "circuit"
var ProviderTableName = "provider"

type Circuit struct {
	BaseDbModel

	Name           string          `gorm:"uniqueIndex:idx_name_organization_id;not null"`
	CID            string          `gorm:"column:c_id;uniqueIndex:idx_circuit_id_organization_id;not null"`
	Status         string          `gorm:"not null;default:Active"`   // Active/Inactive
	CircuitType    string          `gorm:"not null;default:Internet"` // Internet/Intranet
	BandWidth      uint32          `gorm:"not null"`                  //Mbps
	IpAddress      string          `gorm:"not null"`
	Description    *string         `gorm:"column:description;default:null"`
	ProviderID     string          `gorm:"type:uuid;not null"`
	Provider       Provider        `gorm:"constraint:Ondelete:RESTRICT"`
	ASiteID        string          `gorm:"type:uuid;not null"`
	ASite          Site            `gorm:"constraint:Ondelete:RESTRICT;foreignKey:ASiteID"`
	ADeviceID      string          `gorm:"type:uuid;not null"`
	ADevice        Device          `gorm:"constraint:Ondelete:RESTRICT;foreignKey:ADeviceID"`
	AInterfaceID   string          `gorm:"type:uuid;not null"`
	AInterface     DeviceInterface `gorm:"constraint:Ondelete:RESTRICT;foreignKey:AInterfaceID"`
	ZSiteID        string          `gorm:"type:uuid;not null"`
	ZSite          Site            `gorm:"constraint:Ondelete:RESTRICT;foreignKey:ZSiteID"`
	ZDeviceID      string          `gorm:"type:uuid;not null"`
	ZDevice        Device          `gorm:"constraint:Ondelete:RESTRICT;foreignKey:ZDeviceID"`
	ZInterfaceID   string          `gorm:"type:uuid;not null"`
	ZInterface     DeviceInterface `gorm:"constraint:Ondelete:RESTRICT;foreignKey:ZInterfaceID"`
	OrganizationID string          `gorm:"type:uuid;not null"`
	Organization   Organization    `gorm:"constraint:Ondelete:CASCADE"`
}

func (Circuit) TableName() string {
	return CircuitTableName
}

type Provider struct {
	BaseDbModel
	Name        string  `gorm:"unique;not null"`
	Icon        *string `json:"icon"`
	Description *string `json:"description"`
}

func (Provider) TableName() string {
	return ProviderTableName
}
