package infra_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/intend/helpers/bgtask"
	infra_biz "github.com/wangxin688/narvis/server/features/infra/biz"
	"github.com/wangxin688/narvis/server/features/infra/hooks"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	"github.com/wangxin688/narvis/server/tools/errors"
	"github.com/wangxin688/narvis/server/tools/helpers"
	ts "github.com/wangxin688/narvis/server/tools/schemas"
)

// @Tags Infra.Server
// @Summary Create new server
// @X-func {"name": "CreateServer"}
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
	bgtask.BackgroundTask(func() {
		hooks.ServerCreateHooks(newServer)
	})
	c.JSON(http.StatusOK, ts.IdResponse{Id: newServer})
}

// @Tags Infra.Server
// @Summary Update server
// @X-func {"name": "UpdateServer"}
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
	bgtask.BackgroundTask(func() {
		hooks.ServerUpdateHooks(serverId, diff[serverId])
	})
	c.JSON(http.StatusOK, ts.IdResponse{Id: serverId})
}

// @Tags Infra.Server
// @Summary Delete server
// @X-func {"name": "DeleteServer"}
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
	bgtask.BackgroundTask(func() {
		hooks.ServerDeleteHooks(server)
	})
	c.JSON(http.StatusOK, ts.IdResponse{Id: serverId})
}

// @Tags Infra.Server
// @Summary Get server
// @X-func {"name": "GetServer"}
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
// @X-func {"name": "ListServers"}
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

// @Tags Infra.Server
// @Summary Create Server new cli credential
// @X-func {"name": "CreateServerCliCredential"}
// @Description Create Server new cli credential
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted deviceId"
// @Param credential body schemas.CliCredentialCreate true "Credential"
// @Success 200 {object} ts.IdResponse
// @Router /infra/servers/{id}/cli [post]
func createCliCredentialForServer(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var credential schemas.CliCredentialCreate
	serverId := c.Param("id")
	if err = helpers.ValidateUuidString(serverId); err != nil {
		return
	}
	if err = c.ShouldBindJSON(&credential); err != nil {
		return
	}
	id, err := infra_biz.NewCliCredentialService().CreateServerCredential(serverId, &credential)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: id})
}

// @Tags Infra.Server
// @Summary Update Server cli credential
// @X-func {"name": "UpdateServerCliCredential"}
// @Description Update Server cli credential
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted deviceId"
// @Param credential body schemas.CliCredentialUpdate true "Credential"
// @Success 200 {object} ts.IdResponse
// @Router /infra/servers/{id}/cli [put]
func updateCliCredentialForServer(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	id := c.Param("id")
	if err = helpers.ValidateUuidString(id); err != nil {
		return
	}
	var credential schemas.CliCredentialUpdate
	if err = c.ShouldBindJSON(&credential); err != nil {
		return
	}
	_, err = infra_biz.NewCliCredentialService().UpdateServerCredential(id, &credential)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: id})
}

// @Tags Infra.Server
// @Summary Delete Server cli credential
// @X-func {"name": "DeleteServerCliCredential"}
// @Description Delete Server cli credential
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted deviceId"
// @Success 200 {object} ts.IdResponse
// @Router /infra/servers/{id}/cli [delete]
func deleteCliCredentialForServer(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	id := c.Param("id")
	if err = helpers.ValidateUuidString(id); err != nil {
		return
	}
	err = infra_biz.NewCliCredentialService().DeleteServerCredential(id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: id})
}

// @Tags Infra.Server
// @Summary Get Server cli credential
// @X-func {"name": "GetServerCliCredential"}
// @Description Get Server cli credential
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted deviceId"
// @Success 200 {object} schemas.CliCredential
// @Router /infra/servers/{id}/cli [get]
func getCliCredentialForServer(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	id := c.Param("id")
	if err = helpers.ValidateUuidString(id); err != nil {
		return
	}
	credential, err := infra_biz.NewCliCredentialService().GetCredentialByDeviceId(id)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, credential)
}

// @Tags Infra.Server
// @Summary Create server new snmpV2 credential
// @X-func {"name": "CreateServerSnmpV2Credential"}
// @Description Create server new snmpV2 credential
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted deviceId"
// @Param credential body schemas.SnmpV2CredentialCreate true "Credential"
// @Success 200 {object} ts.IdResponse
// @Router /infra/servers/{id}/snmpv2 [post]
func createSnmpV2CredentialForServer(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {

			errors.ResponseErrorHandler(c, err)
		}
	}()
	var credential schemas.SnmpV2CredentialCreate
	serverId := c.Param("id")
	if err = helpers.ValidateUuidString(serverId); err != nil {
		return
	}
	if err = c.ShouldBindJSON(&credential); err != nil {
		return
	}
	id, err := infra_biz.NewSnmpCredentialService().CreateServerSnmpCredential(serverId, &credential)
	if err != nil {
		return
	}
	hooks.ServerSnmpCredCreateHooks(id)
	c.JSON(http.StatusOK, ts.IdResponse{Id: id})
}

// @Tags Infra.Server
// @Summary Update server snmpV2 credential
// @X-func {"name": "UpdateServerSnmpV2Credential"}
// @Description Update server snmpV2 credential
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted deviceId"
// @Param credential body schemas.SnmpV2CredentialUpdate true "Credential"
// @Success 200 {object} ts.IdResponse
// @Router /infra/servers/{id}/snmpv2 [put]
func updateSnmpV2CredentialForServer(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	serverId := c.Param("id")
	if err = helpers.ValidateUuidString(serverId); err != nil {
		return
	}
	var credential schemas.SnmpV2CredentialUpdate
	if err = c.ShouldBindJSON(&credential); err != nil {
		return
	}
	credId, diff, err := infra_biz.NewSnmpCredentialService().UpdateServerSnmpCredential(serverId, &credential)
	if err != nil {
		return
	}
	bgtask.BackgroundTask(func() {
		hooks.ServerSnmpCredUpdateHooks(credId, diff[credId])
	})
	c.JSON(http.StatusOK, ts.IdResponse{Id: serverId})
}

// @Tags Infra.Server
// @Summary Delete server snmpV2 credential
// @X-func {"name": "DeleteServerSnmpV2Credential"}
// @Description Delete server snmpV2 credential
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted deviceId"
// @Success 200 {object} ts.IdResponse
// @Router /infra/servers/{id}/snmpv2 [delete]
func deleteSnmpV2CredentialForServer(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	serverId := c.Param("id")
	if err = helpers.ValidateUuidString(serverId); err != nil {
		return
	}
	cred, err := infra_biz.NewSnmpCredentialService().DeleteServerCredential(serverId)
	if err != nil {
		return
	}
	bgtask.BackgroundTask(func() {
		hooks.ServerSnmpCredDeleteHooks(cred)
	})
	c.JSON(http.StatusOK, ts.IdResponse{Id: serverId})
}

// @Tags Infra.Server
// @Summary Get server snmpV2 credential
// @X-func {"name": "GetServerSnmpV2Credential"}
// @Description Get server snmpV2 credential
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted deviceId"
// @Success 200 {object} schemas.SnmpV2Credential
// @Router /infra/servers/{id}/snmpv2 [get]
func getSnmpV2CredentialForServer(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	serverId := c.Param("id")
	if err = helpers.ValidateUuidString(serverId); err != nil {
		return
	}
	credential, err := infra_biz.NewSnmpCredentialService().GetCredentialByDeviceId(serverId)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, credential)
}
