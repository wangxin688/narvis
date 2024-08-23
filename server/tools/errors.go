package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"

	consts "github.com/wangxin688/narvis/common/constants"
)

type GenericError struct {
	ErrorCode consts.ErrorCode
	ErrorMsg  string
}

func ErrorMessage(message string, args ...interface{}) string {
	if len(args) == 0 {
		return message
	}
	return fmt.Sprintf(message, args...)
}

func NewError(code consts.ErrorCode, msg string, args ...interface{}) *GenericError {

	return &GenericError{
		ErrorCode: code,
		ErrorMsg:  ErrorMessage(msg, args...),
	}
}

// Create a new generic not found error
func NewNotFoundError(resource, field, value string) *GenericError {
	return NewError(consts.ErrorNotFound, string(consts.ErrorNotFoundMsg), resource, field, value)
}

// Create a new generic exist error
func NewExistError(resource, field, value string) *GenericError {
	return NewError(consts.ErrorExist, string(consts.ErrorExistMsg), resource, field, value)
}

// Create a new generic internal server error
func NewInternalServerError(g *gin.Context) *GenericError {
	requestId := g.Value(consts.XRequestID)
	return NewError(consts.ErrorInternalServerError, string(consts.ErrorInternalServerErrorMsg), requestId)
}
