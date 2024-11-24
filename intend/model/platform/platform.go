// Copyright 2024 wangxin.jeffry@gmail.com
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// package platform define built-in supported manufacturer's platform
// platform decide which ssh driver to use, which commands/templates for configuration or parsing
package platform

type Platform string

const (
	CiscoIos      Platform = "ciscoIos"
	CiscoIosXE    Platform = "ciscoXe"
	CiscoIosXR    Platform = "ciscoXr"
	CiscoNexusOS  Platform = "ciscoNxos"
	CiscoASA      Platform = "ciscoAsa"
	Huawei        Platform = "huaweiVrp"
	HuaweiCE      Platform = "huaweiCE"
	HuaweiFM      Platform = "huaweiFM"
	Aruba         Platform = "arubaOs"
	ArubaOSSwitch Platform = "arubaOsswitch"
	Arista        Platform = "arista"
	RuiJie        Platform = "ruijie"
	H3C           Platform = "h3c"
	FortiNet      Platform = "fortinet"
	PaloAlto      Platform = "Panos"
	Juniper       Platform = "juniper"
	Netgear       Platform = "netgear"
	TPLink        Platform = "tp_link"
	Ruckus        Platform = "fastiron"
	F5            Platform = "bigip"
	CheckPoint    Platform = "checkpoint"
	Extreme       Platform = "extreme"
	MikroTik      Platform = "mikrotik"
	Unknown       Platform = "unknown"
)

func SupportPlatform() []Platform {
	return []Platform{
		CiscoIos, CiscoIosXE, CiscoIosXR, CiscoNexusOS,
		Huawei, HuaweiCE, HuaweiFM,
		Aruba, ArubaOSSwitch,
		Arista,
		RuiJie,
		H3C,
		FortiNet,
		PaloAlto,
		Juniper,
		Netgear,
		TPLink,
		Ruckus,
		F5,
		CheckPoint,
		Extreme,
		MikroTik,
		Unknown,
	}
}
