package services

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/core"
)

// router := gin.Default()

// 注册路由时传递 gin.Engine
// router.GET("/route1", routeHandler(router))
// router.POST("/route2", routeHandler(router))

func ReflectRouteToDb(router *gin.Engine) gin.HandlerFunc {
	return func(c *gin.Context) {
		routes := router.Routes()
		for _, route := range routes {
			core.Logger.Info(route.Method + " " + route.Path)
			// add route to database here
		}
	}
}
