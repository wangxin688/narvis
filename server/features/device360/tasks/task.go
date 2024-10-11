package device360_tasks

import (
	"fmt"
	"time"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/intend/metrics"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/pkg/vtm"
	"go.uber.org/zap"
)

func getQueries() map[string]string {
	queries := map[string]string{
		string(metrics.ICMPPing):                  fmt.Sprintf("last_over_time(%s[3m])", metrics.ICMPPing),
		string(metrics.CpuUsage):                  fmt.Sprintf("last_over_time(%s[5m])", metrics.CpuUsage),
		string(metrics.MemoryUsage):               fmt.Sprintf("last_over_time(%s[5m])", metrics.MemoryUsage),
		string(metrics.Temperature):               fmt.Sprintf("last_over_time(%s[5m])", metrics.Temperature),
		string(metrics.FanStatus):                 fmt.Sprintf("last_over_time(%s[5m])", metrics.FanStatus),
		string(metrics.PowerSupplyStatus):         fmt.Sprintf("last_over_time(%s[5m])", metrics.PowerSupplyStatus),
		string(metrics.RxDiscards):                fmt.Sprintf("last_over_time(%s[5m])", metrics.RxDiscards),
		string(metrics.TxDiscards):                fmt.Sprintf("last_over_time(%s[5m])", metrics.TxDiscards),
		string(metrics.RxErrors):                  fmt.Sprintf("last_over_time(%s[5m])", metrics.RxErrors),
		string(metrics.TxErrors):                  fmt.Sprintf("last_over_time(%s[5m])", metrics.TxErrors),
		string(metrics.RxRate):                    fmt.Sprintf("last_over_time(%s[5m])", metrics.RxRate),
		string(metrics.TxRate):                    fmt.Sprintf("last_over_time(%s[5m])", metrics.TxRate),
		string(metrics.OperationalStatus):         fmt.Sprintf("last_over_time(%s[5m])", metrics.OperationalStatus),
		string(metrics.ChannelUtilization):        fmt.Sprintf("last_over_time(%s[5m])", metrics.ChannelUtilization),
		string(metrics.ChannelAssociationClients): fmt.Sprintf("last_over_time(%s[5m])", metrics.ChannelAssociationClients),
		string(metrics.ApStatus):                  fmt.Sprintf("last_over_time(%s[5m])", metrics.ApStatus),
	}
	return queries
}

func queryResults(queries map[string]string) (deviceVectors, apVectors map[string][]*vtm.VectorResponse, err error) {
	if queries == nil {
		core.Logger.Error("[device360Task]: getQueries result is nil")
		return nil, nil, fmt.Errorf("queries is nil")
	}

	vectorRequests := make([]*vtm.VectorRequest, 0)
	for _, v := range queries {
		vectorRequests = append(vectorRequests, &vtm.VectorRequest{
			Step: 180, Query: v})
	}
	resp, err := vtm.NewVtmClient().GetBulkVector(vectorRequests, nil)
	if err != nil {
		return nil, nil, err
	}

	if resp == nil {
		return nil, nil, fmt.Errorf("vtm response is nil")
	}

	for _, v := range resp {
		metricName := v.Metric["__name__"]
		if _, ok := v.Metric["deviceId"]; !ok {
			continue
		}

		if _, ok := v.Metric["apName"]; ok {
			if apVectors == nil {
				apVectors = make(map[string][]*vtm.VectorResponse)
			}

			if _, ok := apVectors[metricName]; !ok {
				apVectors[metricName] = make([]*vtm.VectorResponse, 0)
			}

			apVectors[metricName] = append(apVectors[metricName], v)
		} else {
			if deviceVectors == nil {
				deviceVectors = make(map[string][]*vtm.VectorResponse)
			}

			if _, ok := deviceVectors[metricName]; !ok {
				deviceVectors[metricName] = make([]*vtm.VectorResponse, 0)
			}

			deviceVectors[metricName] = append(deviceVectors[metricName], v)
		}
	}
	return deviceVectors, apVectors, nil
}

