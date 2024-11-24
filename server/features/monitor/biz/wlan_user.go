package metric_biz

import (
	"fmt"
	"strings"
	"time"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/intend/logger"
	"github.com/wangxin688/narvis/server/features/monitor/schemas"
	"github.com/wangxin688/narvis/server/infra"
	"github.com/wangxin688/narvis/server/pkg/contextvar"
	"github.com/wangxin688/narvis/server/pkg/vtm"
	"go.uber.org/zap"
)

type WlanUserService struct{}

func NewWlanUserService() *WlanUserService {
	return &WlanUserService{}
}

func (s *WlanUserService) GetWlanUserTrend(query *schemas.WlanUserTrendRequest) ([]*schemas.WlanUserTrend, error) {
	interval := vtm.CalculateInterval(query.StartedAt.Unix(), query.EndedAt.Unix(), query.DataPoints)
	results := make([]*schemas.WlanUserTrend, 0)
	orgId := contextvar.OrganizationId.Get()
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
	orgId := contextvar.OrganizationId.Get()
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

	macAddrs := lo.Map(results, func(item *schemas.WlanUserItem, index int) string {
		return item.StationMac
	})
	throughput := s.getWlanUserThroughput(macAddrs, query.StartedAt, query.EndedAt)
	if throughput != nil {
		for _, item := range results {
			if throughput, ok := throughput[item.StationMac]; ok {
				response.Users = append(response.Users, &schemas.WlanUserListResult{
					WlanUserItem: *item,
					RxBits:       throughput.RxBits,
					TxBits:       throughput.TxBits,
					TotalBits:    throughput.TotalBits,
					AvgSpeed:     throughput.AvgSpeed,
				})
			}
		}
	}
	return response, nil
}

func (s *WlanUserService) getWlanUserThroughput(macAddr []string, startedAt, endedAt time.Time) map[string]*schemas.WlanUserThroughput {
	macArrayString := strings.Join(macAddr, "','")
	dbResults := make([]*schemas.WlanUserThroughput, 0)
	rawSql := fmt.Sprintf(
		`
			WITH 
				toUnixTimestamp('%s') - toUnixTimestamp('%s') AS time_interval
			SELECT
				stationMac,
				maxRx - minRx AS rxBits,
				maxTx - minTx AS txBits,
			FROM (
				SELECT 
					stationMac,
					-- 首尾值
					anyIf(stationRxBits, ts = min_time) AS minRx,
					anyIf(stationRxBits, ts = max_time) AS maxRx,
					anyIf(stationTxBits, ts = min_time) AS minTx,
					anyIf(stationTxBits, ts = max_time) AS maxTx
				FROM (
					SELECT 
						stationMac,
						ts,
						stationRxBits,
						stationTxBits,
						-- 时间区间首尾
						min(ts) OVER (PARTITION BY stationMac) AS min_time,
						max(ts) OVER (PARTITION BY stationMac) AS max_time
					FROM station_data
					WHERE ts BETWEEN '%s' AND '%s' AND stationMac IN ('%s')
				)
				GROUP BY stationMac
			)
			ORDER BY stationMac;
		`, startedAt, endedAt, startedAt, endedAt, macArrayString,
	)

	err := infra.ClickHouseDB.Raw(rawSql).Scan(&dbResults).Error
	if err != nil {
		logger.Logger.Warn("[getWlanUserThroughput]: get wlan user throughput failed", zap.Error(err))
		return nil
	}

	results := make(map[string]*schemas.WlanUserThroughput)
	timeDelta := endedAt.Sub(startedAt).Seconds()
	for _, result := range dbResults {
		result.TotalBits = result.RxBits + result.TxBits
		result.AvgSpeed = uint64(float64(result.TotalBits) / timeDelta)
		results[result.StationMac] = result
	}
	return results

}

// func (s *WlanUserService) WlanUserDetail(macAddr string, startedAt, endedAt time.Time) {
// 	result := &ckmodel.WlanStation{}
// 	err := infra.ClickHouseDB.Model(&ckmodel.WlanStation{}).
// 		Where("stationMac = ?", macAddr).Order("ts DESC").Limit(1).First(result).Error

// }

func (s *WlanUserService) getWlanUserThroughputSerial(macAddr string, startedAt, endedAt time.Time) string {
	rawsql := fmt.Sprintf(
		`
		WITH
			-- 计算前一条记录的值和时间戳
			lagInFrame(stationRxBits) OVER (PARTITION BY stationMac ORDER BY timestamp) AS prev_rx,
			lagInFrame(stationTxBits) OVER (PARTITION BY stationMac ORDER BY timestamp) AS prev_tx,
			lagInFrame(timestamp) OVER (PARTITION BY stationMac ORDER BY timestamp) AS prev_time
		SELECT
			stationMac,
			timestamp AS current_time,
			prev_time AS previous_time,
			stationRxBits - prev_rx AS rx_diff,
			stationTxBits - prev_tx AS tx_diff,
			toUnixTimestamp(timestamp) - toUnixTimestamp(prev_time) AS time_diff_seconds,
			(stationRxBits - prev_rx) / (toUnixTimestamp(timestamp) - toUnixTimestamp(prev_time)) AS rx_trend_per_second,
			(stationTxBits - prev_tx) / (toUnixTimestamp(timestamp) - toUnixTimestamp(prev_time)) AS tx_trend_per_second
		FROM station_data
		WHERE timestamp BETWEEN '%s' AND '%s' AND stationMac = '%s'
		AND prev_time IS NOT NULL -- 排除第一条记录无前值的情况
		ORDER BY stationMac, current_time;
		`, startedAt, endedAt, macAddr,
	)
	return rawsql

}

func (s *WlanUserService) getWlanUserRSSI(macAddr string, startedAt, endedAt time.Time) {

}

func (s *WlanUserService) getWlanUserSNR(macAddr string, startedAt, endedAt time.Time) {

}

func (s *WlanUserService) getWlanUserLogs(macAddr string, startedAt, endedAt time.Time) {

}
