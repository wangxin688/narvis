package device360_biz

import (
	"regexp"

	"github.com/samber/lo"
	infra_sc "github.com/wangxin688/narvis/server/features/infra/schemas"
)

// func GetDevicePanel(deviceId string) (*schemas.DevicePanel, error) {

// 	interfaces, err := infra_biz.NewDeviceService().GetDeviceInterfaces(deviceId)
// 	if err != nil {
// 		return nil, err
// 	}
// 	tmpInterfaces := excludeInterfaces(interfaces)

// }

func excludeInterfaces(interfaces []*infra_sc.DeviceInterface) []*infra_sc.DeviceInterface {

	mgmtInterfaces := mgmtInterfaces()
	results := make([]*infra_sc.DeviceInterface, 0)
	for _, item := range interfaces {
		if !lo.Contains(mgmtInterfaces, item.IfName) && excludeNotEthernet(item.IfName) {
			results = append(results, item)
		}
	}
	return results
}

func excludeNotEthernet(ifName string) bool {
	rePattern := `\b(?:\w+|\d+)(?:/\d+){1,4}\b"`
	re := regexp.MustCompile(rePattern)
	return re.MatchString(ifName)
}

func mgmtInterfaces() []string {
	return []string{
		"console",
		"stack",
		"blue",         // cisco bluetooth
		"fastethernet", // cisco mgmt interface
		"mgmt",         // huawei ce 6855
		"ethernet0/0/0/0",
		"GigabitEthernet0/0/0",     // me60 mgmt
		"M-GigabitEthernet0/0/0",   // h3c mgmt
		"MethGigabitEthernet0/0/0", // arista mgmt
		"MEth0/0/1",                // huawei mgmt
		"Ethernet0/0/0",            // huawei mgmt
	}
}
