package factory

type SnmpDriver interface {
	SysDescr() (string, error)
	SysObjectID() (string, error)
	SysUpTime() (uint64, error)
	SysName() (string, error)
	ChassisId() (string, error)
	IfPortMode() map[uint64]string
	Interfaces() (interfaces []*DeviceInterface, errors []string)
	LldpNeighbors() (lldp []*LldpNeighbor, errors []string)
	Entities() (entities []*Entity, errors []string)
	MacAddressTable() (macTable []*ArpItem, errors []string)
	ArpTable() (arp []*ArpItem, errors []string)
	Vlans() (vlans []*VlanItem, errors []string)
	Discovery() *DiscoveryResponse
}
