package alerts

import (
	"github.com/samber/lo"
	"github.com/wangxin688/narvis/intend/model/i18n"
)

type AlertNameEnum string

type SeverityEnum string

const (
	SeverityDisaster SeverityEnum = "P1"
	SeverityCritical SeverityEnum = "P2"
	SeverityWarning  SeverityEnum = "P3"
	SeverityInfo     SeverityEnum = "P4"
)

const (
	HighCpuUtilization       AlertNameEnum = "high_cpu_utilization"
	HighMemoryUtilization    AlertNameEnum = "high_memory_utilization"
	HighDiskUtilization      AlertNameEnum = "high_disk_utilization"
	HighSystemLoad           AlertNameEnum = "high_system_load"
	HighChannelUtilization   AlertNameEnum = "high_channel_utilization"
	HighChannelInterference  AlertNameEnum = "high_channel_interference"
	HighChannelNoise         AlertNameEnum = "high_channel_noise"
	HighClientNumber         AlertNameEnum = "high_client_number"
	HighBandwidthUtilization AlertNameEnum = "high_bandwidth_usage"
	HighErrorRate            AlertNameEnum = "high_error_rate"
	HighICMPLatency          AlertNameEnum = "high_icmp_ping_response_time"
	HighICMPPacketLoss       AlertNameEnum = "high_icmp_ping_loss"
	HighTemperature          AlertNameEnum = "high_temperature"
	AbnormalFanStatus        AlertNameEnum = "fan_status_abnormal"
	AbnormalPowerStatus      AlertNameEnum = "power_status_abnormal"
	InterfaceDown            AlertNameEnum = "interface_down"
	InterfaceHalfDuplex      AlertNameEnum = "interface_half_duplex"
	SnmpAgentTimeout         AlertNameEnum = "no_snmp_data_collection"
	NodePingTimeout          AlertNameEnum = "unavailable_by_icmp_ping"
	ApDown                   AlertNameEnum = "wireless_access_point_down"
	Unknown                  AlertNameEnum = "unknown"
	HighSwapSpaceUtilization AlertNameEnum = "high_swap_space_utilization"
	HighSystemLoadAverage    AlertNameEnum = "high_system_load_average"
	HighINodeUtilization     AlertNameEnum = "high_inode_utilization"
)

type AlertName struct {
	Name        AlertNameEnum
	Title       i18n.I18n
	Description i18n.I18n
	Suggestion  i18n.I18n
	Severity    SeverityEnum
}

