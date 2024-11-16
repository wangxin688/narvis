package device360_api

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/middleware"
)

func RegisterDevice360Routes(e *gin.Engine) {

	basePath := "/api/v1"
	router := e.Group(basePath+"/assurance", middleware.AuthMiddleware())
	{
		router.GET("/device-healthy", getDeviceHealthy)
	}
}
