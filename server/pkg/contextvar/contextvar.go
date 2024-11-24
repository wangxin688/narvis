package contextvar

import "github.com/timandy/routine"

var OrganizationId = routine.NewThreadLocal[string]()

var ProxyId = routine.NewThreadLocal[string]()

var XRequestId = routine.NewThreadLocal[string]()

var UserId = routine.NewThreadLocal[string]()

type Diff struct {
	Before any `json:"before"`
	After  any `json:"after"`
}

var OrmDiff = routine.NewThreadLocal[map[string]map[string]*Diff]()
