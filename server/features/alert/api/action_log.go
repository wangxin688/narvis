package alert_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	alert_biz "github.com/wangxin688/narvis/server/features/alert/biz"
	"github.com/wangxin688/narvis/server/features/alert/schemas"
	"github.com/wangxin688/narvis/server/tools/errors"
	ts "github.com/wangxin688/narvis/server/tools/schemas"
)

// @Tags Alert
// @Summary Create Action Log
// @Description Create Action Log
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param object body schemas.ActionLogCreate true "Create Action Log"
// @Success 200 {object} ts.IdResponse
// @Router /alert/action-logs [post]
func createActionLog(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var actionLog schemas.ActionLogCreate
	if err = c.ShouldBindJSON(&actionLog); err != nil {
		return
	}
	logId, err := alert_biz.NewActionLogService().CreateActionLog(&actionLog)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: logId})
}
