package hooks

import (
	"fmt"

	"github.com/wangxin688/narvis/intend/logger"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/pkg/contextvar"
	"github.com/wangxin688/narvis/server/pkg/zbx"
	"github.com/wangxin688/narvis/server/pkg/zbx/zschema"
	"go.uber.org/zap"
)

func SiteHookCreate(siteId string) (string, error) {
	groupId, err := zbx.NewZbxClient().HostGroupCreate(&zschema.HostGroupCreate{
		Name: siteId,
	})
	if err != nil {
		logger.Logger.Error("[siteCreateHooks]:create host group failed", zap.Error(err))
		return "", err
	}
	// ignore potential db update error here because of MVP version, it will be fixed in future version
	_, err = gen.Site.Where(gen.Site.Id.Eq(siteId)).UpdateColumn(gen.Site.MonitorId, groupId)
	if err != nil {
		logger.Logger.Error("[siteCreateHooks]: update site monitorId failed", zap.Error(err))
		return "", err
	}
	logger.Logger.Info("[siteCreateHooks]: create host group success", zap.String("groupId", groupId))
	return groupId, nil
}

func SiteHookUpdate(siteId string, diff map[string]*contextvar.Diff) {
	if diff == nil {
		return
	}
	site, err := gen.Site.Where(gen.Site.Id.Eq(siteId)).First()
	if err != nil {
		logger.Logger.Error("[siteUpdateHooks]: get site failed", zap.Error(err))
		return
	}
	if site.MonitorId == nil && site.Status == "Active" {
		SiteHookCreate(siteId)
	} else {
		status, ok := diff["status"]
		if !ok {
			return
		}
		var hostStatus uint8
		if status.After == "Inactive" {
			hostStatus = 1
		} else if status.After == "Active" {
			hostStatus = 0
		}
		zapi := zbx.NewZbxClient()
		hostIds, err := getSiteHostIds(*site.MonitorId, zapi)
		if err != nil || len(hostIds) == 0 {
			return
		}
		zapi.HostMassUpdate(&zschema.HostMassUpdate{
			Hosts: func() []zschema.HostID {
				result := make([]zschema.HostID, 0)
				for _, v := range hostIds {
					result = append(result, zschema.HostID{
						HostID: v,
					})
				}
				return result
			}(),
			Status: &hostStatus,
		})
	}
}

func SiteHookDelete(site *models.Site) {
	if site.MonitorId != nil && *site.MonitorId != "" {
		groupIds, err := zbx.NewZbxClient().HostGroupDelete([]string{*site.MonitorId})
		if err != nil {
			logger.Logger.Error("[siteDeleteHooks]: delete host group failed", zap.Error(err))
		}
		logger.Logger.Info(fmt.Sprintf("[siteDeleteHooks]: delete host group %v success", groupIds), zap.String("groupId", *site.MonitorId))
	}
}

func getSiteHostIds(monitorId string, client *zbx.Zbx) ([]string, error) {
	result := make([]string, 0)
	req := &zschema.HostGet{
		GroupIDs: &[]string{monitorId},
		Output:   "hostids",
	}
	rsp, err := client.HostGet(req)
	if err != nil {
		logger.Logger.Error("[siteDeleteHooks]: get host id failed", zap.Error(err))
		return result, err
	}
	for _, v := range rsp {
		result = append(result, v.HostID)
	}
	return result, nil
}
