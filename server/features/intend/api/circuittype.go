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
// @Summary Get list of CircuitTypes
// @Param object query schemas.CircuitTypeQuery false "query CircuitTypes"
// @Success 200 {object} schemas.ListResponse{results=[]schemas.CircuitType}
// @Router /intend/circuit-types [get]
func circuitTypeList(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var req its.CircuitTypeQuery
	if err := c.ShouldBindQuery(&req); err != nil {
		return
	}
	count, list := biz.GetCircuitTypes(&req)

	c.JSON(http.StatusOK, schemas.ListResponse{
		Total:   count,
		Results: list,
	})
}
