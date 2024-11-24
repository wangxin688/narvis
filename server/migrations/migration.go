package migrations

import (
	"fmt"

	"github.com/wangxin688/narvis/intend/logger"
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
		&models.WlanStation{},
		&models.Syslog{},
	)
	if err != nil {
		logger.Logger.Fatal("Failed to migrate database", zap.Error(err))
		return err
	}
	createHyperTable(models.SyslogTableName, db)
	createHyperTable(models.AuditLogTableName, db)
	createHyperTable(models.WlanStationTableName, db)
	return nil
}

func createHyperTable(table string, db *gorm.DB) {
	sql := fmt.Sprintf(`
	DO $$
	BEGIN
		IF NOT EXISTS (
			SELECT 1
			FROM timescaledb_information.hypertables
			WHERE hypertable_name = '%s'
		) THEN
			PERFORM  create_hypertable('%s', by_range('time', INTERVAL '1d'), if_not_exists => TRUE);
			PERFORM  add_retention_policy('%s', drop_after => INTERVAL '30d');
			EXECUTE 'alter table %s SET(timescaledb.compress, timescaledb.compress_orderby = ''time DESC'')';
			PERFORM  add_compression_policy('%s', compress_after => INTERVAL '7d');
		END IF;
	END $$;
	`, table, table, table, table, table)
	err := db.Exec(sql).Error
	if err != nil {
		logger.Logger.Fatal("Failed to create hypertable", zap.Error(err))
	}
	logger.Logger.Info("Created hyper table successfully", zap.String("table", table))
}
