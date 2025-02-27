package register

import (
	"github.com/gin-gonic/gin"
	admin_api "github.com/wangxin688/narvis/server/features/admin/api"
	alert_api "github.com/wangxin688/narvis/server/features/alert/api"
	device360_api "github.com/wangxin688/narvis/server/features/device360/api"
	infra_api "github.com/wangxin688/narvis/server/features/infra/api"
	intend_api "github.com/wangxin688/narvis/server/features/intend/api"
	ipam_api "github.com/wangxin688/narvis/server/features/ipam/api"
	monitor_api "github.com/wangxin688/narvis/server/features/monitor/api"
	organization_api "github.com/wangxin688/narvis/server/features/organization/api"
	task_api "github.com/wangxin688/narvis/server/features/task/api"
	webssh_api "github.com/wangxin688/narvis/server/features/webssh/api"
)

func RegisterRouter(e *gin.Engine) {

	organization_api.RegisterOrgRoutes(e)
	intend_api.RegisterIntendRoutes(e)
	admin_api.RegisterAdminRoutes(e)
	admin_api.RegisterLoginRoutes(e)
	infra_api.RegisterInfraRoutes(e)
	alert_api.RegisterAlertRoutes(e)
	task_api.RegisterTaskRoutes(e)
	ipam_api.RegisterIpamRoutes(e)
	webssh_api.RegisterWebSSHRoutes(e)
	monitor_api.RegisterMonitorRoutes(e)
	device360_api.RegisterDevice360Routes(e)
}
