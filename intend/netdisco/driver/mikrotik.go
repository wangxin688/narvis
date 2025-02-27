package driver

import (
	"github.com/wangxin688/narvis/intend/model/snmp"
	"github.com/wangxin688/narvis/intend/netdisco/factory"
)

type MikroTikDriver struct {
	factory.SnmpDiscovery
}

func NewMikroTikDriver(sc *snmp.SnmpConfig) (*MikroTikDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &MikroTikDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}
