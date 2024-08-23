package driver

import "github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"

type ExtremeDriver struct {
	factory.SnmpDiscovery
}

func NewExtremeDriver(sc factory.SnmpConfig) (*ExtremeDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &ExtremeDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}
