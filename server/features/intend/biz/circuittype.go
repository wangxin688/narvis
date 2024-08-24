package biz

import (
	"strings"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/intend/circuittype"
	"github.com/wangxin688/narvis/server/features/intend/schemas"
	"github.com/wangxin688/narvis/server/tools/helpers"
)

func GetCircuitTypes(query *schemas.CircuitTypeQuery) (int64, []circuittype.CircuitType) {
	list := circuittype.GetListCircuitType()
	if query == nil {
		return int64(len(list)), list
	}
	if query.CircuitType != nil {
		list = lo.Filter(list, func(item circuittype.CircuitType, index int) bool {
			return strings.EqualFold(string(item.CircuitType), *query.CircuitType)
		})
	}

	if query.ConnectionType != nil {
		list = lo.Filter(list, func(item circuittype.CircuitType, index int) bool {
			return strings.EqualFold(string(item.ConnectionType), *query.ConnectionType)
		})
	}
	if query.Description != nil {
		list = lo.Filter(list, func(item circuittype.CircuitType, index int) bool {
			return helpers.FuzzySearch(item.ToMap(), *query.Description, true, nil)
		})
	}
	if query.Search != nil {
		list = lo.Filter(list, func(item circuittype.CircuitType, index int) bool {
			return helpers.FuzzySearch(item.ToMap(), *query.Search, true, nil)
		})
	}

	return int64(len(list)), list
}
