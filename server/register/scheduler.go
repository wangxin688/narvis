package register

import (
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/wangxin688/narvis/server/core"
	device360_tasks "github.com/wangxin688/narvis/server/features/device360/tasks"
	"go.uber.org/zap"
)


// 如果未来gin需要在多个容器中运行，使用 Redis/Postgresql 实现分布式锁来确保任务的只运行一次
func RegisterScheduler() {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		core.Logger.Error("[registerScheduler]: failed to create scheduler", zap.Error(err))
		panic(err)
	}
	device360OfflineJob, err := scheduler.NewJob(
		gocron.DurationJob(180*time.Second),
		gocron.NewTask(
			func() {
				device360_tasks.RunDevice360OfflineTask()
			},
		),
	)
	if err != nil {
		core.Logger.Error("[registerScheduler]: failed to create device360 offline job", zap.Error(err))
		panic(err)
	}
	core.Logger.Info("[registerScheduler]: create device360 offline job success", zap.String("jobId", device360OfflineJob.ID().String()))

	scheduler.Start()
	core.Logger.Info("[registerScheduler]: scheduler started")
}
