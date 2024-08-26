package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/features/admin/biz"
	"github.com/wangxin688/narvis/server/features/admin/schemas"
	"github.com/wangxin688/narvis/server/tools/errors"
	ts "github.com/wangxin688/narvis/server/tools/schemas"
)

// @Tags Admin
// @Summary Create Group
// @Description Create User Group
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param group body schemas.GroupCreate true "group"
// @Success 200 {object} ts.IDResponse
// @Failure 409 {object} ts.ErrorResponse
// @Router /admin/groups [post]
func createGroup(c *gin.Context) {
	var group schemas.GroupCreate
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	if err = c.ShouldBindJSON(&group); err != nil {
		return
	}
	newGroup, err := biz.NewGroupService().CreateGroup(&group)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IDResponse{ID: newGroup})
}

// @Tags Admin
// @Summary Get Group
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "group id"
// @Success 200 {object} schemas.Group
// @Router /admin/groups/{id} [get]
func getGroup(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	groupId := c.Param("id")
	group, err := biz.NewGroupService().GetGroupByID(groupId)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, group)
}

// @Tags Admin
// @Summary List Groups
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param object query schemas.GroupQuery false "query groups"
// @Success 200 {object} schemas.ListResponse{results=[]schemas.Group}
// @Router /admin/groups [get]
func listGroups(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var req schemas.GroupQuery
	if err = c.ShouldBindQuery(&req); err != nil {
		return
	}
	count, list, err := biz.NewGroupService().ListGroups(&req)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.ListResponse{
		Total:   count,
		Results: list,
	})
}

// @Tags Admin
// @Summary Update Group
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "group id"
// @Param group body schemas.GroupUpdate true "group"
// @Success 200 {object} ts.IDResponse
// @Router /admin/groups/{id} [patch]
func updateGroup(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	groupId := c.Param("id")
	var group schemas.GroupUpdate
	if err = c.ShouldBindJSON(&group); err != nil {
		return
	}
	err = biz.NewGroupService().UpdateGroup(groupId, &group)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IDResponse{ID: groupId})
}

// @Tags Admin
// @Summary Delete Group
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "group id"
// @Success 200 {object} ts.IDResponse
// @Router /admin/groups/{id} [delete]
func deleteGroup(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	groupId := c.Param("id")
	err = biz.NewGroupService().DeleteGroup(groupId)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IDResponse{ID: groupId})
}
