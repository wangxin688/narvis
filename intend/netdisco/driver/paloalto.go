package driver

import (
	nettyx_snmp "github.com/wangxin688/narvis/intend/model/snmp"
	"github.com/wangxin688/narvis/intend/netdisco/factory"
)

type PaloAltoDriver struct {
	factory.SnmpDiscovery
}

func NewPaloAltoDriver(sc *nettyx_snmp.SnmpConfig) (*PaloAltoDriver, error) {
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
