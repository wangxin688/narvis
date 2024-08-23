package services

import (
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/admin/constants"
	"github.com/wangxin688/narvis/server/models"
)

type RoleService struct {
	gen.IRoleDo
}

func NewRoleService() *RoleService {
	return &RoleService{}
}

func (r *RoleService) CreateAdminRole(organizationID string) (string, error) {
	role := &models.Role{
		Name:           constants.ReserveAdminRoleName,
		Description:    &constants.ReserveAdminRoleDescription,
		OrganizationID: organizationID,
	}
	err := r.Create(role)
	if err != nil {
		return "", err
	}
	return role.ID, nil
}

// func CheckRolePathPermission(roleID string, path string) (bool, error) {
// 	models.Menu
// }
