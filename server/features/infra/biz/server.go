package infra_biz

import (
	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"github.com/wangxin688/narvis/intend/logger"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	infra_utils "github.com/wangxin688/narvis/server/features/infra/utils"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/pkg/contextvar"
	"github.com/wangxin688/narvis/server/tools/errors"
	"github.com/wangxin688/narvis/server/tools/helpers"
	"go.uber.org/zap"
)

type ServerService struct{}

func NewServerService() *ServerService {
	return &ServerService{}
}

func (s *ServerService) CreateServer(server *schemas.ServerCreate) (string, error) {
	orgId := contextvar.OrganizationId.Get()
	ok, err := LicenseUsageDepends(1, orgId)
	if err != nil {
		return "", err
	}
	if !ok {
		return "", errors.NewError(errors.CodeLicenseCountExceeded, errors.MsgLicenseCountExceeded)
	}
	newServer := models.Server{
		Name:           server.Name,
		ManagementIp:   server.ManagementIp,
		Status:         server.Status,
		Manufacturer:   *server.Manufacturer,
		OsVersion:      *server.OsVersion,
		SiteId:         server.SiteId,
		OrganizationId: orgId,
		Cpu:            server.Cpu,
		Memory:         server.Memory,
		Disk:           server.Disk,
		Description:    server.Description,
	}
	if server.RackId != nil && server.RackPosition != nil {
		newServer.RackId = server.RackId
		position, err := infra_utils.SliceUint8ToString(*server.RackPosition)
		if err != nil {
			return "", err
		}
		rackService := NewRackService()
		rack, err := rackService.GetRackById(*server.RackId)
		if err != nil {
			return "", err
		}
		if !rackService.ValidateCreateRackReservation(*server.RackId, rack.UHeight, *server.RackPosition) {
			return "", errors.NewError(errors.CodeRackPositionOccupied, errors.MsgRackPositionOccupied)
		}
		newServer.RackPosition = &position
	}
	err = gen.Server.Create(&newServer)
	if err != nil {
		return "", err
	}
	return newServer.Id, nil
}

func (s *ServerService) UpdateServer(g *gin.Context, serverId string, server *schemas.ServerUpdate) (diff map[string]map[string]*contextvar.Diff, err error) {
	dbServer, err := gen.Server.Where(gen.Server.Id.Eq(serverId), gen.Server.OrganizationId.Eq(contextvar.OrganizationId.Get())).First()
	if err != nil {
		return nil, err
	}
	updateFields := make(map[string]*contextvar.Diff)
	if server.Name != nil && *server.Name != dbServer.Name {
		updateFields["name"] = &contextvar.Diff{Before: dbServer.Name, After: *server.Name}
		dbServer.Name = *server.Name
	}
	if server.Status != nil && *server.Status != dbServer.Status {
		updateFields["status"] = &contextvar.Diff{Before: dbServer.Status, After: *server.Status}
		dbServer.Status = *server.Status
	}
	if server.Manufacturer != nil && *server.Manufacturer != dbServer.Manufacturer {
		updateFields["manufacturer"] = &contextvar.Diff{Before: dbServer.Manufacturer, After: *server.Manufacturer}
		dbServer.Manufacturer = *server.Manufacturer
	}
	if server.OsVersion != nil && *server.OsVersion != dbServer.OsVersion {
		updateFields["osVersion"] = &contextvar.Diff{Before: dbServer.OsVersion, After: *server.OsVersion}
		dbServer.OsVersion = *server.OsVersion
	}
	if server.Description != nil && server.Description != dbServer.Description {
		updateFields["description"] = &contextvar.Diff{Before: dbServer.Description, After: *server.Description}
		dbServer.Description = server.Description
	}
	if server.Cpu != nil && *server.Cpu != dbServer.Cpu {
		updateFields["cpu"] = &contextvar.Diff{Before: dbServer.Cpu, After: *server.Cpu}
		dbServer.Cpu = *server.Cpu
	}
	if server.Memory != nil && *server.Memory != dbServer.Memory {
		updateFields["memory"] = &contextvar.Diff{Before: dbServer.Memory, After: *server.Memory}
		dbServer.Memory = *server.Memory
	}
	if server.Disk != nil && *server.Disk != dbServer.Disk {
		updateFields["disk"] = &contextvar.Diff{Before: dbServer.Disk, After: *server.Disk}
		dbServer.Disk = *server.Disk
	}
	if helpers.HasParams(g, "rackId") && server.RackId != dbServer.RackId {
		updateFields["rackId"] = &contextvar.Diff{Before: dbServer.RackId, After: *server.RackId}
		dbServer.RackId = server.RackId
	}
	if helpers.HasParams(g, "rackPosition") {
		position, err := infra_utils.SliceUint8ToString(*server.RackPosition)
		if err != nil {
			return nil, err
		}
		if position != *dbServer.RackPosition {
			updateFields["rackPosition"] = &contextvar.Diff{Before: dbServer.RackPosition, After: position}
			dbServer.RackPosition = &position
		}
	}

	if len(updateFields) == 0 {
		return nil, nil
	}
	diffValue := make(map[string]map[string]*contextvar.Diff)
	diffValue[serverId] = updateFields
	contextvar.OrmDiff.Set(diffValue)
	err = gen.Server.UnderlyingDB().Save(dbServer).Error
	if err != nil {
		return nil, err
	}
	return diffValue, nil
}

