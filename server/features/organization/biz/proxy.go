package biz

import "github.com/wangxin688/narvis/server/dal/gen"

type ProxyService struct {
	gen.IProxyDo
}

func NewProxyService(proxyDo gen.IProxyDo) *ProxyService {
	return &ProxyService{proxyDo}
}
