package driver

import (
	"strconv"
	"strings"

	"github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"
)

const vlanNameName = ".1.3.6.1.4.1.4526.17.48.21.1.1"
const vlanNameIfIndex = ".1.3.6.1.4.1.4526.17.48.21.1.3"

type NetgearDriver struct {
	factory.SnmpDiscovery
}

func (nd *NetgearDriver) Vlans() (vlan []*factory.VlanItem, errors []string) {
	l2Vlan, err := nd.Session.BulkWalkAll(vlanNameName)
	l2VlanIfIndex, errIfIndex := nd.Session.BulkWalkAll(vlanNameIfIndex)
	if err != nil || errIfIndex != nil {
		errors = append(errors, err.Error())
		errors = append(errors, errIfIndex.Error())
	}
	indexL2Vlan := factory.ExtractString(vlanNameName, l2Vlan)
	indexVlanIndex := factory.ExtractInteger(vlanNameIfIndex, l2VlanIfIndex)

	for i, v := range indexL2Vlan {
		vlanIdString := strings.TrimPrefix(v, ".")
		vlanId, _ := strconv.Atoi(vlanIdString)
		_vlan := &factory.VlanItem{
			VlanId:   uint32(vlanId),
			VlanName: v,
			IfIndex:  indexVlanIndex[i],
		}
		vlan = append(vlan, _vlan)
	}

	return vlan, errors
}

func NewNetgearDriver(sc factory.SnmpConfig) (*NetgearDriver, error) {
	session, err := factory.NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &NetgearDriver{
		factory.SnmpDiscovery{
			Session:   session,
			IpAddress: session.Target},
	}, nil
}
