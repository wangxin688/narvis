package hooks

import (
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/pkg/zbx"
	"github.com/wangxin688/narvis/server/pkg/zbx/zschema"
	"go.uber.org/zap"
)

func CreateZbxProxy(proxy *models.Proxy) {
	zapi := zbx.NewZbxClient()
	res, err := zapi.ProxyCreate(&zschema.ProxyCreate{
		Name:           proxy.Id,
		OperatingMode:  0,
		TlsConnect:     1,
		TlsPskIDentity: proxy.Id,
		TlsAccept:      2,
		TlsPsk:         core.Settings.Jwt.PublicAuthKey,
	})
	if err != nil {
		core.Logger.Error("[proxyCreateHooks]: failed to create proxy", zap.Error(err))
		return
	}
	proxy.ProxyId = &res
	err = gen.Proxy.Save(proxy)
	if err != nil {
		core.Logger.Error("[proxyCreateHooks]: failed to update proxyId to database", zap.Error(err))
	}
	core.Logger.Info("[proxyCreateHooks]: proxy created", zap.String("id", res))
}
