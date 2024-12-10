package fixtures

import (
	"github.com/samber/lo"
	"github.com/wangxin688/narvis/server/config"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/infra"
	"github.com/wangxin688/narvis/server/models"
)

func GetRandomSiteIds(orgId string) ([]string, error) {
	result, err := gen.Site.Select(gen.Site.Id).Where(
		gen.Site.OrganizationId.Eq(orgId),
	).Limit(100).Find()
	if err != nil {
		return nil, err
	}
	siteIds := lo.Map(result, func(item *models.Site, _ int) string {
		return item.Id
	})
	return siteIds, nil
}

func GetRandomDeviceIds(siteId string) ([]string, error) {
	result, err := gen.Device.Select(gen.Device.Id).Where(
		gen.Device.SiteId.Eq(siteId),
	).Limit(500).Find()
	if err != nil {
		return nil, err
	}
	deviceIds := lo.Map(result, func(item *models.Device, _ int) string {
		return item.Id
	})
	return deviceIds, nil
}

func GetRandomInterfaceIds(deviceId string) ([]string, error) {
	result, err := gen.DeviceInterface.Select(gen.DeviceInterface.Id).Where(
		gen.DeviceInterface.DeviceId.Eq(deviceId),
	).Find()
	if err != nil {
		return nil, err
	}
	interfaceIds := lo.Map(result, func(item *models.DeviceInterface, _ int) string {
		return item.Id
	})
	return interfaceIds, nil

}

func GetRandomApIds(siteId string) ([]string, error) {
	result, err := gen.AP.Select(gen.AP.Id).Where(
		gen.AP.SiteId.Eq(siteId),
	).Find()
	if err != nil {
		return nil, err
	}
	apIds := lo.Map(result, func(item *models.AP, _ int) string {
		return item.Id
	})
	return apIds, nil

}

func GetRandomCircuitIds(siteId string) ([]string, error) {
	result, err := gen.Circuit.Select(gen.Circuit.Id).Where(
		gen.Circuit.SiteId.Eq(siteId),
	).Find()
	if err != nil {
		return nil, err
	}
	apIds := lo.Map(result, func(item *models.Circuit, _ int) string {
		return item.Id
	})
	return apIds, nil

}

func GetRandomServerIds(siteId string) ([]string, error) {
	result, err := gen.Server.Select(gen.Server.Id).Where(
		gen.Server.SiteId.Eq(siteId),
	).Find()
	if err != nil {
		return nil, err
	}
	apIds := lo.Map(result, func(item *models.Server, _ int) string {
		return item.Id
	})
	return apIds, nil

}

func GetSiteApNames(siteId string) ([]string, error) {
	result, err := gen.AP.Select(gen.AP.Name).Where(
		gen.AP.SiteId.Eq(siteId),
	).Find()
	if err != nil {
		return nil, err
	}
	apNames := lo.Map(result, func(item *models.AP, _ int) string {
		return item.Name
	})
	return apNames, nil
}

func FixturePrepare() {
	config.InitTestFixtureConfig()
	config.InitLogger()
	err := infra.InitDB()
	if err != nil {
		panic(err)
	}
}
