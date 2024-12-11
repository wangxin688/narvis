package infra_biz

import (
	"slices"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/intend/logger"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	infra_utils "github.com/wangxin688/narvis/server/features/infra/utils"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/pkg/contextvar"
	"github.com/wangxin688/narvis/server/tools/errors"
	"go.uber.org/zap"
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
	orgId := contextvar.OrganizationId.Get()
	err := NewIsolationService().CheckSiteNotFound(rack.SiteId, orgId)
	if err != nil {
		return "", err
	}
	newRack := models.Rack{
		Name:           rack.Name,
		SerialNumber:   rack.SerialNumber,
		UHeight:        *rack.UHeight,
		SiteId:         rack.SiteId,
		DescUnit:       true, // default as true, for backward compatibility
		OrganizationId: orgId,
	}
	err = gen.Rack.Create(&newRack)
	if err != nil {
		return "", err
	}
	return newRack.Id, nil
}

func (r *RackService) UpdateRack(rackId string, rack *schemas.RackUpdate) (err error) {
	dbRack, err := gen.Rack.Where(gen.Rack.Id.Eq(rackId), gen.Rack.OrganizationId.Eq(contextvar.OrganizationId.Get())).First()
	if err != nil {
		return err
	}
	updateFields := make(map[string]*contextvar.Diff)
	if rack.Name != nil && *rack.Name != dbRack.Name {
		updateFields["name"] = &contextvar.Diff{Before: dbRack.Name, After: *rack.Name}
		dbRack.Name = *rack.Name
	}
	if rack.SerialNumber != nil && rack.SerialNumber != dbRack.SerialNumber {
		updateFields["serialNumber"] = &contextvar.Diff{Before: *dbRack.SerialNumber, After: *rack.SerialNumber}
		dbRack.SerialNumber = rack.SerialNumber
	}
	if rack.UHeight != nil && *rack.UHeight != dbRack.UHeight {
		if err := r.validateUpdateRack(rackId, *rack.UHeight); err != nil {
			return err
		}
		updateFields["uHeight"] = &contextvar.Diff{Before: dbRack.UHeight, After: *rack.UHeight}
		dbRack.UHeight = *rack.UHeight
	}
	if len(updateFields) == 0 {
		return nil
	}
	diffValue := make(map[string]map[string]*contextvar.Diff)
	diffValue[rackId] = updateFields
	contextvar.OrmDiff.Set(diffValue)
	err = gen.Rack.UnderlyingDB().Save(dbRack).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *RackService) validateUpdateRack(rackId string, uHeight uint8) error {
	rackDevices, err := gen.Device.Select(gen.Device.RackPosition).Where(
		gen.Device.RackId.Eq(rackId),
		gen.Device.OrganizationId.Eq(contextvar.OrganizationId.Get()),
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
	_, err := gen.Rack.Where(gen.Rack.Id.Eq(rackId), gen.Rack.OrganizationId.Eq(contextvar.OrganizationId.Get())).Delete()
	return err
}

func (r *RackService) GetRackById(rackId string) (*schemas.Rack, error) {
	rack, err := gen.Rack.Select().Where(
		gen.Rack.Id.Eq(rackId),
		gen.Rack.OrganizationId.Eq(contextvar.OrganizationId.Get()),
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
	stmt := gen.Rack.Where(gen.Rack.OrganizationId.Eq(contextvar.OrganizationId.Get()))
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

func (r *RackService) ValidateCreateRackReservation(rackId string, uHeight uint8, positions []uint8) bool {
	devices, err := gen.Device.Select(gen.Device.RackPosition).Where(gen.Device.RackId.Eq(rackId), gen.Device.OrganizationId.Eq(contextvar.OrganizationId.Get())).Find()
	if err != nil {
		return false
	}
	servers, err := gen.Server.Select(gen.Server.RackPosition).Where(gen.Server.RackId.Eq(rackId), gen.Server.OrganizationId.Eq(contextvar.OrganizationId.Get())).Find()
	if err != nil {
		return false
	}
	if len(devices)+len(servers) == 0 && len(positions) <= int(uHeight) {
		return true
	}
	usedPositions := make([]uint8, 0)
	for _, device := range devices {
		if device.RackPosition != nil {
			ps, _ := infra_utils.ParseUint8s(*device.RackPosition)
			usedPositions = append(usedPositions, ps...)
		}
	}
	for _, server := range servers {
		if server.RackPosition != nil {
			ps, _ := infra_utils.ParseUint8s(*server.RackPosition)
			usedPositions = append(usedPositions, ps...)
		}
	}
	// check if all positions are available
	for _, position := range positions {
		if lo.Contains(usedPositions, position) {
			return false
		}
	}
	return true
}

func (r *RackService) ValidateDeviceUpdateRackReservation(rackId string, uHeight uint8, positions []uint8) bool {
	return true
}

// get rack used positions and result sorted asc slice
func (r *RackService) GetRackUsedPositions(rackId string) ([]uint8, error) {
	result := make([]uint8, 0)
	devices, err := gen.Device.Select(gen.Device.Id, gen.Device.RackPosition).Where(gen.Device.RackId.Eq(rackId), gen.Device.OrganizationId.Eq(contextvar.OrganizationId.Get())).Find()
	if err != nil {
		return nil, err
	}
	for _, device := range devices {
		if device.RackPosition != nil {
			ps, _ := infra_utils.ParseUint8s(*device.RackPosition)
			result = append(result, ps...)
		}
	}
	slices.Sort(result)
	return result, nil
}

func (r *RackService) GetRackDevices(rackIds []string) (map[string]*models.Device, error) {
	devices, err := gen.Device.Select(gen.Device.Id, gen.Device.RackPosition).Where(gen.Device.RackId.In(rackIds...), gen.Device.OrganizationId.Eq(contextvar.OrganizationId.Get())).Find()
	if err != nil {
		return nil, err
	}
	result := make(map[string]*models.Device, 0)
	for _, device := range devices {
		result[device.Id] = device
	}
	return result, nil
}

func (r *RackService) GetRackElevation(rackId string) (*schemas.RackElevation, error) {
	rack, err := r.GetRackById(rackId)
	if err != nil {
		return nil, err
	}
	rackElevation, err := r.getRackElevation(rackId)
	if err != nil {
		return nil, err
	}
	usedPositions := make([]uint8, 0)
	for _, item := range rackElevation {
		usedPositions = append(usedPositions, item.Position...)
	}
	totalPositions := make([]uint8, 0)
	for i := uint8(1); i <= rack.UHeight; i++ {
		totalPositions = append(totalPositions, i)
	}
	availablePositions := lo.Without(totalPositions, usedPositions...)
	slices.Sort(availablePositions)
	return &schemas.RackElevation{
		Id:                 rack.Id,
		Name:               rack.Name,
		SerialNumber:       rack.SerialNumber,
		UHeight:            rack.UHeight,
		SiteId:             rack.SiteId,
		Items:              rackElevation,
		AvailablePositions: availablePositions,
	}, nil
}

func (r *RackService) getRackElevation(rackId string) ([]*schemas.RackElevationItem, error) {
	orgId := contextvar.OrganizationId.Get()
	devices, err := gen.Device.Select(
		gen.Device.Id,
		gen.Device.Name,
		gen.AP.DeviceRole,
		gen.Device.RackPosition,
		gen.Device.ManagementIp,
	).Where(gen.Device.RackId.Eq(rackId), gen.Device.OrganizationId.Eq(orgId)).Find()
	if err != nil {
		return nil, err
	}
	if len(devices) == 0 {
		return nil, nil
	}
	deviceIds := lo.Map(devices, func(device *models.Device, _ int) string {
		return device.Id
	})

	opStatus, err := GetDeviceOpStatus(deviceIds, orgId)
	if err != nil {
		logger.Logger.Error("[infraDeviceService]: failed to get device op status", zap.Error(err))
		return nil, err
	}
	result := make([]*schemas.RackElevationItem, 0)
	for _, device := range devices {
		if device.RackPosition != nil {
			ps, _ := infra_utils.ParseUint8s(*device.RackPosition)
			result = append(result, &schemas.RackElevationItem{
				Id:              device.Id,
				Name:            device.Name,
				DeviceRole:      device.DeviceRole,
				ManagementIp:    device.ManagementIp,
				Position:        ps,
				OperatingStatus: opStatus[device.Id],
			})
		}
	}
	return result, nil
}
