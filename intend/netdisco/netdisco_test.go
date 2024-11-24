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

package netdisco_test

import (
	"testing"

	"github.com/gosnmp/gosnmp"
	mem_cache "github.com/wangxin688/narvis/intend/cache"
	"github.com/wangxin688/narvis/intend/logger"
	nettyx_snmp "github.com/wangxin688/narvis/intend/model/snmp"
	"github.com/wangxin688/narvis/intend/netdisco"
)

func TestDriver(t *testing.T) {
	loggerConfig := logger.LogConfig{
		Formatter: "text",
	}
	logger.InitLogger(&loggerConfig)
	mem_cache.InitCache()
	community := "public"
	target := nettyx_snmp.SnmpConfig{
		IpAddress:      "127.0.0.1",
		Version:        gosnmp.Version2c,
		Port:           161,
		Community:      &community,
		Timeout:        3,
		MaxRepetitions: 10,
	}
	disco, err := netdisco.NewNetDisco().Driver(&target)
	if err != nil {
		t.Fatalf("failed to create netdisco: %s", err)
	}
	response := disco.BasicInfo()

	t.Logf("response: %v", response)
}
