package models

import (
	"gorm.io/datatypes"
)

var SiteSearchFields = []string{"name", "site_code", "address"}
var LocationSearchFields = []string{"name"}
var DeviceSearchFields = []string{"name", "management_ip", "chassis_id", "serial_number", "asset_tag"}
var APSearchFields = []string{"name", "mac_address", "serial_number", "management_ip"}

var SiteTableName = "dcim_site"
var LocationTableName = "dcim_location"
var RackTableName = "dcim_rack"
var DeviceTableName = "dcim_device"
var APTableName = "dcim_ap"
var DeviceInterfaceTableName = "dcim_interface"
var LLDPNeighborTableName = "dcim_lldp_neighbor"
var ApLLDPNeighborTableName = "dcim_ap_lldp_neighbor"
var DeviceStackTableName = "dcim_device_stack"
var DeviceConfigTableName = "dcim_device_config"
var DeviceCliCredentialTableName = "dcim_device_cli_credential"
var DeviceSnmpV2CredentialTableName = "dcim_device_snmp_v2_credential"
var DeviceRestconfCredentialTableName = "dcim_device_restconf_credential"

type Site struct {
	BaseDbModel
	Name           string       `gorm:"uniqueIndex:idx_name_organization_id;not null"`
	SiteCode       string       `gorm:"uniqueIndex:idx_site_code_organization_id;not null"`
	Status         string       `gorm:"not null;default:Active"` // Active, Inactive
	Region         string       `gorm:"not null"`
	TimeZone       string       `gorm:"not null"`
	Latitude       string       `gorm:"not null"`
	Longitude      string       `gorm:"not null"`
	Address        string       `gorm:"not null"`
	Description    *string      `gorm:"default:null"`
	OrganizationID string       `gorm:"type:uuid;uniqueIndex:idx_name_organization_id;uniqueIndex:idx_site_code_organization_id;index;"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (Site) TableName() string {
	return SiteTableName
}

type Location struct {
	BaseDbModel
	Name           string       `gorm:"uniqueIndex:idx_name_site_id;not null"`
	Description    *string      `gorm:"default:null"`
	ParentID       *string      `gorm:"type:uuid;default:null"`
	Parent         *Location    `gorm:"constraint:Ondelete:CASCADE;references:ID"`
	SiteID         string       `gorm:"type:uuid;uniqueIndex:idx_name_site_id;not null"`
	Site           Site         `gorm:"constraint:Ondelete:CASCADE"`
	OrganizationID string       `gorm:"type:uuid;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (Location) TableName() string {
	return LocationTableName
}

type Rack struct {
	BaseDbModel
	Name           string       `gorm:"not null"`
	AssetTag       *string      `gorm:"column:asset_tag;uniqueIndex:idx_asset_tag_organization_id"`
	SerialNumber   *string      `gorm:"column:serial_number"`
	UHeight        uint8        `gorm:"type:smallint;default:42"`
	Height         float32      `gorm:"type:float;default:2"`
	Width          float32      `gorm:"type:float;default:0.6"`
	Depth          float32      `gorm:"type:float;default:0.8"`
	DescUnit       bool         `gorm:"default:true"`
	LocationID     *string      `gorm:"type:uuid;default:null"`
	Location       Location     `gorm:"constraint:Ondelete:SET NULL"`
	SiteID         string       `gorm:"type:uuid;not null"`
	Site           Site         `gorm:"constraint:Ondelete:RESTRICT"`
	OrganizationID string       `gorm:"index;not null;uniqueIndex:idx_asset_tag_organization_id"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (Rack) TableName() string {
	return RackTableName
}

type Device struct {
	BaseDbModel

	Name           string       `gorm:"uniqueIndex:idx_name_organization_id;not null"`
	ManagementIp   string       `gorm:"uniqueIndex:idx_management_ip_organization_id;not null;index"`
	Status         string       `gorm:"default:Active"`
	Platform       string       `gorm:"default:Unknown"`
	ProductFamily  string       `gorm:"default:Unknown"`
	DeviceModel    string       `gorm:"default:Unknown"`
	Manufacturer   string       `gorm:"default:Unknown"`
	DeviceRole     string       `gorm:"default:Unknown"`
	ChassisID      *string      `gorm:"uniqueIndex:idx_chassis_id_organization_id;default:null;index"`
	SerialNumber   *string      `gorm:"column:serial_number;uniqueIndex:idx_serial_number_organization_id"`
	AssetTag       *string      `gorm:"column:asset_tag"`
	Description    *string      `gorm:"column:description"`
	OsVersion      *string      `gorm:"column:os_version"`
	OsPatch        *string      `gorm:"column:os_patch"`
	RackID         *string      `gorm:"type:uuid;default:null"`
	Rack           Rack         `gorm:"constraint:Ondelete:SET NULL"`
	RackPosition   *uint8       `gorm:"type:smallint;default:null"`
	RackDirection  string       `gorm:"default:Front"` // Front, Rear
	UHeight        *uint8       `gorm:"type:smallint;default:null"`
	LocationID     *string      `gorm:"type:uuid;default:null"`
	Location       Location     `gorm:"constraint:Ondelete:SET NULL"`
	SiteID         string       `gorm:"type:uuid;index;not null"`
	Site           Site         `gorm:"constraint:Ondelete:RESTRICT"`
	OrganizationID string       `gorm:"type:uuid;uniqueIndex:idx_name_organization_id;uniqueIndex:idx_management_ip_organization_id;uniqueIndex:idx_serial_number_organization_id;uniqueIndex:idx_chassis_id_organization_id;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (Device) TableName() string {
	return DeviceTableName
}

