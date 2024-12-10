package mock

import (
	"github.com/wangxin688/narvis/intend/logger"
	"github.com/wangxin688/narvis/server/tests/fixtures"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func GenerateMockData(orgId string, db *gorm.DB) {
	siteIds, err := fixtures.GetRandomSiteIds(orgId)
	if len(siteIds) > 0 || err != nil {
		for _, siteId := range siteIds {
			mockWlanUser(orgId, siteId)
		}
		return
	}
	mockSite(orgId)
	mockScanDevice(orgId)
	siteIds, err = fixtures.GetRandomSiteIds(orgId)
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to get site ids", zap.Error(err))
		panic(err)
	}
	mockDevice(siteIds, orgId)
	logger.Logger.Info("[bootstrap]: mock data created")
	for _, siteId := range siteIds {
		mockWlanAp(orgId, siteId)
		mockWlanUser(orgId, siteId)
		mockRack(orgId, siteId)
		mockServer(orgId, siteId)
		deviceIds, err := fixtures.GetRandomDeviceIds(siteId)
		if err != nil {
			logger.Logger.Error("[bootstrap]: failed to get site ids", zap.Error(err))
			panic(err)
		}
		for _, deviceId := range deviceIds {
			mockDeviceInterface(siteId, deviceId)
			interfaces, err := fixtures.GetRandomInterfaceIds(deviceId)
			if err != nil {
				panic(err)
			}
			mockCircuit(siteId, orgId, interfaces)
			mockDeviceAlerts(orgId, siteId, deviceId)
		}
		apIds, err := fixtures.GetRandomApIds(siteId)
		if err != nil {
			panic(err)
		}
		for _, apId := range apIds {
			mockApAlerts(orgId, siteId, apId)
		}

		circuitIds, err := fixtures.GetRandomCircuitIds(siteId)
		if err != nil {
			panic(err)
		}
		for _, circuitId := range circuitIds {
			mockCircuitAlerts(orgId, siteId, circuitId)
		}

		serverIds, err := fixtures.GetRandomServerIds(siteId)
		if err != nil {
			panic(err)
		}
		for _, serverId := range serverIds {
			mockServerAlerts(orgId, siteId, serverId)
		}

		mockPrefix(orgId, siteId)
	}
}
