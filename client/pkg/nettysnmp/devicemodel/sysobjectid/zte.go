package sysobjectid

import (
	"github.com/wangxin688/narvis/client/pkg/nettysnmp/devicemodel"
	"github.com/wangxin688/narvis/intend/manufacturer"
	"github.com/wangxin688/narvis/intend/platform"
)

func ZTEDeviceModel(sysObjId string) *devicemodel.DeviceModel {
	stringPlatform := string(platform.ZTE)
	oidMap := map[string]map[string]string{
		".1.3.6.1.4.1.3902.1004.9806.2.1.1": {"platform": stringPlatform, "model": "ZXDSL9806H"},
		".1.3.6.1.4.1.3902.3.100.14":        {"platform": stringPlatform, "model": "ZXR10"},
		".1.3.6.1.4.1.3902.3.100.173":       {"platform": stringPlatform, "model": "ZXR10 2842"},
		".1.3.6.1.4.1.3902.3.100.20":        {"platform": stringPlatform, "model": "ZXR10 5928"},
		".1.3.6.1.4.1.3902.3.100.23":        {"platform": stringPlatform, "model": "ZXR10"},
		".1.3.6.1.4.1.3902.3.100.405":       {"platform": stringPlatform, "model": "ZX10"},
		".1.3.6.1.4.1.3902.3.600.3.1.604":   {"platform": stringPlatform, "model": "ZXR10 9900"},
		".1.3.6.1.4.1.3902.3.600.3.1.724":   {"platform": stringPlatform, "model": "ZXR10 5960"},
	}

	data, ok := oidMap[sysObjId]
	if !ok {
		return &devicemodel.DeviceModel{
			Platform:     platform.TPLink,
			Manufacturer: manufacturer.TPLink,
			DeviceModel:  devicemodel.UnknownDeviceModel,
		}
	}

	return &devicemodel.DeviceModel{
		Platform:     platform.Platform(data["platform"]),
		Manufacturer: manufacturer.TPLink,
		DeviceModel:  data["model"],
	}

}
