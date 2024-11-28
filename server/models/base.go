package models

import (
	"time"
)

type BaseDbModel struct {
	Id        string    `gorm:"column:id;primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"column:createdAt;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updatedAt;autoUpdateTime"`
}

type BaseDbSingleModel struct {
	Id        string    `gorm:"column:id;primary_key;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time `gorm:"column:createdAt;autoCreateTime"`
}

type BaseTimeScaleModel struct {
	Time time.Time `gorm:"column:time;autoCreateTime" json:"time"`
}
