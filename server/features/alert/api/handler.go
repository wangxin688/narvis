package alert_api

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/core"
)

func RegisterAlertRoutes(r *gin.Engine) {
	basePath := core.Settings.System.RouterPrefix
	pubRouter := r.Group(basePath + "/alert")
	{
		pubRouter.POST("/alerts", createAlert)
		pubRouter.POST("/alert-groups", createAlertGroup)
	}
	router := r.Group(basePath + "/alert")
	{
		router.GET("/alerts/:id", getAlert)
		router.GET("/alerts", listAlerts)

		router.POST("/maintenances", createMaintenance)
		router.GET("/maintenances/:id", getMaintenance)
		router.GET("/maintenances", listMaintenances)
		router.PUT("/maintenances/:id", updateMaintenance)
		router.DELETE("/maintenances/:id", deleteMaintenance)

		router.POST("/subscriptions", createSubscription)
		router.GET("/subscriptions/:id", getSubscription)
		router.GET("/subscriptions", listSubscriptions)
		router.PUT("/subscriptions/:id", updateSubscription)
		router.DELETE("/subscriptions/:id", deleteSubscription)
	}
}
