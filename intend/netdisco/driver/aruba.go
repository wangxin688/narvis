package driver

import (
	"fmt"
	"strings"

	mem_cache "github.com/wangxin688/narvis/intend/cache"
	"github.com/wangxin688/narvis/intend/logger"
	intend_device "github.com/wangxin688/narvis/intend/model/device"
	"github.com/wangxin688/narvis/intend/model/snmp"
	"github.com/wangxin688/narvis/intend/model/wlanstation"
	"github.com/wangxin688/narvis/intend/netdisco/factory"

	"go.uber.org/zap"
)

// const wlanApMacAddress string = ".1.3.6.1.4.1.14823.2.2.1.5.2.1.4.1.1" not implement by aruba, replace with snmpIndex

const wlanAPIpAddress string = ".1.3.6.1.4.1.14823.2.2.1.5.2.1.4.1.2"
const wlanAPName string = ".1.3.6.1.4.1.14823.2.2.1.5.2.1.4.1.3"
const wlanAPGroupName string = ".1.3.6.1.4.1.14823.2.2.1.5.2.1.4.1.4"
const wlanAPModel string = ".1.3.6.1.4.1.14823.2.2.1.5.2.1.4.1.5"
const wlanAPSerialNumber string = ".1.3.6.1.4.1.14823.2.2.1.5.2.1.4.1.6"
const wlanAPSwitchIpAddress string = ".1.3.6.1.4.1.14823.2.2.1.5.2.1.4.1.39"
const wlanAPSwVersion string = ".1.3.6.1.4.1.14823.2.2.1.5.2.1.4.1.34"
const wlanAPBssidAPMacAddress = ".1.3.6.1.4.1.14823.2.2.1.5.2.1.7.1.13"

// WLSX-SYSTEMEXT-MIB
const wlsxSysExtHostname string = ".1.3.6.1.4.1.14823.2.2.1.2.1.2.0"
const wlsxSysExtModelName string = ".1.3.6.1.4.1.14823.2.2.1.2.1.3.0"
const wlsxSysExtSwVersion string = ".1.3.6.1.4.1.14823.2.2.1.2.1.28.0"
const wlsxSysExtSerialNumber string = ".1.3.6.1.4.1.14823.2.2.1.2.1.29.0"

const nUserName = ".1.3.6.1.4.1.14823.2.2.1.4.1.2.1.3"
const nUserAssignedVlan = ".1.3.6.1.4.1.14823.2.2.1.4.1.2.1.17"
const nUserApBSSID = ".1.3.6.1.4.1.14823.2.2.1.4.1.2.1.11"
const wlanStaAccessPointESSID = ".1.3.6.1.4.1.14823.2.2.1.5.2.2.1.1.12"
const wlanStaRSSI = ".1.3.6.1.4.1.14823.2.2.1.5.2.2.1.1.14"
const wlanStaUpTime = ".1.3.6.1.4.1.14823.2.2.1.5.2.2.1.1.15"
const wlanStaChannel = ".1.3.6.1.4.1.14823.2.2.1.5.2.2.1.1.6"
const wlanStaTxBytes = ".1.3.6.1.4.1.14823.2.2.1.5.3.2.1.1.3"
const wlanStaRxBytes = ".1.3.6.1.4.1.14823.2.2.1.5.3.2.1.1.5"
const nUserDeviceType = ".1.3.6.1.4.1.14823.2.2.1.4.1.2.1.39"

type ArubaDriver struct {
	factory.SnmpDiscovery
}

type ArubaOSSwitchDriver struct {
	factory.SnmpDiscovery
}

