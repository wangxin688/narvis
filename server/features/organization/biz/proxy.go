package biz

import (
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/organization/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"
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
	global.ProxyId.Set(proxyId)
	return true
}

func (p *ProxyService) CreateProxy(proxy *schemas.ProxyCreate) (*models.Proxy, error) {
	newProxy := &models.Proxy{
		OrganizationId: proxy.OrganizationId,
		Name:           proxy.Name,
		Active:         proxy.Active,
	}

	err := gen.Proxy.Create(newProxy)
	if err != nil {
		return nil, err
	}
	return newProxy, nil
}