func getAlertNameMeta() map[AlertNameEnum]AlertName {

	alertNameMeta := map[AlertNameEnum]AlertName{

		HighCpuUtilization: {
			Name:  HighCpuUtilization,
			Title: i18n.I18n{En: "High CPU utilization", Zh: "CPU利用率过高"},
			Suggestion: i18n.I18n{
				En: "1. Check device load\n2. Check device process cpu utilization\n3. check network traffic load",
				Zh: "检查设备负载\n2. 检查设备进程CPU利用率(cpu/memory)\n3. 检查网络流量"},
			Severity: SeverityWarning,
		},

		HighMemoryUtilization: {
			Name:  HighMemoryUtilization,
			Title: i18n.I18n{En: "High memory utilization", Zh: "内存利用率过高"},
			Suggestion: i18n.I18n{
				En: "1.Check device load\n2.Check device process memory utilization\n3.check memory leak",
				Zh: "检查设备负载\n2. 检查设备进程内存利用率(cpu/memory)\n3. 检查内存泄漏"},
			Severity: SeverityWarning,
		},

		HighDiskUtilization: {
			Name:       HighDiskUtilization,
			Title:      i18n.I18n{En: "High disk utilization", Zh: "磁盘利用率过高"},
			Suggestion: i18n.I18n{En: "1. Check device process disk utilization", Zh: "1. 检查设备磁盘利用率"},
			Severity:   SeverityWarning,
		},
		HighSystemLoad: {
			Name:  HighSystemLoad,
			Title: i18n.I18n{En: "High system load", Zh: "系统负载过高"},
			Suggestion: i18n.I18n{
				En: "1.Check device load\n2.Check device process cpu utilization\n3.check network traffic load",
				Zh: "检查设备负载\n2. 检查设备进程CPU利用率(cpu/memory)\n3. 检查网络流量负载"},
			Severity: SeverityWarning,
		},

		HighChannelUtilization: {
			Name:  HighChannelUtilization,
			Title: i18n.I18n{En: "High channel utilization", Zh: "信道利用率过高"},
			Suggestion: i18n.I18n{
				En: `1.Check co-channel status for the access points in same area\n
                 2.Check airtime utilization for current ap\n
                 3.Check history channel utilization for the area and review wlan performance\n`,
				Zh: `1. 检查AP是否存在同频干扰，确认相邻信道状态\n
                     2. 检查AP空口占用情况\n
                     3. 检查当前区域AP整体信道利用率和历史情况，确认是否是无线容量问题\n`,
			},
			Severity: SeverityInfo,
		},
		HighChannelNoise: {
			Name:       HighChannelNoise,
			Title:      i18n.I18n{En: "High channel noise", Zh: "信道噪声过高"},
			Suggestion: i18n.I18n{En: "1. Check office physical and electronical environment", Zh: "1. 检查办公室物理和电器设备环境"},
			Severity:   SeverityInfo,
		},
		HighChannelInterference: {
			Name:  HighChannelInterference,
			Title: i18n.I18n{En: "High channel interference", Zh: "信道干扰过高"},
			Suggestion: i18n.I18n{
				En: "1. Check office physical and electronical environment\n2. Check the channel monitor for current ap to see if hotpots or radar were founded.",
				Zh: "1. 检查办公室物理环境和电器设备\n2. 检查信道监控，查看AP周边是否存在热点或雷达干扰"},
			Severity: SeverityInfo,
		},
		HighClientNumber: {
			Name:  HighClientNumber,
			Title: i18n.I18n{En: "High client number", Zh: "客户端数量过多"},
			Suggestion: i18n.I18n{
				En: "1. Check Ap transmit power value, make sure it not too high and too low\n2.Check AP density for wlan design\n",
				Zh: "1. 检查AP发射功率是否设置合理\n2. 检查AP密度设计和WLAN点位设计是否合理\n"},
			Severity: SeverityInfo,
		},
		HighBandwidthUtilization: {
			Name:  HighBandwidthUtilization,
			Title: i18n.I18n{En: "High bandwidth utilization", Zh: "带宽利用率过高"},
			Suggestion: i18n.I18n{
				En: `1.check netflow data to find out bandwidth utilization\n
                     2.check history utilization data of the interface to check capacity plan)`,
				Zh: `1. 查看netflow/sflow流量数据，查看带宽利用率使用情况\n
					 2. 检查历史利用率数据，查看带宽利用计划`,
			},
			Severity: SeverityInfo,
		},
		HighErrorRate: {
			Name:       HighErrorRate,
			Title:      i18n.I18n{En: "High error rate", Zh: "错包率过高"},
			Suggestion: i18n.I18n{En: "1.Check interface CRC error to issue physical layer\n2.Check traffic load\n", Zh: "1. 检查接口CRC错包,排查物理层连接问题\n2. 检查流量负载"},
			Severity:   SeverityInfo,
		},
		HighICMPLatency: {
			Name:  HighICMPLatency,
			Title: i18n.I18n{En: "High ICMP latency", Zh: "ICMP延迟过高"},
			Suggestion: i18n.I18n{
				En: `1.Check internet/intranet network packet loss rate\n
                	 2.Check traffic status for inter-connective devices\n
                	 3.Check device load(cpu/memory/process)`,
				Zh: "1. 检查内外网丢包状态\n2. 检查互联端口流量利用率是否过高\n3. 检查网络设备负载",
			},
			Severity: SeverityInfo,
		},
		HighICMPPacketLoss: {
			Name:  HighICMPPacketLoss,
			Title: i18n.I18n{En: "High ICMP packet loss", Zh: "ICMP丢包率过高"},
			Suggestion: i18n.I18n{
				En: `1.Check internet/intranet network packet loss rate\n
						 2.Check traffic status for inter-connective devices\n
						 3.Check device load(cpu/memory/process)`,
				Zh: "1. 检查内外网丢包状态\n2. 检查互联端口流量利用率是否过高\n3. 检查网络设备负载",
			},
			Severity: SeverityWarning,
		},
		HighTemperature: {
			Name:  HighTemperature,
			Title: i18n.I18n{En: "High temperature", Zh: "硬件温度过高"},
			Suggestion: i18n.I18n{
				En: "1.Check Cooling device/sensor status in ServerRoom\n2.Check device hardware status\n",
				Zh: "1. 检查机房冷却设备/传感器状态\n2. 检查设备硬件状态",
			},
			Severity: SeverityWarning,
		},
		AbnormalFanStatus: {
			Name:  AbnormalFanStatus,
			Title: i18n.I18n{En: "Abnormal fan status", Zh: "风扇状态异常"},
			Suggestion: i18n.I18n{
				En: "1.Check Cooling device/sensor status in ServerRoom\n2.Check device hardware status\n",
				Zh: "1. 检查机房冷却设备/传感器状态\n2. 检查设备硬件状态",
			},
			Severity: SeverityWarning,
		},
		AbnormalPowerStatus: {
			Name:  AbnormalPowerStatus,
			Title: i18n.I18n{En: "Abnormal power status", Zh: "电源状态异常"},
			Suggestion: i18n.I18n{
				En: "1.Check UPS/Power status in ServerRoom\n2.Check device hardware status, powerlets connectivity\n",
				Zh: "1. 检查机房电力以及UPS供电状态\n2. 检查设备硬件状态,电源连接状态",
			},
			Severity: SeverityWarning,
		},
		InterfaceDown: {
			Name:  InterfaceDown,
			Title: i18n.I18n{En: "Interface down", Zh: "接口状态异常"},
			Suggestion: i18n.I18n{
				En: `1.check interface cable physical connection status\n
                     2.check device interface hardware(interface/optical-transceiver)`,
				Zh: "1. 检查设备物理连线状态\n2. 检查节点接口硬件(端口、光模块)状态",
			},
			Severity: SeverityWarning,
		},
		InterfaceHalfDuplex: {
			Name:  InterfaceHalfDuplex,
			Title: i18n.I18n{En: "Interface mode in half duplex", Zh: "接口双工状态异常"},
			Suggestion: i18n.I18n{
				En: `1.check interface cable physical connection status\n
                     2.check device interface hardware(interface/optical-transceiver)`,
				Zh: "1. 检查设备物理连线状态\n2. 检查节点接口硬件(端口、光模块)状态",
			},
			Severity: SeverityWarning,
		},
		SnmpAgentTimeout: {
			Name:  SnmpAgentTimeout,
			Title: i18n.I18n{En: "SNMP agent timeout", Zh: "SNMP采集超时"},
			Suggestion: i18n.I18n{
				En: `1.Check node connectivity status\n
                 2.Check SNMP configuration\n
                 3.Check node system load(cpu/memory/process)`,
				Zh: "1. 检查节点连通性是否正常\n2. 检查节点SNMP配置\n3. 检查节点系统负载(cpu/memory/process)",
			},
			Severity: SeverityWarning,
		},
		NodePingTimeout: {
			Name:  NodePingTimeout,
			Title: i18n.I18n{En: "Node ping timeout", Zh: "节点Ping超时"},
			Suggestion: i18n.I18n{
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
			Title: i18n.I18n{En: "AP down", Zh: "AP离线"},
			Suggestion: i18n.I18n{
				En: "1.Check AP connectivity\n2.Check AP hardware status\n",
				Zh: "1. 检查AP连通性\n2. 检查AP硬件状态",
			},
			Severity: SeverityWarning,
		},
		Unknown: {
			Name:  Unknown,
			Title: i18n.I18n{En: "Unknown", Zh: "未知"},
			Suggestion: i18n.I18n{
				En: "1.Unknown",
				Zh: "1. 未知",
			},
			Severity: SeverityInfo,
		},
		HighSwapSpaceUtilization: {
			Name:  HighSwapSpaceUtilization,
			Title: i18n.I18n{En: "High swap space utilization", Zh: "swap空间占用率高"},
			Suggestion: i18n.I18n{
				En: "1.Check swap space utilization",
				Zh: "1. 检查swap空间占用率",
			},
			Severity: SeverityWarning,
		},
		HighSystemLoadAverage: {
			Name:  HighSystemLoadAverage,
			Title: i18n.I18n{En: "High system load average", Zh: "系统负载平均值高"},
			Suggestion: i18n.I18n{
				En: "1.Check system load average",
				Zh: "1. 检查系统负载平均值",
			},
			Severity: SeverityWarning,
		},
		HighINodeUtilization: {
			Name:  HighINodeUtilization,
			Title: i18n.I18n{En: "High inode utilization", Zh: "Inode利用率高"},
			Suggestion: i18n.I18n{
				En: "1.Check inode utilization",
				Zh: "1. 检查inode占用率",
			},
			Severity: SeverityWarning,
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

func GetAlertEnumNames() []AlertNameEnum {
	return []AlertNameEnum{
		HighCpuUtilization,
		HighMemoryUtilization,
		HighDiskUtilization,
		HighSystemLoad,
		HighChannelUtilization,
		HighChannelInterference,
		HighChannelNoise,
		HighClientNumber,
		HighBandwidthUtilization,
		HighErrorRate,
		HighICMPLatency,
		HighICMPPacketLoss,
		HighTemperature,
		AbnormalFanStatus,
		AbnormalPowerStatus,
		InterfaceDown,
		InterfaceHalfDuplex,
		SnmpAgentTimeout,
		NodePingTimeout,
		ApDown,
		Unknown,
		HighSwapSpaceUtilization,
		HighSystemLoadAverage,
		HighINodeUtilization,
	}
}

func GetInterfaceAlertEnumNames() []AlertNameEnum {
	return []AlertNameEnum{
		InterfaceDown,
		InterfaceHalfDuplex,
		HighBandwidthUtilization,
		HighErrorRate,
	}
}

func GetApAlertEnumNames() []AlertNameEnum {
	return []AlertNameEnum{
		ApDown,
		HighChannelUtilization,
		HighChannelInterference,
		HighChannelNoise,
		HighClientNumber,
	}
}
