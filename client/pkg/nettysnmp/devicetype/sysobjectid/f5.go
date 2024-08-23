package sysobjectid

import (
	"github.com/wangxin688/narvis/client/pkg/nettysnmp/devicetype"
	"github.com/wangxin688/narvis/intend/manufacturer"
	"github.com/wangxin688/narvis/intend/platform"
)

func F5DeviceType(sysObjId string) *devicetype.DeviceType {
	// stringPlatform := string(platform.F5)
	oidMap := map[string]map[string]string{}

	data, ok := oidMap[sysObjId]
	if !ok {
		return &devicetype.DeviceType{
			Platform:     platform.F5,
			Manufacturer: manufacturer.F5,
			DeviceType:   devicetype.UnknownDeviceType,
		}
	}

	return &devicetype.DeviceType{
		Platform:     platform.Platform(data["platform"]),
		Manufacturer: manufacturer.F5,
		DeviceType:   data["model"],
	}

}
