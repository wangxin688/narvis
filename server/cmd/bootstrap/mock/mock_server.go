package mock

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/tests/fixtures"
)

func mockServer(orgId string, siteId string) {
	createServers := make([]*models.Server, 0)

	for i := 1; i < 50; i++ {
		name := fmt.Sprintf("MockServer-%s-%d", siteId, i)
		status := lo.Sample([]string{"Active", "Inactive"})
		mgmtIp := fixtures.RandomIpv4PrivateAddress(i, i)
		deviceModel := lo.Sample([]string{"Dell", "HP", "Lenovo", "Microsoft", "IBM"})
		osVersion := lo.Sample([]string{"debian", "ubuntu", "centos", "windows", "redhat"})
		cpu := lo.Sample([]uint8{1, 2, 4, 8, 16, 32, 48, 64})
		mem := lo.Sample([]uint64{1, 2, 4, 8, 16, 32, 48, 64})
		disk := lo.Sample([]uint64{1, 2, 4, 8, 16, 32, 48, 64, 128, 256, 512, 1024})
		createServers = append(createServers, &models.Server{
			Name:           name,
			Status:         status,
			ManagementIp:   mgmtIp,
			Manufacturer:   deviceModel,
			OsVersion:      osVersion,
			Cpu:            cpu,
			Memory:         mem,
			Disk:           disk,
			SiteId:         siteId,
			OrganizationId: orgId,
		})
	}

	err := gen.Server.CreateInBatches(createServers, 50)
	if err != nil {
		panic(err)
	}
}
