package biz

import (
	"strings"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/intend/manufacturer"
	"github.com/wangxin688/narvis/server/features/intend/schemas"
	"github.com/wangxin688/narvis/server/tools/helpers"
)

func GetManufacturers(query *schemas.ManufacturerQuery) (int64, []string) {
	list := manufacturer.SupportedManufacturer()
	count := len(list)
	listString := lo.Map(list, func(item manufacturer.Manufacturer, index int) string {
		return string(item)
	})
	if query == nil {
		return int64(count), listString
	}
	if query.Manufacturer != nil {
		listString = lo.Filter(listString, func(item string, index int) bool {
			return strings.EqualFold(item, *query.Manufacturer)
		})
	}

	if query.Keyword != nil {
		listString = helpers.FuzzySearchList(listString, *query.Keyword, true)
	}
	return int64(len(listString)), listString

}
