package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/core/security"
	"github.com/wangxin688/narvis/server/features/admin/biz"
	"github.com/wangxin688/narvis/server/global"
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

		parts := strings.Split(tokenString, "")
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
			global.UserID.Set(tokenClaims.UserID)
			return
		}
		permissionCheck(tokenClaims.UserID, c.FullPath())
		c.Next()
	}
}

func permissionCheck(userID string, path string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if !checkUserPermission(userID, path) {
			c.AbortWithStatusJSON(
				http.StatusForbidden, errors.GenericError{
					Code:    http.StatusForbidden,
					Message: "permission denied",
					Data:    nil,
				},
			)
			return
		}
		c.Next()
	}
}

func checkUserPermission(userID string, path string) bool {
	return biz.CheckRolePathPermission(userID, path)
}
