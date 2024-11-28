package models

var ServerTableName = "infra_server"
var ServerCredentialTableName = "server_credential"
var ServerSnmpCredentialTableName = "server_snmp_credential"

type Server struct {
	BaseDbModel
	Name           string       `gorm:"column:name;not null"`
	ManagementIp   string       `gorm:"column:managementIp;not null"`
	Manufacturer   string       `gorm:"column:manufacturer;not null"`
	Status         string       `gorm:"column:status;default:Active"`
	OsVersion      string       `gorm:"column:osVersion;not null"`
	RackId         *string      `gorm:"column:rackId;type:uuid;default:null;index"`
	RackPosition   *string      `gorm:"column:rackPosition;default:null"`
	Rack           Rack         `gorm:"constraint:Ondelete:SET NULL"`
	Cpu            uint8        `gorm:"column:Cpu;not null"`
	Memory         uint64       `gorm:"column:memory;not null"`
	Disk           uint64       `gorm:"column:disk;not null"`
	Description    *string      `gorm:"column:description;default:null"`
	MonitorId      *string      `gorm:"column:monitorId;default:null;unique"`
	TemplateId     *string      `gorm:"column:templateId;type:uuid;default:null"`
	Template       Template     `gorm:"constraint:Ondelete:SET NULL"`
	SiteId         string       `gorm:"column:siteId;type:uuid;index;not null"`
	Site           Site         `gorm:"constraint:Ondelete:RESTRICT"`
	OrganizationId string       `gorm:"column:organizationId;type:uuid;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (Server) TableName() string {
	return ServerTableName
}

type ServerCredential struct {
	BaseDbModel
	Username       string       `gorm:"column:username;not null"`
	Password       string       `gorm:"column:password;not null"`
	Port           uint16       `gorm:"column:port;not null;default:22"`
	ServerId       *string      `gorm:"column:serverId;type:uuid;default:null;uniqueIndex:idx_server_id_organization_id;index"` // when device_id is null, the config is global
	Server         Server       `gorm:"constraint:Ondelete:CASCADE"`
	OrganizationId string       `gorm:"column:organizationId;type:uuid;uniqueIndex:idx_server_id_organization_id;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (ServerCredential) TableName() string {
	return ServerCredentialTableName
}

type ServerSnmpCredential struct {
	BaseDbModel
	Community      string       `gorm:"column:community;not null"`
	MaxRepetitions uint8        `gorm:"column:maxRepetitions;type:smallint;not null;default:50"`
	Timeout        uint8        `gorm:"column:timeout;type:smallint;not null;default:10"`
	Port           uint16       `gorm:"column:port;not null;default:161"`
	ServerId       *string      `gorm:"column:serverId;type:uuid;default:null;uniqueIndex:idx_device_id_organization_id;index"`
	Server         Server       `gorm:"constraint:Ondelete:CASCADE"`
	OrganizationId string       `gorm:"column:organizationId;type:uuid;uniqueIndex:idx_device_id_organization_id;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (ServerSnmpCredential) TableName() string {
	return ServerSnmpCredentialTableName
}
