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

type ScanDevice struct {
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	ChassisId      string   `json:"chassisId"`
	ManagementIp   string   `json:"managementIp"`
	Manufacturer   string   `json:"manufacturer"`
	DeviceModel    string   `json:"deviceModel"`
	Platform       string   `json:"platform"`
	OrganizationId string   `json:"organizationId"`
	Errors         []string `json:"errors"`
}

type Device struct {
	DeviceId       string             `json:"deviceId"`
	SiteId         string             `json:"siteId"`
	Name           string             `json:"name"`
	Description    string             `json:"description"`
	ChassisId      *string            `json:"chassisId"`
	ManagementIp   string             `json:"managementIp"`
	Manufacturer   string             `json:"manufacturer"`
	DeviceModel    string             `json:"deviceModel"`
	Platform       string             `json:"platform"`
	OrganizationId string             `json:"organizationId"`
	Interfaces     []*DeviceInterface `json:"interfaces"`
	LldpNeighbors  []*LldpNeighbor    `json:"lldpNeighbors"`
	Entities       []*Entity          `json:"entities"`
	Stacks         []*Stack           `json:"stacks"`
}

type Entity struct {
	EntityPhysicalClass       string `json:"entityPhysicalClass"`
	EntityPhysicalDescr       string `json:"entityPhysicalDescr"`
	EntityPhysicalName        string `json:"entityPhysicalName"`
	EntityPhysicalSoftwareRev string `json:"entityPhysicalSoftwareRev"`
	EntityPhysicalSerialNum   string `json:"entityPhysicalSerialNum"`
}

type Stack struct {
	Id         string `json:"id"`
	Priority   uint32 `json:"priority"`
	Role       string `json:"role"`
	MacAddress string `json:"macAddress"`
}
