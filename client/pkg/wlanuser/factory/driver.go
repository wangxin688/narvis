package wlan_user_factory

type WlanUserDriver interface {
	stationMac() map[string]string
	stationIp() map[string]string
	stationUsername() map[string]string
	stationApMac() map[string]string
	stationSSID() map[string]string
	stationBSSID() map[string]string
	stationChannel() map[string]string
	stationChanBandWidth() map[string]string
	stationRadioType() map[string]string
	stationSNR() map[string]string
	stationRSSI() map[string]string
	stationRxBytes() map[string]string
	stationTxBytes() map[string]string
	stationMaxSpeed() map[string]string
	stationOnlineTime() map[string]string
	Run() []WlanUser
}
