package models

var CircuitSearchFields = []string{"name", "ip_address"}
var ProviderSearchFields = []string{"name"}

var CircuitTableName = "circuit"
var ProviderTableName = "provider"

type Circuit struct {
	BaseDbModel

	Name           string          `gorm:"column:name;uniqueIndex:idx_name_organization_id;not null"`
	CId            string          `gorm:"column:cId;uniqueIndex:idx_circuit_id_organization_id;not null"`
	Status         string          `gorm:"column:status;not null;default:Active"`        // Active/Inactive
	CircuitType    string          `gorm:"column:circuitType;not null;default:Internet"` // Internet/Intranet
	BandWidth      uint32          `gorm:"column:bandWidth;not null"`                    //Mbps
	IpAddress      *string         `gorm:"column:ipAddress;default:null"`
	Description    *string         `gorm:"column:description;default:null"`
	ProviderId     string          `gorm:"column:providerId;type:uuid;not null"`
	Provider       Provider        `gorm:"constraint:Ondelete:RESTRICT"`
	ASiteId        string          `gorm:"column:aSiteId;type:uuid;not null"`
	ASite          Site            `gorm:"constraint:Ondelete:RESTRICT;foreignKey:ASiteId"`
	ADeviceId      string          `gorm:"column:aDeviceId;type:uuid;not null"`
	ADevice        Device          `gorm:"constraint:Ondelete:RESTRICT;foreignKey:ADeviceId"`
	AInterfaceId   string          `gorm:"column:aInterfaceId;type:uuid;not null"`
	AInterface     DeviceInterface `gorm:"constraint:Ondelete:RESTRICT;foreignKey:AInterfaceId"`
	ZSiteId        string          `gorm:"column:zSiteId;type:uuid;not null"`
	ZSite          Site            `gorm:"constraint:Ondelete:RESTRICT;foreignKey:ZSiteId"`
	ZDeviceId      string          `gorm:"column:zDeviceId;type:uuid;not null"`
	ZDevice        Device          `gorm:"constraint:Ondelete:RESTRICT;foreignKey:ZDeviceId"`
	ZInterfaceId   string          `gorm:"column:zInterfaceId;type:uuid;not null"`
	ZInterface     DeviceInterface `gorm:"constraint:Ondelete:RESTRICT;foreignKey:ZInterfaceId"`
	MonitorId      *string         `gorm:"column:monitorId;default:null"`
	OrganizationId string          `gorm:"column:organizationId;type:uuid;not null"`
	Organization   Organization    `gorm:"constraint:Ondelete:CASCADE"`
}

func (Circuit) TableName() string {
	return CircuitTableName
}

type Provider struct {
	BaseDbModel
	Name        string  `gorm:"column:name;unique;not null"`
	Icon        *string `gorm:"column:icon;default:null"`
	Description *string `gorm:"column:description;default:null"`
}

func (Provider) TableName() string {
	return ProviderTableName
}
