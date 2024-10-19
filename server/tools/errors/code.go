package errors

type ErrorCode int
type ErrorMsg string

const ErrorOk ErrorCode = 0
const ErrorOkMsg ErrorMsg = "success"

// Generic error codes
const (
	CodeNotFound                     ErrorCode = 404
	CodeExist                        ErrorCode = 409
	CodeInternalServerError          ErrorCode = 500
	CodeAccessTokenExpired           ErrorCode = 4011
	CodeRefreshTokenExpired          ErrorCode = 4012
	CodeAccessTokenInvalid           ErrorCode = 4013
	CodeAccessTokenInvalidForRefresh ErrorCode = 4014
	CodeTokenMissing                 ErrorCode = 4015
	CodeUnprocessableEntity          ErrorCode = 422
	CodeForbidden                    ErrorCode = 403
	CodeUnauthorized                 ErrorCode = 4010
	CodeBadRequest                   ErrorCode = 400
	CodeTooManyRequests              ErrorCode = 429
	CodeDeleteRestriction            ErrorCode = 10000
)

const (
	MsgNotFound                     ErrorMsg = "%s with %s = %s not found "
	MsgExist                        ErrorMsg = "%s with %s = %s already exist "
	MsgInternalServerError          ErrorMsg = "internal server error, request id: %s"
	MsgAccessTokenExpired           ErrorMsg = "access token expired"
	MsgRefreshTokenExpired          ErrorMsg = "refresh token expired"
	MsgAccessTokenInvalid           ErrorMsg = "access token invalid"
	MsgAccessTokenInvalidForRefresh ErrorMsg = "access token invalid for refresh"
	MsgTokenMissing                 ErrorMsg = "access token not provide"
	MsgUnprocessableEntity          ErrorMsg = "validation error, unprocessable entity"
	MsgForbidden                    ErrorMsg = "forbidden, permission denied"
	MsgUnauthorized                 ErrorMsg = "unauthorized access, no privilege to access the resource"
	MsgBadRequest                   ErrorMsg = "bad request, check the request parameters or request body"
	MsgTooManyRequests              ErrorMsg = "too many requests"
	MsgDeleteRestriction            ErrorMsg = "delete restriction, %s is still associated, please delete %s belongs to it first"
)

// Business error codes

// Organization error codes (100000 - 200000)

const (
	CodeInvalidAuthConfig ErrorCode = 100000
	MsgInvalidAuthConfig  ErrorMsg  = "invalid auth config"
	CodeCreateTenantFail  ErrorCode = 100001
	MsgCreateTenantFail   ErrorMsg  = "create tenant fail %s"
)

// Admin error codes (200000 - 300000)

const (
	CodeInvalidGroupNameForReserve ErrorCode = 200000
	MsgInvalidGroupNameForReserve  ErrorMsg  = "invalid group name for system admin reserve"
	CodePasswordIncorrect          ErrorCode = 200001
	MsgPasswordIncorrect           ErrorMsg  = "password incorrect"
)

// organization error codes (300000 - 400000)

const (
	CodeOrganizationAlreadyExist ErrorCode = 300000
	MsgOrganizationAlreadyExist  ErrorMsg  = "organization already exist"
	CodeLicenseCountExceeded     ErrorCode = 300001
	MsgLicenseCountExceeded      ErrorMsg  = "license count exceeded, can not add infra anymore"
)

// intend error codes (400000 - 500000)

// infra error codes (500000 - 600000)

const (
	CodeCredentialDeviceIdMissing   ErrorCode = 500000
	MsgCredentialDeviceIdMissing    ErrorMsg  = "credential device Id missing, global credential is already created"
	CodeGlobalCredentialMoreThanOne ErrorCode = 500001
	MsgGlobalCredentialMoreThanOne  ErrorMsg  = "global CLI credential already created more than one"
	CodeUpdateRackFailed            ErrorCode = 500002
	MsgUpdateRackFailed             ErrorMsg  = "update rack failed, uHeight should be greater than the devices been occupied"
	CodeRackPositionInconsecutive   ErrorCode = 500003
	MsgRackPositionInconsecutive    ErrorMsg  = "rack position is not consecutive"
	CodeGlobalCredentialDeleteDeny  ErrorCode = 500004
	MsgGlobalCredentialDeleteDeny   ErrorMsg  = "Organization-level credential is not allowed to be deleted"
	CodeIpRangeNotProvided          ErrorCode = 500005
	MsgIpRangeNotProvided           ErrorMsg  = "ip range not provided"
	CodeTaskNameInvalid             ErrorCode = 500006
	MsgTaskNameInvalid              ErrorMsg  = "task name invalid"
	CodeRackPositionOccupied        ErrorCode = 500007
	MsgRackPositionOccupied         ErrorMsg  = "rack position occupied"
	CodeNoDevicesFound              ErrorCode = 500008
	MsgNoDevicesFound              ErrorMsg  = "generate task failed, no devices found for site"
)

// ipam error codes (600000 - 700000)

// circuit error codes (700000 - 800000)

// monitor error codes (800000 - 900000)
const (
	CodeMetricNotDefined ErrorCode = 800000
	MsgMetricNotDefined  ErrorMsg  = "metric: %s not defined"
	CodeQueryBuildFailed ErrorCode = 800001
	MsgQueryBuildFailed  ErrorMsg  = "query %s build failed"
)

// alert error codes (900000 - 1000000)
const (
	CodeAlertStartTimeInFuture          ErrorCode = 900000
	MsgAlertStartTimeInFuture           ErrorMsg  = "alert start time should not be in the future"
	CodeAlertNameNotDefined             ErrorCode = 900001
	MsgAlertNameNotDefined              ErrorMsg  = "alert name: %s not defined"
	CodeAlertHostIdInvalid              ErrorCode = 900002
	MsgAlertHostIdInvalid               ErrorMsg  = "alert hostId: %s invalid"
	CodeApNameTagMissing                ErrorCode = 900003
	MsgApNameTagMissing                 ErrorMsg  = "alert apName tag missing in metrics system"
	CodeInterfaceTagMissing             ErrorCode = 900004
	MsgInterfaceTagMissing              ErrorMsg  = "alert interface tag missing in metrics system"
	CodeAlertGroupMissingOrganizationId ErrorCode = 900005
	MsgAlertGroupMissingOrganizationId  ErrorMsg  = "alert group missing organization id"
)

// webssh error codes (1000000 - 1100000)
const (
	CodeSessionIdEmpty    ErrorCode = 1000000
	MsgSessionIdEmpty     ErrorMsg  = "sessionId empty"
	CodeSessionIdNotFound ErrorCode = 1000001
	MsgSessionIdNotFound  ErrorMsg  = "sessionId not found"
	CodeSessionTimeout    ErrorCode = 1000002
	MsgSessionTimeout     ErrorMsg  = "timeout waiting for proxy websocket connection"
	CodeWebSocketInitFail ErrorCode = 1000003
	MsgWebSocketInitFail  ErrorMsg  = "failed to initialize websocket connection , error: %s"
)
