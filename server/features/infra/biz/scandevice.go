package infra_biz

import (
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/infra/hooks"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/tools"
	"github.com/wangxin688/narvis/server/tools/errors"
	"gorm.io/gorm"
)

type ScanDeviceService struct{}

func NewScanDeviceService() *ScanDeviceService {
	return &ScanDeviceService{}
}

func (s *ScanDeviceService) ListScanDevice(query *schemas.ScanDeviceQuery) (int64, []*schemas.ScanDevice, error) {
	results := make([]*schemas.ScanDevice, 0)
	stmt := gen.ScanDevice.Where(
		gen.ScanDevice.OrganizationId.Eq(global.OrganizationId.Get()),
	)
	if query.ManagementIp != nil {
		stmt = stmt.Where(gen.ScanDevice.ManagementIp.In(*query.ManagementIp...))
	}
	if query.Platform != nil {
		stmt = stmt.Where(gen.ScanDevice.Platform.In(*query.Platform...))
	}
	if query.Manufacturer != nil {
		stmt = stmt.Where(gen.ScanDevice.Manufacturer.In(*query.Manufacturer...))
	}
	if query.DeviceModel != nil {
		stmt = stmt.Where(gen.ScanDevice.DeviceModel.In(*query.DeviceModel...))
	}
	if query.ChassisId != nil {
		stmt = stmt.Where(gen.ScanDevice.ChassisId.In(*query.ChassisId...))
	}
	if query.IsSearchable() {
		keyword := "%" + *query.Keyword + "%"
		stmt = stmt.Where(
			gen.ScanDevice.Name.Like(keyword),
		).Or(
			gen.ScanDevice.ChassisId.Like(keyword),
		).Or(
			gen.ScanDevice.ManagementIp.Like(keyword),
		)
	}

	total, err := stmt.Count()
	if err != nil {
		return 0, results, err
	}
	stmt.UnderlyingDB().Scopes(query.OrderByField())
	stmt.UnderlyingDB().Scopes(query.Pagination())
	list, err := stmt.Find()
	if err != nil {
		return 0, results, err
	}

	for _, item := range list {
		results = append(results, &schemas.ScanDevice{
			Id:           item.Id,
			CreatedAt:    item.CreatedAt,
			Name:         item.Name,
			ManagementIp: item.ManagementIp,
			Platform:     item.Platform,
			DeviceModel:  item.DeviceModel,
			Manufacturer: item.Manufacturer,
			ChassisId:    item.ChassisId,
			Description:  item.Description,
		})
	}
	return total, results, nil
}

func (s *ScanDeviceService) GetById(id string) (*schemas.ScanDevice, error) {
	device, err := gen.ScanDevice.Where(
		gen.ScanDevice.Id.Eq(id),
		gen.ScanDevice.OrganizationId.Eq(global.OrganizationId.Get()),
	).First()
	if err != nil {
		return nil, err
	}

	return &schemas.ScanDevice{
		Id:           device.Id,
		CreatedAt:    device.CreatedAt,
		Name:         device.Name,
		ManagementIp: device.ManagementIp,
		Platform:     device.Platform,
		DeviceModel:  device.DeviceModel,
		Manufacturer: device.Manufacturer,
		ChassisId:    device.ChassisId,
		Description:  device.Description,
	}, nil
}

func (s *ScanDeviceService) UpdateById(id string, update *schemas.ScanDeviceUpdate) (string, error) {
	orgId := global.OrganizationId.Get()
	ok, err := LicenseUsageDepends(1, orgId)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", errors.NewError(errors.CodeLicenseCountExceeded, errors.MsgLicenseCountExceeded)
	}
	device, err := gen.ScanDevice.Where(
		gen.ScanDevice.Id.Eq(id),
		gen.ScanDevice.OrganizationId.Eq(orgId),
	).First()
	if err != nil {
		return "", err
	}
	newDevice := schemas.DeviceCreate{
		Name:         device.Name,
		ManagementIp: device.ManagementIp,
		Status:       update.Status,
		Platform:     &device.Platform,
		DeviceModel:  &device.DeviceModel,
		Manufacturer: &device.Manufacturer,
		ChassisId:    &device.ChassisId,
		Description:  &device.Description,
		Floor:        update.Floor,
		RackId:       update.RackId,
		RackPosition: update.RackPosition,
		SiteId:       update.SiteId,
		DeviceRole:   update.DeviceRole,
	}
	newDeviceId := ""
	err = gen.ScanDevice.UnderlyingDB().Transaction(func(tx *gorm.DB) error {
		newDeviceId, err = NewDeviceService().CreateDevice(&newDevice)
		if err != nil {
			return err
		}
		_, err = gen.ScanDevice.Delete(device)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	tools.BackgroundTask(func() {
		hooks.DeviceCreateHooks(newDeviceId)
	})
	return newDeviceId, nil
}

func (s *ScanDeviceService) BatchUpdate(device *schemas.ScanDeviceBatchUpdate) ([]string, error) {
	orgId := global.OrganizationId.Get()
	ok, err := LicenseUsageDepends(uint32(len(device.Ids)), orgId)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, errors.NewError(errors.CodeLicenseCountExceeded, errors.MsgLicenseCountExceeded)
	}
	for _, id := range device.Ids {
		_, err := s.UpdateById(id, &schemas.ScanDeviceUpdate{
			Status:     device.Status,
			Floor:      device.Floor,
			DeviceRole: device.DeviceRole,
			SiteId:     device.SiteId,
		})
		if err != nil {
			return nil, err
		}
	}
	return device.Ids, nil
}

func (s *ScanDeviceService) GetByScanResult(ips []string, orgId string) (map[string]*models.ScanDevice, error) {
	devices, err := gen.ScanDevice.Select(gen.ScanDevice.Id, gen.ScanDevice.ManagementIp).Where(
		gen.ScanDevice.ManagementIp.In(ips...),
		gen.ScanDevice.OrganizationId.Eq(orgId),
	).Find()
	if err != nil {
		return nil, err
	}
	result := map[string]*models.ScanDevice{}
	for _, device := range devices {
		result[device.ManagementIp] = device
	}
	return result, nil
}

func (s *ScanDeviceService) DeleteById(id string) error {
	_, err := gen.ScanDevice.Where(
		gen.ScanDevice.Id.Eq(id),
		gen.ScanDevice.OrganizationId.Eq(global.OrganizationId.Get()),
	).Delete()
	return err
}

func (s *ScanDeviceService) DeleteByIds(ids []string) error {
	_, err := gen.ScanDevice.Where(
		gen.ScanDevice.Id.In(ids...),
		gen.ScanDevice.OrganizationId.Eq(global.OrganizationId.Get()),
	).Delete()
	return err
}
