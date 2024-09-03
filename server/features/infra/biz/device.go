package biz

import (
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	infra_utils "github.com/wangxin688/narvis/server/features/infra/utils"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"
)

type DeviceService struct{}

func NewDeviceService() *DeviceService {
	return &DeviceService{}
}

func (d *DeviceService) CreateDevice(device *schemas.DeviceCreate) (string, error) {
	newDevice := models.Device{
		Name:           device.Name,
		ManagementIp:   device.ManagementIp,
		Status:         device.Status,
		DeviceModel:    *device.DeviceModel,
		Manufacturer:   *device.Manufacturer,
		DeviceRole:     device.DeviceRole,
		Floor:          device.Floor,
		SiteId:         device.SiteId,
		OrganizationId: global.OrganizationId.Get(),
	}
	if device.RackId != nil && device.RackPosition != nil {
		newDevice.RackId = device.RackId
		position, err := infra_utils.SliceUint8ToString(device.RackPosition)
		if err != nil {
			return "", err
		}
		newDevice.RackPosition = &position
	}

	err := gen.Device.Create(&newDevice)
	if err != nil {
		return "", err
	}
	return newDevice.Id, nil
}
