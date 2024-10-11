package alert_biz

import (
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/alert/schemas"
	"github.com/wangxin688/narvis/server/global"
)

func topAlerts(siteId *string) ([]*schemas.TopX, error) {
	var results []*schemas.TopX

	topAlertStmt := gen.Alert.Select(
		gen.Alert.Id.Count().As("Value"),
		gen.Alert.AlertName.As("Name"),
	).Where(gen.Alert.OrganizationId.Eq(global.OrganizationId.Get()))
	if siteId != nil && *siteId != "" {
		topAlertStmt = topAlertStmt.Where(gen.Alert.SiteId.Eq(*siteId))
	}
	topAlertStmt = topAlertStmt.Group(gen.Alert.AlertName).Order(gen.Alert.Id.Count().Desc())

	err := topAlertStmt.Scan(&results)

	return results, err
}
