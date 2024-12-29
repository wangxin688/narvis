package mock

import (
	"fmt"
	"time"

	"github.com/wangxin688/narvis/intend/intendtask"
	"github.com/wangxin688/narvis/intend/logger"
	"github.com/wangxin688/narvis/server/dal/gen"
	infra_tasks "github.com/wangxin688/narvis/server/features/infra/tasks"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/tests/fixtures"
	"go.uber.org/zap"
	"golang.org/x/exp/rand"
)

func mockDevice(siteIds []string, orgId string) {
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
		{"paloAlto", "Firewall", "Palo Alto", "PA-5050"},
		{"juniper", "Router", "Juniper", "MX-24T"},
	}
	for siteIndex, siteId := range siteIds {
		for deviceIndex, device := range devices {
			name := fmt.Sprintf("%s-%s-%s-%s-%d-%d", device.platform, device.deviceRole, device.manufacturer, device.deviceModel, siteIndex, deviceIndex)
			ip := fixtures.RandomIpv4PrivateAddress(siteIndex, deviceIndex)
			status := []string{"Active", "Inactive"}[rand.Intn(2)]
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

func mockScanDevice(orgId string) {
	mockScanDevices := make([]*models.ScanDevice, 0)
	devices := []struct {
		platform     string
		manufacturer string
		deviceModel  string
	}{
		{"ciscoXe", "Cisco", "C9300-24T"},
		{"ciscoXe", "Cisco", "C9300-48UMX-A"},
		{"ciscoIos", "Cisco", "C2960-24T"},
		{"ciscoXe", "Cisco", "C9800-L"},
		{"ciscoNxos", "Cisco", "N9504"},
		{"huaweiVrp", "Huawei", "S5270-52P"},
		{"arubaOs", "Aruba", "A7030"},
		{"arubaOs", "Aruba", "A7205"},
		{"fortinet", "Fortinet", "60F"},
		{"fortinet", "Fortinet", "90F"},
		{"ruijie", "Ruijie", "RJ-S5310-24T"},
		{"ruijie", "Ruijie", "RJ-WX2560-X"},
		{"h3c", "H3C", "H3C-24T"},
		{"paloAlto", "Palo Alto", "PA-5050"},
		{"juniper", "Juniper", "MX-24T"},
	}
	for i := 0; i < 10; i++ {
		for deviceIndex, device := range devices {
			name := fmt.Sprintf("%s-%s-%s-%d", device.platform, device.manufacturer, device.deviceModel, deviceIndex)
			ip := fixtures.RandomIpv4PrivateAddress(i, deviceIndex)
			chassisId := fixtures.RandomMacAddress()
			mockScanDevices = append(mockScanDevices, &models.ScanDevice{
				Name:           name,
				ManagementIp:   ip,
				DeviceModel:    device.deviceModel,
				Manufacturer:   device.manufacturer,
				Platform:       device.platform,
				ChassisId:      chassisId,
				Description:    "Mock_" + name,
				OrganizationId: orgId,
			})
		}
	}
	err := gen.ScanDevice.CreateInBatches(mockScanDevices, 500)
	if err != nil {
		panic(err)
	}
}

func mockDeviceConfig(siteId string) {
	deviceIds, err := fixtures.GetRandomDeviceIds(siteId)
	if err != nil {
		logger.Logger.Error("[bootstrap]: failed to get site ids", zap.Error(err))
		panic(err)
	}
	for _, deviceId := range deviceIds {
		deviceConfig := &intendtask.ConfigurationBackupTaskResult{
			Configuration: "mockConfig" + fixtures.RandomString(100),
			DeviceId:      deviceId,
			BackupTime:    time.Now().Format(time.RFC3339),
		}

		infra_tasks.ConfigBackUpCallback(deviceConfig)
	}
}
