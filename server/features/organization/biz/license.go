package biz

import "github.com/wangxin688/narvis/server/dal/gen"

func GetOrgLicense(orgId string) (uint32, error) {
	org, err := gen.Organization.Select(gen.Organization.LicenseCount).Where(gen.Organization.Id.Eq(orgId)).First()
	if err != nil {
		return 0, err
	}
	return org.LicenseCount, nil
}
