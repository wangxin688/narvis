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

package intend_device

type MacAddressItem struct {
	MacAddress string `json:"name"`
	IfIndex    uint64 `json:"ifIndex"`
	IfName     string `json:"ifName"`
	IfDescr    string `json:"ifDescr"`
	IpAddress  string `json:"ipAddress"`
	VlanId     uint32 `json:"vlanId"`
}
