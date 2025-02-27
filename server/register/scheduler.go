package register

import (
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/wangxin688/narvis/intend/logger"
	device360_tasks "github.com/wangxin688/narvis/server/features/device360/tasks"
	infra_scheduler "github.com/wangxin688/narvis/server/features/infra/scheduler"
	monitor_scheduler "github.com/wangxin688/narvis/server/features/monitor/scheduler"
	"go.uber.org/zap"
)

// 如果未来gin需要在多个容器中运行，使用 Redis/Postgresql 实现分布式锁来确保任务的只运行一次
func RegisterScheduler() {
	scheduler, err := gocron.NewScheduler()
	if err != nil {
		logger.Logger.Error("[registerScheduler]: failed to create scheduler", zap.Error(err))
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
		logger.Logger.Error("[registerScheduler]: failed to create device360 offline job", zap.Error(err))
	}
	logger.Logger.Info("[registerScheduler]: create device360 offline job success", zap.String("jobId", device360OfflineJob.ID().String()))

	deviceInfoJob, err := scheduler.NewJob(
		gocron.DurationJob(8*time.Hour),
		gocron.NewTask(
			func() {
				infra_scheduler.SyncDeviceScheduler()
			},
		),
	)
	if err != nil {
		logger.Logger.Error("[registerScheduler]: failed to create device info job", zap.Error(err))
	}
	logger.Logger.Info("[registerScheduler]: create device info job success", zap.String("jobId", deviceInfoJob.ID().String()))

	apInfoJob, err := scheduler.NewJob(
		gocron.DurationJob(8*time.Hour),
		gocron.NewTask(
			func() {
				infra_scheduler.SyncApScheduler()
			},
		),
	)
	if err != nil {
		logger.Logger.Error("[registerScheduler]: failed to create ap info job", zap.Error(err))
	}
	logger.Logger.Info("[registerScheduler]: create ap info job success", zap.String("jobId", apInfoJob.ID().String()))

	configBackupJob, err := scheduler.NewJob(
		gocron.DurationJob(24*time.Hour),
		gocron.NewTask(
			func() {
				infra_scheduler.SyncConfigBackupScheduler()
			},
		),
	)
	if err != nil {
		logger.Logger.Error("[registerScheduler]: failed to create config backup job", zap.Error(err))
	}
	logger.Logger.Info("[registerScheduler]: create config backup job success", zap.String("jobId", configBackupJob.ID().String()))

	wlanUserJob, err := scheduler.NewJob(
		gocron.DurationJob(5*time.Minute),
		gocron.NewTask(
			func() {
				monitor_scheduler.WlanUserTask()
			},
		),
	)
	if err != nil {
		logger.Logger.Error("[registerScheduler]: failed to create wlan user job", zap.Error(err))
	}
	logger.Logger.Info("[registerScheduler]: create wlan user job success", zap.String("jobId", wlanUserJob.ID().String()))

	houseKeepingJob, err := scheduler.NewJob(
		gocron.DurationJob(24*time.Hour),
		gocron.NewTask(
			func() {
				infra_scheduler.HouseKeepingResultRecycle()
			},
		),
	)
	if err != nil {
		logger.Logger.Error("[registerScheduler]: failed to create house keeping job", zap.Error(err))
	}
	logger.Logger.Info("[registerScheduler]: create house keeping job success", zap.String("jobId", houseKeepingJob.ID().String()))

	scheduler.Start()
	logger.Logger.Info("[registerScheduler]: scheduler started")
}
