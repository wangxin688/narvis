package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/middleware"
)

func RegisterTaskRoutes(e *gin.Engine) {
	basePath := core.Settings.System.RouterPrefix
	router := e.Group(basePath+"/task", middleware.ProxyAuthMiddleware())
	{
		router.POST("scan-device-basic", scanDeviceBasicInfoCallback)
		router.POST("scan-ap", scanApCallback)
		router.POST("scan-device", scanDeviceDetailCallback)
		router.POST("config-backup", configBackupCallback)
	}
}
