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

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	nettyx_processor "github.com/wangxin688/narvis/intend/helpers/processor"
)

type Ap struct {
	Name            string  `json:"name"`
	MacAddress      string  `json:"macAddress,omitempty"`
	SerialNumber    string  `json:"serialNumber,omitempty"`
	ManagementIp    string  `json:"managementIp"`
	GroupName       *string `json:"groupName,omitempty"`
	DeviceModel     string  `json:"deviceModel"`
	Manufacturer    string  `json:"manufacturer"`
	WlanACIpAddress *string `json:"wlanACIpAddress,omitempty"`
	OsVersion       *string `json:"osVersion,omitempty"`
	SiteId          string  `json:"siteId"`
	OrganizationId  string  `json:"organizationId"`
}

func (a *Ap) CalApHash() string {
	// assume organizationId should never be changed
	hashString := fmt.Sprintf(
		"%s-%s-%s-%s-%s-%s-%s-%s",
		a.Name,
		a.ManagementIp,
		a.SiteId,
		a.MacAddress,
		nettyx_processor.PtrStringToString(a.GroupName),
		nettyx_processor.PtrStringToString(a.WlanACIpAddress),
		a.SerialNumber,
		nettyx_processor.PtrStringToString(a.OsVersion),
	)
	hash := md5.New()
	_, _ = hash.Write([]byte(hashString))
	return hex.EncodeToString(hash.Sum(nil))
}
