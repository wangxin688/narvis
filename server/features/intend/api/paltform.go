package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wangxin688/narvis/server/features/intend/biz"
	its "github.com/wangxin688/narvis/server/features/intend/schemas"
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
// @Param object query schemas.PlatformQuery false "query platforms"
// @Success 200 {object} schemas.ListResponse{results=[]string}
// @Router /intend/platforms [get]
func platformList(c *gin.Context) {
	var err error
	defer func() {
		if err != nil {
			errors.ResponseErrorHandler(c, err)
		}
	}()
	var req its.PlatformQuery
	if err := c.ShouldBindQuery(&req); err != nil {
		return
	}
	count, list := biz.GetPlatforms(&req)

	c.JSON(http.StatusOK, schemas.ListResponse{
		Total:   count,
		Results: list,
	})
}
