package middleware

import (
	"github.com/gin-gonic/gin"
	files "github.com/swaggo/files"
	swg "github.com/swaggo/gin-swagger"
	"github.com/wangxin688/narvis/server/docs"
)

func RegisterOpenAPI(e *gin.Engine) {

	docs.SwaggerInfo.BasePath = "/api/v1"
	e.GET("/swagger/*any", swg.WrapHandler(files.Handler))
}
