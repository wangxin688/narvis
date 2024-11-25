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

package netdisco

import (
	"sync"

	"github.com/wangxin688/narvis/intend/model/snmp"
	"github.com/wangxin688/narvis/intend/netdisco/factory"
)

var MaxScanDeviceChannelPoolSize = 200

func NetworkDeviceDiscovery(targets []*snmp.SnmpConfig) []*factory.DispatchBasicResponse {
	var response []*factory.DispatchBasicResponse

	var wg sync.WaitGroup
	ch := make(chan struct{}, MaxScanDeviceChannelPoolSize)
	for _, target := range targets {
		ch <- struct{}{}
		wg.Add(1)
		go func(target *snmp.SnmpConfig) {
			defer wg.Done()
			targetResponse := discoverBasic(target)

			response = append(response, targetResponse)
			<-ch
		}(target)
	}
	wg.Wait()
	return response

}

func discoverBasic(config *snmp.SnmpConfig) *factory.DispatchBasicResponse {
	var response = &factory.DispatchBasicResponse{}
	response.IpAddress = config.IpAddress
	driver, deviceModel, err := NewNetDisco(config).Driver()
	if err != nil || driver == nil {
		response.SnmpReachable = false
	} else {
		response.SnmpReachable = true
	}
	icmp := driver.IcmpReachable()
	ssh := driver.SshReachable()
	response.IcmpReachable = icmp
	response.SshReachable = ssh
	if !response.SnmpReachable {
		return response
	}
	basic := driver.BasicInfo()
	response.Data = basic
	response.DeviceModel = deviceModel
	return response
}
