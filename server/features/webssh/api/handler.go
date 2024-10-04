package webssh_api

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/middleware"
	"go.uber.org/zap"
)

// @Tags WebSSH
// @Summary WebSSH Server
// @Description WebSSH Server
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param deviceId path string true "device id"
// @Success 200 {string} string "success"
// @Router /webssh/server/{deviceId} [get]
func webSSH(c *gin.Context) {
	err := handleWebSSHRequest(c)
	if err != nil {
		core.Logger.Error("[webssh]: failed to handle webssh request", zap.Error(err))
		return
	}
	//c.String(http.StatusOK, "success")
}

// @Tags WebSSH
// @Summary WebSSH Proxy
// @Description WebSSH Proxy
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param sessionId path string true "session id"
// @Success 200 {string} string "success"
// @Router /webssh/proxy/{sessionId} [post]
func proxyWebSSH(c *gin.Context) {

	err := handleProxyWebSocket(c)
	if err != nil {
		core.Logger.Error("[webssh]: failed to handle proxy webssh request", zap.Error(err))
		return
	}
	//c.String(http.StatusOK, "success")
}

func RegisterWebSSHRoutes(e *gin.Engine) {
	basePath := core.Settings.System.RouterPrefix
	router := e.Group(basePath+"/webssh", middleware.AuthMiddleware())
	{
		router.GET("/server/:deviceId", webSSH)
	}
	router1 := e.Group(basePath+"/webssh", middleware.ProxyAuthMiddleware())
	{
		router1.GET("/proxy/:sessionId", proxyWebSSH)
	}
}
