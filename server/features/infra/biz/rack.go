package infra_biz

import (
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	infra_utils "github.com/wangxin688/narvis/server/features/infra/utils"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/tools/errors"
	ts "github.com/wangxin688/narvis/server/tools/schemas"
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

func (r *RackService) UpdateRack(rackId string, rack *schemas.RackUpdate) (err error) {
	dbRack, err := gen.Rack.Where(gen.Rack.Id.Eq(rackId), gen.Rack.OrganizationId.Eq(global.OrganizationId.Get())).First()
	if err != nil {
		return err
	}
	updateFields := make(map[string]*ts.OrmDiff)
	if rack.Name != nil && *rack.Name != dbRack.Name {
		updateFields["name"] = &ts.OrmDiff{Before: dbRack.Name, After: *rack.Name}
		dbRack.Name = *rack.Name
	}
	if rack.SerialNumber != nil && *rack.SerialNumber != *dbRack.SerialNumber {
		updateFields["serialNumber"] = &ts.OrmDiff{Before: *dbRack.SerialNumber, After: *rack.SerialNumber}
		dbRack.SerialNumber = rack.SerialNumber
	}
	if rack.UHeight != nil && *rack.UHeight != dbRack.UHeight {
		if err := r.validateUpdateRack(rackId, *rack.UHeight); err != nil {
			return err
		}
		updateFields["uHeight"] = &ts.OrmDiff{Before: dbRack.UHeight, After: *rack.UHeight}
		dbRack.UHeight = *rack.UHeight
	}
	if len(updateFields) == 0 {
		return nil
	}
	diffValue := make(map[string]map[string]*ts.OrmDiff)
	diffValue[rackId] = updateFields
	global.OrmDiff.Set(diffValue)
	err = gen.Rack.UnderlyingDB().Save(dbRack).Error
	if err != nil {
		return err
	}
	return nil
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
			devicePositions, _ := infra_utils.ParseUint8s(*device.RackPosition) // comma separated string to uint8 slice, and trust db data is correct
			for _, position := range devicePositions {
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

func (r *RackService) ListRacks(params *schemas.RackQuery) (int64, *[]*schemas.Rack, error) {
	res := make([]*schemas.Rack, 0)
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
	if err != nil || count < 0 {
		return 0, &res, err
	}
	if params.IsSearchable() {
		keyword := "%" + *params.Keyword + "%"
		stmt = stmt.Where(
			gen.Rack.Name.Like(keyword),
		).Or(
			gen.Rack.SerialNumber.Like(keyword),
		)
	}
	stmt.UnderlyingDB().Scopes(params.OrderByField())
	stmt.UnderlyingDB().Scopes(params.Pagination())
	racks, err := stmt.Find()
	if err != nil {
		return 0, &res, err
	}
	for _, rack := range racks {
		res = append(res, &schemas.Rack{
			Id:           rack.Id,
			Name:         rack.Name,
			SerialNumber: rack.SerialNumber,
			UHeight:      rack.UHeight,
			SiteId:       rack.SiteId,
		})
	}
	return count, &res, nil
}

func (r *RackService) ValidateDeviceCreateRackReservation(rackId string, uHeight uint8, positions []uint8) bool {
	devices, err := gen.Device.Select(gen.Device.RackPosition).Where(gen.Device.RackId.Eq(rackId), gen.Device.OrganizationId.Eq(global.OrganizationId.Get())).Find()
	if err != nil {
		return false
	}
	if len(devices) == 0 && len(positions) <= int(uHeight) {
		return true
	}
	usedPositions := make([]uint8, 0)
	for _, device := range devices {
		if device.RackPosition != nil {
			ps, _ := infra_utils.ParseUint8s(*device.RackPosition)
			usedPositions = append(usedPositions, ps...)
		}
	}
	// check if all positions are available

	return true
}
