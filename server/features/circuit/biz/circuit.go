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

func (c *CircuitService) GetDeviceSiteIDByInterfaceID(interfaceID string) (deviceID string, siteID string, err error) {

	di, err := gen.DeviceInterface.Select(gen.DeviceInterface.DeviceID).Where(gen.DeviceInterface.ID.Eq(interfaceID)).First()
	if err != nil {
		return "", "", err
	}
	site, err := gen.Device.Select(gen.Device.SiteID).Where(gen.Device.ID.Eq(di.DeviceID)).First()
	if err != nil {
		return "", "", err
	}
	return di.DeviceID, site.SiteID, nil
}

func (c *CircuitService) CreateCircuit(circuit *schemas.CircuitCreate) (string, error) {

	circuit, err := c.validateCreateCircuit(circuit)
	if err != nil {
		return "", err
	}
	newCircuit := &models.Circuit{
		Name:        circuit.Name,
		CID:         circuit.CID,
		Status:      circuit.Status,
		BandWidth:   circuit.BandWidth,
		IpAddress:   circuit.IpAddress,
		Description: circuit.Description,
		CircuitType: circuit.CircuitType,
		ProviderID:  circuit.ProviderID,
	}
	aSiteID, aDeviceID, err := c.GetDeviceSiteIDByInterfaceID(circuit.AInterfaceID)
	if err != nil {
		return "", err
	}
	newCircuit.ASiteID = aSiteID
	newCircuit.ADeviceID = aDeviceID
	newCircuit.AInterfaceID = circuit.AInterfaceID
	if circuit.ZInterfaceID != nil {
		zSiteID, zDeviceID, err := c.GetDeviceSiteIDByInterfaceID(*circuit.ZInterfaceID)
		if err != nil {
			return "", err
		}
		if zDeviceID == aDeviceID {
			return "", errors.NewError(errors.CodeCircuitSameDevice, errors.MsgCircuitSameDevice)
		}
		newCircuit.ZSiteID = zSiteID
		newCircuit.ZDeviceID = zDeviceID
		newCircuit.ZInterfaceID = *circuit.ZInterfaceID
	}

	err = gen.Circuit.Create(newCircuit)
	if err != nil {
		return "", err
	}
	return newCircuit.ID, nil
}

func (c *CircuitService) UpdateCircuit(circuitID string, circuit *schemas.CircuitUpdate) error {
	updateFields := make(map[string]any)
	if circuit.Name != nil {
		updateFields["name"] = *circuit.Name
	}
	if circuit.CID != nil {
		updateFields["cid"] = *circuit.CID
	}
	if circuit.Status != nil {
		updateFields["status"] = *circuit.Status
	}
	if circuit.CircuitType != nil {
		updateFields["circuit_type"] = *circuit.CircuitType
	}
	if circuit.BandWidth != nil {
		updateFields["band_width"] = *circuit.BandWidth
	}
	if circuit.IpAddress != nil {
		updateFields["ip_address"] = *circuit.IpAddress
	}
	if circuit.Description != nil {
		updateFields["description"] = *circuit.Description
	}
	if circuit.ProviderID != nil {
		updateFields["provider_id"] = *circuit.ProviderID
	}
	if circuit.ZInterfaceID != nil {
		updateFields["z_interface_id"] = *circuit.ZInterfaceID
	}
	if circuit.AInterfaceID != nil {
		updateFields["a_interface_id"] = *circuit.AInterfaceID
	}
	_, err := gen.Circuit.Where(gen.Circuit.ID.Eq(circuitID), gen.Circuit.OrganizationId.Eq(global.OrganizationId.Get())).Updates(updateFields)
	if err != nil {
		return err
	}
	return nil
}

func (c *CircuitService) GetCircuitByID(id string) (*schemas.Circuit, error) {
	circuit, err := gen.Circuit.
		Where(gen.Circuit.ID.Eq(id), gen.Circuit.OrganizationId.Eq(global.OrganizationId.Get())).
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
		ID:          circuit.ID,
		Name:        circuit.Name,
		CID:         circuit.CID,
		Status:      circuit.Status,
		BandWidth:   circuit.BandWidth,
		IpAddress:   circuit.IpAddress,
		Description: circuit.Description,
		CircuitType: circuit.CircuitType,
		Provider:    schemas.ProviderShort{ID: circuit.ProviderID, Name: circuit.Provider.Name},
		CreatedAt:   circuit.CreatedAt,
		UpdatedAt:   circuit.UpdatedAt,
	}, nil
}

func (c *CircuitService) DeleteCircuit(id string) error {
	_, err := gen.Circuit.Where(gen.Circuit.ID.Eq(id), gen.Circuit.OrganizationId.Eq(global.OrganizationId.Get())).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (c *CircuitService) validateCreateCircuit(circuit *schemas.CircuitCreate) (*schemas.CircuitCreate, error) {
	if circuit.AInterfaceID == "" {
		return nil, errors.NewError(errors.CodeCircuitAInterfaceMissing, errors.MsgCircuitAInterfaceMissing)
	}

	if circuit.ZInterfaceID == nil {
		if circuit.AInterfaceID == *circuit.ZInterfaceID {
			return nil, errors.NewError(errors.CodeCircuitSameInterface, errors.MsgCircuitSameInterface)
		}
	}

	if circuit.CircuitType == "Intranet" && circuit.ZInterfaceID == nil {
		return nil, errors.NewError(errors.CodeCircuitZInterfaceMissing, errors.MsgCircuitZInterfaceMissing)
	}
	return circuit, nil
}

func (c *CircuitService) validateUpdateCircuit(circuit *schemas.CircuitUpdate, dbCircuit *models.Circuit) (*schemas.CircuitUpdate, error) {
	if circuit.AInterfaceID != nil && circuit.ZInterfaceID != nil {
		if *circuit.AInterfaceID == *circuit.ZInterfaceID {
			return nil, errors.NewError(errors.CodeCircuitSameInterface, errors.MsgCircuitSameInterface)
		}
		aDeviceID, _, err := c.GetDeviceSiteIDByInterfaceID(*circuit.AInterfaceID)
		if err != nil {
			return nil, err
		}
		zDeviceID, _, err := c.GetDeviceSiteIDByInterfaceID(*circuit.ZInterfaceID)
		if err != nil {
			return nil, err
		}
		if aDeviceID == zDeviceID {
			return nil, errors.NewError(errors.CodeCircuitSameDevice, errors.MsgCircuitSameDevice)
		}
	}

	if dbCircuit.CircuitType == "Internet" {
		if circuit.CircuitType != nil && *circuit.CircuitType != "Internet" {
			if circuit.ZInterfaceID != nil {
				return nil, errors.NewError(errors.CodeCircuitZInterfaceMissing, errors.MsgCircuitZInterfaceMissing)
			}
			if *circuit.CircuitType == "Internet" && circuit.ZInterfaceID != nil {
				return nil, errors.NewError(errors.CodeCircuitZInterfaceNotAllow, errors.MsgCircuitZInterfaceNotAllow)
			}
		}
	}

	return circuit, nil
}
