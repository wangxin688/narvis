package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/core/security"
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
				http.StatusUnauthorized, gin.H{
					"error": gin.H{
						"code":    errors.CodeTokenMissing,
						"message": errors.MsgTokenMissing,
						"detail":  nil,
					},
				},
			)
			return
		}

		parts := strings.Split(tokenString, "")
		if len(parts) != 2 || parts[0] != AuthorizationBearer {
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, gin.H{
					"error": gin.H{
						"code":    errors.CodeAccessTokenInvalid,
						"message": errors.MsgAccessTokenInvalid,
						"detail":  nil,
					},
				},
			)
			return
		}

		tokenErrCode, tokenClaims := security.VerifyAccessToken(parts[1])
		switch tokenErrCode {
		case errors.CodeAccessTokenInvalid:
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, gin.H{
					"error": gin.H{
						"code":    errors.CodeAccessTokenInvalid,
						"message": errors.MsgAccessTokenInvalid,
						"detail":  nil,
					},
				},
			)
			return
		case errors.CodeAccessTokenExpired:
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, gin.H{
					"error": gin.H{
						"code":    errors.CodeAccessTokenExpired,
						"message": errors.MsgAccessTokenExpired,
						"detail":  nil,
					},
				},
			)
			return
		case errors.CodeAccessTokenInvalidForRefresh:
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, gin.H{
					"error": gin.H{
						"code":    errors.CodeAccessTokenInvalidForRefresh,
						"message": errors.MsgAccessTokenInvalidForRefresh,
						"detail":  nil,
					},
				},
			)
			return
		case errors.ErrorOk:
			global.UserID.Set(tokenClaims.UserID)
			return
		}

		c.Next()
	}
}

// func permissionCheck(userID string) gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		if !checkUserPermission(userID) {
// 			c.AbortWithStatusJSON(
// 				http.StatusUnauthorized, gin.H{
// 					"error": gin.H{
// 						"code":    consts.ErrorPermissionDenied,
// 						"message": consts.ErrorPermissionDeniedMsg,
// 						"detail":  nil,
// 					},
// 				},
// 			)
// 			return
// 		}
// 		c.Next()
// 	}
// }

// func checkUserPermission(userID string) bool {

// }
