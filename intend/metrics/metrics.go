package metrics

import (
	"github.com/samber/lo"
	"github.com/wangxin688/narvis/intend/common"
)

type MetricNameEnum string

type ICMPMetrics MetricNameEnum
type CircuitMetrics MetricNameEnum
type InterfaceMetrics MetricNameEnum
type SwitchingMetrics MetricNameEnum
type RoutingMetrics MetricNameEnum
type WirelessMetrics MetricNameEnum
type SecurityMetrics MetricNameEnum
type Device360Metrics MetricNameEnum

const (
	ICMPPing         ICMPMetrics = "icmp_ping"
	ICMPResponseTime ICMPMetrics = "icmp_response_time"
	ICMPPacketLoss   ICMPMetrics = "icmp_loss"
)

const (
	CircuitPing         CircuitMetrics = "circuit_icmp_ping"
	CircuitResponseTime CircuitMetrics = "circuit_icmp_response_time"
	CircuitPacketLoss   CircuitMetrics = "circuit_icmp_loss"
	CircuitRxBytes      CircuitMetrics = "circuit_rx_bytes"
	CircuitTxBytes      CircuitMetrics = "circuit_tx_bytes"
	CircuitRxDiscards   CircuitMetrics = "circuit_rx_discards"
	CircuitTxDiscards   CircuitMetrics = "circuit_tx_discards"
	CircuitRxErrors     CircuitMetrics = "circuit_rx_errors"
	CircuitTxErrors     CircuitMetrics = "circuit_tx_errors"
	CircuitRxRate       CircuitMetrics = "circuit_rx_rate"
	CircuitTxRate       CircuitMetrics = "circuit_tx_rate"
)

// tags: interface, description
const (
	RxBytes           InterfaceMetrics = "rx_bytes"
	TxBytes           InterfaceMetrics = "tx_bytes"
	RxDiscards        InterfaceMetrics = "rx_discards"
	TxDiscards        InterfaceMetrics = "tx_discards"
	RxErrors          InterfaceMetrics = "rx_errors"
	TxErrors          InterfaceMetrics = "tx_errors"
	RxRate            InterfaceMetrics = "rx_rate"
	TxRate            InterfaceMetrics = "tx_rate"
	OperationalStatus InterfaceMetrics = "operational_status"
	HighSpeed         InterfaceMetrics = "high_speed"
	Duplex            InterfaceMetrics = "duplex"
)

const (
	CpuUsage          SwitchingMetrics = "cpu_usage"
	MemoryUsage       SwitchingMetrics = "memory_usage"
	DiskUsage         SwitchingMetrics = "disk_usage"
	SystemLoad        SwitchingMetrics = "system_load"
	FanStatus         SwitchingMetrics = "fan_status"
	PowerSupplyStatus SwitchingMetrics = "power_supply_status"
	Temperature       SwitchingMetrics = "temperature"
	SnmpAgentStatus   SwitchingMetrics = "snmp_agent_status"
	Uptime            SwitchingMetrics = "uptime"
)

// Wireless metric tags
// radio_type: 2.4GHz/5GHz/6GHz
// channel: 1..14, 36-64, 100-140, 149-165
const (
	ApStatus                  WirelessMetrics = "ap_status"
	ApUptime                  WirelessMetrics = "ap_uptime"
	ChannelUtilization        WirelessMetrics = "channel_utilization"
	ChannelNoise              WirelessMetrics = "channel_noise"
	ChannelReceiveErrorRate   WirelessMetrics = "channel_error_rate"
	ChannelRxRate             WirelessMetrics = "channel_rx_rate"
	ChannelTxRate             WirelessMetrics = "channel_tx_rate"
	ChannelRxBytes            WirelessMetrics = "channel_rx_bytes"
	ChannelTxBytes            WirelessMetrics = "channel_tx_bytes"
	ChannelFrameRetryRate     WirelessMetrics = "channel_frame_retry_rate"
	ChannelEirp               WirelessMetrics = "channel_eirp" // transmit power
	ChannelInterferenceRate   WirelessMetrics = "channel_interference_rate"
	ChannelAssociationClients WirelessMetrics = "channel_association_clients"
)

