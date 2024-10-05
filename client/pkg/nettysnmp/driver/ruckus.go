package driver

import "github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"

const ruckusZDWLANAPIPAddr  = ".1.3.6.1.4.1.25053.1.2.2.1.1.2.1.1.10"
const ruckusZDWLANAPDescription  = ".1.3.6.1.4.1.25053.1.2.2.1.1.2.1.1.2"
const ruckusZDWLANAPSWversion = ".1.3.6.1.4.1.25053.1.2.2.1.1.2.1.1.7"
const  ruckusZDWLANAPModel = ".1.3.6.1.4.1.25053.1.2.2.1.1.2.1.1.4"
const  ruckusZDWLANAPSerialNumber = ".1.3.6.1.4.1.25053.1.2.2.1.1.2.1.1.5"

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
