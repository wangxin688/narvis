package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/middleware"
)

func RegisterOrgRoutes(e *gin.Engine) {
	basePath := "/api/v1"
	router := e.Group(basePath+"/org", middleware.AuthMiddleware())
	{
		router.POST("/organizations", orgCreate)
	}
}
