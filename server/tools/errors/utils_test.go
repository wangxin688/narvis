package errors

import (
	"testing"
)

func TestRemoveOrgString(t *testing.T) {
	fields := "name, organizationId"
	values := "string, 215e09ee-10d0-43af-953e-4faca45d57d2"

	fields, values = removeOrgInError(fields, values)
	if fields != "name" || values != "string" {
		t.Error(fields, values)
	}
}
