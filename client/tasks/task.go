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
	"github.com/wangxin688/narvis/client/pkg/nettysnmp"
	"github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"
	"github.com/wangxin688/narvis/client/utils/helpers"
	"github.com/wangxin688/narvis/client/utils/logger"
	"github.com/wangxin688/narvis/client/utils/security"
	"github.com/wangxin688/narvis/intend/intendtask"
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
	targets, err := helpers.CIDRToIPStrings(task.Range)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[ScanDeviceBasicInfo]: received wrong ip range %s", task.Range), zap.Error(err))
		return nil, err
	}
	snmpConfig := factory.BaseSnmpConfig{
		Port:           task.SnmpConfig.Port,
		Version:        gosnmp.Version2c,
		Timeout:        task.SnmpConfig.Timeout,
		MaxRepetitions: int(task.SnmpConfig.MaxRepetitions),
		Community:      &task.SnmpConfig.Community,
	}
	dispatcher := nettysnmp.NewDispatcher(targets, snmpConfig)
	result := dispatcher.DispatchBasic()
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
	snmpConfig := factory.BaseSnmpConfig{
		Port:           task.SnmpConfig.Port,
		Version:        gosnmp.Version2c,
		Timeout:        task.SnmpConfig.Timeout,
		MaxRepetitions: int(task.SnmpConfig.MaxRepetitions),
		Community:      &task.SnmpConfig.Community,
	}
	dispatcher := nettysnmp.NewDispatcher([]string{task.ManagementIp}, snmpConfig)
	result := dispatcher.Dispatch()
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
	resp.Interfaces = lo.Map(r.Data.Interfaces, func(item *factory.DeviceInterface, _ int) *intendtask.DeviceInterface {
		return &intendtask.DeviceInterface{
			IfIndex:       item.IfIndex,
			IfName:        item.IfName,
			IfDescr:       item.IfDescr,
			IfType:        item.IfType,
			IfMtu:         item.IfMtu,
			IfSpeed:       item.IfSpeed,
			IfAdminStatus: item.IfAdminStatus,
			IfOperStatus:  item.IfOperStatus,
			IfLastChange:  item.IfLastChange,
			IfPhysAddr:    &item.IfPhysAddr,
			IfHighSpeed:   item.IfHighSpeed,
			IfIpAddress:   &item.IfIpAddress,
		}
	})
	resp.Vlans = lo.Map(r.Data.Vlans, func(item *factory.VlanItem, _ int) *intendtask.VlanItem {
		return &intendtask.VlanItem{
			VlanId:   item.VlanId,
			VlanName: item.VlanName,
			IfIndex:  item.IfIndex,
			Range:    item.Range,
			Gateway:  item.Gateway,
		}
	})
	resp.LldpNeighbors = lo.Map(r.Data.LldpNeighbors, func(item *factory.LldpNeighbor, _ int) *intendtask.LldpNeighbor {
		return &intendtask.LldpNeighbor{
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
	resp.Entities = lo.Map(r.Data.Entities, func(item *factory.Entity, _ int) *intendtask.Entity {
		return &intendtask.Entity{
			EntityPhysicalClass:       item.EntityPhysicalClass,
			EntityPhysicalDescr:       item.EntityPhysicalDescr,
			EntityPhysicalName:        item.EntityPhysicalName,
			EntityPhysicalSoftwareRev: item.EntityPhysicalSoftwareRev,
			EntityPhysicalSerialNum:   item.EntityPhysicalSerialNum,
		}
	})
	resp.Stacks = lo.Map(r.Data.Stacks, func(item *factory.Stack, _ int) *intendtask.Stack {
		return &intendtask.Stack{
			Id:         item.Id,
			Priority:   item.Priority,
			Role:       item.Role,
			MacAddress: item.MacAddress,
		}
	})
	resp.ArpTable = lo.Map(r.Data.ArpTable, func(item *factory.ArpItem, _ int) *intendtask.ArpItem {
		return &intendtask.ArpItem{
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
	snmpConfig := factory.BaseSnmpConfig{
		Port:           task.SnmpConfig.Port,
		Version:        gosnmp.Version2c,
		Timeout:        task.SnmpConfig.Timeout,
		MaxRepetitions: int(task.SnmpConfig.MaxRepetitions),
		Community:      &task.SnmpConfig.Community,
	}
	dispatcher := nettysnmp.NewDispatcher([]string{task.ManagementIp}, snmpConfig)
	result := dispatcher.DispatchApScan()
	if len(result) != 0 {
		for _, r := range result {
			if !r.SnmpReachable || len(r.Errors) > 0 || r.Data == nil {
				continue
			}
			for _, ap := range r.Data {
				results = append(results, &intendtask.ApScanResponse{
					Name:            ap.Name,
					ManagementIp:    ap.ManagementIp,
					SerialNumber:    StringToPtrString(ap.SerialNumber),
					GroupName:       StringToPtrString(ap.GroupName),
					SiteId:          task.SiteId,
					OrganizationId:  config.Settings.ORGANIZATION_ID,
					DeviceModel:     ap.DeviceModel,
					WlanACIpAddress: StringToPtrString(ap.WlanACIpAddress),
					MacAddress:      StringToPtrString(ap.MacAddress),
					Manufacturer:    string(r.DeviceModel.Manufacturer),
					OsVersion:       StringToPtrString(ap.OsVersion),
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
	token, err := security.ProxyToken(config.Settings.PROXY_ID, config.Settings.SECRET_KEY)
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
	result.HashValue = helpers.StringToMd5(configuration)
	return result

}

