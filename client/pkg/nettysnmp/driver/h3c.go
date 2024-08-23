package driver

import "github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"

type H3CDriver struct {
	factory.SnmpDiscovery
}

func NewH3CDriver(sc factory.SnmpConfig) (*H3CDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &H3CDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}
