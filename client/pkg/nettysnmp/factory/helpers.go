package factory

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"

	"github.com/gosnmp/gosnmp"
	"github.com/samber/lo"
)

var vlanPattern = regexp.MustCompile(`(vl|Vl)[A-Za-z]*(\d+)`)

func ExtractInteger(oid string, values []gosnmp.SnmpPDU) map[string]uint64 {
	result := make(map[string]uint64, len(values))
	for _, v := range values {
		snmpIndex := getSnmpIndex(oid, &v)
		result[snmpIndex] = gosnmp.ToBigInt(v.Value).Uint64()
	}
	return result
}

func ExtractIntegerWithShift(oid string, location int, values []gosnmp.SnmpPDU) map[string]uint64 {
	result := make(map[string]uint64, len(values))
	for _, v := range values {
		snmpIndex := shiftSnmpIndex(getSnmpIndex(oid, &v), location)
		result[snmpIndex] = gosnmp.ToBigInt(v.Value).Uint64()
	}
	return result
}

func extractIfIndex(oid string, values []gosnmp.SnmpPDU) map[string]string {
	result := make(map[string]string, len(values))
	for _, v := range values {
		snmpIndex := getSnmpIndex(oid, &v)
		value := "." + strconv.Itoa(int(gosnmp.ToBigInt(v.Value).Uint64()))
		result[value] = snmpIndex
	}
	return result
}

func ExtractString(oid string, values []gosnmp.SnmpPDU) map[string]string {
	result := make(map[string]string, len(values))
	for _, v := range values {
		snmpIndex := getSnmpIndex(oid, &v)
		result[snmpIndex] = fmt.Sprintf("%s", v.Value)
	}
	return result
}

func ExtractStringWithShift(oid string, location int, values []gosnmp.SnmpPDU) map[string]string {
	result := make(map[string]string, len(values))
	for _, v := range values {
		snmpIndex := shiftSnmpIndex(getSnmpIndex(oid, &v), location)
		result[snmpIndex] = fmt.Sprintf("%s", v.Value)
	}
	return result
}

func ExtractMacAddress(oid string, values []gosnmp.SnmpPDU) map[string]string {
	result := make(map[string]string, len(values))
	for _, v := range values {
		snmpIndex := getSnmpIndex(oid, &v)
		result[snmpIndex] = hex2mac(v.Value.([]byte))
	}
	return result
}

func ExtractMacAddressWithShift(oid string, location int, values []gosnmp.SnmpPDU) map[string]string {
	result := make(map[string]string, len(values))
	for _, v := range values {
		snmpIndex := shiftSnmpIndex(getSnmpIndex(oid, &v), location)
		result[snmpIndex] = hex2mac(v.Value.([]byte))
	}
	return result
}

func getSnmpIndex(oid string, values *gosnmp.SnmpPDU) string {
	return strings.Split(values.Name, oid)[1]
}

// shift snmp index is used for special case for lldp neighbor
// lldpRemChassisID: snmp_index value the last of 2 is local portID's index(last of 1)
func shiftSnmpIndex(index string, location int) string {
	splitValues := strings.Split(index, ".")
	if location < 0 {
		return splitValues[len(splitValues)+location]
	}
	return splitValues[location]
}

func hex2mac(hex []byte) string {
	if len(hex) != 6 {
		return ""
	}
	hexParts := make([]string, 6)
	for i, b := range hex {
		hexParts[i] = fmt.Sprintf("%02x", b)
	}
	return strings.Join(hexParts, ":")
}

func buildOidWithIndex(oid string, index []string) []string {
	results := make([]string, len(index))
	for i, v := range index {
		results[i] = fmt.Sprintf("%s.%s", oid, v)
	}
	return results
}

func netmaskToLength(netmask string) int {
	if strings.Contains(netmask, ".") {
		stringMask := net.IPMask(net.ParseIP(netmask).To4())
		length, _ := stringMask.Size()
		return length
	}
	stringMask := net.IPMask(net.ParseIP(netmask).To16())
	length, _ := stringMask.Size()
	return length
}

// IP-MIB: ipNetToMediaPhysAddress.102.10.90.200.120 = INTEGER: 102
// last 5 is ifIndex
// last 4 is address
func getIfIndexAndAddress(snmpIndex string) (ifIndex uint64, address string) {
	splitData := strings.Split(snmpIndex, ".")
	xLast4 := splitData[len(splitData)-4:]
	xLast5 := splitData[len(splitData)-5]
	address = strings.Join(xLast4, ".")
	_ifIndex, _ := strconv.Atoi(xLast5)
	ifIndex = uint64(_ifIndex)
	return ifIndex, address
}

type VlanIpRange struct {
	VlanId  uint32
	Range   string
	Gateway string
}

