package main

import (
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/models"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func connectDb() *gorm.DB {
	core.SetUpConfig()
	dsn := core.Settings.Postgres.BuildPgDsn()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase:   true,
			SingularTable: true,
		},
	})
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
		&models.MacAddress{},
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
	)
	defer g.Execute()
}
