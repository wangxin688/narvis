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

func SnmpCredCreateHooks(credId string) {
	cred, err := gen.SnmpV2Credential.Where(gen.SnmpV2Credential.Id.Eq(credId)).Preload(gen.SnmpV2Credential.Organization).First()
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[snmpCredCreateHooks]: get snmp cred failed with cred %s", credId), zap.Error(err))
		return
	}
	// create a global macro for given snmp cred
	client := zbx.NewZbxClient()
	if cred.DeviceId == nil || *cred.DeviceId == "" {
		globalMacro := zschema.Macro{
			Macro: getGlobalCommunityMacroName(cred.Organization.EnterpriseCode),
			Value: cred.Community,
		}
		umId, err := client.UserMacroCreateGlobal(&globalMacro)
		if err != nil {
			logger.Logger.Error(fmt.Sprintf("[snmpCredCreateHooks]: create global macro failed with cred %s", credId), zap.Error(err))
			return
		}
		logger.Logger.Info(fmt.Sprintf("[snmpCredCreateHooks]: create global macro success with cred %s", credId))
		_, err = gen.SnmpV2Credential.Where(gen.SnmpV2Credential.Id.Eq(credId)).UpdateColumn(gen.SnmpV2Credential.GlobalMacroId, umId)
		if err != nil {
			logger.Logger.Error(fmt.Sprintf("[snmpCredCreateHooks]: update snmp cred failed with cred %s", credId), zap.Error(err))
			return
		}
		logger.Logger.Info(fmt.Sprintf("[snmpCredCreateHooks]: update snmp cred success with cred %s", credId))
		return
	}

	device, err := gen.Device.Where(gen.Device.Id.Eq(*cred.DeviceId), gen.Device.MonitorId.IsNotNull()).First()
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[snmpCredCreateHooks]: get device failed with cred %s", credId), zap.Error(err))
		return
	}
	zbxInterfaces, err := client.HostInterfaceGet(&zschema.HostInterfaceGet{
		HostIDs: []string{*device.MonitorId},
	})
	if err != nil || len(zbxInterfaces) <= 0 {
		logger.Logger.Error(fmt.Sprintf("[snmpCredCreateHooks]: failed to update delete due to get host interface failed for device %s", device.Id), zap.Error(err))
		return
	}
	hostInterfaceId := zbxInterfaces[0].InterfaceId
	hostInterfaces := make([]zschema.HostInterfaceUpdate, 0)
	hostInterfaces = append(hostInterfaces, zschema.HostInterfaceUpdate{
		InterfaceID: hostInterfaceId,
		Details: &zschema.Details{
			Community:      cred.Community,
			MaxRepetitions: &cred.MaxRepetitions,
		},
	})
	updateSchema := zschema.HostUpdate{HostID: *device.MonitorId}
	updateSchema.Interfaces = &hostInterfaces
	_, err = client.HostUpdate(&updateSchema)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[snmpCredCreateHooks]: update host failed for cred %s", credId), zap.Error(err))
		return
	}
	logger.Logger.Info(fmt.Sprintf("[snmpCredCreateHooks]: update host success for cred %s", credId))

}

func ServerSnmpCredCreateHooks(credId string) {
	cred, err := gen.ServerSnmpCredential.Where(gen.ServerSnmpCredential.Id.Eq(credId)).First()
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[serverSnmpCredCreateHooks]: get server snmp cred failed with cred %s", credId), zap.Error(err))
		return
	}

	client := zbx.NewZbxClient()
	if cred.ServerId == nil || *cred.ServerId == "" {
		logger.Logger.Error(fmt.Sprintf("[serverSnmpCredCreateHooks]: server id is empty for cred %s", credId))
		return
	}
	server, err := gen.Server.Where(gen.Server.Id.Eq(*cred.ServerId), gen.Server.MonitorId.IsNotNull()).First()
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[serverSnmpCredCreateHooks]: get server failed with cred %s", credId), zap.Error(err))
		return
	}
	zbxInterfaces, err := client.HostInterfaceGet(&zschema.HostInterfaceGet{
		HostIDs: []string{*server.MonitorId},
	})
	if err != nil || len(zbxInterfaces) <= 0 {
		logger.Logger.Error(fmt.Sprintf("[serverSnmpCredCreateHooks]: failed to update delete due to get host interface failed for server %s", server.Id), zap.Error(err))
		return
	}
	hostInterfaceId := zbxInterfaces[0].InterfaceId
	hostInterfaces := make([]zschema.HostInterfaceUpdate, 0)
	hostInterfaces = append(hostInterfaces, zschema.HostInterfaceUpdate{
		InterfaceID: hostInterfaceId,
		Details: &zschema.Details{
			Community:      cred.Community,
			MaxRepetitions: &cred.MaxRepetitions,
		},
	})
	updateSchema := zschema.HostUpdate{HostID: *server.MonitorId}
	updateSchema.Interfaces = &hostInterfaces
	_, err = client.HostUpdate(&updateSchema)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[serverSnmpCredCreateHooks]: update host failed for cred %s", credId), zap.Error(err))
		return
	}
	logger.Logger.Info(fmt.Sprintf("[serverSnmpCredCreateHooks]: update host success for cred %s", credId))

}