type DeviceInterface struct {
	BaseDbModel
	IfName      string `gorm:"uniqueIndex:idx_if_name_device_id;not null"`
	IfIndex     uint64 `gorm:"uniqueIndex:idx_if_index_device_id;not null"`
	IfDescr     string `gorm:"default:null"`
	Speed       uint64 `gorm:"default:1000000"`
	Mode        uint64 `gorm:"default:1"`
	Mtu         uint64 `gorm:"column:mtu;default:1500"`
	AdminStatus uint64 `gorm:"default:1"`
	OperStatus  uint64 `gorm:"default:1"`
	LastChange  uint64 `gorm:"default:0"`
	DeviceID    string `gorm:"type:uuid;index;uniqueIndex:idx_if_name_device_id;uniqueIndex:idx_if_index_device_id"`
	Device      Device `gorm:"constraint:Ondelete:CASCADE"`
	SiteID      string `gorm:"type:uuid;index"`
	Site        Site   `gorm:"constraint:Ondelete:CASCADE"`
}

func (DeviceInterface) TableName() string {
	return DeviceInterfaceTableName
}

type DeviceStack struct {
	BaseDbModel
	Priority     uint8   `gorm:"type:smallint;uniqueIndex:idx_priority_device_id;default:0"`
	SerialNumber *string `gorm:"default:null"`
	MacAddress   string  `gorm:"not null;uniqueIndex:idx_mac_address_device_id"`
	DeviceID     string  `gorm:"type:uuid;uniqueIndex:idx_mac_address_device_id;uniqueIndex:idx_priority_device_id;not null"`
	Device       Device  `gorm:"constraint:Ondelete:CASCADE"`
}

func (DeviceStack) TableName() string {
	return DeviceStackTableName
}

type LLDPNeighbor struct {
	BaseDbModel

	SourceInterfaceID string          `gorm:"type:uuid;not null"`
	SourceInterface   DeviceInterface `gorm:"constraint:Ondelete:CASCADE;foreignKey:SourceInterfaceID"`
	SourceDeviceID    string          `gorm:"type:uuid;not null;index"`
	SourceDevice      Device          `gorm:"constraint:Ondelete:CASCADE;foreignKey:SourceDeviceID"`
	TargetInterfaceID string          `gorm:"type:uuid;not null"`
	TargetInterface   DeviceInterface `gorm:"constraint:Ondelete:CASCADE;foreignKey:TargetInterfaceID"`
	TargetDeviceID    string          `gorm:"type:uuid;not null"`
	TargetDevice      Device          `gorm:"constraint:Ondelete:CASCADE;foreignKey:TargetDeviceID"`
	Active            bool            `gorm:"default:true"`
	SiteID            string          `gorm:"type:uuid;not null;index"`
	Site              Site            `gorm:"constraint:Ondelete:CASCADE"`
	OrganizationID    string          `gorm:"type:uuid;not null;index"`
	Organization      Organization    `gorm:"constraint:Ondelete:CASCADE"`
}

func (LLDPNeighbor) TableName() string {
	return LLDPNeighborTableName
}

type ApLLDPNeighbor struct {
	BaseDbModel

	SourceInterfaceID string          `gorm:"type:uuid;not null"`
	SourceInterface   DeviceInterface `gorm:"constraint:Ondelete:CASCADE;foreignKey:SourceInterfaceID"`
	SourceDeviceID    string          `gorm:"type:uuid;not null;index"`
	SourceDevice      Device          `gorm:"constraint:Ondelete:CASCADE;foreignKey:SourceDeviceID"`
	TargetApID        string          `gorm:"type:uuid;not null"`
	TargetAp          AP              `gorm:"constraint:Ondelete:CASCADE;foreignKey:TargetApID"`
	Active            bool            `gorm:"default:true"`
	SiteID            string          `gorm:"type:uuid;not null;index"`
	Site              Site            `gorm:"constraint:Ondelete:CASCADE"`
	OrganizationID    string          `gorm:"type:uuid;not null;index"`
	Organization      Organization    `gorm:"constraint:Ondelete:CASCADE"`
}

