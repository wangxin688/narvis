package factory

import (
	intend_device "github.com/wangxin688/narvis/intend/model/device"
	"github.com/wangxin688/narvis/intend/model/wlanstation"
	"github.com/wangxin688/narvis/intend/netdisco/devicemodel"
)

type VlanAssignItem struct {
	VlanType string `json:"vlanType"`
	VlanId   uint32 `json:"vlanId"`
	IfIndex  uint64 `json:"ifIndex"`
}

type Route struct{}

type Prefix struct{}

type WlanUserResponse struct {
	IpAddress     string                  `json:"ipAddress"`
	SnmpReachable bool                    `json:"snmpReachable"`
	WlanUsers     []*wlanstation.WlanUser `json:"wlanUsers"`
	Errors        []string                `json:"errors"`
}

type DiscoveryResponse struct {
	Hostname        string                           `json:"hostname"`
	SysDescr        string                           `json:"sysDescr"`
	Uptime          uint64                           `json:"uptime"`
	ChassisId       string                           `json:"chassisID"`
	Interfaces      []*intend_device.DeviceInterface `json:"interfaces"`
	LldpNeighbors   []*intend_device.LldpNeighbor    `json:"lldpNeighbors"`
	Entities        []*intend_device.Entity          `json:"entities"`
	Stacks          []*intend_device.Stack           `json:"stacks"`
	Vlans           []*intend_device.VlanItem        `json:"vlans"`
	MacAddressTable []*intend_device.MacAddressItem  `json:"macAddressTable"`
	ArpTable        []*intend_device.ArpItem         `json:"arpTable"`
	Errors          []string                         `json:"errors"`
}

type DiscoveryBasicResponse struct {
	Hostname  string   `json:"hostname"`
	SysDescr  string   `json:"sysDescr"`
	ChassisId string   `json:"chassisID"`
	Errors    []string `json:"errors"`
}

type DispatchResponse struct {
	IpAddress     string                   `json:"ipAddress"`
	Data          *DiscoveryResponse       `json:"data"`
	SnmpReachable bool                     `json:"snmpReachable"`
	IcmpReachable bool                     `json:"icmpReachable"`
	SshReachable  bool                     `json:"sshReachable"`
	SysObjectId   string                   `json:"sysObjectId"`
	DeviceModel   *devicemodel.DeviceModel `json:"deviceModel"`
}

type DispatchBasicResponse struct {
	IpAddress     string                   `json:"ipAddress"`
	Data          *DiscoveryBasicResponse  `json:"data"`
	SnmpReachable bool                     `json:"snmpReachable"`
	IcmpReachable bool                     `json:"icmpReachable"`
	SshReachable  bool                     `json:"sshReachable"`
	SysObjectId   string                   `json:"sysObjectId"`
	DeviceModel   *devicemodel.DeviceModel `json:"deviceModel"`
}

type DispatchApScanResponse struct {
	IpAddress     string                   `json:"ipAddress"`
	Data          []*intend_device.Ap      `json:"data"`
	SnmpReachable bool                     `json:"snmpReachable"`
	IcmpReachable bool                     `json:"icmpReachable"`
	SshReachable  bool                     `json:"sshReachable"`
	SysObjectId   string                   `json:"sysObjectId"`
	DeviceModel   *devicemodel.DeviceModel `json:"deviceModel"`
	Errors        []string                 `json:"errors"`
}

type VlanIpRange struct {
	VlanId  uint32
	Range   string
	Gateway string
}
