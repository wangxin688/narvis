package driver

import (
	"fmt"

	"github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"
)

const ruckusZDWLANAPIPAddr = ".1.3.6.1.4.1.25053.1.2.2.1.1.2.1.1.10"
const ruckusZDWLANAPDescription = ".1.3.6.1.4.1.25053.1.2.2.1.1.2.1.1.2"
const ruckusZDWLANAPSWversion = ".1.3.6.1.4.1.25053.1.2.2.1.1.2.1.1.7"
const ruckusZDWLANAPModel = ".1.3.6.1.4.1.25053.1.2.2.1.1.2.1.1.4"
const ruckusZDWLANAPSerialNumber = ".1.3.6.1.4.1.25053.1.2.2.1.1.2.1.1.5"

type RuckusDriver struct {
	factory.SnmpDiscovery
}

func NewRuckusDriver(sc factory.SnmpConfig) (*RuckusDriver, error) {
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

func (r *RuckusDriver) APs() (ap []*factory.ApItem, errors []string) {

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
		ap = append(ap, &factory.ApItem{
			Name:         indexApName[i],
			ManagementIp: v,
			MacAddress:   indexApMac[i],
			DeviceModel:  indexApType[i],
			SerialNumber: indexApSerialNumber[i],
			OsVersion:    indexApVersion[i],
		})
	}
	return ap, errors
}
