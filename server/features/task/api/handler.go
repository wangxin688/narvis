package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/middleware"
)

func RegisterTaskRoutes(e *gin.Engine) {
	basePath := "/api/v1"
	router := e.Group(basePath+"/task", middleware.ProxyAuthMiddleware())
	{
		router.POST("scan-device-basic", scanDeviceBasicInfoCallback)
		router.POST("scan-aps", scanApCallback)
		router.POST("scan-devices", scanDeviceDetailCallback)
		router.POST("config-backup", configBackupCallback)
		router.POST("wlan-users", wlanUserCallback)
	}
}
