package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/features/intend/biz"
	its "github.com/wangxin688/narvis/server/features/intend/schemas"
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
// @Param object query schemas.DeviceRoleQuery false "query device roles"
// @Success 200 {object} schemas.ListResponse{results=[]devicerole.DeviceRole}
// @Router /intend/device-roles [get]
func deviceRoleList(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var req its.DeviceRoleQuery
	if err := c.ShouldBindQuery(&req); err != nil {
		return
	}
	count, list := biz.GetDeviceRoles(&req)

	c.JSON(http.StatusOK, schemas.ListResponse{
		Total:   count,
		Results: list,
	})
}
