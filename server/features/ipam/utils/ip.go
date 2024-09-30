package ipam_utils

import (
	"net"
	"net/netip"
	"strings"
)

const IPv4 = "IPv4"
const IPv6 = "IPv6"

func CidrVersion(cidr string) string {
	ip, _, _ := net.ParseCIDR(cidr)
	addr, _ := netip.ParseAddr(ip.String())
	if addr.Is4() {
		return IPv4
	}
	return IPv6
}

func CidrSize(cidr string) int64 {
	_, net, _ := net.ParseCIDR(cidr)
	maskSize, _ := net.Mask.Size()
	var prefixLength int
	if CidrVersion(cidr) == IPv4 {
		prefixLength = 32
	} else {
		prefixLength = 128
	}
	// when netmask < 64, ignore size as 0
	return 1 << (prefixLength - maskSize)
}

func TrimGatewayMask(gateway string) string {
	if gateway == "" {
		return ""
	}
	if strings.Contains(gateway, "/") {
		gateway = strings.Split(gateway, "/")[0]
	} else if strings.Contains(gateway, "::") {
		gateway = strings.Split(gateway, "::")[0]
	}
	return gateway
}
