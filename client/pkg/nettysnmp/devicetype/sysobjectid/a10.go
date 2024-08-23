package sysobjectid

import (
	"github.com/wangxin688/narvis/client/pkg/nettysnmp/devicetype"
	"github.com/wangxin688/narvis/intend/manufacturer"
	"github.com/wangxin688/narvis/intend/platform"
)

func A10DeviceType(sysObjId string) *devicetype.DeviceType {
	stringPlatform := string(platform.A10)
	oidMap := map[string]map[string]string{
		"1.3.6.1.4.1.22610.1.3.1":  {"platform": stringPlatform, "model": "AX2100"},
		"1.3.6.1.4.1.22610.1.3.10": {"platform": stringPlatform, "model": "AX3000"},
		"1.3.6.1.4.1.22610.1.3.11": {"platform": stringPlatform, "model": "HITACHIBLADESERVER"},
		"1.3.6.1.4.1.22610.1.3.12": {"platform": stringPlatform, "model": "AX5100"},
		"1.3.6.1.4.1.22610.1.3.13": {"platform": stringPlatform, "model": "SOFTAX"},
		"1.3.6.1.4.1.22610.1.3.14": {"platform": stringPlatform, "model": "AX 3030 ADC"},
		"1.3.6.1.4.1.22610.1.3.15": {"platform": stringPlatform, "model": "AX 1030"},
		"1.3.6.1.4.1.22610.1.3.16": {"platform": stringPlatform, "model": "AX3200-12"},
		"1.3.6.1.4.1.22610.1.3.17": {"platform": stringPlatform, "model": "AX3400"},
		"1.3.6.1.4.1.22610.1.3.18": {"platform": stringPlatform, "model": "AX3530"},
		"1.3.6.1.4.1.22610.1.3.19": {"platform": stringPlatform, "model": "AX5630"},
		"1.3.6.1.4.1.22610.1.3.2":  {"platform": stringPlatform, "model": "AX3100"},
		"1.3.6.1.4.1.22610.1.3.20": {"platform": stringPlatform, "model": "TH6430"},
		"1.3.6.1.4.1.22610.1.3.21": {"platform": stringPlatform, "model": "TH5430"},
		"1.3.6.1.4.1.22610.1.3.22": {"platform": stringPlatform, "model": "THUNDER 3030S"},
		"1.3.6.1.4.1.22610.1.3.23": {"platform": stringPlatform, "model": "THUNDER SERIES 1030S"},
		"1.3.6.1.4.1.22610.1.3.24": {"platform": stringPlatform, "model": "THUNDER SERIES 930"},
		"1.3.6.1.4.1.22610.1.3.25": {"platform": stringPlatform, "model": "TH4430"},
		"1.3.6.1.4.1.22610.1.3.26": {"platform": stringPlatform, "model": "TH5330"},
		"1.3.6.1.4.1.22610.1.3.27": {"platform": stringPlatform, "model": "THUNDER SERIES 4430S"},
		"1.3.6.1.4.1.22610.1.3.28": {"platform": stringPlatform, "model": "TH5630"},
		"1.3.6.1.4.1.22610.1.3.29": {"platform": stringPlatform, "model": "TH6630"},
		"1.3.6.1.4.1.22610.1.3.3":  {"platform": stringPlatform, "model": "AX3200"},
		"1.3.6.1.4.1.22610.1.3.30": {"platform": stringPlatform, "model": "THUNDER SERIES 3430"},
		"1.3.6.1.4.1.22610.1.3.32": {"platform": stringPlatform, "model": "THUNDER SERIES 4440S"},
		"1.3.6.1.4.1.22610.1.3.34": {"platform": stringPlatform, "model": "THUNDER SERIES 1040S"},
		"1.3.6.1.4.1.22610.1.3.35": {"platform": stringPlatform, "model": "THUNDER SERIES 3040S"},
		"1.3.6.1.4.1.22610.1.3.4":  {"platform": stringPlatform, "model": "AX2200"},
		"1.3.6.1.4.1.22610.1.3.44": {"platform": stringPlatform, "model": "THUNDER SERIES 5430S"},
		"1.3.6.1.4.1.22610.1.3.5":  {"platform": stringPlatform, "model": "AX2000"},
		"1.3.6.1.4.1.22610.1.3.51": {"platform": stringPlatform, "model": "THUNDER SERIES 3350S"},
		"1.3.6.1.4.1.22610.1.3.6":  {"platform": stringPlatform, "model": "AX1000"},
		"1.3.6.1.4.1.22610.1.3.7":  {"platform": stringPlatform, "model": "AX5200"},
		"1.3.6.1.4.1.22610.1.3.8":  {"platform": stringPlatform, "model": "AX2500"},
		"1.3.6.1.4.1.22610.1.3.9":  {"platform": stringPlatform, "model": "AX2600"},
	}

	data, ok := oidMap[sysObjId]
	if !ok {
		return &devicetype.DeviceType{
			Platform:     platform.A10,
			Manufacturer: manufacturer.A10,
			DeviceType:   devicetype.UnknownDeviceType,
		}
	}

	return &devicetype.DeviceType{
		Platform:     platform.Platform(data["platform"]),
		Manufacturer: manufacturer.A10,
		DeviceType:   data["model"],
	}

}
