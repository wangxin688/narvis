package models

import (
	"gorm.io/datatypes"
)

var SiteSearchFields = []string{"name", "siteCode", "address"}
var DeviceSearchFields = []string{"name", "managementIp", "chassisId", "serialNumber", "assetTag"}
var APSearchFields = []string{"name", "macAddress", "serialNumber", "managementIp"}

var SiteTableName = "dcim_site"
var RackTableName = "dcim_rack"
var DeviceTableName = "dcim_device"
var APTableName = "dcim_ap"
var DeviceInterfaceTableName = "dcim_interface"
var LLDPNeighborTableName = "dcim_lldp_neighbor"
var ApLLDPNeighborTableName = "dcim_ap_lldp_neighbor"
var DeviceStackTableName = "dcim_device_stack"
var DeviceConfigTableName = "dcim_device_config"
var CliCredentialTableName = "dcim_cli_credential"
var SnmpV2CredentialTableName = "dcim_snmp_v2_credential"
var RestconfCredentialTableName = "dcim_restconf_credential"
var MacAddressTableName = "mac_address"

type Site struct {
	BaseDbModel
	Name           string       `gorm:"column:name;uniqueIndex:idx_name_organization_id;not null"`
	SiteCode       string       `gorm:"column:siteCode;uniqueIndex:idx_site_code_organization_id;not null"`
	Status         string       `gorm:"column:status;not null;default:Active"` // Active, Inactive
	Region         string       `gorm:"column:region;not null"`
	TimeZone       string       `gorm:"column:timeZone;not null"`
	Latitude       float32      `gorm:"column:latitude;not null"`
	Longitude      float32      `gorm:"column:longitude;not null"`
	Address        string       `gorm:"column:address;not null"`
	Description    *string      `gorm:"column:description;default:null"`
	MonitorId      *string      `gorm:"column:monitorId;default:null;unique"`
	OrganizationId string       `gorm:"column:organizationId;type:uuid;uniqueIndex:idx_name_organization_id;uniqueIndex:idx_site_code_organization_id;index;"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (Site) TableName() string {
	return SiteTableName
}

type Rack struct {
	BaseDbModel
	Name           string       `gorm:"column:name;not null"`
	SerialNumber   *string      `gorm:"column:serialNumber"`
	UHeight        uint8        `gorm:"column:uHeight;type:smallint;default:42"`
	DescUnit       bool         `gorm:"column:descUnit;default:true"`
	SiteId         string       `gorm:"column:siteId;type:uuid;not null"`
	Site           Site         `gorm:"constraint:Ondelete:RESTRICT"`
	OrganizationId string       `gorm:"column:organizationId;type:uuid;index;not null"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (Rack) TableName() string {
	return RackTableName
}

type Device struct {
	BaseDbModel

	Name           string                       `gorm:"column:name;not null"`
	ManagementIp   string                       `gorm:"column:managementIp;uniqueIndex:idx_management_ip_organization_id;not null;index"`
	Status         string                       `gorm:"column:status;default:Active"`
	Platform       string                       `gorm:"column:platform;default:Unknown"`
	ProductFamily  string                       `gorm:"column:productFamily;default:Unknown"`
	DeviceModel    string                       `gorm:"column:deviceModel;default:Unknown"`
	Manufacturer   string                       `gorm:"column:manufacturer;default:Unknown"`
	DeviceRole     string                       `gorm:"column:deviceRole;default:Unknown"`
	Floor          *string                      `gorm:"column:floor;default:null"`
	IsRegistered   bool                         `gorm:"column:isRegistered;type:bool;default:false"`
	ChassisId      *string                      `gorm:"column:chassisId;uniqueIndex:idx_chassis_id_organization_id;default:null;index"`
	SerialNumber   *string                      `gorm:"column:serialNumber;uniqueIndex:idx_serial_number_organization_id"`
	Description    *string                      `gorm:"column:description;default:null"`
	OsVersion      *string                      `gorm:"column:osVersion;default:null"`
	OsPatch        *string                      `gorm:"column:osPatch;default:null"`
	RackId         *string                      `gorm:"column:rackId;type:uuid;default:null"`
	Rack           Rack                         `gorm:"constraint:Ondelete:SET NULL"`
	RackPosition   *datatypes.JSONSlice[string] `gorm:"column:rackPosition;type:json;default:null"`
	MonitorId      *string                      `gorm:"column:monitorId;default:null;unique"`
	TemplateId     *string                      `gorm:"column:templateId;type:uuid;default:null"`
	Template       Template                     `gorm:"constraint:Ondelete:SET NULL"`
	SiteId         string                       `gorm:"column:siteId;type:uuid;index;not null"`
	Site           Site                         `gorm:"constraint:Ondelete:RESTRICT"`
	OrganizationId string                       `gorm:"column:organizationId;type:uuid;uniqueIndex:idx_management_ip_organization_id;uniqueIndex:idx_serial_number_organization_id;uniqueIndex:idx_chassis_id_organization_id;index"`
	Organization   Organization                 `gorm:"constraint:Ondelete:CASCADE"`
}

func (Device) TableName() string {
	return DeviceTableName
}

type DeviceInterface struct {
	BaseDbModel
	IfName        string  `gorm:"column:ifName;uniqueIndex:idx_if_name_device_id;not null"`
	IfIndex       uint64  `gorm:"column:ifIndex;uniqueIndex:idx_if_index_device_id;not null"`
	IfDescr       string  `gorm:"column:ifDescr;default:null"`
	IfSpeed       uint64  `gorm:"column:ifSpeed;default:1000"`
	IfType        uint64  `gorm:"column:ifType;default:1"`
	IfMtu         uint64  `gorm:"column:mtu;default:1500"`
	IfAdminStatus uint64  `gorm:"column:ifAdminStatus;default:1"`
	IfOperStatus  uint64  `gorm:"column:ifOperStatus;default:1"`
	IfLastChange  uint64  `gorm:"column:ifLastChange;default:0"`
	IfHighSpeed   uint64  `gorm:"column:ifHighSpeed;default:1000"`
	IfPhyAddr     *string `gorm:"column:ifPhyAddr;default:null"`
	IfIpAddress   *string `gorm:"column:ifIpAddress;default:null"`
	DeviceId      string  `gorm:"column:deviceId;type:uuid;index;uniqueIndex:idx_if_name_device_id;uniqueIndex:idx_if_index_device_id"`
	Device        Device  `gorm:"constraint:Ondelete:CASCADE"`
	SiteId        string  `gorm:"column:siteId;type:uuid;index"`
	Site          Site    `gorm:"constraint:Ondelete:CASCADE"`
}

func (DeviceInterface) TableName() string {
	return DeviceInterfaceTableName
}

type DeviceStack struct {
	BaseDbModel
	Priority     uint8   `gorm:"column:priority;type:smallint;uniqueIndex:idx_priority_device_id;default:0"`
	SerialNumber *string `gorm:"column:serialNumber;default:null"`
	MacAddress   string  `gorm:"column:macAddress;not null;uniqueIndex:idx_mac_address_device_id"`
	DeviceId     string  `gorm:"column:deviceId;type:uuid;uniqueIndex:idx_mac_address_device_id;uniqueIndex:idx_priority_device_id;not null"`
	Device       Device  `gorm:"constraint:Ondelete:CASCADE"`
}

func (DeviceStack) TableName() string {
	return DeviceStackTableName
}

type LLDPNeighbor struct {
	BaseDbModel

	LocalDeviceId  string       `gorm:"column:localDeviceId;type:uuid;not null;index"`
	LocalDevice    Device       `gorm:"constraint:Ondelete:CASCADE;foreignKey:LocalDeviceId"`
	LocalIfName    string       `gorm:"column:localIfName;not null"`
	LocalIfDescr   string       `gorm:"column:localIfDescr;not null"`
	RemoteDeviceId string       `gorm:"column:remoteDeviceId;type:uuid;not null"`
	RemoteDevice   Device       `gorm:"constraint:Ondelete:CASCADE;foreignKey:RemoteDeviceId"`
	RemoteIfName   string       `gorm:"column:remoteIfName;not null"`
	RemoteIfDescr  string       `gorm:"column:remoteIfDescr;not null"`
	HashValue      string       `gorm:"column:hashValue;unique;not null"` // MD5 hash for the local and neighbor
	SiteId         string       `gorm:"type:uuid;not null;index"`
	Site           Site         `gorm:"constraint:Ondelete:CASCADE"`
	OrganizationId string       `gorm:"type:uuid;not null;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (LLDPNeighbor) TableName() string {
	return LLDPNeighborTableName
}

type ApLLDPNeighbor struct {
	BaseDbModel

	LocalDeviceId  string       `gorm:"column:localDeviceId;type:uuid;not null;index"`
	LocalDevice    Device       `gorm:"constraint:Ondelete:CASCADE;foreignKey:LocalDeviceId"`
	LocalIfName    string       `gorm:"column:localIfName;not null"`
	LocalIfDescr   string       `gorm:"column:localIfDescr;not null"`
	RemoteApId     string       `gorm:"column:remoteApId;type:uuid;not null"`
	RemoteAp       AP           `gorm:"constraint:Ondelete:CASCADE;foreignKey:RemoteApId"`
	HashValue      string       `gorm:"column:hashValue;unique;not null"` // MD5 hash for the local and neighbor
	SiteId         string       `gorm:"column:siteId;type:uuid;not null;index"`
	Site           Site         `gorm:"constraint:Ondelete:CASCADE"`
	OrganizationId string       `gorm:"column:organizationId;type:uuid;not null;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (ApLLDPNeighbor) TableName() string {
	return ApLLDPNeighborTableName
}

type DeviceConfig struct {
	BaseDbSingleModel

	Configuration string `gorm:"column:configuration;not null"`
	TotalLines    uint32 `gorm:"column:totalLines;not null;default:0"`
	LinesAdded    uint32 `gorm:"column:linesAdded;not null;default:0"`
	LinesDeleted  uint32 `gorm:"column:linesDeleted;not null;default:0"`
	Md5Checksum   string `gorm:"column:md5Checksum;not null"`
	DeviceId      string `gorm:"column:deviceId;type:uuid;index"`
	Device        Device `gorm:"constraint:Ondelete:CASCADE"`
}

func (DeviceConfig) TableName() string {
	return DeviceConfigTableName
}

type CliCredential struct {
	BaseDbModel
	Username       string       `gorm:"column:username;not null"`
	Password       string       `gorm:"column:password;not null"`
	Port           uint16       `gorm:"column:port;not null;default:22"`
	DeviceId       *string      `gorm:"column:deviceId;type:uuid;default:null;uniqueIndex:idx_device_id_organization_id;index"` // when device_id is null, the config is global
	Device         Device       `gorm:"constraint:Ondelete:CASCADE"`
	OrganizationId string       `gorm:"column:organizationId;type:uuid;uniqueIndex:idx_device_id_organization_id;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (CliCredential) TableName() string {
	return CliCredentialTableName
}

type SnmpV2Credential struct {
	BaseDbModel
	Community      string       `gorm:"column:community;not null"`
	MaxRepetitions uint8        `gorm:"column:maxRepetitions;type:smallint;not null;default:50"`
	Timeout        uint8        `gorm:"column:timeout;type:smallint;not null;default:10"`
	Port           uint16       `gorm:"column:port;not null;default:161"`
	DeviceId       *string      `gorm:"column:deviceId;type:uuid;default:null;uniqueIndex:idx_device_id_organization_id;index"` // when device_id is null, the config is global
	Device         Device       `gorm:"constraint:Ondelete:CASCADE"`
	OrganizationId string       `gorm:"column:organizationId;type:uuid;uniqueIndex:idx_device_id_organization_id;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (SnmpV2Credential) TableName() string {
	return SnmpV2CredentialTableName
}

type RestconfCredential struct {
	BaseDbModel
	Url            string       `gorm:"column:url;not null"`
	Username       string       `gorm:"column:username;not null"`
	Password       string       `gorm:"column:password;not null"`
	DeviceId       *string      `gorm:"column:deviceId;type:uuid;default:null;uniqueIndex:idx_device_id_organization_id;index"` // when device_id is null, the config is global
	Device         Device       `gorm:"constraint:Ondelete:CASCADE"`
	OrganizationId string       `gorm:"column:organizationId;type:uuid;uniqueIndex:idx_device_id_organization_id;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (RestconfCredential) TableName() string {
	return RestconfCredentialTableName
}

type ApCoordinate struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

type AP struct {
	BaseDbModel
	Name           string                            `gorm:"column:name;not null;uniqueIndex:idx_name_site_id"`
	Status         string                            `gorm:"column:status;default:Active"`
	MacAddress     *string                           `gorm:"column:macAddress;type:macaddr;default:null"`
	SerialNumber   *string                           `gorm:"column:serialNumber;default:null"`
	ManagementIp   string                            `gorm:"column:managementIp;not null;uniqueIndex:idx_management_ip_site_id"`
	DeviceModel    string                            `gorm:"column:deviceModel;default:Unknown"`
	Manufacturer   string                            `gorm:"column:manufacturer;default:Unknown"`
	DeviceRole     string                            `gorm:"column:deviceRole;default:WlanAP"`
	OsVersion      *string                           `gorm:"column:osVersion;default:null"`
	GroupName      *string                           `gorm:"column:groupName;default:null"`
	Coordinate     *datatypes.JSONType[ApCoordinate] `gorm:"column:coordinate;type:json;default:null"`
	ActiveWacId    *string                           `gorm:"column:activeWacId;type:uuid;default:null"`
	ActiveWac      Device                            `gorm:"constraint:Ondelete:SET NULL"`
	Floor          *string                           `gorm:"column:floor;default:null"`
	SiteId         string                            `gorm:"column:siteId;type:uuid;uniqueIndex:idx_name_site_id;uniqueIndex:idx_management_ip_site_id;not null"`
	Site           Site                              `gorm:"constraint:Ondelete:RESTRICT"`
	OrganizationId string                            `gorm:"column:organizationId;type:uuid;index"`
	Organization   Organization                      `gorm:"constraint:Ondelete:CASCADE"`
}

func (AP) TableName() string {
	return APTableName
}

type MacAddress struct {
	Id        string `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	MacOUI    string `gorm:"column:macOui;not null" json:"macOui"`
	ShortName string `gorm:"column:shortName;not null" json:"shortName"`
	LongName  string `gorm:"column:longName;not null" json:"longName"`
}

func (MacAddress) TableName() string {
	return MacAddressTableName
}
