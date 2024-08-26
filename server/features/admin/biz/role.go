package biz

import (
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/admin/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/global/constants"
	"github.com/wangxin688/narvis/server/infra"
	"github.com/wangxin688/narvis/server/models"
	"gorm.io/gorm"
)

type RoleService struct{}

func NewRoleService() *RoleService {
	return &RoleService{}
}

func (r *RoleService) CreateAdminRole(organizationID string) (string, error) {
	role := &models.Role{
		Name:           constants.ReserveAdminRoleName,
		Description:    &constants.ReserveAdminRoleDescription,
		OrganizationID: organizationID,
	}
	err := gen.Role.Create(role)
	if err != nil {
		return "", err
	}
	return role.ID, nil
}

// use raw sql for performance
// add redis cache if needed in the future
func CheckRolePathPermission(user *models.User, path string) bool {
	if user.Role.Name == constants.ReserveAdminRoleName {
		return true
	}
	roleID := user.Role.ID
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

func (r *RoleService) ListRoles(params *schemas.RoleQuery) (int64, *schemas.RoleList, error) {
	stmt := gen.Role.Where(gen.Role.OrganizationID.Eq(global.OrganizationID.Get()))
	if params.ID != nil {
		stmt = stmt.Where(gen.Role.ID.In(*params.ID...))
	}
	if params.Name != nil {
		stmt = stmt.Where(gen.Role.Name.In(*params.Name...))
	}

	if params.Keyword != nil {
		stmt.UnderlyingDB().Scopes(params.Search(models.RoleSearchFields))
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
	roles, err := stmt.Find()
	if err != nil {
		return 0, nil, err
	}
	var res schemas.RoleList
	for _, role := range roles {
		res = append(res, schemas.Role{
			ID:          role.ID,
			Name:        role.Name,
			Description: role.Description,
			CreatedAt:   role.CreatedAt,
			UpdatedAt:   role.UpdatedAt,
		})
	}
	return count, &res, nil

}

func (r *RoleService) GetRoleByID(id string) (*schemas.RoleDetail, error) {
	role, err := gen.Role.Where(gen.Role.ID.Eq(id), gen.Role.OrganizationID.Eq(global.OrganizationID.Get())).Preload(gen.Role.Menus).First()
	if err != nil {
		return nil, err
	}
	// TODO: add menu tree to the response
	// if role.Name == constants.ReserveAdminRoleName {
	// 	menus, err := gen.Menu.Find()
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// }

	return &schemas.RoleDetail{
		ID:          role.ID,
		Name:        role.Name,
		Description: role.Description,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}, nil
}

func (r *RoleService) CreateRole(role *schemas.RoleCreate) (string, error) {
	menus, err := gen.Menu.Where(gen.Menu.ID.In(role.Menus...)).Find()
	if err != nil {
		return "", err
	}
	roleModel := &models.Role{
		Name:           role.Name,
		Description:    role.Description,
		OrganizationID: global.OrganizationID.Get(),
	}
	err = gen.Role.Create(roleModel)
	if err != nil {
		return "", err
	}
	gen.Role.Menus.Model(roleModel).Replace(menus...)
	return roleModel.ID, nil
}

func (r *RoleService) UpdateRole(RoleID string, role *schemas.RoleUpdate) error {
	var updateFields = make(map[string]any)
	if role.Name != nil {
		updateFields["name"] = *role.Name
	}
	if role.Description != nil {
		updateFields["description"] = *role.Description
	}
	gen.Role.UnderlyingDB().Transaction(func(tx *gorm.DB) error {
		dbRole, err := gen.Role.Where(gen.Role.ID.Eq(RoleID), gen.Role.OrganizationID.Eq(global.OrganizationID.Get())).First()
		if err != nil {
			return err
		}
		if err := tx.Model(dbRole).Updates(updateFields).Error; err != nil {
			return err
		}
		if role.Menus != nil {
			if len(*role.Menus) == 0 {
				gen.Role.Menus.Model(dbRole).Clear()
			} else {
				dbMenus, err := gen.Menu.Where(gen.Menu.ID.In(*role.Menus...)).Find()
				if err != nil {
					return err
				}
				gen.Role.Menus.Model(dbRole).Replace(dbMenus...)
			}
		}
		gen.Role.UnderlyingDB().Save(updateFields)
		return nil
	})
	return nil
}

func (r *RoleService) DeleteRole(id string) error {
	_, err := gen.Role.Where(gen.Role.ID.Eq(id), gen.Role.OrganizationID.Eq(global.OrganizationID.Get())).Delete()
	if err != nil {
		return err
	}
	return nil
}
