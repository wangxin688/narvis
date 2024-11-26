package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/gosnmp/gosnmp"
	"github.com/samber/lo"
	"github.com/wangxin688/narvis/client/config"
	"github.com/wangxin688/narvis/client/pkg/gossh"
	"github.com/wangxin688/narvis/client/pkg/gowebssh"
	"github.com/wangxin688/narvis/intend/helpers/network"
	"github.com/wangxin688/narvis/intend/helpers/processor"
	"github.com/wangxin688/narvis/intend/helpers/security"
	"github.com/wangxin688/narvis/intend/intendtask"
	"github.com/wangxin688/narvis/intend/logger"
	intend_device "github.com/wangxin688/narvis/intend/model/device"
	"github.com/wangxin688/narvis/intend/model/snmp"
	"github.com/wangxin688/narvis/intend/netdisco"
	"go.uber.org/zap"
)

func scanDeviceBasicInfo(data []byte) ([]*intendtask.DeviceBasicInfoScanResponse, error) {
	results := make([]*intendtask.DeviceBasicInfoScanResponse, 0)
	task := &intendtask.BaseSnmpScanTask{}
	err := json.Unmarshal(data, task)
	if err != nil {
		logger.Logger.Error("[ScanDeviceBasicInfo]: Unmarshal err: ", zap.Error(err))
		return nil, err
	}
	targets, err := network.CIDR2IpStrings(task.Range)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[ScanDeviceBasicInfo]: received wrong ip range %s", task.Range), zap.Error(err))
		return nil, err
	}
	snmpConfigs := make([]*snmp.SnmpConfig, 0)
	for _, target := range targets {
		snmpConfigs = append(snmpConfigs, &snmp.SnmpConfig{
			IpAddress:      target,
			Port:           task.SnmpConfig.Port,
			Timeout:        task.SnmpConfig.Timeout,
			Community:      &task.SnmpConfig.Community,
			Version:        gosnmp.Version2c,
			MaxRepetitions: int(task.SnmpConfig.MaxRepetitions),
		})
	}
	result := netdisco.NetworkDeviceDiscovery(snmpConfigs)
	if len(result) != 0 {
		for _, r := range result {
			if !r.SnmpReachable || r.Data == nil || len(r.Data.Errors) != 0 {
				continue
			}
			results = append(results, &intendtask.DeviceBasicInfoScanResponse{
				Name:           r.Data.Hostname,
				ManagementIp:   r.IpAddress,
				ChassisId:      r.Data.ChassisId,
				Manufacturer:   string(r.DeviceModel.Manufacturer),
				Platform:       string(r.DeviceModel.Platform),
				DeviceModel:    string(r.DeviceModel.DeviceModel),
				Description:    r.Data.SysDescr,
				OrganizationId: config.Settings.ORGANIZATION_ID,
				Errors:         r.Data.Errors,
			})

		}
	}
	return results, nil
}

