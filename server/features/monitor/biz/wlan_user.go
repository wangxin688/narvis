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
	whereCond := "1=1"
	if query.SiteId != nil && *query.SiteId != "" {
		whereCond = fmt.Sprintf("siteId = '%s'", *query.SiteId)
	}
	rawSql := fmt.Sprintf(
		`
	SELECT
		TO_TIMESTAMP(FLOOR(EXTRACT(EPOCH FROM "time") / %d) * %d) AS timestamp,
		"stationESSID" as "ESSID",
		COUNT(DISTINCT "stationMac") AS value
	FROM
		wlan_station
	WHERE
		"time" >= '%s'
		AND "time" <= '%s'
		AND %s
	GROUP BY
		timestamp, "ESSID"
	ORDER BY
		timestamp, "ESSID";
		`,
		interval, interval, query.StartedAt.Format(time.RFC3339), query.EndedAt.Format(time.RFC3339), whereCond,
	)
	err := infra.DB.Raw(rawSql).Scan(&results).Error
	if err != nil {
		logger.Logger.Error("[GetWlanUserTrend]", zap.Error(err))
		return nil, err
	}
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
		SELECT DISTINCT ON ("stationMac")
			TO_TIMESTAMP(EXTRACT(EPOCH FROM "time")) AS "lastSeenAt",
			CASE
				WHEN (EXTRACT(EPOCH FROM TIMESTAMP '%s') - EXTRACT(EPOCH FROM "lastSeenAt")) < 600 THEN TRUE
				ELSE FALSE
			END AS "isActive",
			"stationMac",
			"stationIp",
			"stationUsername",
			"stationApMac",
			"stationApName",
			"stationESSID",
			"stationVlan",
			"stationChannel",   
			"stationChanBandWidth",
			"stationRadioType",
			"stationSNR",
			"stationRSSI",
			"stationMaxSpeed",
			"stationOnlineTime"
		FROM wlan_station
		WHERE
			"organizationId" = '%s'
			AND "time" >= '%s'
			AND "time" <= '%s'
			AND "siteId" = '%s'
			AND %s
		ORDER BY stationMac, "time" DESC;
		LIMIT %d OFFSET %d
		`, query.EndedAt.Format(time.RFC3339), orgId, query.StartedAt, query.EndedAt, query.SiteId, subQuery, (*query.PageInfo.Page-1)**query.PageInfo.PageSize, query.PageInfo.PageSize,
	)
	countRawSql := fmt.Sprintf(
		`
		SELECT
			COUNT(*) AS "totalCount",
			SUM(CASE WHEN isActive = TRUE THEN 1 ELSE 0 END) AS "onlineCount"
		FROM (
			SELECT
				EXTRACT(EPOCH FROM "time") AS "lastSeenAt",
				CASE
					WHEN (EXTRACT(EPOCH FROM TIMESTAMP '%s') - EXTRACT(EPOCH FROM "lastSeenAt")) < 500 THEN TRUE
					ELSE FALSE
				END AS isActive,
				"stationMac",
				"stationIp",
				"stationUsername",
				"stationApMac",
				"stationApName",
				"stationESSID",
				"stationVlan",
				"stationChannel",
				"stationChanBandWidth",
				"stationRadioType",
				"stationSNR",
				"stationRSSI",
				"stationMaxSpeed",
				"stationOnlineTime"
			FROM wlan_station
			WHERE
				organizationId = '%s'
				AND ts >= '%s'
				AND ts <= '%s'
				AND siteId = '%s'
				AND %s
			ORDER BY stationMac, ts DESC
		)
		`, query.EndedAt.Format(time.RFC3339), orgId, query.StartedAt, query.EndedAt, query.SiteId, subQuery,
	)

	err := infra.DB.Raw(countRawSql).Scan(&countUser).Error
	if err != nil {
		return response, nil
	}
	if countUser.TotalCount == 0 {
		return response, nil
	}
	err = infra.DB.Raw(rawSql).Scan(&results).Error
	if err != nil {
		return response, nil
	}
	err = infra.DB.Raw(rawSql).Scan(&results).Error
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
	WITH time_interval AS (
		SELECT 
			EXTRACT(EPOCH FROM TIMESTAMP '%s') - EXTRACT(EPOCH FROM TIMESTAMP '%s') AS interval_in_seconds
		),
		cte_station_data AS (
			SELECT 
				"stationMac",
				"time",
				"stationRxBits",
				"stationTxBits",
				MIN("time") OVER (PARTITION BY "stationMac") AS min_time,
				MAX("time") OVER (PARTITION BY "stationMac") AS max_time
			FROM wlan_station
			WHERE "time" BETWEEN '%s' AND '%s' 
			AND "stationMac" IN ('%s')
		),
		cte_min_max AS (
			SELECT 
				"stationMac",
				MAX(CASE WHEN ts = min_time THEN "stationRxBits" END) AS minRx,
				MAX(CASE WHEN ts = max_time THEN "stationRxBits" END) AS maxRx,
				MAX(CASE WHEN ts = min_time THEN "stationTxBits" END) AS minTx,
				MAX(CASE WHEN ts = max_time THEN "stationTxBits" END) AS maxTx
			FROM cte_station_data
			GROUP BY "stationMac"
		)
		SELECT 
			"stationMac",
			(maxRx - minRx) AS rxBits,
			(maxTx - minTx) AS txBits
		FROM cte_min_max
		ORDER BY "stationMac";
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
		WITH cte_lagged_data AS (
			SELECT
				stationMac,
				timestamp AS current_time,
				LAG(timestamp) OVER (PARTITION BY stationMac ORDER BY timestamp) AS prev_time,
				stationRxBits,
				LAG(stationRxBits) OVER (PARTITION BY stationMac ORDER BY timestamp) AS prev_rx,
				stationTxBits,
				LAG(stationTxBits) OVER (PARTITION BY stationMac ORDER BY timestamp) AS prev_tx
			FROM station_data
			WHERE timestamp BETWEEN '%s' AND '%s' 
			AND stationMac = '%s'
		)
		SELECT
			stationMac,
			current_time,
			prev_time AS previous_time,
			(stationRxBits - prev_rx) AS rx_diff,
			(stationTxBits - prev_tx) AS tx_diff,
			EXTRACT(EPOCH FROM current_time) - EXTRACT(EPOCH FROM prev_time) AS time_diff_seconds,
			(stationRxBits - prev_rx) / NULLIF(EXTRACT(EPOCH FROM current_time) - EXTRACT(EPOCH FROM prev_time), 0) AS rx_trend_per_second,
			(stationTxBits - prev_tx) / NULLIF(EXTRACT(EPOCH FROM current_time) - EXTRACT(EPOCH FROM prev_time), 0) AS tx_trend_per_second
		FROM cte_lagged_data
		WHERE prev_time IS NOT NULL -- Exclude the first record
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
