package main

import (
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/organization/biz"
	"github.com/wangxin688/narvis/server/features/organization/schemas"
	"github.com/wangxin688/narvis/server/global"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
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

func InitOrganization() string {
	gen.SetDefault(connectDb())
	core.SetUpLogger()
	service := biz.NewOrganizationService()

	newOrg, err := service.CreateOrganization(&schemas.OrganizationCreate{
		Name:           "NarvisDemo",
		EnterpriseCode: "narvis-demo",
		DomainName:     "navis-demo@narvis.com",
		Active:         true,
		LicenseCount:   100000,
		AuthType:       0,
		AdminPassword:  "admin123456",
	})
	if err != nil {
		core.Logger.Error("Failed to create organization", zap.Error(err))
		panic(err)
	}
	global.OrganizationID.Set(newOrg.ID)
	core.Logger.Info("Organization created", zap.String("id", newOrg.ID))
	return newOrg.ID
}

func main() {
	InitOrganization()
}
