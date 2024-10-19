package infra_utils

import (
	"github.com/wangxin688/narvis/server/models"
)

func DevicesToIds(devices []*models.Device) []string {
	ids := make([]string, len(devices))
	for i, device := range devices {
		ids[i] = device.Id
	}
	return ids
}

func DevicePlatforms(devices []*models.Device) map[string]string {
	result := make(map[string]string, len(devices))
	for _, device := range devices {
		result[device.Id] = device.Platform
	}
	return result
}

func DeviceManagementIpMap(devices []*models.Device) map[string]string {
	result := make(map[string]string, len(devices))

	for _, device := range devices {
		result[device.Id] = device.ManagementIp
	}
	return result
}
