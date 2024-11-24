package nettysnmp

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
	"github.com/wangxin688/narvis/intend/logger"
	"github.com/wangxin688/narvis/intend/model/manufacturer"
	"github.com/wangxin688/narvis/intend/platform"
	"go.uber.org/zap"
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
	case platform.F5:
		snmpDriver, err = driver.NewF5Driver(snmpConfig)
	case platform.CheckPoint:
		snmpDriver, err = driver.NewCheckPointDriver(snmpConfig)
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
	result, err := session.Get([]string{factory.SysName})
	if err != nil || result == nil {
		logger.Logger.Info("[dispatcher]: SnmpReachable failed", zap.Error(err), zap.String("target", session.Target))
		return false
	}
	return len(result.Variables) > 0
}

// linux need privilege for udp
func (d *Dispatcher) IcmpReachable(address string) bool {
	pinger, err := ping.NewPinger(address)
	if err != nil {
		logger.Logger.Info("[dispatcher]: IcmpReachable failed", zap.String("target", address), zap.Error(err))
		return false
	}
	pinger.Count = 2
	pinger.Interval = time.Duration(100) * time.Millisecond
	pinger.Timeout = time.Second
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
		logger.Logger.Info("[dispatcher]: SshReachable failed", zap.String("target", address), zap.Error(err))
		return false
	}
	if conn == nil {
		logger.Logger.Info("[dispatcher]: SshReachable failed", zap.String("target", address), zap.String("reason", "connection is nil"))
		return false
	}
	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()
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
	if err != nil || session == nil {
		response.SnmpReachable = false
	} else {
		response.SnmpReachable = d.SnmpReachable(session)
	}
	icmp := d.IcmpReachable(config.IpAddress)
	ssh := d.SshReachable(config.IpAddress)
	response.IcmpReachable = icmp
	response.SshReachable = ssh
	if !response.SnmpReachable {
		return response
	}

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
	ch := make(chan struct{}, 100)
	for _, target := range d.Targets {
		ch <- struct{}{}
		wg.Add(1)
		go func(target string) {
			defer wg.Done()

			targetResponse := d.dispatch(factory.SnmpConfig{
				IpAddress:      target,
				BaseSnmpConfig: d.Config,
			})

			responses = append(responses, targetResponse)
			<-ch
		}(target)
	}

	wg.Wait()

	return responses
}

func (d *Dispatcher) dispatchBasic(config factory.SnmpConfig) *factory.DispatchBasicResponse {
	var response = &factory.DispatchBasicResponse{}
	response.IpAddress = config.IpAddress
	session, err := d.Session(&config)
	if err != nil || session == nil {
		response.SnmpReachable = false
	} else {
		response.SnmpReachable = d.SnmpReachable(session)
	}
	icmp := d.IcmpReachable(config.IpAddress)
	ssh := d.SshReachable(config.IpAddress)
	response.IcmpReachable = icmp
	response.SshReachable = ssh
	if !response.SnmpReachable {
		return response
	}
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
	discoveryResponse := driver.DiscoveryBasicInfo()
	response.DeviceModel = deviceType
	response.Data = discoveryResponse
	return response
}

func (d *Dispatcher) DispatchBasic() []*factory.DispatchBasicResponse {
	var responses []*factory.DispatchBasicResponse
	var wg sync.WaitGroup
	ch := make(chan struct{}, 100)
	for _, target := range d.Targets {
		ch <- struct{}{}
		wg.Add(1)
		go func(target string) {
			defer wg.Done()

			targetResponse := d.dispatchBasic(factory.SnmpConfig{
				IpAddress:      target,
				BaseSnmpConfig: d.Config,
			})

			responses = append(responses, targetResponse)
			<-ch
		}(target)
	}

	wg.Wait()
	return responses
}

func (d *Dispatcher) dispatchApScan(config factory.SnmpConfig) *factory.DispatchApScanResponse {
	var response = &factory.DispatchApScanResponse{}
	response.IpAddress = config.IpAddress
	session, err := d.Session(&config)
	if err != nil || session == nil {
		response.SnmpReachable = false
	} else {
		response.SnmpReachable = d.SnmpReachable(session)
	}
	icmp := d.IcmpReachable(config.IpAddress)
	ssh := d.SshReachable(config.IpAddress)
	response.IcmpReachable = icmp
	response.SshReachable = ssh
	if !response.SnmpReachable {
		return response
	}
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
	discoveryResponse, errors := driver.APs()
	response.DeviceModel = deviceType
	if len(errors) > 0 {
		response.Errors = errors
		return response
	}

	response.Data = discoveryResponse
	return response
}

func (d *Dispatcher) DispatchApScan() []*factory.DispatchApScanResponse {
	var responses []*factory.DispatchApScanResponse
	var wg sync.WaitGroup
	ch := make(chan struct{}, 100)
	for _, target := range d.Targets {
		ch <- struct{}{}
		wg.Add(1)
		go func(target string) {
			defer wg.Done()

			targetResponse := d.dispatchApScan(factory.SnmpConfig{
				IpAddress:      target,
				BaseSnmpConfig: d.Config,
			})

			responses = append(responses, targetResponse)
			<-ch
		}(target)
	}

	wg.Wait()
	return responses
}

func (d *Dispatcher) dispatchWlanUser(config factory.SnmpConfig) *factory.WlanUserResponse {
	var response = &factory.WlanUserResponse{}
	session, err := d.Session(&config)
	if err != nil || session == nil {
		response.SnmpReachable = false
	} else {
		response.SnmpReachable = d.SnmpReachable(session)
	}
	if !response.SnmpReachable {
		return response
	}
	sysObjectId := d.SysObjectID(session)
	if sysObjectId == "" {
		return response
	}
	deviceType := GetDeviceModel(sysObjectId)
	driver, err := d.getFactory(deviceType.Platform, config)
	if err != nil {
		return response
	}
	wlanUsers := driver.WlanUsers()
	return wlanUsers
}

func (d *Dispatcher) DispatchWlanUser() []*factory.WlanUserResponse {
	var wg sync.WaitGroup
	ch := make(chan struct{}, 100)
	var responses []*factory.WlanUserResponse
	for _, target := range d.Targets {
		ch <- struct{}{}
		wg.Add(1)
		go func(target string) {
			defer wg.Done()
			response := d.dispatchWlanUser(factory.SnmpConfig{
				IpAddress:      target,
				BaseSnmpConfig: d.Config,
			})
			responses = append(responses, response)
			<-ch
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