func scanDevice(data []byte) (*intendtask.DeviceScanResponse, error) {
	task := &intendtask.BaseSnmpTask{}
	err := json.Unmarshal(data, task)
	if err != nil {
		logger.Logger.Error("[ScanDevice]: Unmarshal err: ", zap.Error(err))
		return nil, err
	}
	snmpConfig := snmp.SnmpConfig{
		IpAddress:      task.ManagementIp,
		Port:           task.SnmpConfig.Port,
		Version:        gosnmp.Version2c,
		Timeout:        task.SnmpConfig.Timeout,
		MaxRepetitions: int(task.SnmpConfig.MaxRepetitions),
		Community:      &task.SnmpConfig.Community,
	}
	result := netdisco.EnrichDevice([]*snmp.SnmpConfig{&snmpConfig})
	if len(result) != 1 {
		return nil, errors.New("no device found")
	}
	r := result[0]
	resp := &intendtask.DeviceScanResponse{
		SiteId:         task.SiteId,
		DeviceId:       task.DeviceId,
		OrganizationId: config.Settings.ORGANIZATION_ID,
		ManagementIp:   task.ManagementIp,
		SshReachable:   r.SshReachable,
		IcmpReachable:  r.IcmpReachable,
		SnmpReachable:  r.SnmpReachable,
	}
	if !r.SnmpReachable {
		resp.Errors = []string{"device snmp not reachable"}
		return resp, nil
	}
	if r.Data == nil || len(r.Data.Errors) > 0 {
		return &intendtask.DeviceScanResponse{
			SiteId:         task.SiteId,
			DeviceId:       task.DeviceId,
			OrganizationId: config.Settings.ORGANIZATION_ID,
			ManagementIp:   task.ManagementIp,
			Errors:         r.Data.Errors,
			SnmpReachable:  r.SnmpReachable,
			SshReachable:   r.SshReachable,
			IcmpReachable:  r.IcmpReachable,
		}, nil
	}
	resp.Manufacturer = string(r.DeviceModel.Manufacturer)
	resp.Platform = string(r.DeviceModel.Platform)
	resp.DeviceModel = string(r.DeviceModel.DeviceModel)
	resp.Description = r.Data.SysDescr
	resp.ChassisId = &r.Data.ChassisId
	resp.Name = r.Data.Hostname
	resp.Interfaces = lo.Map(r.Data.Interfaces, func(item *intend_device.DeviceInterface, _ int) *intend_device.DeviceInterface {
		return &intend_device.DeviceInterface{
			IfIndex:       item.IfIndex,
			IfName:        item.IfName,
			IfDescr:       item.IfDescr,
			IfType:        item.IfType,
			IfMtu:         item.IfMtu,
			IfSpeed:       item.IfSpeed,
			IfAdminStatus: item.IfAdminStatus,
			IfOperStatus:  item.IfOperStatus,
			IfLastChange:  item.IfLastChange,
			IfPhysAddr:    item.IfPhysAddr,
			IfHighSpeed:   item.IfHighSpeed,
			IfIpAddress:   item.IfIpAddress,
		}
	})
	resp.Vlans = lo.Map(r.Data.Vlans, func(item *intend_device.VlanItem, _ int) *intend_device.VlanItem {
		return &intend_device.VlanItem{
			VlanId:   item.VlanId,
			VlanName: item.VlanName,
			IfIndex:  item.IfIndex,
			Network:  item.Network,
			Gateway:  item.Gateway,
		}
	})
	resp.LldpNeighbors = lo.Map(r.Data.LldpNeighbors, func(item *intend_device.LldpNeighbor, _ int) *intend_device.LldpNeighbor {
		return &intend_device.LldpNeighbor{
			LocalChassisId:  item.LocalChassisId,
			LocalHostname:   item.LocalHostname,
			LocalIfName:     item.LocalIfName,
			LocalIfDescr:    item.LocalIfDescr,
			RemoteChassisId: item.RemoteChassisId,
			RemoteHostname:  item.RemoteHostname,
			RemoteIfName:    item.RemoteIfName,
			RemoteIfDescr:   item.RemoteIfDescr,
		}
	})
	resp.Entities = lo.Map(r.Data.Entities, func(item *intend_device.Entity, _ int) *intend_device.Entity {
		return &intend_device.Entity{
			EntityPhysicalClass:       item.EntityPhysicalClass,
			EntityPhysicalDescr:       item.EntityPhysicalDescr,
			EntityPhysicalName:        item.EntityPhysicalName,
			EntityPhysicalSoftwareRev: item.EntityPhysicalSoftwareRev,
			EntityPhysicalSerialNum:   item.EntityPhysicalSerialNum,
		}
	})
	resp.Stacks = lo.Map(r.Data.Stacks, func(item *intend_device.Stack, _ int) *intend_device.Stack {
		return &intend_device.Stack{
			Id:         item.Id,
			Priority:   item.Priority,
			Role:       item.Role,
			MacAddress: item.MacAddress,
		}
	})
	resp.ArpTable = lo.Map(r.Data.ArpTable, func(item *intend_device.ArpItem, _ int) *intend_device.ArpItem {
		return &intend_device.ArpItem{
			IpAddress:  item.IpAddress,
			MacAddress: item.MacAddress,
			Type:       item.Type,
			IfIndex:    item.IfIndex,
			VlanId:     item.VlanId,
			Range:      item.Range,
		}
	})
	return resp, nil
}

func scanAp(data []byte) ([]*intendtask.ApScanResponse, error) {
	results := make([]*intendtask.ApScanResponse, 0)
	task := &intendtask.BaseSnmpTask{}
	err := json.Unmarshal(data, task)
	if err != nil {
		logger.Logger.Error("[ScanAp]: Unmarshal err: ", zap.Error(err))
		return nil, err
	}
	snmpConfig := snmp.SnmpConfig{
		IpAddress:      task.ManagementIp,
		Port:           task.SnmpConfig.Port,
		Version:        gosnmp.Version2c,
		Timeout:        task.SnmpConfig.Timeout,
		MaxRepetitions: int(task.SnmpConfig.MaxRepetitions),
		Community:      &task.SnmpConfig.Community,
	}
	result := netdisco.WlanApDiscovery([]*snmp.SnmpConfig{&snmpConfig})
	if len(result) != 0 {
		for _, r := range result {
			if !r.SnmpReachable || len(r.Errors) > 0 || r.Data == nil {
				continue
			}
			for _, ap := range r.Data {
				results = append(results, &intendtask.ApScanResponse{
					Name:            ap.Name,
					ManagementIp:    ap.ManagementIp,
					SerialNumber:    processor.StringToPtrString(ap.SerialNumber),
					GroupName:       ap.GroupName,
					SiteId:          task.SiteId,
					OrganizationId:  config.Settings.ORGANIZATION_ID,
					DeviceModel:     ap.DeviceModel,
					WlanACIpAddress: ap.WlanACIpAddress,
					MacAddress:      ap.MacAddress,
					Manufacturer:    string(r.DeviceModel.Manufacturer),
					OsVersion:       ap.OsVersion,
				})
			}
		}
	}
	if len(results) == 0 {
		return results, errors.New("no ap found")
	}
	return results, nil

}
func scanMacAddressTable(data []byte) ([]*intendtask.MacAddressTableScanResponse, error) {
	return nil, nil
}

