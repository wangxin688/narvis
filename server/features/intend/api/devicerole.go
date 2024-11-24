package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/intend/model/devicerole"
	"github.com/wangxin688/narvis/server/tools/errors"
	"github.com/wangxin688/narvis/server/tools/schemas"
)

// @Tags Intend
// @Security BearerAuth
// @Summary Get list of device roles
// @X-func {"name": "DeviceRoleList"}
// @Description Get list of device roles
// @Accept json
// @Produce json
// @Success 200 {object} schemas.ListResponse{results=[]string}
// @Router /intend/device-roles [get]
func deviceRoleList(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	count, list := devicerole.GetListDeviceRole()

	c.JSON(http.StatusOK, schemas.ListResponse{
		Total:   count,
		Results: list,
	})
}
