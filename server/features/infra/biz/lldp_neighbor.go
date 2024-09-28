package infra_biz

import (
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/models"
)

type LldpNeighborService struct {
}

func NewLldpNeighborService() *LldpNeighborService {
	return &LldpNeighborService{}
}

func (d *LldpNeighborService) GetDeviceLldpNeighbors(deviceId string) (map[string]*models.LLDPNeighbor, error) {
	lldpList, err := gen.LLDPNeighbor.Where(gen.LLDPNeighbor.LocalDeviceId.Eq(deviceId)).Find()
	if err != nil {
		return nil, err
	}
	lldpMap := make(map[string]*models.LLDPNeighbor)
	for _, lldp := range lldpList {
		lldpMap[lldp.HashValue] = lldp
	}
	return lldpMap, nil
}

func (d *LldpNeighborService) GetApLldpNeighbors(deviceId string) (map[string]*models.ApLLDPNeighbor, error) {
	lldpList, err := gen.ApLLDPNeighbor.Where(gen.ApLLDPNeighbor.LocalDeviceId.Eq(deviceId)).Find()
	if err != nil {
		return nil, err
	}
	lldpMap := make(map[string]*models.ApLLDPNeighbor)
	for _, lldp := range lldpList {
		lldpMap[lldp.HashValue] = lldp
	}
	return lldpMap, nil
}
