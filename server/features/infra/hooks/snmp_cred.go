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

func SnmpCredCreateHooks(credId string) {
	cred, err := gen.SnmpV2Credential.Where(gen.SnmpV2Credential.Id.Eq(credId)).Preload(gen.SnmpV2Credential.Organization).First()
	if err != nil {
		core.Logger.Error(fmt.Sprintf("[snmpCredCreateHooks]: get snmp cred failed with cred %s", credId), zap.Error(err))
		return
	}
	// create a global macro for given snmp cred
	client := zbx.NewZbxClient()
	if cred.DeviceId == nil || *cred.DeviceId == "" {
		globalMacro := zschema.Macro{
			Macro: getGlobalCommunityMacroName(cred.Organization.EnterpriseCode),
			Value: cred.Community,
		}
		_, err := client.UserMacroCreateGlobal(&globalMacro)
		if err != nil {
			core.Logger.Error(fmt.Sprintf("[snmpCredCreateHooks]: create global macro failed with cred %s", credId), zap.Error(err))
			return
		}
		core.Logger.Info(fmt.Sprintf("[snmpCredCreateHooks]: create global macro success with cred %s", credId))
		return
	}

	// device, err := gen.Device.Where(gen.Device.Id.Eq(*cred.DeviceId)).First()
	// if err != nil {
	// 	core.Logger.Error(fmt.Sprintf("[snmpCredCreateHooks]: get device failed with cred %s", credId), zap.Error(err))
	// 	return
	// }
	// globalCred, err := gen.SnmpV2Credential.Where(gen.SnmpV2Credential.DeviceId.IsNull()).First()
	// if err != nil {
	// 	core.Logger.Error(fmt.Sprintf("[snmpCredCreateHooks]: get global cred failed with cred %s", credId), zap.Error(err))
	// 	return
	// }
}

func SnmpCredUpdateHooks(credId string, diff map[string]*ts.OrmDiff) {

}

func SnmpCredDeleteHooks(cred *models.SnmpV2Credential) {}
