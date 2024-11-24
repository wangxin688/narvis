package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/intend/model/manufacturer"
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
// @Success 200 {object} schemas.ListResponse{results=[]string}
// @Router /intend/manufacturers [get]
func manufacturerList(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	list := manufacturer.SupportManufacturer()
	c.JSON(http.StatusOK, schemas.ListResponse{
		Total:   int64(len(list)),
		Results: list,
	})
}
