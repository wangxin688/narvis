package factory

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gosnmp/gosnmp"
	"github.com/samber/lo"
)

type SnmpDiscovery struct {
	Session   *gosnmp.GoSNMP
	IpAddress string
}

type BaseSnmpConfig struct {
	Port           uint16
	Version        gosnmp.SnmpVersion
	Timeout        uint8
	Community      *string
	V3Params       *SnmpV3Params
	MaxRepetitions int
}

type SnmpConfig struct {
	IpAddress string
	BaseSnmpConfig
}

func NewSnmpSession(sc SnmpConfig) (*gosnmp.GoSNMP, error) {
	var instance *gosnmp.GoSNMP
	if !sc.validate() {
		return nil, fmt.Errorf("invalid snmp config parameters for %s", sc.IpAddress)
	} else if sc.Version == gosnmp.Version2c {
		instance = &gosnmp.GoSNMP{
			Target:    sc.IpAddress,
			Port:      sc.Port,
			Timeout:   time.Duration(sc.Timeout) * time.Second,
			Retries:   2,
			MaxOids:   int(sc.MaxRepetitions),
			Version:   sc.Version,
			Community: *sc.Community,
		}
	} else {
		instance = &gosnmp.GoSNMP{
			Target:   sc.IpAddress,
			Port:     sc.Port,
			Timeout:  time.Duration(sc.Timeout) * time.Second,
			Retries:  2,
			MaxOids:  int(sc.MaxRepetitions),
			Version:  sc.Version,
			MsgFlags: gosnmp.AuthPriv,
			SecurityParameters: &gosnmp.UsmSecurityParameters{
				UserName:                 *sc.V3Params.SecurityName,
				AuthenticationProtocol:   sc.V3Params.AuthProtocol,
				AuthenticationPassphrase: *sc.V3Params.AuthPassword,
				PrivacyProtocol:          sc.V3Params.PrivProtocol,
				PrivacyPassphrase:        *sc.V3Params.PrivPassword,
			},
		}
	}
	err := instance.Connect()
	if err != nil {
		return nil, err
	}
	return instance, nil
}

func NewSnmpDiscovery(sc SnmpConfig) (*SnmpDiscovery, error) {
	session, err := NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &SnmpDiscovery{
		Session:   session,
		IpAddress: session.Target,
	}, nil
}

// validate checks if the snmp config is valid
func (c *SnmpConfig) validate() bool {
	switch c.Version {
	case gosnmp.Version2c:
		return c.Community != nil
	case gosnmp.Version3:
		return c.V3Params != nil
	default:
		return false
	}
}

// SysDescr returns the system description from RFC-1213MIB
func (sd *SnmpDiscovery) SysDescr() (string, error) {
	oid := []string{SysDescr}
	result, err := sd.Session.Get(oid)
	if err != nil {
		return "", fmt.Errorf("failed to get system description from %s: %w", sd.IpAddress, err)
	}

	for _, variable := range result.Variables {
		if variable.Type == gosnmp.OctetString {
			return fmt.Sprintf("%s", variable.Value), nil
		}
	}

	return "", fmt.Errorf("failed to get system description from %s: no variable found", sd.IpAddress)
}

// SysObjectID returns the sysObjectID from RFC-1213MIB
func (sd *SnmpDiscovery) SysObjectID() (string, error) {
	oid := []string{SysObjectID}
	result, err := sd.Session.Get(oid)
	if err != nil {
		return "", err
	}

	for _, variable := range result.Variables {
		if variable.Type == gosnmp.OctetString {
			return fmt.Sprintf("%s", variable.Value), nil
		}
	}
	return "", fmt.Errorf("failed to get sysObjectID from %s for OIDs: %v", sd.IpAddress, oid)
}

// SysUpTime returns the sysUpTime from RFC-1213MIB
func (sd *SnmpDiscovery) SysUpTime() (uint64, error) {
	ticksVariables, err := sd.Session.Get([]string{SysUpTime})
	if err != nil {
		return 0, err
	}

	var ticks uint64
	for _, variable := range ticksVariables.Variables {
		if variable.Type == gosnmp.TimeTicks {
			switch v := variable.Value.(type) {
			case uint32:
				ticks = uint64(v)
			case uint16:
				ticks = uint64(v)
			case uint64:
				ticks = v
			default:
				return 0, fmt.Errorf("unexpected type for sysUpTime: %T", variable.Value)
			}
			return ticks, nil
		}
	}

	return 0, fmt.Errorf("failed to get sysUpTime from %s", sd.IpAddress)
}

