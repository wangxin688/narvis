package driver

import "github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"

type TPLinkDriver struct {
	factory.SnmpDiscovery
}

func NewTPLinkDriver(sc factory.SnmpConfig) (*TPLinkDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &TPLinkDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}
