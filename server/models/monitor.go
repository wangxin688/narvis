package models

type Template struct {
	BaseDbModel
	Manufacturer  string `gorm:"column:manufacturer;uniqueIndex:idx_platform_product_family;not null"`
	ProductFamily string `gorm:"column:productFamily;uniqueIndex:idx_platform_product_family;not null"`
	TemplateId    string `gorm:"column:templateId;unique;not null"`
}
