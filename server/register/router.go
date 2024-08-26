package register

import (
	"github.com/gin-gonic/gin"
	admin_api "github.com/wangxin688/narvis/server/features/admin/api"
	intend_api "github.com/wangxin688/narvis/server/features/intend/api"
	organization_api "github.com/wangxin688/narvis/server/features/organization/api"
)

func RegisterRouter(e *gin.Engine) {

	organization_api.RegisterOrgRoutes(e)
	intend_api.RegisterIntendRoutes(e)
	admin_api.RegisterAdminRoutes(e)
	admin_api.RegisterLoginRoutes(e)
}