func aggregateApMetrics(vectors map[string][]*vtm.VectorResponse) map[string]map[string]*ApSchema {
	apMetrics := make(map[string]map[string]*ApSchema)
	for metricName, metricList := range vectors {
		if len(metricList) == 0 {
			continue
		}
		for _, item := range metricList {
			siteId := item.Metric["siteId"]
			apName := item.Metric["apName"]
			if siteId == "" || apName == "" {
				core.Logger.Warn("[device360Task]: siteId or apName is empty", zap.Any("metric", item))
				continue
			}
			delete(item.Metric, "__name__")
			delete(item.Metric, "deviceId")
			delete(item.Metric, "deviceName")
			delete(item.Metric, "channel")
			delete(item.Metric, "radioType")
			if _, ok := apMetrics[siteId]; !ok {
				apMetrics[siteId] = make(map[string]*ApSchema)
			}
			if _, ok := apMetrics[siteId][apName]; !ok {
				apMetrics[siteId][apName] = &ApSchema{
					ChannelUtilization:        make([]float64, 0),
					ChannelAssociationClients: make([]float64, 0),
					ApStatus:                  -1,
					Labels:                    item.Metric,
				}
			}
			if metricName == string(metrics.ApStatus) {
				apMetrics[siteId][apName].ApStatus = stringToFloat64(item.Value[1].(string))
			} else if metricName == string(metrics.ChannelUtilization) {
				apMetrics[siteId][apName].ChannelUtilization = append(apMetrics[siteId][apName].ChannelUtilization, stringToFloat64(item.Value[1].(string)))
			} else if metricName == string(metrics.ChannelAssociationClients) {
				apMetrics[siteId][apName].ChannelAssociationClients = append(apMetrics[siteId][apName].ChannelAssociationClients, stringToFloat64(item.Value[1].(string)))
			}
		}
	}
	return apMetrics
}

func aggregateDeviceMetrics(vectors map[string][]*vtm.VectorResponse) (
	deviceMetrics map[string]*DeviceSchema) {

	deviceMetrics = make(map[string]*DeviceSchema)

	for metricName, metricList := range vectors {
		if len(metricList) == 0 {
			continue
		}
		for _, item := range metricList {
			deviceId := item.Metric["deviceId"]

			if deviceId == "" {
				core.Logger.Warn("[device360Task]: deviceId is empty", zap.Any("metric", item))
				continue
			}

			delete(item.Metric, "__name__")

			if _, ok := deviceMetrics[deviceId]; !ok {
				deviceMetrics[deviceId] = &DeviceSchema{
					ICMPPing:          -1,
					CpuUsage:          make([]float64, 0),
					MemoryUsage:       make([]float64, 0),
					Temperature:       make([]float64, 0),
					FanStatus:         make([]float64, 0),
					PowerSupplyStatus: make([]float64, 0),
					RxDiscards:        make([]float64, 0),
					TxDiscards:        make([]float64, 0),
					RxErrors:          make([]float64, 0),
					TxErrors:          make([]float64, 0),
					RxRate:            make([]float64, 0),
					TxRate:            make([]float64, 0),
					OperationalStatus: make([]float64, 0),
					Labels:            item.Metric,
				}
			}

			switch metricName {
			case string(metrics.ICMPPing):
				deviceMetrics[deviceId].ICMPPing = stringToFloat64(item.Value[1].(string))
			case string(metrics.CpuUsage):
				deviceMetrics[deviceId].CpuUsage = append(deviceMetrics[deviceId].CpuUsage, stringToFloat64(item.Value[1].(string)))
			case string(metrics.MemoryUsage):
				deviceMetrics[deviceId].MemoryUsage = append(deviceMetrics[deviceId].MemoryUsage, stringToFloat64(item.Value[1].(string)))
			case string(metrics.Temperature):
				deviceMetrics[deviceId].Temperature = append(deviceMetrics[deviceId].Temperature, stringToFloat64(item.Value[1].(string)))
			case string(metrics.FanStatus):
				deviceMetrics[deviceId].FanStatus = append(deviceMetrics[deviceId].FanStatus, stringToFloat64(item.Value[1].(string)))
			case string(metrics.PowerSupplyStatus):
				deviceMetrics[deviceId].PowerSupplyStatus = append(deviceMetrics[deviceId].PowerSupplyStatus, stringToFloat64(item.Value[1].(string)))
			case string(metrics.RxDiscards):
				deviceMetrics[deviceId].RxDiscards = append(deviceMetrics[deviceId].RxDiscards, stringToFloat64(item.Value[1].(string)))
			case string(metrics.TxDiscards):
				deviceMetrics[deviceId].TxDiscards = append(deviceMetrics[deviceId].TxDiscards, stringToFloat64(item.Value[1].(string)))
			case string(metrics.RxErrors):
				deviceMetrics[deviceId].RxErrors = append(deviceMetrics[deviceId].RxErrors, stringToFloat64(item.Value[1].(string)))
			case string(metrics.TxErrors):
				deviceMetrics[deviceId].TxErrors = append(deviceMetrics[deviceId].TxErrors, stringToFloat64(item.Value[1].(string)))
			case string(metrics.RxRate):
				deviceMetrics[deviceId].RxRate = append(deviceMetrics[deviceId].RxRate, stringToFloat64(item.Value[1].(string)))
			case string(metrics.TxRate):
				deviceMetrics[deviceId].TxRate = append(deviceMetrics[deviceId].TxRate, stringToFloat64(item.Value[1].(string)))
			case string(metrics.OperationalStatus):
				deviceMetrics[deviceId].OperationalStatus = append(deviceMetrics[deviceId].OperationalStatus, stringToFloat64(item.Value[1].(string)))
			}
		}
	}
	return
}

