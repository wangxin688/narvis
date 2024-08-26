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

func (g *GroupService) CreateAdminGroup(organizationID string, roleID string) (string, error) {

	group := &models.Group{
		OrganizationID: organizationID,
		Name:           constants.ReserveAdminGroupName,
		Description:    &constants.ReserveAdminGroupDescription,
		RoleID:         roleID,
	}
	err := gen.Group.Create(group)
	if err != nil {
		return "", err
	}
	return group.ID, nil
}

func (g *GroupService) CreateGroup(group *schemas.GroupCreate) (string, error) {

	if group.Name == constants.ReserveAdminGroupName {
		return "", errors.NewError(errors.CodeInvalidGroupNameForReserve, errors.MsgInvalidGroupNameForReserve)
	}
	gp := &models.Group{
		Name:           group.Name,
		Description:    group.Description,
		OrganizationID: global.OrganizationID.Get(),
		RoleID:         group.RoleID,
	}

	err := gen.Group.Create(gp)
	if err != nil {
		return "", err
	}
	return gp.ID, nil
}

func (g *GroupService) UpdateGroup(groupID string, group *schemas.GroupUpdate) error {
	var updateFields = make(map[string]string)
	if group.Name != nil {
		updateFields["name"] = *group.Name
	}
	if group.Description != nil {
		updateFields["description"] = *group.Description
	}
	if group.RoleID != nil {
		updateFields["role_id"] = *group.RoleID
	}
	_, err := gen.Group.Where(gen.Group.ID.Eq(groupID), gen.Group.OrganizationID.Eq(global.OrganizationID.Get())).Updates(updateFields)
	if err != nil {
		return err
	}
	return nil
}

func (g *GroupService) GetGroupByID(id string) (*schemas.Group, error) {
	group, err := gen.Group.Where(gen.Group.ID.Eq(id), gen.Group.OrganizationID.Eq(global.OrganizationID.Get())).
		Preload(gen.Group.Role).First()
	if err != nil {
		return nil, err
	}
	return &schemas.Group{
		GroupShort: schemas.GroupShort{
			ID:   group.ID,
			Name: group.Name,
		},
		Description: group.Description,
		CreatedAt:   group.CreatedAt,
		UpdatedAt:   group.UpdatedAt,
		Role: schemas.RoleShort{
			ID:   group.Role.ID,
			Name: group.Role.Name,
		},
	}, nil
}

func (g *GroupService) ListGroups(params *schemas.GroupQuery) (int64, *schemas.GroupList, error) {
	stmt := gen.Group.Where(gen.Group.OrganizationID.Eq(global.OrganizationID.Get()))
	if params.ID != nil {
		stmt = stmt.Where(gen.Group.ID.In(*params.ID...))
	}
	if params.Name != nil {
		stmt = stmt.Where(gen.Group.Name.In(*params.Name...))
	}
	if params.RoleID != nil {
		stmt = stmt.Where(gen.Group.RoleID.In(*params.RoleID...))
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
				ID:   group.ID,
				Name: group.Name,
			},
			Description: group.Description,
			CreatedAt:   group.CreatedAt,
			UpdatedAt:   group.UpdatedAt,
			Role: schemas.RoleShort{
				ID:   group.Role.ID,
				Name: group.Role.Name,
			},
		})
	}
	return count, &res, nil
}

func (g *GroupService) DeleteGroup(id string) error {
	_, err := gen.Group.Where(gen.Group.ID.Eq(id), gen.Group.OrganizationID.Eq(global.OrganizationID.Get())).Delete()
	if err != nil {
		return err
	}
	return nil
}
