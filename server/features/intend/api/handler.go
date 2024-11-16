package api

import (
	"github.com/gin-gonic/gin"
)

func RegisterIntendRoutes(e *gin.Engine) {
	basePath := "/api/v1"
	router := e.Group(basePath + "/intend")
	{
		router.GET("/device-roles", deviceRoleList)
		router.GET("/manufacturers", manufacturerList)
		router.GET("/platforms", platformList)

	}
}
