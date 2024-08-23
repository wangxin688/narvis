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
func (u *UserService) CreateAdminUser(enterpriseCode string, orgID string, password string) (*schemas.User, error) {
	roleService := NewRoleService()
	roleID, err := roleService.CreateAdminRole(orgID)
	if err != nil {
		return nil, err
	}
	newUser := &models.User{
		Username: "Administrator",
		Email:    "admin@" + enterpriseCode + ".com",
		Password: password,
		Group:    models.Group{OrganizationID: orgID, Name: "Administrator Group", RoleID: roleID},
		RoleID:   roleID,
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
			ID:   newUser.Group.ID,
			Name: newUser.Group.Name,
		},
		Role: schemas.RoleShort{
			ID:   newUser.Role.ID,
			Name: newUser.Role.Name,
		},
	}, nil
}

func (u *UserService) GetUserByID(id string) (*schemas.User, error) {
	orgID := global.OrganizationID.Get()
	user, err := u.Where(gen.User.ID.Eq(id), gen.User.OrganizationID.Eq(orgID)).First()
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
			ID:   user.Group.ID,
			Name: user.Group.Name,
		},
		Role: schemas.RoleShort{
			ID:   user.Role.ID,
			Name: user.Role.Name,
		},
	}, nil
}
