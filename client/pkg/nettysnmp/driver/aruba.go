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
