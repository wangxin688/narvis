package infra_biz

import (
	"github.com/samber/lo"
	"github.com/wangxin688/narvis/intend/devicerole"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"
	ts "github.com/wangxin688/narvis/server/tools/schemas"
	"gorm.io/gorm"
)

type SiteService struct{}

func NewSiteService() *SiteService {
	return &SiteService{}
}

func (s *SiteService) Create(site schemas.SiteCreate) (string, error) {
	newSite := &models.Site{
		Name:           site.Name,
		SiteCode:       site.SiteCode,
		Status:         site.Status,
		Region:         site.Region,
		TimeZone:       site.TimeZone,
		Latitude:       site.Latitude,
		Longitude:      site.Longitude,
		Address:        site.Address,
		Description:    site.Description,
		OrganizationId: global.OrganizationId.Get(),
	}
	err := gen.Site.Create(newSite)
	if err != nil {
		return "", err
	}
	return newSite.Id, nil
}

func (s *SiteService) Update(Id string, site *schemas.SiteUpdate) (diff map[string]map[string]*ts.OrmDiff, err error) {
	dbSite, err := gen.Site.Where(gen.Site.Id.Eq(Id), gen.Site.OrganizationId.Eq(global.OrganizationId.Get())).First()
	if err != nil {
		return nil, err
	}
	updateFields := make(map[string]*ts.OrmDiff)
	if site.Name != nil && *site.Name != dbSite.Name {
		updateFields["name"] = &ts.OrmDiff{Before: dbSite.Name, After: *site.Name}
		dbSite.Name = *site.Name
	}
	if site.SiteCode != nil && *site.SiteCode != dbSite.SiteCode {
		updateFields["siteCode"] = &ts.OrmDiff{Before: dbSite.SiteCode, After: *site.SiteCode}
		dbSite.SiteCode = *site.SiteCode
	}
	if site.Region != nil && *site.Region != dbSite.Region {
		updateFields["region"] = &ts.OrmDiff{Before: dbSite.Region, After: *site.Region}
		dbSite.Region = *site.Region
	}
	if site.TimeZone != nil && *site.TimeZone != dbSite.TimeZone {
		updateFields["timeZone"] = &ts.OrmDiff{Before: dbSite.TimeZone, After: *site.TimeZone}
		dbSite.TimeZone = *site.TimeZone
	}
	if site.Latitude != nil && *site.Latitude != dbSite.Latitude {
		updateFields["latitude"] = &ts.OrmDiff{Before: dbSite.Latitude, After: *site.Latitude}
		dbSite.Latitude = *site.Latitude
	}
	if site.Longitude != nil && *site.Longitude != dbSite.Longitude {
		updateFields["longitude"] = &ts.OrmDiff{Before: dbSite.Longitude, After: *site.Longitude}
		dbSite.Longitude = *site.Longitude
	}
	if site.Address != nil && *site.Address != dbSite.Address {
		updateFields["address"] = &ts.OrmDiff{Before: dbSite.Address, After: *site.Address}
		dbSite.Address = *site.Address
	}
	if site.Description != nil && site.Description != dbSite.Description {
		updateFields["description"] = &ts.OrmDiff{Before: dbSite.Description, After: *site.Description}
		dbSite.Description = site.Description
	}
	if site.Status != nil && *site.Status != dbSite.Status {
		updateFields["status"] = &ts.OrmDiff{Before: dbSite.Status, After: *site.Status}
		dbSite.Status = *site.Status
	}
	diffValue := make(map[string]map[string]*ts.OrmDiff)
	diffValue[Id] = updateFields
	global.OrmDiff.Set(diffValue)
	if len(updateFields) == 0 {
		return nil, nil
	}
	err = gen.Site.UnderlyingDB().Transaction(func(tx *gorm.DB) error {
		err := gen.Site.UnderlyingDB().Save(dbSite).Error
		if err != nil {
			return err
		}
		if site.Status != nil && *site.Status == "Inactive" {
			_, err := gen.Device.Where(gen.Device.SiteId.Eq(Id), gen.Device.OrganizationId.Eq(global.OrganizationId.Get())).UpdateColumn(gen.Device.Status, "Inactive")
			if err != nil {
				return err
			}
			_, err = gen.AP.Where(gen.AP.SiteId.Eq(Id), gen.AP.OrganizationId.Eq(global.OrganizationId.Get())).UpdateColumn(gen.AP.Status, "Inactive")
			if err != nil {
				return err
			}
			_, err = gen.Circuit.Where(gen.Circuit.SiteId.Eq(Id), gen.Circuit.OrganizationId.Eq(global.OrganizationId.Get())).UpdateColumn(gen.Circuit.Status, "Inactive")
			if err != nil {
				return err
			}
		}
		_, err = gen.Site.Where(gen.Site.Id.Eq(Id), gen.Site.OrganizationId.Eq(global.OrganizationId.Get())).Updates(updateFields)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return diffValue, nil
}

func (s *SiteService) Delete(Id string) (*models.Site, error) {
	site, err := gen.Site.Where(gen.Site.Id.Eq(Id), gen.Site.OrganizationId.Eq(global.OrganizationId.Get())).First()
	if err != nil {
		return nil, err
	}

	_, err = gen.Site.Delete(site)
	if err != nil {
		return nil, err
	}
	return site, nil
}

func (s *SiteService) GetById(Id string) (schemas.Site, error) {

	site, err := gen.Site.Where(gen.Site.Id.Eq(Id), gen.Site.OrganizationId.Eq(global.OrganizationId.Get())).First()
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
	siteCircuits, err := s.GetCircuitBySites([]string{siteId})
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
		Circuit:      (*siteCircuits)[siteId],
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
	// return gen.Vlan.Where(gen.Vlan.SiteId.Eq(siteId), gen.Vlan.OrganizationId.Eq(global.OrganizationId.Get())).Count()
	return 0, nil
}

func (s *SiteService) GetGatewayCount(siteId string) (int64, error) {
	gatewayRoles := devicerole.GetGatewayRole()
	gatewayDeviceRoles := make([]string, 0)
	for _, gatewayRole := range *gatewayRoles {
		gatewayDeviceRoles = append(gatewayDeviceRoles, string(gatewayRole))
	}
	return gen.Device.Where(gen.Device.DeviceRole.In(gatewayDeviceRoles...), gen.Device.SiteId.Eq(siteId)).Count()
}

func (s *SiteService) GetList(params *schemas.SiteQuery) (int64, *[]*schemas.SiteResponse, error) {
	res := make([]*schemas.SiteResponse, 0)
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
	if params.IsSearchable() {
		keyword := "%" + *params.Keyword + "%"
		stmt = stmt.Where(gen.Site.Name.Like(keyword)).Or(
			gen.Site.SiteCode.Like(keyword),
		).Or(
			gen.Site.Region.Like(keyword),
		).Or(
			gen.Site.Address.Like(keyword),
		)
	}
	count, err := stmt.Count()
	if err != nil || count <= 0 {
		return 0, &res, err
	}
	stmt.UnderlyingDB().Scopes(params.OrderByField())
	stmt.UnderlyingDB().Scopes(params.Pagination())
	sites, err := stmt.Find()
	if err != nil {
		return 0, &res, err
	}
	siteIds := make([]string, 0)
	for _, site := range sites {
		siteIds = append(siteIds, site.Id)
	}
	deviceCount, err := s.GetDeviceApTotalBySites(siteIds)
	if err != nil {
		return 0, nil, err
	}
	circuits, err := s.GetCircuitBySites(siteIds)
	if err != nil {
		return 0, nil, err
	}

	for _, site := range sites {
		res = append(res, &schemas.SiteResponse{
			Site: schemas.Site{
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
			},
			DeviceCount: (*deviceCount)[site.Id],
			Circuit:     (*circuits)[site.Id],
		})
	}
	return count, &res, nil
}

func (s *SiteService) GetCircuitBySites(sites []string) (*map[string][]*schemas.CircuitShort, error) {
	// allocate memory first incase json Unmarshal nil slice to nil value
	results := make(map[string][]*schemas.CircuitShort)
	for _, siteId := range sites {
		if _, ok := results[siteId]; !ok {
			results[siteId] = make([]*schemas.CircuitShort, 0)
		}
	}
	circuits, err := gen.Circuit.Select(
		gen.Circuit.Id,
		gen.Circuit.Name,
		gen.Circuit.Provider,
		gen.Circuit.RxBandWidth,
		gen.Circuit.TxBandWidth,
		gen.Circuit.SiteId,
	).Where(
		gen.Circuit.OrganizationId.Eq(global.OrganizationId.Get()),
		gen.Circuit.SiteId.In(sites...),
	).Find()
	if err != nil {
		return nil, err
	}
	for _, circuit := range circuits {
		results[circuit.SiteId] = append(results[circuit.SiteId], &schemas.CircuitShort{
			Id:          circuit.Id,
			Name:        circuit.Name,
			Provider:    circuit.Provider,
			RxBandWidth: circuit.RxBandWidth,
			TxBandWidth: circuit.TxBandWidth,
		})
	}
	return &results, nil
}

func (s *SiteService) GetDeviceCountBySites(sites []string) (*map[string]int64, error) {
	var results []struct {
		SiteId string
		Count  int64
	}
	err := gen.Device.Select(gen.Device.SiteId.As("SiteId"), gen.Device.Id.Count().As("Count")).
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

func (s *SiteService) GetApCountBySites(sites []string) (*map[string]int64, error) {
	var results []struct {
		SiteId string
		Count  int64
	}
	err := gen.AP.Select(gen.AP.SiteId.As("SiteId"), gen.AP.Id.Count().As("Count")).
		Where(gen.AP.OrganizationId.Eq(global.OrganizationId.Get()), gen.AP.SiteId.In(sites...)).
		Group(gen.AP.SiteId).Scan(&results)
	if err != nil {
		return nil, err
	}
	res := make(map[string]int64)
	for _, result := range results {
		res[result.SiteId] = result.Count
	}
	return &res, nil
}

func (s *SiteService) GetDeviceApTotalBySites(sites []string) (*map[string]int64, error) {
	deviceCount, err := s.GetDeviceCountBySites(sites)
	if err != nil {
		return nil, err
	}
	apCount, err := s.GetApCountBySites(sites)
	if err != nil {
		return nil, err
	}
	res := make(map[string]int64)
	for siteId, count := range *deviceCount {
		res[siteId] = count
	}
	for siteId, count := range *apCount {
		res[siteId] += count
	}
	return &res, nil
}

func (s *SiteService) GetSiteShortMap(siteIds []string) (map[string]*schemas.SiteShort, error) {
	sites, err := gen.Site.Select(
		gen.Site.Id,
		gen.Site.Name,
		gen.Site.SiteCode,
	).Where(gen.Site.Id.In(siteIds...)).Find()
	if err != nil {
		return nil, err
	}
	res := make(map[string]*schemas.SiteShort)
	for _, site := range sites {
		res[site.Id] = &schemas.SiteShort{
			Id:       site.Id,
			Name:     site.Name,
			SiteCode: site.SiteCode,
		}
	}
	return res, nil
}

func (s *SiteService) GetAllActiveSites() ([]*models.Site, error) {

	activeTenants, err := gen.Organization.Select(gen.Organization.Id).Where(gen.Organization.Active.Is(true)).Find()

	if err != nil {
		return nil, err
	}
	tenantIds := lo.Map(activeTenants, func(item *models.Organization, _ int) string {
		return item.Id
	})
	sites := make([]*models.Site, 0)
	err = gen.Site.Select(gen.Site.Id, gen.Site.OrganizationId).
		Where(gen.Site.OrganizationId.In(tenantIds...), gen.Site.Status.Eq("Active")).Scan(&sites)
	if err != nil {
		return nil, err
	}
	return sites, nil

}
