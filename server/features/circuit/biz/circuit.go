package biz

import (
	"github.com/wangxin688/narvis/server/dal/gen"
	schemas "github.com/wangxin688/narvis/server/features/circuit/scheams"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/tools/errors"
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

	circuit, err := c.validateCreateCircuit(circuit)
	if err != nil {
		return "", err
	}
	newCircuit := &models.Circuit{
		Name:        circuit.Name,
		CId:         circuit.CId,
		Status:      circuit.Status,
		BandWidth:   circuit.BandWidth,
		IpAddress:   circuit.IpAddress,
		Description: circuit.Description,
		CircuitType: circuit.CircuitType,
		ProviderId:  circuit.ProviderId,
	}
	aSiteId, aDeviceId, err := c.GetDeviceSiteIdByInterfaceId(circuit.AInterfaceId)
	if err != nil {
		return "", err
	}
	newCircuit.ASiteId = aSiteId
	newCircuit.ADeviceId = aDeviceId
	newCircuit.AInterfaceId = circuit.AInterfaceId
	if circuit.ZInterfaceId != nil {
		zSiteId, zDeviceId, err := c.GetDeviceSiteIdByInterfaceId(*circuit.ZInterfaceId)
		if err != nil {
			return "", err
		}
		if zDeviceId == aDeviceId {
			return "", errors.NewError(errors.CodeCircuitSameDevice, errors.MsgCircuitSameDevice)
		}
		newCircuit.ZSiteId = zSiteId
		newCircuit.ZDeviceId = zDeviceId
		newCircuit.ZInterfaceId = *circuit.ZInterfaceId
	}

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
	if circuit.BandWidth != nil {
		updateFields["bandWidth"] = *circuit.BandWidth
	}
	if circuit.IpAddress != nil {
		updateFields["ipAddress"] = *circuit.IpAddress
	}
	if circuit.Description != nil {
		updateFields["description"] = *circuit.Description
	}
	if circuit.ProviderId != nil {
		updateFields["providerId"] = *circuit.ProviderId
	}
	if circuit.ZInterfaceId != nil {
		updateFields["zInterfaceId"] = *circuit.ZInterfaceId
	}
	if circuit.AInterfaceId != nil {
		updateFields["aInterfaceId"] = *circuit.AInterfaceId
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
		Preload(gen.Circuit.Provider).
		Preload(gen.Circuit.ASite).
		Preload(gen.Circuit.ADevice).
		Preload(gen.Circuit.AInterface).
		Preload(gen.Circuit.ZSite).
		Preload(gen.Circuit.ZDevice).
		Preload(gen.Circuit.ZInterface).
		First()
	if err != nil {
		return nil, err
	}
	return &schemas.Circuit{
		Id:          circuit.Id,
		Name:        circuit.Name,
		CId:         circuit.CId,
		Status:      circuit.Status,
		BandWidth:   circuit.BandWidth,
		IpAddress:   circuit.IpAddress,
		Description: circuit.Description,
		CircuitType: circuit.CircuitType,
		Provider:    schemas.ProviderShort{Id: circuit.ProviderId, Name: circuit.Provider.Name},
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

func (c *CircuitService) validateCreateCircuit(circuit *schemas.CircuitCreate) (*schemas.CircuitCreate, error) {
	if circuit.AInterfaceId == "" {
		return nil, errors.NewError(errors.CodeCircuitAInterfaceMissing, errors.MsgCircuitAInterfaceMissing)
	}

	if circuit.ZInterfaceId == nil {
		if circuit.AInterfaceId == *circuit.ZInterfaceId {
			return nil, errors.NewError(errors.CodeCircuitSameInterface, errors.MsgCircuitSameInterface)
		}
	}

	if circuit.CircuitType == "Intranet" && circuit.ZInterfaceId == nil {
		return nil, errors.NewError(errors.CodeCircuitZInterfaceMissing, errors.MsgCircuitZInterfaceMissing)
	}
	return circuit, nil
}

func (c *CircuitService) validateUpdateCircuit(circuit *schemas.CircuitUpdate, dbCircuit *models.Circuit) (*schemas.CircuitUpdate, error) {
	if circuit.AInterfaceId != nil && circuit.ZInterfaceId != nil {
		if *circuit.AInterfaceId == *circuit.ZInterfaceId {
			return nil, errors.NewError(errors.CodeCircuitSameInterface, errors.MsgCircuitSameInterface)
		}
		aDeviceId, _, err := c.GetDeviceSiteIdByInterfaceId(*circuit.AInterfaceId)
		if err != nil {
			return nil, err
		}
		zDeviceId, _, err := c.GetDeviceSiteIdByInterfaceId(*circuit.ZInterfaceId)
		if err != nil {
			return nil, err
		}
		if aDeviceId == zDeviceId {
			return nil, errors.NewError(errors.CodeCircuitSameDevice, errors.MsgCircuitSameDevice)
		}
	}

	if dbCircuit.CircuitType == "Internet" {
		if circuit.CircuitType != nil && *circuit.CircuitType != "Internet" {
			if circuit.ZInterfaceId != nil {
				return nil, errors.NewError(errors.CodeCircuitZInterfaceMissing, errors.MsgCircuitZInterfaceMissing)
			}
			if *circuit.CircuitType == "Internet" && circuit.ZInterfaceId != nil {
				return nil, errors.NewError(errors.CodeCircuitZInterfaceNotAllow, errors.MsgCircuitZInterfaceNotAllow)
			}
		}
	}

	return circuit, nil
}
