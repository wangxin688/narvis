package monitor_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	metric_biz "github.com/wangxin688/narvis/server/features/monitor/biz"
	"github.com/wangxin688/narvis/server/tools/errors"
)

// @Tags Monitor
// @Summary Get Time Series Data
// @Description Get Time Series Data
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param obj query metric_biz.MetricService true "metric query"
// @Success 200 {object} []vtm.MatrixResponse
// @Router /monitor/time-series [get]
func getTimeSeries(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()

	var query metric_biz.MetricService
	if err = c.ShouldBindQuery(&query); err != nil {
		return
	}
	res, err := query.QueryMatrix()
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, res)
}