// SysName returns the sysName from RFC-1213MIB
func (sd *SnmpDiscovery) SysName() (string, error) {
	sysNameVars, err := sd.Session.Get([]string{SysName})
	if err != nil {
		return "", err
	}

	for _, variable := range sysNameVars.Variables {
		if variable.Type == gosnmp.OctetString {
			return fmt.Sprintf("%s", variable.Value), nil
		}
	}

	return "", fmt.Errorf("failed to get sysName from %s", sd.IpAddress)
}

// ChassisId returns the chassisId from RFC-1213MIB
func (sd *SnmpDiscovery) ChassisId() (string, error) {
	result, err := sd.Session.Get([]string{LldpLocChassisId})
	if err != nil {
		return "", err
	}
	for _, variable := range result.Variables {
		if variable.Type == gosnmp.OctetString {
			return hex2mac(variable.Value.([]byte)), nil
		}
	}
	return "", fmt.Errorf("failed to get ChassisId from %s", sd.IpAddress)
}

func (sd *SnmpDiscovery) IfPortMode() map[uint64]string {
	// TODO: implement me
	return nil
}

// collect interfaces via IF-MIB
func (sd *SnmpDiscovery) Interfaces() (interfaces []*DeviceInterface, errors []string) {
	ifIndex, errIfIndex := sd.Session.BulkWalkAll(IfIndex)
	if len(ifIndex) == 0 || errIfIndex != nil {
		return nil, []string{fmt.Sprintf("failed to get ifIndex from %s", sd.IpAddress)}
	}
	ifName, errIfName := sd.Session.BulkWalkAll(IfDescr)
	ifDesc, errIfDesc := sd.Session.BulkWalkAll(IfAlias)
	ifMtu, errIfMtu := sd.Session.BulkWalkAll(IfMtu)
	ifSpeed, errIfSpeed := sd.Session.BulkWalkAll(IfHighSpeed)
	ifHighSpeed, errIfHighSpeed := sd.Session.BulkWalkAll(IfHighSpeed)
	ifPhysAddr, errIfPhysAddr := sd.Session.BulkWalkAll(IfPhysAddr)
	ifType, errIfType := sd.Session.BulkWalkAll(IfType)
	ifAdminStatus, errIfAdminStatus := sd.Session.BulkWalkAll(IfAdminStatus)
	ifOperStatus, errIfOperStatus := sd.Session.BulkWalkAll(IfOperStatus)
	ifLastChange, errIfLastChange := sd.Session.BulkWalkAll(IfLastChange)
	ifAddrIndex, errIfAddrIndex := sd.Session.BulkWalkAll(IfAdEntIfIndex)
	ifAddrNetMask, errIfAddrNetMask := sd.Session.BulkWalkAll(IfAdEntNetMask)

	if errIfName != nil || errIfDesc != nil ||
		errIfMtu != nil || errIfSpeed != nil || errIfHighSpeed != nil ||
		errIfPhysAddr != nil || errIfType != nil || errIfAdminStatus != nil || errIfLastChange != nil ||
		errIfOperStatus != nil || errIfAddrIndex != nil || errIfAddrNetMask != nil {
		errors = append(errors, errIfName.Error())
		errors = append(errors, errIfDesc.Error())
		errors = append(errors, errIfMtu.Error())
		errors = append(errors, errIfSpeed.Error())
		errors = append(errors, errIfHighSpeed.Error())
		errors = append(errors, errIfPhysAddr.Error())
		errors = append(errors, errIfType.Error())
		errors = append(errors, errIfAdminStatus.Error())
		errors = append(errors, errIfOperStatus.Error())
		errors = append(errors, errIfLastChange.Error())
		errors = append(errors, errIfAddrIndex.Error())
		errors = append(errors, errIfAddrNetMask.Error())
	}
	indexIfIndex := extractInteger(IfIndex, ifIndex)
	indexIfName := extractString(IfDescr, ifName)
	indexIfDesc := extractString(IfAlias, ifDesc)
	indexIfMtu := extractInteger(IfMtu, ifMtu)
	indexIfSpeed := extractInteger(IfHighSpeed, ifSpeed)
	indexIfHighSpeed := extractInteger(IfHighSpeed, ifHighSpeed)
	indexIfPhysAddr := extractMacAddress(IfPhysAddr, ifPhysAddr)
	indexIfType := extractInteger(IfType, ifType)
	indexIfAdminStatus := extractInteger(IfAdminStatus, ifAdminStatus)
	indexIfOperStatus := extractInteger(IfOperStatus, ifOperStatus)
	indexIfLastChange := extractInteger(IfLastChange, ifLastChange)
	indexIfAddrIndex := extractIfIndex(IfAdEntIfIndex, ifAddrIndex)
	indexIfAddrNetMask := extractString(IfAdEntNetMask, ifAddrNetMask)
	for i, v := range indexIfIndex {
		var _ifAddrIndex string
		itemIfAddrIndex := indexIfAddrIndex[i]
		itemIfAddrNetMask := indexIfAddrNetMask[itemIfAddrIndex]
		if indexIfAddrIndex[i] == "" || indexIfAddrNetMask[i] == "" {
			_ifAddrIndex = ""
		} else {
			_ifAddrIndex = strings.TrimPrefix(indexIfAddrIndex[i], ".") + "/" + strconv.Itoa(netmaskToLength(itemIfAddrNetMask))
		}
		iface := DeviceInterface{
			IfIndex:       v,
			IfName:        indexIfName[i],
			IfDescr:       indexIfDesc[i],
			IfType:        indexIfType[i],
			IfMtu:         indexIfMtu[i],
			IfSpeed:       indexIfSpeed[i],
			IfPhysAddr:    indexIfPhysAddr[i],
			IfAdminStatus: indexIfAdminStatus[i],
			IfOperStatus:  indexIfOperStatus[i],
			IfLastChange:  indexIfLastChange[i],
			IfHighSpeed:   indexIfHighSpeed[i],
			IfIpAddress:   _ifAddrIndex,
		}
		interfaces = append(interfaces, &iface)
	}

	return interfaces, errors
}

