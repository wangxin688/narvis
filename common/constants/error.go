package constants

type ErrorCode int
type ErrorMsg string

const ErrorOk ErrorCode = 0
const ErrorOkMsg ErrorMsg = "success"

// Generic error codes
const (
	ErrorNotFound                     ErrorCode = 404
	ErrorExist                        ErrorCode = 409
	ErrorInternalServerError          ErrorCode = 500
	ErrorAccessTokenExpired           ErrorCode = 4011
	ErrorRefreshTokenExpired          ErrorCode = 4012
	ErrorAccessTokenInvalid           ErrorCode = 4013
	ErrorAccessTokenInvalidForRefresh ErrorCode = 4014
	ErrorTokenMissing                 ErrorCode = 4015
	ErrorUnprocessableEntity          ErrorCode = 422
	ErrorForbidden                    ErrorCode = 403
	ErrorUnauthorized                 ErrorCode = 4010
	ErrorBadRequest                   ErrorCode = 400
	ErrorTooManyRequests              ErrorCode = 429
)

const (
	ErrorNotFoundMsg                     ErrorMsg = "not found %s with %s = %s"
	ErrorExistMsg                        ErrorMsg = "already exist %s with %s = %s"
	ErrorInternalServerErrorMsg          ErrorMsg = "internal server error, request id: %s"
	ErrorAccessTokenExpiredMsg           ErrorMsg = "access token expired"
	ErrorRefreshTokenExpiredMsg          ErrorMsg = "refresh token expired"
	ErrorAccessTokenInvalidMsg           ErrorMsg = "access token invalid"
	ErrorAccessTokenInvalidForRefreshMsg ErrorMsg = "access token invalid for refresh"
	ErrorTokenMissingMsg                 ErrorMsg = "access token no provided"
	ErrorUnprocessableEntityMsg          ErrorMsg = "validation error, unprocessable entity"
	ErrorForbiddenMsg                    ErrorMsg = "forbidden, permission denied"
	ErrorUnauthorizedMsg                 ErrorMsg = "unauthorized access, no privilege to access the resource"
	ErrorBadRequestMsg                   ErrorMsg = "bad request, check the request parameters or request body"
	ErrorTooManyRequestsMsg              ErrorMsg = "too many requests"
)

// Business error codes

// Organization error codes (100000 - 200000)

// Admin error codes (200000 - 300000)

// organization error codes (300000 - 400000)

// intend error codes (400000 - 500000)

// dcim error codes (500000 - 600000)

// ipam error codes (600000 - 700000)

// circuit error codes (700000 - 800000)

// monitor error codes (800000 - 900000)

// alert error codes (900000 - 1000000)
