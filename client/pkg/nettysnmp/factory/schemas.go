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

type Vlan struct{}

type Route struct{}

type Prefix struct{}

type ArpItem struct {
	MacAddress string
	Type uint64
}

type DiscoveryResponse struct {
	Hostname        string               `json:"hostname"`
	SysDescr        string               `json:"sysDescr"`
	Uptime          uint64               `json:"uptime"`
	ChassisId       string               `json:"chassisID"`
	Interfaces      []*DeviceInterface   `json:"interfaces"`
	LldpNeighbors   []*LldpNeighbor      `json:"lldpNeighbors"`
	Entities        []*Entity            `json:"entities"`
	Stacks          []*Stack             `json:"stacks"`
	Vlans           []*Vlan              `json:"vlans"`
	MacAddressTable *map[uint64][]string `json:"macAddressTable"`
	ArpTable        *map[string]*ArpItem   `json:"arpTable"`
	Errors          []string             `json:"errors"`
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
