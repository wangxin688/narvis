package biz

import (
	"github.com/samber/lo"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/topology/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"
)

type TopologyService struct {
}

func NewTopologyService() *TopologyService {
	return &TopologyService{}
}

func (s *TopologyService) GetSiteTopology(siteId string) (*schemas.SiteTopology, error) {
	result := &schemas.SiteTopology{
		Nodes: make([]*schemas.Node, 0),
		Lines: make([]*schemas.Line, 0),
	}
	_, err := gen.Site.Select(gen.Site.Id).
		Where(gen.Site.Id.Eq(siteId), gen.Site.OrganizationId.Eq(global.OrganizationId.Get())).First()
	if err != nil {
		return nil, err
	}
	deviceLldpNeighbors, err := s.getLldpNeighbors(siteId)
	if err != nil {
		return nil, err
	}
	if len(deviceLldpNeighbors) == 0 {
		return result, nil
	}
	apLldpNeighbors, err := s.getApLldpNeighbors(siteId)
	if err != nil {
		return nil, err
	}
	devices, err := s.getLLdpDevices(deviceLldpNeighbors)
	if err != nil {
		return nil, err
	}
	aps, err := s.getLldpAps(apLldpNeighbors)
	if err != nil {
		return nil, err
	}
	deviceNodes, deviceLines := deviceLldpToGraph(deviceLldpNeighbors, devices)
	apNodes, apLines := apLldpToGraph(apLldpNeighbors, devices, aps)

	result.Nodes = append(result.Nodes, deviceNodes...)
	result.Nodes = append(result.Nodes, apNodes...)
	result.Lines = append(result.Lines, deviceLines...)
	result.Lines = append(result.Lines, apLines...)
	return result, nil

}

func (s *TopologyService) getTopologyDevices(deviceIds []string) (map[string]*models.Device, error) {
	devices, err := gen.Device.Select(gen.Device.Id, gen.Device.Name, gen.Device.ManagementIp, gen.Device.DeviceRole, gen.Device.Floor).
		Where(gen.Device.Id.In(deviceIds...), gen.Device.OrganizationId.Eq(global.OrganizationId.Get())).
		Find()
	if err != nil {
		return nil, err
	}
	deviceMap := make(map[string]*models.Device)
	for _, device := range devices {
		deviceMap[device.Id] = device
	}
	return deviceMap, nil
}

func (s *TopologyService) getTopologyAps(apIds []string) (map[string]*models.AP, error) {
	aps, err := gen.AP.Select(gen.AP.Id, gen.AP.Name, gen.AP.ManagementIp, gen.AP.DeviceRole, gen.AP.Floor).
		Where(gen.AP.Id.In(apIds...), gen.AP.OrganizationId.Eq(global.OrganizationId.Get())).Find()
	if err != nil {
		return nil, err
	}
	apMap := make(map[string]*models.AP)
	for _, ap := range aps {
		apMap[ap.Id] = ap
	}
	return apMap, nil
}

func (s *TopologyService) getLldpNeighbors(siteId string) ([]*models.LLDPNeighbor, error) {

	neighbors, err := gen.LLDPNeighbor.Where(gen.LLDPNeighbor.SiteId.Eq(siteId),
		gen.LLDPNeighbor.OrganizationId.Eq(global.OrganizationId.Get())).Find()
	if err != nil {
		return nil, err
	}
	if len(neighbors) == 0 {
		return neighbors, nil
	}
	neighbors = deduplicateEdges(neighbors)
	return neighbors, nil
}

func (s *TopologyService) getApLldpNeighbors(siteId string) ([]*models.ApLLDPNeighbor, error) {

	neighbors, err := gen.ApLLDPNeighbor.Where(gen.ApLLDPNeighbor.SiteId.Eq(siteId),
		gen.ApLLDPNeighbor.OrganizationId.Eq(global.OrganizationId.Get())).Find()
	if err != nil {
		return nil, err
	}
	return neighbors, nil
}

func (s *TopologyService) getLLdpDevices(lldp []*models.LLDPNeighbor) (map[string]*models.Device, error) {
	deviceIds := make(map[string]struct{})
	for _, neighbor := range lldp {
		if _, exists := deviceIds[neighbor.LocalDeviceId]; !exists {
			deviceIds[neighbor.LocalDeviceId] = struct{}{}
		}
		if _, exists := deviceIds[neighbor.RemoteDeviceId]; !exists {
			deviceIds[neighbor.RemoteDeviceId] = struct{}{}
		}
	}
	uniqueDeviceIds := lo.MapToSlice(deviceIds, func(key string, _ struct{}) string { return key })
	devices, err := s.getTopologyDevices(uniqueDeviceIds)
	if err != nil {
		return nil, err
	}
	return devices, nil
}

func (s *TopologyService) getLldpAps(lldp []*models.ApLLDPNeighbor) (map[string]*models.AP, error) {
	apIds := make(map[string]struct{})
	for _, neighbor := range lldp {
		if _, exists := apIds[neighbor.RemoteApId]; !exists {
			apIds[neighbor.RemoteApId] = struct{}{}
		}
	}
	uniqueApIds := lo.MapToSlice(apIds, func(key string, _ struct{}) string { return key })
	aps, err := s.getTopologyAps(uniqueApIds)
	if err != nil {
		return nil, err
	}
	return aps, nil
}


// TODO: add operational data to topology nodes
// TODO: add circuit node in topology