func SnmpCredUpdateHooks(credId string, diff map[string]*contextvar.Diff) {
	if len(diff) == 0 {
		logger.Logger.Info("[snmpCredUpdateHooks]: no diff found for cred skip update ", zap.String("credId", credId))
		return
	}
	cred, err := gen.SnmpV2Credential.Where(gen.SnmpV2Credential.Id.Eq(credId)).First()
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[snmpCredUpdateHooks]: get snmp cred failed with cred %s", credId), zap.Error(err))
		return
	}
	client := zbx.NewZbxClient()
	if _community, ok := diff["community"]; ok {
		community := _community.After.(string)
		if cred.GlobalMacroId != nil && *cred.GlobalMacroId != "" {
			updateSchema := zschema.GlobalMacroUpdate{
				GlobalMacroId: *cred.GlobalMacroId,
				Value:         community,
			}
			_, err = client.UserMacroUpdateGlobal(&updateSchema)
			if err != nil {
				logger.Logger.Error(fmt.Sprintf("[snmpCredUpdateHooks]: update global macro failed with cred %s", credId), zap.Error(err))
				return
			}
			logger.Logger.Info(fmt.Sprintf("[snmpCredUpdateHooks]: update global macro success with cred %s", credId))
			return
		}

		if cred.DeviceId == nil || *cred.DeviceId == "" {
			device, err := gen.Device.Where(gen.Device.Id.Eq(*cred.DeviceId), gen.Device.MonitorId.IsNotNull()).First()
			if err != nil {
				logger.Logger.Error(fmt.Sprintf("[snmpCredCreateHooks]: get device failed with cred %s", credId), zap.Error(err))
				return
			}
			zbxInterfaces, err := client.HostInterfaceGet(&zschema.HostInterfaceGet{
				HostIDs: []string{*device.MonitorId},
			})
			if err != nil || len(zbxInterfaces) <= 0 {
				logger.Logger.Error(fmt.Sprintf("[snmpCredCreateHooks]: failed to update delete due to get host interface failed for device %s", device.Id), zap.Error(err))
				return
			}
			hostInterfaceId := zbxInterfaces[0].InterfaceId
			hostInterfaces := make([]zschema.HostInterfaceUpdate, 0)
			hostInterfaces = append(hostInterfaces, zschema.HostInterfaceUpdate{
				InterfaceID: hostInterfaceId,
				Details: &zschema.Details{
					Community:      cred.Community,
					MaxRepetitions: &cred.MaxRepetitions,
				},
			})
			updateSchema := zschema.HostUpdate{HostID: *device.MonitorId}
			updateSchema.Interfaces = &hostInterfaces
			_, err = client.HostUpdate(&updateSchema)
			if err != nil {
				logger.Logger.Error(fmt.Sprintf("[snmpCredCreateHooks]: update host failed for cred %s", credId), zap.Error(err))
				return
			}
			logger.Logger.Info(fmt.Sprintf("[snmpCredCreateHooks]: update host success for cred %s", credId))
			return
		}
	}
}

func ServerSnmpCredUpdateHooks(credId string, diff map[string]*contextvar.Diff) {
	if len(diff) == 0 {
		logger.Logger.Info("[serverSnmpCredUpdateHooks]: no diff found for cred skip update ", zap.String("credId", credId))
		return
	}
	cred, err := gen.ServerSnmpCredential.Where(gen.ServerSnmpCredential.Id.Eq(credId)).First()
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[serverSnmpCredUpdateHooks]: get server snmp cred failed with cred %s", credId), zap.Error(err))
		return
	}
	client := zbx.NewZbxClient()
	if _community, ok := diff["community"]; ok {
		community := _community.After.(string)
		if cred.ServerId == nil || *cred.ServerId == "" {
			server, err := gen.Server.Where(gen.Server.Id.Eq(*cred.ServerId), gen.Server.MonitorId.IsNotNull()).First()
			if err != nil {
				logger.Logger.Error(fmt.Sprintf("[serverSnmpCredCreateHooks]: get server failed with cred %s", credId), zap.Error(err))
				return
			}
			zbxInterfaces, err := client.HostInterfaceGet(&zschema.HostInterfaceGet{
				HostIDs: []string{*server.MonitorId},
			})
			if err != nil || len(zbxInterfaces) <= 0 {
				logger.Logger.Error(fmt.Sprintf("[serverSnmpCredCreateHooks]: failed to update delete due to get host interface failed for server %s", server.Id), zap.Error(err))
				return
			}
			hostInterfaceId := zbxInterfaces[0].InterfaceId
			hostInterfaces := make([]zschema.HostInterfaceUpdate, 0)
			hostInterfaces = append(hostInterfaces, zschema.HostInterfaceUpdate{
				InterfaceID: hostInterfaceId,
				Details: &zschema.Details{
					Community:      community,
					MaxRepetitions: &cred.MaxRepetitions,
				},
			})
			updateSchema := zschema.HostUpdate{HostID: *server.MonitorId}
			updateSchema.Interfaces = &hostInterfaces
			_, err = client.HostUpdate(&updateSchema)
			if err != nil {
				logger.Logger.Error(fmt.Sprintf("[serverSnmpCredCreateHooks]: update host failed for cred %s", credId), zap.Error(err))
				return
			}
		}
	}
}

