package driver

import "github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"

type HuaweiDriver struct {
	factory.SnmpDiscovery
}


func NewHuaweiDriver(sc factory.SnmpConfig) (*HuaweiDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &HuaweiDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}

