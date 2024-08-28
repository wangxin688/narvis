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
func (u *UserService) CreateAdminUser(enterpriseCode string, orgID string, password string) (*schemas.User, error) {
	roleService := NewRoleService()
	roleID, err := roleService.CreateAdminRole(orgID)
	if err != nil {
		return nil, err
	}
	newUser := &models.User{
		Username:       "Administrator",
		Email:          "admin@" + enterpriseCode + ".com",
		Password:       security.GetPasswordHash(password),
		RoleID:         roleID,
		AuthType:       uint8(constants.LocalTenantAuthType),
		OrganizationID: orgID,
	}

	newGroup := &models.Group{
		OrganizationID: orgID,
		Name:           constants.ReserveAdminGroupName,
		Description:    &constants.ReserveAdminGroupDescription,
		RoleID:         roleID,
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
			ID:   newUser.Group.ID,
			Name: newUser.Group.Name,
		},
		Role: schemas.RoleShort{
			ID:   newUser.Role.ID,
			Name: newUser.Role.Name,
		},
		ID:        newUser.ID,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}
	return &user, nil
}

func (u *UserService) GetUserByID(id string) (*schemas.User, error) {
	orgID := global.OrganizationID.Get()
	user, err := u.Where(gen.User.ID.Eq(id), gen.User.OrganizationID.Eq(orgID)).Preload(gen.User.Role).Preload(gen.User.Group).First()
	if err != nil {
		return nil, err
	}

	return &schemas.User{
		Username: user.Username,
		Email:    user.Email,
		AuthType: user.AuthType,
		Role: schemas.RoleShort{
			ID:   user.Role.ID,
			Name: user.Role.Name,
		},
		Group: schemas.GroupShort{
			ID:   user.Group.ID,
			Name: user.Group.Name,
		},
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (u *UserService) CreateUser(user *schemas.UserCreate) (*schemas.User, error) {
	newUser := &models.User{
		Username:       user.Username,
		Email:          user.Email,
		Password:       security.GetPasswordHash(user.Password),
		GroupID:        user.GroupID,
		RoleID:         user.RoleID,
		AuthType:       user.AuthType,
		Avatar:         user.Avatar,
		OrganizationID: global.OrganizationID.Get(),
	}

	err := gen.User.Create(newUser)
	if err != nil {
		return nil, err
	}

	return u.GetUserByID(newUser.ID)
}

func (u *UserService) UpdateUser(userID string, user *schemas.UserUpdate) error {
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
	if user.GroupID != nil {
		updateFields["group_id"] = *user.GroupID
	}
	if user.RoleID != nil {
		updateFields["role_id"] = *user.RoleID
	}
	if user.Status != nil {
		updateFields["status"] = *user.Status
	}
	_, err := gen.User.Select(gen.User.ID.Eq(userID), gen.User.OrganizationID.Eq(global.OrganizationID.Get())).Updates(updateFields)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) DeleteUser(ctx context.Context, userID string) error {
	_, err := gen.User.WithContext(ctx).Select(gen.User.ID.Eq(userID), gen.User.OrganizationID.Eq(global.OrganizationID.Get())).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) ListUsers(params *schemas.UserQuery) (int64, *schemas.UserList, error) {
	orgID := global.OrganizationID.Get()
	stmt := gen.User.Where(gen.User.OrganizationID.Eq(orgID))
	if params.ID != nil {
		stmt = stmt.Where(gen.User.ID.In(*params.ID...))
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
	if params.GroupID != nil {
		stmt = stmt.Where(gen.User.GroupID.In(*params.GroupID...))
	}
	if params.RoleID != nil {
		stmt = stmt.Where(gen.User.RoleID.In(*params.RoleID...))
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
				ID:   user.Role.ID,
				Name: user.Role.Name,
			},
			Group: schemas.GroupShort{
				ID:   user.Group.ID,
				Name: user.Group.Name,
			},
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return count, &userList, nil

}

func VerifyUser(userID string) *models.User {

	user, err := gen.User.Where(gen.User.ID.Eq(userID)).Preload(gen.User.Role).Preload(gen.User.Organization).First()
	if err != nil || user == nil {
		return nil
	}
	if user.Status != "Active" || !user.Organization.Active {
		return nil
	}

	return user
}


func (u *UserService) GetUserMe(userID string) (*schemas.User, error) {
	user, err := u.GetUserByID(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}