const (
	HealthScore         Device360Metrics = "health_score"
	IcmpScore           Device360Metrics = "icmp_score"
	CpuScore            Device360Metrics = "cpu_score"
	MemoryScore         Device360Metrics = "memory_score"
	FanScore            Device360Metrics = "fan_score"
	FanAnomaly          Device360Metrics = "fan_anomaly"
	PowerSupplyScore    Device360Metrics = "power_supply_score"
	PowerAnomaly        Device360Metrics = "power_supply_anomaly"
	TemperatureScore    Device360Metrics = "temperature_score"
	IfErrorScore        Device360Metrics = "interface_error_score"
	IfErrorAnomaly      Device360Metrics = "interface_error_anomaly"
	IfDiscardScore      Device360Metrics = "interface_discard_score"
	IfDiscardAnomaly    Device360Metrics = "interface_discard_anomaly"
	IfTrafficScore      Device360Metrics = "interface_traffic_score"
	IfTrafficAnomaly    Device360Metrics = "interface_traffic_anomaly"
	IfOperStatusScore   Device360Metrics = "interface_oper_status_score"
	IfOperStatusAnomaly Device360Metrics = "interface_oper_status_anomaly"
)

type Metric struct {
	Name                      MetricNameEnum
	Description               common.I18n
	Unit                      *string
	ValueMapping              *map[int]string
	Tags                      []*string
	Legend                    string
	DefaultQueryRangeFunction string
}

