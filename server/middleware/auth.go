package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/features/admin/biz"
	"github.com/wangxin688/narvis/server/pkg/contextvar"
	"github.com/wangxin688/narvis/server/pkg/security"
	"github.com/wangxin688/narvis/server/tools/errors"
)

var AuthorizationString = "Authorization"
var AuthorizationBearer = "Bearer"

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader(AuthorizationString)

		if tokenString == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, errors.GenericError{
					Code:    errors.CodeAccessTokenInvalid,
					Message: errors.MsgAccessTokenInvalid,
					Data:    nil,
				},
			)
			return
		}

		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != AuthorizationBearer {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, errors.GenericError{
					Code:    errors.CodeAccessTokenInvalid,
					Message: errors.MsgAccessTokenInvalid,
					Data:    nil,
				},
			)
			return
		}

		tokenErrCode, tokenClaims := security.VerifyAccessToken(parts[1])
		switch tokenErrCode {
		case errors.CodeAccessTokenInvalid:
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, errors.GenericError{
					Code:    errors.CodeAccessTokenInvalid,
					Message: errors.MsgAccessTokenInvalid,
					Data:    nil,
				},
			)
			return
		case errors.CodeAccessTokenExpired:
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, errors.GenericError{
					Code:    errors.CodeAccessTokenExpired,
					Message: errors.MsgAccessTokenExpired,
					Data:    nil,
				},
			)
			return
		case errors.CodeAccessTokenInvalidForRefresh:
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, errors.GenericError{
					Code:    errors.CodeAccessTokenInvalidForRefresh,
					Message: errors.MsgAccessTokenInvalidForRefresh,
					Data:    nil,
				},
			)
			return
		case errors.ErrorOk:
			contextvar.UserId.Set(tokenClaims.UserId)
			if !checkUserPermission(tokenClaims.UserId, c.FullPath()) {
				c.AbortWithStatusJSON(
					http.StatusForbidden, errors.GenericError{
						Code:    http.StatusForbidden,
						Message: "permission denied",
						Data:    nil,
					},
				)
				return
			}
			return
		}
		c.Next()
	}
}

func checkUserPermission(userID string, path string) bool {
	user := biz.VerifyUser(userID)
	if user == nil {
		return false
	}
	contextvar.OrganizationId.Set(user.OrganizationId)
	return biz.CheckRolePathPermission(user, path)
}

// cookie auth middleware for websocket
func CookieAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token, err := c.Request.Cookie(AuthorizationString)
		if err != nil {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, errors.GenericError{
					Code:    errors.CodeAccessTokenInvalid,
					Message: errors.MsgAccessTokenInvalid,
					Data:    nil,
				},
			)
			return
		}
		tokenString := token.Value
		if tokenString == "" {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, errors.GenericError{
					Code:    errors.CodeAccessTokenInvalid,
					Message: errors.MsgAccessTokenInvalid,
					Data:    nil,
				},
			)
			return
		}

		parts := strings.Split(tokenString, " ")
		if len(parts) != 2 || parts[0] != AuthorizationBearer {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, errors.GenericError{
					Code:    errors.CodeAccessTokenInvalid,
					Message: errors.MsgAccessTokenInvalid,
					Data:    nil,
				},
			)
			return
		}

		tokenErrCode, tokenClaims := security.VerifyAccessToken(parts[1])
		switch tokenErrCode {
		case errors.CodeAccessTokenInvalid:
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, errors.GenericError{
					Code:    errors.CodeAccessTokenInvalid,
					Message: errors.MsgAccessTokenInvalid,
					Data:    nil,
				},
			)
			return
		case errors.CodeAccessTokenExpired:
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, errors.GenericError{
					Code:    errors.CodeAccessTokenExpired,
					Message: errors.MsgAccessTokenExpired,
					Data:    nil,
				},
			)
			return
		case errors.CodeAccessTokenInvalidForRefresh:
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, errors.GenericError{
					Code:    errors.CodeAccessTokenInvalidForRefresh,
					Message: errors.MsgAccessTokenInvalidForRefresh,
					Data:    nil,
				},
			)
			return
		case errors.ErrorOk:
			contextvar.UserId.Set(tokenClaims.UserId)
			if !checkUserPermission(tokenClaims.UserId, c.FullPath()) {
				c.AbortWithStatusJSON(
					http.StatusForbidden, errors.GenericError{
						Code:    http.StatusForbidden,
						Message: "permission denied",
						Data:    nil,
					},
				)
				return
			}
			return
		}
		c.Next()
	}

}
