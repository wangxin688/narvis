package biz

import (
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/global/constants"
	"github.com/wangxin688/narvis/server/models"
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
