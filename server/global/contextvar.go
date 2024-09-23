package global

import (
	"github.com/timandy/routine"
	"github.com/wangxin688/narvis/server/tools/schemas"
)

var OrganizationId = routine.NewThreadLocal[string]()

var ProxyId = routine.NewThreadLocal[string]()

var XRequestId = routine.NewThreadLocal[string]()

var UserId = routine.NewThreadLocal[string]()

var OrmDiff = routine.NewThreadLocal[map[string]map[string]*schemas.OrmDiff]()
