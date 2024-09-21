package biz

import (
	"strings"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/intend/platform"
	"github.com/wangxin688/narvis/server/features/intend/schemas"
	"github.com/wangxin688/narvis/server/tools/helpers"
)

func GetPlatforms(query *schemas.PlatformQuery) (int64, []string) {
	list := platform.SupportedPlatform()
	count := len(list)
	listString := lo.Map(list, func(item platform.Platform, _ int) string {
		return string(item)
	})
	if query == nil {
		return int64(count), listString
	}
	if query.Platform != nil {
		listString = lo.Filter(listString, func(item string, _ int) bool {
			return strings.EqualFold(item, *query.Platform)
		})
	}

	if query.Keyword != nil {
		listString = helpers.FuzzySearchList(listString, *query.Keyword, true)
	}
	return int64(len(listString)), listString
}
