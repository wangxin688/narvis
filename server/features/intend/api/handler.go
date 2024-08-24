package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/core"
)

func RegisterIntendRoutes(e *gin.Engine) {
	basePath := core.Settings.System.RouterPrefix
	router := e.Group(basePath + "/intend")
	{
		router.GET("/device-roles", deviceRoleList)
		router.GET("/circuit-types", circuitTypeList)
	}
}
