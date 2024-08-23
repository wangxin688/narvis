package driver

import "github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"

type JuniperDriver struct {
	factory.SnmpDiscovery
}

func NewJuniperDriver(sc factory.SnmpConfig) (*JuniperDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &JuniperDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}
