package infra_biz

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	infra_utils "github.com/wangxin688/narvis/server/features/infra/utils"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/tools/helpers"
	ts "github.com/wangxin688/narvis/server/tools/schemas"
)

type DeviceService struct{}

func NewDeviceService() *DeviceService {
	return &DeviceService{}
}

func (d *DeviceService) CreateDevice(device *schemas.DeviceCreate) (string, error) {
	newDevice := models.Device{
		Name:           device.Name,
		ManagementIp:   device.ManagementIp,
		Status:         device.Status,
		DeviceModel:    *device.DeviceModel,
		Manufacturer:   *device.Manufacturer,
		DeviceRole:     device.DeviceRole,
		Floor:          device.Floor,
		SiteId:         device.SiteId,
		OrganizationId: global.OrganizationId.Get(),
	}
	if device.RackId != nil && device.RackPosition != nil {
		newDevice.RackId = device.RackId
		position, err := infra_utils.SliceUint8ToString(device.RackPosition)
		if err != nil {
			return "", err
		}
		newDevice.RackPosition = &position
	}

	err := gen.Device.Create(&newDevice)
	if err != nil {
		return "", err
	}
	return newDevice.Id, nil
}

func (d *DeviceService) UpdateDevice(g *gin.Context, deviceId string, device *schemas.DeviceUpdate) (diff map[string]map[string]*ts.OrmDiff, err error) {
	dbDevice, err := gen.Device.Where(gen.Device.Id.Eq(deviceId), gen.Device.OrganizationId.Eq(global.OrganizationId.Get())).First()
	if err != nil {
		return nil, err
	}
	updateFields := make(map[string]*ts.OrmDiff)
	if device.Name != nil && *device.Name != dbDevice.Name {
		updateFields["name"] = &ts.OrmDiff{Before: dbDevice.Name, After: *device.Name}
		dbDevice.Name = *device.Name
	}
	if device.Status != nil && *device.Status != dbDevice.Status {
		updateFields["status"] = &ts.OrmDiff{Before: dbDevice.Status, After: *device.Status}
		dbDevice.Status = *device.Status
	}
	if device.ManagementIp != nil && *device.ManagementIp != dbDevice.ManagementIp {
		updateFields["management_ip"] = &ts.OrmDiff{Before: dbDevice.ManagementIp, After: *device.ManagementIp}
		dbDevice.ManagementIp = *device.ManagementIp
	}
	if device.DeviceModel != nil && *device.DeviceModel != dbDevice.DeviceModel {
		updateFields["device_model"] = &ts.OrmDiff{Before: dbDevice.DeviceModel, After: *device.DeviceModel}
		dbDevice.DeviceModel = *device.DeviceModel
	}
	if device.Manufacturer != nil && *device.Manufacturer != dbDevice.Manufacturer {
		updateFields["manufacturer"] = &ts.OrmDiff{Before: dbDevice.Manufacturer, After: *device.Manufacturer}
		dbDevice.Manufacturer = *device.Manufacturer
	}
	if device.DeviceRole != nil && *device.DeviceRole != dbDevice.DeviceRole {
		updateFields["device_role"] = &ts.OrmDiff{Before: dbDevice.DeviceRole, After: *device.DeviceRole}
		dbDevice.DeviceRole = *device.DeviceRole
	}
	if device.Floor != nil && *device.Floor != *dbDevice.Floor {
		updateFields["floor"] = &ts.OrmDiff{Before: dbDevice.Floor, After: *device.Floor}
		dbDevice.Floor = device.Floor
	}
	if device.Description != nil && *device.Description != *dbDevice.Description {
		updateFields["description"] = &ts.OrmDiff{Before: dbDevice.Description, After: *device.Description}
		dbDevice.Description = device.Description
	}
	if device.OsVersion != nil && *device.OsVersion != *dbDevice.OsVersion {
		updateFields["os_version"] = &ts.OrmDiff{Before: dbDevice.OsVersion, After: *device.OsVersion}
		dbDevice.OsVersion = device.OsVersion
	}
	if helpers.HasParams(g, "rackId") && device.RackId != dbDevice.RackId {
		updateFields["rack_id"] = &ts.OrmDiff{Before: dbDevice.RackId, After: *device.RackId}
		dbDevice.RackId = device.RackId
	}
	if helpers.HasParams(g, "rackPosition") {
		position, err := infra_utils.SliceUint8ToString(*device.RackPosition)
		if err != nil {
			return nil, err
		}
		if position != *dbDevice.RackPosition {
			updateFields["rack_position"] = &ts.OrmDiff{Before: dbDevice.RackPosition, After: position}
			dbDevice.RackPosition = &position
		}
	}
	if len(updateFields) == 0 {
		return nil, nil
	}
	diffValue := make(map[string]map[string]*ts.OrmDiff)
	diffValue[deviceId] = updateFields
	global.OrmDiff.Set(diffValue)
	err = gen.Device.UnderlyingDB().Save(dbDevice).Error
	if err != nil {
		return nil, err
	}
	return diffValue, nil
}

func (d *DeviceService) DeleteDevice(deviceId string) (*models.Device, error) {
	dbDevice, err := gen.Device.Where(gen.Device.Id.Eq(deviceId), gen.Device.OrganizationId.Eq(global.OrganizationId.Get())).First()
	if err != nil {
		return nil, err
	}
	_, err = gen.Device.Delete(dbDevice)
	if err != nil {
		return nil, err
	}
	return dbDevice, nil
}

