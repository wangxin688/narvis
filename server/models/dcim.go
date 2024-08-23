package models

import (
	"github.com/wangxin688/narvis/server/global"
	"gorm.io/datatypes"
)

var SiteSearchFields = []string{"name", "site_code", "address"}
var LocationSearchFields = []string{"name"}
var DeviceSearchFields = []string{"name", "management_ip", "chassis_id", "serial_number", "asset_tag"}
var APSearchFields = []string{"name", "mac_address", "serial_number", "management_ip"}

type Site struct {
	global.BaseDbModel
	Name           string       `gorm:"uniqueIndex:idx_name_organization_id;not null"`
	SiteCode       string       `gorm:"uniqueIndex:idx_site_code_organization_id;not null"`
	Status         string       `gorm:"not null;default:Active"` // Active, Inactive
	Region         string       `gorm:"not null"`
	TimeZone       string       `gorm:"not null"`
	Latitude       string       `gorm:"not null"`
	Longitude      string       `gorm:"not null"`
	Address        string       `gorm:"not null"`
	Description    *string      `gorm:"default:null"`
	OrganizationId string       `gorm:"type:uuid;uniqueIndex:idx_name_organization_id;uniqueIndex:idx_site_code_organization_id;index;"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

type Location struct {
	global.BaseDbModel
	Name           string       `gorm:"uniqueIndex:idx_name_site_id;not null"`
	Description    *string      `gorm:"default:null"`
	ParentId       *string      `gorm:"type:uuid;default:null"`
	Parent         *Location    `gorm:"constraint:Ondelete:CASCADE;references:Id"`
	SiteId         string       `gorm:"type:uuid;uniqueIndex:idx_name_site_id;not null"`
	Site           Site         `gorm:"constraint:Ondelete:CASCADE"`
	OrganizationId string       `gorm:"type:uuid;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

type Rack struct {
	global.BaseDbModel
	Name           string       `gorm:"not null"`
	AssetTag       *string      `gorm:"column:asset_tag;uniqueIndex:idx_asset_tag_organization_id"`
	SerialNumber   *string      `gorm:"column:serial_number"`
	UHeight        uint8        `gorm:"type:smallint;default:42"`
	Height         float32      `gorm:"type:float;default:2"`
	Width          float32      `gorm:"type:float;default:0.6"`
	Depth          float32      `gorm:"type:float;default:0.8"`
	StartingUnit   uint8        `gorm:"type:float;default:1"`
	DescUnit       bool         `gorm:"default:true"`
	LocationId     *string      `gorm:"type:uuid;default:null"`
	Location       Location     `gorm:"constraint:Ondelete:SET NULL"`
	SiteId         string       `gorm:"type:uuid;not null"`
	Site           Site         `gorm:"constraint:Ondelete:RESTRICT"`
	OrganizationId string       `gorm:"index;not null;uniqueIndex:idx_asset_tag_organization_id"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

type Device struct {
	global.BaseDbModel

	Name           string       `gorm:"uniqueIndex:idx_name_organization_id;not null"`
	ManagementIp   string       `gorm:"uniqueIndex:idx_management_ip_organization_id;not null;index"`
	Status         string       `gorm:"default:Active"`
	Platform       string       `gorm:"default:Unknown"`
	ProductFamily  string       `gorm:"default:Unknown"`
	DeviceType     string       `gorm:"default:Unknown"`
	Manufacturer   string       `gorm:"default:Unknown"`
	DeviceRole     string       `gorm:"default:Unknown"`
	ChassisId      *string      `gorm:"uniqueIndex:idx_chassis_id_organization_id;default:null;index"`
	SerialNumber   *string      `gorm:"column:serial_number;uniqueIndex:idx_serial_number_organization_id"`
	AssetTag       *string      `gorm:"column:asset_tag"`
	Description    *string      `gorm:"column:description"`
	RackId         *string      `gorm:"type:uuid;default:null"`
	Rack           Rack         `gorm:"constraint:Ondelete:SET NULL"`
	RackPosition   *uint8       `gorm:"type:smallint;default:null"`
	UHeight        *uint8       `gorm:"type:smallint;default:null"`
	LocationId     *string      `gorm:"type:uuid;default:null"`
	Location       Location     `gorm:"constraint:Ondelete:SET NULL"`
	SiteId         string       `gorm:"type:uuid;index;not null"`
	Site           Site         `gorm:"constraint:Ondelete:RESTRICT"`
	OrganizationId string       `gorm:"type:uuid;uniqueIndex:idx_name_organization_id;uniqueIndex:idx_management_ip_organization_id;uniqueIndex:idx_serial_number_organization_id;uniqueIndex:idx_chassis_id_organization_id;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

type DeviceInterface struct {
	global.BaseDbModel
	IfName      string `gorm:"uniqueIndex:idx_if_name_device_id;not null"`
	IfIndex     uint64 `gorm:"uniqueIndex:idx_if_index_device_id;not null"`
	IfDescr     string `gorm:"default:null"`
	Speed       uint64 `gorm:"default:1000000"`
	Mode        uint64 `gorm:"default:1"`
	Mtu         uint64 `gorm:"column:mtu;default:1500"`
	AdminStatus uint64 `gorm:"default:1"`
	OperStatus  uint64 `gorm:"default:1"`
	LastChange  uint64 `gorm:"default:0"`
	DeviceId    string `gorm:"type:uuid;index;uniqueIndex:idx_if_name_device_id;uniqueIndex:idx_if_index_device_id"`
	Device      Device `gorm:"constraint:Ondelete:CASCADE"`
	SiteId      string `gorm:"type:uuid;index"`
	Site        Site   `gorm:"constraint:Ondelete:CASCADE"`
}

type DeviceStack struct {
	global.BaseDbModel
	Priority     uint8   `gorm:"type:smallint;uniqueIndex:idx_priority_device_id;default:0"`
	SerialNumber *string `gorm:"default:null"`
	MacAddress   string  `gorm:"not null;uniqueIndex:idx_mac_address_device_id"`
	DeviceId     string  `gorm:"type:uuid;uniqueIndex:idx_mac_address_device_id;uniqueIndex:idx_priority_device_id;not null"`
	Device       Device  `gorm:"constraint:Ondelete:CASCADE"`
}

type LLDPNeighbor struct {
	global.BaseDbModel

	SourceInterfaceId string          `gorm:"type:uuid;not null"`
	SourceInterface   DeviceInterface `gorm:"constraint:Ondelete:CASCADE;foreignKey:SourceInterfaceId"`
	SourceDeviceId    string          `gorm:"type:uuid;not null;index"`
	SourceDevice      Device          `gorm:"constraint:Ondelete:CASCADE;foreignKey:SourceDeviceId"`
	TargetInterfaceId string          `gorm:"type:uuid;not null"`
	TargetInterface   DeviceInterface `gorm:"constraint:Ondelete:CASCADE;foreignKey:TargetInterfaceId"`
	TargetDeviceId    string          `gorm:"type:uuid;not null"`
	TargetDevice      Device          `gorm:"constraint:Ondelete:CASCADE;foreignKey:TargetDeviceId"`
	Active            bool            `gorm:"default:true"`
	SiteId            string          `gorm:"type:uuid;not null;index"`
	Site              Site            `gorm:"constraint:Ondelete:CASCADE"`
	OrganizationId    string          `gorm:"type:uuid;not null;index"`
	Organization      Organization    `gorm:"constraint:Ondelete:CASCADE"`
}

type ApLLDPNeighbor struct {
	global.BaseDbModel

	SourceInterfaceId string          `gorm:"type:uuid;not null"`
	SourceInterface   DeviceInterface `gorm:"constraint:Ondelete:CASCADE;foreignKey:SourceInterfaceId"`
	SourceDeviceId    string          `gorm:"type:uuid;not null;index"`
	SourceDevice      Device          `gorm:"constraint:Ondelete:CASCADE;foreignKey:SourceDeviceId"`
	TargetApId        string          `gorm:"type:uuid;not null"`
	TargetAp          AP              `gorm:"constraint:Ondelete:CASCADE;foreignKey:TargetApId"`
	Active            bool            `gorm:"default:true"`
	SiteId            string          `gorm:"type:uuid;not null;index"`
	Site              Site            `gorm:"constraint:Ondelete:CASCADE"`
	OrganizationId    string          `gorm:"type:uuid;not null;index"`
	Organization      Organization    `gorm:"constraint:Ondelete:CASCADE"`
}

type DeviceConfig struct {
	global.BaseDbSingleModel

	Configuration string `gorm:"not null"`
	TotalLines    uint32 `gorm:"not null;default:0"`
	LinesAdded    uint32 `gorm:"not null;default:0"`
	LinesDeleted  uint32 `gorm:"not null;default:0"`
	Md5Checksum   string `gorm:"not null"`
	DeviceId      string `gorm:"type:uuid;index"`
	Device        Device `gorm:"constraint:Ondelete:CASCADE"`
}

type DeviceCliCredential struct {
	global.BaseDbModel
	Username       string       `gorm:"not null"`
	Password       string       `gorm:"not null"`
	DeviceId       *string      `gorm:"type:uuid;default:null;uniqueIndex:idx_device_id_organization_id;index"` // when device_id is null, the config is global
	Device         Device       `gorm:"constraint:Ondelete:CASCADE"`
	OrganizationID string       `gorm:"type:uuid;uniqueIndex:idx_device_id_organization_id;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

type DeviceSnmpV2Credential struct {
	global.BaseDbModel
	Community      string       `gorm:"not null"`
	MaxRepetitions uint8        `gorm:"type:smallint;not null;default:50"`
	Timeout        uint8        `gorm:"type:smallint;not null;default:10"`
	DeviceId       *string      `gorm:"type:uuid;default:null;uniqueIndex:idx_device_id_organization_id;index"` // when device_id is null, the config is global
	Device         Device       `gorm:"constraint:Ondelete:CASCADE"`
	OrganizationID string       `gorm:"type:uuid;uniqueIndex:idx_device_id_organization_id;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

type DeviceRestconfCredential struct {
	global.BaseDbModel
	Url            string       `gorm:"not null"`
	Username       string       `gorm:"not null"`
	Password       string       `gorm:"not null"`
	DeviceId       *string      `gorm:"type:uuid;default:null;uniqueIndex:idx_device_id_organization_id;index"` // when device_id is null, the config is global
	Device         Device       `gorm:"constraint:Ondelete:CASCADE"`
	OrganizationID string       `gorm:"type:uuid;uniqueIndex:idx_device_id_organization_id;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

type ApCoordinate struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

type AP struct {
	global.BaseDbModel
	Name           string                            `gorm:"not null;uniqueIndex:idx_name_site_id"`
	Status         string                            `gorm:"default:Active"`
	MacAddress     *string                           `gorm:"macaddr;default:null"`
	SerialNumber   *string                           `gorm:"column:serial_number;default:null"`
	ManagementIP   *string                           `gorm:"column:management_ip;default:null"`
	DeviceType     string                            `gorm:"default:Unknown"`
	Manufacturer   string                            `gorm:"default:Unknown"`
	DeviceRole     string                            `gorm:"default:WlanAP"`
	Version        *string                           `gorm:"default:null"`
	GroupName      *string                           `gorm:"default:null"`
	Coordinate     *datatypes.JSONType[ApCoordinate] `gorm:"type:json;default:null"`
	ActiveWacId    *string                           `gorm:"type:uuid;default:null"`
	ActiveWac      Device                            `gorm:"constraint:Ondelete:SET NULL"`
	LocationId     *string                           `gorm:"type:uuid;default:null"`
	Location       Location                          `gorm:"constraint:Ondelete:SET NULL"`
	SiteId         string                            `gorm:"type:uuid;uniqueIndex:idx_name_site_id;not null"`
	Site           Site                              `gorm:"constraint:Ondelete:RESTRICT"`
	OrganizationID string                            `gorm:"type:uuid;index"`
	Organization   Organization                      `gorm:"constraint:Ondelete:CASCADE"`
}
