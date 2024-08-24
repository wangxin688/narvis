package biz

import (
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/global/constants"
	"github.com/wangxin688/narvis/server/infra"
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

// use raw sql for performance
// add redis cache if needed in the future
func CheckRolePathPermission(roleID string, path string) bool {
	dbRole, err := gen.Q.Role.Where(gen.Role.ID.Eq(roleID)).First()
	if err != nil {
		return false
	}
	if dbRole.Name == constants.ReserveAdminRoleName {
		return true
	}
	var result string
	stmt := `
	SELECT rm.role_id  
	FROM role_menus rm  
	JOIN menus m ON rm.menu_id = m.id  
	JOIN menu_permissions mp ON m.id = mp.menu_id  
	JOIN permissions p ON mp.permission_id = p.id  
	WHERE rm.role_id = ? AND p.path = ?  
	LIMIT 1;
	`
	infra.DB.Raw(stmt, roleID, path).Scan(&result)
	return result != ""
}
