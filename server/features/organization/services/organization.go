package services

import (
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/organization/schemas"
)

type OrganizationService struct {
	gen.IOrganizationDo
}

func NewOrganizationService() *OrganizationService {
	return &OrganizationService{}
}

func (o *OrganizationService) validateLocalAuth(authConfig *schemas.AuthConfig) bool {
	return authConfig.Password != nil
}

func (o *OrganizationService) validateOauth2Auth(authConfig *schemas.SlackAuthConfig) bool {
	return authConfig.ClientId != "" && authConfig.ClientSecret != ""
}

// func (o *OrganizationService) CreateOrganization(organization *schemas.OrganizationCreate) (error, *schemas.Organization) {

// }
