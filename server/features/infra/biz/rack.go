package biz

import (
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/tools/errors"
)

type RackService struct{}

func NewRackService() *RackService {
	return &RackService{}
}

func (r *RackService) CreateRack(rack *schemas.RackCreate) (string, error) {
	if rack.UHeight == nil {
		rack.UHeight = new(uint8)
		*rack.UHeight = 42
	}
	newRack := models.Rack{
		Name:           rack.Name,
		SerialNumber:   rack.SerialNumber,
		UHeight:        *rack.UHeight,
		SiteId:         rack.SiteId,
		DescUnit:       true, // default as true, for backward compatibility
		OrganizationId: global.OrganizationId.Get(),
	}
	err := gen.Rack.Create(&newRack)
	if err != nil {
		return "", err
	}
	return newRack.Id, nil
}

func (r *RackService) UpdateRack(rackId string, rack *schemas.RackUpdate) error {
	updateFields := make(map[string]any)
	if rack.Name != nil {
		updateFields["name"] = rack.Name
	}
	if rack.SerialNumber != nil {
		updateFields["serialNumber"] = rack.SerialNumber
	}
	if rack.UHeight != nil {
		if err := r.validateUpdateRack(rackId, *rack.UHeight); err != nil {
			return err
		}
		updateFields["uHeight"] = rack.UHeight
	}
	_, err := gen.Rack.Select(gen.Rack.Id.Eq(rackId), gen.Rack.OrganizationId.Eq(global.OrganizationId.Get())).Updates(updateFields)
	return err
}

func (r *RackService) validateUpdateRack(rackId string, uHeight uint8) error {
	rackDevices, err := gen.Device.Select(gen.Device.RackPosition).Where(
		gen.Device.RackId.Eq(rackId),
		gen.Device.OrganizationId.Eq(global.OrganizationId.Get()),
	).Find()
	if err != nil {
		return err
	}
	for _, device := range rackDevices {
		if device.RackPosition != nil {
			for _, position := range *device.RackPosition {
				if position > uHeight {
					return errors.NewError(errors.CodeUpdateRackFailed, errors.MsgUpdateRackFailed)
				}
			}
		}
	}
	return nil
}

func (r *RackService) DeleteRack(rackId string) error {
	_, err := gen.Rack.Where(gen.Rack.Id.Eq(rackId), gen.Rack.OrganizationId.Eq(global.OrganizationId.Get())).Delete()
	return err
}

func (r *RackService) GetRackByID(rackID string) (*schemas.Rack, error) {
	rack, err := gen.Rack.Select().Where(
		gen.Rack.Id.Eq(rackID),
		gen.Rack.OrganizationId.Eq(global.OrganizationId.Get()),
	).First()
	if err != nil {
		return nil, err
	}
	return &schemas.Rack{
		Id:           rack.Id,
		Name:         rack.Name,
		SerialNumber: rack.SerialNumber,
		UHeight:      rack.UHeight,
		SiteId:       rack.SiteId,
	}, nil
}

func (r *RackService) ListRacks(params *schemas.RackQuery) (int64, *schemas.RackList, error) {
	stmt := gen.Rack.Where(gen.Rack.OrganizationId.Eq(global.OrganizationId.Get()))
	if params.SiteId != nil {
		stmt = stmt.Where(gen.Rack.SiteId.Eq(*params.SiteId))
	}
	if params.Name != nil {
		stmt = stmt.Where(gen.Rack.Name.In(*params.Name...))
	}
	if params.SerialNumber != nil {
		stmt = stmt.Where(gen.Rack.SerialNumber.In(*params.SerialNumber...))
	}
	count, err := stmt.Count()
	if err != nil || count <= 0 {
		return 0, nil, err
	}
	if params.Keyword != nil {
		stmt.UnderlyingDB().Scopes(params.Search(models.RackSearchFields))
	}
	stmt.UnderlyingDB().Scopes(params.OrderByField())
	stmt.UnderlyingDB().Scopes(params.LimitOffset())
	racks, err := stmt.Find()
	if err != nil {
		nilResult := schemas.RackList{}
		return 0, &nilResult, err
	}
	var list schemas.RackList
	for _, rack := range racks {
		list = append(list, schemas.Rack{
			Id:           rack.Id,
			Name:         rack.Name,
			SerialNumber: rack.SerialNumber,
			UHeight:      rack.UHeight,
			SiteId:       rack.SiteId,
		})
	}
	return count, &list, nil
}
