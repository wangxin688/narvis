package infra_api

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/middleware"
)

func RegisterInfraRoutes(e *gin.Engine) {

	basePath := core.Settings.System.RouterPrefix
	router := e.Group(basePath+"/infra", middleware.AuthMiddleware())
	{
		router.POST("/sites", createSite)
		router.GET("/sites", listSites)
		router.GET("/sites/:id", getSite)
		router.PUT("/sites/:id", updateSite)
		router.DELETE("/sites/:id", deleteSite)

		router.POST("/racks", createRack)
		router.GET("/racks", listRacks)
		router.GET("/racks/:id", getRack)
		router.PUT("/racks/:id", updateRack)
		router.DELETE("/racks/:id", deleteRack)

		router.POST("/circuits", createCircuit)
		router.GET("/circuits", listCircuit)
		router.GET("/circuits/:id", getCircuit)
		router.PUT("/circuits/:id", updateCircuit)
		router.DELETE("/circuits/:id", deleteCircuit)

	}
}
