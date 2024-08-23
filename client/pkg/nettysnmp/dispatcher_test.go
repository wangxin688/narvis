package nettygo

import (
	"fmt"
	"testing"

	"github.com/gosnmp/gosnmp"
	"github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"
)

func TestDispatcher(t *testing.T) {
	community := "public"
	dispatcher := NewDispatcher([]string{"127.0.0.1"}, factory.BaseSnmpConfig{
		Port:           161,
		Version:        gosnmp.Version2c,
		Timeout:        2,
		MaxRepetitions: 50,
		Community:      &community,
	},
	)
	response := dispatcher.Dispatch()
	fmt.Printf("%v\n", response)
}
