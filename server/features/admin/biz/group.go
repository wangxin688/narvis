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
	gen.IGroupDo
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
	err := g.Create(group)
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

	err := g.IGroupDo.Create(gp)
	if err != nil {
		return "", err
	}
	return gp.ID, nil
}
