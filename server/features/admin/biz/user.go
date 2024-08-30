package biz

import (
	"context"

	"github.com/wangxin688/narvis/server/core/security"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/admin/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/global/constants"
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
		Username:       "Administrator",
		Email:          "admin@" + enterpriseCode + ".com",
		Password:       security.GetPasswordHash(password),
		RoleId:         roleId,
		AuthType:       uint8(constants.LocalTenantAuthType),
		OrganizationId: orgId,
	}

	newGroup := &models.Group{
		OrganizationId: orgId,
		Name:           constants.ReserveAdminGroupName,
		Description:    &constants.ReserveAdminGroupDescription,
		RoleId:         roleId,
		User:           []models.User{*newUser},
	}

	err = gen.Group.Create(newGroup)
	if err != nil {
		return nil, err
	}
	user := schemas.User{
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
		Id:        newUser.Id,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}
	return &user, nil
}

func (u *UserService) GetUserById(id string) (*schemas.User, error) {
	orgId := global.OrganizationId.Get()
	user, err := u.Where(gen.User.Id.Eq(id), gen.User.OrganizationId.Eq(orgId)).Preload(gen.User.Role).Preload(gen.User.Group).First()
	if err != nil {
		return nil, err
	}

	return &schemas.User{
		Username: user.Username,
		Email:    user.Email,
		AuthType: user.AuthType,
		Role: schemas.RoleShort{
			Id:   user.Role.Id,
			Name: user.Role.Name,
		},
		Group: schemas.GroupShort{
			Id:   user.Group.Id,
			Name: user.Group.Name,
		},
		Id:        user.Id,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (u *UserService) CreateUser(user *schemas.UserCreate) (*schemas.User, error) {
	newUser := &models.User{
		Username:       user.Username,
		Email:          user.Email,
		Password:       security.GetPasswordHash(user.Password),
		GroupId:        user.GroupId,
		RoleId:         user.RoleId,
		AuthType:       user.AuthType,
		Avatar:         user.Avatar,
		OrganizationId: global.OrganizationId.Get(),
	}

	err := gen.User.Create(newUser)
	if err != nil {
		return nil, err
	}

	return u.GetUserById(newUser.Id)
}

func (u *UserService) UpdateUser(userId string, user *schemas.UserUpdate) error {
	var updateFields = make(map[string]any)
	if user.Username != nil {
		updateFields["username"] = *user.Username
	}
	if user.Email != nil {
		updateFields["email"] = *user.Email
	}
	if user.Password != nil {
		updateFields["password"] = security.GetPasswordHash(*user.Password)
	}
	if user.Avatar != nil {
		updateFields["avatar"] = *user.Avatar
	}
	if user.GroupId != nil {
		updateFields["group_id"] = *user.GroupId
	}
	if user.RoleId != nil {
		updateFields["role_id"] = *user.RoleId
	}
	if user.Status != nil {
		updateFields["status"] = *user.Status
	}
	_, err := gen.User.Select(gen.User.Id.Eq(userId), gen.User.OrganizationId.Eq(global.OrganizationId.Get())).Updates(updateFields)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) DeleteUser(ctx context.Context, userId string) error {
	_, err := gen.User.WithContext(ctx).Select(gen.User.Id.Eq(userId), gen.User.OrganizationId.Eq(global.OrganizationId.Get())).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) ListUsers(params *schemas.UserQuery) (int64, *schemas.UserList, error) {
	orgId := global.OrganizationId.Get()
	stmt := gen.User.Where(gen.User.OrganizationId.Eq(orgId))
	if params.Id != nil {
		stmt = stmt.Where(gen.User.Id.In(*params.Id...))
	}
	if params.Username != nil {
		stmt = stmt.Where(gen.User.Username.In(*params.Username...))
	}
	if params.Email != nil {
		stmt = stmt.Where(gen.User.Email.In(*params.Email...))
	}
	if params.Status != nil {
		stmt = stmt.Where(gen.User.Status.Eq(*params.Status))
	}
	if params.GroupId != nil {
		stmt = stmt.Where(gen.User.GroupId.In(*params.GroupId...))
	}
	if params.RoleId != nil {
		stmt = stmt.Where(gen.User.RoleId.In(*params.RoleId...))
	}
	if params.AuthType != nil {
		stmt = stmt.Where(gen.User.AuthType.Eq(*params.AuthType))
	}

	if params.Keyword != nil {
		stmt.UnderlyingDB().Scopes(params.Search(models.UserSearchFields))
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

	users, err := stmt.Preload(gen.User.Role).Preload(gen.User.Group).Find()
	if err != nil {
		return 0, nil, err
	}
	var userList schemas.UserList
	for _, user := range users {
		userList = append(userList, &schemas.User{
			Username: user.Username,
			Email:    user.Email,
			AuthType: user.AuthType,
			Role: schemas.RoleShort{
				Id:   user.Role.Id,
				Name: user.Role.Name,
			},
			Group: schemas.GroupShort{
				Id:   user.Group.Id,
				Name: user.Group.Name,
			},
			Id:        user.Id,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return count, &userList, nil

}

func VerifyUser(userId string) *models.User {

	user, err := gen.User.Where(gen.User.Id.Eq(userId)).Preload(gen.User.Role).Preload(gen.User.Organization).First()
	if err != nil || user == nil {
		return nil
	}
	if user.Status != "Active" || !user.Organization.Active {
		return nil
	}

	return user
}

func (u *UserService) GetUserMe(userId string) (*schemas.User, error) {
	user, err := u.GetUserById(userId)
	if err != nil {
		return nil, err
	}
	return user, nil
}
