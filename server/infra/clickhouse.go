package infra

import (
	"github.com/wangxin688/narvis/server/core"
	"go.uber.org/zap"
	"gorm.io/driver/clickhouse"
	"gorm.io/gorm"
)

var ClickHouseDB *gorm.DB

func InitClickHouseDB() error {
	dsn := core.Settings.ClickHouse.BuildClickHouseDsn()
	db, err := gorm.Open(clickhouse.Open(dsn), &gorm.Config{})
	if err != nil {
		core.Logger.Fatal("[infraConnectDb]: failed to connect clickhouse", zap.Error(err))
		return err
	}
	ClickHouseDB = db
	return nil
}
