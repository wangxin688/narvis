package webssh_api

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/intend/logger"
	"github.com/wangxin688/narvis/server/middleware"
	"github.com/wangxin688/narvis/server/tools/errors"
	"go.uber.org/zap"
)

type WebSSHServer struct {
	Rows int `form:"rows"`
	Cols int `form:"cols"`
}

// @Tags WebSSH
// @Summary WebSSH Server
// @X-func {"name": "CreateWebSSHSession"}
// @Description WebSSH Server
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param deviceId path string true "device id"
// @Param obj query WebSSHServer true "web ssh query"
// @Success 200 {string} string "success"
// @Router /webssh/server/{deviceId} [get]
func webSSH(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	err = handleWebSSHRequest(c)
	if err != nil {
		logger.Logger.Error("[webssh]: failed to handle webssh request", zap.Error(err))
		return
	}

}

// @Tags WebSSH
// @Summary WebSSH Proxy
// @X-func {"name": "ProxyWebSSHCallback"}
// @Description WebSSH Proxy
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param sessionId path string true "session id"
// @Success 200 {string} string "success"
// @Router /webssh/proxy/{sessionId} [get]
func proxyWebSSH(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()

	err = handleProxyWebSocket(c)
	if err != nil {
		logger.Logger.Error("[webssh]: failed to handle proxy webssh request", zap.Error(err))
		return
	}
	//c.String(http.StatusOK, "success")
}

func RegisterWebSSHRoutes(e *gin.Engine) {
	basePath := "/api/v1"
	router := e.Group(basePath+"/webssh", middleware.CookieAuthMiddleware())
	{
		router.GET("/server/:deviceId", webSSH)
	}
	router1 := e.Group(basePath+"/webssh", middleware.ProxyAuthMiddleware())
	{
		router1.GET("/proxy/:sessionId", proxyWebSSH)
	}
}
