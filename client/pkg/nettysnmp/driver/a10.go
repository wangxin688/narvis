package driver

import "github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"

type A10Driver struct {
	factory.SnmpDiscovery
}

func NewA10Driver(sc factory.SnmpConfig) (*A10Driver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &A10Driver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}
