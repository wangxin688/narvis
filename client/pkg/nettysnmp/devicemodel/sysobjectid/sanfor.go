package sysobjectid

import (
	"github.com/wangxin688/narvis/client/pkg/nettysnmp/devicemodel"
	"github.com/wangxin688/narvis/intend/manufacturer"
	"github.com/wangxin688/narvis/intend/platform"
)

func SangforDeviceModel(sysObjId string) *devicemodel.DeviceModel {
	// stringPlatform := string(platform.Sangfor)
	oidMap := map[string]map[string]string{}

	data, ok := oidMap[sysObjId]
	if !ok {
		return &devicemodel.DeviceModel{
			Platform:     platform.Sangfor,
			Manufacturer: manufacturer.Sangfor,
			DeviceModel:  devicemodel.UnknownDeviceModel,
		}
	}

	return &devicemodel.DeviceModel{
		Platform:     platform.Platform(data["platform"]),
		Manufacturer: manufacturer.Sangfor,
		DeviceModel:  data["model"],
	}

}
