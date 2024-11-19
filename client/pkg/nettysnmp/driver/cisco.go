package driver

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"
	mem_cache "github.com/wangxin688/narvis/client/utils/cache"
)

const vtpVlanName = ".1.3.6.1.4.1.9.9.46.1.3.1.1.4"
const vtpVlanIfIndex = ".1.3.6.1.4.1.9.9.46.1.3.1.1.18"

// https://pastebin.com/PaP5yfYU
// https://mibs.observium.org/mib/AIRESPACE-WIRELESS-MIB/
// const bsnAPDot3MacAddress = ".1.3.6.1.4.1.14179.2.2.1.1.1" // Radio mac address for clients associated with the AP

const bsnAPEthernetMacAddress = ".1.3.6.1.4.1.14179.2.2.1.1.33" // Ethernet mac address connected to switch
const bsnApIpAddress = ".1.3.6.1.4.1.14179.2.2.1.1.19"
const bsnAPName = ".1.3.6.1.4.1.14179.2.2.1.1.3"
const bsnAPModel = ".1.3.6.1.4.1.14179.2.2.1.1.16"
const bsnAPSerialNumber = ".1.3.6.1.4.1.14179.2.2.1.1.17"

// const bsnAPPrimaryMwarName = ".1.3.6.1.4.1.14179.2.2.1.1.10"
const bsnAPSoftwareVersion = ".1.3.6.1.4.1.14179.2.2.1.1.8"

// CISCO-VLAN-MEMBERSHIP-MIB && CISCO-VTP-MIB
const vmVlanType = ".1.3.6.1.4.1.9.9.68.1.2.2.1.1"
const vmVlan = ".1.3.6.1.4.1.9.9.68.1.2.2.1.2"
const vlanTrunkPortIfIndex = ".1.3.6.1.4.1.9.9.46.1.6.1.1.1"
const vlanTrunkPortDynamicStatus = ".1.3.6.1.4.1.9.9.46.1.6.1.1.14"

// some oid has bug if version lower than ios-xe 17.9.5
const bsnMobileStationSsid = ".1.3.6.1.4.1.14179.2.1.4.1.7"
const bsnMobileStationAPMacAddr = ".1.3.6.1.4.1.14179.2.1.4.1.4"
const bsnMobileStationIpAddress = ".1.3.6.1.4.1.14179.2.1.4.1.2"
const bsnMobileStationUserName = ".1.3.6.1.4.1.14179.2.1.4.1.3"
const bsnMobileStationRSSI = ".1.3.6.1.4.1.14179.2.1.6.1.1"
const bsnMobileStationBytesReceived = ".1.3.6.1.4.1.14179.2.1.6.1.2"
const bsnMobileStationSnr = ".1.3.6.1.4.1.14179.2.1.6.1.26"
const bsnMobileStationBytesSent = ".1.3.6.1.4.1.14179.2.1.6.1.3"
const cldcClientUpTime = ".1.3.6.1.4.1.9.9.599.1.3.1.1.15"

const cldcClientDeviceType = ".1.3.6.1.4.1.9.9.599.1.3.1.1.44"
const bsnMobileStationVlanId = ".1.3.6.1.4.1.14179.2.1.4.1.29"
const cldcClientChannel = ".1.3.6.1.4.1.9.9.599.1.3.1.1.35"

