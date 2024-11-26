package fixtures

import (
	"fmt"
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
