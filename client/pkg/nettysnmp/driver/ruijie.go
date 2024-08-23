package driver

import "github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"

type RuiJieDriver struct {
	factory.SnmpDiscovery
}

func NewRuiJieDriver(sc factory.SnmpConfig) (*RuiJieDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &RuiJieDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target,
		},
	}, nil
}
