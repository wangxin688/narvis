package infra_tasks

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/wangxin688/narvis/intend/intendtask"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/dal/gen"
	infra_biz "github.com/wangxin688/narvis/server/features/infra/biz"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	infra_utils "github.com/wangxin688/narvis/server/features/infra/utils"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/pkg/rmq"
	"github.com/wangxin688/narvis/server/tools/errors"
	"go.uber.org/zap"
)

func GenerateSNMPTask(siteId, taskName, callback string) ([]string, error) {
	taskIds := make([]string, 0)
	orgId := global.OrganizationId.Get()
	devices, err := getDevicesByTaskName(siteId, taskName)
	if err != nil {
		return taskIds, err
	}
	if len(devices) == 0 {
		core.Logger.Info("[taskGeneration]: no devices found", zap.String("taskName", taskName), zap.String("siteId", siteId))
		return taskIds, errors.NewError(errors.CodeNoDevicesFound, errors.MsgNoDevicesFound)
	}
	deviceIds := infra_utils.DevicesToIds(devices)
	deviceManagementIp := infra_utils.DeviceManagementIpMap(devices)
	snmpConfigs, err := infra_biz.NewSnmpCredentialService().GetCredentialByDeviceIds(deviceIds)
	if err != nil {
		return taskIds, err
	}
	newTasks := make([]*models.TaskResult, 0)
	for deviceId, snmpConfig := range snmpConfigs {
		taskId := uuid.New().String()
		task := &intendtask.BaseSnmpTask{
			TaskId:   taskId,
			TaskName: taskName,
			SnmpConfig: &intendtask.SnmpV2Credential{
				Community:      snmpConfig.Community,
				Port:           snmpConfig.Port,
				Timeout:        snmpConfig.Timeout,
				MaxRepetitions: snmpConfig.MaxRepetitions,
			},
			SiteId:       siteId,
			DeviceId:     deviceId,
			ManagementIp: deviceManagementIp[deviceId],
			Callback:     callback,
		}
		taskBytes, err := json.Marshal(task)
		if err != nil {
			return taskIds, err
		}
		err = rmq.PublishProxyMessage(taskBytes, orgId)
		if err != nil {
			return taskIds, err
		}
		taskIds = append(taskIds, taskId)
		newTasks = append(newTasks, &models.TaskResult{
			BaseDbModel: models.BaseDbModel{
				Id: taskId,
			},
			Name:           taskName,
			Status:         "InProgress",
			OrganizationId: orgId,
			TaskRaw:        string(taskBytes),
		})
	}
	err = gen.TaskResult.CreateInBatches(newTasks, len(newTasks))
	if err != nil {
		return taskIds, err
	}
	return taskIds, nil
}

func getDevicesByTaskName(siteId, taskName string) ([]*models.Device, error) {
	switch taskName {
	case intendtask.ScanAp:
		return infra_biz.NewDeviceService().GetActiveWlanAC(siteId)
	case intendtask.WlanUser:
		return infra_biz.NewDeviceService().GetActiveWlanAC(siteId)
	case intendtask.ScanDevice:
		return infra_biz.NewDeviceService().GetActiveDevices(siteId)
	case intendtask.ScanMacAddressTable:
		return infra_biz.NewDeviceService().GetActiveDevices(siteId)
	case intendtask.ScanIPAM:
		return infra_biz.NewDeviceService().GetCoreSwitch(siteId)
	case intendtask.ConfigurationBackup:
		return infra_biz.NewDeviceService().GetActiveDevices(siteId)
	default:
		return nil, errors.NewError(errors.CodeTaskNameInvalid, errors.MsgTaskNameInvalid)
	}
}

