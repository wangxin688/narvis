package infra_tasks

import (
	"github.com/google/uuid"
	"github.com/wangxin688/narvis/intend/intendtask"
	infra_biz "github.com/wangxin688/narvis/server/features/infra/biz"
	infra_utils "github.com/wangxin688/narvis/server/features/infra/utils"
)

func generateTask(siteId, taskName, callback string) ([]*intendtask.BaseSnmpTask, error) {
	results := make([]*intendtask.BaseSnmpTask, 0)

	devices, err := infra_biz.NewDeviceService().GetActiveDevices(siteId)
	if err != nil {
		return results, err
	}
	if len(devices) == 0 {
		return results, nil
	}
	deviceIds := infra_utils.DevicesToIds(devices)
	deviceManagementIp := infra_utils.DeviceManagementIpMap(devices)
	snmpConfigs, err := infra_biz.NewSnmpCredentialService().GetCredentialByDeviceIds(deviceIds)
	if err != nil {
		return results, err
	}
	for deviceId, snmpConfig := range snmpConfigs {
		results = append(results, &intendtask.BaseSnmpTask{
			TaskId:   uuid.New().String(),
			TaskName: taskName,
			SnmpConfig: &intendtask.SnmpV2Credential{
				Community:      snmpConfig.Community,
				Port:           snmpConfig.Port,
				Timeout:        snmpConfig.Timeout,
				MaxRepetitions: snmpConfig.MaxRepetitions,
			},
			DeviceId:     deviceId,
			ManagementIp: deviceManagementIp[deviceId],
			Callback:     callback,
		})
	}
	return results, nil
}

// use ip range to scan devices
func generateScanDeviceTask(ipRange, community string, port uint16) *intendtask.BaseSnmpScanTask {
	return &intendtask.BaseSnmpScanTask{
		TaskId:   uuid.New().String(),
		TaskName: intendtask.ScanDeviceBasicInfo,
		Range:    ipRange,
		SnmpConfig: &intendtask.SnmpV2Credential{
			Community:      community,
			Port:           port,
			Timeout:        10,
			MaxRepetitions: 50,
		},
		Callback: intendtask.ScanDeviceBasicInfoCallback,
	}
}
