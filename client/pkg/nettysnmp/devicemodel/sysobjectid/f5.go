package sysobjectid

import (
	"github.com/wangxin688/narvis/client/pkg/nettysnmp/devicemodel"
	"github.com/wangxin688/narvis/intend/model/manufacturer"
	"github.com/wangxin688/narvis/intend/platform"
)

func F5DeviceModel(sysObjId string) *devicemodel.DeviceModel {
	// stringPlatform := string(platform.F5)
	oidMap := map[string]map[string]string{}

	data, ok := oidMap[sysObjId]
	if !ok {
		return &devicemodel.DeviceModel{
			Platform:     platform.F5,
			Manufacturer: manufacturer.F5,
			DeviceModel:  devicemodel.UnknownDeviceModel,
		}
	}

	return &devicemodel.DeviceModel{
		Platform:     platform.Platform(data["platform"]),
		Manufacturer: manufacturer.F5,
		DeviceModel:  data["model"],
	}

}
