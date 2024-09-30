package task_biz

import (
	"github.com/wangxin688/narvis/intend/intendtask"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/models"
	"go.uber.org/zap"
	"gorm.io/datatypes"
)

func UpdateApScanResult(taskId string, scanAp []*intendtask.ApScanResponse) error {
	Status := "Success"
	result := models.Result{}
	if scanAp == nil {
		Status = "Failed"
		result.Errors = append(result.Errors, "scanAp receive empty response from proxy")
	}
	_, err := gen.TaskResult.Where(gen.TaskResult.Id.Eq(taskId)).UpdateColumns(
		map[string]any{
			"Status": Status,
			"result": datatypes.NewJSONType(result),
		},
	)
	if err != nil {
		core.Logger.Error("[updateTaskResult]: update ap scan task failed", zap.Error(err))
		return err
	}
	return nil
}
