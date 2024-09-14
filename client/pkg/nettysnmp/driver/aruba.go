package driver

import "github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"

const wlanApMacAddress string = ".1.3.6.1.4.1.14823.2.2.1.5.2.1.4.1.1"
const wlanApIpAddress string = ".1.3.6.1.4.1.14823.2.2.1.5.2.1.3.1.2"
const wlanApName string = ".1.3.6.1.4.1.14823.2.2.1.5.2.1.3.1.3"
const wlanApGroupName string = ".1.3.6.1.4.1.14823.2.2.1.5.2.1.3.1.5"
const wlanApModel string = ".1.3.6.1.4.1.14823.2.2.1.5.2.1.3.1.5"
const wlanApSerialNumber string = ".1.3.6.1.4.1.14823.2.2.1.5.2.1.3.1.6"
const wlanApSwitchIpAddress string = ".1.3.6.1.4.1.14823.2.2.1.5.2.1.3.1.39"

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

func (ad *ArubaDriver) ScanAp() (ap []*factory.ApItem, errors []string) {
	results := make([]*factory.ApItem, 0)
	return results, nil
}
