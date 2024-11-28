package mock

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/intend/model/devicerole"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/tests/fixtures"
)

func mockWlanAp(orgId, siteId string) {
	createAps := make([]*models.AP, 0)

	for i := 1; i < 50; i++ {
		name := fmt.Sprintf("MockAP-%s-%d", siteId, i)
		status := lo.Sample([]string{"Active", "Inactive"})
		macAddr := fixtures.RandomMacAddress()
		serialNumber := lo.Sample([]string{fixtures.RandomString(16), ""})
		mgmtIp := fixtures.RandomIpv4PrivateAddress(i, i)
		deviceModel := string(devicerole.WlanAP)
		osVersion := lo.Sample([]string{"v8.1.0", "v6.0.1", ""})
		Floor := lo.Sample([]string{"F1", "F2", "F3", "F4"})
		createAps = append(createAps, &models.AP{
			Name:           name,
			Status:         status,
			MacAddress:     macAddr,
			SerialNumber:   &serialNumber,
			ManagementIp:   mgmtIp,
			DeviceModel:    deviceModel,
			OsVersion:      &osVersion,
			Floor:          &Floor,
			SiteId:         siteId,
			OrganizationId: orgId,
		})
	}
	err := gen.AP.CreateInBatches(createAps, 50)
	if err != nil {
		panic(err)
	}
}
