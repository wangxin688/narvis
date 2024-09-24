package task_biz

import (
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/dal/gen"
	"go.uber.org/zap"
)

func UpdateApScanResult(taskId string) error {
	_, err := gen.TaskResult.Where(gen.TaskResult.Id.Eq(taskId)).UpdateColumn(gen.TaskResult.Status, "Success")
	if err != nil {
		core.Logger.Error("[updateTaskResult]: update ap scan task failed", zap.Error(err))
		return err
	}
	return nil
}
