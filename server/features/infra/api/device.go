package infra_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	infra_biz "github.com/wangxin688/narvis/server/features/infra/biz"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
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
	if err = infra_biz.NewDeviceService().UpdateDevice(c, deviceId, &device); err != nil {
		return
	}
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
	err = infra_biz.NewDeviceService().DeleteDevice(deviceId)
	if err != nil {
		return
	}
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