func extractVlanId(ifName string) (uint32, bool) {
	if match := vlanPattern.FindStringSubmatch(ifName); match != nil {
		if len(match) == 3 {
			if num, err := strconv.Atoi(match[2]); err == nil {
				return uint32(num), true
			}
			return 0, false
		}
	}
	return 0, false
}

func ipAddrToNet(addr string) (string, error) {
	if addr == "" {
		return "", nil
	}
	ip, ipNet, err := net.ParseCIDR(addr)
	if err != nil {
		return "", err
	}
	networkIp := ip.Mask(ipNet.Mask)
	size, _ := ipNet.Mask.Size()
	networkStr := fmt.Sprintf("%s/%d", networkIp.String(), size)
	return networkStr, nil
}

func getIfVlanIpRange(interfaces []*DeviceInterface) map[uint64]*VlanIpRange {
	var results = make(map[uint64]*VlanIpRange)
	for _, iface := range interfaces {
		if iface.IfType == "portVirtual" {
			vlanIpRange := &VlanIpRange{}
			vlan, ok := extractVlanId(iface.IfName)
			if ok {
				vlanIpRange.VlanId = vlan
			}
			inet, err := ipAddrToNet(iface.IfIpAddress)
			if err == nil {
				vlanIpRange.Range = inet
				vlanIpRange.Gateway = iface.IfIpAddress
			}
			if vlan != 0 && inet != "" {
				results[iface.IfIndex] = vlanIpRange
			}
		}
	}
	return results
}

func EnrichArpInfo(arp []*ArpItem, interfaces []*DeviceInterface) []*ArpItem {
	for len(arp) <= 0 || len(interfaces) <= 0 {
		return arp
	}
	ifVlanIpRange := getIfVlanIpRange(interfaces)
	for _, item := range arp {
		inner := ifVlanIpRange[item.IfIndex]
		item.VlanId = inner.VlanId
		item.Range = inner.Range
	}
	return arp
}

func EnrichVlanInfo(vlan []*VlanItem, interfaces []*DeviceInterface) []*VlanItem {
	if len(vlan) <= 0 || len(interfaces) <= 0 {
		return vlan
	}
	ifVlanIpRange := getIfVlanIpRange(interfaces)
	for _, vl := range vlan {
		ipRange := ifVlanIpRange[vl.IfIndex]
		vl.Range = ipRange.Range
		vl.Gateway = ipRange.Gateway
	}
	return vlan
}

func lldpNeighborInterfaces(lldp []*LldpNeighbor) []string {
	results := make([]string, len(lldp))
	for _, item := range lldp {
		results = append(results, item.LocalIfName)
	}
	return results
}

func RemoveNonLocalMacAddress(mac *map[uint64][]string, interfaces []*DeviceInterface, lldp []*LldpNeighbor) *map[uint64][]string {
	interfaceNames := make(map[uint64]string)
	interfaceTypes := make(map[uint64]string)
	for _, iface := range interfaces {
		interfaceNames[iface.IfIndex] = iface.IfName
		interfaceTypes[iface.IfIndex] = iface.IfType
	}
	lldpInterfaces := lldpNeighborInterfaces(lldp)
	for index := range *mac {
		if lo.Contains(lldpInterfaces, interfaceNames[index]) || interfaceTypes[index] != "ethernetCsmacd" {
			delete(*mac, index)
		}
	}
	return mac
}

func EnrichMacAddress(mac *map[uint64][]string, interfaces []*DeviceInterface, lldp []*LldpNeighbor, arp []*ArpItem) []*MacAddressItem {
	results := make([]*MacAddressItem, 0)
	interfaceMapping := make(map[uint64]*DeviceInterface)
	arpMapping := make(map[string]*ArpItem)
	for _, iface := range interfaces {
		interfaceMapping[iface.IfIndex] = iface
	}
	for _, arpItem := range arp {
		arpMapping[arpItem.MacAddress] = arpItem
	}
	lldpInterfaces := lldpNeighborInterfaces(lldp)
	for index, value := range *mac {
		if len(value) == 0 {
			continue
		} else if lo.Contains(lldpInterfaces, interfaceMapping[index].IfName) || interfaceMapping[index].IfType != "ethernetCsmacd" {
			// remove mac address is not learned from access port.
			continue
		} else {
			for _, v := range value {
				ip := "0.0.0.0"
				vlanId := uint32(0)
				if item, ok := arpMapping[v]; !ok {
					continue
				} else {
					ip = item.IpAddress
					vlanId = item.VlanId
				}
				results = append(results, &MacAddressItem{
					MacAddress: v,
					IfIndex:    index,
					IfName:     interfaceMapping[index].IfName,
					IfDescr:    interfaceMapping[index].IfDescr,
					IpAddress:  ip,
					VlanId:     vlanId,
				})
			}
		}
	}
	return results
}
