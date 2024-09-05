package ipam_biz

import (
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/ipam/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"
)

type IpAddressService struct {
}

func NewIpAddressService() *IpAddressService {
	return &IpAddressService{}
}

func (i *IpAddressService) CreateIpAddress(ip *schemas.IpAddressCreate) (string, error) {
	newIpAddress := &models.IpAddress{
		OrganizationId: global.OrganizationId.Get(),
		Address:        ip.Address,
		Status:         ip.Status,
		MacAddress:     ip.MacAddress,
		Type:           ip.Type,
		Description:    ip.Description,
		SiteId:         ip.SiteId,
	}
	err := gen.IpAddress.Create(newIpAddress)
	if err != nil {
		return "", err
	}
	return newIpAddress.Id, nil
}

func (i *IpAddressService) GetById(ipId string) (*schemas.IpAddress, error) {
	ip, err := gen.IpAddress.Select().Where(gen.IpAddress.Id.Eq(ipId), gen.IpAddress.OrganizationId.Eq(global.OrganizationId.Get())).First()
	if err != nil {
		return nil, err
	}
	return &schemas.IpAddress{
		Id:          ip.Id,
		CreatedAt:   ip.CreatedAt,
		UpdatedAt:   ip.UpdatedAt,
		Address:     ip.Address,
		Status:      ip.Status,
		MacAddress:  ip.MacAddress,
		Type:        ip.Type,
		Description: ip.Description,
		SiteId:      ip.SiteId,
	}, nil
}

// func (i *IpAddressService) ListIpAddresses(query *schemas.IpAddressQuery) (int64, *[]*schemas.IpAddress, error) {
// 	res := make([]*schemas.IpAddress, 0)
// 	stmt := gen.IpAddress.Where(gen.IpAddress.OrganizationId.Eq(global.OrganizationId.Get()))
// 	if query.SiteId != nil {
// 		stmt = stmt.Where(gen.IpAddress.SiteId.Eq(*query.SiteId))
// 	}
// 	if query.Address != nil {
// 		stmt = stmt.Where(gen.IpAddress.Address.In(*query.Address...))
// 	}
// 	if query.Status != nil {
// 		stmt = stmt.Where(gen.IpAddress.Status.In(*query.Status...))
// 	}
// 	// if query.Type != nil {
// 	// 	stmt = stmt.Where(gen.IpAddress.Type.In(*query.Type...))
// 	// }
// 	if query.Range != nil {

// 	}
// 	total, err := stmt.Count()
// 	if err != nil {
// 		return 0, nil, err
// 	}
// 	stmt = stmt.Offset(query.Offset).Limit(query.Limit)
// 	err = stmt.Find(&res)
// 	if err != nil {
// 		return 0, nil, err
// 	}
// 	return total, &res, nil

// }

func (i *IpAddressService) UpdateIpAddress(ipId string, ip *schemas.IpAddressUpdate) error {
	var updateFields = make(map[string]any)
	if ip.Address != nil {
		updateFields["address"] = *ip.Address
	}
	if ip.Status != nil {
		updateFields["status"] = *ip.Status
	}
	if ip.MacAddress != nil {
		updateFields["macAddress"] = *ip.MacAddress
	}
	if ip.Type != nil {
		updateFields["type"] = *ip.Type
	}
	if ip.Description != nil {
		updateFields["description"] = *ip.Description
	}
	if ip.SiteId != nil && *ip.SiteId != "" {
		updateFields["siteId"] = *ip.SiteId
	}

	_, err := gen.IpAddress.Select(gen.IpAddress.Id.Eq(ipId), gen.IpAddress.OrganizationId.Eq(global.OrganizationId.Get())).Updates(updateFields)
	if err != nil {
		return err
	}
	return nil
}

func (i *IpAddressService) DeleteIpAddress(ipId string) error {
	_, err := gen.IpAddress.Select(gen.IpAddress.Id.Eq(ipId), gen.IpAddress.OrganizationId.Eq(global.OrganizationId.Get())).Delete()
	if err != nil {
		return err
	}
	return nil
}
