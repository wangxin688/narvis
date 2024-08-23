package driver

import "github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"

type F5Driver struct {
	factory.SnmpDiscovery
}

func NewF5Driver(sc factory.SnmpConfig) (*F5Driver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &F5Driver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}
