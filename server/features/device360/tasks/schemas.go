package device360_tasks

type Vector struct {
	Name      string
	Label     map[string]string
	Timestamp string
	Values    float64
}

type DeviceSchema struct {
	ICMPPing          float64           `json:"icmp_ping"`
	CpuUtilization    []float64         `json:"cpu_utilization"`
	MemoryUtilization []float64         `json:"memory_utilization"`
	Temperature       []float64         `json:"temperature"`
	FanStatus         []float64         `json:"fan_status"`
	PowerSupplyStatus []float64         `json:"power_supply_status"`
	RxDiscards        []float64         `json:"rx_discards"`
	TxDiscards        []float64         `json:"tx_discards"`
	RxErrors          []float64         `json:"rx_errors"`
	TxErrors          []float64         `json:"tx_errors"`
	RxRate            []float64         `json:"rx_rate"`
	TxRate            []float64         `json:"tx_rate"`
	OperationalStatus []float64         `json:"operational_status"`
	Labels            map[string]string `json:"labels"`
}

type ApSchema struct {
	ChannelUtilization        []float64         `json:"channel_utilization"`
	ChannelAssociationClients []float64         `json:"channel_associated_clients"`
	ApStatus                  float64           `json:"ap_status"`
	Labels                    map[string]string `json:"labels"`
}
