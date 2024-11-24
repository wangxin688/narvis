package fixtures

import (
	"fmt"

	"golang.org/x/exp/rand"
)

func RandomIpv4PrivateAddress() string {
	return fmt.Sprintf("10.%d.%d.%d", rand.Intn(256), rand.Intn(256), rand.Intn(256))
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