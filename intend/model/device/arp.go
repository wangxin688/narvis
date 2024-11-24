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

type ArpItem struct {
	IpAddress  string `json:"ipAddress"`
	MacAddress string `json:"macAddress"`
	Type       string `json:"type"` // one of other/invalid/dynamic/static
	IfIndex    uint64 `json:"ifIndex"`
	VlanId     uint32 `json:"vlanId"`
	Range      string `json:"range"`
	HashValue  string `json:"hashValue"`
}

func GetArpTypeValue(arpType uint64) string {
	IpNetToMediaTypeValueMapping := map[uint64]string{
		1: "other",
		2: "invalid",
		3: "dynamic",
		4: "static",
	}
	return IpNetToMediaTypeValueMapping[arpType]
}
