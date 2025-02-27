package driver

import (
	"github.com/wangxin688/narvis/intend/model/snmp"
	"github.com/wangxin688/narvis/intend/netdisco/factory"
)

type TPLinkDriver struct {
	factory.SnmpDiscovery
}

func NewTPLinkDriver(sc *snmp.SnmpConfig) (*TPLinkDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &TPLinkDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}
