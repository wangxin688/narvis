package biz

import (
	"github.com/wangxin688/narvis/server/core/security"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/admin/schemas"
	"github.com/wangxin688/narvis/server/tools/errors"
)

type RBACService struct {
}

func NewRBACService() *RBACService {
	return &RBACService{}
}

func (r *RBACService) PasswordLogin(login schemas.Oauth2PasswordRequest) (*security.AccessToken, error) {
	var orgID string
	if login.DomainName != nil || login.EnterpriseCode != nil {
		if login.DomainName != nil {
			org, err := gen.Organization.Select(gen.Organization.ID, gen.Organization.Active).Where(gen.Organization.DomainName.Eq(*login.DomainName)).First()
			if err != nil {
				return nil, errors.NewError(errors.CodeNotFound, errors.MsgNotFound, "organization", "domain_name", *login.DomainName)
			}
			if !org.Active {
				return nil, errors.NewError(errors.CodeNotFound, errors.MsgNotFound, "organization", "domain_name", *login.DomainName)
			}
			orgID = org.ID
		}
		if login.EnterpriseCode != nil {
			org, err := gen.Organization.Select(gen.Organization.ID, gen.Organization.Active).Where(gen.Organization.EnterpriseCode.Eq(*login.EnterpriseCode)).First()
			if err != nil {
				return nil, errors.NewError(errors.CodeNotFound, errors.MsgNotFound, "organization", "enterprise_code", *login.EnterpriseCode)
			}
			if !org.Active {
				return nil, errors.NewError(errors.CodeNotFound, errors.MsgNotFound, "organization", "enterprise_code", *login.EnterpriseCode)
			}
			orgID = org.ID
		}
	}
	if orgID != "" {
		user, err := gen.User.Where(gen.User.Email.Eq(login.Username), gen.User.OrganizationID.Eq(orgID)).First()
		if err != nil {
			return nil, errors.NewError(errors.CodeNotFound, errors.MsgNotFound, "user", "email", login.Username)
		}
		if user.Status != "Active" {
			return nil, errors.NewError(errors.CodeForbidden, errors.MsgForbidden)
		}
		if !security.VerifyPasswordHash(login.Password, user.Password) {
			return nil, errors.NewError(errors.CodePasswordIncorrect, errors.MsgPasswordIncorrect)
		}
		return security.GenerateTokenResponse(user.ID, user.Username), nil
	}
	users, err := gen.User.Where(gen.User.Email.Eq(login.Username)).Find()
	if err != nil {
		return nil, errors.NewError(errors.CodeNotFound, errors.MsgNotFound, "user", "email", login.Username)
	}
	if len(users) > 1 {
		return nil, errors.NewError(errors.CodeNotFound, errors.MsgNotFound, "user", "email", login.Username)
	}
	user := users[0]
	if user.Status == "Active" && security.VerifyPasswordHash(login.Password, user.Password) {
		return security.GenerateTokenResponse(users[0].ID, users[0].Username), nil
	}
	if user.Status != "Active" {
		return nil, errors.NewError(errors.CodeForbidden, errors.MsgForbidden)
	}
	return nil, errors.NewError(errors.CodePasswordIncorrect, errors.MsgPasswordIncorrect)
}
