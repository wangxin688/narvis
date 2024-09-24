package driver

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"
)

const vtpVlanName = ".1.3.6.1.4.1.9.9.46.1.3.1.1.4"
const vtpVlanIfIndex = ".1.3.6.1.4.1.9.9.46.1.3.1.1.18"

// https://pastebin.com/PaP5yfYU
// https://mibs.observium.org/mib/AIRESPACE-WIRELESS-MIB/
const bsnAPDot3MacAddress = ".1.3.6.1.4.1.14179.2.2.1.1.1"

// const bsnAPEthernetMacAddress  = ".1.3.6.1.4.1.14179.2.2.1.1.33" // which one to use?
const bsnApIpAddress = ".1.3.6.1.4.1.14179.2.2.1.1.19"
const bsnAPName = ".1.3.6.1.4.1.14179.2.2.1.1.3"
const bsnAPType = ".1.3.6.1.4.1.14179.2.2.1.1.22"
const bsnAPSerialNumber = ".1.3.6.1.4.1.14179.2.2.1.1.17"
const bsnAPPrimaryMwarName = ".1.3.6.1.4.1.14179.2.2.1.1.10"

type CiscoBaseDriver struct {
	factory.SnmpDiscovery
}

type CiscoIosDriver struct {
	CiscoBaseDriver
}

type CiscoIosXRDriver struct {
	CiscoBaseDriver
}

type CiscoNexusOSDriver struct {
	CiscoBaseDriver
}

type CiscoIosXEDriver struct {
	CiscoBaseDriver
}

func NewCiscoIosDriver(sc factory.SnmpConfig) (*CiscoIosDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &CiscoIosDriver{
		CiscoBaseDriver{
			factory.SnmpDiscovery{
				Session:   session,
				IpAddress: session.Target},
		},
	}, nil
}

func NewCiscoIosXRDriver(sc factory.SnmpConfig) (*CiscoIosXRDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &CiscoIosXRDriver{
		CiscoBaseDriver{
			factory.SnmpDiscovery{
				Session:   session,
				IpAddress: session.Target},
		},
	}, nil
}

func NewCiscoNexusOSDriver(sc factory.SnmpConfig) (*CiscoNexusOSDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &CiscoNexusOSDriver{
		CiscoBaseDriver{
			factory.SnmpDiscovery{
				Session:   session,
				IpAddress: session.Target},
		},
	}, nil
}

func NewCiscoIosXEDriver(sc factory.SnmpConfig) (*CiscoIosXEDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &CiscoIosXEDriver{
		CiscoBaseDriver{
			factory.SnmpDiscovery{
				Session:   session,
				IpAddress: session.Target},
		},
	}, nil
}

func (cd *CiscoBaseDriver) Vlans() (vlan []*factory.VlanItem, errors []string) {
	l2Vlan, err := cd.Session.BulkWalkAll(vtpVlanName)
	l2VlanIfIndex, errIfIndex := cd.Session.BulkWalkAll(vtpVlanIfIndex)
	if err != nil || errIfIndex != nil {
		errors = append(errors, err.Error())
		errors = append(errors, errIfIndex.Error())
	}
	indexL2Vlan := factory.ExtractString(vtpVlanName, l2Vlan)
	indexVlanIndex := factory.ExtractInteger(vtpVlanIfIndex, l2VlanIfIndex)
	for i, v := range indexL2Vlan {
		vlanIdStrings := strings.Split(i, ".")
		vlanIdString := vlanIdStrings[len(vlanIdStrings)-1]
		vlanId, _ := strconv.Atoi(vlanIdString)
		_vlan := &factory.VlanItem{
			VlanId:   uint32(vlanId),
			VlanName: v,
			IfIndex:  indexVlanIndex[i],
		}
		vlan = append(vlan, _vlan)
	}

	return vlan, errors
}

