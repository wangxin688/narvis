package models

import (
	"github.com/wangxin688/narvis/server/global"
	"gorm.io/datatypes"
)

var AuditLogTableName = "audit_log"

type PostChange struct {
}

type AuditLog struct {
	global.BaseDbSingleModel

	ObjectID       string         `gorm:"type:uuid;not null"`
	ObjectType     string         `gorm:"type:varchar(255);not null"`
	RequestID      *string        `gorm:"type:uuid;default:null"`
	UserID         *string        `gorm:"type:uuid;default:null"`
	Action         string         `gorm:"not null"`
	Data           datatypes.JSON `gorm:"type:json;"`
	OrganizationID string         `gorm:"type:uuid;not null"`
}

func (AuditLog) TableName() string {
	return AuditLogTableName
}
