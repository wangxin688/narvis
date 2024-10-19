package infra_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/intend/intendtask"
	infra_biz "github.com/wangxin688/narvis/server/features/infra/biz"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	infra_tasks "github.com/wangxin688/narvis/server/features/infra/tasks"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/tools/errors"
	"github.com/wangxin688/narvis/server/tools/helpers"
	ts "github.com/wangxin688/narvis/server/tools/schemas"
)

// @Tags Infra
// @Summary Scan Device Create
// @Description Scan Device Create
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param data body schemas.ScanDeviceCreate true "data"
// @Success 200 {object} schemas.ScanDeviceCreateResponse
// @Router /infra/scan-devices [post]
func createScanDevice(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var scanDeviceCreate schemas.ScanDeviceCreate
	if err = c.ShouldBindJSON(&scanDeviceCreate); err != nil {
		return
	}
	scanDeviceCreate.SetDefaultValue()
	orgId := global.OrganizationId.Get()
	taskIds, err := infra_tasks.CreateScanTask(&scanDeviceCreate, orgId)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, schemas.ScanDeviceCreateResponse{TaskIds: taskIds})
}

// @Tags Infra
// @Summary Scan Device Update
// @Description Scan Device Update
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param id path string true "uuid formatted scanDeviceId"
// @Param data body schemas.ScanDeviceUpdate true "data"
// @Success 200 {object} ts.IdResponse
// @Router /infra/scan-devices/{id} [put]
func updateScanDevice(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var scanDeviceUpdate schemas.ScanDeviceUpdate
	id := c.Param("id")
	if err = helpers.ValidateUuidString(id); err != nil {
		return
	}
	if err = c.ShouldBindJSON(&scanDeviceUpdate); err != nil {
		return
	}
	id, err = infra_biz.NewScanDeviceService().UpdateById(id, &scanDeviceUpdate)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: id})
}

// @Tags Infra
// @Summary Scan Device Batch Update
// @Description Scan Device Batch Update
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param data body schemas.ScanDeviceBatchUpdate true "data"
// @Success 200 {object} ts.IdsResponse
// @Router /infra/scan-devices [put]
func batchUpdateScanDevice(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var scanDeviceBatchUpdate schemas.ScanDeviceBatchUpdate
	if err = c.ShouldBindJSON(&scanDeviceBatchUpdate); err != nil {
		return
	}
	ids, err := infra_biz.NewScanDeviceService().BatchUpdate(&scanDeviceBatchUpdate)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdsResponse{Ids: ids})
}

// @Tags Infra
// @Summary Scan Device List
// @Description Scan Device List
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param object query schemas.ScanDeviceQuery true "query"
// @Success 200 {object} ts.ListResponse{results=[]schemas.ScanDevice}
// @Router /infra/scan-devices [get]
func listScanDevices(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var scanDeviceQuery schemas.ScanDeviceQuery
	if err = c.ShouldBindQuery(&scanDeviceQuery); err != nil {
		return
	}
	count, scanDevices, err := infra_biz.NewScanDeviceService().ListScanDevice(&scanDeviceQuery)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.ListResponse{Total: count, Results: scanDevices})
}

// @Tags Infra
// @Summary Scan Device Get
// @Description Scan Device Get
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param id path string true "uuid formatted scanDeviceId"
// @Success 200 {object} schemas.ScanDevice
// @Router /infra/scan-devices/{id} [get]
func getScanDevice(c *gin.Context) {
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
	scanDevice, err := infra_biz.NewScanDeviceService().GetById(id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, scanDevice)
}

// @Tags Infra
// @Summary Scan Device Delete
// @Description Scan Device Delete
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param id path string true "uuid formatted scanDeviceId"
// @Success 200 {object} ts.IdResponse
// @Router /infra/scan-devices/{id} [delete]
func deleteScanDevice(c *gin.Context) {
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
	err = infra_biz.NewScanDeviceService().DeleteById(id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: id})
}

// @Tags Infra
// @Summary Scan Device Batch Delete
// @Description Scan Device Batch Delete
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param ids body []string true "ids"
// @Success 200 {object} ts.IdsResponse
// @Router /infra/scan-devices [delete]
func batchDeleteScanDevice(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var ids []string
	if err = c.ShouldBindJSON(&ids); err != nil {
		return
	}
	err = infra_biz.NewScanDeviceService().DeleteByIds(ids)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdsResponse{Ids: ids})
}

// @Tags Infra
// @Summary Scan AP
// @Description Scan AP
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param data body schemas.ScanApCreate true "data"
// @Success 200 {object} ts.IdsResponse
// @Router /infra/scan-aps [post]
func createScanAP(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var scanAP schemas.ScanApCreate
	if err = c.ShouldBindJSON(&scanAP); err != nil {
		return
	}

	taskIds, err := infra_tasks.GenerateSNMPTask(scanAP.SiteId, intendtask.ScanAp, intendtask.ScanApCallback)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdsResponse{Ids: taskIds})
}

// @Tags Infra
// @Summary Scan device details
// @Description Scan device details
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param data body schemas.ScanDeviceDetailTask true "data"
// @Success 200 {object} ts.IdsResponse
// @Router /infra/scan-device-details [post]
func scanDeviceDetails(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var scanDeviceDetailTask schemas.ScanDeviceDetailTask
	if err = c.ShouldBindJSON(&scanDeviceDetailTask); err != nil {
		return
	}
	taskIds, err := infra_tasks.GenerateSNMPTask(scanDeviceDetailTask.SiteId, intendtask.ScanDevice, intendtask.ScanDeviceCallback)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdsResponse{Ids: taskIds})
}

// @Tags Infra
// @Summary Device Configuration backup
// @Description Device Configuration backup
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param data body schemas.ConfigBackUpCreate true "data"
// @Success 200 {object} ts.IdsResponse
// @Router /infra/device-config-backup [post]

func configBackUp(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var configBackUpCreate schemas.ConfigBackUpCreate
	if err = c.ShouldBindJSON(&configBackUpCreate); err != nil {
		return
	}
	taskIds, err := infra_tasks.ConfigBackUpTask(
		configBackUpCreate.SiteId,
		intendtask.ConfigurationBackup,
		intendtask.ConfigurationBackupCallback)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdsResponse{Ids: taskIds})
}
