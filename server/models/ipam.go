package models

var IpSearchFields = []string{"address"}
var PrefixSearchFields = []string{"prefix", "vlanName"}

var PrefixTableName = "ipam_prefix"
var IpAddressTableName = "ipam_ip_address"

type Prefix struct {
	BaseDbModel

	Range          string       `gorm:"column:range;type:cidr;index;not null"`
	Version        string       `gorm:"column:version;not null"`     // IPv4 or IPv6
	Type           string       `gorm:"column:type;default:Dynamic"` // Dynamic or Static
	VlanId         *uint32      `gorm:"column:vlanId;default:null;index"`
	VlanName       *string      `gorm:"column:vlanName;default:null"`
	Gateway        *string      `gorm:"column:gateway;default:null"`
	SiteId         string       `gorm:"column:siteId;type:uuid;index;index"`
	Site           Site         `gorm:"constraint:Ondelete:CASCADE"`
	OrganizationId string       `gorm:"column:organizationId;type:uuid;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
	// TODO: Add utilization
	// TODO: Add child prefixes
	// TODO: Add family
}

func (Prefix) TableName() string {
	return PrefixTableName
}

type IpAddress struct {
	BaseDbModel

	Address        string       `gorm:"column:address;type:inet;not null;uniqueIndex:idx_ip_address_site_id;index"`
	Status         string       `gorm:"column:status;default:Active"` // Active or Reserved
	MacAddress     *string      `gorm:"column:macAddress;type:macaddr"`
	Vlan           *uint32      `gorm:"column:vlan;default:null"`
	Range          *string      `gorm:"column:range;type:cidr;default:null;index"`
	Description    *string      `gorm:"column:description;default:null"`
	Type           string       `gorm:"column:type;default:Dynamic"` // Dynamic or Static or Gateway or Broadcast or NetworkId
	SiteId         string       `gorm:"column:siteId;type:uuid;uniqueIndex:idx_ip_address_site_id;index"`
	Site           Site         `gorm:"constraint:Ondelete:CASCADE"`
	OrganizationId string       `gorm:"column:organizationId;type:uuid;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (IpAddress) TableName() string {
	return IpAddressTableName
}
