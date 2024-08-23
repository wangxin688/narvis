package models

import "github.com/wangxin688/narvis/server/global"

var IpSearchFields = []string{"address"}
var BlockSearchFields = []string{"prefix"}
var PrefixSearchFields = []string{"prefix"}
var VlanSearchFields = []string{"name", "vid"}

var BlockTableName = "ipam_block"
var PrefixTableName = "ipam_prefix"
var IpAddressTableName = "ipam_ip_address"
var VlanTableName = "ipam_vlan"

type Block struct {
	global.BaseDbModel

	Prefix         string       `gorm:"type:cidr;not null"`
	Description    *string      `gorm:"default:null"`
	OrganizationID string       `gorm:"type:uuid;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

// method: family, child_prefixes, utilization

func (Block) TableName() string {
	return BlockTableName
}

type Prefix struct {
	Prefix         string       `gorm:"type:cidr;index;not null;uniqueIndex:idx_prefix_organization_id;index"`
	Version        string       `gorm:"not null"` // IPv4 or IPv6
	Description    *string      `gorm:"default:null"`
	Status         *string      `gorm:"default:Active"`
	IsPool         *bool        `gorm:"default:false"`
	MarkAsFull     *bool        `gorm:"default:false"`
	VlanID         *string      `gorm:"type:uuid;default:null"`
	Vlan           Vlan         `gorm:"constraint:Ondelete:SET NULL"`
	SiteID         *string      `gorm:"type:uuid;index"`
	Site           Site         `gorm:"constraint:Ondelete:SET NULL"`
	OrganizationID string       `gorm:"type:uuid;uniqueIndex:idx_prefix_organization_id;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
	// TODO: Add utilization
	// TODO: Add child prefixes
	// TODO: Add family
}

func (Prefix) TableName() string {
	return PrefixTableName
}

type IpAddress struct {
	global.BaseDbModel

	Address        string       `gorm:"type:inet;not null;uniqueIndex:idx_address_organization_id;index"`
	Status         string       `gorm:"default:Active"`
	DnsName        *string      `gorm:"default:null"`
	AssignTo       *string      `gorm:"default:null"`
	OrganizationID string       `gorm:"type:uuid;uniqueIndex:idx_address_organization_id;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (IpAddress) TableName() string {
	return IpAddressTableName
}

type Vlan struct {
	global.BaseDbModel
	Name           string       `gorm:"not null"`
	Vid            uint32       `gorm:"not null;uniqueIndex:idx_vid_site_id"` // 1-4094 and vxlan range
	Description    *string      `gorm:"default:null"`
	Status         string       `gorm:"default:Active"`
	SiteID         string       `gorm:"type:uuid;uniqueIndex:idx_vid_site_id;index"`
	Site           Site         `gorm:"constraint:Ondelete:CASCADE"`
	OrganizationID string       `gorm:"type:uuid;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (Vlan) TableName() string {
	return VlanTableName
}
