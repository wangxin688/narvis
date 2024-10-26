package infra_biz

import (
	"fmt"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/tools/helpers"
	"go.uber.org/zap"
)

type ApService struct{}

func NewApService() *ApService {
	return &ApService{}
}

func (s *ApService) GetApList(query *schemas.ApQuery) (int64, *[]*schemas.AP, error) {
	res := make([]*schemas.AP, 0)
	orgId := global.OrganizationId.Get()
	stmt := gen.AP.Where(gen.AP.OrganizationId.Eq(orgId))
	if query.Name != nil {
		stmt = stmt.Where(gen.AP.Name.In(*query.Name...))
	}
	if query.ManagementIp != nil {
		stmt = stmt.Where(gen.AP.ManagementIp.In(*query.ManagementIp...))
	}
	if query.SiteId != nil {
		stmt = stmt.Where(gen.AP.SiteId.Eq(*query.SiteId))
	}
	if query.DeviceModel != nil {
		stmt = stmt.Where(gen.AP.DeviceModel.In(*query.DeviceModel...))
	}
	if query.Floor != nil {
		stmt = stmt.Where(gen.AP.Floor.Eq(*query.Floor))
	}
	if query.SerialNumber != nil {
		stmt = stmt.Where(gen.AP.SerialNumber.Eq(*query.SerialNumber))
	}
	if query.IsSearchable() {
		searchString := "%" + *query.Keyword + "%"
		stmt = stmt.Where(gen.AP.Name.Like(searchString)).Or(gen.AP.ManagementIp.Like(searchString))
	}

	count, err := stmt.Count()
	if err != nil || count <= 0 {
		return 0, &res, err
	}
	stmt.UnderlyingDB().Scopes(query.OrderByField())
	stmt.UnderlyingDB().Scopes(query.Pagination())
	aps, err := stmt.Find()
	if err != nil {
		return 0, &res, err
	}
	apIds := lo.Map(aps, func(ap *models.AP, _ int) string { return ap.Id })
	opStatus, err := GetApOpStatus(apIds, orgId)
	if err != nil {
		core.Logger.Error("[metricService]: failed to get ap operation status", zap.Error(err))
	}

	for _, ap := range aps {
		res = append(res, &schemas.AP{
			Id:        ap.Id,
			CreatedAt: ap.CreatedAt,
			UpdatedAt: ap.UpdatedAt,
			Name:      ap.Name,
			Status:    ap.Status,
			OperStatus: func() string {
				if status, ok := opStatus[ap.Id]; ok {
					return status
				}
				return "nodata"
			}(),
			MacAddress:      ap.MacAddress,
			SerialNumber:    ap.SerialNumber,
			ManagementIp:    ap.ManagementIp,
			DeviceRole:      ap.DeviceRole,
			Manufacturer:    ap.Manufacturer,
			DeviceModel:     ap.DeviceModel,
			OsVersion:       ap.OsVersion,
			Floor:           ap.Floor,
			GroupName:       ap.GroupName,
			CoordinateX:     ap.CoordinateX,
			CoordinateY:     ap.CoordinateY,
			CoordinateZ:     ap.CoordinateZ,
			WlanACIpAddress: ap.WlanACIpAddress,
			SiteId:          ap.SiteId,
		})
	}
	return count, &res, nil

}

