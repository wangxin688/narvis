package task_biz

import (
	"github.com/wangxin688/narvis/intend/intendtask"
	"github.com/wangxin688/narvis/intend/logger"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/models"
	"go.uber.org/zap"
	"gorm.io/datatypes"
)

func UpdateConfigBackupResult(taskId string, configBackUp *intendtask.ConfigurationBackupTaskResult) error {
	Status := "Success"
	result := models.Result{}
	if configBackUp.Error != "" {
		Status = "Failed"
		result.Errors = append(result.Errors, configBackUp.Error)
	}
	_, err := gen.TaskResult.Where(gen.TaskResult.Id.Eq(taskId)).UpdateColumns(
		map[string]any{
			"Status": Status,
			"result": datatypes.NewJSONType(result),
		},
	)
	if err != nil {
		logger.Logger.Error("[updateTaskResult]: update config backup task failed", zap.Error(err))
		return err
	}
	return nil
}
