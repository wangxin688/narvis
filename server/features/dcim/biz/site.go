package biz

import (
	"github.com/wangxin688/narvis/intend/devicerole"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/dcim/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"
)

type SiteService struct{}

func NewSiteService() *SiteService {
	return &SiteService{}
}

func (s *SiteService) Create(site schemas.SiteCreate) (string, error) {
	newSite := &models.Site{
		Name:        site.Name,
		SiteCode:    site.SiteCode,
		Status:      site.Status,
		Region:      site.Region,
		TimeZone:    site.TimeZone,
		Latitude:    site.Latitude,
		Longitude:   site.Longitude,
		Address:     site.Address,
		Description: site.Description,
	}
	err := gen.Site.Create(newSite)
	if err != nil {
		return "", err
	}
	return newSite.Id, nil
}

func (s *SiteService) Update(Id string, site schemas.SiteUpdate) error {
	updateFields := make(map[string]any)
	if site.Name != nil {
		updateFields["name"] = site.Name
	}
	if site.SiteCode != nil {
		updateFields["siteCode"] = site.SiteCode
	}
	if site.Region != nil {
		updateFields["region"] = site.Region
	}
	if site.TimeZone != nil {
		updateFields["timeZone"] = site.TimeZone
	}
	if site.Latitude != nil {
		updateFields["latitude"] = site.Latitude
	}
	if site.Longitude != nil {
		updateFields["longitude"] = site.Longitude
	}
	if site.Address != nil {
		updateFields["address"] = site.Address
	}
	if site.Description != nil {
		updateFields["description"] = site.Description
	}
	_, err := gen.Site.Select(gen.Site.Id.Eq(Id), gen.Site.OrganizationId.Eq(global.OrganizationId.Get())).Updates(updateFields)
	if err != nil {
		return err
	}
	return nil
}

func (s *SiteService) Delete(Id string) error {
	_, err := gen.Site.Select(gen.Site.Id.Eq(Id), gen.Site.OrganizationId.Eq(global.OrganizationId.Get())).Delete()
	if err != nil {
		return err
	}
	return nil
}

func (s *SiteService) GetById(Id string) (schemas.Site, error) {

	site, err := gen.Site.Select(gen.Site.Id.Eq(Id), gen.Site.OrganizationId.Eq(global.OrganizationId.Get())).First()
	if err != nil {
		return schemas.Site{}, err
	}
	return schemas.Site{
		Id:          site.Id,
		CreatedAt:   site.CreatedAt,
		UpdatedAt:   site.UpdatedAt,
		Name:        site.Name,
		SiteCode:    site.SiteCode,
		Status:      site.Status,
		Region:      site.Region,
		TimeZone:    site.TimeZone,
		Latitude:    site.Latitude,
		Longitude:   site.Longitude,
		Address:     site.Address,
		Description: site.Description,
	}, nil
}

func (s *SiteService) GetSwitchCount(siteId string) (int64, error) {
	swDeviceRoles := devicerole.GetSwitchDeviceRole()
	switchDeviceRoles := make([]string, 0)
	for _, swDeviceRole := range *swDeviceRoles {
		switchDeviceRoles = append(switchDeviceRoles, string(swDeviceRole))
	}
	return gen.Device.Select(gen.Device.DeviceRole.In(switchDeviceRoles...), gen.Device.SiteId.Eq(siteId)).Count()
}
