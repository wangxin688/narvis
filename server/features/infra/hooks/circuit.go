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

func CircuitCreateHooks(circuitId string) {
	circuit, err := gen.Circuit.Where(gen.Circuit.Id.Eq(circuitId)).First()
	if err != nil {
		core.Logger.Error(fmt.Sprintf("[circuitCreateHooks]: get circuit failed with circuit %s", circuitId), zap.Error(err))
		return
	}
	proxy := proxySelect(circuit.OrganizationId)
	template, err := circuitTemplateSelect()
	if err != nil {
		core.Logger.Error(fmt.Sprintf("[circuitCreateHooks]: get template failed with circuit %s", circuitId), zap.Error(err))
		return
	}
	hostGroupId, err := getHostGroupId(circuit.SiteId)
	if err != nil {
		core.Logger.Error(fmt.Sprintf("[deviceCreateHooks]: getHostGroupId for device failed with device %s", circuitId),
			zap.Error(err))
		return
	}
	if hostGroupId == nil || *hostGroupId == "" {
		core.Logger.Info(fmt.Sprintf("[circuitCreateHooks]: hostGroupId is empty for circuit %s, create it now", circuitId))
		hgId, err := SiteHookCreate(circuit.SiteId)
		if err != nil {
			core.Logger.Error(fmt.Sprintf("[circuitCreateHooks]: create host group failed for siteId %s", circuit.SiteId),
				zap.Error(err))
			return
		}
		hostGroupId = &hgId
	}
	if (circuit.Ipv4Address == nil && circuit.Ipv6Address == nil) || (*circuit.Ipv4Address == "" && *circuit.Ipv6Address == "") {
		core.Logger.Info(fmt.Sprintf("[circuitCreateHooks]: skip create host for circuit %s has no address", circuitId))
	}
	hostId, err := zbx.NewZbxClient().HostCreate(&zschema.HostCreate{
		Host:        "c_" + circuitId,
		Interfaces:  genCircuitInterfaces(circuit),
		Tags:        genCircuitTags(circuit),
		Templates:   genTemplates(template),
		Groups:      genGroups(*hostGroupId),
		MonitoredBy: 1,
		ProxyID:     &proxy,
		Status:      getHostMonitorStatus(circuit.Status),
	})
	if err != nil {
		core.Logger.Error(fmt.Sprintf("[circuitCreateHooks]: create host failed for circuit %s", circuitId),
			zap.Error(err))
		return
	}
	core.Logger.Info(fmt.Sprintf("[circuitCreateHooks]: create host success for circuit %s", circuitId),
		zap.String("hostId", hostId))
	gen.Circuit.Where(gen.Circuit.Id.Eq(circuitId)).Update(gen.Circuit.MonitorId, hostId)

	circuitDeviceCreateHooks(circuit, proxy)
}

func circuitDeviceCreateHooks(circuit *models.Circuit, proxy string) {
	circuitDevice, err := gen.Device.Where(gen.Device.Id.Eq(circuit.DeviceId)).First()
	if err != nil {
		core.Logger.Error(
			fmt.Sprintf("[circuitHostCreateHooks]: get circuit connected device %s failed with circuit %s",
				circuit.DeviceId, circuit.Id),
			zap.Error(err))
		return
	}
	deviceInterface, err := gen.DeviceInterface.Where(gen.DeviceInterface.Id.Eq(circuit.InterfaceId)).First()
	if err != nil {
		core.Logger.Error(
			fmt.Sprintf("[circuitHostCreateHooks]: get circuit connected interface %s failed with circuit %s",
				circuit.InterfaceId, circuit.Id),
			zap.Error(err))
		return
	}

	community, port := snmpV2CommunitySelect(circuit.DeviceId)
	template, err := circuitHostTemplateSelect()
	if err != nil {
		core.Logger.Error(fmt.Sprintf("[circuitHostCreateHooks]: get template failed with circuit %s", circuit.Id), zap.Error(err))
		return
	}
	hostGroupId, err := getHostGroupId(circuit.SiteId)
	if err != nil {
		core.Logger.Error(fmt.Sprintf("[circuitHostCreateHooks]: getHostGroupId for device failed with device %s", circuit.DeviceId), zap.Error(err))
		return
	}
	if hostGroupId == nil || *hostGroupId == "" {
		core.Logger.Info(fmt.Sprintf("[circuitHostCreateHooks]: hostGroupId is empty for circuit %s, create it now", circuit.Id))
		hgId, err := SiteHookCreate(circuit.SiteId)
		if err != nil {
			core.Logger.Error(fmt.Sprintf("[circuitHostCreateHooks]: create host group failed for siteId %s", circuit.SiteId), zap.Error(err))
			return
		}
		hostGroupId = &hgId
	}
	hostId, err := zbx.NewZbxClient().HostCreate(&zschema.HostCreate{
		Host:        "cd_" + circuit.DeviceId,
		Interfaces:  genHostInterfaces(circuitDevice.ManagementIp, community, port),
		Tags:        genDeviceTags(circuitDevice),
		Macros:      genCircuitMacros(circuit, deviceInterface),
		Templates:   genTemplates(template),
		Groups:      genGroups(*hostGroupId),
		MonitoredBy: 1,
		ProxyID:     &proxy,
		Status:      getHostMonitorStatus(circuit.Status),
	})
	if err != nil {
		core.Logger.Error(fmt.Sprintf("[circuitHostCreateHooks]: create host failed for circuit %s", circuit.Id), zap.Error(err))
		return
	}
	core.Logger.Info(fmt.Sprintf("[circuitHostCreateHooks]: create host success for circuit %s", circuit.Id), zap.String("hostId", hostId))
	gen.Circuit.Where(gen.Circuit.Id.Eq(circuit.Id)).Update(gen.Circuit.MonitorHostId, hostId)
}

