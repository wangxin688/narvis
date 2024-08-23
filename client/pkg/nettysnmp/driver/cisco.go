package driver

import "github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"

type CiscoIosDriver struct {
	factory.SnmpDiscovery
}

type CiscoIosXRDriver struct {
	factory.SnmpDiscovery
}

type CiscoNexusOSDriver struct {
	factory.SnmpDiscovery
}

type CiscoIosXEDriver struct {
	factory.SnmpDiscovery
}

func NewCiscoIosDriver(sc factory.SnmpConfig) (*CiscoIosDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &CiscoIosDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}

func NewCiscoIosXRDriver(sc factory.SnmpConfig) (*CiscoIosXRDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &CiscoIosXRDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}

func NewCiscoNexusOSDriver(sc factory.SnmpConfig) (*CiscoNexusOSDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &CiscoNexusOSDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}

func NewCiscoIosXEDriver(sc factory.SnmpConfig) (*CiscoIosXEDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &CiscoIosXEDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}
