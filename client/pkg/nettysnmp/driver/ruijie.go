package driver

import (
	"fmt"

	"github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"
	mem_cache "github.com/wangxin688/narvis/intend/cache"
)

type RuiJieDriver struct {
	factory.SnmpDiscovery
}

const ruijieApMacAddr = ".1.3.6.1.4.1.4881.1.1.10.2.56.2.1.1.1.1"
const ruijieApApName = ".1.3.6.1.4.1.4881.1.1.10.2.56.2.1.1.1.2"
const ruijieApApgName = ".1.3.6.1.4.1.4881.1.1.10.2.56.2.1.1.1.3"
const ruijieApApSn = ".1.3.6.1.4.1.4881.1.1.10.2.56.2.1.1.1.32"
const ruijieApApIp = ".1.3.6.1.4.1.4881.1.1.10.2.56.2.1.1.1.33"
const ruijieApApSwVer = ".1.3.6.1.4.1.4881.1.1.10.2.56.2.1.1.1.37"
const ruijieApApPID = ".1.3.6.1.4.1.4881.1.1.10.2.56.2.1.1.1.39"

const ruijieStaMacAddr = ".1.3.6.1.4.1.4881.1.1.10.2.56.5.1.1.1.1"
const ruijieStaApMacAddr = ".1.3.6.1.4.1.4881.1.1.10.2.56.5.1.1.1.2"
const ruijieStaVlan = ".1.3.6.1.4.1.4881.1.1.10.2.56.5.1.1.1.3"
const ruijieStaIp = ".1.3.6.1.4.1.4881.1.1.10.2.56.5.1.1.1.5"
const ruijieStaSsid = ".1.3.6.1.4.1.4881.1.1.10.2.56.5.1.1.1.7"
const ruijieStaLinkRate = ".1.3.6.1.4.1.4881.1.1.10.2.56.5.1.1.1.18"
const ruijieStaCurChan = ".1.3.6.1.4.1.4881.1.1.10.2.56.5.1.1.1.19"
const ruijieStaRssi = ".1.3.6.1.4.1.4881.1.1.10.2.56.5.1.1.1.21"
const ruijieStaUsername = ".1.3.6.1.4.1.4881.1.1.10.2.56.5.1.1.1.22"
const ruijieStaOnlineTime = ".1.3.6.1.4.1.4881.1.1.10.2.56.5.1.1.1.24"

func NewRuiJieDriver(sc factory.SnmpConfig) (*RuiJieDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &RuiJieDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target,
		},
	}, nil
}

func (d *RuiJieDriver) APs() (ap []*factory.ApItem, errors []string) {
	apIP, errApIP := d.Session.BulkWalkAll(ruijieApApIp)
	apMac, errApMac := d.Session.BulkWalkAll(ruijieApMacAddr)
	apName, errApName := d.Session.BulkWalkAll(ruijieApApName)
	apGroup, errApGroup := d.Session.BulkWalkAll(ruijieApApgName)
	apSerialNumber, errApSerialNumber := d.Session.BulkWalkAll(ruijieApApSn)
	apVersion, errApVersion := d.Session.BulkWalkAll(ruijieApApSwVer)
	apPID, errApPID := d.Session.BulkWalkAll(ruijieApApPID)
	if errApMac != nil || errApName != nil || errApGroup != nil || errApSerialNumber != nil || errApVersion != nil || errApPID != nil || errApIP != nil {
		errors = append(errors, errApMac.Error())
		errors = append(errors, errApName.Error())
		errors = append(errors, errApGroup.Error())
		errors = append(errors, errApSerialNumber.Error())
		errors = append(errors, errApVersion.Error())
		errors = append(errors, errApPID.Error())
		errors = append(errors, errApIP.Error())
	}
	indexApIP := factory.ExtractString(ruijieApApIp, apIP)
	indexApMac := factory.ExtractMacAddress(ruijieApMacAddr, apMac)
	indexApName := factory.ExtractString(ruijieApApName, apName)
	indexApGroup := factory.ExtractString(ruijieApApgName, apGroup)
	indexApSerialNumber := factory.ExtractString(ruijieApApSn, apSerialNumber)
	indexApVersion := factory.ExtractString(ruijieApApSwVer, apVersion)
	indexApPID := factory.ExtractString(ruijieApApPID, apPID)
	for i, v := range indexApIP {
		ap_mac := indexApMac[i]
		ap_name := indexApName[i]
		mem_cache.MemCache.SetDefault(ap_mac, ap_name)
		ap = append(ap, &factory.ApItem{
			Name:         indexApName[i],
			ManagementIp: v,
			MacAddress:   indexApMac[i],
			DeviceModel:  indexApPID[i],
			SerialNumber: indexApSerialNumber[i],
			OsVersion:    indexApVersion[i],
			GroupName:    indexApGroup[i],
		})
	}
	return ap, errors
}