func (s *ApService) GetById(id string) (*schemas.AP, error) {
	orgId := global.OrganizationId.Get()
	ap, err := gen.AP.Where(gen.AP.OrganizationId.Eq(orgId), gen.AP.Id.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	opStatus, err := GetApOpStatus([]string{ap.Id}, orgId)
	if err != nil {
		core.Logger.Error("[metricService]: failed to get ap operation status", zap.Error(err))
	}
	return &schemas.AP{
		Id:        ap.Id,
		CreatedAt: ap.CreatedAt,
		UpdatedAt: ap.UpdatedAt,
		Name:      ap.Name,
		Status:    ap.Status,
		OperStatus: func() string {
			if status, ok := opStatus[ap.Id]; ok {
				return status
			}
			return "nodata"
		}(),
		MacAddress:      ap.MacAddress,
		SerialNumber:    ap.SerialNumber,
		ManagementIp:    ap.ManagementIp,
		DeviceRole:      ap.DeviceRole,
		Manufacturer:    ap.Manufacturer,
		DeviceModel:     ap.DeviceModel,
		OsVersion:       ap.OsVersion,
		Floor:           ap.Floor,
		GroupName:       ap.GroupName,
		CoordinateX:     ap.CoordinateX,
		CoordinateY:     ap.CoordinateY,
		CoordinateZ:     ap.CoordinateZ,
		WlanACIpAddress: ap.WlanACIpAddress,
		SiteId:          ap.SiteId,
	}, nil
}

func (s *ApService) UpdateApById(id string, ap *schemas.ApUpdate) error {

	dbAp, err := gen.AP.Where(gen.AP.OrganizationId.Eq(global.OrganizationId.Get()), gen.AP.Id.Eq(id)).First()
	if err != nil {
		return err
	}
	if ap.Status != nil {
		dbAp.Status = *ap.Status
	}

	if ap.Floor != nil {
		dbAp.Floor = ap.Floor
	}
	if ap.CoordinateX != nil {
		dbAp.CoordinateX = ap.CoordinateX
	}
	if ap.CoordinateY != nil {
		dbAp.CoordinateY = ap.CoordinateY
	}
	if ap.CoordinateZ != nil {
		dbAp.CoordinateZ = ap.CoordinateZ
	}
	err = gen.AP.Save(dbAp)
	return err
}

func (s *ApService) BatchUpdateAp(ap *schemas.ApBatchUpdate) ([]string, error) {
	dbAps, err := gen.AP.Where(gen.AP.OrganizationId.Eq(global.OrganizationId.Get()), gen.AP.Id.In(ap.Ids...)).Find()
	if err != nil {
		return nil, err
	}
	for _, dbAp := range dbAps {
		if ap.Status != nil {
			dbAp.Status = *ap.Status
		}
		if ap.Floor != nil {
			dbAp.Floor = ap.Floor
		}
		if ap.CoordinateX != nil {
			dbAp.CoordinateX = ap.CoordinateX
		}
		if ap.CoordinateY != nil {
			dbAp.CoordinateY = ap.CoordinateY
		}
		if ap.CoordinateZ != nil {
			dbAp.CoordinateZ = ap.CoordinateZ
		}
		err = gen.AP.Save(dbAp)
		if err != nil {
			return nil, err
		}
	}
	return ap.Ids, nil
}

func (s *ApService) DeleteApByIds(ids []string) error {
	_, err := gen.AP.Where(gen.AP.OrganizationId.Eq(global.OrganizationId.Get()), gen.AP.Id.In(ids...)).Delete()
	return err
}

func (s *ApService) DeleteApById(id string) error {
	_, err := gen.AP.Where(gen.AP.OrganizationId.Eq(global.OrganizationId.Get()), gen.AP.Id.Eq(id)).Delete()
	return err
}

func (s *ApService) GetApShortMap(apIds []string) (map[string]*schemas.APShort, error) {
	aps, err := gen.AP.Select(
		gen.AP.Id,
		gen.AP.Name,
		gen.AP.ManagementIp,
	).Where(gen.AP.Id.In(apIds...)).Find()

	if err != nil {
		return nil, err
	}
	res := make(map[string]*schemas.APShort)
	for _, ap := range aps {
		res[ap.Id] = &schemas.APShort{
			Id:           ap.Id,
			Name:         ap.Name,
			ManagementIP: ap.ManagementIp,
		}
	}
	return res, nil
}

func (s *ApService) SearchApByKeyword(keyword string, orgId string) ([]string, error) {
	result := make([]string, 0)
	stmt := gen.AP.Select(gen.AP.Id).Where(gen.AP.OrganizationId.Eq(orgId))
	keyword = "%" + keyword + "%"
	stmt = stmt.Where(gen.AP.Name.Like(keyword)).Or(gen.AP.ManagementIp.Like(keyword))
	err := stmt.Scan(&result)
	return result, err
}

func (s *ApService) GetByIpsAndSiteId(ips []string, siteId string, orgId string) (map[string]*models.AP, error) {

	aps, err := gen.AP.Where(
		gen.AP.OrganizationId.Eq(orgId),
		gen.AP.ManagementIp.In(ips...),
		gen.AP.SiteId.Eq(siteId),
	).Find()
	if err != nil {
		return nil, err
	}

	res := make(map[string]*models.AP)
	for _, ap := range aps {
		res[ap.ManagementIp] = ap
	}
	return res, nil
}

// GetApByMacAddresses: get ap by mac address, return map[ap.macAddress]apModel
func (s *ApService) GetApByMacAddresses(macAddresses []string, orgId string) (map[string]*models.AP, error) {
	aps, err := gen.AP.Where(
		gen.AP.OrganizationId.Eq(orgId),
		gen.AP.MacAddress.In(macAddresses...),
	).Find()
	if err != nil {
		return nil, err
	}
	res := make(map[string]*models.AP)
	for _, ap := range aps {
		res[*ap.MacAddress] = ap
	}
	return res, nil
}

func (s *ApService) CalApHash(ap *models.AP) string {
	hashString := fmt.Sprintf(
		"%s-%s-%s-%s-%s-%s", ap.Name, ap.ManagementIp,
		helpers.PtrStringToString(ap.MacAddress),
		helpers.PtrStringToString(ap.GroupName),
		helpers.PtrStringToString(ap.WlanACIpAddress),
		helpers.PtrStringToString(ap.SerialNumber),
	)
	return helpers.StringToMd5(hashString)
}

func (s *ApService) GetApIdsByNames(apNames []string, siteId string) ([]string, error) {
	result := make([]string, 0)
	err := gen.AP.Select(gen.AP.Id).Where(
		gen.AP.Name.In(apNames...),
		gen.AP.SiteId.Eq(siteId),
	).Scan(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (s *ApService) GetAllApIdsBySiteId(siteId string) ([]string, error) {
	result := make([]string, 0)
	err := gen.AP.Select(gen.AP.Id).Where(
		gen.AP.SiteId.Eq(siteId),
	).Scan(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