func calcHealthScore(deviceMetrics map[string]*DeviceSchema, timestamp int64) []*vtm.Metric {
	var results []*vtm.Metric
	for _, device := range deviceMetrics {
		if device == nil {
			core.Logger.Warn("[device360Task]: device is nil")
			continue
		}
		icmpScore := calcIcmpScore(device.ICMPPing)
		cpuScore := calcCpuScore(device.CpuUsage)
		memoryScore := calcMemScore(device.MemoryUsage)
		tempScore := calcTemperatureScore(device.Temperature)
		fanScore, fanAnomaly := calcFanSore(device.FanStatus)
		powerSupplyScore, powerSupplyAnomaly := calcPowerSupplyScore(device.PowerSupplyStatus)
		rxDiscardScore, rxDiscardAnomaly := calcIfPacketScore(device.RxDiscards)
		txDiscardScore, txDiscardAnomaly := calcIfPacketScore(device.TxDiscards)
		rxErrorScore, rxErrorAnomaly := calcIfPacketScore(device.RxErrors)
		txErrorScore, txErrorAnomaly := calcIfPacketScore(device.TxErrors)
		rxRateScore, rxRateAnomaly := calcIfTrafficScore(device.RxRate)
		txRateScore, txRateAnomaly := calcIfTrafficScore(device.TxRate)
		operationalStatusScore, operationalStatusAnomaly := calcIfOpStatusScore(device.OperationalStatus)
		device360Score := lo.Min([]float64{icmpScore, cpuScore, memoryScore, tempScore, fanScore, powerSupplyScore, rxDiscardScore, txDiscardScore, rxErrorScore, txErrorScore, rxRateScore, txRateScore, operationalStatusScore})
		results = append(results, &vtm.Metric{
			Metric:    string(metrics.HealthScore),
			Labels:    device.Labels,
			Value:     device360Score,
			Timestamp: timestamp,
		})
		results = append(results, &vtm.Metric{
			Metric:    string(metrics.IcmpScore),
			Labels:    device.Labels,
			Value:     icmpScore,
			Timestamp: timestamp,
		})
		results = append(results, &vtm.Metric{
			Metric:    string(metrics.CpuScore),
			Labels:    device.Labels,
			Value:     cpuScore,
			Timestamp: timestamp,
		})
		results = append(results, &vtm.Metric{
			Metric:    string(metrics.MemoryScore),
			Labels:    device.Labels,
			Value:     memoryScore,
			Timestamp: timestamp,
		})
		results = append(results, &vtm.Metric{
			Metric:    string(metrics.TemperatureScore),
			Labels:    device.Labels,
			Value:     tempScore,
			Timestamp: timestamp,
		})
		results = append(results, &vtm.Metric{
			Metric:    string(metrics.FanScore),
			Labels:    device.Labels,
			Value:     fanScore,
			Timestamp: timestamp,
		})
		results = append(results, &vtm.Metric{
			Metric:    string(metrics.FanAnomaly),
			Labels:    device.Labels,
			Value:     float64(fanAnomaly),
			Timestamp: timestamp,
		})
		results = append(results, &vtm.Metric{
			Metric:    string(metrics.PowerSupplyScore),
			Labels:    device.Labels,
			Value:     powerSupplyScore,
			Timestamp: timestamp,
		})
		results = append(results, &vtm.Metric{
			Metric:    string(metrics.PowerAnomaly),
			Labels:    device.Labels,
			Value:     float64(powerSupplyAnomaly),
			Timestamp: timestamp,
		})
		results = append(results, &vtm.Metric{
			Metric:    string(metrics.IfErrorScore),
			Labels:    device.Labels,
			Value:     lo.Min([]float64{rxErrorScore, txErrorScore}),
			Timestamp: timestamp,
		})
		results = append(results, &vtm.Metric{
			Metric:    string(metrics.IfErrorAnomaly),
			Labels:    device.Labels,
			Value:     float64(rxErrorAnomaly) + float64(txErrorAnomaly),
			Timestamp: timestamp,
		})
		results = append(results, &vtm.Metric{
			Metric:    string(metrics.IfDiscardScore),
			Labels:    device.Labels,
			Value:     lo.Min([]float64{rxDiscardScore, txDiscardScore}),
			Timestamp: timestamp,
		})
		results = append(results, &vtm.Metric{
			Metric:    string(metrics.IfDiscardAnomaly),
			Labels:    device.Labels,
			Value:     float64(rxDiscardAnomaly) + float64(txDiscardAnomaly),
			Timestamp: timestamp,
		})
		results = append(results, &vtm.Metric{
			Metric:    string(metrics.IfTrafficScore),
			Labels:    device.Labels,
			Value:     lo.Min([]float64{rxRateScore, txRateScore}),
			Timestamp: timestamp,
		})
		results = append(results, &vtm.Metric{
			Metric:    string(metrics.IfTrafficAnomaly),
			Labels:    device.Labels,
			Value:     float64(rxRateAnomaly) + float64(txRateAnomaly),
			Timestamp: timestamp,
		})
		results = append(results, &vtm.Metric{
			Metric:    string(metrics.IfOperStatusScore),
			Labels:    device.Labels,
			Value:     operationalStatusScore,
			Timestamp: timestamp,
		})
		results = append(results, &vtm.Metric{
			Metric:    string(metrics.IfOperStatusAnomaly),
			Labels:    device.Labels,
			Value:     float64(operationalStatusAnomaly),
			Timestamp: timestamp,
		})
	}
	return results
}

