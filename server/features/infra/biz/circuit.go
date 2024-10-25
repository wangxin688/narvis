package infra_biz

import (
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"
	ts "github.com/wangxin688/narvis/server/tools/schemas"
)

type CircuitService struct{}

func NewCircuitService() *CircuitService {
	return &CircuitService{}
}

func (c *CircuitService) GetDeviceSiteIdByInterfaceId(interfaceId string) (deviceId string, siteId string, err error) {

	di, err := gen.DeviceInterface.Select(gen.DeviceInterface.DeviceId).Where(gen.DeviceInterface.Id.Eq(interfaceId)).First()
	if err != nil {
		return "", "", err
	}
	site, err := gen.Device.Select(gen.Device.SiteId).Where(gen.Device.Id.Eq(di.DeviceId)).First()
	if err != nil {
		return "", "", err
	}
	return di.DeviceId, site.SiteId, nil
}

func (c *CircuitService) CreateCircuit(circuit *schemas.CircuitCreate) (string, error) {
	newCircuit := &models.Circuit{
		Name:           circuit.Name,
		CId:            circuit.CId,
		Status:         circuit.Status,
		RxBandWidth:    circuit.RxBandWidth,
		TxBandWidth:    circuit.TxBandWidth,
		Ipv4Address:    circuit.Ipv4Address,
		Description:    circuit.Description,
		CircuitType:    circuit.CircuitType,
		Provider:       circuit.Provider,
		OrganizationId: global.OrganizationId.Get(),
	}
	deviceId, siteId, err := c.GetDeviceSiteIdByInterfaceId(circuit.InterfaceId)
	if err != nil {
		return "", err
	}
	newCircuit.SiteId = siteId
	newCircuit.DeviceId = deviceId
	newCircuit.InterfaceId = circuit.InterfaceId
	err = gen.Circuit.Create(newCircuit)
	if err != nil {
		return "", err
	}
	return newCircuit.Id, nil
}

