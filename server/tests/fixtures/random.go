package fixtures

import (
	"fmt"
	"net"
	"time"

	"golang.org/x/exp/rand"
)

func RandomIpv4PrivateAddress(siteIndex, deviceIndex int) string {
	return fmt.Sprintf("10.%d.%d.%d", rand.Intn(256), siteIndex, deviceIndex)
}

func RandomIpv4() string {
	return fmt.Sprint(rand.Intn(256), ".", rand.Intn(256), ".", rand.Intn(256), ".", rand.Intn(256))
}

func RandomMacAddress() string {
	return fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

func RandomVlanId() int {
	return rand.Intn(4095)
}

func RandomIpv6Address() string {
	return fmt.Sprintf("%x:%x:%x:%x:%x:%x:%x:%x", rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256), rand.Intn(256))
}

func RandomString(length int) string {
	charset := "abcdefghijklmnopqrstuvwxyz" +
		"ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"0123456789"
	seededRand := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}

func GenerateRFC1918Prefix() (string, string) {
	rfc1918Ranges := []struct {
		base string
		mask int
	}{
		{"10.0.0.0", 8},
		{"172.16.0.0", 12},
		{"192.168.0.0", 16},
		{"100.64.0.0", 11},
		{"30.0.0.0", 8},
	}
	selectedRange := rfc1918Ranges[rand.Intn(len(rfc1918Ranges))]

	// 生成子网掩码，最小为 /21，最大为 /32
	minSubnetMask := 21
	maxSubnetMask := 32
	subnetMask := rand.Intn(maxSubnetMask-minSubnetMask+1) + minSubnetMask

	// 根据选择的范围和掩码生成合法网络号
	baseIP := net.ParseIP(selectedRange.base).To4()
	if baseIP == nil {
		panic("Invalid RFC1918 base IP")
	}

	// 计算网络地址偏移的范围
	subnetSize := 1 << (32 - subnetMask) // 当前子网的 IP 数量
	offsetRange := 1 << (subnetMask - selectedRange.mask)
	randomOffset := rand.Intn(offsetRange) * subnetSize

	// 计算网络号（确保对齐到子网边界）
	networkIP := make(net.IP, len(baseIP))
	copy(networkIP, baseIP)
	for i := 3; i >= 0; i-- {
		networkIP[i] += byte((randomOffset >> (8 * (3 - i))) & 0xFF)
	}
	// 返回合法的 CIDR
	return fmt.Sprintf("%s/%d", networkIP.String(), subnetMask), string(nextIP(baseIP, 1))
}

func nextIP(ip net.IP, inc uint) string {
	i := ip.To4()
	v := uint(i[0])<<24 + uint(i[1])<<16 + uint(i[2])<<8 + uint(i[3])
	v += inc
	v3 := byte(v & 0xFF)
	v2 := byte((v >> 8) & 0xFF)
	v1 := byte((v >> 16) & 0xFF)
	v0 := byte((v >> 24) & 0xFF)
	return net.IPv4(v0, v1, v2, v3).To4().String()
}

func GenerateRandomIPsByPrefix(prefix string, percentage float64) ([]string, error) {
	_, ipNet, err := net.ParseCIDR(prefix)
	if err != nil {
		return nil, fmt.Errorf("invalid prefix: %w", err)
	}

	// 计算总的 IP 数量
	maskSize, bits := ipNet.Mask.Size()
	totalIPs := 1 << (bits - maskSize)
	selectedCount := int(float64(totalIPs) * percentage)

	// 获取网络起始地址
	baseIP := ipNet.IP.To4()
	if baseIP == nil {
		return nil, fmt.Errorf("not an IPv4 CIDR: %s", prefix)
	}

	// 生成随机 IP
	rand.Seed(uint64(time.Now().UnixNano()))
	ipSet := make(map[string]struct{})
	for len(ipSet) < selectedCount {
		offset := rand.Intn(totalIPs)
		ip := make(net.IP, len(baseIP))
		copy(ip, baseIP)

		for i := 3; i >= 0; i-- {
			ip[i] += byte(offset & 0xFF)
			offset >>= 8
		}

		ipSet[ip.String()] = struct{}{}
	}

	// 转换为列表返回
	ips := make([]string, 0, len(ipSet))
	for ip := range ipSet {
		ips = append(ips, ip)
	}
	return ips, nil
}

func MockTimeBeforeNow() time.Time {
	now := time.Now()

	// 定义时间范围（8小时）
	maxDuration := 8 * time.Hour

	// 随机生成一个时间偏移（单位为秒）
	rand.Seed(uint64(time.Now().UnixNano()))
	randomOffset := time.Duration(rand.Intn(int(maxDuration.Seconds()))) * time.Second

	// 计算随机时间
	randomTime := now.Add(-randomOffset)
	return randomTime
}
