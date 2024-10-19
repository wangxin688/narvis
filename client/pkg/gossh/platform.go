package gossh

import (
	"fmt"

	scrapligo_platform "github.com/scrapli/scrapligo/platform"
	"github.com/wangxin688/narvis/intend/platform"
)

func getScrapliPlatform(platform_ string) (string, error) {
	platformMapping := map[string]string{
		string(platform.CiscoIos):     scrapligo_platform.CiscoIosxe,
		string(platform.CiscoIosXE):   scrapligo_platform.CiscoIosxe,
		string(platform.CiscoIosXR):   scrapligo_platform.CiscoIosxr,
		string(platform.CiscoNexusOS): scrapligo_platform.CiscoNxos,
		string(platform.Huawei):       scrapligo_platform.HuaweiVrp,
		string(platform.Aruba):        scrapligo_platform.ArubaWlc,
		string(platform.Arista):       scrapligo_platform.AristaEos,
		string(platform.RuiJie):       scrapligo_platform.RuijieRgos,
		string(platform.H3C):          scrapligo_platform.HpComware,
		string(platform.PaloAlto):     scrapligo_platform.PaloAltoPanos,
		// string(platform.FortiNet): scrapligo_platform.Fortigate,
		// string(platform.Netgear): scrapligo_platform.N,
		string(platform.TPLink): scrapligo_platform.CiscoIosxe,
		// string(platform.Ruckus): scrapligo_platform.Ruckus,
		string(platform.Juniper): scrapligo_platform.JuniperJunos,
		// string(platform.CheckPoint): scrapligo_platform.Checkpoint,
	}
	if platform_, ok := platformMapping[platform_]; ok {
		return platform_, nil
	}
	return "", fmt.Errorf("unsupported ssh platform %s", platform_)

}
