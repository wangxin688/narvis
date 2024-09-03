package nettygo

import (
	"fmt"
	"net"
	"strings"
	"sync"
	"time"

	"github.com/go-ping/ping"
	"github.com/gosnmp/gosnmp"
	dt "github.com/wangxin688/narvis/client/pkg/nettysnmp/devicemodel"
	s "github.com/wangxin688/narvis/client/pkg/nettysnmp/devicemodel/sysobjectid"
	"github.com/wangxin688/narvis/client/pkg/nettysnmp/driver"
	"github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"
	"github.com/wangxin688/narvis/intend/manufacturer"
	"github.com/wangxin688/narvis/intend/platform"
)

type Dispatcher struct {
	Targets []string
	Config  factory.BaseSnmpConfig
}

func GetDeviceModel(sysObjId string) *dt.DeviceModel {
	privateEnterPriseId := strings.Split(strings.Split(sysObjId, ".1.3.6.1.4.1.")[1], ".")[0]
	enterprise := manufacturer.GetManufacturerByEnterpriseId(privateEnterPriseId)
	if enterprise == manufacturer.Unknown {
		return &dt.DeviceModel{
			Platform:     platform.Unknown,
			Manufacturer: manufacturer.Unknown,
			DeviceModel:  dt.UnknownDeviceModel,
		}
	}
	return GetDeviceModelFromManufacturer(enterprise, sysObjId)
}

func GetDeviceModelFromManufacturer(mf manufacturer.Manufacturer, sysObjId string) *dt.DeviceModel {
	switch mf {
	case manufacturer.Cisco:
		return s.CiscoDeviceModel(sysObjId)
	case manufacturer.Huawei:
		return s.HuaweiDeviceModel(sysObjId)
	case manufacturer.Aruba:
		return s.ArubaDeviceModel(sysObjId)
	case manufacturer.Arista:
		return s.AristaDeviceModel(sysObjId)
	case manufacturer.H3C:
		return s.H3CDeviceModel(sysObjId)
	case manufacturer.RuiJie:
		return s.RuiJieDeviceModel(sysObjId)
	case manufacturer.PaloAlto:
		return s.PaloAltoDeviceModel(sysObjId)
	case manufacturer.FortiNet:
		return s.FortiNetDeviceModel(sysObjId)
	case manufacturer.Netgear:
		return s.NetgearDeviceModel(sysObjId)
	case manufacturer.TPLink:
		return s.TPLinkDeviceModel(sysObjId)
	case manufacturer.Ruckus:
		return s.RuckusDeviceModel(sysObjId)
	case manufacturer.Juniper:
		return s.JuniperDeviceModel(sysObjId)
	case manufacturer.CheckPoint:
		return s.CheckPointDeviceModel(sysObjId)
	case manufacturer.Sangfor:
		return s.SangforDeviceModel(sysObjId)
	case manufacturer.A10:
		return s.A10DeviceModel(sysObjId)
	case manufacturer.F5:
		return s.F5DeviceModel(sysObjId)
	case manufacturer.Extreme:
		return s.ExtremeDeviceModel(sysObjId)
	case manufacturer.MikroTik:
		return s.MikroTikDeviceModel(sysObjId)
	}

	return &dt.DeviceModel{
		Platform:     platform.Unknown,
		Manufacturer: mf,
		DeviceModel:  dt.UnknownDeviceModel,
	}
}

