package sysobjectid

import (
	"github.com/wangxin688/narvis/client/pkg/nettysnmp/devicetype"
	"github.com/wangxin688/narvis/intend/manufacturer"
	"github.com/wangxin688/narvis/intend/platform"
)

func MikroTikDeviceType(sysObjId string) *devicetype.DeviceType {
	stringPlatform := string(platform.MikroTik)
	oidMap := map[string]map[string]string{
		"1.3.6.1.4.1.14988.1": {"platform": stringPlatform, "model": "RB1200"},
	}

	data, ok := oidMap[sysObjId]
	if !ok {
		return &devicetype.DeviceType{
			Platform:     platform.MikroTik,
			Manufacturer: manufacturer.MikroTik,
			DeviceType:   devicetype.UnknownDeviceType,
		}
	}

	return &devicetype.DeviceType{
		Platform:     platform.Platform(data["platform"]),
		Manufacturer: manufacturer.MikroTik,
		DeviceType:   data["model"],
	}

}
