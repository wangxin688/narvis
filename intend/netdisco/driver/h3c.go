package driver

import (
	"fmt"

	mem_cache "github.com/wangxin688/narvis/intend/cache"
	intend_device "github.com/wangxin688/narvis/intend/model/device"
	"github.com/wangxin688/narvis/intend/model/snmp"
	"github.com/wangxin688/narvis/intend/model/wlanstation"
	"github.com/wangxin688/narvis/intend/netdisco/factory"
)

// WLAN AP-MIB: https://www.h3c.com/cn/d_202401/2021242_30005_0.htm#_Toc146378887
// need more test for devices
// HH3C-DOT11-APMT-MIB:

const hh3cDot11CurrAPIPAddress string = ".1.3.6.1.4.1.25506.2.75.2.1.2.1.2"
const hh3cDot11CurrAPMacAddress string = ".1.3.6.1.4.1.25506.2.75.2.1.2.1.3"
const hh3cDot11CurrAPName string = ".1.3.6.1.4.1.25506.2.75.2.1.2.1.8"
const hh3cDot11CurrAPModelName string = ".1.3.6.1.4.1.25506.2.75.2.1.2.1.9"
const hh3cDot11CurrAPTemplateName string = ".1.3.6.1.4.1.25506.2.75.2.1.2.1.6"
const hh3cDot11CurrAPSoftwareVersion string = ".1.3.6.1.4.1.25506.2.75.2.1.2.1.11"

const hh3cDot11StationMAC = ".1.3.6.1.4.1.25506.2.75.3.1.1.1.1"
const hh3cDot11StationIPAddress = ".1.3.6.1.4.1.25506.2.75.3.1.1.1.2"
const hh3cDot11StationUserName = ".1.3.6.1.4.1.25506.2.75.3.1.1.1.3"
const hh3cDot11StationSSIDName = ".1.3.6.1.4.1.25506.2.75.3.1.1.1.12"
const hh3cDot11StationVlanId = ".1.3.6.1.4.1.25506.2.75.3.1.1.1.11"
const hh3cDot11StationRxSNR = ".1.3.6.1.4.1.25506.2.75.3.1.1.1.17"
const hh3cDot11StationAssociateAPMACAddressCM = ".1.3.6.1.4.1.25506.2.75.3.1.1.1.40"
const hh3cDot11StationRSSI = ".1.3.6.1.4.1.25506.2.75.3.1.1.1.7"
const hh3cDot11StationChannel = ".1.3.6.1.4.1.25506.2.75.3.1.1.1.8"
const hh3cDot11StationTxSpeed = ".1.3.6.1.4.1.25506.2.75.3.1.1.1.24"
const hh3cDot11StationRxSpeed = ".1.3.6.1.4.1.25506.2.75.3.1.1.1.25"
const hh3cDot11StationMaxRate = ". 1.3.6.1.4.1.25506.2.75.3.1.1.1.33" // (Mbps)
const hh3cDot11StationAssTime = ".1.3.6.1.4.1.25506.2.75.3.1.1.1.30"
const hh3cDot11StationVendorName = ".1.3.6.1.4.1.25506.2.75.3.1.1.1.20"

type H3CDriver struct {
	factory.SnmpDiscovery
}

func NewH3CDriver(sc *snmp.SnmpConfig) (*H3CDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &H3CDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}

func (hd *H3CDriver) APs() (ap []*intend_device.Ap, errors []string) {
	apIp, errApIp := hd.Session.BulkWalkAll(hh3cDot11CurrAPIPAddress)
	if len(apIp) == 0 || errApIp != nil {
		return nil, []string{fmt.Sprintf("failed to get ap ipAddress from %s", hd.IpAddress)}
	}
	apMac, errApMac := hd.Session.BulkWalkAll(hh3cDot11CurrAPMacAddress)
	apName, errApName := hd.Session.BulkWalkAll(hh3cDot11CurrAPName)
	apModel, errApModel := hd.Session.BulkWalkAll(hh3cDot11CurrAPModelName)
	apGroup, errApGroup := hd.Session.BulkWalkAll(hh3cDot11CurrAPTemplateName)
	apVersion, errApVersion := hd.Session.BulkWalkAll(hh3cDot11CurrAPSoftwareVersion)
	if errApMac != nil || errApName != nil || errApModel != nil || errApGroup != nil || errApVersion != nil {
		errors = append(errors, errApMac.Error())
		errors = append(errors, errApName.Error())
		errors = append(errors, errApModel.Error())
		errors = append(errors, errApGroup.Error())
	}
	indexApIP := factory.ExtractString(hh3cDot11CurrAPIPAddress, apIp)
	indexApMac := factory.ExtractMacAddress(hh3cDot11CurrAPMacAddress, apMac)
	indexApName := factory.ExtractString(hh3cDot11CurrAPName, apName)
	indexApModel := factory.ExtractString(hh3cDot11CurrAPModelName, apModel)
	indexApGroup := factory.ExtractString(hh3cDot11CurrAPTemplateName, apGroup)
	indexApVersion := factory.ExtractString(hh3cDot11CurrAPSoftwareVersion, apVersion)
	for i, v := range indexApIP {
		ap_mac := indexApMac[i]
		ap_name := indexApName[i]
		group_name := indexApGroup[i]
		acIp := hd.IpAddress
		apVer := indexApVersion[i]
		mem_cache.MemCache.SetDefault(ap_mac, ap_name)
		ap = append(ap, &intend_device.Ap{
			Name:            indexApName[i],
			ManagementIp:    v,
			MacAddress:      indexApMac[i],
			DeviceModel:     indexApModel[i],
			GroupName:       &group_name,
			WlanACIpAddress: &acIp,
			OsVersion:       &apVer,
		})
	}
	return ap, errors
}

