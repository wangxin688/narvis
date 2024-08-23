package driver

import "github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"

type HuaweiDriver struct {
	factory.SnmpDiscovery
}

type HuaweiVrpDriver struct {
	factory.SnmpDiscovery
}

type HuaweiVrpV8Driver struct {
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

func NewHuaweiVrpDriver(sc factory.SnmpConfig) (*HuaweiVrpDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &HuaweiVrpDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}

func NewHuaweiVrpV8Driver(sc factory.SnmpConfig) (*HuaweiVrpV8Driver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &HuaweiVrpV8Driver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}
