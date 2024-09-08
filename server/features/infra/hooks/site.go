package hooks

import (
	"fmt"

	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/pkg/zbx"
	"github.com/wangxin688/narvis/server/pkg/zbx/zschema"
	ts "github.com/wangxin688/narvis/server/tools/schemas"
	"go.uber.org/zap"
)

func SiteHookCreate(siteId string) {
	groupId, err := zbx.NewZbxClient().HostGroupCreate(&zschema.HostGroupCreate{
		Name: siteId,
	})
	if err != nil {
		core.Logger.Error("site create hooks: create host group failed", zap.Error(err))
	}
	gen.Site.Where(gen.Site.Id.Eq(siteId)).UpdateColumn(gen.Site.MonitorId, groupId)
	core.Logger.Info("site create hooks: create host group success", zap.String("groupId", groupId))
}

func SiteHookUpdate(siteId string, diff map[string]*ts.OrmDiff) error {
	if diff == nil {
		return nil
	}
	site, err := gen.Site.Where(gen.Site.Id.Eq(siteId)).First()
	if err != nil {
		core.Logger.Error("site update hooks: get site failed", zap.Error(err))
		return err
	}
	if site.MonitorId == nil && site.Status == "Active" {
		SiteHookCreate(siteId)
	}
	// if site.
	return nil
}

func SiteHookDelete(site *models.Site) {
	if site.MonitorId != nil && *site.MonitorId != "" {
		groupIds, err := zbx.NewZbxClient().HostGroupDelete([]string{*site.MonitorId})
		if err != nil {
			core.Logger.Error("site delete hooks: delete host group failed", zap.Error(err))
		}
		core.Logger.Info(fmt.Sprintf("site delete hooks: delete host group %v success", groupIds), zap.String("groupId", *site.MonitorId))
	}
}