func getMetricMeta() map[MetricNameEnum]Metric {
	var ms string = "ms"
	var percent string = "%"
	var bps string = "bps"
	var cs string = "c/s"
	var mbps string = "Mbps"
	var celsius string = "°C"
	var seconds string = "s"
	var dBm string = "dBm"

	var opStatus map[int]string = map[int]string{
		-1: "nodata",
		0:  "down",
		1:  "up",
	}

	var duplexStatus map[int]string = map[int]string{
		1: "unknown",
		2: "half",
		3: "full",
	}

	var entityStatus map[int]string = map[int]string{
		-1: "nodata",
		1:  "normal",
		2:  "abnormal",
	}

	metricMeta := map[MetricNameEnum]Metric{
		MetricNameEnum(ICMPPing): {
			Name:                      MetricNameEnum(ICMPPing),
			Description:               common.I18n{En: "ICMP ping", Zh: "ICMP连通性"},
			Unit:                      nil,
			ValueMapping:              &opStatus,
			Legend:                    "{device_name} ICMP Ping",
			DefaultQueryRangeFunction: "min_over_time",
		},
		MetricNameEnum(ICMPResponseTime): {
			Name:                      MetricNameEnum(ICMPResponseTime),
			Description:               common.I18n{En: "ICMP response time", Zh: "ICMP响应时间"},
			Unit:                      &ms,
			ValueMapping:              nil,
			Legend:                    "{device_name} ICMP Response Time",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(ICMPPacketLoss): {
			Name:                      MetricNameEnum(ICMPPacketLoss),
			Description:               common.I18n{En: "ICMP packet loss", Zh: "ICMP丢包率"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{device_name} ICMP Packet Loss",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(RxBytes): {
			Name:                      MetricNameEnum(RxBytes),
			Description:               common.I18n{En: "Rx Bytes", Zh: "下行流量"},
			Unit:                      &bps,
			ValueMapping:              nil,
			Legend:                    "{device_name} {interface}:{description} Rx Bytes",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(TxBytes): {
			Name:                      MetricNameEnum(TxBytes),
			Description:               common.I18n{En: "Tx Bytes", Zh: "上行流量"},
			Unit:                      &bps,
			ValueMapping:              nil,
			Legend:                    "{device_name} {interface}:{description} Tx Bytes",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(RxDiscards): {
			Name:                      MetricNameEnum(RxDiscards),
			Description:               common.I18n{En: "Rx Packet Discards", Zh: "下行丢包数"},
			Unit:                      &cs,
			ValueMapping:              nil,
			Legend:                    "{device_name} {interface}:{description} Rx Packet Discards",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(TxDiscards): {
			Name:                      MetricNameEnum(TxDiscards),
			Description:               common.I18n{En: "Tx Packet Discards", Zh: "上行丢包数"},
			Unit:                      &cs,
			ValueMapping:              nil,
			Legend:                    "{device_name} {interface}:{description} Tx Packet Discards",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(RxErrors): {
			Name:                      MetricNameEnum(RxErrors),
			Description:               common.I18n{En: "Rx Packet Errors", Zh: "下行错包数"},
			Unit:                      &cs,
			ValueMapping:              nil,
			Legend:                    "{device_name} {interface}:{description} Rx Packet Errors",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(TxErrors): {
			Name:                      MetricNameEnum(TxErrors),
			Description:               common.I18n{En: "Tx Packet Errors", Zh: "上行错包数"},
			Unit:                      &cs,
			ValueMapping:              nil,
			Legend:                    "{device_name} {interface}:{description} Tx Packet Errors",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(RxRate): {
			Name:                      MetricNameEnum(RxRate),
			Description:               common.I18n{En: "Rx Rate", Zh: "下行带宽利用率"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{device_name} {interface}:{description} Rx Rate",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(TxRate): {
			Name:                      MetricNameEnum(TxRate),
			Description:               common.I18n{En: "Tx Rate", Zh: "上行带宽利用率"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{device_name} {interface}:{description} Tx Rate",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(OperationalStatus): {
			Name:                      MetricNameEnum(OperationalStatus),
			Description:               common.I18n{En: "Operational Status", Zh: "运行状态"},
			Unit:                      nil,
			ValueMapping:              &opStatus,
			Legend:                    "{device_name} {interface}:{description} Operational Status",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(HighSpeed): {
			Name:                      MetricNameEnum(HighSpeed),
			Description:               common.I18n{En: "Port Speed", Zh: "接口速率"},
			Unit:                      &mbps,
			ValueMapping:              nil,
			Legend:                    "{device_name} {interface}:{description} Port Speed",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(Duplex): {
			Name:                      MetricNameEnum(Duplex),
			Description:               common.I18n{En: "Duplex Status", Zh: "接口双工状态"},
			Unit:                      nil,
			ValueMapping:              &duplexStatus,
			Legend:                    "{device_name} {interface}:{description} Duplex Status",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(CpuUsage): {
			Name:                      MetricNameEnum(CpuUsage),
			Description:               common.I18n{En: "CPU Usage", Zh: "CPU利用率"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{device_name} CPU#{cpu} CPU Usage",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(MemoryUsage): {
			Name:                      MetricNameEnum(MemoryUsage),
			Description:               common.I18n{En: "Memory Usage", Zh: "内存利用率"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{device_name} Memory Usage",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(DiskUsage): {
			Name:                      MetricNameEnum(DiskUsage),
			Description:               common.I18n{En: "Disk Usage", Zh: "磁盘利用率"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{device_name} Disk Usage",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(SystemLoad): {
			Name:                      MetricNameEnum(SystemLoad),
			Description:               common.I18n{En: "System Load", Zh: "系统负载"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{device_name} System Load",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(FanStatus): {
			Name:                      MetricNameEnum(FanStatus),
			Description:               common.I18n{En: "Fan Status", Zh: "风扇状态"},
			Unit:                      nil,
			ValueMapping:              &entityStatus,
			Legend:                    "{device_name} Fan#{entity} Status",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(PowerSupplyStatus): {
			Name:                      MetricNameEnum(PowerSupplyStatus),
			Description:               common.I18n{En: "Power Supply Status", Zh: "电源状态"},
			Unit:                      nil,
			ValueMapping:              &entityStatus,
			Legend:                    "{device_name} Power Supply#{entity} Status",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(Temperature): {
			Name:                      MetricNameEnum(Temperature),
			Description:               common.I18n{En: "Temperature", Zh: "温度"},
			Unit:                      &celsius,
			ValueMapping:              nil,
			Legend:                    "{device_name} Temperature",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(SnmpAgentStatus): {
			Name:                      MetricNameEnum(SnmpAgentStatus),
			Description:               common.I18n{En: "SNMP Agent Status", Zh: "SNMP Agent状态"},
			Unit:                      nil,
			ValueMapping:              &entityStatus,
			Legend:                    "{device_name} SNMP Agent Status",
			DefaultQueryRangeFunction: "min_over_time",
		},
		MetricNameEnum(Uptime): {
			Name:                      MetricNameEnum(Uptime),
			Description:               common.I18n{En: "Uptime", Zh: "运行时间"},
			Unit:                      &seconds,
			ValueMapping:              nil,
			Legend:                    "{device_name} Uptime",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(ApStatus): {
			Name:                      MetricNameEnum(ApStatus),
			Description:               common.I18n{En: "AP Status", Zh: "AP运行状态"},
			Unit:                      nil,
			ValueMapping:              &opStatus,
			Legend:                    "{device_name} Status",
			DefaultQueryRangeFunction: "min_over_time",
		},
		MetricNameEnum(ApUptime): {
			Name:                      MetricNameEnum(ApUptime),
			Description:               common.I18n{En: "AP Uptime", Zh: "AP运行时间"},
			Unit:                      &seconds,
			ValueMapping:              nil,
			Legend:                    "{device_name} Uptime",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(ChannelUtilization): {
			Name:                      MetricNameEnum(ChannelUtilization),
			Description:               common.I18n{En: "Channel Utilization", Zh: "信道利用率"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{device_name}(channel) Channel Utilization",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(ChannelNoise): {
			Name:                      MetricNameEnum(ChannelNoise),
			Description:               common.I18n{En: "Channel Noise", Zh: "信道噪声"},
			Unit:                      &dBm,
			ValueMapping:              nil,
			Legend:                    "{device_name}(channel) Channel Noise",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(ChannelReceiveErrorRate): {
			Name:                      MetricNameEnum(ChannelReceiveErrorRate),
			Description:               common.I18n{En: "Channel Receive Error Rate", Zh: "信道接收错误率"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{device_name}(channel) Channel Receive Error Rate",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(ChannelRxRate): {
			Name:                      MetricNameEnum(ChannelRxRate),
			Description:               common.I18n{En: "Channel Rx Rate", Zh: "信道上行速率"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{device_name}(channel) Channel Rx Rate",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(ChannelTxRate): {
			Name:                      MetricNameEnum(ChannelTxRate),
			Description:               common.I18n{En: "Channel Tx Rate", Zh: "信道下行速率"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{device_name}(channel) Channel Tx Rate",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(ChannelRxBytes): {
			Name:                      MetricNameEnum(ChannelRxBytes),
			Description:               common.I18n{En: "Channel Rx Bytes", Zh: "信道上行流量"},
			Unit:                      &bps,
			ValueMapping:              nil,
			Legend:                    "{device_name}(channel) Channel Rx Bytes",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(ChannelTxBytes): {
			Name:                      MetricNameEnum(ChannelTxBytes),
			Description:               common.I18n{En: "Channel Tx Bytes", Zh: "信道下行流量"},
			Unit:                      &bps,
			ValueMapping:              nil,
			Legend:                    "{device_name}(channel) Channel Tx Bytes",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(ChannelFrameRetryRate): {
			Name:                      MetricNameEnum(ChannelFrameRetryRate),
			Description:               common.I18n{En: "Channel Frame Retry", Zh: "信道帧重传率"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{device_name}(channel) Channel Frame Retry",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(ChannelEirp): {
			Name:                      MetricNameEnum(ChannelEirp),
			Description:               common.I18n{En: "Channel Transmit Power", Zh: "信道发射功率"},
			Unit:                      &dBm,
			ValueMapping:              nil,
			Legend:                    "{device_name}(channel) Channel Transmit Power",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(ChannelInterferenceRate): {
			Name:                      MetricNameEnum(ChannelInterferenceRate),
			Description:               common.I18n{En: "Channel Interference", Zh: "信道干扰"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{device_name}(channel) Channel Interference",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(ChannelAssociationClients): {
			Name:                      MetricNameEnum(ChannelAssociationClients),
			Description:               common.I18n{En: "Channel Association Clients", Zh: "信道关联客户端"},
			Unit:                      nil,
			ValueMapping:              nil,
			Legend:                    "{device_name}(channel) Channel Association Clients",
			DefaultQueryRangeFunction: "max_over_time",
		},
	}
	return metricMeta
}

func GetListMetric() (int, []Metric) {

	metricMeta := getMetricMeta()
	return len(metricMeta), lo.Values(metricMeta)
}

func GetMetric(metricName string) *Metric {

	metricMeta := getMetricMeta()

	if _, ok := metricMeta[MetricNameEnum(metricName)]; !ok {
		return nil
	}

	result := metricMeta[MetricNameEnum(metricName)]
	return &result
}
