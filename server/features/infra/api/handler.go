package infra_api

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/middleware"
)

func RegisterInfraRoutes(e *gin.Engine) {

	basePath := "/api/v1"
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

		router.POST("/devices", createDevice)
		router.GET("/devices", listDevices)
		router.GET("/devices/:id", getDevice)
		router.PUT("/devices/:id", updateDevice)
		router.DELETE("/devices/:id", deleteDevice)
		router.GET("/devices/:id/interfaces", getDeviceInterfaces)
		router.GET("/devices/:id/panels", getDevicePanel)

		router.POST("/devices/:id/restconf", createRestconfCredential)
		router.GET("/devices/:id/restconf", getRestconfCredential)
		router.PUT("/devices/:id/restconf", updateRestconfCredential)
		router.DELETE("/devices/:id/restconf", deleteRestconfCredential)

		router.POST("/devices/:id/cli", createCliCredential)
		router.GET("/devices/:id/cli", getCliCredential)
		router.PUT("/devices/:id/cli", updateCliCredential)
		router.DELETE("/devices/:id/cli", deleteCliCredential)

		router.POST("/devices/:id/snmpv2", createSnmpV2Credential)
		router.GET("/devices/:id/snmpv2", getSnmpV2Credential)
		router.PUT("/devices/:id/snmpv2", updateSnmpV2Credential)
		router.DELETE("/devices/:id/snmpv2", deleteSnmpV2Credential)

		router.GET("/devices/:id/config", getDeviceBackupConfig)

		router.PUT("/interfaces/:id", updateDeviceInterface)

		router.GET("/aps/:id", getAp)
		router.GET("/aps", listAp)
		router.PUT("/aps/:id", updateAp)
		router.PUT("/aps", batchUpdateAp)
		router.DELETE("/aps/:id", deleteAp)
		router.DELETE("/aps", batchDeleteAp)

		router.POST("/circuits", createCircuit)
		router.GET("/circuits", listCircuit)
		router.GET("/circuits/:id", getCircuit)
		router.PUT("/circuits/:id", updateCircuit)
		router.DELETE("/circuits/:id", deleteCircuit)

		router.POST("/scan-devices", createScanDevice)
		router.GET("/scan-devices/:id", getScanDevice)
		router.GET("/scan-devices", listScanDevices)
		router.PUT("/scan-devices/:id", updateScanDevice)
		router.PUT("/scan-devices", batchUpdateScanDevice)
		router.DELETE("/scan-devices/:id", deleteScanDevice)
		router.DELETE("/scan-devices", batchDeleteScanDevice)

		router.POST("/scan-aps", createScanAP)
		router.POST("/scan-device-details", scanDeviceDetails)
		router.POST("/config-backup", configBackUp)

		router.POST("/servers", createServer)
		router.GET("/servers", listServers)
		router.GET("/servers/:id", getServer)
		router.PUT("/servers/:id", updateServer)
		router.DELETE("/servers/:id", deleteServer)

		router.POST("/servers/:id/cli", createCliCredentialForServer)
		router.GET("/servers/:id/cli", getCliCredentialForServer)
		router.DELETE("/servers/:id/cli", deleteCliCredentialForServer)
		router.PUT("/servers/:id/cli", updateCliCredentialForServer)

		router.POST("servers/:id/snmpv2", createSnmpV2CredentialForServer)
		router.GET("servers/:id/snmpv2", getSnmpV2CredentialForServer)
		router.PUT("servers/:id/snmpv2", updateSnmpV2CredentialForServer)
		router.DELETE("servers/:id/snmpv2", deleteSnmpV2CredentialForServer)
	}
}
