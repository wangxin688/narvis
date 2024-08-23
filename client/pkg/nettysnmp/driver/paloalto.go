package driver

import "github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"

type PaloAltoDriver struct {
	factory.SnmpDiscovery
}

func NewPaloAltoDriver(sc factory.SnmpConfig) (*PaloAltoDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &PaloAltoDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}
