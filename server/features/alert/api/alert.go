package alert_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	alert_biz "github.com/wangxin688/narvis/server/features/alert/biz"
	"github.com/wangxin688/narvis/server/features/alert/schemas"
	"github.com/wangxin688/narvis/server/tools/errors"
	"github.com/wangxin688/narvis/server/tools/helpers"
	ts "github.com/wangxin688/narvis/server/tools/schemas"
)

// @Tags Alert
// @Summary Create Alert
// @X-func {"name": "CreateAlert"}
// @Description Create Alert
// @Accept json
// @Produce json
// @Param data body schemas.AlertCreate true "data"
// @Success 200 {object} ts.IdResponse
// @Router /alert/alerts [post]
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

// @Tags Alert
// @Summary Get Alert
// @X-func {"name": "GetAlert"}
// @Description Get Alert
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted alertId"
// @Success 200 {object} schemas.Alert
// @Router /alert/alerts/{id} [get]
func getAlert(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	alertId := c.Param("id")
	if err = helpers.ValidateUuidString(alertId); err != nil {
		return
	}
	alert, err := alert_biz.NewAlertService().GetById(alertId)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, alert)
}

// @Tags Alert
// @Summary List Alerts
// @X-func {"name": "ListAlerts"}
// @Description List Alerts
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param object query schemas.AlertQuery true "query"
// @Success 200 {object} ts.ListResponse{results=[]schemas.Alert}
// @Router /alert/alerts [get]
func listAlerts(c *gin.Context) {
	var err error

	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var query schemas.AlertQuery
	if err = c.ShouldBindQuery(&query); err != nil {
		return
	}

	count, alerts, err := alert_biz.NewAlertService().ListAlerts(query)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.ListResponse{Total: count, Results: alerts})
}
