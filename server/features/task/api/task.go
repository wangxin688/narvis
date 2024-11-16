package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/intend/intendtask"
	infra_tasks "github.com/wangxin688/narvis/server/features/infra/tasks"
	task_biz "github.com/wangxin688/narvis/server/features/task/biz"
	"github.com/wangxin688/narvis/server/tools/errors"
	"github.com/wangxin688/narvis/server/tools/helpers"
	ts "github.com/wangxin688/narvis/server/tools/schemas"
)

const xTaskID = "X-Task-ID"

// @Tags Task
// @Summary ScanDevice BasicInfo Callback
// @X-func {"name": "ScanDeviceBasicInfoCallback"}
// @Description Scan device basic information callback
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body []intendtask.DeviceBasicInfoScanResponse true "data"
// @Success 200 {object} ts.SuccessResponse
// @Router /task/scan-device-basic [post]
func scanDeviceBasicInfoCallback(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var scanDevices []*intendtask.DeviceBasicInfoScanResponse
	taskId := c.GetHeader(xTaskID)
	err = helpers.ValidateUuidString(taskId)
	if err != nil {
		return
	}
	if err = c.ShouldBindJSON(&scanDevices); err != nil {
		return
	}
	err = task_biz.UpdateScanDeviceBasicResult(taskId, scanDevices)
	if err != nil {
		return
	}
	if scanDevices != nil {
		err = infra_tasks.DeviceBasicInfoScanCallback(scanDevices)
		if err != nil {
			return
		}
	}
	c.JSON(http.StatusOK, ts.SuccessResponse{Status: "ok"})
}

// @Tags Task
// @Summary ScanAP Callback
// @X-func {"name": "ScanApCallback"}
// @Description Scan AP callback
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body []intendtask.ApScanResponse true "data"
// @Success 200 {object} ts.SuccessResponse
// @Router /task/scan-ap [post]
func scanApCallback(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var scanAp []*intendtask.ApScanResponse
	taskId := c.GetHeader(xTaskID)
	err = helpers.ValidateUuidString(taskId)
	if err != nil {
		return
	}
	if err = c.ShouldBindJSON(&scanAp); err != nil {
		return
	}
	err = task_biz.UpdateApScanResult(taskId, scanAp)
	if err != nil {
		return
	}
	if scanAp != nil {
		err = infra_tasks.ScanApCallback(scanAp)
		if err != nil {
			return
		}
	}
	c.JSON(http.StatusOK, ts.SuccessResponse{Status: "ok"})
}

// @Tags Task
// @Summary ScanDeviceDetail Callback
// @X-func {"name": "ScanDeviceDetailCallback"}
// @Description Scan device detail callback
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body intendtask.DeviceScanResponse true "data"
// @Success 200 {object} ts.SuccessResponse
// @Router /task/scan-device [post]
func scanDeviceDetailCallback(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var scanDevices *intendtask.DeviceScanResponse
	taskId := c.GetHeader(xTaskID)
	err = helpers.ValidateUuidString(taskId)
	if err != nil {
		return
	}
	if err = c.ShouldBindJSON(&scanDevices); err != nil {
		return
	}
	err = task_biz.UpdateScanDeviceResult(taskId, scanDevices)
	if err != nil {
		return
	}
	if scanDevices != nil && len(scanDevices.Errors) == 0 {
		err = infra_tasks.DeviceScanCallback(scanDevices)
		if err != nil {
			return
		}
	}
	c.JSON(http.StatusOK, ts.SuccessResponse{Status: "ok"})
}

// @Tags Task
// @Summary ConfigurationBackup Callback
// @X-func {"name": "ConfigBackupCallback"}
// @Description Configuration backup callback
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body []intendtask.ConfigurationBackupTaskResult true "data"
// @Success 200 {object} ts.SuccessResponse
// @Router /task/config-backup [post]
func configBackupCallback(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var configBackUp *intendtask.ConfigurationBackupTaskResult
	taskId := c.GetHeader(xTaskID)
	err = helpers.ValidateUuidString(taskId)
	if err != nil {
		return
	}
	if err = c.ShouldBindJSON(&configBackUp); err != nil {
		return
	}
	err = task_biz.UpdateConfigBackupResult(taskId, configBackUp)
	if err != nil {
		return
	}
	if configBackUp != nil && configBackUp.Error == "" {
		err = infra_tasks.ConfigBackUpCallback(configBackUp)
		if err != nil {
			return
		}
	}
	c.JSON(http.StatusOK, ts.SuccessResponse{Status: "ok"})
}
