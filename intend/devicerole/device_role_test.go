package devicerole_test

import (
	"testing"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/intend/devicerole"
)

func TestDeviceRole(t *testing.T) {
	count, list := devicerole.GetListDeviceRole()
	result := lo.Filter(list, func(item devicerole.DeviceRole, index int) bool {
		return string(item.DeviceRole) == "Unknown"
	})
	t.Log(count, result)
}