func (sd *SnmpDiscovery) LldpNeighbors() (lldp []*LldpNeighbor, errors []string) {

	localChassisId, err := sd.ChassisId()
	if err != nil {
		errors = append(errors, err.Error())
		return nil, errors
	}
	hostname, _ := sd.SysName()
	localIfName, errIfName := sd.Session.BulkWalkAll(LldpLocPortId)
	localIfDescr, errIfDescr := sd.Session.BulkWalkAll(LldpLocPortDesc)
	remoteChassisId, errRemChassisId := sd.Session.BulkWalkAll(LldpRemChassisId)
	remoteHostname, errRemHostname := sd.Session.BulkWalkAll(LldpRemSysName)
	remoteIfName, errRemIfName := sd.Session.BulkWalkAll(LldpRemPortId)
	remoteIfDescr, errRemIfDescr := sd.Session.BulkWalkAll(LldpRemPortDesc)
	if errIfName != nil || errIfDescr != nil || errRemChassisId != nil || errRemIfName != nil || errRemIfDescr != nil {
		errors = append(errors, errIfName.Error())
		errors = append(errors, errIfDescr.Error())
		errors = append(errors, errRemHostname.Error())
		errors = append(errors, errRemChassisId.Error())
		errors = append(errors, errRemIfName.Error())
		errors = append(errors, errRemIfDescr.Error())
	}
	IndexIfName := extractString(LldpLocPortId, localIfName)
	IndexIfDescr := extractString(LldpLocPortDesc, localIfDescr)
	IndexRemChassisId := extractMacAddressWithShift(LldpRemChassisId, -2, remoteChassisId)
	IndexRemoteHostname := extractStringWithShift(LldpRemSysName, -2, remoteHostname)
	IndexRemoteIfName := extractStringWithShift(LldpRemPortId, -2, remoteIfName)
	IndexRemoteIfDescr := extractStringWithShift(LldpRemPortDesc, -2, remoteIfDescr)

	for i, v := range IndexRemChassisId {
		lldp = append(lldp, &LldpNeighbor{
			LocalChassisId:  localChassisId,
			LocalHostname:   hostname,
			LocalIfName:     IndexIfName[i],
			LocalIfDescr:    IndexIfDescr[i],
			RemoteChassisId: v,
			RemoteHostname:  IndexRemoteHostname[i],
			RemoteIfName:    IndexRemoteIfName[i],
			RemoteIfDescr:   IndexRemoteIfDescr[i],
		})
	}
	return lldp, errors
}

