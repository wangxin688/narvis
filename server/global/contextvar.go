package global

import (
	"github.com/timandy/routine"
)

var OrganizationID = routine.NewThreadLocal[string]()

var XRequestID = routine.NewThreadLocal[string]()

var UserID = routine.NewThreadLocal[string]()
