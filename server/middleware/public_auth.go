package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
		
		c.Next()
	}
}
