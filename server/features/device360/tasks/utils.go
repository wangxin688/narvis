package device360_tasks

import (
	"strconv"

	"github.com/samber/lo"
)

const channelBusyThreshold float64 = 70
const ChannelAssociationClientsThreshold float64 = 70
const cpuHighThreshold float64 = 90
const memHighThreshold float64 = 90
const temperatureHighThreshold float64 = 75
const packErrDiscardThreshold float64 = 10
const trafficHighThreshold float64 = 90
const trafficCriticalThreshold float64 = 98

func calcIcmpScore(icmp float64) float64 {
	if icmp == -1 {
		return -1
	}
	if icmp == 1 {
		return 10
	}
	return 0
}

func calcApStatusScore(status float64) float64 {
	if status == -1 {
		return -1
	}
	if status == 1 {
		return 10
	}
	return 0
}

func calcApRadioScore(chanBusy []float64, chanClients []float64) float64 {
	var score float64 = 10
	if len(chanBusy) == 0 || len(chanClients) == 0 {
		return -1
	}
	if len(chanBusy) > 0 && lo.Max(chanBusy) >= channelBusyThreshold {
		score -= 4
	}
	if len(chanClients) > 0 && lo.Max(chanClients) >= ChannelAssociationClientsThreshold {
		score -= 3
	}
	return score
}

func calcCpuScore(cpuUsage []float64) float64 {
	if len(cpuUsage) == 0 {
		return -1
	}
	if lo.Max(cpuUsage) >= cpuHighThreshold {
		return 1
	}
	return 10
}

func calcMemScore(memUsage []float64) float64 {
	if len(memUsage) == 0 {
		return -1
	}
	if lo.Max(memUsage) >= memHighThreshold {
		return 1
	}
	return 10
}

func calcTemperatureScore(temperature []float64) float64 {
	if len(temperature) == 0 {
		return -1
	}
	if lo.Max(temperature) >= temperatureHighThreshold {
		return 8
	}
	return 10
}

func calcFanSore(fanStatus []float64) (score float64, anomalyCount int) {
	if len(fanStatus) == 0 {
		return -1, 0
	}
	if lo.Max(fanStatus) == 0 {
		return 10, 0
	}
	return 5, lo.Count(fanStatus, 2)
}

func calcPowerSupplyScore(powerSupplyStatus []float64) (score float64, anomalyCount int) {
	if len(powerSupplyStatus) == 0 {
		return -1, 0
	}
	if lo.Max(powerSupplyStatus) == 0 {
		return 10, 0
	}
	return 5, lo.Count(powerSupplyStatus, 2)
}

func calcIfPacketScore(packets []float64) (score float64, anomalyCount int) {
	if len(packets) == 0 {
		return -1, 0
	}
	for _, packet := range packets {
		if packet >= packErrDiscardThreshold {
			anomalyCount++
		}
	}
	if anomalyCount == 0 {
		return 10, 0
	}
	if anomalyCount == len(packets) {
		return 1, anomalyCount
	}
	return 8, anomalyCount
}

func calcIfTrafficScore(traffic []float64) (score float64, anomalyCount int) {
	if len(traffic) == 0 {
		return -1, 0
	}
	abLevel1 := 0
	abLevel2 := 0
	for _, item := range traffic {
		if item < trafficHighThreshold {
			continue
		}
		if trafficHighThreshold <= item && item <= trafficCriticalThreshold {
			abLevel1++
		}
		if item > trafficCriticalThreshold {
			abLevel2++
		}
	}
	if abLevel2 > 0 {
		return 1, abLevel2 + abLevel1
	}
	if abLevel1 > 0 {
		return 8, abLevel1
	}
	return 10, 0
}

func calcIfOpStatusScore(opStatus []float64) (score float64, anomalyCount int) {
	if len(opStatus) == 0 {
		return -1, 0
	}
	for _, item := range opStatus {
		if item == 0 {
			anomalyCount++
		}
	}
	if anomalyCount == 0 {
		return 10, 0
	}
	if anomalyCount == len(opStatus) {
		return 1, anomalyCount
	}
	return 8, anomalyCount
}

func stringToFloat64(str string) float64 {
	result, _ := strconv.ParseFloat(str, 64)
	return result
}
