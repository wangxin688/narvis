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

func (r *RoleService) CreateAdminRole(organizationId string) (string, error) {
	role := &models.Role{
		Name:           constants.ReserveAdminRoleName,
		Description:    &constants.ReserveAdminRoleDescription,
		OrganizationId: organizationId,
	}
	err := r.Create(role)
	if err != nil {
		return "", err
	}
	return role.Id, nil
}


// func CheckRolePathPermission(roleId string, path string) (bool, error) {
// 	models.Menu
// }