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
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/pkg/rmq"
	"github.com/wangxin688/narvis/server/tools/errors"
	"go.uber.org/zap"
)

func GenerateTask(siteId, taskName, callback string) ([]*intendtask.BaseSnmpTask, error) {
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
			SiteId:       siteId,
			DeviceId:     deviceId,
			ManagementIp: deviceManagementIp[deviceId],
			Callback:     callback,
		})
	}
	return results, nil
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
		}
	}
	err := gen.TaskResult.CreateInBatches(newTasks, taskLen)
	if err != nil {
		core.Logger.Error("[CreateScanTask]: create task result to db failed", zap.Error(err))
	}
	return taskIds, nil
}
