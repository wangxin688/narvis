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
		&models.User{},
		&models.Role{},
		&models.Group{},
		&models.Permission{},
		&models.Menu{},
		&models.Site{},
		&models.Location{},
		&models.Rack{},
		&models.Device{},
		&models.DeviceInterface{},
		&models.LLDPNeighbor{},
		&models.DeviceConfig{},
		&models.DeviceCliCredential{},
		&models.DeviceSnmpV2Credential{},
		&models.DeviceRestconfCredential{},
		&models.Provider{},
		&models.Circuit{},
		&models.Block{},
		&models.Prefix{},
		&models.IpAddress{},
		&models.Vlan{},
		&models.AlertGroup{},
		&models.Alert{},
		&models.ActionLog{},
		&models.Subscription{},
		&models.Maintenance{},
		&models.RootCause{},
		&models.SubscriptionRecord{},
		&models.Template{},
	)
	if err != nil {
		core.Logger.Fatal("Failed to migrate database", zap.Error(err))
		return err
	}
	return nil
}
