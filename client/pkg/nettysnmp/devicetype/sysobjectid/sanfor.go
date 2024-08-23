package sysobjectid

import (
	"github.com/wangxin688/narvis/client/pkg/nettysnmp/devicetype"
	"github.com/wangxin688/narvis/intend/manufacturer"
	"github.com/wangxin688/narvis/intend/platform"
)

func SangforDeviceType(sysObjId string) *devicetype.DeviceType {
	// stringPlatform := string(platform.Sangfor)
	oidMap := map[string]map[string]string{}

	data, ok := oidMap[sysObjId]
	if !ok {
		return &devicetype.DeviceType{
			Platform:     platform.Sangfor,
			Manufacturer: manufacturer.Sangfor,
			DeviceType:   devicetype.UnknownDeviceType,
		}
	}

	return &devicetype.DeviceType{
		Platform:     platform.Platform(data["platform"]),
		Manufacturer: manufacturer.Sangfor,
		DeviceType:   data["model"],
	}

}
