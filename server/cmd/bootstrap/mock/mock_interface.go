package mock

import (
	"github.com/wangxin688/narvis/intend/helpers/processor"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/tests/fixtures"
	"gorm.io/gorm"
)

func mockDeviceInterface(db *gorm.DB, siteId, deviceId string) {

	createInterfaces := make([]*models.DeviceInterface, 0)

	interfaces := []struct {
		ifName        string
		ifIndex       uint64
		ifDescr       string
		ifSpeed       uint64
		ifType        string
		ifMtu         uint64
		ifAdminStatus string
		ifOperStatus  string
		ifLastChange  uint64
		ifHighSpeed   uint64
		ifPhysAddr    *string
		ifIpAddress   *string
	}{
		{"eth0", 1, "eth0", 1000, "ethernetCsmacd", 1500, "up", "up", 0, 1000, processor.StringToPtrString(fixtures.RandomMacAddress()), processor.StringToPtrString(fixtures.RandomIpv4())},
		{"eth1", 2, "eth1", 1000, "ethernetCsmacd", 1500, "up", "up", 0, 1000, processor.StringToPtrString(fixtures.RandomMacAddress()), processor.StringToPtrString(fixtures.RandomIpv4())},
		{"eth2", 3, "eth2", 1000, "ethernetCsmacd", 1500, "up", "up", 0, 1000, processor.StringToPtrString(fixtures.RandomMacAddress()), processor.StringToPtrString(fixtures.RandomIpv4())},
		{"eth3", 4, "eth3", 1000, "ethernetCsmacd", 1500, "up", "up", 0, 1000, processor.StringToPtrString(fixtures.RandomMacAddress()), processor.StringToPtrString(fixtures.RandomIpv4())},
		{"eth4", 5, "eth4", 1000, "ethernetCsmacd", 1500, "up", "up", 0, 1000, processor.StringToPtrString(fixtures.RandomMacAddress()), nil},
		{"eth5", 6, "eth5", 1000, "ethernetCsmacd", 1500, "up", "up", 0, 1000, nil, processor.StringToPtrString(fixtures.RandomIpv4())},
		{"eth6", 7, "eth6", 1000, "ethernetCsmacd", 1500, "up", "up", 0, 1000, nil, nil},
		{"eth7", 8, "eth7", 1000, "ethernetCsmacd", 1500, "up", "up", 0, 1000, processor.StringToPtrString(fixtures.RandomMacAddress()), processor.StringToPtrString(fixtures.RandomIpv4())},
		{"Vlan10", 9, "Vlan10", 1000, "propVirtual", 1500, "up", "up", 0, 1000, processor.StringToPtrString(fixtures.RandomMacAddress()), processor.StringToPtrString(fixtures.RandomIpv4())},
		{"Vlan100", 10, "Vlan100", 1000, "propVirtual", 1500, "up", "up", 0, 1000, processor.StringToPtrString(fixtures.RandomMacAddress()), processor.StringToPtrString(fixtures.RandomIpv4())},
		{"GigabitEthernet0/0/1", 11, "GigabitEthernet0/0/1", 1000, "ethernetCsmacd", 1500, "up", "up", 0, 1000, processor.StringToPtrString(fixtures.RandomMacAddress()), processor.StringToPtrString(fixtures.RandomIpv4())},
		{"GigabitEthernet0/0/2", 12, "GigabitEthernet0/0/2", 1000, "ethernetCsmacd", 1500, "up", "up", 0, 1000, processor.StringToPtrString(fixtures.RandomMacAddress()), processor.StringToPtrString(fixtures.RandomIpv4())},
		{"XGigabitEthernet0/0/1", 13, "GigabitEthernet0/0/3", 10000, "ethernetCsmacd", 1500, "up", "up", 0, 10000, processor.StringToPtrString(fixtures.RandomMacAddress()), processor.StringToPtrString(fixtures.RandomIpv4())},
		{"GigabitEthernet0/0/2", 14, "GigabitEthernet0/0/1", 10000, "ethernetCsmacd", 9200, "up", "up", 0, 10000, processor.StringToPtrString(fixtures.RandomMacAddress()), processor.StringToPtrString(fixtures.RandomIpv4())},
	}

	for _, iface := range interfaces {
		createInterfaces = append(createInterfaces, &models.DeviceInterface{
			IfName:        iface.ifName,
			IfIndex:       iface.ifIndex,
			IfDescr:       iface.ifDescr,
			IfSpeed:       iface.ifSpeed,
			IfType:        iface.ifType,
			IfMtu:         iface.ifMtu,
			IfAdminStatus: iface.ifAdminStatus,
			IfOperStatus:  iface.ifOperStatus,
			IfLastChange:  iface.ifLastChange,
			IfHighSpeed:   iface.ifHighSpeed,
			IfPhysAddr:    iface.ifPhysAddr,
			IfIpAddress:   iface.ifIpAddress,
			DeviceId:      deviceId,
			SiteId:        siteId,
		})
	}
	err := gen.DeviceInterface.CreateInBatches(createInterfaces, 100)
	if err != nil {
		panic(err)
	}
}
