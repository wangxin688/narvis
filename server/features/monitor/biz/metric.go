package biz

import "time"

type MetricService struct {
	StartedAt time.Time
	EndedAt   *time.Time
	DataPoints uint
	SiteId     string
	DeviceId   string
	ApName     string
	IfName     string
	RadioType  string
	FuncName   string
}
