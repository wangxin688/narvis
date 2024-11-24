package infra

import (
	nlog "github.com/wangxin688/narvis/intend/logger"
	"github.com/wangxin688/narvis/server/config"
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
	logMode := config.Settings.Postgres.LogMode
	dsn := config.Settings.Postgres.BuildPgDsn()
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
		nlog.Logger.Fatal("[infraConnectDb]: failed to connect database", zap.Error(err))
	}
	sqlDB, err := db.DB()
	if err != nil {
		nlog.Logger.Fatal("[infraConnectDb]: failed to connect database", zap.Error(err))
		return err
	}
	sqlDB.SetMaxIdleConns(config.Settings.Postgres.MaxIdleConns)
	sqlDB.SetMaxOpenConns(config.Settings.Postgres.MaxOpenConns)

	// set db to global variable

	registerAuditLogMixin(db)
	DB = db
	return nil
}

func registerAuditLogMixin(db *gorm.DB) {

	newAuditLogMixin := audit.NewAuditLogMixin()
	registeredTables := []string{
		models.SiteTableName,
		models.RackTableName,
		models.DeviceTableName,
		models.APTableName,
		models.CliCredentialTableName,
		models.SnmpV2CredentialTableName,
		models.RestconfCredentialTableName,
		models.CircuitTableName,
		models.PrefixTableName,
		models.ServerTableName,
		models.ServerCredentialTableName,
		models.ServerSnmpCredentialTableName,
	}
	newAuditLogMixin.AuditTableRegister(registeredTables)
	newAuditLogMixin.RegisterCallbacks(db)
}

func AutoMigration(db *gorm.DB) error {
	err = migrations.Migrate(db)
	if err != nil {
		nlog.Logger.Fatal("[infraConnectDb]: failed to migrate database", zap.Error(err))
		return err
	}
	nlog.Logger.Info("[infraConnectDb]: database migrated")
	return nil
}
