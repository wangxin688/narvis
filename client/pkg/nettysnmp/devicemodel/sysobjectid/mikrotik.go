package sysobjectid

import (
	"github.com/wangxin688/narvis/client/pkg/nettysnmp/devicemodel"
	"github.com/wangxin688/narvis/intend/model/manufacturer"
	"github.com/wangxin688/narvis/intend/platform"
)

func MikroTikDeviceModel(sysObjId string) *devicemodel.DeviceModel {
	stringPlatform := string(platform.MikroTik)
	oidMap := map[string]map[string]string{
		"1.3.6.1.4.1.14988.1": {"platform": stringPlatform, "model": "RB1200"},
	}

	data, ok := oidMap[sysObjId]
	if !ok {
		return &devicemodel.DeviceModel{
			Platform:     platform.MikroTik,
			Manufacturer: manufacturer.MikroTik,
			DeviceModel:  devicemodel.UnknownDeviceModel,
		}
	}

	return &devicemodel.DeviceModel{
		Platform:     platform.Platform(data["platform"]),
		Manufacturer: manufacturer.MikroTik,
		DeviceModel:  data["model"],
	}

}