func (d *Dispatcher) getFactory(platformType platform.Platform, snmpConfig factory.SnmpConfig) (factory.SnmpDriver, error) {
	var snmpDriver factory.SnmpDriver
	var err error

	switch platformType {
	case platform.CiscoIos:
		snmpDriver, err = driver.NewCiscoIosDriver(snmpConfig)
	case platform.CiscoIosXE:
		snmpDriver, err = driver.NewCiscoIosXEDriver(snmpConfig)
	case platform.CiscoIosXR:
		snmpDriver, err = driver.NewCiscoIosXRDriver(snmpConfig)
	case platform.CiscoNexusOS:
		snmpDriver, err = driver.NewCiscoNexusOSDriver(snmpConfig)
	case platform.Huawei:
		snmpDriver, err = driver.NewHuaweiDriver(snmpConfig)
	case platform.Aruba:
		snmpDriver, err = driver.NewArubaDriver(snmpConfig)
	case platform.ArubaOSSwitch:
		snmpDriver, err = driver.NewArubaOSSwitchDriver(snmpConfig)
	case platform.Arista:
		snmpDriver, err = driver.NewAristaDriver(snmpConfig)
	case platform.RuiJie:
		snmpDriver, err = driver.NewRuiJieDriver(snmpConfig)
	case platform.H3C:
		snmpDriver, err = driver.NewH3CDriver(snmpConfig)
	case platform.FortiNet:
		snmpDriver, err = driver.NewFortiNetDriver(snmpConfig)
	case platform.PaloAlto:
		snmpDriver, err = driver.NewPaloAltoDriver(snmpConfig)
	case platform.Juniper:
		snmpDriver, err = driver.NewJuniperDriver(snmpConfig)
	case platform.Netgear:
		snmpDriver, err = driver.NewNetgearDriver(snmpConfig)
	case platform.TPLink:
		snmpDriver, err = driver.NewTPLinkDriver(snmpConfig)
	case platform.Ruckus:
		snmpDriver, err = driver.NewRuckusDriver(snmpConfig)
	case platform.Sangfor:
		snmpDriver, err = driver.NewSangforDriver(snmpConfig)
	case platform.A10:
		snmpDriver, err = driver.NewA10Driver(snmpConfig)
	case platform.F5:
		snmpDriver, err = driver.NewF5Driver(snmpConfig)
	case platform.CheckPoint:
		snmpDriver, err = driver.NewCheckPointDriver(snmpConfig)
	case platform.ZTE:
		snmpDriver, err = driver.NewZTEDriver(snmpConfig)
	case platform.Extreme:
		snmpDriver, err = driver.NewExtremeDriver(snmpConfig)
	case platform.MikroTik:
		snmpDriver, err = driver.NewMikroTikDriver(snmpConfig)
	case platform.Unknown:
		snmpDriver, err = factory.NewSnmpDiscovery(snmpConfig)
	default:
		snmpDriver, err = factory.NewSnmpDiscovery(snmpConfig)
	}

	if err != nil {
		return nil, err
	}

	return snmpDriver, nil
}

func (d *Dispatcher) SnmpReachable(session *gosnmp.GoSNMP) bool {
	result, err := session.GetNext([]string{".1"})
	if err != nil {
		return false
	}
	return len(result.Variables) > 0
}

// linux need privilege for udp 
func (d *Dispatcher) IcmpReachable(address string) bool {
	pinger, err := ping.NewPinger(address)
	if err != nil {
		return false
	}
	pinger.Count = 2
	pinger.Interval = time.Duration(100) * time.Millisecond
	err = pinger.Run()
	if err != nil {
		return false
	}
	return pinger.Statistics().PacketsRecv > 0
}

func (d *Dispatcher) SshReachable(address string) bool {
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", address+":22", timeout)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func (d *Dispatcher) Session(config *factory.SnmpConfig) (*gosnmp.GoSNMP, error) {
	session, err := factory.NewSnmpSession(*config)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (d *Dispatcher) SysObjectID(session *gosnmp.GoSNMP) string {
	oid := []string{factory.SysObjectID}
	result, err := session.Get(oid)
	if err != nil {
		return ""
	}

	for _, variable := range result.Variables {
		if variable.Type == gosnmp.ObjectIdentifier {
			return fmt.Sprintf("%s", variable.Value)
		}
	}
	return ""
}

func (d *Dispatcher) dispatch(config factory.SnmpConfig) *factory.DispatchResponse {
	var response = &factory.DispatchResponse{}
	response.IpAddress = config.IpAddress
	session, err := d.Session(&config)
	if err != nil {
		response.SnmpReachable = false
	}
	icmp := d.IcmpReachable(config.IpAddress)
	ssh := d.SshReachable(config.IpAddress)
	response.IcmpReachable = icmp
	response.SshReachable = ssh
	response.SnmpReachable = d.SnmpReachable(session)

	sysObjectId := d.SysObjectID(session)
	if sysObjectId == "" {
		return response
	}
	response.SysObjectId = sysObjectId
	deviceType := GetDeviceModel(sysObjectId)
	driver, err := d.getFactory(deviceType.Platform, config)
	if err != nil {
		return response
	}
	discoveryResponse := driver.Discovery()
	response.DeviceModel = deviceType
	response.Data = discoveryResponse
	return response
}

func (d *Dispatcher) Dispatch() []*factory.DispatchResponse {
	var responses []*factory.DispatchResponse
	var wg sync.WaitGroup

	for _, target := range d.Targets {
		wg.Add(1)
		go func(target string) {
			defer wg.Done()

			targetResponse := d.dispatch(factory.SnmpConfig{
				IpAddress:      target,
				BaseSnmpConfig: d.Config,
			})

			responses = append(responses, targetResponse)
		}(target)
	}

	wg.Wait()

	return responses
}

func NewDispatcher(targets []string, config factory.BaseSnmpConfig) *Dispatcher {
	return &Dispatcher{
		Targets: targets,
		Config:  config,
	}
}
