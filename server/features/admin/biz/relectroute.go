package biz

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/intend/logger"
)

// router := gin.Default()

// 注册路由时传递 gin.Engine
// router.GET("/route1", routeHandler(router))
// router.POST("/route2", routeHandler(router))

func ReflectRouteToDb(router *gin.Engine) gin.HandlerFunc {
	return func(_ *gin.Context) {
		routes := router.Routes()
		for _, route := range routes {
			logger.Logger.Info(route.Method + " " + route.Path)
			// add route to database here
		}
	}
}
