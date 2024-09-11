package infra_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	infra_biz "github.com/wangxin688/narvis/server/features/infra/biz"
	"github.com/wangxin688/narvis/server/features/infra/hooks"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	"github.com/wangxin688/narvis/server/tools"
	"github.com/wangxin688/narvis/server/tools/errors"
	"github.com/wangxin688/narvis/server/tools/helpers"
	ts "github.com/wangxin688/narvis/server/tools/schemas"
)

// @Tags Infra
// @Summary Create device
// @Description Create device
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param device body schemas.DeviceCreate true "device"
// @Success 200 {object} ts.IdResponse
// @Router /infra/devices [post]
func createDevice(c *gin.Context) {
	var device schemas.DeviceCreate
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	if err = c.ShouldBindJSON(&device); err != nil {
		return
	}
	newDevice, err := infra_biz.NewDeviceService().CreateDevice(&device)
	if err != nil {
		return
	}
	tools.BackgroundTask(func() {
		hooks.DeviceCreateHooks(newDevice)
	})
	c.JSON(http.StatusOK, ts.IdResponse{Id: newDevice})

}

// @Tags Infra
// @Summary Update device
// @Description Update device
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param device body schemas.DeviceUpdate true "device"
// @Success 200 {object} ts.IdResponse
// @Router /infra/devices/{id} [put]
func updateDevice(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	deviceId := c.Param("id")
	var device schemas.DeviceUpdate
	if err = c.ShouldBindJSON(&device); err != nil {
		return
	}
	diff, err := infra_biz.NewDeviceService().UpdateDevice(c, deviceId, &device)
	if err != nil {
		return
	}
	tools.BackgroundTask(func() {
		hooks.DeviceUpdateHooks(deviceId, diff[deviceId])
	})
	c.JSON(http.StatusOK, ts.IdResponse{Id: deviceId})
}

// @Tags Infra
// @Summary Get device
// @Description get device
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} schemas.Device
// @Router /infra/devices/{id} [get]
func getDevice(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	deviceId := c.Param("id")
	if err = helpers.ValidateUuidString(deviceId); err != nil {
		return
	}
	device, err := infra_biz.NewDeviceService().GetById(deviceId)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, device)
}

// @Tags Infra
// @Summary Delete device
// @Description Delete device
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} ts.IdResponse
// @Router /infra/devices/{id} [delete]
func deleteDevice(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	deviceId := c.Param("id")
	if err = helpers.ValidateUuidString(deviceId); err != nil {
		return
	}
	device, err := infra_biz.NewDeviceService().DeleteDevice(deviceId)
	if err != nil {
		return
	}
	tools.BackgroundTask(func() {
		hooks.DeviceDeleteHooks(device)
	})
	c.JSON(http.StatusOK, ts.IdResponse{Id: deviceId})
}

// @Tags Infra
// @Summary List devices
// @Description List devices
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param query query schemas.DeviceQuery true "query"
// @Success 200 {object} ts.ListResponse{data=[]schemas.Device}
// @Router /infra/devices [get]
func listDevices(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var query schemas.DeviceQuery
	if err = c.ShouldBindQuery(&query); err != nil {
		return
	}
	count, devices, err := infra_biz.NewDeviceService().GetDeviceList(&query)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.ListResponse{Total: count, Results: devices})
}

// @Tags Infra
// @Summary Get device interfaces
// @Description Get device interfaces
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} []schemas.DeviceInterface
// @Router /infra/devices/{id}/interfaces [get]
func getDeviceInterfaces(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	deviceId := c.Param("id")
	if err = helpers.ValidateUuidString(deviceId); err != nil {
		return
	}
	interfaces, err := infra_biz.NewDeviceService().GetDeviceInterfaces(deviceId)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, interfaces)
}

// @Tags Infra
// @Summary Create new device restconf credential
// @Description Create new device restconf credential
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "deviceId"
// @Param credential body schemas.RestconfCredentialCreate true "Credential"
// @Success 200 {object} ts.IdResponse
// @Router /infra/devices/{id}/restconf [post]
func createRestconfCredential(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var credential schemas.RestconfCredentialCreate
	deviceId := c.Param("id")
	if err = helpers.ValidateUuidString(deviceId); err != nil {
		return
	}
	if err = c.ShouldBindJSON(&credential); err != nil {
		return
	}
	id, err := infra_biz.NewRestConfCredentialService().CreateCredential(deviceId, &credential)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: id})
}

// @Tags Infra
// @Summary Get device restconf credential
// @Description Get device restconf credential
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "deviceId"
// @Success 200 {object} schemas.RestconfCredential
// @Router /infra/devices/{id}/restconf [get]
func getRestconfCredential(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	id := c.Param("id")
	if err = helpers.ValidateUuidString(id); err != nil {
		return
	}
	credential, err := infra_biz.NewRestConfCredentialService().GetCredentialByDeviceId(id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, credential)
}

// @Tags Infra
// @Summary Update device restconf credential
// @Description Update device restconf credential
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "deviceId"
// @Param credential body schemas.RestconfCredentialUpdate true "Credential"
// @Success 200 {object} ts.IdResponse
// @Router /infra/devices/{id}/restconf [put]
func updateRestconfCredential(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	id := c.Param("id")
	if err = helpers.ValidateUuidString(id); err != nil {
		return
	}
	var credential schemas.RestconfCredentialUpdate
	if err = c.ShouldBindJSON(&credential); err != nil {
		return
	}
	_, err = infra_biz.NewRestConfCredentialService().UpdateCredential(id, &credential)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: id})
}

