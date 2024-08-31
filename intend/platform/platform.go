package platform

type Platform string

const (
	CiscoIos      Platform = "ciscoIos"
	CiscoIosXE    Platform = "ciscoXe"
	CiscoIosXR    Platform = "ciscoXr"
	CiscoNexusOS  Platform = "ciscoNxos"
	CiscoASA      Platform = "ciscoAsa"
	Huawei        Platform = "huawei"
	Aruba         Platform = "arubaOs"
	ArubaOSSwitch Platform = "arubaOsswitch"
	Arista        Platform = "aristaEos"
	RuiJie        Platform = "ruijieOs"
	H3C           Platform = "hpComware"
	FortiNet      Platform = "fortinet"
	PaloAlto      Platform = "Panos"
	Juniper       Platform = "Junos"
	Netgear       Platform = "prosafe"
	TPLink        Platform = "jetstream"
	Ruckus        Platform = "fastiron"
	Sangfor       Platform = "sangfor"
	A10           Platform = "a10"
	F5            Platform = "bigip"
	CheckPoint    Platform = "checkpoint"
	ZTE           Platform = "zxrOs"
	Extreme       Platform = "extreme"
	MikroTik      Platform = "mikrotik"
	Unknown       Platform = "unknown"
)

func SupportedPlatform() []Platform {
	return []Platform{
		CiscoIos,
		CiscoIosXE,
		CiscoIosXR,
		CiscoNexusOS,
		Huawei,
		Aruba,
		ArubaOSSwitch,
		Arista,
		RuiJie,
		H3C,
		FortiNet,
		PaloAlto,
		Juniper,
		Netgear,
		TPLink,
		Ruckus,
		Sangfor,
		A10,
		F5,
		CheckPoint,
		ZTE,
		Extreme,
		MikroTik,
		Unknown,
	}
}
