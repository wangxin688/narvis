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
// @Summary Create new user
// @Description Create new user
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param user body schemas.UserCreate true "user"
// @Success 200 {object} ts.IdResponse
// @Router /admin/users [post]
func createUser(c *gin.Context) {
	var user schemas.UserCreate
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	if err = c.ShouldBindJSON(&user); err != nil {
		return
	}
	newUser, err := biz.NewUserService().CreateUser(&user)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: newUser.Id})
}

// @Tags Admin
// @Summary Get user
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Success 200 {object} schemas.User
// @Router /admin/users/{id} [get]
func getUser(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	userId := c.Param("id")
	user, err := biz.NewUserService().GetUserById(userId)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, user)
}

// @Tags Admin
// @Summary List users
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param object query schemas.UserQuery false "query users"
// @Success 200 {object} schemas.ListResponse{results=[]schemas.User}
// @Router /admin/users [get]
func listUsers(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var req schemas.UserQuery
	if err = c.ShouldBindQuery(&req); err != nil {
		return
	}
	count, list, err := biz.NewUserService().ListUsers(&req)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.ListResponse{
		Total:   count,
		Results: list,
	})
}

// @Tags Admin
// @Summary Update user
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Param user body schemas.UserUpdate true "user"
// @Success 200 {object} ts.IdResponse
// @Router /admin/users/{id} [patch]
func updateUser(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	userId := c.Param("id")
	var user schemas.UserUpdate
	if err = c.ShouldBindJSON(&user); err != nil {
		return
	}
	err = biz.NewUserService().UpdateUser(userId, &user)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: userId})
}

// @Tags Admin
// @Summary Delete user
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path string true "user id"
// @Success 200 {object} ts.IdResponse
// @Router /admin/users/{id} [delete]
func deleteUser(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	userId := c.Param("id")
	err = biz.NewUserService().DeleteUser(c, userId)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, ts.IdResponse{Id: userId})
}
