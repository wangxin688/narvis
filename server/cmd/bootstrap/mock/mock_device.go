package mock

import (
	"fmt"

	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/tests/fixtures"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
)

func MockDevice(db *gorm.DB, siteIds []string, orgId string) {
	createDevices := make([]*models.Device, 0)
	devices := []struct {
		platform     string
		deviceRole   string
		manufacturer string
		deviceModel  string
	}{
		{"ciscoXe", "Switch", "Cisco", "C9300-24T"},
		{"ciscoXe", "Switch", "Cisco", "C9300-48UMX-A"},
		{"ciscoIos", "Switch", "Cisco", "C2960-24T"},
		{"ciscoXe", "WlanAC", "Cisco", "C9800-L"},
		{"ciscoNxos", "Switch", "Cisco", "N9504"},
		{"huaweiVrp", "Switch", "Huawei", "S5270-52P"},
		{"arubaOs", "WlanAC", "Aruba", "A7030"},
		{"arubaOs", "WlanAC", "Aruba", "A7205"},
		{"fortinet", "Firewall", "Fortinet", "60F"},
		{"fortinet", "Firewall", "Fortinet", "90F"},
		{"ruijie", "Switch", "Ruijie", "RJ-S5310-24T"},
		{"ruijie", "WlanAC", "Ruijie", "RJ-WX2560-X"},
		{"h3c", "Switch", "H3C", "H3C-24T"},
	}
	for siteIndex, siteId := range siteIds {
		for deviceIndex, device := range devices {
			name := fmt.Sprintf("%s-%s-%s-%s-%d-%d", device.platform, device.deviceRole, device.manufacturer, device.deviceModel, siteIndex, deviceIndex)
			ip := fixtures.RandomIpv4PrivateAddress(siteIndex, deviceIndex)
			status := []string{"Active", "InActive"}[rand.Intn(2)]
			chassisId := fixtures.RandomMacAddress()
			floor := fmt.Sprintf("%dF", siteIndex)
			createDevices = append(createDevices, &models.Device{
				Name:           name,
				ManagementIp:   ip,
				Status:         status,
				DeviceModel:    device.deviceModel,
				Manufacturer:   device.manufacturer,
				Platform:       device.platform,
				DeviceRole:     device.deviceRole,
				ChassisId:      &chassisId,
				SiteId:         siteId,
				OrganizationId: orgId,
				Floor:          &floor,
				SerialNumber:   &name,
			})
		}
	}
	err := gen.Device.CreateInBatches(createDevices, 100)
	if err != nil {
		panic(err)
	}

}
