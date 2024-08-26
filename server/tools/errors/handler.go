package errors

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/global"
	"go.uber.org/zap"
)

var pgConflictRegexp = regexp.MustCompile(`\((.*?)\)=\((.*?)\)`)

type GenericError struct {
	Code    ErrorCode
	Data    any
	Message ErrorMsg
}

func (e *GenericError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func formatErrorMsg(message ErrorMsg, args ...any) string {
	return fmt.Sprintf(string(message), args...)
}

func NewError(code ErrorCode, message ErrorMsg, args ...any) *GenericError {
	return &GenericError{
		Code:    code,
		Message: ErrorMsg(formatErrorMsg(message, args...)),
	}
}

func NewErrorWithData(code ErrorCode, message ErrorMsg, data any, args ...any) *GenericError {
	return &GenericError{
		Code:    code,
		Data:    data,
		Message: ErrorMsg(formatErrorMsg(message, args...)),
	}
}

func ResponseErrorHandler(g *gin.Context, e error) {
	var generalError *GenericError
	var validationError validator.ValidationErrors
	var pgError *pgconn.PgError
	switch {
	case errors.As(e, &generalError):
		if generalError == nil {
			core.Logger.Error("unknown error", zap.Error(e))
			g.AbortWithStatusJSON(http.StatusInternalServerError, NewError(CodeInternalServerError, MsgInternalServerError, global.XRequestID.Get()))
			return
		}
		if generalError.Code <= 500 {
			g.AbortWithStatusJSON(int(generalError.Code), generalError)
			return
		}
		g.AbortWithStatusJSON(http.StatusBadRequest, generalError)
		return
	case errors.As(e, &pgError):
		if pgError == nil {
			core.Logger.Error("unknown error", zap.Error(e))
			g.AbortWithStatusJSON(http.StatusInternalServerError, NewError(CodeInternalServerError, MsgInternalServerError, global.XRequestID.Get()))
			return
		}
		if pgError.Code == "23505" {
			var fields, values string
			matches := pgConflictRegexp.FindStringSubmatch(pgError.Detail)
			if len(matches) >= 3 {
				fields = matches[1]
				values = matches[2]
				fields, values = removeOrgInError(fields, values)
			}

			g.AbortWithStatusJSON(http.StatusConflict, NewError(CodeExist, MsgExist, pgError.TableName, fields, values))
			return
		}
		g.AbortWithStatusJSON(http.StatusInternalServerError, NewError(CodeInternalServerError, MsgInternalServerError, global.XRequestID.Get()))
		return
	case errors.As(e, &validationError):
		if validationError == nil {
			core.Logger.Error("unknown error", zap.Error(e))
			g.AbortWithStatusJSON(http.StatusInternalServerError, NewError(CodeInternalServerError, MsgInternalServerError, global.XRequestID.Get()))
			return
		}
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, NewErrorWithData(CodeUnprocessableEntity, MsgUnprocessableEntity, e.Error()))
		return
	default:
		core.Logger.Error("unknown error", zap.Error(e))
		g.AbortWithStatusJSON(http.StatusInternalServerError, NewError(CodeInternalServerError, MsgInternalServerError, global.XRequestID.Get()))
		return
	}

}
