package models

type Template struct {
	BaseDbModel
	TemplateName string `gorm:"column:templateName;not null"`
	Platform     string `gorm:"column:platform;uniqueIndex:idx_platform_device_role;not null"`
	DeviceRole   string `gorm:"column:deviceRole;uniqueIndex:idx_platform_device_role;not null"`
	TemplateId   string `gorm:"column:templateId;not null"`
}
