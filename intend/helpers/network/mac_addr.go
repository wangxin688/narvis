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

// package network mac_addr provide various method for MacAddress processing

package network

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// MacAddressValidator takes a mac address string and returns a normalized mac address
// string (xx:xx:xx:xx:xx:xx) or an error if the mac address is invalid.
// if the mac address is empty, it will return an empty string and no error.
func MacAddressValidator(mac string) (string, error) {
	if mac == "" {
		return "", nil
	}
	mac = strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(mac, "-", ""), ".", ""), ":", ""))
	if len(mac) != 12 {
		return "", fmt.Errorf("invalid mac address %s, length: %d", mac, len(mac))
	}
	re := regexp.MustCompile("^[0-9a-fA-F]{12}$")
	if !re.MatchString(mac) {
		return "", fmt.Errorf("invalid mac address: %s", mac)
	}
	// no for loop here because of performance
	return mac[:2] + ":" + mac[2:4] + ":" + mac[4:6] + ":" + mac[6:8] + ":" + mac[8:10] + ":" + mac[10:], nil
}

// Hex2Mac converts a 6-byte slice into a MAC address string formatted as "xx:xx:xx:xx:xx:xx".
// If the input byte slice is not exactly 6 bytes long, an empty string is returned.
func Hex2Mac(hex []byte) string {
	if len(hex) != 6 {
		return ""
	}
	hexParts := make([]string, 6)
	for i, b := range hex {
		hexParts[i] = fmt.Sprintf("%02x", b)
	}
	return strings.Join(hexParts, ":")
}

// convert snmpIndex octet mac string to hex mac string
// support case 1, use snmpIndex directly: `.101.88.24.36.0.1`
// support case 2, extract partial of snmpIndex without prefix `101.88.24.36.0.1`
func OctetString2HexStringMac(snmpIndex string) string {
	macLength := 6
	if strings.HasPrefix(snmpIndex, ".") {
		snmpIndex = strings.TrimPrefix(snmpIndex, ".")
	}
	parts := strings.Split(snmpIndex, ".")
	if len(parts) != macLength {
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
