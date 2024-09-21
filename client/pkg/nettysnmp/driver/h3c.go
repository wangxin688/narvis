package driver

import "github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"

// WLAN AP-MIB: https://www.h3c.com/cn/d_202401/2021242_30005_0.htm#_Toc146378887
// HH3C-DOT11-APMT-MIB:

const hh3cDot11ApMacAddress string = ".1.3.6.1.4.1.25506.2.75.2.1.1.1.2"

// const hh3cDot11APID =

type H3CDriver struct {
	factory.SnmpDiscovery
}

func NewH3CDriver(sc factory.SnmpConfig) (*H3CDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &H3CDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}
