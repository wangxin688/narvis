package driver

import "github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"

type SangforDriver struct {
	factory.SnmpDiscovery
}

func NewSangforDriver(sc factory.SnmpConfig) (*SangforDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &SangforDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}
