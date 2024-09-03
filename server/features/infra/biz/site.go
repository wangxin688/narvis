package biz

import (
	"github.com/wangxin688/narvis/intend/devicerole"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"
)

type SiteService struct{}

func NewSiteService() *SiteService {
	return &SiteService{}
}

func (s *SiteService) Create(site schemas.SiteCreate) (string, error) {
	newSite := &models.Site{
		Name:        site.Name,
		SiteCode:    site.SiteCode,
		Status:      site.Status,
		Region:      site.Region,
		TimeZone:    site.TimeZone,
		Latitude:    site.Latitude,
		Longitude:   site.Longitude,
		Address:     site.Address,
		Description: site.Description,
	}
	err := gen.Site.Create(newSite)
	if err != nil {
		return "", err
	}
	return newSite.Id, nil
}

func (s *SiteService) Update(Id string, site *schemas.SiteUpdate) error {
	updateFields := make(map[string]any)
	if site.Name != nil {
		updateFields["name"] = site.Name
	}
	if site.SiteCode != nil {
		updateFields["siteCode"] = site.SiteCode
	}
	if site.Region != nil {
		updateFields["region"] = site.Region
	}
	if site.TimeZone != nil {
		updateFields["timeZone"] = site.TimeZone
	}
	if site.Latitude != nil {
		updateFields["latitude"] = site.Latitude
	}
	if site.Longitude != nil {
		updateFields["longitude"] = site.Longitude
	}
	if site.Address != nil {
		updateFields["address"] = site.Address
	}
	if site.Description != nil {
		updateFields["description"] = site.Description
	}
	_, err := gen.Site.Select(gen.Site.Id.Eq(Id), gen.Site.OrganizationId.Eq(global.OrganizationId.Get())).Updates(updateFields)
	if err != nil {
		return err
	}
	return nil
}

