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

// GinRecovery recovers from panics in the gin framework.

package middleware

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/pkg/contextvar"
	"go.uber.org/zap"
)

func GinRecovery(logger *zap.Logger, logStack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenConn bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						brokenConn = strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer")
					}
				}

				dump, _ := httputil.DumpRequest(c.Request, false)
				if brokenConn {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(dump)),
					)
					c.Error(err.(error)) //nolint: errcheck
					c.Abort()
					return
				}
				stack := ""
				if logStack {
					stack = string(debug.Stack())
				}

				logger.Error("[Recovery from panic]",
					zap.Any("error", err),
					zap.String("request", string(dump)),
					zap.String("stack", stack),
				)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code":    500,
					"message": fmt.Sprintf("Internal Server Error, requestId: %s", contextvar.XRequestId.Get()),
					"data":    err,
				})
			}
		}()
		c.Next()
	}
}
