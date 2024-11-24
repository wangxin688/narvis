package metrics

import (
	"github.com/samber/lo"
	"github.com/wangxin688/narvis/intend/model/i18n"
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
type ServerMetrics MetricNameEnum

const (
	ICMPPing         ICMPMetrics = "icmp_ping"
	ICMPResponseTime ICMPMetrics = "icmp_response_time"
	ICMPPacketLoss   ICMPMetrics = "icmp_loss"
)

const (
	CircuitPing         CircuitMetrics = "circuit_icmp_ping"
	CircuitResponseTime CircuitMetrics = "circuit_icmp_response_time"
	CircuitPacketLoss   CircuitMetrics = "circuit_icmp_loss"
	CircuitRxBits       CircuitMetrics = "circuit_rx_bits"
	CircuitTxBits       CircuitMetrics = "circuit_tx_bits"
	CircuitRxDiscards   CircuitMetrics = "circuit_rx_discards"
	CircuitTxDiscards   CircuitMetrics = "circuit_tx_discards"
	CircuitRxErrors     CircuitMetrics = "circuit_rx_errors"
	CircuitTxErrors     CircuitMetrics = "circuit_tx_errors"
	CircuitRxRate       CircuitMetrics = "circuit_rx_rate"
	CircuitTxRate       CircuitMetrics = "circuit_tx_rate"
)

// tags: interface, description
const (
	RxBits            InterfaceMetrics = "rx_bits"
	TxBits            InterfaceMetrics = "tx_bits"
	RxDiscards        InterfaceMetrics = "rx_discards"
	TxDiscards        InterfaceMetrics = "tx_discards"
	RxErrors          InterfaceMetrics = "rx_errors"
	TxErrors          InterfaceMetrics = "tx_errors"
	RxRate            InterfaceMetrics = "rx_rate"
	TxRate            InterfaceMetrics = "tx_rate"
	OperationalStatus InterfaceMetrics = "operational_status"
	HighSpeed         InterfaceMetrics = "high_speed"
	DuplexStatus      InterfaceMetrics = "duplex_status"
)

const (
	CpuUtilization    SwitchingMetrics = "cpu_utilization"
	MemoryUtilization SwitchingMetrics = "memory_utilization"
	DiskUsage         SwitchingMetrics = "space_utilization"
	SystemLoad        SwitchingMetrics = "system_load"
	FanStatus         SwitchingMetrics = "fan_status"
	PowerSupplyStatus SwitchingMetrics = "power_supply_status"
	Temperature       SwitchingMetrics = "temperature"
	SnmpAgentStatus   SwitchingMetrics = "snmp_agent_availability"
	Uptime            SwitchingMetrics = "uptime"
)

// Wireless metric tags
// radio_type: 2.4GHz/5GHz/6GHz
// channel: 1..14, 36-64, 100-140, 149-165
// h3c: no way to get client per-channel, only per-ap
const (
	ApStatus                  WirelessMetrics = "ap_status"
	ApUptime                  WirelessMetrics = "ap_uptime"
	ApCpuUtilization          WirelessMetrics = "ap_cpu_utilization"
	ApMemoryUtilization       WirelessMetrics = "ap_memory_utilization"
	ApBootstraps              WirelessMetrics = "ap_bootstraps"
	ChannelUtilization        WirelessMetrics = "channel_utilization"
	ChannelNoise              WirelessMetrics = "channel_noise"
	ChannelTransmitPower      WirelessMetrics = "channel_transmit_power" // transmit power
	ChannelInterferenceRate   WirelessMetrics = "channel_interference_rate"
	ChannelAssociationClients WirelessMetrics = "channel_associated_clients"
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
const (
	LoadAverage1m            ServerMetrics = "load_average_1m"
	LoadAverage5m            ServerMetrics = "load_average_5m"
	LoadAverage15m           ServerMetrics = "load_average_15m"
	ContextSwitchesPerSecond ServerMetrics = "context_switches_per_second"
	InterruptsPerSecond      ServerMetrics = "interrupts_per_second"
	SwapSpaceUtilization     ServerMetrics = "swap_space_utilization"
	DiskReadRate             ServerMetrics = "disk_read_rate"
	DiskWriteRate            ServerMetrics = "disk_write_rate"
	DiskUtilization          ServerMetrics = "disk_utilization"
	InodesFree               ServerMetrics = "inodes_free"
	FileSystemUtilization    ServerMetrics = "file_system_utilization"
)

type Metric struct {
	Name                      MetricNameEnum
	Description               i18n.I18n
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
		2:  "down",
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
			Description:               i18n.I18n{En: "ICMP ping", Zh: "ICMP连通性"},
			Unit:                      nil,
			ValueMapping:              &opStatus,
			Legend:                    "{deviceName} ICMP Ping",
			DefaultQueryRangeFunction: "min_over_time",
		},
		MetricNameEnum(ICMPResponseTime): {
			Name:                      MetricNameEnum(ICMPResponseTime),
			Description:               i18n.I18n{En: "ICMP response time", Zh: "ICMP响应时间"},
			Unit:                      &ms,
			ValueMapping:              nil,
			Legend:                    "{deviceName} ICMP Response Time",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(ICMPPacketLoss): {
			Name:                      MetricNameEnum(ICMPPacketLoss),
			Description:               i18n.I18n{En: "ICMP packet loss", Zh: "ICMP丢包率"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{deviceName} ICMP Packet Loss",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(RxBits): {
			Name:                      MetricNameEnum(RxBits),
			Description:               i18n.I18n{En: "Rx Bits", Zh: "下行流量"},
			Unit:                      &bps,
			ValueMapping:              nil,
			Legend:                    "{deviceName} {interface}:{description} Rx Bits",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(TxBits): {
			Name:                      MetricNameEnum(TxBits),
			Description:               i18n.I18n{En: "Tx Bits", Zh: "上行流量"},
			Unit:                      &bps,
			ValueMapping:              nil,
			Legend:                    "{deviceName} {interface}:{description} Tx Bits",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(RxDiscards): {
			Name:                      MetricNameEnum(RxDiscards),
			Description:               i18n.I18n{En: "Rx Packet Discards", Zh: "下行丢包数"},
			Unit:                      &cs,
			ValueMapping:              nil,
			Legend:                    "{deviceName} {interface}:{description} Rx Packet Discards",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(TxDiscards): {
			Name:                      MetricNameEnum(TxDiscards),
			Description:               i18n.I18n{En: "Tx Packet Discards", Zh: "上行丢包数"},
			Unit:                      &cs,
			ValueMapping:              nil,
			Legend:                    "{deviceName} {interface}:{description} Tx Packet Discards",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(RxErrors): {
			Name:                      MetricNameEnum(RxErrors),
			Description:               i18n.I18n{En: "Rx Packet Errors", Zh: "下行错包数"},
			Unit:                      &cs,
			ValueMapping:              nil,
			Legend:                    "{deviceName} {interface}:{description} Rx Packet Errors",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(TxErrors): {
			Name:                      MetricNameEnum(TxErrors),
			Description:               i18n.I18n{En: "Tx Packet Errors", Zh: "上行错包数"},
			Unit:                      &cs,
			ValueMapping:              nil,
			Legend:                    "{deviceName} {interface}:{description} Tx Packet Errors",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(RxRate): {
			Name:                      MetricNameEnum(RxRate),
			Description:               i18n.I18n{En: "Rx Rate", Zh: "下行带宽利用率"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{deviceName} {interface}:{description} Rx Rate",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(TxRate): {
			Name:                      MetricNameEnum(TxRate),
			Description:               i18n.I18n{En: "Tx Rate", Zh: "上行带宽利用率"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{deviceName} {interface}:{description} Tx Rate",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(OperationalStatus): {
			Name:                      MetricNameEnum(OperationalStatus),
			Description:               i18n.I18n{En: "Operational Status", Zh: "运行状态"},
			Unit:                      nil,
			ValueMapping:              &opStatus,
			Legend:                    "{deviceName} {interface}:{description} Operational Status",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(HighSpeed): {
			Name:                      MetricNameEnum(HighSpeed),
			Description:               i18n.I18n{En: "Port Speed", Zh: "接口速率"},
			Unit:                      &mbps,
			ValueMapping:              nil,
			Legend:                    "{deviceName} {interface}:{description} Port Speed",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(DuplexStatus): {
			Name:                      MetricNameEnum(DuplexStatus),
			Description:               i18n.I18n{En: "Duplex Status", Zh: "接口双工状态"},
			Unit:                      nil,
			ValueMapping:              &duplexStatus,
			Legend:                    "{deviceName} {interface}:{description} Duplex Status",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(CpuUtilization): {
			Name:                      MetricNameEnum(CpuUtilization),
			Description:               i18n.I18n{En: "CPU Usage", Zh: "CPU利用率"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{deviceName} CPU#{cpu} CPU Usage",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(MemoryUtilization): {
			Name:                      MetricNameEnum(MemoryUtilization),
			Description:               i18n.I18n{En: "Memory Usage", Zh: "内存利用率"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{deviceName} Memory Usage",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(DiskUsage): {
			Name:                      MetricNameEnum(DiskUsage),
			Description:               i18n.I18n{En: "Disk Usage", Zh: "磁盘利用率"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{deviceName} Disk Usage",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(SystemLoad): {
			Name:                      MetricNameEnum(SystemLoad),
			Description:               i18n.I18n{En: "System Load", Zh: "系统负载"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{deviceName} System Load",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(FanStatus): {
			Name:                      MetricNameEnum(FanStatus),
			Description:               i18n.I18n{En: "Fan Status", Zh: "风扇状态"},
			Unit:                      nil,
			ValueMapping:              &entityStatus,
			Legend:                    "{deviceName} Fan#{entity} Status",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(PowerSupplyStatus): {
			Name:                      MetricNameEnum(PowerSupplyStatus),
			Description:               i18n.I18n{En: "Power Supply Status", Zh: "电源状态"},
			Unit:                      nil,
			ValueMapping:              &entityStatus,
			Legend:                    "{deviceName} Power Supply#{entity} Status",
			DefaultQueryRangeFunction: "max_over_time",
		},

		MetricNameEnum(Temperature): {
			Name:                      MetricNameEnum(Temperature),
			Description:               i18n.I18n{En: "Temperature", Zh: "温度"},
			Unit:                      &celsius,
			ValueMapping:              nil,
			Legend:                    "{deviceName} Temperature",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(SnmpAgentStatus): {
			Name:                      MetricNameEnum(SnmpAgentStatus),
			Description:               i18n.I18n{En: "SNMP Agent Status", Zh: "SNMP Agent状态"},
			Unit:                      nil,
			ValueMapping:              &entityStatus,
			Legend:                    "{deviceName} SNMP Agent Status",
			DefaultQueryRangeFunction: "min_over_time",
		},
		MetricNameEnum(Uptime): {
			Name:                      MetricNameEnum(Uptime),
			Description:               i18n.I18n{En: "Uptime", Zh: "运行时间"},
			Unit:                      &seconds,
			ValueMapping:              nil,
			Legend:                    "{deviceName} Uptime",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(ApStatus): {
			Name:                      MetricNameEnum(ApStatus),
			Description:               i18n.I18n{En: "AP Status", Zh: "AP运行状态"},
			Unit:                      nil,
			ValueMapping:              &opStatus,
			Legend:                    "{apName} Status",
			DefaultQueryRangeFunction: "min_over_time",
		},
		MetricNameEnum(ApUptime): {
			Name:                      MetricNameEnum(ApUptime),
			Description:               i18n.I18n{En: "AP Uptime", Zh: "AP运行时间"},
			Unit:                      &seconds,
			ValueMapping:              nil,
			Legend:                    "{name} Uptime",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(ChannelUtilization): {
			Name:                      MetricNameEnum(ChannelUtilization),
			Description:               i18n.I18n{En: "Channel Utilization", Zh: "信道利用率"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{apName}(channel) Channel Utilization",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(ChannelNoise): {
			Name:                      MetricNameEnum(ChannelNoise),
			Description:               i18n.I18n{En: "Channel Noise", Zh: "信道噪声"},
			Unit:                      &dBm,
			ValueMapping:              nil,
			Legend:                    "{apName}(channel) Channel Noise",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(ChannelTransmitPower): {
			Name:                      MetricNameEnum(ChannelTransmitPower),
			Description:               i18n.I18n{En: "Channel Transmit Power", Zh: "信道发射功率"},
			Unit:                      &dBm,
			ValueMapping:              nil,
			Legend:                    "{apName}(channel) Channel Transmit Power",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(ChannelInterferenceRate): {
			Name:                      MetricNameEnum(ChannelInterferenceRate),
			Description:               i18n.I18n{En: "Channel Interference", Zh: "信道干扰"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{apName}(channel) Channel Interference",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(ChannelAssociationClients): {
			Name:                      MetricNameEnum(ChannelAssociationClients),
			Description:               i18n.I18n{En: "Channel Association Clients", Zh: "信道关联客户端"},
			Unit:                      nil,
			ValueMapping:              nil,
			Legend:                    "{apName}(channel) Channel Association Clients",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(ApCpuUtilization): {
			Name:                      MetricNameEnum(ApCpuUtilization),
			Description:               i18n.I18n{En: "AP CPU utilization", Zh: "AP CPU利用率"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{apName} AP CPU utilization",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(ApMemoryUtilization): {
			Name:                      MetricNameEnum(ApMemoryUtilization),
			Description:               i18n.I18n{En: "AP Memory utilization", Zh: "AP 内存利用率"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{apName} AP Memory utilization",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(LoadAverage1m): {
			Name:                      MetricNameEnum(LoadAverage1m),
			Description:               i18n.I18n{En: "Load average 1m", Zh: "负载平均值 1m"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{serverName} Load average 1m",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(LoadAverage5m): {
			Name:                      MetricNameEnum(LoadAverage5m),
			Description:               i18n.I18n{En: "Load average 5m", Zh: "负载平均值 5m"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{serverName} Load average 5m",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(LoadAverage15m): {
			Name:                      MetricNameEnum(LoadAverage15m),
			Description:               i18n.I18n{En: "Load average 15m", Zh: "负载平均值 15m"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{serverName} Load average 15m",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(ContextSwitchesPerSecond): {
			Name:                      MetricNameEnum(ContextSwitchesPerSecond),
			Description:               i18n.I18n{En: "Context Switches per second", Zh: "上下文切换频率"},
			Unit:                      nil,
			ValueMapping:              nil,
			Legend:                    "{serverName} Context Switches per second",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(InterruptsPerSecond): {
			Name:                      MetricNameEnum(InterruptsPerSecond),
			Description:               i18n.I18n{En: "Interrupts per second", Zh: "中断频率"},
			Unit:                      nil,
			ValueMapping:              nil,
			Legend:                    "{serverName} Interrupts per second",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(SwapSpaceUtilization): {
			Name:                      MetricNameEnum(SwapSpaceUtilization),
			Description:               i18n.I18n{En: "Swap Space Utilization", Zh: "交换空间占用率"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{serverName} Swap Space Utilization",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(DiskReadRate): {
			Name:                      MetricNameEnum(DiskReadRate),
			Description:               i18n.I18n{En: "Disk Read Rate", Zh: "磁盘读取速率"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{serverName} Disk Read Rate",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(DiskWriteRate): {
			Name:                      MetricNameEnum(DiskWriteRate),
			Description:               i18n.I18n{En: "Disk Write Rate", Zh: "磁盘写入速率"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{serverName} Disk Write Rate",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(DiskUtilization): {
			Name:                      MetricNameEnum(DiskUtilization),
			Description:               i18n.I18n{En: "Disk Utilization", Zh: "磁盘利用率"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{serverName} Disk Utilization",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(InodesFree): {
			Name:                      MetricNameEnum(InodesFree),
			Description:               i18n.I18n{En: "Inodes Free", Zh: "inode空闲数量"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{serverName} Inodes Free",
			DefaultQueryRangeFunction: "max_over_time",
		},
		MetricNameEnum(FileSystemUtilization): {
			Name:                      MetricNameEnum(FileSystemUtilization),
			Description:               i18n.I18n{En: "File System Utilization", Zh: "文件系统利用率"},
			Unit:                      &percent,
			ValueMapping:              nil,
			Legend:                    "{serverName} {filesystem} File System Utilization",
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
