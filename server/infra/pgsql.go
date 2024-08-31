package infra

import (
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/migrations"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/pkg/audit"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB
var err error

func InitDB() error {
	var db = DB
	logMode := core.Settings.Postgres.LogMode
	dsn := core.Settings.Postgres.BuildPgDsn()
	if logMode == "debug" {
		logLevel := logger.Info
		db.Logger = db.Logger.LogMode(logLevel)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			// LogMode sets the logger for gorm. Default value is silent.
			Logger:         logger.Default.LogMode(logLevel),
			TranslateError: true,
		})
	} else {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}
	if err != nil {
		core.Logger.Fatal("Failed to connect database", zap.Error(err))
	}
	sqlDB, err := db.DB()
	if err != nil {
		core.Logger.Fatal("Failed to connect database", zap.Error(err))
		return err
	}
	sqlDB.SetMaxIdleConns(core.Settings.Postgres.MaxIdleConns)
	sqlDB.SetMaxOpenConns(core.Settings.Postgres.MaxOpenConns)
	err = db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error
	if err != nil {
		core.Logger.Fatal("failed to create extension: %v", zap.Error(err))
	}

	// set db to global variable
	err = migrations.Migrate(db)
	if err != nil {
		core.Logger.Fatal("Failed to migrate database", zap.Error(err))
		return err
	}

	registerAuditLogMixin()
	DB = db
	return nil
}

func registerAuditLogMixin() {

	newAuditLogMixin := audit.NewAuditLogMixin()
	registeredTables := []string{
		models.SiteTableName,
		models.LocationTableName,
		models.RackTableName,
		models.DeviceTableName,
		models.APTableName,
		models.CliCredentialTableName,
		models.SnmpV2CredentialTableName,
		models.RestconfCredentialTableName,
		models.CircuitTableName,
		models.BlockTableName,
		models.PrefixTableName,
	}
	newAuditLogMixin.AuditTableRegister(registeredTables)

}
