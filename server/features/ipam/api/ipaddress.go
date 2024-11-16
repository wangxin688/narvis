package ipam_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	ipam_biz "github.com/wangxin688/narvis/server/features/ipam/biz"
	"github.com/wangxin688/narvis/server/features/ipam/schemas"
	"github.com/wangxin688/narvis/server/tools/errors"
	"github.com/wangxin688/narvis/server/tools/helpers"
	ts "github.com/wangxin688/narvis/server/tools/schemas"
)

// @Tags IPAM
// @Summary Create ip address
// @X-func {"name": "CreateIpAddress"}
// @Description Create ip address
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param ip body schemas.IpAddressCreate true "ip"
// @Success 200 {object} ts.IdResponse
// @Router /ipam/ip-addresses [post]
func createIpAddress(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var ip schemas.IpAddressCreate
	if err = c.ShouldBindJSON(&ip); err != nil {
		return
	}
	ipId, err := ipam_biz.NewIpAddressService().CreateIpAddress(&ip)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: ipId})
}

// @Tags IPAM
// @Summary Get ip address
// @X-func {"name": "GetIpAddress"}
// @Description Get ip address
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted ipAddressId"
// @Success 200 {object} schemas.IpAddress
// @Router /ipam/ip-addresses/{id} [get]
func getIpAddress(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	ipId := c.Param("id")
	if err = helpers.ValidateUuidString(ipId); err != nil {
		return
	}
	ip, err := ipam_biz.NewIpAddressService().GetById(ipId)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ip)
}

// @Tags IPAM
// @Summary Get ip address list
// @X-func {"name": "ListIpAddresses"}
// @Description Get ip address list
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param object query schemas.IpAddressQuery true "query ip"
// @Success 200 {object} ts.ListResponse{results=[]schemas.IpAddress}
// @Router /ipam/ip-addresses [get]
func getIpAddressList(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var query schemas.IpAddressQuery
	if err = c.ShouldBindQuery(&query); err != nil {
		return
	}
	count, ipList, err := ipam_biz.NewIpAddressService().ListIpAddresses(&query)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.ListResponse{Total: count, Results: ipList})
}

// @Tags IPAM
// @Summary Delete ip address
// @X-func {"name": "DeleteIpAddress"}
// @Description Delete ip address
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted ipAddressId"
// @Success 200 {object} ts.IdResponse
// @Router /ipam/ip-addresses/{id} [delete]
func deleteIpAddress(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	ipId := c.Param("id")
	if err = helpers.ValidateUuidString(ipId); err != nil {
		return
	}
	err = ipam_biz.NewIpAddressService().DeleteIpAddress(ipId)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: ipId})
}

// @Tags IPAM
// @Summary Update ip address
// @X-func {"name": "UpdateIpAddress"}
// @Description Update ip address
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted ipAddressId"
// @Param ip body schemas.IpAddressUpdate true "ip"
// @Success 200 {object} schemas.IpAddress
// @Router /ipam/ip-addresses/{id} [put]
func updateIpAddress(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	ipId := c.Param("id")
	if err = helpers.ValidateUuidString(ipId); err != nil {
		return
	}
	var ip schemas.IpAddressUpdate
	if err = c.ShouldBindJSON(&ip); err != nil {
		return
	}
	err = ipam_biz.NewIpAddressService().UpdateIpAddress(ipId, &ip)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ip)
}
