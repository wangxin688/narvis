package monitor_api

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/middleware"
)

func RegisterMonitorRoutes(r *gin.Engine) {
	basePath := "/api/v1"
	router := r.Group(basePath+"/monitor", middleware.AuthMiddleware())
	{
		router.GET("/time-series", getTimeSeries)
	}
}
