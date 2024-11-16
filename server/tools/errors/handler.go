package errors

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/global"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var pgConflictRegexp = regexp.MustCompile(`\((.*?)\)=\((.*?)\)`)
var pgViolateRegexp = regexp.MustCompile(`\(([^)]+)\)=\(([^)]+)\).*?\"([^"]+)\"`)

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

// ResponseErrorHandler is a gin middleware to handle errors.
// The error is always one of the following:
//   - GenericError
//   - validator.ValidationErrors
//   - pgconn.PgError
//   - gorm.ErrRecordNotFound
//   - other errors
//
// The middleware will log the error and return the error to the client
// with the correct status code.
func ResponseErrorHandler(g *gin.Context, e error) {
	var generalError *GenericError
	var validationError validator.ValidationErrors
	var pgError *pgconn.PgError
	switch {
	case errors.As(e, &generalError):
		if generalError == nil {
			core.Logger.Warn("[errResponseHandler]: unknown error", zap.Error(e), zap.String("X-Request-ID", global.XRequestId.Get()))
			g.AbortWithStatusJSON(http.StatusInternalServerError, NewError(CodeInternalServerError, MsgInternalServerError, global.XRequestId.Get()))
			return
		}
		core.Logger.Warn("[errResponseHandler]: general error", zap.Error(e), zap.Int("code", int(generalError.Code)))
		if generalError.Code <= 500 {
			g.AbortWithStatusJSON(int(generalError.Code), generalError)
			return
		}
		g.AbortWithStatusJSON(http.StatusBadRequest, generalError)
		return
	case errors.As(e, &pgError):
		if pgError == nil {
			core.Logger.Error("[errResponseHandler]: unknown error", zap.Error(e), zap.String("X-Request-ID", global.XRequestId.Get()))
			g.AbortWithStatusJSON(http.StatusInternalServerError, NewError(CodeInternalServerError, MsgInternalServerError, global.XRequestId.Get()))
			return
		}
		core.Logger.Warn("[errResponseHandler]: conflict error", zap.Error(e), zap.String("tableName", pgError.TableName))
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
		if pgError.Code == "23503" {
			var key, value, table string
			matches := pgViolateRegexp.FindStringSubmatch(pgError.Detail)
			if len(matches) >= 3 {
				key = matches[1]
				value = matches[2]
				table = matches[3]
				if strings.Contains(pgError.Message, "insert or update") {
					g.AbortWithStatusJSON(http.StatusNotFound, NewError(CodeNotFound, MsgNotFound, table, key, value))
					return
				}
				if strings.Contains(pgError.Message, "update or delete") {
					g.AbortWithStatusJSON(http.StatusForbidden, NewError(CodeDeleteRestriction, MsgDeleteRestriction, table, table))
					return
				}
				g.AbortWithStatusJSON(http.StatusUnprocessableEntity, NewError(CodeUnprocessableEntity, MsgUnprocessableEntity, pgError.Detail))
				return
			}
			g.AbortWithStatusJSON(http.StatusUnprocessableEntity, NewError(CodeUnprocessableEntity, MsgUnprocessableEntity, pgError.Detail))
			return
		}
		g.AbortWithStatusJSON(http.StatusInternalServerError, NewError(CodeInternalServerError, MsgInternalServerError, global.XRequestId.Get()))
		return
	case errors.As(e, &validationError):
		if validationError == nil {
			core.Logger.Error("[errResponseHandler]: unknown error", zap.Error(e), zap.String("X-Request-ID", global.XRequestId.Get()))
			g.AbortWithStatusJSON(http.StatusInternalServerError, NewError(CodeInternalServerError, MsgInternalServerError, global.XRequestId.Get()))
			return
		}
		core.Logger.Warn("[errResponseHandler]: validation error", zap.Error(e))
		g.AbortWithStatusJSON(http.StatusUnprocessableEntity, NewErrorWithData(CodeUnprocessableEntity, MsgUnprocessableEntity, e.Error()))
		return
	case errors.Is(e, gorm.ErrRecordNotFound):
		g.AbortWithStatusJSON(http.StatusNotFound, NewError(CodeNotFound, "record not found"))
		return
	default:
		core.Logger.Error("[errResponseHandler]: unknown error", zap.Error(e), zap.String("X-Request-ID", global.XRequestId.Get()))
		g.AbortWithStatusJSON(http.StatusInternalServerError, NewError(CodeInternalServerError, MsgInternalServerError, global.XRequestId.Get()))
		return
	}

}
