package driver

import "github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"


type AristaDriver struct {
	factory.SnmpDiscovery
}

func NewAristaDriver(sc factory.SnmpConfig) (*AristaDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &AristaDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}
