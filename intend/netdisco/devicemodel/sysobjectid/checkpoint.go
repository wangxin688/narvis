package sysobjectid

import (
	manufacturer "github.com/wangxin688/narvis/intend/model/manufacturer"
	platform "github.com/wangxin688/narvis/intend/model/platform"
	"github.com/wangxin688/narvis/intend/netdisco/devicemodel"
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