func (d *DeviceService) GetById(deviceId string) (*schemas.Device, error) {
	device, err := gen.Device.Select(gen.Device.Id.Eq(deviceId), gen.Device.OrganizationId.Eq(global.OrganizationId.Get())).First()
	if err != nil {
		return nil, err
	}
	return &schemas.Device{
		Id:            device.Id,
		CreatedAt:     device.CreatedAt,
		UpdatedAt:     device.UpdatedAt,
		Name:          device.Name,
		ManagementIp:  device.ManagementIp,
		Platform:      device.Platform,
		ProductFamily: device.ProductFamily,
		Status:        device.Status,
		OperStatus:    "",
		DeviceModel:   device.DeviceModel,
		Manufacturer:  device.Manufacturer,
		DeviceRole:    device.DeviceRole,
		Floor:         device.Floor,
		OsPatch:       device.OsPatch,
		OsVersion:     device.OsVersion,
		Description:   device.Description,
		RackId:        device.RackId,
		RackPosition: func() *[]uint8 {
			if device.RackPosition == nil {
				return nil
			} else {
				position, _ := infra_utils.ParseUint8s(*device.RackPosition)
				return &position
			}
		}(),
		SiteId: device.SiteId,
	}, nil
}

func (d *DeviceService) GetDeviceList(query *schemas.DeviceQuery) (int64, *[]*schemas.Device, error) {
	res := make([]*schemas.Device, 0)
	stmt := gen.Device.Where(gen.Device.OrganizationId.Eq(global.OrganizationId.Get()))
	if query.Name != nil {
		stmt = stmt.Where(gen.Device.Name.In(*query.Name...))
	}

	if query.Status != nil {
		stmt = stmt.Where(gen.Device.Status.Eq(*query.Status))
	}

	if query.SiteId != nil {
		stmt = stmt.Where(gen.Device.SiteId.Eq(*query.SiteId))
	}

	if query.DeviceRole != nil {
		stmt = stmt.Where(gen.Device.DeviceRole.In(*query.DeviceRole...))
	}

	if query.DeviceModel != nil {
		stmt = stmt.Where(gen.Device.DeviceModel.In(*query.DeviceModel...))
	}

	if query.Manufacturer != nil {
		stmt = stmt.Where(gen.Device.Manufacturer.In(*query.Manufacturer...))
	}

	if query.RackId != nil {
		stmt = stmt.Where(gen.Device.RackId.Eq(*query.RackId))
	}

	if query.Floor != nil {
		stmt = stmt.Where(gen.Device.Floor.Eq(*query.Floor))
	}
	if query.SerialNumber != nil {
		stmt = stmt.Where(gen.Device.SerialNumber.Eq(*query.SerialNumber))
	}
	if query.IsSearchable() {
		keyword := "%" + *query.Keyword + "%"
		stmt = stmt.Where(
			gen.Device.Name.Like(keyword),
		).Or(
			gen.Device.ChassisId.Like(keyword),
		).Or(
			gen.Device.SerialNumber.Like(keyword),
		).Or(
			gen.Device.ManagementIp.Like(keyword),
		)
	}

	count, err := stmt.Count()
	if err != nil && count < 0 {
		return 0, &res, err
	}
	stmt.UnderlyingDB().Scopes(query.OrderByField())
	stmt.UnderlyingDB().Scopes(query.Pagination())
	list, err := stmt.Find()
	if err != nil {
		return 0, &res, err
	}

	for _, item := range list {
		res = append(res, &schemas.Device{
			Id:            item.Id,
			CreatedAt:     item.CreatedAt,
			UpdatedAt:     item.UpdatedAt,
			Name:          item.Name,
			ManagementIp:  item.ManagementIp,
			Platform:      item.Platform,
			ProductFamily: item.ProductFamily,
			Status:        item.Status,
			OperStatus:    "",
			IsRegistered:  item.IsRegistered,
			DeviceModel:   item.DeviceModel,
			Manufacturer:  item.Manufacturer,
			DeviceRole:    item.DeviceRole,
			Floor:         item.Floor,
			OsPatch:       item.OsPatch,
			OsVersion:     item.OsVersion,
			Description:   item.Description,
			RackId:        item.RackId,
			RackPosition: func() *[]uint8 {
				if item.RackPosition == nil {
					return nil
				} else {
					position, _ := infra_utils.ParseUint8s(*item.RackPosition)
					return &position
				}
			}(),
			SiteId: item.SiteId,
		})
	}
	return count, &res, nil
}

func (d *DeviceService) GetDeviceInterfaces(deviceId string) (*[]*schemas.DeviceInterface, error) {

	list, err := gen.DeviceInterface.Where(gen.DeviceInterface.DeviceId.Eq(deviceId)).Find()

	if err != nil {
		return nil, err
	}
	res := make([]*schemas.DeviceInterface, 0)
	for _, item := range list {
		res = append(res, &schemas.DeviceInterface{
			Id:            item.Id,
			CreatedAt:     item.CreatedAt,
			UpdatedAt:     item.UpdatedAt,
			IfIndex:       item.IfIndex,
			IfName:        item.IfName,
			IfDescr:       item.IfDescr,
			IfType:        item.IfType,
			IfMtu:         item.IfMtu,
			IfSpeed:       item.IfSpeed,
			IfPhysAddr:    *item.IfPhysAddr,
			IfAdminStatus: item.IfAdminStatus,
			IfOperStatus:  item.IfOperStatus,
			IfHighSpeed:   item.IfHighSpeed,
			IfLastChange:  item.IfLastChange,
		})
	}
	return &res, nil
}
