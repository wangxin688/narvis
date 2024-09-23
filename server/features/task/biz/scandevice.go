package task_biz

import (
	"github.com/wangxin688/narvis/intend/intendtask"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/models"
	"go.uber.org/zap"
	"gorm.io/datatypes"
)

func UpdateScanDeviceBasicResult(taskId string, scanDevices []*intendtask.DeviceBasicInfoScanResponse) error {
	var status string
	var errors []string
	var updateValues models.Result

	for _, scanDevice := range scanDevices {
		if len(scanDevice.Errors) > 0 {
			errors = append(errors, scanDevice.Errors...)
		}
	}
	if len(errors) == 0 {
		status = "Success"
		updateValues = models.Result{
			Data:   nil,
			Errors: errors,
		}
	} else {
		status = "Failed"
		updateValues = models.Result{
			Data:   scanDevices,
			Errors: errors,
		}
	}
	_, err := gen.TaskResult.Where(gen.TaskResult.Id.Eq(taskId)).UpdateColumns(map[string]any{
		"Status": status,
		"result": datatypes.NewJSONType(updateValues),
	})
	if err != nil {
		core.Logger.Error("[updateTaskResult]: update scan device basic info task failed", zap.Error(err))
	}
	return nil
}
