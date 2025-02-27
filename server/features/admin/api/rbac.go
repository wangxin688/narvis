package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/features/admin/biz"
	"github.com/wangxin688/narvis/server/features/admin/schemas"
	"github.com/wangxin688/narvis/server/tools/errors"
)

// @Tags Auth
// @Summary Username Password Login
// @X-func {"name": "PasswordLogin"}
// @param body body schemas.Oauth2PasswordRequest true "Username Password Login"
// @Success 200 {object} security.AccessToken
// @Router /login/password [post]
func passwordLogin(c *gin.Context) {
	var req schemas.Oauth2PasswordRequest
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	if err = c.ShouldBind(&req); err != nil {
		return
	}
	token, err := biz.NewRBACService().PasswordLogin(req)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, token)
}

// @Tags Auth
// @Summary Generate refresh token
// @X-func {"name": "CreateRefreshToken"}
// @param body body schemas.Oauth2RefreshRequest true "Generate refresh token"
// @Success 200 {object} security.AccessToken
// @Router /login/refresh [post]
func refreshToken(c *gin.Context) {
	var req schemas.Oauth2RefreshRequest
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	if err = c.ShouldBind(&req); err != nil {
		return
	}
	token, err := biz.NewRBACService().CreateRefreshToken(req.RefreshToken)
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, token)
}

// google auth https://console.cloud.google.com/apis/credentials/consent/edit;newAppInternalUser=false?hl=zh-cn&project=smart-seer-431515-a0
