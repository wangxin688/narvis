package middleware

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/wangxin688/narvis/server/global"
	"go.uber.org/zap"
)

var XRequestIdHeaderName = "X-Request-ID"

func ZapLoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 开始时间
		startTime := time.Now()
		// 获取或生成request id
		requestId := global.XRequestId.Get()
		if requestId == "" {
			_requestId := uuid.New().String()
			global.XRequestId.Set(_requestId)
			c.Writer.Header().Set(XRequestIdHeaderName, _requestId)
			requestId = _requestId
		}
		// 处理请求
		c.Next()
		// 结束时间
		endTime := time.Now()
		// 执行时间
		latencyTime := endTime.Sub(startTime)
		// 请求方式
		reqMethod := c.Request.Method
		// 请求路由
		reqPath := c.Request.URL.Path
		// 状态码
		statusCode := c.Writer.Status()
		// 请求IP
		clientIP := c.ClientIP()
		// 日志格式
		logger.Info(requestId,
			zap.Int("status", statusCode),
			zap.String("method", reqMethod),
			zap.String("path", reqPath),
			zap.String("client_ip", clientIP),
			zap.Duration("latency", latencyTime),
		)

		if len(c.Errors) > 0 {
			for _, e := range c.Errors.Errors() {
				zap.L().Error("Request error", zap.String("requestId", requestId), zap.String("error", e))
			}
		}
	}
}

// GinRecovery recovers from panics in the gin framework.
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
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
