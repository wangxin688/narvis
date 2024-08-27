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
	IfIndex       uint64
	IfName        string
	IfDescr       string
	IfType        uint64
	IfMtu         uint64
	IfSpeed       uint64
	IfPhysAddr    string
	IfAdminStatus uint64
	IfOperStatus  uint64
	IfLastChange  uint64
	IfHighSpeed   uint64
	IfIpAddress   string
}

type LldpNeighbor struct {
	LocalChassisId  string
	LocalHostname   string
	LocalIfName     string
	LocalIfDescr    string
	RemoteChassisId string
	RemoteHostname  string
	RemoteIfName    string
	RemoteIfDescr   string
}

type Entity struct {
	EntityPhysicalClass       uint64
	EntityPhysicalDescr       string
	EntityPhysicalName        string
	EntityPhysicalSoftwareRev string
	EntityPhysicalSerialNum   string
}

type Stack struct {
	Id         string
	Priority   uint32
	Role       string
	MacAddress string
}

type Vlan struct{}

type Route struct{}

type Prefix struct{}

type DiscoveryResponse struct {
	HostName        string
	SysDescr        string
	Uptime          uint64
	SysObjectID     string
	ChassisId       string
	Interfaces      []*DeviceInterface
	LldpNeighbors   []*LldpNeighbor
	Entities        []*Entity
	Stack           []*Stack
	Vlans           []*Vlan
	MacAddressTable *map[uint64][]string
	ArpTable        *map[string]string
	Errors          []string
}

type DispatchResponse struct {
	IpAddress     string
	Data          *DiscoveryResponse
	SnmpReachable bool
	IcmpReachable bool
	SshReachable  bool
	SysObjectId   string
	DeviceModel   *devicemodel.DeviceModel
}
