package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/intend/model/platform"
	"github.com/wangxin688/narvis/server/tools/errors"
	"github.com/wangxin688/narvis/server/tools/schemas"
)

// @Tags Intend
// @Security BearerAuth
// @Summary Get list of Platforms
// @X-func {"name": "PlatformList"}
// @Description Get list of Platforms
// @Accept json
// @Produce json
// @Success 200 {object} schemas.ListResponse{results=[]string}
// @Router /intend/platforms [get]
func platformList(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	list := platform.SupportPlatform()
	c.JSON(http.StatusOK, schemas.ListResponse{
		Total:   int64(len(list)),
		Results: list,
	})
}