func (cd *CiscoBaseDriver) APs() (ap []*factory.ApItem, errors []string) {
	apIP, errApIP := cd.Session.BulkWalkAll(bsnApIpAddress)
	if len(apIP) == 0 || errApIP != nil {
		return nil, []string{fmt.Sprintf("failed to get ap ipAddress from %s", cd.IpAddress)}
	}
	apMac, errApMac := cd.Session.BulkWalkAll(bsnAPDot3MacAddress)
	apName, errApName := cd.Session.BulkWalkAll(bsnAPName)
	apType, errApType := cd.Session.BulkWalkAll(bsnAPType)
	apSerialNumber, errApSerialNumber := cd.Session.BulkWalkAll(bsnAPSerialNumber)
	apPrimaryMwarName, errApPrimaryMwarName := cd.Session.BulkWalkAll(bsnAPPrimaryMwarName)
	if errApMac != nil || errApName != nil || errApType != nil || errApSerialNumber != nil || errApPrimaryMwarName != nil {
		errors = append(errors, errApMac.Error())
		errors = append(errors, errApName.Error())
		errors = append(errors, errApType.Error())
		errors = append(errors, errApSerialNumber.Error())
		errors = append(errors, errApPrimaryMwarName.Error())
	}
	indexApIP := factory.ExtractString(bsnApIpAddress, apIP)
	indexApMac := factory.ExtractString(bsnAPDot3MacAddress, apMac)
	indexApName := factory.ExtractString(bsnAPName, apName)
	indexApType := factory.ExtractString(bsnAPType, apType)
	indexApSerialNumber := factory.ExtractString(bsnAPSerialNumber, apSerialNumber)
	indexApPrimaryMwarName := factory.ExtractString(bsnAPPrimaryMwarName, apPrimaryMwarName)
	for i, v := range indexApIP {
		ap = append(ap, &factory.ApItem{
			Name:            indexApName[i],
			ManagementIp:    v,
			MacAddress:      indexApMac[i],
			DeviceModel:     indexApType[i],
			SerialNumber:    indexApSerialNumber[i],
			WlanACIpAddress: indexApPrimaryMwarName[i],
		})
	}
	return ap, errors

}

func (cd *CiscoBaseDriver) Discovery() *factory.DiscoveryResponse {
	sysDescr, sysError := cd.SysDescr()
	sysUpTime, sysUpTimeError := cd.SysUpTime()
	sysName, sysNameError := cd.SysName()
	chassisId, chassisIdError := cd.ChassisId()
	interfaces, interfacesError := cd.Interfaces()
	entities, entitiesError := cd.Entities()
	lldp, lldpError := cd.LldpNeighbors()
	macAddress, macAddressError := cd.MacAddressTable()
	arp, arpError := cd.ArpTable()
	arp = factory.EnrichArpInfo(arp, interfaces)
	vlan, VlanError := cd.Vlans()
	vlan = factory.EnrichVlanInfo(vlan, interfaces)
	macAddress_ := factory.EnrichMacAddress(macAddress, interfaces, lldp, arp)
	response := &factory.DiscoveryResponse{
		SysDescr:        sysDescr,
		Uptime:          sysUpTime,
		Hostname:        sysName,
		ChassisId:       chassisId,
		Interfaces:      interfaces,
		LldpNeighbors:   lldp,
		Entities:        entities,
		MacAddressTable: macAddress_,
		ArpTable:        arp,
		Vlans:           vlan,
	}
	if sysError != nil {
		response.Errors = append(response.Errors, sysError.Error())
	}
	if sysUpTimeError != nil {
		response.Errors = append(response.Errors, sysUpTimeError.Error())
	}
	if sysNameError != nil {
		response.Errors = append(response.Errors, sysNameError.Error())
	}
	if chassisIdError != nil {
		response.Errors = append(response.Errors, chassisIdError.Error())
	}
	if interfacesError != nil {
		response.Errors = append(response.Errors, interfacesError...)
	}
	if entitiesError != nil {
		response.Errors = append(response.Errors, entitiesError...)
	}
	if lldpError != nil {
		response.Errors = append(response.Errors, lldpError...)
	}
	if macAddressError != nil {
		response.Errors = append(response.Errors, macAddressError...)
	}
	if arpError != nil {
		response.Errors = append(response.Errors, arpError...)
	}
	if VlanError != nil {
		response.Errors = append(response.Errors, VlanError...)
	}
	return response
}
