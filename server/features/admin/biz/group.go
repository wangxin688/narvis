package biz

import (
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/admin/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/global/constants"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/tools/errors"
)

type GroupService struct {
}

func NewGroupService() *GroupService {
	return &GroupService{}
}

func (g *GroupService) CreateAdminGroup(organizationId string, roleId string) (string, error) {

	group := &models.Group{
		OrganizationId: organizationId,
		Name:           constants.ReserveAdminGroupName,
		Description:    &constants.ReserveAdminGroupDescription,
		RoleId:         roleId,
	}
	err := gen.Group.Create(group)
	if err != nil {
		return "", err
	}
	return group.Id, nil
}

func (g *GroupService) CreateGroup(group *schemas.GroupCreate) (string, error) {

	if group.Name == constants.ReserveAdminGroupName {
		return "", errors.NewError(errors.CodeInvalidGroupNameForReserve, errors.MsgInvalidGroupNameForReserve)
	}
	gp := &models.Group{
		Name:           group.Name,
		Description:    group.Description,
		OrganizationId: global.OrganizationId.Get(),
		RoleId:         group.RoleId,
	}

	err := gen.Group.Create(gp)
	if err != nil {
		return "", err
	}
	return gp.Id, nil
}

func (g *GroupService) UpdateGroup(groupId string, group *schemas.GroupUpdate) error {
	var updateFields = make(map[string]string)
	if group.Name != nil {
		updateFields["name"] = *group.Name
	}
	if group.Description != nil {
		updateFields["description"] = *group.Description
	}
	if group.RoleId != nil {
		updateFields["role_id"] = *group.RoleId
	}
	_, err := gen.Group.Where(gen.Group.Id.Eq(groupId), gen.Group.OrganizationId.Eq(global.OrganizationId.Get())).Updates(updateFields)
	if err != nil {
		return err
	}
	return nil
}

func (g *GroupService) GetGroupById(id string) (*schemas.Group, error) {
	group, err := gen.Group.Where(gen.Group.Id.Eq(id), gen.Group.OrganizationId.Eq(global.OrganizationId.Get())).
		Preload(gen.Group.Role).First()
	if err != nil {
		return nil, err
	}
	return &schemas.Group{
		GroupShort: schemas.GroupShort{
			Id:   group.Id,
			Name: group.Name,
		},
		Description: group.Description,
		CreatedAt:   group.CreatedAt,
		UpdatedAt:   group.UpdatedAt,
		Role: schemas.RoleShort{
			Id:   group.Role.Id,
			Name: group.Role.Name,
		},
	}, nil
}

func (g *GroupService) ListGroups(params *schemas.GroupQuery) (int64, *schemas.GroupList, error) {
	stmt := gen.Group.Where(gen.Group.OrganizationId.Eq(global.OrganizationId.Get()))
	if params.Id != nil {
		stmt = stmt.Where(gen.Group.Id.In(*params.Id...))
	}
	if params.Name != nil {
		stmt = stmt.Where(gen.Group.Name.In(*params.Name...))
	}
	if params.RoleId != nil {
		stmt = stmt.Where(gen.Group.RoleId.In(*params.RoleId...))
	}
	if params.Keyword != nil {
		stmt.UnderlyingDB().Scopes(params.Search(models.GroupSearchFields))
	}
	count, err := stmt.Count()
	if err != nil {
		return 0, nil, err
	}
	if count <= 0 {
		return 0, nil, nil
	}
	if params.OrderBy != nil {
		stmt.UnderlyingDB().Scopes(params.OrderByField())
	}
	stmt.UnderlyingDB().Scopes(params.LimitOffset())
	groups, err := stmt.Preload(gen.Group.Role).Find()
	if err != nil {
		return 0, nil, err
	}
	var res schemas.GroupList
	for _, group := range groups {
		res = append(res, &schemas.Group{
			GroupShort: schemas.GroupShort{
				Id:   group.Id,
				Name: group.Name,
			},
			Description: group.Description,
			CreatedAt:   group.CreatedAt,
			UpdatedAt:   group.UpdatedAt,
			Role: schemas.RoleShort{
				Id:   group.Role.Id,
				Name: group.Role.Name,
			},
		})
	}
	return count, &res, nil
}

func (g *GroupService) DeleteGroup(id string) error {
	_, err := gen.Group.Where(gen.Group.Id.Eq(id), gen.Group.OrganizationId.Eq(global.OrganizationId.Get())).Delete()
	if err != nil {
		return err
	}
	return nil
}
