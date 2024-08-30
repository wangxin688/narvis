package models

var IpSearchFields = []string{"address"}
var BlockSearchFields = []string{"prefix"}
var PrefixSearchFields = []string{"prefix"}
var VlanSearchFields = []string{"name", "vid"}

var BlockTableName = "ipam_block"
var PrefixTableName = "ipam_prefix"
var IpAddressTableName = "ipam_ip_address"
var VlanTableName = "ipam_vlan"

type Block struct {
	BaseDbModel

	Prefix         string       `gorm:"column:prefix;type:cidr;not null"`
	Description    *string      `gorm:"column:description;default:null"`
	OrganizationId string       `gorm:"column:organizationId;type:uuid;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

// method: family, child_prefixes, utilization

func (Block) TableName() string {
	return BlockTableName
}

type Prefix struct {
	Prefix         string       `gorm:"column:prefix;type:cidr;index;not null;uniqueIndex:idx_prefix_organization_id;index"`
	Version        string       `gorm:"column:version;not null"` // IPv4 or IPv6
	Description    *string      `gorm:"column:description;default:null"`
	Status         *string      `gorm:"column:status;default:Active"`
	IsPool         *bool        `gorm:"column:isPool;default:false"`
	MarkAsFull     *bool        `gorm:"column:markAsFull;default:false"`
	VlanId         *string      `gorm:"column:VlanId;type:uuid;default:null"`
	Vlan           Vlan         `gorm:"constraint:Ondelete:SET NULL"`
	SiteId         *string      `gorm:"column:siteId;type:uuid;index"`
	Site           Site         `gorm:"constraint:Ondelete:SET NULL"`
	OrganizationId string       `gorm:"column:organizationId;type:uuid;uniqueIndex:idx_prefix_organization_id;index"`
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

	Address        string       `gorm:"column:address;type:inet;not null;uniqueIndex:idx_address_organization_id;index"`
	Status         string       `gorm:"column:status;default:Active"`
	DnsName        *string      `gorm:"column:dnsName;default:null"`
	AssignTo       *string      `gorm:"column:assignTo;default:null"`
	OrganizationId string       `gorm:"column:organizationId;type:uuid;uniqueIndex:idx_address_organization_id;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (IpAddress) TableName() string {
	return IpAddressTableName
}

type Vlan struct {
	BaseDbModel
	Name           string       `gorm:"column:name;not null"`
	Vid            uint32       `gorm:"column:vid;not null;uniqueIndex:idx_vid_site_id"` // 1-4094 and vxlan range
	Description    *string      `gorm:"column:description;default:null"`
	Status         string       `gorm:"column:status;default:Active"`
	SiteId         string       `gorm:"column:siteId;type:uuid;uniqueIndex:idx_vid_site_id;index"`
	Site           Site         `gorm:"constraint:Ondelete:CASCADE"`
	OrganizationId string       `gorm:"column:organizationId;type:uuid;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (Vlan) TableName() string {
	return VlanTableName
}
