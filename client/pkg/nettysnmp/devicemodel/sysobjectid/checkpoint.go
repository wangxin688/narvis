package sysobjectid

import (
	"github.com/wangxin688/narvis/client/pkg/nettysnmp/devicemodel"
	"github.com/wangxin688/narvis/intend/model/manufacturer"
	"github.com/wangxin688/narvis/intend/platform"
)

func CheckPointDeviceModel(sysObjId string) *devicemodel.DeviceModel {
	// stringPlatform := string(platform.CheckPoint)
	oidMap := map[string]map[string]string{}

	data, ok := oidMap[sysObjId]
	if !ok {
		return &devicemodel.DeviceModel{
			Platform:     platform.CheckPoint,
			Manufacturer: manufacturer.CheckPoint,
			DeviceModel:  devicemodel.UnknownDeviceModel,
		}
	}

	return &devicemodel.DeviceModel{
		Platform:     platform.Platform(data["platform"]),
		Manufacturer: manufacturer.CheckPoint,
		DeviceModel:  data["model"],
	}

}
