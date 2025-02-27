package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/intend/helpers/security"
	"github.com/wangxin688/narvis/server/config"
	"github.com/wangxin688/narvis/server/features/organization/biz"
	"github.com/wangxin688/narvis/server/tools/errors"
)

func ProxyAuthMiddleware() gin.HandlerFunc {
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
		proxyId, secretKey, err := security.VerifyAgentToken(parts[1], config.Settings.Jwt.PublicAuthKey)
		if err != nil || secretKey != config.Settings.Jwt.PublicAuthKey {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, errors.GenericError{
					Code:    errors.CodeAccessTokenInvalid,
					Message: errors.MsgAccessTokenInvalid,
					Data:    nil,
				},
			)
			return
		}
		if !biz.NewProxyService().VerifyProxy(proxyId) {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, errors.GenericError{
					Code:    errors.CodeAccessTokenInvalid,
					Message: errors.MsgAccessTokenInvalid,
					Data:    nil,
				},
			)
			return
		}
		c.Next()
	}
}

func PublicAuthMiddleware() gin.HandlerFunc {
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
		if parts[1] != config.Settings.Jwt.PublicAuthKey {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, errors.GenericError{
					Code:    errors.CodeAccessTokenInvalid,
					Message: errors.MsgAccessTokenInvalid,
					Data:    nil,
				},
			)
			return
		}
		c.Next()
	}
}
