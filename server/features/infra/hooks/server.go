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

func ServerCreateHooks(serverId string) {
	server, err := gen.Server.Where(gen.Server.Id.Eq(serverId)).First()
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[serverCreateHooks]: for server failed with serverId %s", serverId), zap.Error(err))
		return
	}
	proxy := proxySelect(server.OrganizationId)
	template, err := serverTemplateSelect(server.OsVersion)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[serverCreateHooks]: for server failed with serverId %s", serverId), zap.Error(err))
		return
	}
	community, port := serverSnmpV2CommunitySelect(serverId)
	hostGroupId, err := getHostGroupId(server.SiteId)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[serverCreateHooks]: getHostGroupId for server failed with serverId %s", serverId), zap.Error(err))
		return
	}
	if hostGroupId == nil || *hostGroupId == "" {
		logger.Logger.Info(fmt.Sprintf("[serverCreateHooks]: hostGroupId is empty for server %s, create it now", serverId))
		hgId, err := SiteHookCreate(server.SiteId)
		if err != nil {
			logger.Logger.Error(fmt.Sprintf("[serverCreateHooks]: create host group failed for siteId %s", server.SiteId), zap.Error(err))
			return
		}
		hostGroupId = &hgId
	}
	hostId, err := zbx.NewZbxClient().HostCreate(&zschema.HostCreate{
		Host:        "s_" + serverId,
		Interfaces:  genHostInterfaces(server.ManagementIp, community, port),
		Groups:      genGroups(*hostGroupId),
		Tags:        genServerTags(server),
		MonitoredBy: 1,
		ProxyID:     &proxy,
		Status:      getHostMonitorStatus(server.Status),
		Templates:   genTemplates(template),
	})
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[serverCreateHooks]: create host failed for server %s", serverId), zap.Error(err))
		return
	}
	logger.Logger.Info(fmt.Sprintf("[serverCreateHooks]: create host success for server %s", serverId))
	// ignore potential db update error here because of MVP version, it will be fixed in future version
	_, err = gen.Server.Where(gen.Server.Id.Eq(serverId)).Update(gen.Server.MonitorId, hostId)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[serverCreateHooks]: update hostId failed for server %s", serverId), zap.Error(err))
		return
	}
}

func ServerUpdateHooks(serverId string, diff map[string]*contextvar.Diff) {
	if len(diff) == 0 {
		logger.Logger.Info("[serverUpdateHooks]: no diff found for server skip update ", zap.String("serverId", serverId))
		return
	}
	server, err := gen.Server.Where(gen.Server.Id.Eq(serverId)).First()
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[serverUpdateHooks]: get server failed with serverId %s", serverId), zap.Error(err))
		return
	}
	client := zbx.NewZbxClient()
	if server.MonitorId == nil || *server.MonitorId == "" {
		logger.Logger.Info(fmt.Sprintf("[serverUpdateHooks]: hostId is empty for server %s, create it now", serverId))
		ServerCreateHooks(serverId)
		return
	}
	updateSchema := zschema.HostUpdate{HostID: *server.MonitorId}
	for k, v := range diff {
		switch k {
		case "status":
			*updateSchema.Status = getHostMonitorStatus(v.After.(string))
		case "osVersion":
			template, err := serverTemplateSelect(v.After.(string))
			if err != nil {
				logger.Logger.Error(fmt.Sprintf("[serverCreateHooks]: osVersion %s not support in monitoring yet", server.OsVersion), zap.Error(err))
				return
			}
			*updateSchema.TemplateClear = genTemplates(template)
		case "managementIp":
			zbxInterfaces, err := client.HostInterfaceGet(&zschema.HostInterfaceGet{
				HostIDs: []string{*server.MonitorId},
			})
			if err != nil {
				logger.Logger.Error(fmt.Sprintf("[serverUpdateHooks]: get host interfaces failed for server %s", serverId), zap.Error(err))
				return
			}
			hostInterfaceId := zbxInterfaces[0].InterfaceId
			hostInterfaces := make([]zschema.HostInterfaceUpdate, 0)
			updateIp := getPureIp(server.ManagementIp)
			hostInterfaces = append(hostInterfaces, zschema.HostInterfaceUpdate{
				InterfaceID: hostInterfaceId,
				Ip:          &updateIp,
			})
			updateSchema.Interfaces = &hostInterfaces
		}
	}
	_, err = client.HostUpdate(&updateSchema)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[serverUpdateHooks]: update host failed for server %s", serverId), zap.Error(err))
		return
	}
}

func ServerDeleteHooks(server *models.Server) {
	if server == nil || server.MonitorId == nil || *server.MonitorId == "" {
		return
	}
	_, err := zbx.NewZbxClient().HostDelete([]string{*server.MonitorId})
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[serverDeleteHooks]: delete host failed for server %s", server.Id), zap.Error(err))
	}
	logger.Logger.Info(fmt.Sprintf("[serverDeleteHooks]: delete host success for server %s", server.Id))
}
