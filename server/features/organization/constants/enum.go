package constants

type TenantAuthTypeEnum string

const (
	LocalTenantAuthType TenantAuthTypeEnum = "local"
	SlackTenantAuthType TenantAuthTypeEnum = "slack"
	TeamsTenantAuthType TenantAuthTypeEnum = "teams"
	GooglTenantAuthType TenantAuthTypeEnum = "google"
)