func (c *CircuitService) UpdateCircuit(circuitId string, circuit *schemas.CircuitUpdate) (diff map[string]map[string]*ts.OrmDiff, err error) {
	dbCircuit, err := gen.Circuit.Where(gen.Circuit.Id.Eq(circuitId), gen.Circuit.OrganizationId.Eq(global.OrganizationId.Get())).First()
	if err != nil {
		return nil, err
	}
	updateFields := make(map[string]*ts.OrmDiff)
	if circuit.Name != nil && *circuit.Name != dbCircuit.Name {
		updateFields["name"] = &ts.OrmDiff{Before: dbCircuit.Name, After: *circuit.Name}
		dbCircuit.Name = *circuit.Name
	}
	if circuit.CId != nil && *circuit.CId != *dbCircuit.CId {
		updateFields["cid"] = &ts.OrmDiff{Before: *dbCircuit.CId, After: *circuit.CId}
		dbCircuit.CId = circuit.CId
	}
	if circuit.Status != nil && *circuit.Status != dbCircuit.Status {
		updateFields["status"] = &ts.OrmDiff{Before: dbCircuit.Status, After: *circuit.Status}
		dbCircuit.Status = *circuit.Status
	}
	if circuit.CircuitType != nil && *circuit.CircuitType != dbCircuit.CircuitType {
		updateFields["circuitType"] = &ts.OrmDiff{Before: dbCircuit.CircuitType, After: *circuit.CircuitType}
		dbCircuit.CircuitType = *circuit.CircuitType
	}
	if circuit.RxBandWidth != nil && *circuit.RxBandWidth != dbCircuit.RxBandWidth {
		updateFields["rxBandWidth"] = &ts.OrmDiff{Before: dbCircuit.RxBandWidth, After: *circuit.RxBandWidth}
		dbCircuit.RxBandWidth = *circuit.RxBandWidth
	}
	if circuit.TxBandWidth != nil && *circuit.TxBandWidth != dbCircuit.TxBandWidth {
		updateFields["txBandWidth"] = &ts.OrmDiff{Before: dbCircuit.TxBandWidth, After: *circuit.TxBandWidth}
		dbCircuit.TxBandWidth = *circuit.TxBandWidth
	}
	if circuit.Ipv4Address != nil && *circuit.Ipv4Address != *dbCircuit.Ipv4Address {
		updateFields["ipv4Address"] = &ts.OrmDiff{Before: *dbCircuit.Ipv4Address, After: *circuit.Ipv4Address}
		dbCircuit.Ipv4Address = circuit.Ipv4Address
	}
	if circuit.Ipv6Address != nil && *circuit.Ipv6Address != *dbCircuit.Ipv6Address {
		updateFields["ipv6Address"] = &ts.OrmDiff{Before: *dbCircuit.Ipv6Address, After: *circuit.Ipv6Address}
		dbCircuit.Ipv6Address = circuit.Ipv6Address
	}
	if circuit.Description != nil && *circuit.Description != *dbCircuit.Description {
		updateFields["description"] = &ts.OrmDiff{Before: dbCircuit.Description, After: *circuit.Description}
		dbCircuit.Description = circuit.Description
	}
	if circuit.Provider != nil && *circuit.Provider != dbCircuit.Provider {
		updateFields["provider"] = &ts.OrmDiff{Before: dbCircuit.Provider, After: *circuit.Provider}
		dbCircuit.Provider = *circuit.Provider
	}
	if circuit.InterfaceId != nil && *circuit.InterfaceId != dbCircuit.InterfaceId {
		updateFields["InterfaceId"] = &ts.OrmDiff{Before: dbCircuit.InterfaceId, After: *circuit.InterfaceId}
		dbCircuit.InterfaceId = *circuit.InterfaceId
	}

	if len(updateFields) == 0 {
		return nil, nil
	}
	err = gen.Circuit.UnderlyingDB().Save(dbCircuit).Error
	if err != nil {
		return nil, err
	}
	diffValue := make(map[string]map[string]*ts.OrmDiff)
	diffValue[circuitId] = updateFields
	global.OrmDiff.Set(diffValue)
	return diffValue, nil
}

func (c *CircuitService) GetCircuitById(id string) (*schemas.Circuit, error) {
	circuit, err := gen.Circuit.
		Where(gen.Circuit.Id.Eq(id), gen.Circuit.OrganizationId.Eq(global.OrganizationId.Get())).
		Preload(gen.Circuit.Site).
		Preload(gen.Circuit.Device).
		Preload(gen.Circuit.DeviceInterface).
		First()
	if err != nil {
		return nil, err
	}
	return &schemas.Circuit{
		Id:          circuit.Id,
		Name:        circuit.Name,
		CId:         *circuit.CId,
		Status:      circuit.Status,
		RxBandWidth: circuit.RxBandWidth,
		TxBandWidth: circuit.TxBandWidth,
		Ipv4Address: circuit.Ipv4Address,
		Ipv6Address: circuit.Ipv6Address,
		Description: circuit.Description,
		CircuitType: circuit.CircuitType,
		Provider:    circuit.Provider,
		CreatedAt:   circuit.CreatedAt,
		UpdatedAt:   circuit.UpdatedAt,
	}, nil
}

func (c *CircuitService) DeleteCircuit(id string) (*models.Circuit, error) {
	dbCircuit, err := gen.Circuit.Where(gen.Circuit.Id.Eq(id), gen.Circuit.OrganizationId.Eq(global.OrganizationId.Get())).First()
	if err != nil {
		return nil, err
	}
	_, err = gen.Circuit.Delete(dbCircuit)
	if err != nil {
		return nil, err
	}
	return dbCircuit, nil
}