func (s *ServerService) DeleteServer(serverId string) (*models.Server, error) {
	dbServer, err := gen.Server.Where(gen.Server.Id.Eq(serverId), gen.Server.OrganizationId.Eq(contextvar.OrganizationId.Get())).First()
	if err != nil {
		return nil, err
	}
	_, err = gen.Server.Delete(dbServer)
	if err != nil {
		return nil, err
	}
	return dbServer, nil
}

func (s *ServerService) GetAllServerIdsBySiteId(siteId string) ([]string, error) {
	var result []string
	err := gen.Server.Select(gen.Server.Id).Where(gen.Server.SiteId.Eq(siteId), gen.Server.OrganizationId.Eq(contextvar.OrganizationId.Get())).Scan(&result)
	return result, err
}

func (s *ServerService) GetById(serverId string) (*schemas.Server, error) {
	orgId := contextvar.OrganizationId.Get()
	server, err := gen.Server.Where(gen.Server.Id.Eq(serverId), gen.Server.OrganizationId.Eq(orgId)).First()
	if err != nil {
		return nil, err
	}
	opStatus, err := GetServerOpStatus([]string{serverId}, orgId)
	if err != nil {
		logger.Logger.Error("[infraDeviceService]: failed to get server op status", zap.Error(err))
	}
	return &schemas.Server{
		Id:           server.Id,
		Name:         server.Name,
		Status:       server.Status,
		ManagementIp: server.ManagementIp,
		OperStatus: func() string {
			if value, ok := opStatus[serverId]; ok {
				return value
			}
			return "nodata"
		}(),
		Manufacturer: server.Manufacturer,
		OsVersion:    server.OsVersion,
		Description:  server.Description,
		Cpu:          server.Cpu,
		Memory:       server.Memory,
		Disk:         server.Disk,
		RackId:       server.RackId,
		RackPosition: func() *[]uint8 {
			if server.RackPosition == nil {
				return nil
			}
			position, _ := infra_utils.ParseUint8s(*server.RackPosition)
			return &position
		}(),
		SiteId: server.SiteId,
	}, nil
}

func (s *ServerService) GetServerList(query *schemas.ServerQuery) (int64, *[]*schemas.Server, error) {
	orgId := contextvar.OrganizationId.Get()
	res := make([]*schemas.Server, 0)
	stmt := gen.Server.Where(gen.Server.OrganizationId.Eq(orgId))
	if query.SiteId != nil {
		stmt = stmt.Where(gen.Server.SiteId.Eq(*query.SiteId))
	}
	if query.Name != nil {
		stmt = stmt.Where(gen.Server.Name.In(*query.Name...))
	}
	if query.Status != nil {
		stmt = stmt.Where(gen.Server.Status.Eq(*query.Status))
	}
	if query.Manufacturer != nil {
		stmt = stmt.Where(gen.Server.Manufacturer.In(*query.Manufacturer...))
	}
	if query.OsVersion != nil {
		stmt = stmt.Where(gen.Server.OsVersion.In(*query.OsVersion...))
	}
	if query.RackId != nil {
		stmt = stmt.Where(gen.Server.RackId.Eq(*query.RackId))
	}
	if query.ManagementIp != nil {
		stmt = stmt.Where(gen.Server.ManagementIp.In(*query.ManagementIp...))
	}
	if query.IsSearchable() {
		searchString := "%" + *query.Keyword + "%"
		stmt = stmt.Where(gen.Server.Name.Like(searchString)).
			Or(gen.Server.ManagementIp.Like(searchString))
	}
	count, err := stmt.Count()
	if err != nil || count <= 0 {
		return 0, &res, err
	}
	stmt.UnderlyingDB().Scopes(query.OrderByField())
	stmt.UnderlyingDB().Scopes(query.Pagination())
	list, err := stmt.Find()
	if err != nil {
		return 0, &res, err
	}
	serverIds := lo.Map(list, func(item *models.Server, _ int) string {
		return item.Id
	})
	opStatus, err := GetServerOpStatus(serverIds, orgId)
	if err != nil {
		logger.Logger.Error("[infraDeviceService]: failed to get server op status", zap.Error(err))
	}
	for _, server := range list {
		res = append(res, &schemas.Server{
			Id:           server.Id,
			Name:         server.Name,
			Status:       server.Status,
			ManagementIp: server.ManagementIp,
			OperStatus: func() string {
				if value, ok := opStatus[server.Id]; ok {
					return value
				}
				return "nodata"
			}(),
			Manufacturer: server.Manufacturer,
			OsVersion:    server.OsVersion,
			Description:  server.Description,
			Cpu:          server.Cpu,
			Memory:       server.Memory,
			Disk:         server.Disk,
			RackId:       server.RackId,
			RackPosition: func() *[]uint8 {
				if server.RackPosition == nil {
					return nil
				}
				position, _ := infra_utils.ParseUint8s(*server.RackPosition)
				return &position
			}(),
			SiteId: server.SiteId,
		})
	}
	return count, &res, nil
}
