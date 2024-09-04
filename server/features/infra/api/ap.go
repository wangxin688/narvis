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

// @Tags Infra
// @Summary Get ap
// @Description get ap
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param id path int true "id"
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

// @Tags Infra
// @Summary List aps
// @Description List aps
// @Security BearerAuth
// @Accept  json
// @Produce  json
// @Param object query schemas.ApQuery true "query aps"
// @Success 200 {object} []ts.ListResponse{results=[]schemas.AP}
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
