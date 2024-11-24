package factory

import (
	"github.com/gosnmp/gosnmp"
	"github.com/wangxin688/narvis/client/pkg/nettysnmp/devicemodel"
)

type SnmpV3Params struct {
	ContextEngineId *string
	SecurityName    *string
	SecurityLevel   gosnmp.SnmpV3SecurityModel
	AuthProtocol    gosnmp.SnmpV3AuthProtocol
	AuthPassword    *string
	PrivProtocol    gosnmp.SnmpV3PrivProtocol
	PrivPassword    *string
}

type DeviceInterface struct {
	IfIndex       uint64 `json:"ifIndex"`
	IfName        string `json:"ifName"`
	IfDescr       string `json:"ifDescr"`
	IfType        string `json:"ifType"`
	IfMtu         uint64 `json:"ifMtu"`
	IfSpeed       uint64 `json:"ifSpeed"`
	IfPhysAddr    string `json:"ifPhysAddr"`
	IfAdminStatus string `json:"ifAdminStatus"`
	IfOperStatus  string `json:"ifOperStatus"`
	IfLastChange  uint64 `json:"ifLastChange"`
	IfHighSpeed   uint64 `json:"ifHighSpeed"`
	IfIpAddress   string `json:"ifIpAddress"`
}

type LldpNeighbor struct {
	LocalChassisId  string `json:"localChassisId"`
	LocalHostname   string `json:"localHostname"`
	LocalIfName     string `json:"localIfName"`
	LocalIfDescr    string `json:"localIfDescr"`
	RemoteChassisId string `json:"remoteChassisId"`
	RemoteHostname  string `json:"remoteHostname"`
	RemoteIfName    string `json:"remoteIfName"`
	RemoteIfDescr   string `json:"remoteIfDescr"`
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

type VlanItem struct {
	VlanId   uint32 `json:"vlanId"`
	VlanName string `json:"vlanName"`
	IfIndex  uint64 `json:"ifIndex"`
	Range    string `json:"range"`
	Gateway  string `json:"gateway"`
}

type VlanAssignItem struct {
	VlanType string `json:"vlanType"`
	VlanId   uint32 `json:"vlanId"`
	IfIndex  uint64 `json:"ifIndex"`
}

type Route struct{}

type Prefix struct{}

type ArpItem struct {
	IpAddress  string `json:"ipAddress"`
	MacAddress string `json:"macAddress"`
	Type       string `json:"type"`
	IfIndex    uint64 `json:"ifIndex"`
	VlanId     uint32 `json:"vlanId"`
	Range      string `json:"range"`
}
type MacAddressItem struct {
	MacAddress string `json:"name"`
	IfIndex    uint64 `json:"ifIndex"`
	IfName     string `json:"ifName"`
	IfDescr    string `json:"ifDescr"`
	IpAddress  string `json:"ipAddress"`
	VlanId     uint32 `json:"vlanId"`
}

type ApItem struct {
	Name            string `json:"name"`
	MacAddress      string `json:"macAddress"`
	SerialNumber    string `json:"serialNumber"`
	ManagementIp    string `json:"managementIp"`
	GroupName       string `json:"groupName"`
	DeviceModel     string `json:"deviceModel"`
	WlanACIpAddress string `json:"wlanACIpAddress"`
	OsVersion       string `json:"osVersion"`
}

type WlanUser struct {
	StationMac           string  `json:"stationMac"`                     // 终端MAC
	StationIp            string  `json:"stationIp"`                      // 终端IP
	StationUsername      string  `json:"stationUsername"`                // 终端用户名
	StationApMac         *string `json:"stationApMac,omitempty"`         // AP MAC
	StationApName        *string `json:"stationApName,omitempty"`        // AP 名称
	StationESSID         string  `json:"stationESSID"`                   // ESSID
	StationVlan          *uint64 `json:"stationVlan,omitempty"`          // VLAN
	StationChannel       uint64  `json:"stationChannel,omitempty"`       // 信道
	StationChanBandWidth *string `json:"stationChanBandWidth,omitempty"` // 信道带宽
	StationRadioType     string  `json:"stationRadioType"`               // radio类型
	StationSNR           *uint64 `json:"stationSNR,omitempty"`           // 终端NR
	StationRSSI          uint64  `json:"stationRSSI"`                    // 终端RSSI
	StationRxBits        uint64  `json:"stationRxBytes"`                 // 终端下行流量
	StationTxBits        uint64  `json:"stationTxBytes"`                 // 终端上行流量
	StationMaxSpeed      *uint64 `json:"stationMaxSpeed,omitempty"`      // 终端协商速率
	StationOnlineTime    uint64  `json:"stationOnlineTime"`              // 终端在线时间
}

type WlanUserResponse struct {
	IpAddress     string      `json:"ipAddress"`
	SnmpReachable bool        `json:"snmpReachable"`
	WlanUsers     []*WlanUser `json:"wlanUsers"`
	Errors        []string    `json:"errors"`
}

type DiscoveryResponse struct {
	Hostname        string             `json:"hostname"`
	SysDescr        string             `json:"sysDescr"`
	Uptime          uint64             `json:"uptime"`
	ChassisId       string             `json:"chassisID"`
	Interfaces      []*DeviceInterface `json:"interfaces"`
	LldpNeighbors   []*LldpNeighbor    `json:"lldpNeighbors"`
	Entities        []*Entity          `json:"entities"`
	Stacks          []*Stack           `json:"stacks"`
	Vlans           []*VlanItem        `json:"vlans"`
	MacAddressTable []*MacAddressItem  `json:"macAddressTable"`
	ArpTable        []*ArpItem         `json:"arpTable"`
	Errors          []string           `json:"errors"`
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
	Data          []*ApItem                `json:"data"`
	SnmpReachable bool                     `json:"snmpReachable"`
	IcmpReachable bool                     `json:"icmpReachable"`
	SshReachable  bool                     `json:"sshReachable"`
	SysObjectId   string                   `json:"sysObjectId"`
	DeviceModel   *devicemodel.DeviceModel `json:"deviceModel"`
	Errors        []string                 `json:"errors"`
}
