package global

import (
	"github.com/timandy/routine"
)

var OrganizationId = routine.NewThreadLocal[string]()

var XRequestId = routine.NewThreadLocal[string]()

var UserId = routine.NewThreadLocal[string]()
