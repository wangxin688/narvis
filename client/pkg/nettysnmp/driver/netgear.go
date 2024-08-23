package driver

import "github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"

type NetgearDriver struct {
	factory.SnmpDiscovery
}

func NewNetgearDriver(sc factory.SnmpConfig) (*NetgearDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &NetgearDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}
