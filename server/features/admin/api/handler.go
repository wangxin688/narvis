package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/middleware"
)

func RegisterAdminRoutes(e *gin.Engine) {
	basePath := core.Settings.System.RouterPrefix
	router := e.Group(basePath+"/admin", middleware.AuthMiddleware())
	{

		router.POST("/users", createUser)
		router.GET("/users", listUsers)
		router.GET("/users/:id", getUser)
		router.PUT("/users/:id", updateUser)
		router.DELETE("/users/:id", deleteUser)

		router.POST("/groups", createGroup)
		router.GET("/groups", listGroups)
		router.GET("/groups/:id", getGroup)
		router.PUT("/groups/:id", updateGroup)
		router.DELETE("/groups/:id", deleteGroup)

		router.POST("/roles", createRole)
		router.GET("/roles", listRoles)
		router.GET("/roles/:id", getRole)
		router.PUT("/roles/:id", updateRole)
		router.DELETE("/roles/:id", deleteRole)
	}
}

func RegisterLoginRoutes(e *gin.Engine) {
	basePath := core.Settings.System.RouterPrefix
	router := e.Group(basePath + "/login")
	{
		router.POST("/password", passwordLogin)
	}
}