func (s *SiteService) Delete(Id string) error {
	_, err := gen.Site.Select(gen.Site.Id.Eq(Id), gen.Site.OrganizationId.Eq(global.OrganizationId.Get())).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (s *SiteService) GetById(Id string) (schemas.Site, error) {

	site, err := gen.Site.Select(gen.Site.Id.Eq(Id), gen.Site.OrganizationId.Eq(global.OrganizationId.Get())).First()
	if err != nil {
		return schemas.Site{}, err
	}
	return schemas.Site{
		Id:          site.Id,
		CreatedAt:   site.CreatedAt,
		UpdatedAt:   site.UpdatedAt,
		Name:        site.Name,
		SiteCode:    site.SiteCode,
		Status:      site.Status,
		Region:      site.Region,
		TimeZone:    site.TimeZone,
		Latitude:    site.Latitude,
		Longitude:   site.Longitude,
		Address:     site.Address,
		Description: site.Description,
	}, nil
}

func (s *SiteService) GetSiteDetail(siteId string) (*schemas.SiteDetail, error) {
	site, err := s.GetById(siteId)
	if err != nil {
		return &schemas.SiteDetail{}, err
	}
	switchCount, err := s.GetSwitchCount(siteId)
	if err != nil {
		return &schemas.SiteDetail{}, err
	}
	apCount, err := s.GetApCount(siteId)
	if err != nil {
		return &schemas.SiteDetail{}, err
	}
	rackCount, err := s.GetRackCount(siteId)
	if err != nil {
		return &schemas.SiteDetail{}, err
	}
	circuitCount, err := s.GetCircuitCount(siteId)
	if err != nil {
		return &schemas.SiteDetail{}, err
	}
	vlanCount, err := s.GetVlanCount(siteId)
	if err != nil {
		return &schemas.SiteDetail{}, err
	}
	gatewayCount, err := s.GetGatewayCount(siteId)
	if err != nil {
		return &schemas.SiteDetail{}, err
	}
	return &schemas.SiteDetail{
		Site:         site,
		SwitchCount:  switchCount,
		ApCount:      apCount,
		RackCount:    rackCount,
		CircuitCount: circuitCount,
		VlanCount:    vlanCount,
		GatewayCount: gatewayCount,
	}, nil
}

func (s *SiteService) GetSwitchCount(siteId string) (int64, error) {
	swDeviceRoles := devicerole.GetSwitchDeviceRole()
	switchDeviceRoles := make([]string, 0)
	for _, swDeviceRole := range *swDeviceRoles {
		switchDeviceRoles = append(switchDeviceRoles, string(swDeviceRole))
	}
	return gen.Device.Where(gen.Device.DeviceRole.In(switchDeviceRoles...), gen.Device.SiteId.Eq(siteId)).Count()
}

func (s *SiteService) GetApCount(siteId string) (int64, error) {
	return gen.AP.Where(gen.AP.SiteId.Eq(siteId), gen.AP.OrganizationId.Eq(global.OrganizationId.Get())).Count()
}

func (s *SiteService) GetRackCount(siteId string) (int64, error) {
	return gen.Rack.Where(gen.Rack.SiteId.Eq(siteId), gen.Rack.OrganizationId.Eq(global.OrganizationId.Get())).Count()
}

func (s *SiteService) GetCircuitCount(siteId string) (int64, error) {
	return gen.Circuit.Where(gen.Circuit.SiteId.Eq(siteId), gen.Circuit.OrganizationId.Eq(global.OrganizationId.Get())).Count()
}

func (s *SiteService) GetVlanCount(siteId string) (int64, error) {
	return gen.Vlan.Where(gen.Vlan.SiteId.Eq(siteId), gen.Vlan.OrganizationId.Eq(global.OrganizationId.Get())).Count()
}

func (s *SiteService) GetGatewayCount(siteId string) (int64, error) {
	gatewayRoles := devicerole.GetGatewayRole()
	gatewayDeviceRoles := make([]string, 0)
	for _, gatewayRole := range *gatewayRoles {
		gatewayDeviceRoles = append(gatewayDeviceRoles, string(gatewayRole))
	}
	return gen.Device.Where(gen.Device.DeviceRole.In(gatewayDeviceRoles...), gen.Device.SiteId.Eq(siteId)).Count()
}

func (s *SiteService) GetList(params *schemas.SiteQuery) (int64, *schemas.SiteList, error) {
	stmt := gen.Site.Where(gen.Site.OrganizationId.Eq(global.OrganizationId.Get()))
	if params.Id != nil {
		stmt = stmt.Where(gen.Site.Id.In(*params.Id...))
	}
	if params.Name != nil {
		stmt = stmt.Where(gen.Site.Name.In(*params.Name...))
	}
	if params.SiteCode != nil {
		stmt = stmt.Where(gen.Site.SiteCode.In(*params.SiteCode...))
	}
	if params.Status != nil {
		stmt = stmt.Where(gen.Site.Status.Eq(*params.Status))
	}
	if params.Region != nil {
		stmt = stmt.Where(gen.Site.Region.In(*params.Region...))
	}
	if params.Keyword != nil {
		stmt.UnderlyingDB().Scopes(params.Search(models.SiteSearchFields))
	}

	count, err := stmt.Count()
	if err != nil {
		return 0, nil, err
	}
	if count <= 0 {
		return 0, nil, nil
	}
	stmt.UnderlyingDB().Scopes(params.OrderByField())
	stmt.UnderlyingDB().Scopes(params.LimitOffset())
	sites, err := stmt.Find()
	if err != nil {
		return 0, nil, err
	}
	var res schemas.SiteList
	for _, site := range sites {
		res = append(res, schemas.Site{
			Id:          site.Id,
			CreatedAt:   site.CreatedAt,
			UpdatedAt:   site.UpdatedAt,
			Name:        site.Name,
			SiteCode:    site.SiteCode,
			Status:      site.Status,
			Region:      site.Region,
			TimeZone:    site.TimeZone,
			Latitude:    site.Latitude,
			Longitude:   site.Longitude,
			Address:     site.Address,
			Description: site.Description,
		})
	}
	return count, &res, nil
}

func (s *SiteService) GetCircuitBySites(sites []string) (*map[string]*schemas.CircuitShort, error) {
	circuits, err := gen.Circuit.Select(
		gen.Circuit.Id,
		gen.Circuit.Name,
		gen.Circuit.Provider,
		gen.Circuit.RxBandWidth,
		gen.Circuit.TxBandWidth,
	).Where(
		gen.Circuit.OrganizationId.Eq(global.OrganizationId.Get()),
		gen.Circuit.SiteId.In(sites...),
	).Find()
	if err != nil {
		return nil, err
	}
	results := make(map[string]*schemas.CircuitShort)
	for _, circuit := range circuits {
		results[circuit.Id] = &schemas.CircuitShort{
			Id:          circuit.Id,
			Name:        circuit.Name,
			Provider:    circuit.Provider,
			RxBandWidth: circuit.RxBandWidth,
			TxBandWidth: circuit.TxBandWidth,
		}
	}
	return &results, nil
}

func (s *SiteService) GetDeviceCountBySites(sites []string) (*map[string]int64, error) {
	var results []struct {
		SiteId string
		Count  int64
	}
	err := gen.Device.Select(gen.Device.SiteId, gen.Device.Id.Count().As("count")).
		Where(gen.Device.OrganizationId.Eq(global.OrganizationId.Get()), gen.Device.SiteId.In(sites...)).
		Group(gen.Device.SiteId).Scan(&results)
	if err != nil {
		return nil, err
	}
	res := make(map[string]int64)
	for _, result := range results {
		res[result.SiteId] = result.Count
	}
	return &res, nil
}
