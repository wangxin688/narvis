package mock

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/tests/fixtures"
)

func mockRack(orgId string, siteId string) {
	createRacks := make([]*models.Rack, 0)

	for i := 1; i < 10; i++ {
		name := fmt.Sprintf("MockRack-%s-%d", siteId, i)
		serial := fixtures.RandomString(16)
		uHeight := lo.Sample([]uint8{42, 48, 24})
		descUnit := true
		createRacks = append(createRacks, &models.Rack{
			Name:           name,
			SerialNumber:   &serial,
			UHeight:        uHeight,
			DescUnit:       descUnit,
			SiteId:         siteId,
			OrganizationId: orgId,
		})
	}
	err := gen.Rack.CreateInBatches(createRacks, 10)
	if err != nil {
		panic(err)
	}
}
