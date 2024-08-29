package nettygo

import (
	"encoding/json"
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
	result, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("failed to marshal response: %s", err)
	}
	fmt.Printf("%s\n", result)
}
