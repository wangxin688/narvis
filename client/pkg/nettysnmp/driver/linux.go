package driver

import "github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"

type LinuxDriver struct {
	factory.SnmpDiscovery
}

func NewLinuxDriver(sc factory.SnmpConfig) (*LinuxDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &LinuxDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}