// const bsnMobileStationProtocol = ".1.3.6.1.4.1.14179.2.1.4.1.25"

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
	if err != nil {
		return nil, []string{fmt.Sprintf("failed to get vlan from %s", cd.IpAddress)}
	}
	l2VlanIfIndex, errIfIndex := cd.Session.BulkWalkAll(vtpVlanIfIndex)
	if errIfIndex != nil {
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

func (cd *CiscoBaseDriver) VlanAssign() (vlan []*factory.VlanAssignItem, errors []string) {
	vlanType, errVlanType := cd.Session.BulkWalkAll(vmVlanType)
	if errVlanType != nil {
		return nil, []string{fmt.Sprintf("failed to get vlan assignment from %s", cd.IpAddress)}
	}

	vlanId, errVlanId := cd.Session.BulkWalkAll(vmVlan)
	if errVlanId != nil {
		return nil, []string{fmt.Sprintf("failed to get vlan assignment from %s", cd.IpAddress)}
	}
	indexVlanType := factory.ExtractInteger(vmVlanType, vlanType)
	indexVlanId := factory.ExtractInteger(vmVlan, vlanId)
	for i, v := range indexVlanType {
		ifIndexString := strings.TrimPrefix(i, ".")
		ifIndex, _ := strconv.Atoi(ifIndexString)
		_vlan := &factory.VlanAssignItem{
			VlanType: factory.GetCiscoVlanMemberShipTypeValue(v),
			IfIndex:  uint64(ifIndex),
			VlanId:   uint32(indexVlanId[i]),
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
	apMac, errApMac := cd.Session.BulkWalkAll(bsnAPEthernetMacAddress)
	apName, errApName := cd.Session.BulkWalkAll(bsnAPName)
	apType, errApType := cd.Session.BulkWalkAll(bsnAPModel)
	apSerialNumber, errApSerialNumber := cd.Session.BulkWalkAll(bsnAPSerialNumber)
	// apPrimaryMwarName, errApPrimaryMwarName := cd.Session.BulkWalkAll(bsnAPPrimaryMwarName)
	apVersion, errApVersion := cd.Session.BulkWalkAll(bsnAPSoftwareVersion)
	if errApMac != nil || errApName != nil || errApType != nil || errApSerialNumber != nil || errApVersion != nil {
		errors = append(errors, errApMac.Error())
		errors = append(errors, errApName.Error())
		errors = append(errors, errApType.Error())
		errors = append(errors, errApSerialNumber.Error())
		errors = append(errors, errApVersion.Error())
	}
	indexApIP := factory.ExtractString(bsnApIpAddress, apIP)
	indexApMac := factory.ExtractMacAddress(bsnAPEthernetMacAddress, apMac)
	indexApName := factory.ExtractString(bsnAPName, apName)
	indexApType := factory.ExtractString(bsnAPModel, apType)
	indexApSerialNumber := factory.ExtractString(bsnAPSerialNumber, apSerialNumber)
	// indexApPrimaryMwarName := factory.ExtractString(bsnAPPrimaryMwarName, apPrimaryMwarName)
	indexApVersion := factory.ExtractString(bsnAPSoftwareVersion, apVersion)
	for i, v := range indexApIP {
		apMac := indexApMac[i]
		apName := indexApName[i]
		mem_cache.MemCache.SetDefault(apMac, apName)
		ap = append(ap, &factory.ApItem{
			Name:            indexApName[i],
			ManagementIp:    v,
			MacAddress:      indexApMac[i],
			DeviceModel:     indexApType[i],
			SerialNumber:    indexApSerialNumber[i],
			WlanACIpAddress: cd.IpAddress,
			OsVersion:       indexApVersion[i],
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

func (cd *CiscoBaseDriver) WlanUsers() *factory.WlanUserResponse {
	results := make([]*factory.WlanUser, 0)
	errors := make([]string, 0)
	userNames, err := cd.Session.BulkWalkAll(bsnMobileStationUserName)
	if err != nil {
		return &factory.WlanUserResponse{
			Errors:    []string{fmt.Sprintf("failed to get user name from %s", cd.IpAddress)},
			WlanUsers: results,
		}
	}
	userUptime, errUptime := cd.Session.BulkWalkAll(cldcClientUpTime)
	userIp, errUserIp := cd.Session.BulkWalkAll(bsnMobileStationIpAddress)
	userAssignedVlan, errAssignedVlan := cd.Session.BulkWalkAll(bsnMobileStationVlanId)
	userRSSI, errRSSI := cd.Session.BulkWalkAll(bsnMobileStationRSSI)
	apMac, errApMac := cd.Session.BulkWalkAll(bsnMobileStationAPMacAddr)
	userESSID, errESSID := cd.Session.BulkWalkAll(bsnMobileStationSsid)
	userChannel, errChannel := cd.Session.BulkWalkAll(cldcClientChannel)
	userTxBytes, errTxBytes := cd.Session.BulkWalkAll(bsnMobileStationBytesSent)
	userRxBytes, errRxBytes := cd.Session.BulkWalkAll(bsnMobileStationBytesReceived)
	userSNR, errSnr := cd.Session.BulkWalkAll(bsnMobileStationSnr)

	if errUptime != nil || errRSSI != nil ||
		errUserIp != nil || errAssignedVlan != nil || errSnr != nil ||
		errApMac != nil || errESSID != nil || errChannel != nil || errTxBytes != nil || errRxBytes != nil {
		errors = append(errors, errUptime.Error())
		errors = append(errors, errAssignedVlan.Error())
		errors = append(errors, errRSSI.Error())
		errors = append(errors, errApMac.Error())
		errors = append(errors, errESSID.Error())
		errors = append(errors, errChannel.Error())
		errors = append(errors, errTxBytes.Error())
		errors = append(errors, errRxBytes.Error())
	}
	indexUserName := factory.ExtractString(bsnMobileStationUserName, userNames)
	indexUserIp := factory.ExtractString(bsnMobileStationIpAddress, userIp)
	indexUserVlan := factory.ExtractInteger(bsnMobileStationVlanId, userAssignedVlan)
	indexRssi := factory.ExtractInteger(bsnMobileStationRSSI, userRSSI)
	indexApMAc := factory.ExtractMacAddress(bsnMobileStationAPMacAddr, apMac)
	indexESSID := factory.ExtractString(bsnMobileStationSsid, userESSID)
	indexChannel := factory.ExtractInteger(cldcClientChannel, userChannel)
	indexTxBytes := factory.ExtractInteger(bsnMobileStationBytesSent, userTxBytes)
	indexRxBytes := factory.ExtractInteger(bsnMobileStationBytesReceived, userRxBytes)
	indexSnr := factory.ExtractInteger(bsnMobileStationSnr, userSNR)
	indexUptime := factory.ExtractInteger(cldcClientUpTime, userUptime)
	for i, v := range indexUserName {
		vlan := indexUserVlan[i]
		channel := indexChannel[i]
		ap_mac := indexApMAc[i]
		ap_name := cd.getAPName(ap_mac)
		snr := indexSnr[i]
		results = append(results, &factory.WlanUser{
			StationMac:        factory.StringToHexMac(i),
			StationApMac:      &ap_mac,
			StationApName:     &ap_name,
			StationIp:         indexUserIp[i],
			StationUsername:   v,
			StationESSID:      indexESSID[i],
			StationRSSI:       indexRssi[i],
			StationSNR:        &snr,
			StationVlan:       &vlan,
			StationOnlineTime: indexUptime[i],
			StationChannel:    channel,
			StationRxBytes:    indexRxBytes[i],
			StationTxBytes:    indexTxBytes[i],
			StationRadioType:  factory.ChannelToRadioType(channel),
		})
	}

	return &factory.WlanUserResponse{
		WlanUsers: results,
		Errors:    errors,
	}
}

func (cd *CiscoBaseDriver) getAPName(apMac string) string {
	apName, ok := mem_cache.MemCache.Get(apMac)
	if !ok {
		return ""
	}
	return apName.(string)
}
