package models

var WlanStationTableName = "wlan_station" // wlan_stationTable

type WlanStation struct {
	BaseTimeScaleModel
	OrganizationId       string  `gorm:"column:organizationId;type:uuid;not null"`
	SiteId               string  `gorm:"column:siteId;type:uuid;not null"`
	StationMac           string  `gorm:"column:stationMac;not null"`
	StationIp            string  `gorm:"column:stationIp;not null"`
	StationUsername      string  `gorm:"column:stationUsername;not null"`
	StationApMac         *string `gorm:"column:stationApMac;default:null"`
	StationApName        *string `gorm:"column:stationApName;default:null"`
	StationESSID         string  `gorm:"column:stationESSID;not null"`
	StationVlan          *uint64 `gorm:"column:stationVlan;default:null"`
	StationChannel       uint64  `gorm:"column:stationChannel;not null"`
	StationChanBandWidth *string `gorm:"column:stationChanBandWidth;default:null"`
	StationRadioType     string  `gorm:"column:stationRadioType;not null"`
	StationSNR           *uint64 `gorm:"column:stationSNR;default:null"`
	StationRSSI          uint64  `gorm:"column:stationRSSI;not null"`
	StationRxBits        uint64  `gorm:"column:stationRxBits;not null"`
	StationTxBits        uint64  `gorm:"column:stationTxBits;not null"`
	StationMaxSpeed      *uint64 `gorm:"column:stationMaxSpeed;default:null"`
	StationOnlineTime    uint64  `gorm:"column:stationOnlineTime;not null"`
}

func (WlanStation) TableName() string {
	return WlanStationTableName
}
