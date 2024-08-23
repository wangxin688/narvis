package factory

import (
	"fmt"
	"strings"

	"github.com/gosnmp/gosnmp"
)

func extractInteger(oid string, values []gosnmp.SnmpPDU) map[string]uint64 {
	result := make(map[string]uint64, len(values))
	for _, v := range values {
		snmpIndex := getSnmpIndex(oid, &v)
		result[snmpIndex] = gosnmp.ToBigInt(v.Value).Uint64()
	}
	return result
}

func extractIntegerWithShift(oid string, location int, values []gosnmp.SnmpPDU) map[string]uint64 {
	result := make(map[string]uint64, len(values))
	for _, v := range values {
		snmpIndex := shiftSnmpIndex(getSnmpIndex(oid, &v), location)
		result[snmpIndex] = gosnmp.ToBigInt(v.Value).Uint64()
	}
	return result
}

func extractString(oid string, values []gosnmp.SnmpPDU) map[string]string {
	result := make(map[string]string, len(values))
	for _, v := range values {
		snmpIndex := getSnmpIndex(oid, &v)
		result[snmpIndex] = fmt.Sprintf("%s", v.Value)
	}
	return result
}

func extractStringWithShift(oid string, location int, values []gosnmp.SnmpPDU) map[string]string {
	result := make(map[string]string, len(values))
	for _, v := range values {
		snmpIndex := shiftSnmpIndex(getSnmpIndex(oid, &v), location)
		result[snmpIndex] = fmt.Sprintf("%s", v.Value)
	}
	return result
}

func extractMacAddress(oid string, values []gosnmp.SnmpPDU) map[string]string {
	result := make(map[string]string, len(values))
	for _, v := range values {
		snmpIndex := getSnmpIndex(oid, &v)
		result[snmpIndex] = hex2mac(v.Value.([]byte))
	}
	return result
}

func extractMacAddressWithShift(oid string, location int, values []gosnmp.SnmpPDU) map[string]string {
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
	return strings.Split(index, ".")[location]
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
