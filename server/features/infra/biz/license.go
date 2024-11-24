package infra_biz

import (
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/organization/biz"
	"github.com/wangxin688/narvis/server/pkg/contextvar"
)

func licenseUsage(orgId string) (uint32, error) {
	deviceCount, err := gen.Device.Where(
		gen.Device.OrganizationId.Eq(orgId)).Count()

	if err != nil {
		return 0, err
	}
	apCount, err := gen.AP.Where(
		gen.AP.OrganizationId.Eq(contextvar.OrganizationId.Get())).Count()
	if err != nil {
		return 0, err
	}
	serverCount, err := gen.Server.Where(
		gen.Server.OrganizationId.Eq(contextvar.OrganizationId.Get())).Count()
	if err != nil {
		return 0, err
	}
	return uint32(deviceCount) + uint32(apCount) + uint32(serverCount), nil
}

// LicenseUsageDepends checks if the new device can be added to the organization
// without exceeding the license count.
//
// It first gets the organization's license count, then gets the current used
// license count, and finally checks if the new device can be added without
// exceeding the license count.
func LicenseUsageDepends(newDevice uint32, orgId string) (bool, error) {
	// get the organization's license count
	orgLicenseCount, err := biz.GetOrgLicense(orgId)
	if err != nil {
		return false, err
	}

	// get the current used license count
	usedLicenseCount, err := licenseUsage(orgId)
	if err != nil {
		return false, err
	}

	// check if the new device can be added without exceeding the license count
	if usedLicenseCount+newDevice > orgLicenseCount {
		return false, nil
	}

	return true, nil
}
