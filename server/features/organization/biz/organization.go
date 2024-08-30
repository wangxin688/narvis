package biz

import (
	"fmt"

	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/admin/biz"
	"github.com/wangxin688/narvis/server/features/organization/schemas"
	"github.com/wangxin688/narvis/server/global/constants"
	"github.com/wangxin688/narvis/server/models"
	e "github.com/wangxin688/narvis/server/tools/errors"
	"go.uber.org/zap"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type OrganizationService struct {
	gen.IOrganizationDo
}

func NewOrganizationService() *OrganizationService {
	return &OrganizationService{}
}

func (o *OrganizationService) CreateOrganization(organization *schemas.OrganizationCreate) (*schemas.Organization, error) {

	if organization.AuthType == uint8(constants.SlackTenantAuthType) ||
		organization.AuthType == uint8(constants.TeamsTenantAuthType) ||
		organization.AuthType == uint8(constants.GooglTenantAuthType) {
		if organization.AuthConfig.ClientId == "" || organization.AuthConfig.ClientSecret == "" {
			return nil, &e.GenericError{Code: e.CodeInvalidAuthConfig, Message: e.MsgInvalidAuthConfig}
		}
	}
	organizationModel := &models.Organization{
		Name:           organization.Name,
		Active:         organization.Active,
		EnterpriseCode: organization.EnterpriseCode,
		DomainName:     organization.DomainName,
		LicenseCount:   organization.LicenseCount,
		AuthType:       organization.AuthType,
	}
	if organization.AuthType != uint8(constants.LocalTenantAuthType) {
		clientId := organization.AuthConfig.ClientId
		clientSecret := organization.AuthConfig.ClientSecret
		authConfig := datatypes.NewJSONType(
			models.AuthConfig{ClientId: clientId, ClientSecret: clientSecret},
		)
		organizationModel.AuthConfig = &authConfig
	}
	if o.validateExist(organization) {
		core.Logger.Error(fmt.Sprintf("failed to create organization %v, enterprise_code or domain_name already exist", organization))
		return nil, &e.GenericError{Code: e.CodeOrganizationAlreadyExist, Message: e.MsgOrganizationAlreadyExist}
	}

	gen.Organization.UnderlyingDB().Transaction(func(tx *gorm.DB) error {
		err := gen.Organization.Create(organizationModel)
		if err != nil {
			core.Logger.Error(fmt.Sprintf("failed to create organization %v", organization), zap.Error(err))
			return err
		}
		core.Logger.Info(fmt.Sprintf("create organization %s %s", organization.Name, organizationModel.Id))
		userService := biz.NewUserService()
		user, err := userService.CreateAdminUser(organization.EnterpriseCode, organizationModel.Id, organization.AdminPassword)
		if err != nil {
			core.Logger.Error(fmt.Sprintf("failed to create admin user %v", organization), zap.Error(err))
			return err
		}
		core.Logger.Info(fmt.Sprintf("create admin user %s", user.Id))
		return nil
	})
	return &schemas.Organization{
		Id:             organizationModel.Id,
		CreatedAt:      organizationModel.CreatedAt,
		UpdatedAt:      organizationModel.UpdatedAt,
		Name:           organization.Name,
		Active:         organization.Active,
		EnterpriseCode: organization.EnterpriseCode,
		DomainName:     organization.DomainName,
		LicenseCount:   organization.LicenseCount,
		AuthType:       organization.AuthType,
	}, nil
}

func (o *OrganizationService) GetByName(enterpriseCode string) (*models.Organization, error) {

	organization, err := gen.Organization.Where(gen.Organization.EnterpriseCode.Eq(enterpriseCode)).First()
	if err != nil {
		return nil, err
	}
	return organization, nil
}

func (o *OrganizationService) GetById(orgId string) (*models.Organization, error) {

	organization, err := gen.Organization.Where(gen.Organization.Id.Eq(orgId)).First()
	if err != nil {
		return nil, err
	}
	return organization, nil
}

func (o *OrganizationService) GetByDomainName(domainName string) (*models.Organization, error) {

	organization, err := gen.Organization.Where(gen.Organization.DomainName.Eq(domainName)).First()
	if err != nil {
		return nil, err
	}
	return organization, nil
}

// func (o *OrganizationService) UpdateOrganization(ordId string, organization *schemas.OrganizationUpdate) (*schemas.OrganizationUpdate, *errors.AppError) {

// }

func (o *OrganizationService) DeleteOrganization(orgId string) error {
	_, err := o.IOrganizationDo.Where(gen.Organization.Id.Eq(orgId)).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (o *OrganizationService) ListOrganizations(organization *schemas.OrganizationQuery) (*schemas.OrganizationUpdate, error) {

	return nil, nil
}

func (o *OrganizationService) validateExist(organization *schemas.OrganizationCreate) bool {

	org, err := o.GetByName(organization.EnterpriseCode)
	if err != nil {
		return false
	}
	org1, err := o.GetByDomainName(organization.DomainName)
	if err != nil {
		return false
	}
	return org != nil && org1 != nil

}
