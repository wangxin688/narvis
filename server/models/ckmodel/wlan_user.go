package ckmodel

import "time"

var WlanStationTableName = "wlan_station" // wlan_stationTable

type WlanStation struct {
	Ts                      time.Time `gorm:"column:ts;type:DateTime64(3, 'UTC');default:now64(3)"`
	OrganizationId          string    `gorm:"column:organizationId"`
	SiteId                  string    `gorm:"column:siteId"`
	StationMac              string    `gorm:"column:stationMac"`
	StationIp               string    `gorm:"column:stationIp"`
	StationUsername         string    `gorm:"column:stationUsername"`
	StationApMac            string    `gorm:"column:stationApMac"`
	StationApName           string    `gorm:"column:stationApName"`
	StationESSID            string    `gorm:"column:stationESSID"`
	StationChannel          uint16    `gorm:"column:stationChannel"`
	StationChannelBandWidth string    `gorm:"column:stationChannelBandWidth"`
	StationSNR              uint16    `gorm:"column:stationSNR"`
	StationRSSI             int16     `gorm:"column:stationRSSI"`
	StationRxBits           uint64    `gorm:"column:stationRxBits"`
	StationTxBits           uint64    `gorm:"column:stationTxBits"`
	StationMaxSpeed         uint32    `gorm:"column:stationMaxSpeed"`
	StationOnlineTime       uint32    `gorm:"column:stationOnlineTime"`
}

func (WlanStation) TableName() string {
	return WlanStationTableName
}
