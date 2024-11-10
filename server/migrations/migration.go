package migrations

import (
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.Organization{},
		&models.Proxy{},
		&models.AuditLog{},
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.Menu{},
		&models.Site{},
		&models.Rack{},
		&models.Device{},
		&models.DeviceStack{},
		&models.DeviceInterface{},
		&models.LLDPNeighbor{},
		&models.ApLLDPNeighbor{},
		&models.DeviceConfig{},
		&models.CliCredential{},
		&models.SnmpV2Credential{},
		&models.RestconfCredential{},
		&models.AP{},
		&models.Server{},
		&models.ServerCredential{},
		&models.ServerSnmpCredential{},
		&models.MacAddress{},
		&models.ScanDevice{},
		&models.Circuit{},
		&models.Prefix{},
		&models.IpAddress{},
		&models.AlertGroup{},
		&models.Alert{},
		&models.AlertActionLog{},
		&models.Subscription{},
		&models.Maintenance{},
		&models.RootCause{},
		&models.SubscriptionRecord{},
		&models.Template{},
		&models.TaskResult{},
	)
	if err != nil {
		core.Logger.Fatal("Failed to migrate database", zap.Error(err))
		return err
	}
	return nil
}
