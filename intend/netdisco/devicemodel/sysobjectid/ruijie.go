package sysobjectid

import (
	manufacturer "github.com/wangxin688/narvis/intend/model/manufacturer"
	platform "github.com/wangxin688/narvis/intend/model/platform"
	"github.com/wangxin688/narvis/intend/netdisco/devicemodel"
)

func RuiJieDeviceModel(sysObjId string) *devicemodel.DeviceModel {
	stringPlatform := string(platform.RuiJie)
	oidMap := map[string]map[string]string{
		".1.3.6.1.4.1.4881.1.1.10.1.1":   {"platform": stringPlatform, "model": "S2126G"},
		".1.3.6.1.4.1.4881.1.1.10.1.106": {"platform": stringPlatform, "model": "S2628G-E"},
		".1.3.6.1.4.1.4881.1.1.10.1.11":  {"platform": stringPlatform, "model": "S21-STACKING"},
		".1.3.6.1.4.1.4881.1.1.10.1.12":  {"platform": stringPlatform, "model": "S3550-24"},
		".1.3.6.1.4.1.4881.1.1.10.1.13":  {"platform": stringPlatform, "model": "S3550-48"},
		".1.3.6.1.4.1.4881.1.1.10.1.15":  {"platform": stringPlatform, "model": "S3550-12SFP-GT"},
		".1.3.6.1.4.1.4881.1.1.10.1.16":  {"platform": stringPlatform, "model": "S6806"},
		".1.3.6.1.4.1.4881.1.1.10.1.17":  {"platform": stringPlatform, "model": "S6810"},
		".1.3.6.1.4.1.4881.1.1.10.1.18":  {"platform": stringPlatform, "model": "S2126S"},
		".1.3.6.1.4.1.4881.1.1.10.1.19":  {"platform": stringPlatform, "model": "S2126S-STACKING"},
		".1.3.6.1.4.1.4881.1.1.10.1.2":   {"platform": stringPlatform, "model": "S2126GL3"},
		".1.3.6.1.4.1.4881.1.1.10.1.20":  {"platform": stringPlatform, "model": "S1908PLUS"},
		".1.3.6.1.4.1.4881.1.1.10.1.21":  {"platform": stringPlatform, "model": "S1916PLUS"},
		".1.3.6.1.4.1.4881.1.1.10.1.22":  {"platform": stringPlatform, "model": "S6506"},
		".1.3.6.1.4.1.4881.1.1.10.1.23":  {"platform": stringPlatform, "model": "S2126S-08"},
		".1.3.6.1.4.1.4881.1.1.10.1.24":  {"platform": stringPlatform, "model": "S2126S-16"},
		".1.3.6.1.4.1.4881.1.1.10.1.25":  {"platform": stringPlatform, "model": "S6806E"},
		".1.3.6.1.4.1.4881.1.1.10.1.26":  {"platform": stringPlatform, "model": "S6810E"},
		".1.3.6.1.4.1.4881.1.1.10.1.27":  {"platform": stringPlatform, "model": "S2026G"},
		".1.3.6.1.4.1.4881.1.1.10.1.28":  {"platform": stringPlatform, "model": "S3750-24"},
		".1.3.6.1.4.1.4881.1.1.10.1.29":  {"platform": stringPlatform, "model": "S3750-48"},
		".1.3.6.1.4.1.4881.1.1.10.1.3":   {"platform": stringPlatform, "model": "S2150G"},
		".1.3.6.1.4.1.4881.1.1.10.1.30":  {"platform": stringPlatform, "model": "S2126"},
		".1.3.6.1.4.1.4881.1.1.10.1.31":  {"platform": stringPlatform, "model": "S2126-STACKING"},
		".1.3.6.1.4.1.4881.1.1.10.1.32":  {"platform": stringPlatform, "model": "S2026F"},
		".1.3.6.1.4.1.4881.1.1.10.1.33":  {"platform": stringPlatform, "model": "S3760-48"},
		".1.3.6.1.4.1.4881.1.1.10.1.34":  {"platform": stringPlatform, "model": "S3760-12SFP-GT"},
		".1.3.6.1.4.1.4881.1.1.10.1.35":  {"platform": stringPlatform, "model": "S4009"},
		".1.3.6.1.4.1.4881.1.1.10.1.36":  {"platform": stringPlatform, "model": "S3526"},
		".1.3.6.1.4.1.4881.1.1.10.1.37":  {"platform": stringPlatform, "model": "S3512G"},
		".1.3.6.1.4.1.4881.1.1.10.1.38":  {"platform": stringPlatform, "model": "HCL-12GCS-L3"},
		".1.3.6.1.4.1.4881.1.1.10.1.39":  {"platform": stringPlatform, "model": "HCL-24GS-L3"},
		".1.3.6.1.4.1.4881.1.1.10.1.4":   {"platform": stringPlatform, "model": "S2150GL3"},
		".1.3.6.1.4.1.4881.1.1.10.1.40":  {"platform": stringPlatform, "model": "HCL-48TMS-2S-S"},
		".1.3.6.1.4.1.4881.1.1.10.1.41":  {"platform": stringPlatform, "model": "S5750-24GT-12SFP"},
		".1.3.6.1.4.1.4881.1.1.10.1.42":  {"platform": stringPlatform, "model": "S5750P-24GT-12SFP"},
		".1.3.6.1.4.1.4881.1.1.10.1.43":  {"platform": stringPlatform, "model": "S8606"},
		".1.3.6.1.4.1.4881.1.1.10.1.44":  {"platform": stringPlatform, "model": "S8610"},
		".1.3.6.1.4.1.4881.1.1.10.1.45":  {"platform": stringPlatform, "model": "S9610"},
		".1.3.6.1.4.1.4881.1.1.10.1.46":  {"platform": stringPlatform, "model": "S9620"},
		".1.3.6.1.4.1.4881.1.1.10.1.47":  {"platform": stringPlatform, "model": "S2924"},
		".1.3.6.1.4.1.4881.1.1.10.1.48":  {"platform": stringPlatform, "model": "S3760-24"},
		".1.3.6.1.4.1.4881.1.1.10.1.49":  {"platform": stringPlatform, "model": "S3760-48V2"},
		".1.3.6.1.4.1.4881.1.1.10.1.5":   {"platform": stringPlatform, "model": "S4909"},
		".1.3.6.1.4.1.4881.1.1.10.1.50":  {"platform": stringPlatform, "model": "S3750E-24"},
		".1.3.6.1.4.1.4881.1.1.10.1.51":  {"platform": stringPlatform, "model": "S3750E-48"},
		".1.3.6.1.4.1.4881.1.1.10.1.52":  {"platform": stringPlatform, "model": "S3750E-12SFP-GT"},
		".1.3.6.1.4.1.4881.1.1.10.1.53":  {"platform": stringPlatform, "model": "S5750S-24GT-12SFP"},
		".1.3.6.1.4.1.4881.1.1.10.1.54":  {"platform": stringPlatform, "model": "S2128G"},
		".1.3.6.1.4.1.4881.1.1.10.1.55":  {"platform": stringPlatform, "model": "S2927XG"},
		".1.3.6.1.4.1.4881.1.1.10.1.56":  {"platform": stringPlatform, "model": "S3512GPLUS"},
		".1.3.6.1.4.1.4881.1.1.10.1.57":  {"platform": stringPlatform, "model": "S6604"},
		".1.3.6.1.4.1.4881.1.1.10.1.58":  {"platform": stringPlatform, "model": "S6606"},
		".1.3.6.1.4.1.4881.1.1.10.1.59":  {"platform": stringPlatform, "model": "S6610"},
		".1.3.6.1.4.1.4881.1.1.10.1.6":   {"platform": stringPlatform, "model": "S3550-12G"},
		".1.3.6.1.4.1.4881.1.1.10.1.60":  {"platform": stringPlatform, "model": "S5750-24SFP-12GT"},
		".1.3.6.1.4.1.4881.1.1.10.1.61":  {"platform": stringPlatform, "model": "S5750-48GT-4SFP"},
		".1.3.6.1.4.1.4881.1.1.10.1.62":  {"platform": stringPlatform, "model": "S5750S-48GT-4SFP"},
		".1.3.6.1.4.1.4881.1.1.10.1.63":  {"platform": stringPlatform, "model": "S2328G"},
		".1.3.6.1.4.1.4881.1.1.10.1.64":  {"platform": stringPlatform, "model": "S3250-48"},
		".1.3.6.1.4.1.4881.1.1.10.1.66":  {"platform": stringPlatform, "model": "S2951XG"},
		".1.3.6.1.4.1.4881.1.1.10.1.67":  {"platform": stringPlatform, "model": "S3750-24-UB"},
		".1.3.6.1.4.1.4881.1.1.10.1.68":  {"platform": stringPlatform, "model": "S3750-48-UB"},
		".1.3.6.1.4.1.4881.1.1.10.1.69":  {"platform": stringPlatform, "model": "SCG5510"},
		".1.3.6.1.4.1.4881.1.1.10.1.70":  {"platform": stringPlatform, "model": "S2052G"},
		".1.3.6.1.4.1.4881.1.1.10.1.71":  {"platform": stringPlatform, "model": "S2352G"},
		".1.3.6.1.4.1.4881.1.1.10.1.72":  {"platform": stringPlatform, "model": "S8614"},
		".1.3.6.1.4.1.4881.1.1.10.1.73":  {"platform": stringPlatform, "model": "S5650-24GT-4SFP"},
		".1.3.6.1.4.1.4881.1.1.10.1.74":  {"platform": stringPlatform, "model": "S5650-27XG"},
		".1.3.6.1.4.1.4881.1.1.10.1.75":  {"platform": stringPlatform, "model": "S5650-51XG"},
		".1.3.6.1.4.1.4881.1.1.10.1.76":  {"platform": stringPlatform, "model": "S5450-28GT"},
		".1.3.6.1.4.1.4881.1.1.10.1.77":  {"platform": stringPlatform, "model": "S3760E-24"},
		".1.3.6.1.4.1.4881.1.1.10.1.78":  {"platform": stringPlatform, "model": "S3250P-24"},
		".1.3.6.1.4.1.4881.1.1.10.1.79":  {"platform": stringPlatform, "model": "S2928G"},
		".1.3.6.1.4.1.4881.1.1.10.1.8":   {"platform": stringPlatform, "model": "S3550-24G"},
		".1.3.6.1.4.1.4881.1.1.10.1.80":  {"platform": stringPlatform, "model": "S2952G"},
		".1.3.6.1.4.1.4881.1.1.10.1.81":  {"platform": stringPlatform, "model": "S2028G"},
		".1.3.6.1.4.1.4881.1.1.10.1.82":  {"platform": stringPlatform, "model": "S2528G"},
		".1.3.6.1.4.1.4881.1.1.10.1.83":  {"platform": stringPlatform, "model": "S2552G"},
		".1.3.6.1.4.1.4881.1.1.10.1.84":  {"platform": stringPlatform, "model": "S5750R-48GT-4SFP"},
		".1.3.6.1.4.1.4881.1.1.10.1.85":  {"platform": stringPlatform, "model": "S5750P-48GT-4SFP"},
		".1.3.6.1.4.1.4881.1.1.10.1.86":  {"platform": stringPlatform, "model": "S5750R-24GT-4SFP"},
		".1.3.6.1.4.1.4881.1.1.10.1.87":  {"platform": stringPlatform, "model": "S5750P-24GT-4SFP"},
		".1.3.6.1.4.1.4881.1.1.10.1.88":  {"platform": stringPlatform, "model": "S5750-24GT-4SFP"},
		".1.3.6.1.4.1.4881.1.1.10.1.89":  {"platform": stringPlatform, "model": "S5750S-24GT-4SFP"},
		".1.3.6.1.4.1.4881.1.1.10.1.92":  {"platform": stringPlatform, "model": "S5750-48GT-4SFP-A"},
		".1.3.6.1.4.1.4881.1.1.10.1.93":  {"platform": stringPlatform, "model": "S5750-48GT-4SFP-AP"},
		".1.3.6.1.4.1.4881.1.1.10.1.95":  {"platform": stringPlatform, "model": "NM2X-24ESW"},
		".1.3.6.1.4.1.4881.1.1.10.1.96":  {"platform": stringPlatform, "model": "NM2X-16ESW"},
		".1.3.6.1.4.1.4881.1.1.10.1.98":  {"platform": stringPlatform, "model": "S3760E-24"},
		".1.3.6.1.4.1.4881.1.2.1.1.1":    {"platform": stringPlatform, "model": "R2620"},
		".1.3.6.1.4.1.4881.1.2.1.1.10":   {"platform": stringPlatform, "model": "R2632"},
		".1.3.6.1.4.1.4881.1.2.1.1.11":   {"platform": stringPlatform, "model": "R1762"},
		".1.3.6.1.4.1.4881.1.2.1.1.12":   {"platform": stringPlatform, "model": "RCMS"},
		".1.3.6.1.4.1.4881.1.2.1.1.13":   {"platform": stringPlatform, "model": "HCL-R1762"},
		".1.3.6.1.4.1.4881.1.2.1.1.14":   {"platform": stringPlatform, "model": "HCL-R2632"},
		".1.3.6.1.4.1.4881.1.2.1.1.15":   {"platform": stringPlatform, "model": "HCL-R2692"},
		".1.3.6.1.4.1.4881.1.2.1.1.16":   {"platform": stringPlatform, "model": "HCL-R3642"},
		".1.3.6.1.4.1.4881.1.2.1.1.17":   {"platform": stringPlatform, "model": "HCL-R3662"},
		".1.3.6.1.4.1.4881.1.2.1.1.18":   {"platform": stringPlatform, "model": "R3740"},
		".1.3.6.1.4.1.4881.1.2.1.1.19":   {"platform": stringPlatform, "model": "NBR2000"},
		".1.3.6.1.4.1.4881.1.2.1.1.2":    {"platform": stringPlatform, "model": "R2624"},
		".1.3.6.1.4.1.4881.1.2.1.1.20":   {"platform": stringPlatform, "model": "NBR300"},
		".1.3.6.1.4.1.4881.1.2.1.1.21":   {"platform": stringPlatform, "model": "NBR1200"},
		".1.3.6.1.4.1.4881.1.2.1.1.22":   {"platform": stringPlatform, "model": "NBR1500"},
		".1.3.6.1.4.1.4881.1.2.1.1.23":   {"platform": stringPlatform, "model": "R2716"},
		".1.3.6.1.4.1.4881.1.2.1.1.24":   {"platform": stringPlatform, "model": "R2724"},
		".1.3.6.1.4.1.4881.1.2.1.1.25":   {"platform": stringPlatform, "model": "R3802"},
		".1.3.6.1.4.1.4881.1.2.1.1.26":   {"platform": stringPlatform, "model": "R3804"},
		".1.3.6.1.4.1.4881.1.2.1.1.27":   {"platform": stringPlatform, "model": "RSR50-20"},
		".1.3.6.1.4.1.4881.1.2.1.1.28":   {"platform": stringPlatform, "model": "RSR50-40"},
		".1.3.6.1.4.1.4881.1.2.1.1.29":   {"platform": stringPlatform, "model": "RSR50-80"},
		".1.3.6.1.4.1.4881.1.2.1.1.3":    {"platform": stringPlatform, "model": "R2690"},
		".1.3.6.1.4.1.4881.1.2.1.1.30":   {"platform": stringPlatform, "model": "NPE50-20"},
		".1.3.6.1.4.1.4881.1.2.1.1.31":   {"platform": stringPlatform, "model": "RSR10-02"},
		".1.3.6.1.4.1.4881.1.2.1.1.32":   {"platform": stringPlatform, "model": "RSR20-04"},
		".1.3.6.1.4.1.4881.1.2.1.1.33":   {"platform": stringPlatform, "model": "VPN120"},
		".1.3.6.1.4.1.4881.1.2.1.1.34":   {"platform": stringPlatform, "model": "NPE80"},
		".1.3.6.1.4.1.4881.1.2.1.1.35":   {"platform": stringPlatform, "model": "RSR20-24"},
		".1.3.6.1.4.1.4881.1.2.1.1.36":   {"platform": stringPlatform, "model": "NM2-16ESW"},
		".1.3.6.1.4.1.4881.1.2.1.1.37":   {"platform": stringPlatform, "model": "NM2-24ESW"},
		".1.3.6.1.4.1.4881.1.2.1.1.38":   {"platform": stringPlatform, "model": "NMX-24ESW"},
		".1.3.6.1.4.1.4881.1.2.1.1.39":   {"platform": stringPlatform, "model": "NMX-24ESW-L2"},
		".1.3.6.1.4.1.4881.1.2.1.1.4":    {"platform": stringPlatform, "model": "R2692"},
		".1.3.6.1.4.1.4881.1.2.1.1.40":   {"platform": stringPlatform, "model": "NMX-24ESW-3GEL3"},
		".1.3.6.1.4.1.4881.1.2.1.1.41":   {"platform": stringPlatform, "model": "RSR20-14"},
		".1.3.6.1.4.1.4881.1.2.1.1.42":   {"platform": stringPlatform, "model": "RSR30-44"},
		".1.3.6.1.4.1.4881.1.2.1.1.43":   {"platform": stringPlatform, "model": "R2700V2V3"},
		".1.3.6.1.4.1.4881.1.2.1.1.44":   {"platform": stringPlatform, "model": "R2700V5"},
		".1.3.6.1.4.1.4881.1.2.1.1.45":   {"platform": stringPlatform, "model": "NPE50-40"},
		".1.3.6.1.4.1.4881.1.2.1.1.46":   {"platform": stringPlatform, "model": "RSR20-18"},
		".1.3.6.1.4.1.4881.1.2.1.1.5":    {"platform": stringPlatform, "model": "R3642"},
		".1.3.6.1.4.1.4881.1.2.1.1.6":    {"platform": stringPlatform, "model": "R3662"},
		".1.3.6.1.4.1.4881.1.2.1.1.7":    {"platform": stringPlatform, "model": "NBR1000"},
		".1.3.6.1.4.1.4881.1.2.1.1.8":    {"platform": stringPlatform, "model": "NBR200"},
		".1.3.6.1.4.1.4881.1.2.1.1.9":    {"platform": stringPlatform, "model": "SECVPN100"},
		".1.3.6.1.4.1.4881.1.3.1.1.1":    {"platform": stringPlatform, "model": "WGP500"},
	}

	data, ok := oidMap[sysObjId]
	if !ok {
		return &devicemodel.DeviceModel{
			Platform:     platform.RuiJie,
			Manufacturer: manufacturer.RuiJie,
			DeviceModel:  devicemodel.UnknownDeviceModel,
		}
	}

	return &devicemodel.DeviceModel{
		Platform:     platform.Platform(data["platform"]),
		Manufacturer: manufacturer.RuiJie,
		DeviceModel:  data["model"],
	}

}
