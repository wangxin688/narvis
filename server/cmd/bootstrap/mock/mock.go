package mock

import (
	"github.com/wangxin688/narvis/intend/logger"
	"github.com/wangxin688/narvis/server/tests/fixtures"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GenerateMockData(orgId string, db *gorm.DB) {
	siteIds, err := fixtures.GetRandomSiteIds(orgId)
	if len(siteIds) >0 || err != nil{
		return
	}
	MockSite(db, orgId)
	siteIds, err = fixtures.GetRandomSiteIds(orgId)	
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to get site ids", zap.Error(err))
		panic(err)
	}
	
	MockDevice(db, siteIds, orgId)
	logger.Logger.Info("[bootstrap]: mock data created")
	for _, siteId := range siteIds {
		MockWlanAp(db, orgId, siteId)
		deviceIds, err := fixtures.GetRandomDeviceIds(siteId)
		if err != nil {
			logger.Logger.Error("[bootstrap]: failed to get site ids", zap.Error(err))
			panic(err)
		}
		for _, deviceId := range deviceIds {
			MockDeviceInterface(db, siteId, deviceId)
			interfaces, err := fixtures.GetRandomInterfaceIds(deviceId)
			if err != nil {
				panic(err)
			}
			MockCircuit(db, siteId, orgId, interfaces)
		}

	}
}
