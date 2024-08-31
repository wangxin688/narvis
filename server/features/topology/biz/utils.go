package biz

import (
	"github.com/wangxin688/narvis/intend/devicerole"
	"github.com/wangxin688/narvis/server/features/topology/schemas"
	"github.com/wangxin688/narvis/server/models"
)

func CreateDeviceNode(device models.Device) *schemas.Node {
	dr := devicerole.GetDeviceRole(device.DeviceRole)
	return &schemas.Node{
		Id:           device.Id,
		Name:         device.Name,
		ManagementIp: device.ManagementIp,
		DeviceRole:   device.DeviceRole,
		Weight:       dr.Weight,
	}
}

func CreateApNode(ap models.AP) *schemas.Node {
	return &schemas.Node{
		Id:           ap.Id,
		Name:         ap.Name,
		ManagementIp: ap.ManagementIp,
		DeviceRole:   ap.DeviceRole,
		Weight:       0,
	}
}
