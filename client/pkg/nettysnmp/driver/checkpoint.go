package driver

import "github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"



type CheckPointDriver struct {
	factory.SnmpDiscovery
}

func NewCheckPointDriver(sc factory.SnmpConfig) (*CheckPointDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &CheckPointDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}
