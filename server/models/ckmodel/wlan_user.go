package ckmodel

import "time"

type WlanStation struct {
	Ts                      time.Time `gorm:"ts"`
	OrganizationId          string    `gorm:"organizationId"`
	SiteId                  string    `gorm:"siteId"`
	StationMac              string    `gorm:"stationMac"`
	StationIp               string    `gorm:"stationIp"`
	StationUsername         string    `gorm:"stationUsername"`
	StationApMac            string    `gorm:"stationApMac"`
	StationApName           string    `gorm:"stationApName"`
	StationESSID            string    `gorm:"stationESSID"`
	StationChannel          uint16    `gorm:"stationChannel"`
	StationChannelBandWidth string    `gorm:"stationChannelBandwidth"`
	StationSNR              uint16    `gorm:"stationSNR"`
	StationRSSI             int16     `gorm:"stationRSSI"`
	StationRxBits           uint64    `gorm:"stationRxBits"`
	StationTxBits           uint64    `gorm:"stationTxBits"`
	StationMaxSpeed         uint32    `gorm:"stationMaxSpeed"`
	StationOnlineTime       uint32    `gorm:"stationOnlineTime"`
}
