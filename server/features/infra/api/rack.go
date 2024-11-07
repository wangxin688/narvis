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

// @Tags Infra
// @Summary Create new rack
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

// @Tags Infra
// @Summary Get rack
// @Description Get rack
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param rackId path string true "rackId"
// @Success 200 {object} schemas.RackElevation
// @Router /infra/racks/{id} [get]
func getRack(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	rackId := c.Param("rackId")
	if err := helpers.ValidateUuidString(rackId); err != nil {
		return
	}
	rack, err := biz.NewRackService().GetRackElevation(rackId)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, rack)
}

// @Tags Infra
// @Summary List racks
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

// @Tags Infra
// @Summary Update rack
// @Description Update rack
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param rackId path string true "rackId"
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
	rackId := c.Param("rackId")
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

// @Tags Infra
// @Summary Delete rack
// @Description Delete rack
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param rackId path string true "rackId"
// @Success 200 {object} ts.IdResponse
// @Router /infra/racks/{id} [delete]
func deleteRack(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	rackId := c.Param("rackId")
	if err = helpers.ValidateUuidString(rackId); err != nil {
		return
	}
	if err = biz.NewRackService().DeleteRack(rackId); err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: rackId})
}
