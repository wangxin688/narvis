package alerts

import (
	"github.com/samber/lo"
	"github.com/wangxin688/narvis/intend/common"
	"github.com/wangxin688/narvis/intend/metrics"
)

type AlertNameEnum string

type SeverityEnum string

const (
	SeverityDisaster SeverityEnum = "DISASTER"
	SeverityCritical SeverityEnum = "CRITICAL"
	SeverityWarning  SeverityEnum = "WARNING"
	SeverityInfo     SeverityEnum = "INFO"
)

const (
	HighCpuUsage            AlertNameEnum = "high_cpu_usage"
	HighMemoryUsage         AlertNameEnum = "high_memory_usage"
	HighDiskUsage           AlertNameEnum = "high_disk_usage"
	HighSystemLoad          AlertNameEnum = "high_system_load"
	HighChannelUsage        AlertNameEnum = "high_channel_usage"
	HighChannelInterference AlertNameEnum = "high_channel_interference"
	HighChannelNoise        AlertNameEnum = "high_channel_noise"
	HighClientNumber        AlertNameEnum = "high_client_number"
	HighBandwidthUsage      AlertNameEnum = "high_bandwidth_usage"
	HighErrorRate           AlertNameEnum = "high_error_rate"
	HighICMPLatency         AlertNameEnum = "high_icmp_latency"
	HighICMPPacketLoss      AlertNameEnum = "high_icmp_packet_loss"
	HighTemperature         AlertNameEnum = "high_temperature"
	AbnormalFanStatus       AlertNameEnum = "abnormal_fan_status"
	AbnormalPowerStatus     AlertNameEnum = "abnormal_power_status"
	InterfaceDown           AlertNameEnum = "interface_down"
	SnmpAgentTimeout        AlertNameEnum = "snmp_agent_timeout"
	NodePingTimeout         AlertNameEnum = "node_ping_timeout"
	ApDown                  AlertNameEnum = "ap_down"
	Unknown                 AlertNameEnum = "unknown"
)

type AlertName struct {
	Name        AlertNameEnum
	Title       common.I18n
	Description common.I18n
	Suggestion  common.I18n
	Severity    SeverityEnum
}

