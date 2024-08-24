package constants

type TenantAuthTypeEnum uint8

const (
	LocalTenantAuthType  TenantAuthTypeEnum = 0
	SlackTenantAuthType  TenantAuthTypeEnum = 1
	TeamsTenantAuthType  TenantAuthTypeEnum = 2
	GooglTenantAuthType  TenantAuthTypeEnum = 3
	GithubTenantAuthType TenantAuthTypeEnum = 4
)
