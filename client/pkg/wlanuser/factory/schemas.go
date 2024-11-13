package wlan_user_factory

type WlanUser struct {
	StationMac           string  `json:"stationMac"`                     // 终端MAC
	StationIp            string  `json:"stationIp"`                      // 终端IP
	StationUsername      string  `json:"stationUsername"`                // 终端用户名
	StationApMac         *string `json:"stationApMac,omitempty"`         // AP MAC
	StationApName        *string `json:"stationApName,omitempty"`        // AP 名称
	StationSSID          string  `json:"stationSSID"`                    // ESSID
	StationBSSID         *string `json:"stationBSSID"`                   // BSSID
	StationChannel       string  `json:"stationChannel,omitempty"`       // 信道
	StationChanBandWidth *string `json:"stationChanBandWidth,omitempty"` // 信道带宽
	StationRadioType     string  `json:"stationRadioType,omitempty"`     // radio类型
	StationSNR           string  `json:"stationSNR"`                     // 终端NR
	StationRSSI          string  `json:"stationRSSI"`                    // 终端RSSI
	StationRxBytes       string  `json:"stationRxBytes"`                 // 终端下行流量
	StationTxBytes       string  `json:"stationTxBytes"`                 // 终端上行流量
	StationMaxSpeed      *string `json:"stationMaxSpeed"`                // 终端协商速率
	StationOnlineTime    *string `json:"stationOnlineTime"`              // 终端在线时间
}

func (w *WlanUser) Enrich() {

}
