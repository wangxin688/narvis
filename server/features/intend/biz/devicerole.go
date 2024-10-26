package biz

import (
	"strings"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/intend/devicerole"
	"github.com/wangxin688/narvis/server/features/intend/schemas"
	"github.com/wangxin688/narvis/server/tools/helpers"
)

func GetDeviceRoles(query *schemas.DeviceRoleQuery) (int64, []devicerole.DeviceRole) {
	count, list := devicerole.GetListDeviceRole()
	if query == nil {
		return int64(count), list
	}
	if query.Name != nil {
		list = lo.Filter(list, func(item devicerole.DeviceRole, _ int) bool {
			return strings.EqualFold(string(item.DeviceRole), *query.Name)
		})
	}
	if query.Keyword != nil {
		list = lo.Filter(list, func(item devicerole.DeviceRole, _ int) bool {
			return helpers.FuzzySearch(item.ToMap(), *query.Keyword, true, nil)
		})
	}

	return int64(len(list)), list
}
