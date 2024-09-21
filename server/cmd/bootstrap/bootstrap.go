package main

import (
	"encoding/json"
	"os"

	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/organization/biz"
	"github.com/wangxin688/narvis/server/features/organization/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	InitOrganization()
	InitMacAddress()
}

func connectDB() *gorm.DB {
	core.SetUpConfig()
	dsn := core.Settings.Postgres.BuildPgDsn()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		core.Logger.Fatal("[bootstrap]: failed to connect database", zap.Error(err))
	}
	return db
}

func InitOrganization() string {
	gen.SetDefault(connectDB())
	core.SetUpLogger()
	service := biz.NewOrganizationService()

	org, err := gen.Organization.Where(gen.Organization.Name.Eq("NarvisDemo")).Find()
	if err != nil {
		core.Logger.Error("[bootstrap]: failed to get organization", zap.Error(err))
		panic(err)
	}

	if len(org) > 0 {
		global.OrganizationId.Set(org[0].Id)
		core.Logger.Info("[bootstrap]: organization already exists", zap.String("id", org[0].Id))
		return org[0].Id
	}

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
		core.Logger.Error("[bootstrap]: failed to create organization", zap.Error(err))
		panic(err)
	}
	global.OrganizationId.Set(newOrg.Id)
	core.Logger.Info("[bootstrap]: organization created", zap.String("id", newOrg.Id))
	return newOrg.Id
}

func InitMacAddress() {

	gen.SetDefault(connectDB())
	core.SetUpLogger()
	core.SetUpConfig()
	mac, err := gen.MacAddress.Count()
	if err != nil {
		core.Logger.Error("[bootstrap]: failed to get mac address", zap.Error(err))
		panic(err)
	}
	if mac >= 1 {
		core.Logger.Info("[bootstrap]: mac address already exists")
		return
	}
	macAddressFilePath := core.ProjectPath + "/cmd/bootstrap/appdata/mac_address.json"
	file, err := os.Open(macAddressFilePath)
	if err != nil {
		core.Logger.Error("[bootstrap]: failed to open mac address file", zap.Error(err))
		panic(err)
	}
	defer file.Close()
	var macAddresses []*models.MacAddress
	if err := json.NewDecoder(file).Decode(&macAddresses); err != nil {
		core.Logger.Error("[bootstrap]: failed to decode mac address file", zap.Error(err))
		panic(err)
	}
	err = gen.MacAddress.CreateInBatches(macAddresses, 100)
	if err != nil {
		core.Logger.Error("[bootstrap]: failed to create mac address", zap.Error(err))
		panic(err)
	}
	core.Logger.Info("[bootstrap]: mac address created")
}
