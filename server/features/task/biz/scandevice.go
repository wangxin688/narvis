package task_biz

import (
	"github.com/wangxin688/narvis/intend/intendtask"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/models"
	"gorm.io/datatypes"
)

func UpdateScanDeviceBasicResult(taskId string, scanDevices []*intendtask.DeviceBasicInfoScanResponse) error {
	result := models.Result{}
	status := "Success"
	if scanDevices == nil {
		status = "Failed"
		result.Errors = append(result.Errors, "scanDevices receive empty response from proxy")
	} else {
		for _, scanDevice := range scanDevices {
			result.Errors = append(result.Errors, scanDevice.Errors...)
		}
	}
	if len(result.Errors) > 0 {
		status = "Failed"
		result.Data = scanDevices
	}

	_, err := gen.TaskResult.Where(gen.TaskResult.Id.Eq(taskId)).UpdateColumns(map[string]any{
		"Status": status,
		"result": datatypes.NewJSONType(result),
	})
	if err != nil {
		return err
	}
	return nil
}

func UpdateScanDeviceResult(taskId string, device *intendtask.DeviceScanResponse) error {
	result := models.Result{}
	status := "Success"
	if device == nil {
		status = "Failed"
		result.Errors = append(result.Errors, "device receive empty response from proxy")
	} else {
		result.Errors = append(result.Errors, device.Errors...)
	}
	if len(result.Errors) > 0 {
		status = "Failed"
		result.Data = device
	}

	_, err := gen.TaskResult.Where(gen.TaskResult.Id.Eq(taskId)).UpdateColumns(map[string]any{
		"Status": status,
		"result": datatypes.NewJSONType(result),
	})
	if err != nil {
		return err
	}
	return nil
}
