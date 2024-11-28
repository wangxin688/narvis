package driver

import (
	"fmt"

	mem_cache "github.com/wangxin688/narvis/intend/cache"
	intend_device "github.com/wangxin688/narvis/intend/model/device"
	"github.com/wangxin688/narvis/intend/model/snmp"
	"github.com/wangxin688/narvis/intend/model/wlanstation"
	"github.com/wangxin688/narvis/intend/netdisco/factory"
)

const ruckusZDWLANAPIPAddr = ".1.3.6.1.4.1.25053.1.2.2.1.1.2.1.1.10"
const ruckusZDWLANAPDescription = ".1.3.6.1.4.1.25053.1.2.2.1.1.2.1.1.2"
const ruckusZDWLANAPSWversion = ".1.3.6.1.4.1.25053.1.2.2.1.1.2.1.1.7"
const ruckusZDWLANAPModel = ".1.3.6.1.4.1.25053.1.2.2.1.1.2.1.1.4"
const ruckusZDWLANAPSerialNumber = ".1.3.6.1.4.1.25053.1.2.2.1.1.2.1.1.5"

const ruckusZDWLANStaMacAddr = "1.3.6.1.4.1.25053.1.2.2.1.1.3.1.1.1"
const ruckusZDWLANStaAPMacAddr = ".1.3.6.1.4.1.25053.1.2.2.1.1.3.1.1.2"
const ruckusZDWLANStaSSID = ".1.3.6.1.4.1.25053.1.2.2.1.1.3.1.1.4"
const ruckusZDWLANStaUser = ".1.3.6.1.4.1.25053.1.2.2.1.1.3.1.1.5"
const ruckusZDWLANStaChannel = ".1.3.6.1.4.1.25053.1.2.2.1.1.3.1.1.7"
const ruckusZDWLANStaIPAddr = ".1.3.6.1.4.1.25053.1.2.2.1.1.3.1.1.8"
const ruckusZDWLANStaAvgRSSI = ".1.3.6.1.4.1.25053.1.2.2.1.1.3.1.1.9"
const ruckusZDWLANStaSNR = ".1.3.6.1.4.1.25053.1.2.2.1.1.3.1.1.21"
const ruckusZDWLANStaRxBytes = ".1.3.6.1.4.1.25053.1.2.2.1.1.3.1.1.11"
const ruckusZDWLANStaTxBytes = ".1.3.6.1.4.1.25053.1.2.2.1.1.3.1.1.13"
const ruckusZDWLANStaAssocTime = ".1.3.6.1.4.1.25053.1.2.2.1.1.3.1.1.15"
const ruckusZDWLANStaVlanID = ".1.3.6.1.4.1.25053.1.2.2.1.1.3.1.1.30"

type RuckusDriver struct {
	factory.SnmpDiscovery
}

func NewRuckusDriver(sc *snmp.SnmpConfig) (*RuckusDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &RuckusDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}

func (r *RuckusDriver) APs() (ap []*intend_device.Ap, errors []string) {

	apIP, errApIP := r.Session.BulkWalkAll(ruckusZDWLANAPIPAddr)
	if len(apIP) == 0 || errApIP != nil {
		return nil, []string{fmt.Sprintf("failed to get ap ipAddress from %s", r.IpAddress)}
	}
	apMac, errApMac := r.Session.BulkWalkAll(ruckusZDWLANAPDescription)
	apName, errApName := r.Session.BulkWalkAll(ruckusZDWLANAPSWversion)
	apType, errApType := r.Session.BulkWalkAll(ruckusZDWLANAPModel)
	apVersion, errApVersion := r.Session.BulkWalkAll(ruckusZDWLANAPSWversion)
	apSerialNumber, errApSerialNumber := r.Session.BulkWalkAll(ruckusZDWLANAPSerialNumber)
	if errApMac != nil || errApName != nil || errApType != nil || errApSerialNumber != nil || errApVersion != nil {
		errors = append(errors, errApMac.Error())
		errors = append(errors, errApName.Error())
		errors = append(errors, errApType.Error())
		errors = append(errors, errApVersion.Error())
		errors = append(errors, errApSerialNumber.Error())
	}
	indexApIP := factory.ExtractString(ruckusZDWLANAPIPAddr, apIP)
	indexApMac := factory.ExtractMacAddress(ruckusZDWLANAPDescription, apMac)
	indexApName := factory.ExtractString(ruckusZDWLANAPSWversion, apName)
	indexApType := factory.ExtractString(ruckusZDWLANAPModel, apType)
	indexApSerialNumber := factory.ExtractString(ruckusZDWLANAPSerialNumber, apSerialNumber)
	indexApVersion := factory.ExtractString(ruckusZDWLANAPSWversion, apVersion)
	for i, v := range indexApIP {
		ap_name := indexApName[i]
		ap_mac := indexApMac[i]
		mem_cache.MemCache.SetDefault(ap_mac, ap_name)
		apVer := indexApVersion[i]
		ap = append(ap, &intend_device.Ap{
			Name:         indexApName[i],
			ManagementIp: v,
			MacAddress:   indexApMac[i],
			DeviceModel:  indexApType[i],
			SerialNumber: indexApSerialNumber[i],
			OsVersion:    &apVer,
		})
	}
	return ap, errors
}

