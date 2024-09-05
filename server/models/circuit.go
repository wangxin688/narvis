package models

var CircuitSearchFields = []string{"name", "Ipv4Address", "Ipv6Address"}

var CircuitTableName = "circuit"
var ProviderTableName = "provider"

type Circuit struct {
	BaseDbModel

	Name            string          `gorm:"column:name;uniqueIndex:idx_name_site_id;not null"`
	CId             *string         `gorm:"column:cId;default:null"`
	Status          string          `gorm:"column:status;not null;default:Active"` // Active/Inactive
	CircuitType     string          `gorm:"column:circuitType;not null;default:Internet"`
	RxBandWidth     uint32          `gorm:"column:rxBandWidth;not null"` //Mbps
	TxBandWidth     uint32          `gorm:"column:txBandWidth;not null"`
	Ipv4Address     *string         `gorm:"column:ipv4Address;default:null"`
	Ipv6Address     *string         `gorm:"column:ipv6Address;default:null"`
	Description     *string         `gorm:"column:description;default:null"`
	Provider        string          `gorm:"column:provider;not null;default:Unknown"`
	SiteId          string          `gorm:"column:siteId;type:uuid;not null"`
	Site            Site            `gorm:"constraint:Ondelete:RESTRICT;foreignKey:siteId"`
	DeviceId        string          `gorm:"column:deviceId;type:uuid;not null"`
	Device          Device          `gorm:"constraint:Ondelete:RESTRICT;foreignKey:deviceId"`
	InterfaceId     string          `gorm:"column:interfaceId;type:uuid;not null"`
	DeviceInterface DeviceInterface `gorm:"constraint:Ondelete:RESTRICT;foreignKey:interfaceId"`
	MonitorId       *string         `gorm:"column:monitorId;default:null"`
	OrganizationId  string          `gorm:"column:organizationId;uniqueIndex:idx_name_site_id;type:uuid;not null"`
	Organization    Organization    `gorm:"constraint:Ondelete:CASCADE"`
}

func (Circuit) TableName() string {
	return CircuitTableName
}
