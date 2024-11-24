package mock

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/wangxin688/narvis/server/tests/fixtures"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

func MockSite(db *gorm.DB) {
	orgId, err := fixtures.GetOrgId()
	if err != nil {
		panic(err)
	}
	regions := []struct {
		region    string
		timeZone  string
		latitude  float64
		longitude float64
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
		status := []string{"Active", "InActive"}[rand.Intn(2)]

		// Random time generation
		createdAt := time.Date(2023, 1, rand.Intn(10)+1, rand.Intn(24), rand.Intn(60), rand.Intn(60), 0, time.UTC)
		updatedAt := createdAt.Add(time.Duration(rand.Intn(3600)) * time.Second) // Ensure updatedAt > createdAt

		sql := fmt.Sprintf("INSERT INTO infra_site (id, \"createdAt\", \"updatedAt\", name, \"siteCode\", status, region, \"timeZone\", latitude, longitude, address, \"organizationId\") VALUES "+
			"('%s', '%s', '%s', 'mock%d', 'mock%d', '%s', '%s', '%s', %.2f, %.2f, 'mock_address_%d', '%s');\n",
			uuid.NewString(), createdAt.Format("2006-01-02 15:04:05"), updatedAt.Format("2006-01-02 15:04:05"), i, i, status, region.region, region.timeZone, region.latitude, region.longitude, i, orgId)
		_ = db.Exec(sql).Error
	}
}
