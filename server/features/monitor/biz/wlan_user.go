package metric_biz

import (
	"fmt"
	"time"

	"github.com/wangxin688/narvis/intend/intendtask"
	"github.com/wangxin688/narvis/server/features/monitor/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/infra"
	"github.com/wangxin688/narvis/server/pkg/vtm"
)

func GetWlanUserTrend(query *schemas.WlanUserTrendRequest) ([]*schemas.WlanUserTrend, error) {
	interval := vtm.CalculateInterval(query.StartedAt.Unix(), query.EndedAt.Unix(), query.DataPoints)
	results := make([]*schemas.WlanUserTrend, 0)
	orgId := global.OrganizationId.Get()
	selectString := fmt.Sprintf(
		`
		stationESSID,
		toStartOfInterval(time, INTERVAL %d second) as timestamp,
		uniq(stationMac) as value
		`, interval,
	)
	stmt := infra.ClickHouseDB.Table("wlan_station").Select(selectString).Where(
		"ts >= ? AND ts <= ? AND organizationId = ?", query.StartedAt, query.EndedAt, orgId,
	)
	if query.SiteId != nil && *query.SiteId != "" {
		stmt = stmt.Where("siteId = ?", *query.SiteId)
	}
	stmt = stmt.Group("stationESSID, timestamp").Order("stationESSID ASC, timestamp ASC")
	err := stmt.Scan(&results).Error
	if err != nil {
		return nil, err
	}
	// TODO：transform data schema to fit frontend
	return results, nil
}

func ListWlanUsers(query *schemas.WlanUserQuery) (int64, []*intendtask.WlanUserItem, error) {
	results := make([]*intendtask.WlanUserItem, 0)
	orgId := global.OrganizationId.Get()
	count := int64(0)
	stmt := infra.ClickHouseDB.Table("wlan_station").Where(
		"ts >= ? AND ts <= ? AND organizationId = ?", query.StartedAt, query.EndedAt, orgId,
	)
	if query.SiteId != nil && *query.SiteId != "" {
		stmt = stmt.Where("siteId = ?", *query.SiteId)
	}
	if query.StationMac != nil && *query.StationMac != "" {
		stmt = stmt.Where("stationMac = ?", *query.StationMac)
	}
	if query.StationESSID != nil && *query.StationESSID != "" {
		stmt = stmt.Where("stationESSID = ?", *query.StationESSID)
	}
	if query.ApName != nil && *query.ApName != "" {
		stmt = stmt.Where("apName = ?", *query.ApName)
	}
	if query.IsSearchable() {
		keyword := "%" + *query.Keyword + "%"
		stmt = stmt.Where("stationMac LIKE ? OR stationIp LIKE ? OR stationUsername LIKE", keyword, keyword, keyword)
	}
	err := stmt.Count(&count).Error
	if err != nil {
		return 0, nil, err
	}
	err = stmt.Scan(&results).Error
	if err != nil || count <= 0 {
		return 0, results, err
	}
	stmt.Scopes(query.OrderByField())
	stmt.Scopes(query.Pagination())
	err = stmt.Scan(&results).Error
	if err != nil {
		return 0, nil, err
	}
	return count, results, nil

}

func WlanUserDetail(macAddr string) {

}

func getWlanUserThroughput(macAddr string, startedAt, endedAt time.Time) {

}

func getWlanUserRSSI(macAddr string, startedAt, endedAt time.Time) {

}

func getWlanUserSNR(macAddr string, startedAt, endedAt time.Time) {

}

func getWlanUserLogs(macAddr string, startedAt, endedAt time.Time) {

}
