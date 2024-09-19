package tasks

import (
	"encoding/json"
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
			if len(r.Data.Errors) > 0 {
				continue
			} else {
				results = append(results, &intendtask.DeviceBasicInfoScanResponse{
					Name:           r.Data.Hostname,
					ManagementIp:   r.IpAddress,
					Manufacturer:   string(r.DeviceModel.Manufacturer),
					Platform:       string(r.DeviceModel.Platform),
					DeviceModel:    string(r.DeviceModel.DeviceModel),
					Description:    r.Data.SysDescr,
					OrganizationId: config.Settings.ORGANIZATION_ID,
				})
			}
		}
	}
	return results, nil
}

func scanDevice(data []byte) ([]*intendtask.DeviceScanResponse, error) {
	return nil, nil
}

func scanAp(data []byte) ([]*intendtask.ApScanResponse, error) {
	return nil, nil
}

func scanMacAddressTable(data []byte) ([]*intendtask.MacAddressTableScanResponse, error) {
	return nil, nil
}