func (sd *SnmpDiscovery) Entities() (entities []*Entity, errors []string) {
	entPhysicalClass, err := sd.Session.BulkWalkAll(EntPhysicalClass)
	if len(entPhysicalClass) == 0 || err != nil {
		if err != nil {
			errors = append(errors, err.Error())
		} else {
			errors = append(errors, fmt.Sprintf("get entity physical class failed from %s, No entities found", sd.IpAddress))
		}
		return nil, errors
	}
	IndexEntPhysicalClass := extractInteger(EntPhysicalClass, entPhysicalClass)
	FilteredIndexEntPhysicalClass := lo.PickByValues(IndexEntPhysicalClass, []uint64{3})
	chassidIndex := lo.Keys(FilteredIndexEntPhysicalClass)
	if len(chassidIndex) == 0 {
		errors = append(errors, fmt.Sprintf("get entity physical class failed from %s, No entities found", sd.IpAddress))
		return nil, errors
	}
	entityPhysicalDescr, errEntityPhysicalDescr := sd.Session.Get(buildOidWithIndex(EntPhysicalDescr, chassidIndex))
	entityPhysicalName, errEntityPhysicalName := sd.Session.Get(buildOidWithIndex(EntPhysicalName, chassidIndex))
	entityPhysicalSoftwareRev, errEntityPhysicalSoftwareRev := sd.Session.Get(buildOidWithIndex(EntPhysicalSoftwareRev, chassidIndex))
	entityPhysicalSerialNum, errEntityPhysicalSerialNum := sd.Session.Get(buildOidWithIndex(EntPhysicalSerialNum, chassidIndex))

	if errEntityPhysicalDescr != nil || errEntityPhysicalName != nil || errEntityPhysicalSoftwareRev != nil || errEntityPhysicalSerialNum != nil {
		errors = append(errors, errEntityPhysicalDescr.Error())
		errors = append(errors, errEntityPhysicalName.Error())
		errors = append(errors, errEntityPhysicalSoftwareRev.Error())
		errors = append(errors, errEntityPhysicalSerialNum.Error())
	}
	IndexEntityPhysicalDescr := extractString(EntPhysicalDescr, entityPhysicalDescr.Variables)
	IndexEntityPhysicalName := extractString(EntPhysicalName, entityPhysicalName.Variables)
	IndexEntityPhysicalSoftwareRev := extractString(EntPhysicalSoftwareRev, entityPhysicalSoftwareRev.Variables)
	IndexEntityPhysicalSerialNum := extractString(EntPhysicalSerialNum, entityPhysicalSerialNum.Variables)
	for i, v := range FilteredIndexEntPhysicalClass {
		entities = append(entities, &Entity{
			EntityPhysicalClass:       v,
			EntityPhysicalDescr:       IndexEntityPhysicalDescr[i],
			EntityPhysicalName:        IndexEntityPhysicalName[i],
			EntityPhysicalSoftwareRev: IndexEntityPhysicalSoftwareRev[i],
			EntityPhysicalSerialNum:   IndexEntityPhysicalSerialNum[i],
		})
	}
	return entities, errors
}

