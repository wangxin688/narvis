package mock

import (
	"fmt"

	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/models"
	"golang.org/x/exp/rand"
)

func mockSite(orgId string) {

	sites := make([]*models.Site, 0)
	regions := []struct {
		region    string
		timeZone  string
		latitude  float32
		longitude float32
	}{
		{"China/Shanghai", "Asia/Shanghai", 31.23, 121.47},
		{"USA/New York", "America/New_York", 40.71, -74.01},
		{"Germany/Berlin", "Europe/Berlin", 52.52, 13.41},
		{"India/Mumbai", "Asia/Kolkata", 19.07, 72.87},
		{"Australia/Sydney", "Australia/Sydney", -33.87, 151.21},
		{"Japan/Tokyo", "Asia/Tokyo", 35.68, 139.69},
		{"Canada/Toronto", "America/Toronto", 43.65, -79.38},
		{"France/Paris", "Europe/Paris", 48.86, 2.35},
		{"Brazil/Sao Paulo", "America/Sao_Paulo", -23.55, -46.63},
		{"South Africa/Johannesburg", "Africa/Johannesburg", -26.2, 28.04},
	}

	for i := 1; i <= 50; i++ {
		region := regions[rand.Intn(len(regions))]
		status := []string{"Active", "Inactive"}[rand.Intn(2)]

		sites = append(sites, &models.Site{
			Name:           fmt.Sprintf("mock_Site_%d", i),
			SiteCode:       fmt.Sprintf("mock_siteCode_%d", i),
			Status:         status,
			Region:         region.region,
			TimeZone:       region.timeZone,
			Latitude:       region.latitude,
			Longitude:      region.longitude,
			Address:        fmt.Sprintf("mock_address %d", i),
			Description:    nil,
			OrganizationId: orgId,
		})
	}
	err := gen.Site.CreateInBatches(sites, 50)
	if err != nil {
		panic(err)
	}
}
