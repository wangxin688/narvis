package fixtures_test

import (
	"fmt"
	"testing"

	"github.com/wangxin688/narvis/server/tests/fixtures"
)

func TestRandomPrefix(t *testing.T) {
	a,b := fixtures.GenerateRFC1918Prefix()
	fmt.Println(a, b)
}
