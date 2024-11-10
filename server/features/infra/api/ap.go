package infra_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	infra_biz "github.com/wangxin688/narvis/server/features/infra/biz"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	"github.com/wangxin688/narvis/server/tools/errors"
	"github.com/wangxin688/narvis/server/tools/helpers"
	ts "github.com/wangxin688/narvis/server/tools/schemas"
)

// @Tags Infra.AP
// @Summary Get ap
// @Description get ap
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param id path string true "uuid formatted ap id"
// @Success 200 {object} schemas.AP
// @Router /infra/aps/{id} [get]
func getAp(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()

	apId := c.Param("id")

	if err = helpers.ValidateUuidString(apId); err != nil {
		return
	}
	ap, err := infra_biz.NewApService().GetById(apId)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ap)

}

// @Tags Infra.AP
// @Summary List aps
// @Description List aps
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param object query schemas.ApQuery true "query aps"
// @Success 200 {object} ts.ListResponse{results=[]schemas.AP}
// @Router /infra/aps [get]
func listAp(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var query schemas.ApQuery
	if err = c.ShouldBindQuery(&query); err != nil {
		return
	}
	count, aps, err := infra_biz.NewApService().GetApList(&query)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.ListResponse{Total: count, Results: aps})
}

// @Tags Infra.AP
// @Summary Update ap
// @Description Update ap
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted ap id"
// @Param ap body schemas.ApUpdate true "ap"
// @Success 200 {object} ts.IdResponse
// @Router /infra/aps/{id} [put]
func updateAp(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	apId := c.Param("id")
	if err = helpers.ValidateUuidString(apId); err != nil {
		return
	}
	var ap schemas.ApUpdate
	if err = c.ShouldBindJSON(&ap); err != nil {
		return
	}
	err = infra_biz.NewApService().UpdateApById(apId, &ap)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: apId})
}

// @Tags Infra.AP
// @Summary Batch Update ap
// @Description Batch Update ap
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param aps body []schemas.ApBatchUpdate true "aps"
// @Success 200 {object} ts.IdsResponse
// @Router /infra/aps [put]
func batchUpdateAp(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var ap schemas.ApBatchUpdate
	if err = c.ShouldBindJSON(&ap); err != nil {
		return
	}
	ids, err := infra_biz.NewApService().BatchUpdateAp(&ap)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdsResponse{Ids: ids})
}

// @Tags Infra.AP
// @Summary Batch Delete ap
// @Description Batch Delete ap
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param ids body []string true "ids"
// @Success 200 {object} ts.IdsResponse
// @Router /infra/aps [delete]
func batchDeleteAp(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var ids []string
	if err = c.ShouldBindJSON(&ids); err != nil {
		return
	}
	err = infra_biz.NewApService().DeleteApByIds(ids)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdsResponse{Ids: ids})
}

// @Tags Infra.AP
// @Summary Delete ap
// @Description Delete ap
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted ap id"
// @Success 200 {object} ts.IdResponse
// @Router /infra/aps/{id} [delete]
func deleteAp(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	apId := c.Param("id")
	if err = helpers.ValidateUuidString(apId); err != nil {
		return
	}
	err = infra_biz.NewApService().DeleteApById(apId)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: apId})
}
