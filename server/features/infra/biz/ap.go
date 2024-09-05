package infra_biz

import (
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	"github.com/wangxin688/narvis/server/global"
)

type ApService struct{}

func NewApService() *ApService {
	return &ApService{}
}

func (s *ApService) GetApList(query *schemas.ApQuery) (int64, *[]*schemas.AP, error) {
	stmt := gen.AP.Where(gen.AP.OrganizationId.Eq(global.OrganizationId.Get()))
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
	if err != nil || count < 0 {
		return 0, nil, err
	}
	stmt.UnderlyingDB().Scopes(query.OrderByField())
	stmt.UnderlyingDB().Scopes(query.LimitOffset())
	aps, err := stmt.Find()
	if err != nil {
		return 0, nil, err
	}
	var res []*schemas.AP
	for _, ap := range aps {
		res = append(res, &schemas.AP{
			Id:           ap.Id,
			CreatedAt:    ap.CreatedAt,
			UpdatedAt:    ap.UpdatedAt,
			Name:         ap.Name,
			Status:       ap.Status,
			OperStatus:   "",
			MacAddress:   ap.MacAddress,
			SerialNumber: ap.SerialNumber,
			ManagementIp: ap.ManagementIp,
			DeviceRole:   ap.DeviceRole,
			Manufacturer: ap.Manufacturer,
			DeviceModel:  ap.DeviceModel,
			OsVersion:    ap.OsVersion,
			Floor:        ap.Floor,
			GroupName:    ap.GroupName,
			Coordinate: &schemas.ApCoordinate{
				X: ap.Coordinate.Data().X,
				Y: ap.Coordinate.Data().Y,
				Z: ap.Coordinate.Data().Z,
			},
			ActiveWacId: ap.ActiveWacId,
			SiteId:      ap.SiteId,
		})
	}
	return count, &res, nil

}

func (s *ApService) GetById(id string) (*schemas.AP, error) {

	ap, err := gen.AP.Where(gen.AP.OrganizationId.Eq(global.OrganizationId.Get()), gen.AP.Id.Eq(id)).First()
	if err != nil {
		return nil, err
	}

	return &schemas.AP{
		Id:           ap.Id,
		CreatedAt:    ap.CreatedAt,
		UpdatedAt:    ap.UpdatedAt,
		Name:         ap.Name,
		Status:       ap.Status,
		OperStatus:   "",
		MacAddress:   ap.MacAddress,
		SerialNumber: ap.SerialNumber,
		ManagementIp: ap.ManagementIp,
		DeviceRole:   ap.DeviceRole,
		Manufacturer: ap.Manufacturer,
		DeviceModel:  ap.DeviceModel,
		OsVersion:    ap.OsVersion,
		Floor:        ap.Floor,
		GroupName:    ap.GroupName,
		Coordinate: &schemas.ApCoordinate{
			X: ap.Coordinate.Data().X,
			Y: ap.Coordinate.Data().Y,
			Z: ap.Coordinate.Data().Z,
		},
		ActiveWacId: ap.ActiveWacId,
		SiteId:      ap.SiteId,
	}, nil
}
