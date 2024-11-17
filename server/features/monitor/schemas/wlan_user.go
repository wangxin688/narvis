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
	SiteId          *string   `form:"siteId" binding:"omitempty,uuid"`
	StartedAt       time.Time `form:"startedAt" binding:"required,datetime"`
	EndedAt         time.Time `form:"endedAt" binding:"required,datetime"`
	StationMac      *string   `form:"stationMac" binding:"omitempty,mac"`
	StationIp       *string   `form:"stationIp" binding:"omitempty,ip"`
	StationUsername *string   `form:"stationUsername" binding:"omitempty"`
	ApName          *string   `form:"apName" binding:"omitempty"`
	StationESSID    *string   `form:"stationESSID" binding:"omitempty"`
}
