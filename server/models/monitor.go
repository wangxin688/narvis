package models

type Template struct {
	BaseDbModel
	Name         string `gorm:"column:Name;not null"`
	Manufacturer string `gorm:"column:manufacturer;uniqueIndex:idx_platform_product_family;not null"`
	DeviceRole   string `gorm:"column:deviceRole;uniqueIndex:idx_platform_device_role;not null"`
	TemplateId   string `gorm:"column:templateId;unique;not null"`
}
