package sysobjectid

import (
	"github.com/wangxin688/narvis/client/pkg/nettysnmp/devicetype"
	"github.com/wangxin688/narvis/intend/manufacturer"
	"github.com/wangxin688/narvis/intend/platform"
)

func ArubaDeviceType(sysObjId string) *devicetype.DeviceType {
	stringPlatform := string(platform.Aruba)
	oidMap := map[string]map[string]string{
		".1.3.6.1.4.1.14823.1.1.1":    {"platform": stringPlatform, "model": "5000"},
		".1.3.6.1.4.1.14823.1.1.10":   {"platform": stringPlatform, "model": "200"},
		".1.3.6.1.4.1.14823.1.1.11":   {"platform": stringPlatform, "model": "2424"},
		".1.3.6.1.4.1.14823.1.1.12":   {"platform": stringPlatform, "model": "ARUBA 6000-SC3"},
		".1.3.6.1.4.1.14823.1.1.13":   {"platform": stringPlatform, "model": "3200"},
		".1.3.6.1.4.1.14823.1.1.14":   {"platform": stringPlatform, "model": "3200-8"},
		".1.3.6.1.4.1.14823.1.1.15":   {"platform": stringPlatform, "model": "3400"},
		".1.3.6.1.4.1.14823.1.1.16":   {"platform": stringPlatform, "model": "3400-32"},
		".1.3.6.1.4.1.14823.1.1.17":   {"platform": stringPlatform, "model": "3600"},
		".1.3.6.1.4.1.14823.1.1.18":   {"platform": stringPlatform, "model": "3600-64"},
		".1.3.6.1.4.1.14823.1.1.19":   {"platform": stringPlatform, "model": "650-US"},
		".1.3.6.1.4.1.14823.1.1.2":    {"platform": stringPlatform, "model": "2400"},
		".1.3.6.1.4.1.14823.1.1.20":   {"platform": stringPlatform, "model": "651"},
		".1.3.6.1.4.1.14823.1.1.23":   {"platform": stringPlatform, "model": "620"},
		".1.3.6.1.4.1.14823.1.1.24":   {"platform": stringPlatform, "model": "ARUBA S3500-24P"},
		".1.3.6.1.4.1.14823.1.1.25":   {"platform": stringPlatform, "model": "ARUBA S3500-24T"},
		".1.3.6.1.4.1.14823.1.1.26":   {"platform": stringPlatform, "model": "S3500-48P"},
		".1.3.6.1.4.1.14823.1.1.27":   {"platform": stringPlatform, "model": "ARUBA S3500-48T"},
		".1.3.6.1.4.1.14823.1.1.28":   {"platform": stringPlatform, "model": "ARUBA S2500-24P"},
		".1.3.6.1.4.1.14823.1.1.29":   {"platform": stringPlatform, "model": "ARUBA S2500-24T"},
		".1.3.6.1.4.1.14823.1.1.3":    {"platform": stringPlatform, "model": "800"},
		".1.3.6.1.4.1.14823.1.1.30":   {"platform": stringPlatform, "model": "ARUBA S2500-48P"},
		".1.3.6.1.4.1.14823.1.1.31":   {"platform": stringPlatform, "model": "ARUBA S2500-48T"},
		".1.3.6.1.4.1.14823.1.1.32":   {"platform": stringPlatform, "model": "7210"},
		".1.3.6.1.4.1.14823.1.1.33":   {"platform": stringPlatform, "model": "7220"},
		".1.3.6.1.4.1.14823.1.1.34":   {"platform": stringPlatform, "model": "7240"},
		".1.3.6.1.4.1.14823.1.1.35":   {"platform": stringPlatform, "model": "ARUBA S3500-24F"},
		".1.3.6.1.4.1.14823.1.1.36":   {"platform": stringPlatform, "model": "ARUBA S1500-48P"},
		".1.3.6.1.4.1.14823.1.1.37":   {"platform": stringPlatform, "model": "ARUBA S1500-24P"},
		".1.3.6.1.4.1.14823.1.1.38":   {"platform": stringPlatform, "model": "S1500-12P"},
		".1.3.6.1.4.1.14823.1.1.39":   {"platform": stringPlatform, "model": "7005"},
		".1.3.6.1.4.1.14823.1.1.4":    {"platform": stringPlatform, "model": "6000"},
		".1.3.6.1.4.1.14823.1.1.40":   {"platform": stringPlatform, "model": "7010-US"},
		".1.3.6.1.4.1.14823.1.1.41":   {"platform": stringPlatform, "model": "7030"},
		".1.3.6.1.4.1.14823.1.1.42":   {"platform": stringPlatform, "model": "7205"},
		".1.3.6.1.4.1.14823.1.1.43":   {"platform": stringPlatform, "model": "ARUBA A7024"},
		".1.3.6.1.4.1.14823.1.1.44":   {"platform": stringPlatform, "model": "ARUBA A7105"},
		".1.3.6.1.4.1.14823.1.1.45":   {"platform": stringPlatform, "model": "ARUBA A9900"},
		".1.3.6.1.4.1.14823.1.1.46":   {"platform": stringPlatform, "model": "ARUBA A9980"},
		".1.3.6.1.4.1.14823.1.1.47":   {"platform": stringPlatform, "model": "7240XM"},
		".1.3.6.1.4.1.14823.1.1.48":   {"platform": stringPlatform, "model": "7008"},
		".1.3.6.1.4.1.14823.1.1.5":    {"platform": stringPlatform, "model": "2450"},
		".1.3.6.1.4.1.14823.1.1.50":   {"platform": stringPlatform, "model": "MOBILITY MASTER VIRTUAL APPLIANCE"},
		".1.3.6.1.4.1.14823.1.1.52":   {"platform": stringPlatform, "model": "MOBILITY CONTROLLER VIRTUAL APPLIANCE"},
		".1.3.6.1.4.1.14823.1.1.54":   {"platform": stringPlatform, "model": "MOBILITY MASTER VIRTUAL APPLIANCE"},
		".1.3.6.1.4.1.14823.1.1.55":   {"platform": stringPlatform, "model": "MOBILITY MASTER VIRTUAL APPLIANCE"},
		".1.3.6.1.4.1.14823.1.1.6":    {"platform": stringPlatform, "model": "850"},
		".1.3.6.1.4.1.14823.1.1.7":    {"platform": stringPlatform, "model": "2400E"},
		".1.3.6.1.4.1.14823.1.1.8":    {"platform": stringPlatform, "model": "800E"},
		".1.3.6.1.4.1.14823.1.1.9":    {"platform": stringPlatform, "model": "804"},
		".1.3.6.1.4.1.14823.1.1.99":   {"platform": stringPlatform, "model": "2400-F"},
		".1.3.6.1.4.1.14823.1.1.9999": {"platform": stringPlatform, "model": "6000"},
		".1.3.6.1.4.1.14823.1.2.1":    {"platform": stringPlatform, "model": "ARUBA A50"},
		".1.3.6.1.4.1.14823.1.2.10":   {"platform": stringPlatform, "model": "ARUBA AP-80M"},
		".1.3.6.1.4.1.14823.1.2.102":  {"platform": stringPlatform, "model": "AP-535"},
		".1.3.6.1.4.1.14823.1.2.107":  {"platform": stringPlatform, "model": "IAP-515"},
		".1.3.6.1.4.1.14823.1.2.11":   {"platform": stringPlatform, "model": "ARUBA AP-WG102"},
		".1.3.6.1.4.1.14823.1.2.111":  {"platform": stringPlatform, "model": "AP-505"},
		".1.3.6.1.4.1.14823.1.2.114":  {"platform": stringPlatform, "model": "AP-575"},
		".1.3.6.1.4.1.14823.1.2.115":  {"platform": stringPlatform, "model": "AP-577"},
		".1.3.6.1.4.1.14823.1.2.12":   {"platform": stringPlatform, "model": "ARUBA AP-40"},
		".1.3.6.1.4.1.14823.1.2.13":   {"platform": stringPlatform, "model": "ARUBA AP-41"},
		".1.3.6.1.4.1.14823.1.2.14":   {"platform": stringPlatform, "model": "ARUBA AP-65"},
		".1.3.6.1.4.1.14823.1.2.15":   {"platform": stringPlatform, "model": "ARUBA AP-MW1700"},
		".1.3.6.1.4.1.14823.1.2.16":   {"platform": stringPlatform, "model": "ARUBA AP-DUOWJ"},
		".1.3.6.1.4.1.14823.1.2.17":   {"platform": stringPlatform, "model": "ARUBA AP-DUO"},
		".1.3.6.1.4.1.14823.1.2.18":   {"platform": stringPlatform, "model": "ARUBA AP-80MB"},
		".1.3.6.1.4.1.14823.1.2.19":   {"platform": stringPlatform, "model": "ARUBA AP-80SB"},
		".1.3.6.1.4.1.14823.1.2.2":    {"platform": stringPlatform, "model": "ARUBA A52"},
		".1.3.6.1.4.1.14823.1.2.20":   {"platform": stringPlatform, "model": "ARUBA AP-85"},
		".1.3.6.1.4.1.14823.1.2.21":   {"platform": stringPlatform, "model": "ARUBA AP-124"},
		".1.3.6.1.4.1.14823.1.2.22":   {"platform": stringPlatform, "model": "ARUBA AP-125"},
		".1.3.6.1.4.1.14823.1.2.23":   {"platform": stringPlatform, "model": "ARUBA AP-120"},
		".1.3.6.1.4.1.14823.1.2.24":   {"platform": stringPlatform, "model": "ARUBA AP-121"},
		".1.3.6.1.4.1.14823.1.2.25":   {"platform": stringPlatform, "model": "ARUBA AP-1250"},
		".1.3.6.1.4.1.14823.1.2.26":   {"platform": stringPlatform, "model": "ARUBA AP-120ABG"},
		".1.3.6.1.4.1.14823.1.2.27":   {"platform": stringPlatform, "model": "ARUBA AP-121ABG"},
		".1.3.6.1.4.1.14823.1.2.28":   {"platform": stringPlatform, "model": "ARUBA AP-124ABG"},
		".1.3.6.1.4.1.14823.1.2.29":   {"platform": stringPlatform, "model": "ARUBA AP-125ABG"},
		".1.3.6.1.4.1.14823.1.2.3":    {"platform": stringPlatform, "model": "ARUBA AP-60"},
		".1.3.6.1.4.1.14823.1.2.30":   {"platform": stringPlatform, "model": "ARUBA RAP-5WN"},
		".1.3.6.1.4.1.14823.1.2.31":   {"platform": stringPlatform, "model": "ARUBA RAP-5"},
		".1.3.6.1.4.1.14823.1.2.32":   {"platform": stringPlatform, "model": "ARUBA RAP-2WG"},
		".1.3.6.1.4.1.14823.1.2.33":   {"platform": stringPlatform, "model": "ARUBA RESERVED4 ACCESS POINT"},
		".1.3.6.1.4.1.14823.1.2.34":   {"platform": stringPlatform, "model": "IAP-105"},
		".1.3.6.1.4.1.14823.1.2.35":   {"platform": stringPlatform, "model": "ARUBA AP-65WB"},
		".1.3.6.1.4.1.14823.1.2.36":   {"platform": stringPlatform, "model": "ARUBA AP-651"},
		".1.3.6.1.4.1.14823.1.2.37":   {"platform": stringPlatform, "model": "ARUBA RESERVED6 ACCESS POINT"},
		".1.3.6.1.4.1.14823.1.2.38":   {"platform": stringPlatform, "model": "ARUBA AP-60P"},
		".1.3.6.1.4.1.14823.1.2.39":   {"platform": stringPlatform, "model": "ARUBA RESERVED7 ACCESS POINT"},
		".1.3.6.1.4.1.14823.1.2.4":    {"platform": stringPlatform, "model": "ARUBA AP-61"},
		".1.3.6.1.4.1.14823.1.2.40":   {"platform": stringPlatform, "model": "ARUBA IAP-92"},
		".1.3.6.1.4.1.14823.1.2.41":   {"platform": stringPlatform, "model": "ARUBA IAP-93"},
		".1.3.6.1.4.1.14823.1.2.42":   {"platform": stringPlatform, "model": "ARUBA AP-68"},
		".1.3.6.1.4.1.14823.1.2.43":   {"platform": stringPlatform, "model": "ARUBA AP-68P"},
		".1.3.6.1.4.1.14823.1.2.44":   {"platform": stringPlatform, "model": "ARUBA AP-175P"},
		".1.3.6.1.4.1.14823.1.2.45":   {"platform": stringPlatform, "model": "ARUBA AP-175AC"},
		".1.3.6.1.4.1.14823.1.2.46":   {"platform": stringPlatform, "model": "ARUBA AP-175DC"},
		".1.3.6.1.4.1.14823.1.2.47":   {"platform": stringPlatform, "model": "IAP-134"},
		".1.3.6.1.4.1.14823.1.2.48":   {"platform": stringPlatform, "model": "IAP-135"},
		".1.3.6.1.4.1.14823.1.2.49":   {"platform": stringPlatform, "model": "ARUBA RESERVED8 ACCESS POINT"},
		".1.3.6.1.4.1.14823.1.2.5":    {"platform": stringPlatform, "model": "ARUBA AP-70"},
		".1.3.6.1.4.1.14823.1.2.50":   {"platform": stringPlatform, "model": "ARUBA AP-93H"},
		".1.3.6.1.4.1.14823.1.2.51":   {"platform": stringPlatform, "model": "ARUBA RAP-3WN"},
		".1.3.6.1.4.1.14823.1.2.52":   {"platform": stringPlatform, "model": "ARUBA RAP-3WNP"},
		".1.3.6.1.4.1.14823.1.2.53":   {"platform": stringPlatform, "model": "IAP-104"},
		".1.3.6.1.4.1.14823.1.2.54":   {"platform": stringPlatform, "model": "ARUBA RAP-155"},
		".1.3.6.1.4.1.14823.1.2.55":   {"platform": stringPlatform, "model": "ARUBA RAP-155P"},
		".1.3.6.1.4.1.14823.1.2.56":   {"platform": stringPlatform, "model": "ARUBA RAP-108"},
		".1.3.6.1.4.1.14823.1.2.57":   {"platform": stringPlatform, "model": "ARUBA RAP-109"},
		".1.3.6.1.4.1.14823.1.2.58":   {"platform": stringPlatform, "model": "IAP-224"},
		".1.3.6.1.4.1.14823.1.2.59":   {"platform": stringPlatform, "model": "IAP-225"},
		".1.3.6.1.4.1.14823.1.2.6":    {"platform": stringPlatform, "model": "ARUBA AP-61-WJ"},
		".1.3.6.1.4.1.14823.1.2.60":   {"platform": stringPlatform, "model": "IAP-114"},
		".1.3.6.1.4.1.14823.1.2.61":   {"platform": stringPlatform, "model": "IAP-115"},
		".1.3.6.1.4.1.14823.1.2.62":   {"platform": stringPlatform, "model": "ARUBA RAP-109L"},
		".1.3.6.1.4.1.14823.1.2.63":   {"platform": stringPlatform, "model": "IAP-274"},
		".1.3.6.1.4.1.14823.1.2.64":   {"platform": stringPlatform, "model": "IAP-275"},
		".1.3.6.1.4.1.14823.1.2.65":   {"platform": stringPlatform, "model": "ARUBA AP-214A"},
		".1.3.6.1.4.1.14823.1.2.66":   {"platform": stringPlatform, "model": "ARUBA AP-215A"},
		".1.3.6.1.4.1.14823.1.2.67":   {"platform": stringPlatform, "model": "IAP-204"},
		".1.3.6.1.4.1.14823.1.2.68":   {"platform": stringPlatform, "model": "IAP-205"},
		".1.3.6.1.4.1.14823.1.2.69":   {"platform": stringPlatform, "model": "IAP-103"},
		".1.3.6.1.4.1.14823.1.2.7":    {"platform": stringPlatform, "model": "ARUBA A2E"},
		".1.3.6.1.4.1.14823.1.2.70":   {"platform": stringPlatform, "model": "ARUBA AP-103H"},
		".1.3.6.1.4.1.14823.1.2.71":   {"platform": stringPlatform, "model": "IAP VIRTUAL CONTROLLER"},
		".1.3.6.1.4.1.14823.1.2.72":   {"platform": stringPlatform, "model": "ARUBA AP-277"},
		".1.3.6.1.4.1.14823.1.2.73":   {"platform": stringPlatform, "model": "IAP-214"},
		".1.3.6.1.4.1.14823.1.2.74":   {"platform": stringPlatform, "model": "IAP-215"},
		".1.3.6.1.4.1.14823.1.2.75":   {"platform": stringPlatform, "model": "IAP-228"},
		".1.3.6.1.4.1.14823.1.2.76":   {"platform": stringPlatform, "model": "IAP-205H"},
		".1.3.6.1.4.1.14823.1.2.77":   {"platform": stringPlatform, "model": "IAP-324"},
		".1.3.6.1.4.1.14823.1.2.78":   {"platform": stringPlatform, "model": "IAP-325"},
		".1.3.6.1.4.1.14823.1.2.79":   {"platform": stringPlatform, "model": "IAP-334"},
		".1.3.6.1.4.1.14823.1.2.8":    {"platform": stringPlatform, "model": "ARUBA AP-1200"},
		".1.3.6.1.4.1.14823.1.2.80":   {"platform": stringPlatform, "model": "IAP-335"},
		".1.3.6.1.4.1.14823.1.2.81":   {"platform": stringPlatform, "model": "IAP-314"},
		".1.3.6.1.4.1.14823.1.2.82":   {"platform": stringPlatform, "model": "IAP-315"},
		".1.3.6.1.4.1.14823.1.2.84":   {"platform": stringPlatform, "model": "IAP-207"},
		".1.3.6.1.4.1.14823.1.2.85":   {"platform": stringPlatform, "model": "IAP-304"},
		".1.3.6.1.4.1.14823.1.2.86":   {"platform": stringPlatform, "model": "IAP-305"},
		".1.3.6.1.4.1.14823.1.2.87":   {"platform": stringPlatform, "model": "IAP-303H"},
		".1.3.6.1.4.1.14823.1.2.88":   {"platform": stringPlatform, "model": "IAP-365"},
		".1.3.6.1.4.1.14823.1.2.9":    {"platform": stringPlatform, "model": "ARUBA AP-80S"},
		".1.3.6.1.4.1.14823.1.2.9999": {"platform": stringPlatform, "model": "ARUBA AP-UNDEFINED"},
		".1.3.6.1.4.1.14823.1.3.1":    {"platform": stringPlatform, "model": "ZMASTER"},
		".1.3.6.1.4.1.14823.1.6.1":    {"platform": stringPlatform, "model": "C2000"},
		".1.3.6.1.4.1.14823.1.1.59":   {"platform": stringPlatform, "model": "Aruba9240"},
	}

	data, ok := oidMap[sysObjId]
	if !ok {
		return &devicetype.DeviceType{
			Platform:     platform.Aruba,
			Manufacturer: manufacturer.Aruba,
			DeviceType:   devicetype.UnknownDeviceType,
		}
	}

	return &devicetype.DeviceType{
		Platform:     platform.Platform(data["platform"]),
		Manufacturer: manufacturer.Aruba,
		DeviceType:   data["model"],
	}

}
