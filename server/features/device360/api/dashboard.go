package device360_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	device360_biz "github.com/wangxin688/narvis/server/features/device360/biz"
	"github.com/wangxin688/narvis/server/features/device360/schemas"
	"github.com/wangxin688/narvis/server/tools/errors"
)

// @Tags Assurance
// @Summary Get Device Healthy
// @X-func {"name": "GetDeviceHealthy"}
// @Description Get Device Healthy
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param object query schemas.DeviceHealthQuery true "query"
// @Success 200 {object} schemas.HealthHeatMap
// @Router /assurance/device-healthy [get]
func getDeviceHealthy(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	query := &schemas.DeviceHealthQuery{}
	if err = c.ShouldBindQuery(query); err != nil {
		return
	}
	healthy, err := device360_biz.GetHealth(query)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, healthy)
}
