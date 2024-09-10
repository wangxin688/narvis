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

func DeviceCreateHooks(deviceId string) {
	device, err := gen.Device.Where(gen.Device.Id.Eq(deviceId)).First()
	if err != nil {
		core.Logger.Error(fmt.Sprintf("[deviceCreateHooks]: get device failed with device %s", deviceId), zap.Error(err))
		return
	}
	proxy := proxySelect(device.OrganizationId)
	template, err := deviceTemplateSelect(device)
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
		Host:        "d_" + deviceId,
		Interfaces:  genHostInterfaces(device.ManagementIp, community, port),
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
		return
	}
	device, err := gen.Device.Where(gen.Device.Id.Eq(deviceId)).First()
	if err != nil {
		core.Logger.Error(fmt.Sprintf("[deviceUpdateHooks]: get device failed with device %s", deviceId), zap.Error(err))
		return
	}
	client := zbx.NewZbxClient()
	if device.MonitorId == nil || *device.MonitorId == "" {
		core.Logger.Info(fmt.Sprintf("[deviceUpdateHooks]: hostId is empty for device %s, create it now", deviceId))
		DeviceCreateHooks(deviceId)
		return
	}
	updateSchema := zschema.HostUpdate{HostID: *device.MonitorId}
	for k, v := range diff {
		switch k {
		case "status":
			*updateSchema.Status = getHostMonitorStatus(v.After.(string))
		case "deviceRole", "manufacturer":
			template, err := deviceTemplateSelect(device)
			if err != nil {
				core.Logger.Error(fmt.Sprintf("[deviceCreateHooks]: device role %s and manufacturer %s not support in monitoring yet", device.DeviceRole, device.Manufacturer), zap.Error(err))
				return
			}
			*updateSchema.TemplateClear = genTemplates(template)
		case "managementIp":
			zbxInterfaces, err := client.HostInterfaceGet(&zschema.HostInterfaceGet{
				HostIDs: []string{*device.MonitorId},
			})
			if err != nil || len(zbxInterfaces) <= 0 {
				core.Logger.Error(fmt.Sprintf("[deviceUpdateHooks]: failed to update delete due to get host interface failed for device %s", deviceId), zap.Error(err))
				return
			}
			hostInterfaceId := zbxInterfaces[0].InterfaceId
			hostInterfaces := make([]zschema.HostInterfaceUpdate, 0)
			updateIp := getPureIp(device.ManagementIp)
			hostInterfaces = append(hostInterfaces, zschema.HostInterfaceUpdate{
				InterfaceID: hostInterfaceId,
				Ip:          &updateIp,
			})
			updateSchema.Interfaces = &hostInterfaces
		}
	}
	_, err = client.HostUpdate(&updateSchema)
	if err != nil {
		core.Logger.Error(fmt.Sprintf("[deviceUpdateHooks]: update host failed for device %s", deviceId), zap.Error(err))
		return
	}
	core.Logger.Info(fmt.Sprintf("[deviceUpdateHooks]: update host success for device %s", deviceId))
}

func DeviceDeleteHooks(device *models.Device) {
	if device == nil || device.MonitorId == nil || *device.MonitorId == "" {
		return
	}
	_, err := zbx.NewZbxClient().HostDelete([]string{*device.MonitorId})
	if err != nil {
		core.Logger.Error(fmt.Sprintf("[deviceDeleteHooks]: delete host failed for device %s", device.Id), zap.Error(err))
	}
	core.Logger.Info(fmt.Sprintf("[deviceDeleteHooks]: delete host success for device %s", device.Id))
}