func (ApLLDPNeighbor) TableName() string {
	return ApLLDPNeighborTableName
}

type DeviceConfig struct {
	BaseDbSingleModel

	Configuration string `gorm:"not null"`
	TotalLines    uint32 `gorm:"not null;default:0"`
	LinesAdded    uint32 `gorm:"not null;default:0"`
	LinesDeleted  uint32 `gorm:"not null;default:0"`
	Md5Checksum   string `gorm:"not null"`
	DeviceID      string `gorm:"type:uuid;index"`
	Device        Device `gorm:"constraint:Ondelete:CASCADE"`
}

func (DeviceConfig) TableName() string {
	return DeviceConfigTableName
}

type DeviceCliCredential struct {
	BaseDbModel
	Username       string       `gorm:"not null"`
	Password       string       `gorm:"not null"`
	Port           uint16       `gorm:"not null;default:22"`
	DeviceID       *string      `gorm:"type:uuid;default:null;uniqueIndex:idx_device_id_organization_id;index"` // when device_id is null, the config is global
	Device         Device       `gorm:"constraint:Ondelete:CASCADE"`
	OrganizationID string       `gorm:"type:uuid;uniqueIndex:idx_device_id_organization_id;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (DeviceCliCredential) TableName() string {
	return DeviceCliCredentialTableName
}

type DeviceSnmpV2Credential struct {
	BaseDbModel
	Community      string       `gorm:"not null"`
	MaxRepetitions uint8        `gorm:"type:smallint;not null;default:50"`
	Timeout        uint8        `gorm:"type:smallint;not null;default:10"`
	Port           uint16       `gorm:"not null;default:161"`
	DeviceID       *string      `gorm:"type:uuid;default:null;uniqueIndex:idx_device_id_organization_id;index"` // when device_id is null, the config is global
	Device         Device       `gorm:"constraint:Ondelete:CASCADE"`
	OrganizationID string       `gorm:"type:uuid;uniqueIndex:idx_device_id_organization_id;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (DeviceSnmpV2Credential) TableName() string {
	return DeviceSnmpV2CredentialTableName
}

type DeviceRestconfCredential struct {
	BaseDbModel
	Url            string       `gorm:"not null"`
	Username       string       `gorm:"not null"`
	Password       string       `gorm:"not null"`
	DeviceID       *string      `gorm:"type:uuid;default:null;uniqueIndex:idx_device_id_organization_id;index"` // when device_id is null, the config is global
	Device         Device       `gorm:"constraint:Ondelete:CASCADE"`
	OrganizationID string       `gorm:"type:uuid;uniqueIndex:idx_device_id_organization_id;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (DeviceRestconfCredential) TableName() string {
	return DeviceRestconfCredentialTableName
}

type ApCoordinate struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

type AP struct {
	BaseDbModel
	Name           string                            `gorm:"not null;uniqueIndex:idx_name_site_id"`
	Status         string                            `gorm:"default:Active"`
	MacAddress     *string                           `gorm:"macaddr;default:null"`
	SerialNumber   *string                           `gorm:"column:serial_number;default:null"`
	ManagementIP   *string                           `gorm:"column:management_ip;default:null"`
	DeviceModel    string                            `gorm:"default:Unknown"`
	Manufacturer   string                            `gorm:"default:Unknown"`
	DeviceRole     string                            `gorm:"default:WlanAP"`
	Version        *string                           `gorm:"default:null"`
	GroupName      *string                           `gorm:"default:null"`
	Coordinate     *datatypes.JSONType[ApCoordinate] `gorm:"type:json;default:null"`
	ActiveWacID    *string                           `gorm:"type:uuid;default:null"`
	ActiveWac      Device                            `gorm:"constraint:Ondelete:SET NULL"`
	LocationID     *string                           `gorm:"type:uuid;default:null"`
	Location       Location                          `gorm:"constraint:Ondelete:SET NULL"`
	SiteID         string                            `gorm:"type:uuid;uniqueIndex:idx_name_site_id;not null"`
	Site           Site                              `gorm:"constraint:Ondelete:RESTRICT"`
	OrganizationID string                            `gorm:"type:uuid;index"`
	Organization   Organization                      `gorm:"constraint:Ondelete:CASCADE"`
}

func (AP) TableName() string {
	return APTableName
}

type MacAddress struct {
	ID        string `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	MacOUI    string `gorm:"column:mac_oui;not null"`
	ShortName string `gorm:"column:short_name;not null"`
	LongName  string `gorm:"column:long_name;not null"`
}
