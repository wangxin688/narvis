package models

import (
	"github.com/wangxin688/narvis/server/global"
	"gorm.io/datatypes"
)

type PostChange struct {
}

type AuditLog struct {
	global.BaseDbSingleModel

	ObjectId   string          `gorm:"type:uuid;not null"`
	ObjectType string          `gorm:"type:varchar(255);not null"`
	UserId     string          `gorm:"type:uuid;not null"`
	Action     uint8           `gorm:"type:smallint;not null"`
	PreChange  *datatypes.JSON `gorm:"type:json;default:null"`
	Diff       *datatypes.JSON `gorm:"type:json;default:null"`
	PostChange *datatypes.JSON `gorm:"type:json;default:null"`
}
