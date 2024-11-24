package models

import (
	"gorm.io/datatypes"
)

var AuditLogTableName = "audit_log"

type PostChange struct {
}

type AuditLog struct {
	BaseTimeScaleModel

	ObjectId       string         `gorm:"column:objectId;type:uuid;not null"`
	ObjectType     string         `gorm:"column:objectType;not null"`
	RequestId      *string        `gorm:"column:requestId;type:uuid;default:null"`
	UserId         *string        `gorm:"column:userId;type:uuid;default:null"`
	Action         string         `gorm:"column:action;not null"`
	Data           datatypes.JSON `gorm:"column:data;type:json;"`
	Diff           datatypes.JSON `gorm:"column:diff;type:json;"` // Only for update change
	OrganizationId string         `gorm:"column:organizationId;type:uuid;not null"`
}

func (AuditLog) TableName() string {
	return AuditLogTableName
}
