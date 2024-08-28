package platform

type Platform string

const (
	CiscoIos      Platform = "cisco_ios"
	CiscoIosXE    Platform = "cisco_xe"
	CiscoIosXR    Platform = "cisco_xr"
	CiscoNexusOS  Platform = "cisco_nxos"
	CiscoASA      Platform = "cisco_asa"
	Huawei        Platform = "huawei"
	HuaweiVrp     Platform = "huawei_vrp"
	HuaweiVrpV8   Platform = "huawei_vrpv8"
	Aruba         Platform = "aruba_os"
	ArubaOSSwitch Platform = "aruba_osswitch"
	Arista        Platform = "arista_eos"
	RuiJie        Platform = "ruijie_os"
	H3C           Platform = "hp_comware"
	FortiNet      Platform = "fortinet"
	PaloAlto      Platform = "paloalto_panos"
	Juniper       Platform = "juniper_junos"
	Netgear       Platform = "netgear_prosafe"
	TPLink        Platform = "tplink_jetstream"
	Ruckus        Platform = "ruckus_fastiron"
	Sangfor       Platform = "sangfor"
	A10           Platform = "a10"
	F5            Platform = "f5_bigip"
	CheckPoint    Platform = "checkpoint"
	ZTE           Platform = "zte_zxros"
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
		HuaweiVrp,
		HuaweiVrpV8,
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
