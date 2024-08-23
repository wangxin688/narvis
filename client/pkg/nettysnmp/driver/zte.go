package driver

import "github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"

type ZTEDriver struct {
	factory.SnmpDiscovery
}

func NewZTEDriver(sc factory.SnmpConfig) (*ZTEDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &ZTEDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}
