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
// @Summary Get list of Manufacturers
// @X-func {"name": "ManufacturerList"}
// @Description Get list of Manufacturers
// @Accept json
// @Produce json
// @Param object query schemas.ManufacturerQuery false "query manufacturers"
// @Success 200 {object} schemas.ListResponse{results=[]string}
// @Router /intend/manufacturers [get]
func manufacturerList(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var req its.ManufacturerQuery
	if err := c.ShouldBindQuery(&req); err != nil {
		return
	}
	count, list := biz.GetManufacturers(&req)

	c.JSON(http.StatusOK, schemas.ListResponse{
		Total:   count,
		Results: list,
	})
}