func calcAp360(apMetrics map[string]map[string]*ApSchema, timestamp int64) []*vtm.Metric {
	results := make([]*vtm.Metric, 0)
	for _, apsBySite := range apMetrics {
		for _, ap := range apsBySite {
			var score float64
			if ap.ApStatus == 0 && ap.ChannelUtilization == nil && ap.ChannelAssociationClients == nil {
				score = -1
			} else {
				apStatusScore := calcApStatusScore(ap.ApStatus)
				apRadioScore := calcApRadioScore(ap.ChannelUtilization, ap.ChannelAssociationClients)
				if apStatusScore == -1 || apRadioScore == -1 {
					score = -1
				} else {
					score = lo.Min([]float64{apStatusScore, apRadioScore})
				}
			}

			results = append(results, &vtm.Metric{
				Metric:    string(metrics.HealthScore),
				Labels:    ap.Labels,
				Value:     score,
				Timestamp: timestamp,
			})
		}
	}
	return results
}

func RunDevice360OfflineTask() {
	timestamp := time.Now().UTC().Unix()
	queries := getQueries()
	deviceVectors, apVectors, err := queryResults(queries)
	if err != nil {
		core.Logger.Error("[device360OfflineTask]: failed to get vector from victoriaMetrics", zap.Error(err))
		return
	}

	aggDeviceMetrics := aggregateDeviceMetrics(deviceVectors)
	aggApMetrics := aggregateApMetrics(apVectors)

	deviceHealthScores := calcHealthScore(aggDeviceMetrics, timestamp)
	apHealthScores := calcAp360(aggApMetrics, timestamp)

	vtmClient := vtm.NewVtmClient()
	err = vtmClient.BulkImportMetrics(deviceHealthScores, nil)
	if err != nil {
		core.Logger.Error("[device360OfflineTask]failed to import metrics to vtm", zap.Error(err))
	}
	err = vtmClient.BulkImportMetrics(apHealthScores, nil)
	if err != nil {
		core.Logger.Error("[device360OfflineTask]failed to import metrics to vtm", zap.Error(err))
	}
}