// when delete host cred, use global cred as default
// global cred will not be deleted
func SnmpCredDeleteHooks(cred *models.SnmpV2Credential) {
	globalCred, err := gen.SnmpV2Credential.Where(
		gen.SnmpV2Credential.OrganizationId.Eq(cred.OrganizationId),
		gen.SnmpV2Credential.DeviceId.IsNull()).First()
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[snmpCredDeleteHooks]: get snmp global cred failed with cred %s", cred.Id), zap.Error(err))
		return
	}
	device, err := gen.Device.Where(gen.Device.Id.Eq(*cred.DeviceId), gen.Device.MonitorId.IsNotNull()).First()
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[snmpCredDeleteHooks]: get device failed with cred %s", cred.Id), zap.Error(err))
		return
	}

	client := zbx.NewZbxClient()
	zbxInterfaces, err := client.HostInterfaceGet(&zschema.HostInterfaceGet{
		HostIDs: []string{*device.MonitorId},
	})
	if err != nil || len(zbxInterfaces) <= 0 {
		logger.Logger.Error(fmt.Sprintf("[snmpCredDeleteHooks]: failed to update delete due to get host interface failed for device %s", device.Id), zap.Error(err))
		return
	}
	hostInterfaceId := zbxInterfaces[0].InterfaceId
	hostInterfaces := make([]zschema.HostInterfaceUpdate, 0)
	hostInterfaces = append(hostInterfaces, zschema.HostInterfaceUpdate{
		InterfaceID: hostInterfaceId,
		Details: &zschema.Details{
			Community:      globalCred.Community,
			MaxRepetitions: &globalCred.MaxRepetitions,
		},
	})
	updateSchema := zschema.HostUpdate{HostID: *device.MonitorId}
	updateSchema.Interfaces = &hostInterfaces
	_, err = client.HostUpdate(&updateSchema)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[snmpCredDeleteHooks]: update host failed for cred %s", cred.Id), zap.Error(err))
		return
	}
	logger.Logger.Info(fmt.Sprintf("[snmpCredDeleteHooks]: update host success for cred %s", cred.Id))
}

func ServerSnmpCredDeleteHooks(cred *models.ServerSnmpCredential) {
	globalCred, err := gen.ServerSnmpCredential.Where(
		gen.ServerSnmpCredential.OrganizationId.Eq(cred.OrganizationId),
		gen.ServerSnmpCredential.ServerId.IsNull()).First()
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[serverSnmpCredDeleteHooks]: get server snmp global cred failed with cred %s", cred.Id), zap.Error(err))
		return
	}
	server, err := gen.Server.Where(gen.Server.Id.Eq(*cred.ServerId), gen.Server.MonitorId.IsNotNull()).First()
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[serverSnmpCredDeleteHooks]: get server failed with cred %s", cred.Id), zap.Error(err))
		return
	}

	client := zbx.NewZbxClient()
	zbxInterfaces, err := client.HostInterfaceGet(&zschema.HostInterfaceGet{
		HostIDs: []string{*server.MonitorId},
	})
	if err != nil || len(zbxInterfaces) <= 0 {
		logger.Logger.Error(fmt.Sprintf("[serverSnmpCredDeleteHooks]: failed to update delete due to get host interface failed for server %s", server.Id), zap.Error(err))
		return
	}
	hostInterfaceId := zbxInterfaces[0].InterfaceId
	hostInterfaces := make([]zschema.HostInterfaceUpdate, 0)
	hostInterfaces = append(hostInterfaces, zschema.HostInterfaceUpdate{
		InterfaceID: hostInterfaceId,
		Details: &zschema.Details{
			Community:      globalCred.Community,
			MaxRepetitions: &globalCred.MaxRepetitions,
		},
	})
	updateSchema := zschema.HostUpdate{HostID: *server.MonitorId}
	updateSchema.Interfaces = &hostInterfaces
	_, err = client.HostUpdate(&updateSchema)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[serverSnmpCredDeleteHooks]: update host failed for cred %s", cred.Id), zap.Error(err))
		return
	}
	logger.Logger.Info(fmt.Sprintf("[serverSnmpCredDeleteHooks]: update host success for cred %s", cred.Id))
}
