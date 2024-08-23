package services

import (
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/admin/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"
)

type UserService struct {
	gen.IUserDo
}

func NewUserService() *UserService {
	return &UserService{}
}

// when new organization created, create a new admin user
func (u *UserService) CreateAdminUser(enterpriseCode string, orgId string, password string) (*schemas.User, error) {
	roleService := NewRoleService()
	roleId, err := roleService.CreateAdminRole(orgId)
	if err != nil {
		return nil, err
	}
	newUser := &models.User{
		Username: "Administrator",
		Email:    "admin@" + enterpriseCode + ".com",
		Password: password,
		Group:    models.Group{OrganizationId: orgId, Name: "Administrator Group", RoleId: roleId},
		RoleId:   roleId,
		AuthType: 1,
	}

	err = u.Create(newUser)
	if err != nil {
		return nil, err
	}
	return &schemas.User{
		Username: newUser.Username,
		Email:    newUser.Email,
		AuthType: newUser.AuthType,
		Group: schemas.GroupShort{
			Id:   newUser.Group.Id,
			Name: newUser.Group.Name,
		},
		Role: schemas.RoleShort{
			Id:   newUser.Role.Id,
			Name: newUser.Role.Name,
		},
	}, nil
}

func (u *UserService) GetUserById(id string) (*schemas.User, error) {
	orgId := global.OrganizationId.Get()
	user, err := u.Where(gen.User.Id.Eq(id), gen.User.OrganizationId.Eq(orgId)).First()
	// Preload(gen.User.Group).
	// Preload(gen.User.Role).First()
	if err != nil {
		return nil, err
	}

	return &schemas.User{
		Username: user.Username,
		Email:    user.Email,
		AuthType: user.AuthType,
		Group: schemas.GroupShort{
			Id:   user.Group.Id,
			Name: user.Group.Name,
		},
		Role: schemas.RoleShort{
			Id:   user.Role.Id,
			Name: user.Role.Name,
		},
	}, nil
}
