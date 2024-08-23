package models

import "github.com/wangxin688/narvis/server/global"

type Template struct {
	global.BaseDbModel
	Manufacturer  string `gorm:"uniqueIndex:idx_platform_product_family;not null"`
	ProductFamily string `gorm:"uniqueIndex:idx_platform_product_family;not null"`
	TemplateID    string `gorm:"unique;not null"`
}
