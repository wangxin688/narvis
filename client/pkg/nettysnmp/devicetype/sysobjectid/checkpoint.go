package sysobjectid

import (
	"github.com/wangxin688/narvis/client/pkg/nettysnmp/devicetype"
	"github.com/wangxin688/narvis/intend/manufacturer"
	"github.com/wangxin688/narvis/intend/platform"
)

func CheckPointDeviceType(sysObjId string) *devicetype.DeviceType {
	// stringPlatform := string(platform.CheckPoint)
	oidMap := map[string]map[string]string{}

	data, ok := oidMap[sysObjId]
	if !ok {
		return &devicetype.DeviceType{
			Platform:     platform.CheckPoint,
			Manufacturer: manufacturer.CheckPoint,
			DeviceType:   devicetype.UnknownDeviceType,
		}
	}

	return &devicetype.DeviceType{
		Platform:     platform.Platform(data["platform"]),
		Manufacturer: manufacturer.CheckPoint,
		DeviceType:   data["model"],
	}

}
