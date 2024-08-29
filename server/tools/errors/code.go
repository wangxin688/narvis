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
)

// intend error codes (400000 - 500000)

// dcim error codes (500000 - 600000)

// ipam error codes (600000 - 700000)

// circuit error codes (700000 - 800000)

const (
	CodeCircuitZInterfaceMissing ErrorCode = 700000
	MsgCircuitZInterfaceMissing  ErrorMsg  = "z side interface ID missing when create circuit type is Intranet"
	CodeCircuitSameInterface     ErrorCode = 700001
	MsgCircuitSameInterface      ErrorMsg  = "same interface ID in a side and z side is not allowed"
	CodeCircuitAInterfaceMissing ErrorCode = 700002
	MsgCircuitAInterfaceMissing  ErrorMsg  = "a side interface ID missing when create circuit type is Intranet"
	CodeCircuitSameDevice        ErrorCode = 700003
	MsgCircuitSameDevice         ErrorMsg  = "same device ID in a side and z side is not allowed"
	CodeCircuitZInterfaceNotAllow ErrorCode = 700004
	MsgCircuitZInterfaceNotAllow ErrorMsg  = "z side interface ID not allowed for internet circuit type"
)

// monitor error codes (800000 - 900000)

// alert error codes (900000 - 1000000)
