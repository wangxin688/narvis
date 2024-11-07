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
	newAdminRole := models.Role{
		Name:           constants.ReserveAdminRoleName,
		Description:    &constants.ReserveAdminRoleDescription,
		OrganizationId: orgId,
	}
	err := gen.Role.Create(&newAdminRole)
	if err != nil {
		return nil, err
	}
	newUser := &models.User{
		Username:       "Administrator",
		Email:          "admin@" + enterpriseCode + ".com",
		Password:       security.GetPasswordHash(password),
		RoleId:         newAdminRole.Id,
		AuthType:       uint8(constants.LocalTenantAuthType),
		OrganizationId: orgId,
	}

	err = gen.User.Create(newUser)
	if err != nil {
		return nil, err
	}

	user := schemas.User{
		Username: newUser.Username,
		Email:    newUser.Email,
		AuthType: newUser.AuthType,
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
	user, err := u.Where(gen.User.Id.Eq(id), gen.User.OrganizationId.Eq(orgId)).Preload(gen.User.Role).First()
	if err != nil {
		return nil, err
	}

	return &schemas.User{
		Username: user.Username,
		Email:    user.Email,
		AuthType: user.AuthType,
		Avatar:   user.Avatar,
		Role: schemas.RoleShort{
			Id:   user.Role.Id,
			Name: user.Role.Name,
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
	if user.RoleId != nil {
		updateFields["roleId"] = *user.RoleId
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

func (u *UserService) ListUsers(params *schemas.UserQuery) (int64, *[]*schemas.User, error) {
	res := make([]*schemas.User, 0)
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
	if params.RoleId != nil {
		stmt = stmt.Where(gen.User.RoleId.In(*params.RoleId...))
	}
	if params.AuthType != nil {
		stmt = stmt.Where(gen.User.AuthType.Eq(*params.AuthType))
	}

	if params.IsSearchable() {
		searchString := "%" + *params.Keyword + "%"
		stmt = stmt.Where(gen.User.Username.Like(searchString)).Or(gen.User.Email.Like(searchString))
	}
	count, err := stmt.Count()
	if err != nil || count <= 0 {
		return 0, &res, err
	}
	stmt.UnderlyingDB().Scopes(params.OrderByField())
	stmt.UnderlyingDB().Scopes(params.Pagination())

	users, err := stmt.Preload(gen.User.Role).Find()
	if err != nil {
		return 0, &res, err
	}
	for _, user := range users {
		res = append(res, &schemas.User{
			Username: user.Username,
			Email:    user.Email,
			AuthType: user.AuthType,
			Role: schemas.RoleShort{
				Id:   user.Role.Id,
				Name: user.Role.Name,
			},
			Id:        user.Id,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return count, &res, nil

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
