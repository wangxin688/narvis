package commands

import (
	"fmt"

	"github.com/wangxin688/narvis/intend/platform"
)

func ShowConfigurationCmd(plat string) (string, error) {
	commandsMapping := map[string]string{
		string(platform.CiscoIos):      "show running-config",
		string(platform.CiscoIosXE):    "show running-config",
		string(platform.CiscoIosXR):    "show running-config",
		string(platform.CiscoNexusOS):  "show running-config",
		string(platform.Huawei):        "display current-configuration",
		string(platform.HuaweiCE):      "display current-configuration",
		string(platform.HuaweiFM):      "display current-configuration",
		string(platform.Aruba):         "show running-config",
		string(platform.ArubaOSSwitch): "show running-config",
		string(platform.Arista):        "show running-config",
		string(platform.RuiJie):        "show running-config",
		string(platform.H3C):           "display current-configuration",
		string(platform.PaloAlto):      "show config running",
		string(platform.FortiNet):      "show full-configuration",
		string(platform.Netgear):       "show running-config",
		string(platform.TPLink):        "show running-config",
		string(platform.Ruckus):        "show running-config",
		string(platform.Juniper):       "show running-config",
		string(platform.CheckPoint):    "show configuration",
		string(platform.Extreme):       "show running-config",
	}
	if command, ok := commandsMapping[plat]; ok {
		return command, nil
	}
	return "", fmt.Errorf("unsupported platform: %s", plat)
}
