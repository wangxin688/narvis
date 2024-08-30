package biz

import (
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/global"
)

type TopologyService struct {
}

func NewTopologyService() *TopologyService {
	return &TopologyService{}
}

func (s *TopologyService) GetLldpNeighbors(siteID string) {
	nbrs, err := gen.LLDPNeighbor.Where(gen.LLDPNeighbor.SiteID.Eq(siteID), gen.LLDPNeighbor.OrganizationId.Eq(global.OrganizationId.Get())).
		Preload(gen.LLDPNeighbor.Device).Select()

}
