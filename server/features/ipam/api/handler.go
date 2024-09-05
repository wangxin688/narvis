package ipam_api

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/middleware"
)

func RegisterIpamRoutes(e *gin.Engine) {

	basePath := core.Settings.System.RouterPrefix
	router := e.Group(basePath+"/ipam", middleware.AuthMiddleware())
	{

		router.GET("/prefixes", getPrefixList)
		router.POST("/prefixes", createPrefix)
		router.GET("/prefixes/:id", getPrefix)
		router.PUT("/prefixes/:id", updatePrefix)
		router.DELETE("/prefixes/:id", deletePrefix)
	}
}
