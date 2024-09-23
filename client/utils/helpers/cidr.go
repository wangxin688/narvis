package helpers

import (
	"net"
)

func CIDRToIPStrings(cidrStr string) ([]string, error) {
	_, ipNet, err := net.ParseCIDR(cidrStr)
	if err != nil {
		return nil, err
	}
	size, _ := ipNet.Mask.Size()

	var ips []string
	for ip := ipNet.IP.Mask(ipNet.Mask); ipNet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}
	if size >= 30 {
		return ips, nil
	}
	// remove network address and broadcast address
	return ips[1 : len(ips)-1], nil
}
func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
