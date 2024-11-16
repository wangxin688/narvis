package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/middleware"
)

func RegisterAdminRoutes(e *gin.Engine) {
	basePath := "/api/v1"
	router := e.Group(basePath+"/admin", middleware.AuthMiddleware())
	{

		router.POST("/users", createUser)
		router.GET("/users", listUsers)
		router.GET("/users/me", getUserMe)
		router.GET("/users/:id", getUser)
		router.PUT("/users/me", updateUserMe)
		router.PUT("/users/:id", updateUser)
		router.DELETE("/users/:id", deleteUser)

		router.POST("/roles", createRole)
		router.GET("/roles", listRoles)
		router.GET("/roles/:id", getRole)
		router.PUT("/roles/:id", updateRole)
		router.DELETE("/roles/:id", deleteRole)
	}
}

func RegisterLoginRoutes(e *gin.Engine) {
	basePath := "/api/v1"
	router := e.Group(basePath + "/login")
	{
		router.POST("/password", passwordLogin)
	}
	router.POST("/refresh", refreshToken, middleware.AuthMiddleware())
}
