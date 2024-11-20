package infra_biz

import (
	"regexp"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	"github.com/wangxin688/narvis/server/tools/helpers"
)

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

func GetDevicePanel(deviceId string) ([]*schemas.DeviceInterface, error) {

	list, err := gen.DeviceInterface.Where(gen.DeviceInterface.DeviceId.Eq(deviceId),
		gen.DeviceInterface.IfType.Eq("ethernetCsmacd")).Find()

	if err != nil {
		return nil, err
	}
	res := make([]*schemas.DeviceInterface, 0)
	mgmtInterfaces := mgmtInterfaces()
	for _, item := range list {
		if !lo.Contains(mgmtInterfaces, item.IfName) && excludeNotEthernet(item.IfName) {

		}
		res = append(res, &schemas.DeviceInterface{
			Id:            item.Id,
			CreatedAt:     item.CreatedAt,
			UpdatedAt:     item.UpdatedAt,
			IfIndex:       item.IfIndex,
			IfName:        item.IfName,
			IfDescr:       item.IfDescr,
			IfType:        item.IfType,
			IfMtu:         item.IfMtu,
			IfSpeed:       item.IfSpeed,
			IfPhysAddr:    item.IfPhysAddr,
			IfAdminStatus: item.IfAdminStatus,
			IfOperStatus:  item.IfOperStatus,
			IfHighSpeed:   item.IfHighSpeed,
			IfLastChange:  helpers.TimeTicksToDuration(item.IfLastChange),
			IfIpAddress:   item.IfIpAddress,
		})
	}
	return res, nil
}
