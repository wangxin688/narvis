package biz

import (
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"
)

type TopologyService struct {
}

func NewTopologyService() *TopologyService {
	return &TopologyService{}
}

func (s *TopologyService) GetTopologyDevices(deviceIds []string) (*map[string]*models.Device, error) {
	devices, err := gen.Device.Select(gen.Device.Id, gen.Device.Name, gen.Device.ManagementIp, gen.Device.DeviceRole).
		Where(gen.Device.Id.In(deviceIds...), gen.Device.OrganizationId.Eq(global.OrganizationId.Get())).
		Find()
	if err != nil {
		return nil, err
	}
	deviceMap := make(map[string]*models.Device)
	for _, device := range devices {
		deviceMap[device.Id] = device
	}
	return &deviceMap, nil
}

func (s *TopologyService) GetLldpNeighbors(siteId string) {

	neighbors, err := gen.LLDPNeighbor.Where(gen.LLDPNeighbor.SiteId.Eq(siteId), gen.LLDPNeighbor.OrganizationId.Eq(global.OrganizationId.Get())).Find()
	if err != nil {
		return
	}
	deviceIds := make([]string, 0)
	for _, neighbor := range neighbors {
		deviceIds = append(deviceIds, neighbor.LocalDeviceId)
		deviceIds = append(deviceIds, neighbor.RemoteDeviceId)
	}
	devices, err := s.GetTopologyDevices(deviceIds)
	if err != nil {
		return
	}

}
