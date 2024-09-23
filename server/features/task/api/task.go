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
// @Description Scan device basic information callback
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param rack body []intendtask.DeviceBasicInfoScanResponse true "rack"
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
	err = infra_tasks.DeviceBasicInfoScanCallback(scanDevices)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.SuccessResponse{Status: "ok"})
}
