package nettyx_wlanstation

type WlanUser struct {
	StationMac           string  `json:"stationMac"`                     // 终端MAC
	StationIp            string  `json:"stationIp"`                      // 终端IP
	StationUsername      string  `json:"stationUsername"`                // 终端用户名
	StationApMac         *string `json:"stationApMac,omitempty"`         // AP MAC
	StationApName        *string `json:"stationApName,omitempty"`        // AP 名称
	StationESSID         string  `json:"stationESSID"`                   // ESSID
	StationVlan          *uint64 `json:"stationVlan,omitempty"`          // VLAN
	StationChannel       uint64  `json:"stationChannel,omitempty"`       // 信道
	StationChanBandWidth *string `json:"stationChanBandWidth,omitempty"` // 信道带宽
	StationRadioType     string  `json:"stationRadioType"`               // radio类型
	StationSNR           *uint64 `json:"stationSNR,omitempty"`           // 终端NR
	StationRSSI          uint64  `json:"stationRSSI"`                    // 终端RSSI
	StationRxBits        uint64  `json:"stationRxBits"`                  // 终端下行流量
	StationTxBits        uint64  `json:"stationTxBits"`                  // 终端上行流量
	StationMaxSpeed      *uint64 `json:"stationMaxSpeed,omitempty"`      // 终端协商速率
	StationOnlineTime    uint64  `json:"stationOnlineTime"`              // 终端在线时间
}