// webSSHTask starts a webssh session with the target device and
// relays the input and output between the websocket and the ssh connection.
func webSSHTask(data []byte) error {
	task := &intendtask.WebSSHTask{}
	err := json.Unmarshal(data, task)
	if err != nil {
		logger.Logger.Error("[webSSHTask]: Unmarshal err: ", zap.Error(err))
		return err
	}
	// Get the token from the proxy server
	token, err := security.GenerateAgentToken(config.Settings.PROXY_ID, config.Settings.SECRET_KEY, config.Settings.SECRET_KEY)
	if err != nil {
		logger.Logger.Error("[webSSHTask]: failed to get token", zap.Error(err))
		return err
	}
	sessionId := task.SessionId
	// Dial to the websocket server
	wsConn, _, err := websocket.DefaultDialer.Dial(
		config.Settings.WebSocketUrl()+intendtask.WebSocketCbUrl+"/"+sessionId, http.Header{"Authorization": {"Bearer " + token}})
	if err != nil {
		logger.Logger.Error("[webSSHTask]: failed to dial to server", zap.Error(err))
		return err
	}
	logger.Logger.Info(fmt.Sprintf("[webSSHTask]: dial to server success with sessionId: %s", sessionId))
	defer wsConn.Close()
	// Set the read deadline to 60 seconds
	wsConn.SetReadDeadline(time.Now().Add(60 * time.Second))
	terminal := gowebssh.NewTerminal(wsConn, gowebssh.Options{
		Addr:     task.ManagementIP,
		Port:     task.Port,
		User:     task.Username,
		Password: task.Password,
		Rows:     task.Rows,
		Cols:     task.Cols,
	})

	// Start the session
	terminal.Run()
	return nil
}

func configurationBackupTask(data []byte) *intendtask.ConfigurationBackupTaskResult {
	task := &intendtask.ConfigurationBackupTask{}
	err := json.Unmarshal(data, task)
	result := &intendtask.ConfigurationBackupTaskResult{
		DeviceId: task.DeviceId,
	}
	if err != nil {
		logger.Logger.Warn("[ConfigurationBackupTask]: Unmarshal err: ", zap.Error(err))
		result.Error = err.Error()
		return result
	}
	connectionInfo := gossh.ConnectionInfo{
		ManagementIp: task.ManagementIp,
		Port:         int(task.Port),
		Platform:     task.Platform,
		Username:     task.Username,
		Password:     task.Password,
		Timeout:      15,
	}
	sshConn, err := gossh.NewConnection(&connectionInfo)
	if err != nil {
		logger.Logger.Warn("[ConfigurationBackupTask]: failed to create ssh connection", zap.String("managementIp", task.ManagementIp), zap.Error(err))
		result.Error = err.Error()
		return result
	}
	configuration, err := sshConn.ShowRunningConfig()
	if err != nil {
		logger.Logger.Warn("[ConfigurationBackupTask]: failed to get running config", zap.String("managementIp", task.ManagementIp), zap.Error(err))
		result.Error = err.Error()
		return result
	}
	result.Configuration = configuration
	result.BackupTime = time.Now().UTC().String()
	result.HashValue = processor.String2Md5(configuration)
	return result

}

func wlanUserTask(data []byte) *intendtask.WlanUserTaskResult {
	task := &intendtask.BaseSnmpTask{}
	err := json.Unmarshal(data, task)
	result := &intendtask.WlanUserTaskResult{
		DeviceId:       task.DeviceId,
		SiteId:         task.SiteId,
		OrganizationId: config.Settings.ORGANIZATION_ID,
	}
	if err != nil {
		logger.Logger.Error("[WlanUserTask]: Unmarshal err: ", zap.Error(err))
		result.Errors = []string{fmt.Sprintf("failed to unmarshal task, error: %s", err.Error())}
		return result
	}
	snmpConfig := snmp.SnmpConfig{
		Port:           task.SnmpConfig.Port,
		Version:        gosnmp.Version2c,
		Timeout:        task.SnmpConfig.Timeout,
		MaxRepetitions: int(task.SnmpConfig.MaxRepetitions),
		Community:      &task.SnmpConfig.Community,
	}
	driver, _, err := netdisco.NewNetDisco(&snmpConfig).Driver()
	if err != nil || driver == nil {
		result.Errors = []string{fmt.Sprintf("not support for target device %s", task.ManagementIp)}
	}
	response, errors := driver.WlanUsers()
	if len(errors) > 0 {
		result.Errors = errors
		return result
	}
	result.WlanUsers = response
	return result
}