func (c *CircuitService) ListCircuit(query *schemas.CircuitQuery) (int64, *[]*schemas.Circuit, error) {
	res := make([]*schemas.Circuit, 0)
	stmt := gen.Circuit.Where(gen.Circuit.OrganizationId.Eq(global.OrganizationId.Get()))
	if query.Name != nil {
		stmt = stmt.Where(gen.Circuit.Name.In(*query.Name...))
	}
	if query.Status != nil {
		stmt = stmt.Where(gen.Circuit.Status.Eq(*query.Status))
	}
	if query.Provider != nil {
		stmt = stmt.Where(gen.Circuit.Provider.In(*query.Provider...))
	}
	if query.CircuitType != nil {
		stmt = stmt.Where(gen.Circuit.CircuitType.In(*query.CircuitType...))
	}
	if query.SiteId != nil {
		stmt = stmt.Where(gen.Circuit.SiteId.In(*query.SiteId...))
	}
	if query.DeviceId != nil {
		stmt = stmt.Where(gen.Circuit.DeviceId.In(*query.DeviceId...))
	}
	if query.InterfaceId != nil {
		stmt = stmt.Where(gen.Circuit.InterfaceId.In(*query.InterfaceId...))
	}
	if query.Ipv4Address != nil {
		stmt = stmt.Where(gen.Circuit.Ipv4Address.In(*query.Ipv4Address...))
	}
	if query.Ipv6Address != nil {
		stmt = stmt.Where(gen.Circuit.Ipv6Address.In(*query.Ipv6Address...))
	}
	if query.IsSearchable() {
		keyword := "%" + *query.Keyword + "%"
		stmt = stmt.Where(gen.Circuit.Name.Like(keyword)).Or(
			gen.Circuit.Ipv4Address.Like(keyword)).Or(
			gen.Circuit.Ipv6Address.Like(keyword))
	}

	count, err := stmt.Count()
	if err != nil || count < 0 {
		return 0, &res, err
	}

	stmt.UnderlyingDB().Scopes(query.OrderByField())
	stmt.UnderlyingDB().Scopes(query.Pagination())

	list, err := stmt.Find()
	if err != nil {
		return 0, &res, err
	}
	for _, item := range list {
		res = append(res, &schemas.Circuit{
			Id:          item.Id,
			Name:        item.Name,
			CId:         *item.CId,
			Status:      item.Status,
			RxBandWidth: item.RxBandWidth,
			TxBandWidth: item.TxBandWidth,
			Ipv4Address: item.Ipv4Address,
			Ipv6Address: item.Ipv6Address,
			Description: item.Description,
			CircuitType: item.CircuitType,
			Provider:    item.Provider,
			CreatedAt:   item.CreatedAt,
			UpdatedAt:   item.UpdatedAt,
		})
	}
	return count, &res, nil
}

func (c *CircuitService) GetCircuitShortMap(cIds []string) (map[string]*schemas.CircuitShort, error) {
	circuits, err := gen.Circuit.Select(
		gen.Circuit.Id,
		gen.Circuit.Name,
		gen.Circuit.RxBandWidth,
		gen.Circuit.TxBandWidth,
		gen.Circuit.Provider,
	).Where(gen.Circuit.Id.In(cIds...)).Find()
	if err != nil {
		return nil, err
	}

	res := make(map[string]*schemas.CircuitShort)
	for _, circuit := range circuits {
		res[circuit.Id] = &schemas.CircuitShort{
			Id:          circuit.Id,
			Name:        circuit.Name,
			RxBandWidth: circuit.RxBandWidth,
			TxBandWidth: circuit.TxBandWidth,
			Provider:    circuit.Provider,
		}
	}
	return res, nil
}

func (d *CircuitService) SearchCircuitByKeyword(keyword string, orgId string) ([]string, error) {
	result := make([]string, 0)
	stmt := gen.Circuit.Select(gen.Circuit.Id).Where(gen.Circuit.OrganizationId.Eq(orgId))
	keyword = "%" + keyword + "%"
	stmt = stmt.Where(gen.Circuit.Name.Like(keyword)).Or(
		gen.Circuit.Ipv4Address.Like(keyword)).Or(
		gen.Circuit.Ipv6Address.Like(keyword))
	err := stmt.Scan(&result)
	return result, err
}