// use ip range to scan devices
func generateScanDeviceTask(ipRange, community string, port uint16, timeout uint8, maxRepetitions uint8) *intendtask.BaseSnmpScanTask {
	return &intendtask.BaseSnmpScanTask{
		TaskId:   uuid.New().String(),
		TaskName: intendtask.ScanDeviceBasicInfo,
		Range:    ipRange,
		SnmpConfig: &intendtask.SnmpV2Credential{
			Community:      community,
			Port:           port,
			Timeout:        timeout,
			MaxRepetitions: maxRepetitions,
		},
		Callback: intendtask.ScanDeviceBasicInfoCallback,
	}
}

func CreateScanTask(sd *schemas.ScanDeviceCreate, orgId string) ([]string, error) {
	taskLen := len(sd.Range)
	taskIds := make([]string, taskLen)
	if len(sd.Range) == 0 {
		return taskIds, errors.NewError(errors.CodeIpRangeNotProvided, errors.MsgIpRangeNotProvided)
	}
	newTasks := make([]*models.TaskResult, taskLen)
	for index, ipRange := range sd.Range {
		task := generateScanDeviceTask(
			ipRange, *sd.Community, *sd.Port, *sd.Timeout, *sd.MaxRepetitions,
		)
		taskIds[index] = task.TaskId
		taskByte, err := json.Marshal(task)
		if err != nil {
			core.Logger.Error("[CreateScanTask]: marshal task failed", zap.Error(err))
			continue
		}
		err = rmq.PublishProxyMessage(taskByte, orgId)
		if err != nil {
			core.Logger.Error("[CreateScanTask]: publish task failed", zap.Error(err))
			continue
		}
		newTasks[index] = &models.TaskResult{
			BaseDbModel: models.BaseDbModel{
				Id: task.TaskId,
			},
			Name:           intendtask.ScanDeviceBasicInfo,
			Status:         "InProgress",
			OrganizationId: orgId,
			TaskRaw: 	   string(taskByte),
		}
	}
	err := gen.TaskResult.CreateInBatches(newTasks, taskLen)
	if err != nil {
		core.Logger.Error("[CreateScanTask]: create task result to db failed", zap.Error(err))
	}
	return taskIds, nil
}

func ConfigBackUpTask(siteId, taskName, callback string) ([]string, error) {
	taskIds := make([]string, 0)
	orgId := global.OrganizationId.Get()
	devices, err := getDevicesByTaskName(siteId, taskName)
	if err != nil {
		return taskIds, err
	}
	deviceIds := infra_utils.DevicesToIds(devices)
	devicePlatforms := infra_utils.DevicePlatforms(devices)
	deviceManagementIp := infra_utils.DeviceManagementIpMap(devices)
	sshCredential, err := infra_biz.NewCliCredentialService().GetCredentialByDeviceIds(deviceIds)
	if err != nil {
		return taskIds, err
	}
	newTasks := make([]*models.TaskResult, 0)
	for deviceId, sshConfig := range sshCredential {
		taskId := uuid.New().String()
		task := &intendtask.ConfigurationBackupTask{
			TaskId:       taskId,
			TaskName:     taskName,
			DeviceId:     deviceId,
			ManagementIp: deviceManagementIp[deviceId],
			Username:     sshConfig.Username,
			Password:     sshConfig.Password,
			Port:         sshConfig.Port,
			Platform:     devicePlatforms[deviceId],
			Callback:     callback,
		}
		taskBytes, err := json.Marshal(task)
		if err != nil {
			return taskIds, err
		}
		err = rmq.PublishProxyMessage(taskBytes, orgId)
		if err != nil {
			return taskIds, err
		}
		newTasks = append(newTasks, &models.TaskResult{
			BaseDbModel: models.BaseDbModel{
				Id: taskId,
			},
			Name:           taskName,
			Status:         "InProgress",
			OrganizationId: orgId,
			TaskRaw:		string(taskBytes),
		})
		taskIds = append(taskIds, taskId)
	}
	err = gen.TaskResult.CreateInBatches(newTasks, len(newTasks))
	if err != nil {
		return taskIds, err
	}
	return taskIds, nil

}
