package schemas

import (
	"time"

	ts "github.com/wangxin688/narvis/server/tools/schemas"
)

type WlanUserTrend struct {
	ESSID     string    `json:"essid"`
	Timestamp time.Time `json:"timestamp"`
	Value     uint      `json:"value"`
}

type WlanUserTrendRequest struct {
	SiteId     *string   `form:"siteId" binding:"uuid,omitempty"`
	StartedAt  time.Time `form:"startedAt" binding:"required,datetime"`
	EndedAt    time.Time `form:"endedAt" binding:"required,datetime"`
	DataPoints uint      `form:"dataPoints" binding:"required,gte=300"`
}

type WlanUserQuery struct {
	ts.PageInfo
	SiteId          string    `form:"siteId" binding:"uuid"`
	StartedAt       time.Time `form:"startedAt" binding:"required,datetime"`
	EndedAt         time.Time `form:"endedAt" binding:"required,datetime"`
	StationMac      *string   `form:"stationMac" binding:"omitempty,mac"`  // support only one of StationMac or StationIp or StationUsername
	StationIp       *string   `form:"stationIp" binding:"omitempty,ip"`    // support only one of StationMac or StationIp or StationUsername
	StationUsername *string   `form:"stationUsername" binding:"omitempty"` // support only one of StationMac or StationIp or StationUsername
	ApName          *string   `form:"apName" binding:"omitempty"`
	StationESSID    *string   `form:"stationESSID" binding:"omitempty"`
}

type WlanUserItem struct {
	StationMac           string    `json:"stationMac"`                     // 终端MAC
	StationIp            string    `json:"stationIp"`                      // 终端IP
	StationUsername      string    `json:"stationUsername"`                // 终端用户名
	StationApMac         *string   `json:"stationApMac,omitempty"`         // AP MAC
	StationApName        *string   `json:"stationApName,omitempty"`        // AP 名称
	StationESSID         string    `json:"stationESSID"`                   // ESSID
	StationVlan          *uint64   `json:"stationVlan,omitempty"`          // VLAN
	StationChannel       uint64    `json:"stationChannel,omitempty"`       // 信道
	StationChanBandWidth *string   `json:"stationChanBandWidth,omitempty"` // 信道带宽
	StationRadioType     string    `json:"stationRadioType"`               // radio类型
	StationSNR           *uint64   `json:"stationSNR,omitempty"`           // 终端NR
	StationRSSI          uint64    `json:"stationRSSI"`                    // 终端RSSI
	StationMaxSpeed      *uint64   `json:"stationMaxSpeed,omitempty"`      // 终端协商速率
	StationOnlineTime    uint64    `json:"stationOnlineTime"`              // 终端在线时间
	LastSeenAt           time.Time `json:"lastSeenAt"`                     // 最后在线时间
	Online               bool      `json:"online"`                         // 是否在线
}

type WlanUserListResult struct {
	WlanUserItem
	TotalBits uint64 `json:"totalBits"`
	RxBits    uint64 `json:"rxBits"`
	TxBits    uint64 `json:"txBits"`
	AvgSpeed  uint64 `json:"avgSpeed"`
}

type WlanUserThroughput struct {
	StationMac string `json:"stationMac"`
	TotalBits  uint64 `json:"totalBits"`
	RxBits     uint64 `json:"rxBits"`
	TxBits     uint64 `json:"txBits"`
	AvgSpeed   uint64 `json:"avgSpeed"`
}

type WlanUserListResponse struct {
	Users   []*WlanUserListResult `json:"users"`
	Total   int64                 `json:"total"`
	Online  int64                 `json:"online"`
	Offline int64                 `json:"offline"`
}

