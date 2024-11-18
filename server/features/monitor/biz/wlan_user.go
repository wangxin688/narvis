package metric_biz

import (
	"fmt"
	"time"

	"github.com/wangxin688/narvis/server/features/monitor/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/infra"
	"github.com/wangxin688/narvis/server/pkg/vtm"
)

type WlanUserService struct{}

func NewWlanUserService() *WlanUserService {
	return &WlanUserService{}
}

func (s *WlanUserService) GetWlanUserTrend(query *schemas.WlanUserTrendRequest) ([]*schemas.WlanUserTrend, error) {
	interval := vtm.CalculateInterval(query.StartedAt.Unix(), query.EndedAt.Unix(), query.DataPoints)
	results := make([]*schemas.WlanUserTrend, 0)
	orgId := global.OrganizationId.Get()
	selectString := fmt.Sprintf(
		`
		stationESSID,
		toStartOfInterval(ts, INTERVAL %d second) as timestamp,
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

func (s *WlanUserService) ListWlanUsers(query *schemas.WlanUserQuery) (*schemas.WlanUserListResponse, error) {
	response := &schemas.WlanUserListResponse{}
	results := make([]*schemas.WlanUserItem, 0)
	orgId := global.OrganizationId.Get()
	var countUser struct {
		OnlineCount int64 `json:"onlineCount"`
		TotalCount  int64 `json:"totalCount"`
	}
	subQuery := ""
	if query.StationMac != nil && *query.StationMac != "" {
		subQuery = fmt.Sprintf("stationMac = '%s'", *query.StationMac)
	} else if query.StationIp != nil && *query.StationIp != "" {
		subQuery = fmt.Sprintf("stationIp = '%s'", *query.StationIp)
	} else if query.StationUsername != nil && *query.StationUsername != "" {
		subQuery = fmt.Sprintf("stationUsername = '%s'", *query.StationUsername)
	} else {
		subQuery = "1=1"
	}
	if query.ApName != nil && *query.ApName != "" {
		subQuery += fmt.Sprintf(" AND stationApName = '%s'", *query.ApName)
	}
	if query.StationESSID != nil && *query.StationESSID != "" {
		subQuery += fmt.Sprintf(" AND stationESSID = '%s'", *query.StationESSID)
	}

	if query.PageInfo.Page == nil || *query.PageInfo.Page == 0 {
		query.PageInfo.Page = new(int)
		*query.PageInfo.Page = 1
	}
	if query.PageInfo.PageSize == nil || *query.PageInfo.PageSize == 0 {
		query.PageInfo.PageSize = new(int)
		*query.PageInfo.PageSize = 10
	}
	rawSql := fmt.Sprintf(
		`
		SELECT
			*
		FROM (
			SELECT
				toUnixTimestamp(ts) AS lastSeenAt,
				if ((toUnit32(endedAt)-toUint32(lastSeenAt))<500, TRUE, FALSE)) AS isActive,
				stationMac,
				stationIp,
				stationUsername,
				stationApMac,
				stationApName,
				stationESSID,
				stationVlan,
				stationChannel,
				stationChanBandWidth,
				stationRadioType,
				stationSNR,
				stationRSSI,
				stationRxBits,
				stationTxBits,
				stationMaxSpeed,
				stationOnlineTime
			FROM
				wlan_station
			WHERE
				organizationId = '%s'
				AND ts >= '%s'
				AND ts <= '%s'
				AND siteId = '%s'
			ORDER BY stationMac, ts DESC
			LIMIT 1 BY stationMac
		)
		WHERE %s LIMIT %d OFFSET %d
		`, orgId, query.StartedAt, query.EndedAt, query.SiteId, subQuery, (*query.PageInfo.Page-1)**query.PageInfo.PageSize, query.PageInfo.PageSize,
	)
	countRawSql := fmt.Sprintf(
		`
		SELECT
			COUNT(*) AS totalCount,
			SUM(CASE WHEN isActive = TRUE THEN 1 ELSE 0 END) AS onlineCount
		FROM (
			SELECT
				toUnixTimestamp(ts) AS lastSeenAt,
				if ((toUnit32(endedAt)-toUint32(lastSeenAt))<500, TRUE, FALSE)) AS isActive,
				stationMac,
				stationIp,
				stationUsername,
				stationApMac,
				stationApName,
				stationESSID,
				stationVlan,
				stationChannel,
				stationChanBandWidth,
				stationRadioType,
				stationSNR,
				stationRSSI,
				stationRxBits,
				stationTxBits,
				stationMaxSpeed,
				stationOnlineTime
			FROM
				wlan_station
			WHERE
				organizationId = '%s'
				AND ts >= '%s'
				AND ts <= '%s'
				AND siteId = '%s'
			ORDER BY stationMac, ts DESC
			LIMIT 1 BY stationMac
		)
		WHERE %s 
		`, orgId, query.StartedAt, query.EndedAt, query.SiteId, subQuery,
	)

	err := infra.ClickHouseDB.Raw(countRawSql).Scan(&countUser).Error
	if err != nil {
		return response, nil
	}
	if countUser.TotalCount == 0 {
		return response, nil
	}
	err = infra.ClickHouseDB.Raw(rawSql).Scan(&results).Error
	if err != nil {
		return response, nil
	}
	err = infra.ClickHouseDB.Raw(rawSql).Scan(&results).Error
	if err != nil {
		return response, nil
	}
	response.Total = countUser.TotalCount
	response.Online = countUser.OnlineCount
	response.Offline = countUser.TotalCount - countUser.OnlineCount
	response.Users = results
	return response, nil
	// TODO：confirm throughput calculate
}

func (s *WlanUserService) WlanUserDetail(macAddr string) {
	
}

func (s *WlanUserService) getWlanUserThroughput(macAddr string, startedAt, endedAt time.Time) {

}

func (s *WlanUserService) getWlanUserRSSI(macAddr string, startedAt, endedAt time.Time) {

}

func (s *WlanUserService) getWlanUserSNR(macAddr string, startedAt, endedAt time.Time) {

}

func (s *WlanUserService) getWlanUserLogs(macAddr string, startedAt, endedAt time.Time) {

}
