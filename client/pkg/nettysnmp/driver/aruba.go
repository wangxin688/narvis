package driver

import (
	"fmt"

	"github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"
)

// const wlanApMacAddress string = ".1.3.6.1.4.1.14823.2.2.1.5.2.1.4.1.1" not implement by aruba, replace with snmpIndex

const wlanAPIpAddress string = ".1.3.6.1.4.1.14823.2.2.1.5.2.1.4.1.2"
const wlanAPName string = ".1.3.6.1.4.1.14823.2.2.1.5.2.1.4.1.3"
const wlanAPGroupName string = ".1.3.6.1.4.1.14823.2.2.1.5.2.1.4.1.4"
const wlanAPModel string = ".1.3.6.1.4.1.14823.2.2.1.5.2.1.4.1.5"
const wlanAPSerialNumber string = ".1.3.6.1.4.1.14823.2.2.1.5.2.1.4.1.6"
const wlanAPSwitchIpAddress string = ".1.3.6.1.4.1.14823.2.2.1.5.2.1.4.1.39"
const wlanAPSwVersion string = ".1.3.6.1.4.1.14823.2.2.1.5.2.1.4.1.34"

// WLSX-SYSTEMEXT-MIB
const wlsxSysExtHostname string = ".1.3.6.1.4.1.14823.2.2.1.2.1.2.0"
const wlsxSysExtModelName string = ".1.3.6.1.4.1.14823.2.2.1.2.1.3.0"
const wlsxSysExtSwVersion string = ".1.3.6.1.4.1.14823.2.2.1.2.1.28.0"
const wlsxSysExtSerialNumber string = ".1.3.6.1.4.1.14823.2.2.1.2.1.29.0"

type ArubaDriver struct {
	factory.SnmpDiscovery
}

type ArubaOSSwitchDriver struct {
	factory.SnmpDiscovery
}

func NewArubaDriver(sc factory.SnmpConfig) (*ArubaDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &ArubaDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}

func NewArubaOSSwitchDriver(sc factory.SnmpConfig) (*ArubaOSSwitchDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &ArubaOSSwitchDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}

func (ad *ArubaDriver) Entities() (entities []*factory.Entity, errors []string) {
	hostname, errHostname := ad.Session.Get([]string{wlsxSysExtHostname})
	if errHostname != nil {
		errors = append(errors, errHostname.Error())
	}
	modelName, errModelName := ad.Session.Get([]string{wlsxSysExtModelName})
	swVersion, errSwVersion := ad.Session.Get([]string{wlsxSysExtSwVersion})
	serialNumber, errSerialNumber := ad.Session.Get([]string{wlsxSysExtSerialNumber})
	if errModelName != nil || errSwVersion != nil || errSerialNumber != nil {
		errors = append(errors, errModelName.Error())
		errors = append(errors, errSwVersion.Error())
		errors = append(errors, errSerialNumber.Error())
		return nil, errors
	}
	return []*factory.Entity{
		{
			EntityPhysicalClass:       "chassis",
			EntityPhysicalName:        fmt.Sprintf("%s", hostname.Variables[0].Value),
			EntityPhysicalSoftwareRev: fmt.Sprintf("%s", swVersion.Variables[0].Value),
			EntityPhysicalSerialNum:   fmt.Sprintf("%s", serialNumber.Variables[0].Value),
			EntityPhysicalDescr:       fmt.Sprintf("%s", modelName.Variables[0].Value),
		},
	}, nil

}

func (ad *ArubaDriver) APs() (ap []*factory.ApItem, errors []string) {
	apIp, errApIP := ad.Session.BulkWalkAll(wlanAPIpAddress)
	if len(apIp) == 0 || errApIP != nil {
		return nil, []string{fmt.Sprintf("failed to get ap ipAddress from %s", ad.IpAddress)}
	}
	apName, errApName := ad.Session.BulkWalkAll(wlanAPName)
	apGroupName, errApGroupName := ad.Session.BulkWalkAll(wlanAPGroupName)
	apModel, errApModel := ad.Session.BulkWalkAll(wlanAPModel)
	apSerialNumber, errApSerialNumber := ad.Session.BulkWalkAll(wlanAPSerialNumber)
	switchIp, errSwitchIp := ad.Session.BulkWalkAll(wlanAPSwitchIpAddress)
	apVersion, errApVersion := ad.Session.BulkWalkAll(wlanAPSwVersion)

	if errApName != nil || errApGroupName != nil || errApModel != nil || errApSerialNumber != nil || errSwitchIp != nil || errApVersion != nil {
		errors = append(errors, errApName.Error())
		errors = append(errors, errApGroupName.Error())
		errors = append(errors, errApModel.Error())
		errors = append(errors, errApSerialNumber.Error())
		errors = append(errors, errSwitchIp.Error())
		errors = append(errors, errApVersion.Error())
	}
	indexApIp := factory.ExtractString(wlanAPIpAddress, apIp)
	indexApName := factory.ExtractString(wlanAPName, apName)
	indexApGroupName := factory.ExtractString(wlanAPGroupName, apGroupName)
	indexApModel := factory.ExtractString(wlanAPModel, apModel)
	indexApSerialNumber := factory.ExtractString(wlanAPSerialNumber, apSerialNumber)
	indexSwitchIp := factory.ExtractString(wlanAPSwitchIpAddress, switchIp)
	indexApVersion := factory.ExtractString(wlanAPSwVersion, apVersion)
	for i, v := range indexApIp {
		ap = append(ap, &factory.ApItem{
			Name:            indexApName[i],
			ManagementIp:    v,
			MacAddress:      factory.StringToHexMac(i),
			GroupName:       indexApGroupName[i],
			DeviceModel:     indexApModel[i],
			SerialNumber:    indexApSerialNumber[i],
			WlanACIpAddress: indexSwitchIp[i],
			OsVersion:       indexApVersion[i],
		})
	}
	return ap, errors
}

func (ad *ArubaDriver) Discovery() *factory.DiscoveryResponse {
	sysDescr, sysError := ad.SysDescr()
	sysUpTime, sysUpTimeError := ad.SysUpTime()
	sysName, sysNameError := ad.SysName()
	chassisId, chassisIdError := ad.ChassisId()
	interfaces, interfacesError := ad.Interfaces()
	entities, entitiesError := ad.Entities()
	lldp, lldpError := ad.LldpNeighbors()
	macAddress, macAddressError := ad.MacAddressTable()
	arp, arpError := ad.ArpTable()
	arp = factory.EnrichArpInfo(arp, interfaces)
	vlan, VlanError := ad.Vlans()
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
