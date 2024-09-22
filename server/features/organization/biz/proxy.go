package biz

import (
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/global"
)

type ProxyService struct{}

func NewProxyService() *ProxyService {
	return &ProxyService{}
}

func (p *ProxyService) VerifyProxy(proxyId string) bool {
	proxy, err := gen.Proxy.Select(
		gen.Proxy.Id, gen.Proxy.Active, gen.Proxy.OrganizationId,
	).Where(gen.Proxy.Id.Eq(proxyId)).First()
	if err != nil {
		return false
	}
	if !proxy.Active {
		return false
	}
	global.OrganizationId.Set(proxy.OrganizationId)
	return true
}
