package snmp

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"

	"github.com/gosnmp/gosnmp"
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

func StringToHexMac(snmpIndex string) string {
	parts := strings.Split(strings.TrimPrefix(snmpIndex, "."), ".")
	if len(parts) != 6 {
		return ""
	}
	var macAddress strings.Builder
	for i, part := range parts {
		decimal, err := strconv.Atoi(part)
		if err != nil {
			return ""
		}
		hex := decimalToHex(decimal)
		macAddress.WriteString(hex)
		if i < 5 {
			macAddress.WriteString(":")
		}
	}
	return macAddress.String()
}

func decimalToHex(decimal int) string {
	hex := strconv.FormatInt(int64(decimal), 16)
	if len(hex) == 1 {
		hex = "0" + hex
	}
	return hex
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
