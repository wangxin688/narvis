package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/core"
)

func RegisterAdminRoutes(e *gin.Engine) {
	basePath := core.Settings.System.RouterPrefix
	router := e.Group(basePath)
	{
		router.POST("/login/password", passwordLogin)
	}
}
