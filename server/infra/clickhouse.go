package infra

import (
	"fmt"

	"github.com/wangxin688/narvis/server/core"
	"go.uber.org/zap"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

var ClickHouseDB *gorm.DB

func InitClickHouseDB() error {
	if core.Settings == nil {
		core.Logger.Fatal("[infraConnectDb]: core settings are not initialized")
		return fmt.Errorf("core settings are not initialized")
	}

	dsn := core.Settings.ClickHouse.BuildClickHouseDsn()
	if dsn == "" {
		core.Logger.Fatal("[infraConnectDb]: failed to build ClickHouse DSN")
		return fmt.Errorf("failed to build ClickHouse DSN")
	}

	db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		core.Logger.Fatal("[infraConnectDb]: failed to connect clickhouse", zap.Error(err))
		return err
	}

	sqlDB, err := db.DB()
	if err != nil {
		core.Logger.Fatal("[infraConnectDb]: failed to get generic database object", zap.Error(err))
		return err
	}

	err = sqlDB.Ping()
	if err != nil {
		core.Logger.Fatal("[infraConnectDb]: failed to ping clickhouse", zap.Error(err))
		return err
	}

	ClickHouseDB = db
	return nil
}