func NewArubaDriver(sc *snmp.SnmpConfig) (*ArubaDriver, error) {
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

func NewArubaOSSwitchDriver(sc *snmp.SnmpConfig) (*ArubaOSSwitchDriver, error) {
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

func (ad *ArubaDriver) Entities() (entities []*intend_device.Entity, errors []string) {
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
	return []*intend_device.Entity{
		{
			EntityPhysicalClass:       "chassis",
			EntityPhysicalName:        fmt.Sprintf("%s", hostname.Variables[0].Value),
			EntityPhysicalSoftwareRev: fmt.Sprintf("%s", swVersion.Variables[0].Value),
			EntityPhysicalSerialNum:   fmt.Sprintf("%s", serialNumber.Variables[0].Value),
			EntityPhysicalDescr:       fmt.Sprintf("%s", modelName.Variables[0].Value),
		},
	}, nil

}

func (ad *ArubaDriver) APs() (ap []*intend_device.Ap, errors []string) {
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
	ad.cacheBssidApMac()

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
		apMac := factory.StringToHexMac(i)
		apName := indexApName[i]
		mem_cache.MemCache.SetDefault(apMac, apName)
		groupName := indexApGroupName[i]
		switchIp := indexSwitchIp[i]
		apVersion := indexApVersion[i]
		ap = append(ap, &intend_device.Ap{
			Name:            apName,
			ManagementIp:    v,
			MacAddress:      apMac,
			GroupName:       &groupName,
			DeviceModel:     indexApModel[i],
			SerialNumber:    indexApSerialNumber[i],
			WlanACIpAddress: &switchIp,
			OsVersion:       &apVersion,
		})
	}
	return ap, errors
}

func (ad *ArubaDriver) WlanUsers() (wlanUsers []*wlanstation.WlanUser, errors []string) {
	results := make([]*wlanstation.WlanUser, 0)
	userNames, err := ad.Session.BulkWalkAll(nUserName)
	if err != nil {
		return nil, []string{fmt.Sprintf("failed to get users from %s", ad.IpAddress), err.Error()}
	}
	userUptime, errUptime := ad.Session.BulkWalkAll(wlanStaUpTime)
	userAssignedVlan, errAssignedVlan := ad.Session.BulkWalkAll(nUserAssignedVlan)
	userRSSI, errRSSI := ad.Session.BulkWalkAll(wlanStaRSSI)
	userBSSID, errBSSID := ad.Session.BulkWalkAll(nUserApBSSID)
	userESSID, errESSID := ad.Session.BulkWalkAll(wlanStaAccessPointESSID)
	userChannel, errChannel := ad.Session.BulkWalkAll(wlanStaChannel)
	userTxBytes, errTxBytes := ad.Session.BulkWalkAll(wlanStaTxBytes)
	userRxBytes, errRxBytes := ad.Session.BulkWalkAll(wlanStaRxBytes)

	if errUptime != nil || errAssignedVlan != nil || errRSSI != nil ||
		errBSSID != nil || errESSID != nil || errChannel != nil || errTxBytes != nil || errRxBytes != nil {
		errors = append(errors, errUptime.Error())
		errors = append(errors, errAssignedVlan.Error())
		errors = append(errors, errBSSID.Error())
		errors = append(errors, errESSID.Error())
		errors = append(errors, errChannel.Error())
		errors = append(errors, errTxBytes.Error())
		errors = append(errors, errRxBytes.Error())
	}
	indexUserName := factory.ExtractString(nUserName, userNames)
	indexUserUptime := factory.ExtractInteger(wlanStaUpTime, userUptime)
	indexUserVlan := factory.ExtractInteger(nUserAssignedVlan, userAssignedVlan)
	indexBSSID := factory.ExtractMacAddress(nUserApBSSID, userBSSID)
	indexESSID := factory.ExtractString(wlanStaAccessPointESSID, userESSID)
	indexRSSI := factory.ExtractInteger(wlanStaRSSI, userRSSI)
	indexChannel := factory.ExtractInteger(wlanStaChannel, userChannel)
	indexTxBytes := factory.ExtractInteger(wlanStaTxBytes, userTxBytes)
	indexRxBytes := factory.ExtractInteger(wlanStaRxBytes, userRxBytes)
	for i, v := range indexUserName {
		mac, ip, macIndex := factory.SnmpIndexToMacAndIp(i)
		bssid := indexBSSID[i]
		vlan := uint16(indexUserVlan[i])
		channel := uint16(indexChannel[i])
		ap_name := ad.getApNameByBssid(bssid)
		user := wlanstation.WlanUser{
			StationMac:        mac,
			StationIp:         ip,
			StationUsername:   v,
			StationESSID:      indexESSID[macIndex],
			StationApName:     &ap_name,
			StationRSSI:       int8(indexRSSI[macIndex]) * -1,
			StationVlan:       &vlan,
			StationOnlineTime: indexUserUptime[macIndex],
			StationChannel:    channel,
			StationRxBits:     indexRxBytes[macIndex] * 8,
			StationTxBits:     indexTxBytes[macIndex] * 8,
			StationRadioType:  factory.ChannelToRadioType(channel),
		}
		results = append(results, &user)
	}
	return results, errors
}

func (ad *ArubaDriver) cacheBssidApMac() {
	bssidResult, err := ad.Session.BulkWalkAll(wlanAPBssidAPMacAddress)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("failed to get bssid ap mac from %s", ad.IpAddress), zap.Error(err))
	}
	indexBssid := factory.ExtractMacAddress(wlanAPBssidAPMacAddress, bssidResult)
	for i, v := range indexBssid {
		bssidOid := strings.Split(i, ".")
		bssid := factory.StringToHexMac("." + strings.Join(bssidOid[len(bssidOid)-6:], "."))
		mem_cache.MemCache.SetDefault(bssid, v)
	}
}

func (ad *ArubaDriver) getApNameByBssid(bssid string) (apName string) {
	apMac, ok := mem_cache.MemCache.Get(bssid)
	if !ok {
		return ""
	}
	apMacString := apMac.(string)
	apNameInterface, ok := mem_cache.MemCache.Get(apMacString)
	if !ok {
		return
	}
	return apNameInterface.(string)
}
