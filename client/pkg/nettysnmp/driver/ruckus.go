package driver

import "github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"

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
