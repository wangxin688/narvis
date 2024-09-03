package infra_api

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/core"
)

func RegisterInfraRoutes(e *gin.Engine) {

	basePath := core.Settings.System.RouterPrefix
	router := e.Group(basePath + "/infra")
	{

	}
}
