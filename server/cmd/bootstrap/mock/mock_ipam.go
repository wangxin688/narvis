package mock

import (
	"fmt"

	"golang.org/x/exp/rand"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/tests/fixtures"
)

func mockPrefix(orgId string, siteId string) {
	createPrefix := make([]*models.Prefix, 0)
	newIpAddrs := make([]*models.IpAddress, 0)
	for i := 0; i < 10; i++ {
		newRange, gateway := fixtures.GenerateRFC1918Prefix()
		version := "IPv4"
		type_ := lo.Sample([]string{"Dynamic", "Static"})
		vlanId := uint32(rand.Intn(4094))
		vlanName := fmt.Sprintf("Mock_vlan_%d", vlanId)
		createPrefix = append(createPrefix, &models.Prefix{
			Range:          newRange,
			Version:        version,
			Type:           type_,
			VlanId:         &vlanId,
			VlanName:       &vlanName,
			Gateway:        &gateway,
			SiteId:         siteId,
			OrganizationId: orgId,
		})
		randomPercent := rand.Float64()
		ipAddrs, err := fixtures.GenerateRandomIPsByPrefix(newRange, randomPercent)
		if err != nil {
			panic(err)
		}
		for _, addr := range ipAddrs {
			macAddr := fixtures.RandomMacAddress()
			mockDescr := fmt.Sprintf("Mock descr %s", addr)
			newIpAddrs = append(newIpAddrs, &models.IpAddress{
				Address:        addr,
				Status:         lo.Sample([]string{"Active", "Reserved"}),
				MacAddress:     &macAddr,
				Vlan:           &vlanId,
				Range:          &newRange,
				Description:    &mockDescr,
				Type:           lo.Sample([]string{"Dynamic", "Static"}),
				SiteId:         siteId,
				OrganizationId: orgId,
			})
		}

	}
	err := gen.Prefix.CreateInBatches(createPrefix, 50)
	if err != nil {
		panic(err)
	}
	err = gen.IpAddress.CreateInBatches(newIpAddrs, 500)
	if err != nil {
		panic(err)
	}
}