func (d *RuiJieDriver) WlanUsers() *factory.WlanUserResponse {
	results := make([]*factory.WlanUser, 0)
	errors := make([]string, 0)
	userNames, err := d.Session.BulkWalkAll(ruijieStaUsername)
	if err != nil {
		return &factory.WlanUserResponse{
			Errors:    []string{fmt.Sprintf("failed to get users from %s", d.IpAddress), err.Error()},
			WlanUsers: results,
		}
	}
	userApMac, errApMac := d.Session.BulkWalkAll(ruijieStaApMacAddr)
	userUptime, errUptime := d.Session.BulkWalkAll(ruijieStaOnlineTime)
	userAssignedVlan, errAssignedVlan := d.Session.BulkWalkAll(ruijieStaVlan)
	userRSSI, errRSSI := d.Session.BulkWalkAll(ruijieStaRssi)
	userESSID, errESSID := d.Session.BulkWalkAll(ruijieStaSsid)
	userChannel, errChannel := d.Session.BulkWalkAll(ruijieStaCurChan)
	// userTxBytes, errTxBytes := d.Session.BulkWalkAll(hwWlanStaWirelessTxBytes)
	// userRxBytes, errRxBytes := d.Session.BulkWalkAll(hwWlanStaWirelessRxBytes)
	userIp, errIp := d.Session.BulkWalkAll(ruijieStaIp)

	if errUptime != nil || errAssignedVlan != nil || errRSSI != nil ||
		errApMac != nil || errESSID != nil || errIp != nil ||
		errChannel != nil {
		errors = append(errors, errUptime.Error())
		errors = append(errors, errAssignedVlan.Error())
		errors = append(errors, errESSID.Error())
		errors = append(errors, errChannel.Error())
		errors = append(errors, errIp.Error())
	}
	indexApMac := factory.ExtractMacAddress(ruijieStaApMacAddr, userApMac)
	indexUserName := factory.ExtractString(ruijieStaUsername, userNames)
	indexUserUptime := factory.ExtractInteger(ruijieStaOnlineTime, userUptime)
	indexUserVlan := factory.ExtractInteger(ruijieStaVlan, userAssignedVlan)
	indexESSID := factory.ExtractString(ruijieStaSsid, userESSID)
	indexRSSI := factory.ExtractInteger(ruijieStaRssi, userRSSI)
	indexChannel := factory.ExtractInteger(ruijieStaCurChan, userChannel)
	// indexTxBytes := factory.ExtractInteger(wlanStaTxBytes, userTxBytes)
	// indexRxBytes := factory.ExtractInteger(wlanStaRxBytes, userRxBytes)
	indexIp := factory.ExtractString(hwWlanStaIP, userIp)
	for i, v := range indexUserName {
		vlan := indexUserVlan[i]
		channel := indexChannel[i]
		ap_mac := indexApMac[i]
		ap_name := d.getAPName(ap_mac)
		user := factory.WlanUser{
			StationMac:        factory.StringToHexMac(i),
			StationIp:         indexIp[i],
			StationUsername:   v,
			StationApMac:      &ap_mac,
			StationApName:     &ap_name,
			StationESSID:      indexESSID[i],
			StationRSSI:       indexRSSI[i],
			StationVlan:       &vlan,
			StationOnlineTime: indexUserUptime[i],
			StationChannel:    channel,
			StationRadioType:  factory.ChannelToRadioType(channel),
		}
		results = append(results, &user)
	}
	return &factory.WlanUserResponse{
		WlanUsers: results,
		Errors:    errors,
	}
}

func (d *RuiJieDriver) getAPName(apMac string) string {
	apName, ok := mem_cache.MemCache.Get(apMac)
	if !ok {
		return ""
	}
	return apName.(string)
}
