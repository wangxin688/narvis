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
