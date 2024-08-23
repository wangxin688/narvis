package main

import (
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/models"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func connectDb() *gorm.DB {
	core.SetUpConfig()
	dsn := core.Settings.Postgres.BuildPgDsn()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		core.Logger.Fatal("Failed to connect database", zap.Error(err))
	}
	return db
}

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:        "../../dal/gen",
		Mode:           gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,
		FieldNullable:  true,
		FieldCoverable: true,
	})
	g.UseDB(connectDb())
	g.ApplyBasic(
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
	defer g.Execute()
}