// @Tags Infra
// @Summary Delete device restconf credential
// @Description Delete device restconf credential
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "deviceId"
// @Success 200 {object} ts.IdResponse
// @Router /infra/devices/{id}/restconf [delete]
func deleteRestconfCredential(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	id := c.Param("id")
	if err = helpers.ValidateUuidString(id); err != nil {
		return
	}
	err = infra_biz.NewRestConfCredentialService().DeleteCredential(id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: id})
}

// @Tags Infra
// @Summary Create device new cli credential
// @Description Create device new cli credential
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "DeviceId"
// @Param credential body schemas.CliCredentialCreate true "Credential"
// @Success 200 {object} ts.IdResponse
// @Router /infra/devices/{id}/cli [post]
func createCliCredential(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var credential schemas.CliCredentialCreate
	deviceId := c.Param("id")
	if err = helpers.ValidateUuidString(deviceId); err != nil {
		return
	}
	if err = c.ShouldBindJSON(&credential); err != nil {
		return
	}
	id, err := infra_biz.NewCliCredentialService().CreateCredential(deviceId, &credential)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: id})
}

// @Tags Infra
// @Summary Get device cli credential
// @Description Get device cli credential
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "deviceId"
// @Success 200 {object} schemas.CliCredential
// @Router /infra/devices/{id}/cli [get]
func getCliCredential(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	id := c.Param("id")
	if err = helpers.ValidateUuidString(id); err != nil {
		return
	}
	credential, err := infra_biz.NewCliCredentialService().GetCredentialByDeviceId(id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, credential)
}

// @Tags Infra
// @Summary Update device cli credential
// @Description Update device cli credential
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "deviceId"
// @Param credential body schemas.CliCredentialUpdate true "Credential"
// @Success 200 {object} ts.IdResponse
// @Router /infra/devices/{id}/cli [put]
func updateCliCredential(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	id := c.Param("id")
	if err = helpers.ValidateUuidString(id); err != nil {
		return
	}
	var credential schemas.CliCredentialUpdate
	if err = c.ShouldBindJSON(&credential); err != nil {
		return
	}
	_, err = infra_biz.NewCliCredentialService().UpdateCredential(id, &credential)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: id})
}

// @Tags Infra
// @Summary Delete device cli credential
// @Description Delete device cli credential
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "deviceId"
// @Success 200 {object} ts.IdResponse
// @Router /infra/devices/{id}/cli [delete]
func deleteCliCredential(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	id := c.Param("id")
	if err = helpers.ValidateUuidString(id); err != nil {
		return
	}
	err = infra_biz.NewCliCredentialService().DeleteCredential(id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: id})
}

// @Tags Infra
// @Summary Create new device snmpV2 credential
// @Description Create device new snmpV2 credential
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "deviceId"
// @Param credential body schemas.SnmpV2CredentialCreate true "Credential"
// @Success 200 {object} ts.IdResponse
// @Router /infra/devices/{id}/snmpv2 [post]
func createSnmpV2Credential(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var credential schemas.SnmpV2CredentialCreate
	deviceId := c.Param("id")
	if err = helpers.ValidateUuidString(deviceId); err != nil {
		return
	}
	if err = c.ShouldBindJSON(&credential); err != nil {
		return
	}
	id, err := infra_biz.NewSnmpCredentialService().CreateSnmpCredential(deviceId, &credential)
	if err != nil {
		return
	}
	hooks.SnmpCredCreateHooks(id)
	c.JSON(http.StatusOK, ts.IdResponse{Id: deviceId})
}

// @Tags Infra
// @Summary Get device snmpV2 credential
// @Description Get device snmpV2 credential
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "deviceId"
// @Success 200 {object} schemas.SnmpV2Credential
// @Router /infra/devices/{id}/snmpv2 [get]
func getSnmpV2Credential(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	id := c.Param("id")
	if err = helpers.ValidateUuidString(id); err != nil {
		return
	}
	credential, err := infra_biz.NewSnmpCredentialService().GetCredentialByDeviceId(id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, credential)
}

// @Tags Infra
// @Summary Update device snmpV2 credential
// @Description Update device snmpV2 credential
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "deviceId"
// @Param credential body schemas.SnmpV2CredentialUpdate true "Credential"
// @Success 200 {object} ts.IdResponse
// @Router /infra/devices/{id}/snmpv2 [put]
func updateSnmpV2Credential(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	deviceId := c.Param("id")
	if err = helpers.ValidateUuidString(deviceId); err != nil {
		return
	}
	var credential schemas.SnmpV2CredentialUpdate
	if err = c.ShouldBindJSON(&credential); err != nil {
		return
	}
	credId, diff, err := infra_biz.NewSnmpCredentialService().UpdateSnmpCredential(deviceId, &credential)
	if err != nil {
		return
	}
	tools.BackgroundTask(func() {
		hooks.SnmpCredUpdateHooks(credId, diff[credId])
	})
	c.JSON(http.StatusOK, ts.IdResponse{Id: deviceId})
}

// @Tags Infra
// @Summary Delete device snmpV2 credential
// @Description Delete device snmpV2 credential
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "deviceId"
// @Success 200 {object} ts.IdResponse
// @Router /infra/devices/{id}/snmpv2 [delete]
func deleteSnmpV2Credential(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	id := c.Param("id")
	if err = helpers.ValidateUuidString(id); err != nil {
		return
	}
	cred, err := infra_biz.NewSnmpCredentialService().DeleteCredential(id)
	if err != nil {
		return
	}
	tools.BackgroundTask(func() {
		hooks.SnmpCredDeleteHooks(cred)
	})
	c.JSON(http.StatusOK, ts.IdResponse{Id: id})
}