func (r *RuckusDriver) WlanUsers() (wlanUsers []*wlanstation.WlanUser, errors []string) {

	results := make([]*wlanstation.WlanUser, 0)
	userNames, err := r.Session.BulkWalkAll(ruckusZDWLANStaUser)
	if err != nil {
		return nil, []string{fmt.Sprintf("failed to get users from %s", r.IpAddress)}
	}
	apMac, errMac := r.Session.BulkWalkAll(ruckusZDWLANStaAPMacAddr)
	userUptime, errUptime := r.Session.BulkWalkAll(ruckusZDWLANStaAssocTime)
	userAssignedVlan, errAssignedVlan := r.Session.BulkWalkAll(ruckusZDWLANStaVlanID)
	userRSSI, errRSSI := r.Session.BulkWalkAll(ruckusZDWLANStaAvgRSSI)
	userSnr, errSnr := r.Session.BulkWalkAll(ruckusZDWLANStaSNR)
	userESSID, errESSID := r.Session.BulkWalkAll(ruckusZDWLANStaSSID)
	userChannel, errChannel := r.Session.BulkWalkAll(ruckusZDWLANStaChannel)
	userTxBytes, errTxBytes := r.Session.BulkWalkAll(ruckusZDWLANStaTxBytes)
	userRxBytes, errRxBytes := r.Session.BulkWalkAll(ruckusZDWLANStaRxBytes)
	userIp, errIp := r.Session.BulkWalkAll(ruckusZDWLANStaIPAddr)

	if errUptime != nil || errAssignedVlan != nil || errRSSI != nil || errSnr != nil ||
		errESSID != nil || errIp != nil || errMac != nil ||
		errChannel != nil || errTxBytes != nil || errRxBytes != nil {
		errors = append(errors, errUptime.Error())
		errors = append(errors, errAssignedVlan.Error())
		errors = append(errors, errSnr.Error())
		errors = append(errors, errESSID.Error())
		errors = append(errors, errChannel.Error())
		errors = append(errors, errTxBytes.Error())
		errors = append(errors, errRxBytes.Error())
		errors = append(errors, errIp.Error())
		errors = append(errors, errMac.Error())
	}
	indexApMac := factory.ExtractMacAddress(ruckusZDWLANStaAPMacAddr, apMac)
	indexUserName := factory.ExtractString(ruckusZDWLANStaUser, userNames)
	indexUserUptime := factory.ExtractInteger(ruckusZDWLANStaAssocTime, userUptime)
	indexUserVlan := factory.ExtractInteger(ruckusZDWLANStaVlanID, userAssignedVlan)
	indexSnr := factory.ExtractInteger(ruckusZDWLANStaSNR, userSnr)
	indexESSID := factory.ExtractString(ruckusZDWLANStaSSID, userESSID)
	indexRSSI := factory.ExtractInteger(ruckusZDWLANStaAvgRSSI, userRSSI)
	indexChannel := factory.ExtractInteger(ruckusZDWLANStaChannel, userChannel)
	indexTxBytes := factory.ExtractInteger(ruckusZDWLANStaTxBytes, userTxBytes)
	indexRxBytes := factory.ExtractInteger(ruckusZDWLANStaRxBytes, userRxBytes)
	indexIp := factory.ExtractString(ruckusZDWLANStaIPAddr, userIp)
	for i, v := range indexUserName {

		vlan := uint16(indexUserVlan[i])
		channel := uint16(indexChannel[i])
		snr := uint8(indexSnr[i])
		ap_mac := indexApMac[i]
		ap_name := r.getApName(ap_mac)
		user := wlanstation.WlanUser{
			StationMac:        factory.StringToHexMac(i),
			StationIp:         indexIp[i],
			StationUsername:   v,
			StationApMac:      &ap_mac,
			StationApName:     &ap_name,
			StationESSID:      indexESSID[i],
			StationSNR:        &snr,
			StationRSSI:       int8(indexRSSI[i]) * -1,
			StationVlan:       &vlan,
			StationOnlineTime: indexUserUptime[i],
			StationChannel:    channel,
			StationRxBits:     indexRxBytes[i] * 8,
			StationTxBits:     indexTxBytes[i] * 8,
			StationRadioType:  factory.ChannelToRadioType(channel),
		}
		results = append(results, &user)
	}
	return results, errors
}

func (r *RuckusDriver) getApName(apMac string) string {
	apName, ok := mem_cache.MemCache.Get(apMac)
	if !ok {
		return ""
	}
	return apName.(string)
}
