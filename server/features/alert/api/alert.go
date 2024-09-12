package alert_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	alert_biz "github.com/wangxin688/narvis/server/features/alert/biz"
	"github.com/wangxin688/narvis/server/features/alert/schemas"
	"github.com/wangxin688/narvis/server/tools/errors"
)

// @Tags Create Alert
// @Summary Create Alert
// @Description Create Alert
// @Accept json
// @Produce json
// @Param data body schemas.AlertCreate true "data"
// @Success 200 {object} ts.IdResponse
// @Router /alert [post]
func createAlert(c *gin.Context) {
	var err error

	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()

	var alert schemas.AlertCreate
	if err = c.ShouldBindJSON(&alert); err != nil {
		return
	}

	if err = alert.Validate(); err != nil {
		return
	}
	newAlert, err := alert_biz.NewAlertService().CreateAlert(&alert)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, newAlert.Id)
}
