package infra_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	biz "github.com/wangxin688/narvis/server/features/infra/biz"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	"github.com/wangxin688/narvis/server/tools/errors"
	"github.com/wangxin688/narvis/server/tools/helpers"
	ts "github.com/wangxin688/narvis/server/tools/schemas"
)

// @Tags Infra.Rack
// @Summary Create new rack
// @X-func {"name": "CreateRack"}
// @Description Create new rack
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param rack body schemas.RackCreate true "rack"
// @Success 200 {object} ts.IdResponse
// @Router /infra/racks [post]
func createRack(c *gin.Context) {
	var rack schemas.RackCreate
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	if err = c.ShouldBindJSON(&rack); err != nil {
		return
	}
	newRack, err := biz.NewRackService().CreateRack(&rack)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: newRack})
}

// @Tags Infra.Rack
// @Summary Get rack
// @X-func {"name": "GetRack"}
// @Description Get rack
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "rackId"
// @Success 200 {object} schemas.RackElevation
// @Router /infra/racks/{id} [get]
func getRack(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	rackId := c.Param("id")
	if err := helpers.ValidateUuidString(rackId); err != nil {
		return
	}
	rack, err := biz.NewRackService().GetRackElevation(rackId)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, rack)
}

// @Tags Infra.Rack
// @Summary List racks
// @X-func {"name": "ListRacks"}
// @Description List racks
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param object query schemas.RackQuery false "query racks"
// @Success 200 {object} ts.ListResponse{results=[]schemas.Rack}
// @Router /infra/racks [get]
func listRacks(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var req schemas.RackQuery
	if err = c.ShouldBindQuery(&req); err != nil {
		return
	}
	count, list, err := biz.NewRackService().ListRacks(&req)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.ListResponse{
		Total:   count,
		Results: list,
	})
}

// @Tags Infra.Rack
// @Summary Update rack
// @X-func {"name": "UpdateRack"}
// @Description Update rack
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "rackId"
// @Param rack body schemas.RackUpdate true "rack"
// @Success 200 {object} ts.IdResponse
// @Router /infra/racks/{id} [put]
func updateRack(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	rackId := c.Param("id")
	var rack schemas.RackUpdate
	if err = c.ShouldBindJSON(&rack); err != nil {
		return
	}
	err = biz.NewRackService().UpdateRack(rackId, &rack)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: rackId})
}

// @Tags Infra.Rack
// @Summary Delete rack
// @X-func {"name": "DeleteRack"}
// @Description Delete rack
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "rackId"
// @Success 200 {object} ts.IdResponse
// @Router /infra/racks/{id} [delete]
func deleteRack(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	rackId := c.Param("id")
	if err = helpers.ValidateUuidString(rackId); err != nil {
		return
	}
	if err = biz.NewRackService().DeleteRack(rackId); err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: rackId})
}
