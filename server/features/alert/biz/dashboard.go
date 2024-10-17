package alert_biz

import (
	"time"

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

func alertTrend(siteId *string, startedAtGte time.Time, startedAtLte time.Time) ([]*schemas.TrendItem, error) {
	trendStmt := gen.Alert.Select(
		gen.Alert.Id.Count().As("value"),
		gen.Alert.Severity.As("severity"),
		gen.Alert.StartedAt.Date().As("time"),
	).Where(
		gen.Alert.StartedAt.Between(startedAtGte, startedAtLte),
		gen.Alert.OrganizationId.Eq(global.OrganizationId.Get()),
	).Group(gen.Alert.Severity, gen.Alert.StartedAt.Date()).Order(gen.Alert.StartedAt.Desc())

	if siteId != nil && *siteId != "" {
		trendStmt = trendStmt.Where(gen.Alert.SiteId.Eq(*siteId))
	}

	var results []*schemas.TrendItem
	err := trendStmt.Scan(&results)
	return results, err
}
