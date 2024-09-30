package ipam_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	biz "github.com/wangxin688/narvis/server/features/ipam/biz"
	"github.com/wangxin688/narvis/server/features/ipam/schemas"
	"github.com/wangxin688/narvis/server/tools/errors"
	"github.com/wangxin688/narvis/server/tools/helpers"
	ts "github.com/wangxin688/narvis/server/tools/schemas"
)

// @Tags IPAM
// @Summary Create prefix
// @Description Create prefix
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param prefix body schemas.PrefixCreate true "Prefix"
// @Success 200 {object} ts.IdResponse
// @Router /ipam/prefixes [post]
func createPrefix(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var prefix schemas.PrefixCreate
	if err = c.ShouldBindJSON(&prefix); err != nil {
		return
	}
	prefixId, err := biz.NewPrefixService().CreatePrefix(&prefix)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: prefixId})
}

// @Tags IPAM
// @Summary Get prefix
// @Description Get prefix
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted prefixId"
// @Success 200 {object} schemas.Prefix
// @Router /ipam/prefixes/{id} [get]
func getPrefix(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	prefixId := c.Param("id")
	if err = helpers.ValidateUuidString(prefixId); err != nil {
		return
	}
	prefix, err := biz.NewPrefixService().GetById(prefixId)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, prefix)
}

// @Tags IPAM
// @Summary Get prefix list
// @Description Get prefix list
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param object query schemas.PrefixQuery true "query prefix"
// @Success 200 {object} ts.ListResponse{results=[]schemas.Prefix}
// @Router /ipam/prefixes [get]
func getPrefixList(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var query schemas.PrefixQuery
	if err = c.ShouldBindQuery(&query); err != nil {
		return
	}
	count, prefixList, err := biz.NewPrefixService().ListPrefix(&query)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.ListResponse{Total: count, Results: prefixList})
}

// @Tags IPAM
// @Summary Update prefix
// @Description Update prefix
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted prefixId"
// @Param prefix body schemas.PrefixUpdate true "Prefix"
// @Success 200 {object} ts.IdResponse
// @Router /ipam/prefixes/{id} [put]
func updatePrefix(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	prefixId := c.Param("id")
	if err = helpers.ValidateUuidString(prefixId); err != nil {
		return
	}
	var prefix schemas.PrefixUpdate
	if err = c.ShouldBindJSON(&prefix); err != nil {
		return
	}
	err = biz.NewPrefixService().UpdatePrefix(c, prefixId, &prefix)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: prefixId})
}

// @Tags IPAM
// @Summary Delete prefix
// @Description Delete prefix
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted prefixId"
// @Success 200 {object} ts.IdResponse
// @Router /ipam/prefixes/{id} [delete]
func deletePrefix(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	prefixId := c.Param("id")
	if err = helpers.ValidateUuidString(prefixId); err != nil {
		return
	}
	err = biz.NewPrefixService().DeletePrefix(prefixId)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: prefixId})
}
