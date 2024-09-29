package manufacturer

import "github.com/wangxin688/narvis/intend/platform"

type Manufacturer string

const (
	Cisco      Manufacturer = "Cisco"
	Huawei     Manufacturer = "Huawei"
	Aruba      Manufacturer = "Aruba"
	Arista     Manufacturer = "Arista"
	H3C        Manufacturer = "H3C"
	RuiJie     Manufacturer = "Ruijie"
	PaloAlto   Manufacturer = "Palo Alto"
	FortiNet   Manufacturer = "Fortinet"
	Netgear    Manufacturer = "Netgear"
	TPLink     Manufacturer = "TP-Link"
	Ruckus     Manufacturer = "Ruckus"
	Juniper    Manufacturer = "Juniper"
	CheckPoint Manufacturer = "Check Point"
	Sangfor    Manufacturer = "Sangfor"
	A10        Manufacturer = "A10 Networks"
	F5         Manufacturer = "F5 Networks"
	Extreme    Manufacturer = "Extreme"
	MikroTik   Manufacturer = "MikroTik"
	Unknown    Manufacturer = "Unknown"
)

func SupportedManufacturer() []Manufacturer {
	return []Manufacturer{
		Cisco,
		Huawei,
		Aruba,
		Arista,
		H3C,
		RuiJie,
		PaloAlto,
		FortiNet,
		Netgear,
		TPLink,
		Ruckus,
		Juniper,
		CheckPoint,
		Sangfor,
		A10,
		F5,
		Extreme,
		MikroTik,
		Unknown,
	}
}

func GetManufacturerByEnterpriseId(etpId string) Manufacturer {
	idMapping := map[string]Manufacturer{
		"2011":  Huawei,
		"56813": "Huawei",
		"9":     Cisco,
		"14823": Aruba,
		"30065": Arista,
		"4881":  RuiJie,
		"61878": H3C,
		"25461": PaloAlto,
		"12356": FortiNet,
		"2636":  Juniper,
		"4562":  Netgear,
		"11863": TPLink,
		"25053": Ruckus,
		"2620":  CheckPoint,
		"30547": Sangfor,
		"22610": A10,
		"40842": A10,
		"12276": F5,
		"1916":  Extreme,
		"14988": MikroTik,
	}
	if etpId == "" {
		return Unknown
	}
	mf, ok := idMapping[etpId]
	if !ok {
		return Unknown
	}
	return mf
}

func GetAllManufacturerPlatform() map[Manufacturer][]platform.Platform {
	return map[Manufacturer][]platform.Platform{
		Cisco:      {platform.CiscoIos, platform.CiscoIosXE, platform.CiscoIosXR, platform.CiscoNexusOS},
		Huawei:     {platform.Huawei},
		Aruba:      {platform.Aruba, platform.ArubaOSSwitch},
		Arista:     {platform.Arista},
		RuiJie:     {platform.RuiJie},
		H3C:        {platform.H3C},
		PaloAlto:   {platform.PaloAlto},
		FortiNet:   {platform.FortiNet},
		Netgear:    {platform.Netgear},
		TPLink:     {platform.TPLink},
		Ruckus:     {platform.Ruckus},
		Juniper:    {platform.Juniper},
		CheckPoint: {platform.CheckPoint},
		Sangfor:    {platform.Sangfor},
		A10:        {platform.A10},
		F5:         {platform.F5},
		Extreme:    {platform.Extreme},
		MikroTik:   {platform.MikroTik},
		Unknown:    {platform.Unknown},
	}
}

func GetManufacturerPlatform(mf Manufacturer) []platform.Platform {
	meta := GetAllManufacturerPlatform()
	return meta[mf]
}

func GetAllPlatformManufacturer(plt platform.Platform) map[platform.Platform]Manufacturer {
	return map[platform.Platform]Manufacturer{
		platform.CiscoIos:      Cisco,
		platform.CiscoIosXE:    Cisco,
		platform.CiscoIosXR:    Cisco,
		platform.CiscoNexusOS:  Cisco,
		platform.Huawei:        Huawei,
		platform.Aruba:         Aruba,
		platform.ArubaOSSwitch: Aruba,
		platform.Arista:        Arista,
		platform.RuiJie:        RuiJie,
		platform.H3C:           H3C,
		platform.PaloAlto:      PaloAlto,
		platform.FortiNet:      FortiNet,
		platform.Netgear:       Netgear,
		platform.TPLink:        TPLink,
		platform.Ruckus:        Ruckus,
		platform.Juniper:       Juniper,
		platform.CheckPoint:    CheckPoint,
		platform.Sangfor:       Sangfor,
		platform.A10:           A10,
		platform.F5:            F5,
		platform.Extreme:       Extreme,
		platform.MikroTik:      MikroTik,
		platform.Unknown:       Unknown,
	}

}

func GetPlatformManufacturer(plt platform.Platform) Manufacturer {
	meta := GetAllPlatformManufacturer(plt)
	return meta[plt]
}
