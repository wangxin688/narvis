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

// use raw sql for performance
// add redis cache if needed in the future
func CheckRolePathPermission(user *models.User, path string) bool {
	if user.Role.Name == constants.ReserveAdminRoleName {
		return true
	}
	roleId := user.Role.Id
	var result string
	stmt := `
	SELECT rm.roleId  
	FROM role_menus rm  
	JOIN menus m ON rm.menuId = m.id  
	JOIN menu_permission mp ON m.id = mp.menuId  
	JOIN permission p ON mp.permissionId = p.id  
	WHERE rm.roleId = ? AND p.path = ?  
	LIMIT 1;
	`
	infra.DB.Raw(stmt, roleId, path).Scan(&result)
	return result != ""
}

func (r *RoleService) ListRoles(params *schemas.RoleQuery) (int64, *[]*schemas.Role, error) {
	res := make([]*schemas.Role, 0)
	stmt := gen.Role.Where(gen.Role.OrganizationId.Eq(global.OrganizationId.Get()))
	if params.Id != nil {
		stmt = stmt.Where(gen.Role.Id.In(*params.Id...))
	}
	if params.Name != nil {
		stmt = stmt.Where(gen.Role.Name.In(*params.Name...))
	}

	if params.IsSearchable() {
		searchString := "%" + *params.Keyword + "%"
		stmt = stmt.Where(
			gen.Role.Name.Like(searchString),
		).Or(gen.Role.Description.Like(searchString))
	}
	count, err := stmt.Count()
	if err != nil || count <= 0 {
		return 0, &res, err
	}
	stmt.UnderlyingDB().Scopes(params.OrderByField())
	stmt.UnderlyingDB().Scopes(params.Pagination())
	roles, err := stmt.Find()
	if err != nil {
		return 0, &res, err
	}
	for _, role := range roles {
		res = append(res, &schemas.Role{
			Id:          role.Id,
			Name:        role.Name,
			Description: role.Description,
			CreatedAt:   role.CreatedAt,
			UpdatedAt:   role.UpdatedAt,
		})
	}
	return count, &res, nil

}

func (r *RoleService) GetRoleById(id string) (*schemas.RoleDetail, error) {
	role, err := gen.Role.Where(gen.Role.Id.Eq(id), gen.Role.OrganizationId.Eq(global.OrganizationId.Get())).Preload(gen.Role.Menus).First()
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
		Id:          role.Id,
		Name:        role.Name,
		Description: role.Description,
		CreatedAt:   role.CreatedAt,
		UpdatedAt:   role.UpdatedAt,
	}, nil
}

func (r *RoleService) CreateRole(role *schemas.RoleCreate) (string, error) {
	menus, err := gen.Menu.Where(gen.Menu.Id.In(role.Menus...)).Find()
	if err != nil {
		return "", err
	}
	roleModel := &models.Role{
		Name:           role.Name,
		Description:    role.Description,
		OrganizationId: global.OrganizationId.Get(),
	}
	err = gen.Role.UnderlyingDB().Transaction(func(_ *gorm.DB) error {
		err = gen.Role.Create(roleModel)
		if err != nil {
			return err
		}
		err = gen.Role.Menus.Model(roleModel).Replace(menus...)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return roleModel.Id, nil
}

func (r *RoleService) UpdateRole(RoleId string, role *schemas.RoleUpdate) error {
	var updateFields = make(map[string]any)
	if role.Name != nil {
		updateFields["name"] = *role.Name
	}
	if role.Description != nil {
		updateFields["description"] = *role.Description
	}
	err := gen.Role.UnderlyingDB().Transaction(func(tx *gorm.DB) error {
		dbRole, err := gen.Role.Where(gen.Role.Id.Eq(RoleId), gen.Role.OrganizationId.Eq(global.OrganizationId.Get())).First()
		if err != nil {
			return err
		}
		if err := tx.Model(dbRole).Updates(updateFields).Error; err != nil {
			return err
		}
		if role.Menus != nil {
			if len(*role.Menus) == 0 {
				err = gen.Role.Menus.Model(dbRole).Clear()
				if err != nil {
					return err
				}
			} else {
				dbMenus, err := gen.Menu.Where(gen.Menu.Id.In(*role.Menus...)).Find()
				if err != nil {
					return err
				}
				err = gen.Role.Menus.Model(dbRole).Replace(dbMenus...)
				if err != nil {
					return err
				}
			}
		}
		gen.Role.UnderlyingDB().Save(updateFields)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *RoleService) DeleteRole(id string) error {
	_, err := gen.Role.Where(gen.Role.Id.Eq(id), gen.Role.OrganizationId.Eq(global.OrganizationId.Get())).Delete()
	if err != nil {
		return err
	}
	return nil
}
