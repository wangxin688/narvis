package driver

import "github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"

type FortiNetDriver struct {
	factory.SnmpDiscovery
}

func NewFortiNetDriver(sc factory.SnmpConfig) (*FortiNetDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &FortiNetDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}
