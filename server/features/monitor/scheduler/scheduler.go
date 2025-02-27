package monitor_scheduler

import (
	"github.com/wangxin688/narvis/intend/intendtask"
	"github.com/wangxin688/narvis/intend/logger"
	infra_biz "github.com/wangxin688/narvis/server/features/infra/biz"
	infra_tasks "github.com/wangxin688/narvis/server/features/infra/tasks"
	"github.com/wangxin688/narvis/server/pkg/contextvar"
	"go.uber.org/zap"
)

func WlanUserTask() {
	allSites, err := infra_biz.NewSiteService().GetAllActiveSites()
	if err != nil {
		logger.Logger.Warn("[WlanUserScheduler]: get all sites failed", zap.Error(err))
	}
	for _, site := range allSites {
		siteId := site.Id
		orgId := site.OrganizationId
		contextvar.OrganizationId.Set(orgId)
		ids, err := infra_tasks.GenerateSNMPTask(siteId, intendtask.WlanUser, intendtask.WlanUserCallback)
		if err != nil {
			logger.Logger.Warn("[WlanUserScheduler]: generate snmp task failed", zap.Error(err))
		}
		if len(ids) > 0 {
			logger.Logger.Info("[WlanUserScheduler]: generate snmp task success", zap.Any("taskIds", ids))
		}

	}
}
