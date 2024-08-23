package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	consts "github.com/wangxin688/narvis/common/constants"
	"github.com/wangxin688/narvis/server/core/security"
	"github.com/wangxin688/narvis/server/global"
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
						"code":    consts.ErrorTokenMissing,
						"message": consts.ErrorTokenMissingMsg,
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
						"code":    consts.ErrorAccessTokenInvalid,
						"message": consts.ErrorAccessTokenInvalidMsg,
						"detail":  nil,
					},
				},
			)
			return
		}

		tokenErrCode, tokenClaims := security.VerifyAccessToken(parts[1])
		switch tokenErrCode {
		case consts.ErrorAccessTokenInvalid:
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, gin.H{
					"error": gin.H{
						"code":    consts.ErrorAccessTokenInvalid,
						"message": consts.ErrorAccessTokenInvalidMsg,
						"detail":  nil,
					},
				},
			)
			return
		case consts.ErrorAccessTokenInvalidForRefresh:
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, gin.H{
					"error": gin.H{
						"code":    consts.ErrorAccessTokenExpired,
						"message": consts.ErrorAccessTokenExpiredMsg,
						"detail":  nil,
					},
				},
			)
			return
		case consts.ErrorAccessTokenExpired:
			c.AbortWithStatusJSON(
				http.StatusUnauthorized, gin.H{
					"error": gin.H{
						"code":    consts.ErrorAccessTokenExpired,
						"message": consts.ErrorAccessTokenExpiredMsg,
						"detail":  nil,
					},
				},
			)
			return
		case consts.ErrorOk:
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