func (sd *SnmpDiscovery) MacAddressTable() (macTable *map[uint64][]string, errors []string) {
	dot1dBasePortIndex, err := sd.Session.BulkWalkAll(Dot1dBasePortIfIndex)
	if err != nil {
		return nil, []string{err.Error()}
	}
	dot1dTpFdbAddress, errDot1dTpFdbAddress := sd.Session.BulkWalkAll(Dot1dTpFdbAddress)
	dot1dTpFdbPort, errDot1dTpFdbPort := sd.Session.BulkWalkAll(Dot1dTpFdbPort)
	if errDot1dTpFdbAddress != nil || errDot1dTpFdbPort != nil {
		errors = append(errors, errDot1dTpFdbAddress.Error())
		errors = append(errors, errDot1dTpFdbPort.Error())
	}

	_IndexDot1dBasePortIfIndex := extractInteger(Dot1dBasePortIfIndex, dot1dBasePortIndex)
	IndexDot1dBasePortIfIndex := lo.MapValues(_IndexDot1dBasePortIfIndex, func(x uint64, _ string) string {
		return strconv.FormatUint(x, 10)
	})
	IndexDot1dBasePortIfIndex["0"] = "0"
	indexDot1dTpFdbAddress := make(map[string]string)
	for _, v := range dot1dTpFdbAddress {
		splitValue := strings.Split(v.Name, ".")
		index := strings.Join(splitValue[len(splitValue)-6:], ".") // 最后六位时MAC地址地址的十进制表示
		if mac := hex2mac(v.Value.([]byte)); mac != "" {
			indexDot1dTpFdbAddress[index] = mac
		}
	}
	indexDot1dTpFdbPort := make(map[string]string)
	for _, v := range dot1dTpFdbPort {
		splitValue := strings.Split(v.Name, ".")
		index := strings.Join(splitValue[len(splitValue)-6:], ".")
		indexDot1dTpFdbPort[index] = strconv.FormatUint(gosnmp.ToBigInt(v.Value).Uint64(), 10)
	}

	result := make(map[string][]string)
	for p := range IndexDot1dBasePortIfIndex {
		result[p] = make([]string, 0)
	}

	for portIndex, macAddress := range indexDot1dTpFdbAddress {
		if _, ok := indexDot1dTpFdbPort[portIndex]; ok {
			result[IndexDot1dBasePortIfIndex["."+indexDot1dTpFdbPort[portIndex]]] = append(result[portIndex], macAddress)
		}
	}

	_macTable := lo.MapKeys(result, func(_ []string, x string) uint64 {
		index, _ := strconv.ParseUint(x, 10, 64)
		return index
	})

	return &_macTable, errors
}

func (sd *SnmpDiscovery) ArpTable() (arp *map[string]string, error error) {
	arpTable, err := sd.Session.BulkWalkAll(IpNetToMediaPhysAddress)
	if err != nil {
		return nil, err
	}
	arpMap := extractMacAddress(IpNetToMediaPhysAddress, arpTable)
	_arp := lo.MapKeys(arpMap, func(_ string, x string) string {
		splitData := strings.Split(x, ".")
		x_last_4 := splitData[len(splitData)-4:]
		return strings.Join(x_last_4, ".")
	})
	return &_arp, nil
}

func (sd *SnmpDiscovery) Discovery() *DiscoveryResponse {
	sysDescr, sysError := sd.SysDescr()
	sysObjectId, sysObjectIdError := sd.SysObjectID()
	sysUpTime, sysUpTimeError := sd.SysUpTime()
	sysName, sysNameError := sd.SysName()
	chassisId, chassisIdError := sd.ChassisId()
	interfaces, interfacesError := sd.Interfaces()
	entities, entitiesError := sd.Entities()
	lldp, lldpError := sd.LldpNeighbors()
	macAddress, macAddressError := sd.MacAddressTable()
	arp, arpError := sd.ArpTable()

	response := &DiscoveryResponse{
		SysDescr:        sysDescr,
		SysObjectID:     sysObjectId,
		Uptime:          sysUpTime,
		HostName:        sysName,
		ChassisId:       chassisId,
		Interfaces:      interfaces,
		LldpNeighbors:   lldp,
		Entities:        entities,
		MacAddressTable: macAddress,
		ArpTable:        arp,
	}
	if sysError != nil {
		response.Errors = append(response.Errors, sysError.Error())
	}
	if sysObjectIdError != nil {
		response.Errors = append(response.Errors, sysObjectIdError.Error())
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
		response.Errors = append(response.Errors, arpError.Error())
	}
	return response
}
