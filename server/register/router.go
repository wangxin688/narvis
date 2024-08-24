package register

import (
	"github.com/gin-gonic/gin"
	intend_api "github.com/wangxin688/narvis/server/features/intend/api"
	organization_api "github.com/wangxin688/narvis/server/features/organization/api"
)

func RegisterRouter(e *gin.Engine) {

	organization_api.RegisterOrgRoutes(e)
	intend_api.RegisterIntendRoutes(e)
}
