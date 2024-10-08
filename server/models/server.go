package models

type VirtualMachine struct {
	BaseDbModel
	VCPUs          uint8        `gorm:"column:vCPUs;not null"`
	Memory         uint16       `gorm:"column:memory;not null"`
	Disk           uint16       `gorm:"column:disk;not null"`
	
	SiteId         string       `gorm:"column:siteId;type:uuid;index;not null"`
	Site           Site         `gorm:"constraint:Ondelete:RESTRICT"`
	OrganizationId string       `gorm:"column:organizationId;type:uuid;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}
