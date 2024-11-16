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
// @Summary Create Alert Group
// @X-func {"name": "CreateAlertGroup"}
// @Description Create Alert Group
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param data body schemas.AlertGroupCreate true "data"
// @Success 200 {object} ts.IdResponse
// @Router /alert/alert-groups [post]
func createAlertGroup(c *gin.Context) {
	var err error

	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()

	var alertGroup schemas.AlertGroupCreate
	if err = c.ShouldBindJSON(&alertGroup); err != nil {
		return
	}
	newAlertGroup, err := alert_biz.NewAlertGroupService().CreateAlertGroup(&alertGroup)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: newAlertGroup.Id})
}
