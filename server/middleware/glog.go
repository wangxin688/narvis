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
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wangxin688/narvis/server/pkg/contextvar"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var XRequestIdHeaderName = "X-Request-ID"

func ZapLoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		requestID := contextvar.XRequestId.Get()
		if requestID == "" {
			requestID = uuid.New().String()
			contextvar.XRequestId.Set(requestID)
			c.Writer.Header().Set(XRequestIdHeaderName, requestID)
		}

		c.Next()

		endTime := time.Now()
		latency := endTime.Sub(startTime)
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()

		fields := []zapcore.Field{
			zap.Int("status", statusCode),
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.String("client_ip", clientIP),
			zap.Duration("latency", latency),
		}

		logger.Info(requestID, fields...)

		if len(c.Errors) > 0 {
			for _, err := range c.Errors.Errors() {
				logger.Error(requestID, zap.String("error", err))
			}
		}
	}
}
