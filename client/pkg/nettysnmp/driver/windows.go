package driver

import "github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"

type WindowsDriver struct {
	factory.SnmpDiscovery
}


func NewWindowsDriver(sc factory.SnmpConfig) (*WindowsDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &WindowsDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}