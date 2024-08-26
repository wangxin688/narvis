package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/features/organization/biz"
	"github.com/wangxin688/narvis/server/features/organization/schemas"
	"github.com/wangxin688/narvis/server/tools/errors"
)

// @Tags Organization
// @Security BearerAuth
// @Summary Create organization
// @Param body body schemas.Organization true "Create organization"
// @Success 200 {object} schemas.Organization
// @Router /org/organizations [post]
func orgCreate(c *gin.Context) {
	var org schemas.OrganizationCreate
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	if err = c.ShouldBindJSON(&org); err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest, gin.H{
				"error": err.Error(),
			},
		)
	}
	newOrg, err := biz.NewOrganizationService().CreateOrganization(&org)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, newOrg)
}
