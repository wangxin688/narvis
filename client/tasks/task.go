package tasks

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gosnmp/gosnmp"
	"github.com/wangxin688/narvis/client/config"
	"github.com/wangxin688/narvis/client/pkg/nettysnmp"
	"github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"
	"github.com/wangxin688/narvis/client/utils/helpers"
	"github.com/wangxin688/narvis/client/utils/logger"
	"github.com/wangxin688/narvis/intend/intendtask"
)

func scanDeviceBasicInfo(data []byte) ([]*intendtask.DeviceBasicInfoScanResponse, error) {
	results := make([]*intendtask.DeviceBasicInfoScanResponse, 0)
	task := &intendtask.BaseSnmpScanTask{}
	err := json.Unmarshal(data, task)
	if err != nil {
		logger.Logger.Error("[ScanDeviceBasicInfo]: Unmarshal err: ", err)
		return nil, err
	}
	targets, err := helpers.CIDRToIPStrings(task.Range)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[ScanDeviceBasicInfo]: received wrong ip range %s", task.Range), err)
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
			if !r.SnmpReachable || (r.Data != nil && len(r.Data.Errors) > 0) {
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

func scanDevice(data []byte) ([]*intendtask.DeviceScanResponse, error) {
	return nil, nil
}

func scanAp(data []byte) ([]*intendtask.ApScanResponse, error) {
	results := make([]*intendtask.ApScanResponse, 0)
	task := &intendtask.BaseSnmpTask{}
	err := json.Unmarshal(data, task)
	if err != nil {
		logger.Logger.Error("[ScanAp]: Unmarshal err: ", err)
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
					ManagementIp:    task.ManagementIp,
					SerialNumber:    StringToPtrString(ap.SerialNumber),
					GroupName:       StringToPtrString(ap.GroupName),
					SiteId:          task.SiteId,
					OrganizationId:  config.Settings.ORGANIZATION_ID,
					DeviceModel:     ap.DeviceModel,
					WlanACIpAddress: StringToPtrString(ap.WlanACIpAddress),
					MacAddress:      StringToPtrString(ap.MacAddress),
					Manufacturer:    string(r.DeviceModel.Manufacturer),
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
