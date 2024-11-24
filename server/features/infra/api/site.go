package infra_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/intend/helpers/bgtask"
	"github.com/wangxin688/narvis/intend/logger"
	biz "github.com/wangxin688/narvis/server/features/infra/biz"
	"github.com/wangxin688/narvis/server/features/infra/hooks"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	"github.com/wangxin688/narvis/server/tools/errors"
	"github.com/wangxin688/narvis/server/tools/helpers"
	ts "github.com/wangxin688/narvis/server/tools/schemas"
	"go.uber.org/zap"
)

// @Tags Infra.Site
// @Summary Create new site
// @X-func {"name": "CreateSite"}
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
	bgtask.BackgroundTask(func() {
		_, err = hooks.SiteHookCreate(newSite)
		if err != nil {
			logger.Logger.Error("[siteCreateHooks]:create host group failed", zap.Error(err))
		}
	})
	c.JSON(http.StatusOK, ts.IdResponse{Id: newSite})
}

// @Tags Infra.Site
// @Summary Get site
// @X-func {"name": "GetSite"}
// @Description Get site
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted siteId"
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

// @Tags Infra.Site
// @Summary List sites
// @X-func {"name": "ListSites"}
// @Description List sites
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param object query schemas.SiteQuery false "query sites"
// @Success 200 {object} ts.ListResponse{results=[]schemas.SiteResponse}
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

// @Tags Infra.Site
// @Summary Update site
// @X-func {"name": "UpdateSite"}
// @Description Update site
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted siteId"
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
	diff, err := biz.NewSiteService().Update(siteId, &site)
	if err != nil {
		return
	}
	bgtask.BackgroundTask(func() {
		hooks.SiteHookUpdate(siteId, diff[siteId])
	})
	c.JSON(http.StatusOK, ts.IdResponse{Id: siteId})
}

// @Tags Infra.Site
// @Summary Delete site
// @X-func {"name": "DeleteSite"}
// @Description Delete site
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted siteId"
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
	site, err := biz.NewSiteService().Delete(siteId)
	if err != nil {
		return
	}
	bgtask.BackgroundTask(func() {
		hooks.SiteHookDelete(site)
	})
	c.JSON(http.StatusOK, ts.IdResponse{Id: siteId})
}
