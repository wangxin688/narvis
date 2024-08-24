package models

import (
	"time"
)

type BaseDbModel struct {
	ID        string    `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

type BaseDbSingleModel struct {
	ID        string    `gorm:"primary_key;type:uuid;default:uuid_generate_v4()" json:"id"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
