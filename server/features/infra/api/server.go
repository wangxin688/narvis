package infra_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	infra_biz "github.com/wangxin688/narvis/server/features/infra/biz"
	"github.com/wangxin688/narvis/server/features/infra/hooks"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	"github.com/wangxin688/narvis/server/tools"
	"github.com/wangxin688/narvis/server/tools/errors"
	"github.com/wangxin688/narvis/server/tools/helpers"
	ts "github.com/wangxin688/narvis/server/tools/schemas"
)

// @Tags Infra.Server
// @Summary Create new server
// @Description Create new server
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param server body schemas.ServerCreate true "server"
// @Success 200 {object} ts.IdResponse
// @Router /infra/servers [post]
func createServer(c *gin.Context) {
	var server schemas.ServerCreate
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	if err = c.ShouldBindJSON(&server); err != nil {
		return
	}
	newServer, err := infra_biz.NewServerService().CreateServer(&server)
	if err != nil {
		return
	}
	tools.BackgroundTask(func() {
		hooks.ServerCreateHooks(newServer)
	})
	c.JSON(http.StatusOK, ts.IdResponse{Id: newServer})
}

// @Tags Infra.Server
// @Summary Update server
// @Description Update server
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted serverId"
// @Param server body schemas.ServerUpdate true "server"
// @Success 200 {object} ts.IdResponse
// @Router /infra/servers/{id} [put]
func updateServer(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	serverId := c.Param("id")
	if err := helpers.ValidateUuidString(serverId); err != nil {
		return
	}
	var server schemas.ServerUpdate
	if err = c.ShouldBindJSON(&server); err != nil {
		return
	}
	diff, err := infra_biz.NewServerService().UpdateServer(c, serverId, &server)
	if err != nil {
		return
	}
	tools.BackgroundTask(func() {
		hooks.ServerUpdateHooks(serverId, diff[serverId])
	})
	c.JSON(http.StatusOK, ts.IdResponse{Id: serverId})
}

// @Tags Infra.Server
// @Summary Delete server
// @Description Delete server
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted serverId"
// @Success 200 {object} ts.IdResponse
// @Router /infra/servers/{id} [delete]
func deleteServer(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	serverId := c.Param("id")
	if err := helpers.ValidateUuidString(serverId); err != nil {
		return
	}
	server, err := infra_biz.NewServerService().DeleteServer(serverId)
	if err != nil {
		return
	}
	tools.BackgroundTask(func() {
		hooks.ServerDeleteHooks(server)
	})
	c.JSON(http.StatusOK, ts.IdResponse{Id: serverId})
}

// @Tags Infra.Server
// @Summary Get server
// @Description Get server
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted serverId"
// @Success 200 {object} schemas.Server
// @Router /infra/servers/{id} [get]
func getServer(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	serverId := c.Param("id")
	if err := helpers.ValidateUuidString(serverId); err != nil {
		return
	}
	server, err := infra_biz.NewServerService().GetById(serverId)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, server)
}

// @Tags Infra.Server
// @Summary Get servers
// @Description Get servers
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param query query schemas.ServerQuery true "query"
// @Success 200 {object} ts.ListResponse{results=[]schemas.Server}
// @Router /infra/servers [get]
func listServers(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var query schemas.ServerQuery
	if err = c.ShouldBindQuery(&query); err != nil {
		return
	}
	count, servers, err := infra_biz.NewServerService().GetServerList(&query)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.ListResponse{Total: count, Results: servers})
}