func CircuitUpdateHooks(circuitId string, diff map[string]*ts.OrmDiff) {
	if len(diff) == 0 {
		core.Logger.Info("[circuitUpdateHooks]: no diff found for circuit skip update ", zap.String("circuitId", circuitId))
		return
	}
	circuit, err := gen.Circuit.Where(gen.Circuit.Id.Eq(circuitId)).First()
	if err != nil {
		core.Logger.Error(fmt.Sprintf("[circuitUpdateHooks]: get circuit failed for circuit %s", circuitId), zap.Error(err))
		return
	}
	client := zbx.NewZbxClient()
	if circuit.MonitorId == nil || *circuit.MonitorId == "" {
		core.Logger.Error(fmt.Sprintf("[circuitUpdateHooks]: monitorId is empty for circuit %s, create it now", circuitId))
		CircuitCreateHooks(circuit.Id)
		return
	}
	if circuit.MonitorHostId == nil || *circuit.MonitorHostId == "" {
		core.Logger.Error(fmt.Sprintf("[circuitUpdateHooks]: monitorHostId is empty for circuit %s, create it now", circuitId))
		proxy := proxySelect(circuit.OrganizationId)
		circuitDeviceCreateHooks(circuit, proxy)
		return
	}
	updateSchema := zschema.HostUpdate{HostID: *circuit.MonitorId}
	if len(diff) == 1 && diff["status"] != nil {
		*updateSchema.Status = getHostMonitorStatus(diff["status"].After.(string))
		_, err := client.HostUpdate(&updateSchema)
		if err != nil {
			core.Logger.Error(fmt.Sprintf("[circuitUpdateHooks]: update host failed for circuit %s", circuitId), zap.Error(err))
			return
		}
		return
	} else if len(diff) > 1 {
		deleteHosts := make([]string, 0)
		if circuit.MonitorId != nil && *circuit.MonitorId != "" {
			deleteHosts = append(deleteHosts, *circuit.MonitorId)
		}
		if circuit.MonitorHostId != nil && *circuit.MonitorHostId != "" {
			deleteHosts = append(deleteHosts, *circuit.MonitorHostId)
		}
		if len(deleteHosts) > 0 {
			_, err := client.HostDelete([]string{*circuit.MonitorId, *circuit.MonitorHostId})
			if err != nil {
				core.Logger.Error(fmt.Sprintf("[circuitUpdateHooks]: delete host failed for circuit %s", circuitId), zap.Error(err))
				return
			}
			CircuitCreateHooks(circuit.Id)
			circuitDeviceCreateHooks(circuit, proxySelect(circuit.OrganizationId))
		}
	}
}

func CircuitDeleteHooks(circuit *models.Circuit) {
	deleteHosts := make([]string, 0)
	if circuit.MonitorId != nil && *circuit.MonitorId != "" {
		deleteHosts = append(deleteHosts, *circuit.MonitorId)
	}
	if circuit.MonitorHostId != nil && *circuit.MonitorHostId != "" {
		deleteHosts = append(deleteHosts, *circuit.MonitorHostId)
	}
	if len(deleteHosts) > 0 {
		_, err := zbx.NewZbxClient().HostDelete([]string{*circuit.MonitorId, *circuit.MonitorHostId})
		if err != nil {
			core.Logger.Error(fmt.Sprintf("[circuitDeleteHooks]: delete host failed for circuit %s", circuit.Id), zap.Error(err))
			return
		}
	}
	core.Logger.Info(fmt.Sprintf("[circuitDeleteHooks]: delete host success for circuit %s", circuit.Id))
}
