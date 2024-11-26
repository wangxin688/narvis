package mock

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/server/dal/gen"
	infra_biz "github.com/wangxin688/narvis/server/features/infra/biz"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/tests/fixtures"
	"gorm.io/gorm"
)

func MockCircuit(db *gorm.DB, siteId, orgId string, interfaceId []string) {
	createCircuits := make([]*models.Circuit, 0)
	for _, iface := range interfaceId {
		circuitService := infra_biz.NewCircuitService()
		name := fmt.Sprintf("MockCircuit-%s", fixtures.RandomString(8))
		cid := fixtures.RandomString(10)
		status := lo.Sample([]string{"Active", "Inactive"})
		ctype := lo.Sample([]string{"Internet", "MPLS", "IEPL", "DPLC", "DarkFiber", "ADSL"})
		rxBandWidth := lo.Sample([]int32{100, 500, 1000, 2000, 10000})
		ipv4 := fixtures.RandomIpv4()
		ipv6 := fixtures.RandomIpv6Address()
		provider := lo.Sample([]string{"CU", "CT", "CM", "GTT", "NTT", "ATT"})
		deviceId, _, err := circuitService.GetDeviceSiteIdByInterfaceId(iface)
		if err != nil {
			panic(err)
		}
		createCircuits = append(createCircuits, &models.Circuit{
			Name:           name,
			CId:            &cid,
			Status:         status,
			CircuitType:    ctype,
			RxBandWidth:    uint32(rxBandWidth),
			TxBandWidth:    uint32(rxBandWidth),
			Ipv4Address:    &ipv4,
			Ipv6Address:    &ipv6,
			InterfaceId:    iface,
			DeviceId:       deviceId,
			SiteId:         siteId,
			OrganizationId: orgId,
			Provider:       provider,
		})
	}

	err := gen.Circuit.CreateInBatches(createCircuits, 50)
	if err != nil {
		panic(err)
	}
}
