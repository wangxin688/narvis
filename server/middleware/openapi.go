package middleware

import (
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	swg "github.com/swaggo/gin-swagger"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/docs"
)

func RegisterOpenAPI(e *gin.Engine) {

	docs.SwaggerInfo.BasePath = core.Settings.System.RouterPrefix
	e.GET("/swagger/*any", swg.WrapHandler(files.Handler))
}
