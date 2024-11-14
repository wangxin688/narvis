package task_biz

import (
	"github.com/wangxin688/narvis/intend/intendtask"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/models"
	"go.uber.org/zap"
	"gorm.io/datatypes"
)

func UpdateWlanUserResult(taskId string, wlanUser *intendtask.WlanUserTaskResult) error {
	Status := "Success"
	result := models.Result{}
	if wlanUser.Errors != nil {
		Status = "Failed"
		result.Errors = wlanUser.Errors
	}
	_, err := gen.TaskResult.Where(gen.TaskResult.Id.Eq(taskId)).UpdateColumns(
		map[string]any{
			"Status": Status,
			"result": datatypes.NewJSONType(result),
		},
	)
	if err != nil {
		core.Logger.Error("[updateTaskResult]: update wlan user task failed", zap.Error(err))
		return err
	}
	return nil
}
