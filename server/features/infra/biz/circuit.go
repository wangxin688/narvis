package infra_biz

import (
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"
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
		Name:        circuit.Name,
		CId:         circuit.CId,
		Status:      circuit.Status,
		RxBandWidth: circuit.RxBandWidth,
		TxBandWidth: circuit.TxBandWidth,
		Ipv4Address: circuit.Ipv4Address,
		Description: circuit.Description,
		CircuitType: circuit.CircuitType,
		Provider:    circuit.Provider,
		OrganizationId: global.OrganizationId.Get(),
	}
	siteId, deviceId, err := c.GetDeviceSiteIdByInterfaceId(circuit.InterfaceId)
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

func (c *CircuitService) UpdateCircuit(circuitId string, circuit *schemas.CircuitUpdate) error {
	updateFields := make(map[string]any)
	if circuit.Name != nil {
		updateFields["name"] = *circuit.Name
	}
	if circuit.CId != nil {
		updateFields["cid"] = *circuit.CId
	}
	if circuit.Status != nil {
		updateFields["status"] = *circuit.Status
	}
	if circuit.CircuitType != nil {
		updateFields["circuitType"] = *circuit.CircuitType
	}
	if circuit.RxBandWidth != nil {
		updateFields["rxBandWidth"] = *circuit.RxBandWidth
	}
	if circuit.TxBandWidth != nil {
		updateFields["txBandWidth"] = *circuit.TxBandWidth
	}
	if circuit.Ipv4Address != nil {
		updateFields["ipv4Address"] = *circuit.Ipv4Address
	}
	if circuit.Ipv6Address != nil {
		updateFields["ipv6Address"] = *circuit.Ipv6Address
	}
	if circuit.Description != nil {
		updateFields["description"] = *circuit.Description
	}
	if circuit.Provider != nil {
		updateFields["provider"] = *circuit.Provider
	}
	if circuit.InterfaceId != nil {
		updateFields["InterfaceId"] = *circuit.InterfaceId
	}
	_, err := gen.Circuit.Where(gen.Circuit.Id.Eq(circuitId), gen.Circuit.OrganizationId.Eq(global.OrganizationId.Get())).Updates(updateFields)
	if err != nil {
		return err
	}
	return nil
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

func (c *CircuitService) DeleteCircuit(id string) error {
	_, err := gen.Circuit.Where(gen.Circuit.Id.Eq(id), gen.Circuit.OrganizationId.Eq(global.OrganizationId.Get())).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (c *CircuitService) ListCircuit(query *schemas.CircuitQuery) (int64, *schemas.CircuitList, error) {
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
	if query.Keyword != nil {
		stmt.UnderlyingDB().Scopes(query.Search(models.CircuitSearchFields))
	}

	count, err := stmt.Count()
	if err != nil || count < 0 {
		return 0, nil, err
	}

	stmt.UnderlyingDB().Scopes(query.OrderByField())
	stmt.UnderlyingDB().Scopes(query.LimitOffset())

	list, err := stmt.Find()
	if err != nil {
		return 0, nil, err
	}
	var res schemas.CircuitList

	for _, item := range list {
		res = append(res, schemas.Circuit{
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