func getAlertNameMeta() map[AlertNameEnum]AlertName {

	alertNameMeta := map[AlertNameEnum]AlertName{

		HighCpuUsage: {
			Name:  HighCpuUsage,
			Title: common.I18n{En: "High CPU usage", Zh: "CPU利用率过高"},
			Suggestion: common.I18n{
				En: "1. Check device load\n2. Check device process cpu utilization\n3. check network traffic load",
				Zh: "检查设备负载\n2. 检查设备进程CPU利用率(cpu/memory)\n3. 检查网络流量"},
			Severity: SeverityWarning,
		},

		HighMemoryUsage: {
			Name:  HighMemoryUsage,
			Title: common.I18n{En: "High memory usage", Zh: "内存利用率过高"},
			Suggestion: common.I18n{
				En: "1.Check device load\n2.Check device process memory utilization\n3.check memory leak",
				Zh: "检查设备负载\n2. 检查设备进程内存利用率(cpu/memory)\n3. 检查内存泄漏"},
			Severity: SeverityWarning,
		},

		HighDiskUsage: {
			Name:       HighDiskUsage,
			Title:      common.I18n{En: "High disk usage", Zh: "磁盘利用率过高"},
			Suggestion: common.I18n{En: "1. Check device process disk utilization", Zh: "1. 检查设备磁盘利用率"},
			Severity:   SeverityWarning,
		},
		HighSystemLoad: {
			Name:  HighSystemLoad,
			Title: common.I18n{En: "High system load", Zh: "系统负载过高"},
			Suggestion: common.I18n{
				En: "1.Check device load\n2.Check device process cpu utilization\n3.check network traffic load",
				Zh: "检查设备负载\n2. 检查设备进程CPU利用率(cpu/memory)\n3. 检查网络流量负载"},
			Severity: SeverityWarning,
		},

		HighChannelUsage: {
			Name:  HighChannelUsage,
			Title: common.I18n{En: "High channel usage", Zh: "信道利用率过高"},
			Suggestion: common.I18n{
				En: `1.Check co-channel status for the access points in same area\n
                 2.Check airtime usage for current ap\n
                 3.Check history channel usage for the area and review wlan performance\n`,
				Zh: `1. 检查AP是否存在同频干扰，确认相邻信道状态\n
                     2. 检查AP空口占用情况\n
                     3. 检查当前区域AP整体信道利用率和历史情况，确认是否是无线容量问题\n`,
			},
			Severity: SeverityInfo,
		},
		HighChannelNoise: {
			Name:       HighChannelNoise,
			Title:      common.I18n{En: "High channel noise", Zh: "信道噪声过高"},
			Suggestion: common.I18n{En: "1. Check office physical and electronical environment", Zh: "1. 检查办公室物理和电器设备环境"},
			Severity:   SeverityInfo,
		},
		HighChannelInterference: {
			Name:  HighChannelInterference,
			Title: common.I18n{En: "High channel interference", Zh: "信道干扰过高"},
			Suggestion: common.I18n{
				En: "1. Check office physical and electronical environment\n2. Check the channel monitor for current ap to see if hotpots or radar were founded.",
				Zh: "1. 检查办公室物理环境和电器设备\n2. 检查信道监控，查看AP周边是否存在热点或雷达干扰"},
			Severity: SeverityInfo,
		},
		HighClientNumber: {
			Name:  HighClientNumber,
			Title: common.I18n{En: "High client number", Zh: "客户端数量过多"},
			Suggestion: common.I18n{
				En: "1. Check Ap transmit power value, make sure it not too high and too low\n2.Check AP density for wlan design\n",
				Zh: "1. 检查AP发射功率是否设置合理\n2. 检查AP密度设计和WLAN点位设计是否合理\n"},
			Severity: SeverityInfo,
		},
		HighBandwidthUsage: {
			Name:  HighBandwidthUsage,
			Title: common.I18n{En: "High bandwidth usage", Zh: "带宽利用率过高"},
			Suggestion: common.I18n{
				En: `1.check netflow data to find out bandwidth usage\n
                     2.check history utilization data of the interface to check capacity plan)`,
				Zh: `1. 查看netflow/sflow流量数据，查看带宽利用率使用情况\n
					 2. 检查历史利用率数据，查看带宽利用计划`,
			},
			Severity: SeverityInfo,
		},
		HighErrorRate: {
			Name:       HighErrorRate,
			Title:      common.I18n{En: "High error rate", Zh: "错包率过高"},
			Suggestion: common.I18n{En: "1.Check interface CRC error to issue physical layer\n2.Check traffic load\n", Zh: "1. 检查接口CRC错包,排查物理层连接问题\n2. 检查流量负载"},
			Severity:   SeverityInfo,
		},
		HighICMPLatency: {
			Name:  HighICMPLatency,
			Title: common.I18n{En: "High ICMP latency", Zh: "ICMP延迟过高"},
			Suggestion: common.I18n{
				En: `1.Check internet/intranet network packet loss rate\n
                	 2.Check traffic status for inter-connective devices\n
                	 3.Check device load(cpu/memory/process)`,
				Zh: "1. 检查内外网丢包状态\n2. 检查互联端口流量利用率是否过高\n3. 检查网络设备负载",
			},
			Severity: SeverityInfo,
		},
		HighICMPPacketLoss: {
			Name:  HighICMPPacketLoss,
			Title: common.I18n{En: "High ICMP packet loss", Zh: "ICMP丢包率过高"},
			Suggestion: common.I18n{
				En: `1.Check internet/intranet network packet loss rate\n
						 2.Check traffic status for inter-connective devices\n
						 3.Check device load(cpu/memory/process)`,
				Zh: "1. 检查内外网丢包状态\n2. 检查互联端口流量利用率是否过高\n3. 检查网络设备负载",
			},
			Severity: SeverityWarning,
		},
		HighTemperature: {
			Name:  HighTemperature,
			Title: common.I18n{En: "High temperature", Zh: "硬件温度过高"},
			Suggestion: common.I18n{
				En: "1.Check Cooling device/sensor status in ServerRoom\n2.Check device hardware status\n",
				Zh: "1. 检查机房冷却设备/传感器状态\n2. 检查设备硬件状态",
			},
			Severity: SeverityWarning,
		},
		AbnormalFanStatus: {
			Name:  AbnormalFanStatus,
			Title: common.I18n{En: "Abnormal fan status", Zh: "风扇状态异常"},
			Suggestion: common.I18n{
				En: "1.Check Cooling device/sensor status in ServerRoom\n2.Check device hardware status\n",
				Zh: "1. 检查机房冷却设备/传感器状态\n2. 检查设备硬件状态",
			},
			Severity: SeverityWarning,
		},
		AbnormalPowerStatus: {
			Name:  AbnormalPowerStatus,
			Title: common.I18n{En: "Abnormal power status", Zh: "电源状态异常"},
			Suggestion: common.I18n{
				En: "1.Check UPS/Power status in ServerRoom\n2.Check device hardware status, powerlets connectivity\n",
				Zh: "1. 检查机房电力以及UPS供电状态\n2. 检查设备硬件状态,电源连接状态",
			},
			Severity: SeverityWarning,
		},
		InterfaceDown: {
			Name:  InterfaceDown,
			Title: common.I18n{En: "Interface down", Zh: "接口状态异常"},
			Suggestion: common.I18n{
				En: `1.check interface cable physical connection status\n
                     2.check device interface hardware(interface/optical-transceiver)`,
				Zh: "1. 检查设备物理连线状态\n2. 检查节点接口硬件(端口、光模块)状态",
			},
			Severity: SeverityWarning,
		},
		SnmpAgentTimeout: {
			Name:  SnmpAgentTimeout,
			Title: common.I18n{En: "SNMP agent timeout", Zh: "SNMP采集超时"},
			Suggestion: common.I18n{
				En: `1.Check node connectivity status\n
                 2.Check SNMP configuration\n
                 3.Check node system load(cpu/memory/process)`,
				Zh: "1. 检查节点连通性是否正常\n2. 检查节点SNMP配置\n3. 检查节点系统负载(cpu/memory/process)",
			},
			Severity: SeverityWarning,
		},
		NodePingTimeout: {
			Name:  NodePingTimeout,
			Title: common.I18n{En: "Node ping timeout", Zh: "节点Ping超时"},
			Suggestion: common.I18n{
				En: `
                1.Check internet/intranet network connectivity\n
                2.Check ip routing reachability\n
                3.Check device hardware/power availability`,
				Zh: "1. 检查内外网连通性是否正常\n2. 检查IP路由是否可达\n3. 检查设备硬件/电源是否可用",
			},
			Severity: SeverityCritical,
		},
		ApDown: {
			Name:  ApDown,
			Title: common.I18n{En: "AP down", Zh: "AP离线"},
			Suggestion: common.I18n{
				En: "1.Check AP connectivity\n2.Check AP hardware status\n",
				Zh: "1. 检查AP连通性\n2. 检查AP硬件状态",
			},
			Severity: SeverityWarning,
		},
		Unknown: {
			Name:  Unknown,
			Title: common.I18n{En: "Unknown", Zh: "未知"},
			Suggestion: common.I18n{
				En: "1.Unknown",
				Zh: "1. 未知",
			},
			Severity: SeverityInfo,
		},
	}
	return alertNameMeta
}

func GetListALertName() (int, []AlertName) {

	alertNameMeta := getAlertNameMeta()
	return len(alertNameMeta), lo.Values(alertNameMeta)
}

func GetAlertName(alertName string) AlertName {
	alertNameMeta := getAlertNameMeta()

	if result, ok := alertNameMeta[AlertNameEnum(alertName)]; ok {
		return result
	}

	return getAlertNameMeta()[Unknown]
}
