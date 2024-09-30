package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/features/admin/biz"
	"github.com/wangxin688/narvis/server/features/admin/schemas"
	"github.com/wangxin688/narvis/server/tools/errors"
	"github.com/wangxin688/narvis/server/tools/helpers"
	ts "github.com/wangxin688/narvis/server/tools/schemas"
)

// @Tags Admin
// @Summary Create role
// @Description create role
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param role body schemas.RoleCreate true "role"
// @Success 200 {object} schemas.Role
// @Router /admin/roles [post]
func createRole(c *gin.Context) {
	var role schemas.RoleCreate
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	if err = c.ShouldBindJSON(&role); err != nil {
		return
	}

	newRole, err := biz.NewRoleService().CreateRole(&role)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, newRole)
}

// @Tags Admin
// @Summary Get role
// @Description get role
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted role id"
// @Success 200 {object} schemas.RoleDetail
// @Router /admin/roles/{id} [get]
func getRole(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	roleId := c.Param("id")
	if err := helpers.ValidateUuidString(roleId); err != nil {
		return
	}
	role, err := biz.NewRoleService().GetRoleById(roleId)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, role)
}

// @Tags Admin
// @Summary List roles
// @Description list roles
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param object query schemas.RoleQuery false "query roles"
// @Success 200 {object} schemas.ListResponse{results=[]schemas.Role}
// @Router /admin/roles [get]
func listRoles(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var req schemas.RoleQuery
	if err = c.ShouldBindQuery(&req); err != nil {
		return
	}
	count, list, err := biz.NewRoleService().ListRoles(&req)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.ListResponse{
		Total:   count,
		Results: list,
	})
}

// @Tags Admin
// @Summary Update role
// @Description update role
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted role id"
// @Param role body schemas.RoleUpdate true "role"
// @Success 200 {object} schemas.Role
// @Router /admin/roles/{id} [put]
func updateRole(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	roleId := c.Param("id")
	if err := helpers.ValidateUuidString(roleId); err != nil {
		return
	}
	var role schemas.RoleUpdate
	if err = c.ShouldBindJSON(&role); err != nil {
		return
	}
	err = biz.NewRoleService().UpdateRole(roleId, &role)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: roleId})
}

// @Tags Admin
// @Summary Delete role
// @Description delete role
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "uuid formatted role id"
// @Success 200 {object} schemas.Role
// @Router /admin/roles/{id} [delete]
func deleteRole(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	roleId := c.Param("id")
	if err := helpers.ValidateUuidString(roleId); err != nil {
		return
	}
	err = biz.NewRoleService().DeleteRole(roleId)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: roleId})

}
