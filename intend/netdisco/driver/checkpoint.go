package driver

import (
	nettyx_snmp "github.com/wangxin688/narvis/intend/model/snmp"
	"github.com/wangxin688/narvis/intend/netdisco/factory"
)

type CheckPointDriver struct {
	factory.SnmpDiscovery
}

func NewCheckPointDriver(sc *nettyx_snmp.SnmpConfig) (*CheckPointDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &CheckPointDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}