func (hd *H3CDriver) WlanUsers() (wlanUsers []*wlanstation.WlanUser, errors []string) {

	results := make([]*wlanstation.WlanUser, 0)
	userNames, err := hd.Session.BulkWalkAll(hh3cDot11StationUserName)
	if err != nil {
		return nil, []string{fmt.Sprintf("failed to get wlan user name from %s", hd.IpAddress)}
	}
	apMac, errMac := hd.Session.BulkWalkAll(hh3cDot11StationAssociateAPMACAddressCM)
	userUptime, errUptime := hd.Session.BulkWalkAll(hh3cDot11StationAssTime)
	userAssignedVlan, errAssignedVlan := hd.Session.BulkWalkAll(hh3cDot11StationVlanId)
	userRSSI, errRSSI := hd.Session.BulkWalkAll(hh3cDot11StationRSSI)
	userSnr, errSnr := hd.Session.BulkWalkAll(hh3cDot11StationRxSNR)
	userESSID, errESSID := hd.Session.BulkWalkAll(hh3cDot11StationSSIDName)
	userChannel, errChannel := hd.Session.BulkWalkAll(hh3cDot11StationChannel)
	userTxBytes, errTxBytes := hd.Session.BulkWalkAll(hh3cDot11StationTxSpeed)
	userRxBytes, errRxBytes := hd.Session.BulkWalkAll(hh3cDot11StationRxSpeed)
	userIp, errIp := hd.Session.BulkWalkAll(hh3cDot11StationIPAddress)
	userSpeed, errSpeed := hd.Session.BulkWalkAll(hh3cDot11StationMaxRate)

	if errUptime != nil || errAssignedVlan != nil || errRSSI != nil || errSnr != nil ||
		errESSID != nil || errIp != nil || errMac != nil || errSpeed != nil ||
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
		errors = append(errors, errSpeed.Error())
	}
	indexApMac := factory.ExtractMacAddress(hh3cDot11StationAssociateAPMACAddressCM, apMac)
	indexUserName := factory.ExtractString(hh3cDot11StationUserName, userNames)
	indexUserUptime := factory.ExtractInteger(hh3cDot11StationAssTime, userUptime)
	indexUserVlan := factory.ExtractInteger(hh3cDot11StationVlanId, userAssignedVlan)
	indexSnr := factory.ExtractInteger(hh3cDot11StationRxSNR, userSnr)
	indexESSID := factory.ExtractString(hh3cDot11StationSSIDName, userESSID)
	indexRSSI := factory.ExtractInteger(hh3cDot11StationRSSI, userRSSI)
	indexChannel := factory.ExtractInteger(hh3cDot11StationChannel, userChannel)
	indexTxBytes := factory.ExtractInteger(hh3cDot11StationTxSpeed, userTxBytes)
	indexRxBytes := factory.ExtractInteger(hh3cDot11StationRxSpeed, userRxBytes)
	indexIp := factory.ExtractString(hh3cDot11StationIPAddress, userIp)
	indexSpeed := factory.ExtractInteger(hh3cDot11StationMaxRate, userSpeed)
	for i, v := range indexUserName {

		vlan := uint16(indexUserVlan[i])
		channel := uint16(indexChannel[i])
		snr := uint8(indexSnr[i])
		ap_mac := indexApMac[i]
		ap_name := hd.getApName(ap_mac)
		speed := indexSpeed[i]
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
			StationMaxSpeed:   &speed,
		}
		results = append(results, &user)
	}
	return results, errors
}

func (hd *H3CDriver) getApName(apMac string) string {
	apName, ok := mem_cache.MemCache.Get(apMac)
	if !ok {
		return ""
	}
	return apName.(string)
}
