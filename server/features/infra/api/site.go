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
// @Summary Create new site
// @Description Create new site
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param site body schemas.SiteCreate true "site"
// @Success 200 {object} ts.IdResponse
// @Router /infra/sites [post]
func createSite(c *gin.Context) {
	var site schemas.SiteCreate
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	if err = c.ShouldBindJSON(&site); err != nil {
		return
	}
	newSite, err := biz.NewSiteService().Create(site)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: newSite})
}

// @Tags Infra
// @Summary Get site
// @Description Get site
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "site id"
// @Success 200 {object} schemas.SiteDetail
// @Router /infra/sites/{id} [get]
func getSite(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	siteId := c.Param("id")
	if err := helpers.ValidateUuidString(siteId); err != nil {
		return
	}
	site, err := biz.NewSiteService().GetSiteDetail(siteId)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, site)
}

// @Tags Infra
// @Summary List sites
// @Description List sites
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param object query schemas.SiteQuery false "query sites"
// @Success 200 {object} schemas.ListResponse{results=[]schemas.Site}
// @Router /infra/sites [get]
func listSites(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var req schemas.SiteQuery
	if err = c.ShouldBindQuery(&req); err != nil {
		return
	}
	count, list, err := biz.NewSiteService().GetList(&req)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.ListResponse{
		Total:   count,
		Results: list,
	})
}

// @Tags Infra
// @Summary Update site
// @Description Update site
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "site id"
// @Param site body schemas.SiteUpdate true "site"
// @Success 200 {object} ts.IdResponse
// @Router /infra/sites/{id} [put]
func updateSite(c *gin.Context) {
	var site schemas.SiteUpdate
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	siteId := c.Param("id")
	if err := helpers.ValidateUuidString(siteId); err != nil {
		return
	}
	if err = c.ShouldBindJSON(&site); err != nil {
		return
	}
	err = biz.NewSiteService().Update(siteId, &site)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: siteId})
}

// @Tags Infra
// @Summary Delete site
// @Description Delete site
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "site id"
// @Success 200 {object} ts.IdResponse
// @Router /infra/sites/{id} [delete]
func deleteSite(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	siteId := c.Param("id")
	if err := helpers.ValidateUuidString(siteId); err != nil {
		return
	}
	err = biz.NewSiteService().Delete(siteId)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: siteId})
}
