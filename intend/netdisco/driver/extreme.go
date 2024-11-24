package driver

import (
	nettyx_snmp "github.com/wangxin688/narvis/intend/model/snmp"
	"github.com/wangxin688/narvis/intend/netdisco/factory"
)

type ExtremeDriver struct {
	factory.SnmpDiscovery
}

func NewExtremeDriver(sc *nettyx_snmp.SnmpConfig) (*ExtremeDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &ExtremeDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}
