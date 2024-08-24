package api

import (
	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/features/organization/schemas"
)

// @Tags Organization
// @Security ApiKeyAuth
// @Summary Create organization
// @Param body body schemas.Organization true "Create organization"
// @Success 200 {object} schemas.Organization
// @Router /api/organizations [post]
func OrgCreate(c *gin.Context) {
	var org schemas.OrganizationCreate

	def func() {
		if err != nil {
	}

	if err := c.ShouldBindJSON(&org); err != nil {
}
