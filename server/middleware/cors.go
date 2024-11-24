// Copyright 2024 wangxin.jeffry@gmail.com
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samber/lo"
	"github.com/wangxin688/narvis/server/config"
)

func corsNoStrict() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-ID")
		c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PUT,PATCH")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, New-Token, New-Expires-At")
		c.Header("Access-Control-Allow-Credentials", "true")

		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		c.Next()
	}
}

func CorsMiddleware() gin.HandlerFunc {
	if len(config.Settings.Cors.AllowedOrigins) == 1 && config.Settings.Cors.AllowedOrigins[0] == "*" {
		return corsNoStrict()
	}
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if checkCors(origin) {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization,Token,X-Token,X-User-ID")
			c.Header("Access-Control-Allow-Methods", "POST,GET,OPTIONS,DELETE,PUT,PATCH")
			c.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type,New-Token,New-Expires-At")
			c.Header("Access-Control-Allow-Credentials", "true")
		} else {
			c.AbortWithStatus(http.StatusForbidden)
		}
		c.Next()
	}
}

func checkCors(currentOrigin string) bool {
	return lo.Contains(config.Settings.Cors.AllowedOrigins, currentOrigin)
}
