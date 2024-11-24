package infra_scheduler

import (
	"time"

	"github.com/wangxin688/narvis/intend/intendtask"
	"github.com/wangxin688/narvis/intend/logger"
	"github.com/wangxin688/narvis/server/dal/gen"
	infra_biz "github.com/wangxin688/narvis/server/features/infra/biz"
	infra_tasks "github.com/wangxin688/narvis/server/features/infra/tasks"
	"github.com/wangxin688/narvis/server/pkg/contextvar"
	"go.uber.org/zap"
)

func SyncDeviceScheduler() {
	allSites, err := infra_biz.NewSiteService().GetAllActiveSites()
	if err != nil {
		logger.Logger.Warn("[syncDeviceScheduler]: get all sites failed", zap.Error(err))
	}

	for _, site := range allSites {
		siteId := site.Id
		orgId := site.OrganizationId
		contextvar.OrganizationId.Set(orgId)
		ids, err := infra_tasks.GenerateSNMPTask(siteId, intendtask.ScanDevice, intendtask.ScanDeviceCallback)
		if err != nil {
			logger.Logger.Warn("[syncDeviceScheduler]: generate snmp task failed", zap.Error(err))
		}
		if len(ids) > 0 {
			logger.Logger.Info("[syncDeviceScheduler]: generate snmp task success", zap.Any("taskIds", ids))
		}

	}
}

func SyncApScheduler() {
	allSites, err := infra_biz.NewSiteService().GetAllActiveSites()
	if err != nil {
		logger.Logger.Error("[syncApScheduler]: get all sites failed", zap.Error(err))
	}

	for _, site := range allSites {
		siteId := site.Id
		orgId := site.OrganizationId
		contextvar.OrganizationId.Set(orgId)
		ids, err := infra_tasks.GenerateSNMPTask(siteId, intendtask.ScanAp, intendtask.ScanApCallback)
		if err != nil {
			logger.Logger.Warn("[syncApScheduler]: generate snmp task failed", zap.Error(err))
		}
		if len(ids) > 0 {
			logger.Logger.Warn("[syncApScheduler]: generate snmp task success", zap.Any("taskIds", ids))
		}
	}
}

func SyncConfigBackupScheduler() {

	allSites, err := infra_biz.NewSiteService().GetAllActiveSites()
	if err != nil {
		logger.Logger.Error("[syncConfigBackupScheduler]: get all sites failed", zap.Error(err))
	}

	for _, site := range allSites {
		siteId := site.Id
		orgId := site.OrganizationId
		contextvar.OrganizationId.Set(orgId)
		ids, err := infra_tasks.ConfigBackUpTask(siteId, intendtask.ConfigurationBackup, intendtask.ConfigurationBackupCallback)
		if err != nil {
			logger.Logger.Warn("[syncConfigBackupScheduler]: generate snmp task failed", zap.Error(err))
		}
		if len(ids) > 0 {
			logger.Logger.Warn("[syncConfigBackupScheduler]: generate snmp task success", zap.Any("taskIds", ids))
		}
	}
}

// delete task results older than 30 days
func HouseKeepingResultRecycle() {
	_, err := gen.TaskResult.Where(
		gen.TaskResult.CreatedAt.Lt(time.Now().AddDate(0, 0, -30)),
	).Delete()
	if err != nil {
		logger.Logger.Error("[houseKeepingResultRecycle]: delete history task results failed", zap.Error(err))
	}
}
