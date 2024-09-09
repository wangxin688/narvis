package hooks

import (
	"fmt"

	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/pkg/zbx"
	"github.com/wangxin688/narvis/server/pkg/zbx/zschema"
	ts "github.com/wangxin688/narvis/server/tools/schemas"
	"go.uber.org/zap"
)

func DeviceCreateHooks(deviceId string) {
	device, err := gen.Device.Where(gen.Device.Id.Eq(deviceId)).First()
	if err != nil {
		return
	}
	proxy := proxySelect(device.OrganizationId)
	template, err := templateSelect(device)
	if err != nil {
		core.Logger.Error(fmt.Sprintf("[deviceCreateHooks]: device role %s and manufacturer %s not support in monitoring yet", device.DeviceRole, device.Manufacturer), zap.Error(err))
		return
	}
	community, port := snmpV2CommunitySelect(deviceId)
	hostGroupId, err := getHostGroupId(device.SiteId)
	if err != nil {
		core.Logger.Error(fmt.Sprintf("[deviceCreateHooks]: getHostGroupId for device failed with device %s", deviceId), zap.Error(err))
		return
	}
	if hostGroupId == nil || *hostGroupId == "" {
		core.Logger.Info(fmt.Sprintf("[deviceCreateHooks]: hostGroupId is empty for device %s, create it now", deviceId))
		hgId, err := SiteHookCreate(device.SiteId)
		if err != nil {
			core.Logger.Error(fmt.Sprintf("[deviceCreateHooks]: create host group failed for siteId %s", device.SiteId), zap.Error(err))
			return
		}
		hostGroupId = &hgId
	}
	hostId, err := zbx.NewZbxClient().HostCreate(&zschema.HostCreate{
		Host:        deviceId,
		Interfaces:  genHostInterfaces(device, community, port),
		Groups:      genGroups(*hostGroupId),
		Tags:        genDeviceTags(device),
		MonitoredBy: 1,
		ProxyID:     &proxy,
		Status:      getHostMonitorStatus(device.Status),
		Templates:   genTemplates(template),
	})
	if err != nil {
		core.Logger.Error(fmt.Sprintf("[deviceCreateHooks]: create host failed for device %s", deviceId), zap.Error(err))
		return
	}
	core.Logger.Info(fmt.Sprintf("[deviceCreateHooks]: create host success for device %s with hostId %s", deviceId, hostId))
	// ignore potential db update error here because of MVP version, it will be fixed in future version
	gen.Device.Where(gen.Device.Id.Eq(deviceId)).Update(gen.Device.MonitorId, hostId)
}

func DeviceUpdateHooks(deviceId string, diff map[string]*ts.OrmDiff) {
	if len(diff) == 0 {
		core.Logger.Info("[deviceUpdateHooks]: no diff found for device skip update ", zap.String("deviceId", deviceId))
	}
	device, err := gen.Device.Where(gen.Device.Id.Eq(deviceId)).First()
	if err != nil {
		core.Logger.Error(fmt.Sprintf("[deviceUpdateHooks]: get device failed with device %s", deviceId), zap.Error(err))
		return
	}
	updateSchema := zschema.HostUpdate{}
	for k, v := range diff {
		switch k {
		case "status":
			*updateSchema.Status = getHostMonitorStatus(v.After.(string))
		case "deviceRole", "manufacturer":
			template, err := templateSelect(device)
			if err != nil {
				core.Logger.Error(fmt.Sprintf("[deviceCreateHooks]: device role %s and manufacturer %s not support in monitoring yet", device.DeviceRole, device.Manufacturer), zap.Error(err))
				return
			}
			*updateSchema.TemplateClear = genTemplates(template)
			// case "managementIp":
			// 	*updateSchema. = v.After.(string)

		}
		// TODO: complete other fields
	}
}
