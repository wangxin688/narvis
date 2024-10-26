package infra_biz

import (
	"strings"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/intend/metrics"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/features/monitor/schemas"
	"github.com/wangxin688/narvis/server/pkg/vtm"
	"go.uber.org/zap"
)

func OpStatusMapping(status string) string {
	switch status {
	case "1":
		return "up"
	case "0":
		return "down"
	default:
		return "nodata"
	}
}
func OpStatusReverseMapping(status string) int {
	switch status {
	case "up":
		return 1
	case "down":
		return 0
	default:
		return -1
	}
}

func GetApOpStatus(apIds []string, orgId string) (map[string]string, error) {
	apMap, err := NewApService().GetApShortMap(apIds)
	if err != nil {
		return nil, err
	}
	apNames := []string{}
	apNameIdMap := make(map[string]string)
	res := make(map[string]string)
	for _, ap := range apMap {
		res[ap.Id] = "nodata"
		apNames = append(apNames, ap.Name)
		apNameIdMap[ap.Name] = ap.Id
	}
	ql := vtm.NewPromQLBuilder(string(metrics.ApStatus))
	query, err := ql.WithFuncName("last_over_time").WithLabels(
		vtm.Label{
			Name:    "apName",
			Value:   strings.Join(apNames, "|"),
			Matcher: vtm.LikeMatcher,
		},
	).WithWindow("5m").Build()

	if err != nil {
		core.Logger.Error("[metricService]: failed to build ap operation status query", zap.Error(err))
		return res, err
	}
	vectors, err := vtm.NewVtmClient().GetVector(&vtm.VectorRequest{Query: query, Step: 60}, &orgId)
	if err != nil {
		core.Logger.Error("[metricService]: failed to get ap operation status", zap.Error(err))
		return res, err
	}
	if len(vectors) == 0 {
		return res, nil
	}
	for _, v := range vectors {
		res[apNameIdMap[v.Metric["apName"]]] = OpStatusMapping(v.Value[1].(string))
	}
	return res, nil

}

func GetDeviceOpStatus(deviceIds []string, orgId string) (map[string]string, error) {
	res := make(map[string]string)
	for _, device := range deviceIds {
		res[device] = "nodata"
	}
	ql := vtm.NewPromQLBuilder(string(metrics.ICMPPing))
	query, err := ql.WithFuncName("last_over_time").WithLabels(
		vtm.Label{
			Name:    "deviceId",
			Value:   strings.Join(deviceIds, "|"),
			Matcher: vtm.LikeMatcher,
		},
	).WithWindow("5m").Build()

	if err != nil {
		core.Logger.Error("[metricService]: failed to build device operation status query", zap.Error(err))
		return res, err
	}
	vectors, err := vtm.NewVtmClient().GetVector(&vtm.VectorRequest{Query: query, Step: 60}, &orgId)
	if err != nil {
		core.Logger.Error("[metricService]: failed to get device operation status", zap.Error(err))
		return res, err
	}
	if len(vectors) == 0 {
		return res, nil
	}
	for _, v := range vectors {
		res[v.Metric["deviceId"]] = OpStatusMapping(v.Value[1].(string))
	}
	return res, nil
}

// GetApIdsByOpStatus get ap Ids by ap operational status
func GetApIdsByOpStatus(siteId string, opStatus string, orgId string) ([]string, error) {
	query, err := vtm.NewPromQLBuilder(string(metrics.ApStatus)).
		WithFuncName("last_over_time").
		WithLabels(vtm.Label{Name: "siteId", Value: siteId, Matcher: vtm.EqualMatcher}).
		WithWindow("5m").Build()

	if err != nil {
		core.Logger.Error("[metricService]: failed to build ap operation status query", zap.Error(err))
		return nil, err
	}
	vectors, err := vtm.NewVtmClient().GetVector(&vtm.VectorRequest{Query: query, Step: 60}, nil)
	if err != nil {
		core.Logger.Error("[metricService]: failed to get ap operation status", zap.Error(err))
		return nil, err
	}
	if len(vectors) == 0 {
		if opStatus == "nodata" {
			return NewApService().GetAllApIdsBySiteId(siteId)
		}
		return []string{}, nil
	}
	if opStatus == "nodata" {
		allApNameWithData := lo.Map(vectors, func(v *vtm.VectorResponse, _ int) string {
			return v.Metric["apName"]
		})
		allApIds, err := NewApService().GetAllApIdsBySiteId(siteId)
		if err != nil {
			core.Logger.Error("[metricService]: failed to get ap ids by site id", zap.String("siteId", siteId), zap.Error(err))
			return nil, err
		}
		apIds, err := NewApService().GetApIdsByNames(allApNameWithData, orgId)
		if err != nil {
			core.Logger.Error("[metricService]: failed to get ap ids by names", zap.Error(err))
			return nil, err
		}
		subApIds := lo.Intersect(allApIds, apIds)
		return subApIds, nil
	}
	var apNames []string
	if opStatus == "up" {
		apNames = lo.Map(vectors, func(v *vtm.VectorResponse, _ int) string {
			if v.Value[1] == "1" {
				return v.Metric["apName"]
			}
			return ""
		})
		apNames = lo.Filter(apNames, func(item string, _ int) bool {
			return item != ""
		})
	} else if opStatus == "down" {
		apNames = lo.Map(vectors, func(v *vtm.VectorResponse, _ int) string {
			if v.Value[1] == "0" {
				return v.Metric["apName"]
			}
			return ""
		})
		apNames = lo.Filter(apNames, func(item string, _ int) bool {
			return item != ""
		})
	}
	if len(apNames) == 0 {
		return apNames, nil
	}

	apIds, err := NewApService().GetApIdsByNames(apNames, orgId)
	if err != nil {
		core.Logger.Error("[metricService]: failed to get ap ids by names", zap.Error(err))
		return nil, err
	}
	return apIds, nil
}

// GetDeviceIdsByOpStatus get device Ids by device operational status
func GetDeviceIdsByOpStatus(siteId string, opStatus string, orgId string) ([]string, error) {
	query, err := vtm.NewPromQLBuilder(string(metrics.ICMPPing)).
		WithFuncName("last_over_time").
		WithLabels(vtm.Label{Name: "siteId", Value: siteId, Matcher: vtm.EqualMatcher}).
		WithWindow("5m").Build()

	if err != nil {
		core.Logger.Error("[metricService]: failed to build device operation status query", zap.Error(err))
		return nil, err
	}
	vectors, err := vtm.NewVtmClient().GetVector(&vtm.VectorRequest{Query: query, Step: 60}, nil)
	if err != nil {
		core.Logger.Error("[metricService]: failed to get device operation status", zap.Error(err))
		return nil, err
	}
	if len(vectors) == 0 {
		if opStatus == "nodata" {
			return NewDeviceService().GetAllDeviceIdsBySiteId(siteId)
		}
		return []string{}, nil
	}
	if opStatus == "nodata" {
		allDeviceIdsWithData := lo.Map(vectors, func(v *vtm.VectorResponse, _ int) string {
			return v.Metric["deviceId"]
		})
		allDeviceIds, err := NewDeviceService().GetAllDeviceIdsBySiteId(siteId)
		if err != nil {
			core.Logger.Error("[metricService]: failed to get device ids by site id", zap.String("siteId", siteId), zap.Error(err))
			return nil, err
		}
		subDeviceIds := lo.Intersect(allDeviceIds, allDeviceIdsWithData)
		return subDeviceIds, nil
	}
	var deviceIds []string
	if opStatus == "up" {
		deviceIds = lo.Map(vectors, func(v *vtm.VectorResponse, _ int) string {
			if v.Value[1] == "1" {
				return v.Metric["deviceId"]
			}
			return ""
		})
		deviceIds = lo.Filter(deviceIds, func(item string, _ int) bool {
			return item != ""
		})
	} else if opStatus == "down" {
		deviceIds = lo.Map(vectors, func(v *vtm.VectorResponse, _ int) string {
			if v.Value[1] == "0" {
				return v.Metric["deviceId"]
			}
			return ""
		})
		deviceIds = lo.Filter(deviceIds, func(item string, _ int) bool {
			return item != ""
		})
	}
	return deviceIds, nil
}

func GetApAssociatedClients(apNames []string, siteId string, orgId string) (map[string]*schemas.ApClientItem, error) {
	apNamesString := strings.Join(apNames, "|")
	query, err := vtm.NewPromQLBuilder(string(metrics.ChannelAssociationClients)).
		WithFuncName("last_over_time").WithLabels(
		vtm.Label{
			Name:    "apName",
			Value:   apNamesString,
			Matcher: vtm.LikeMatcher,
		},
	).WithLabels(vtm.Label{Name: "siteId", Value: siteId, Matcher: vtm.EqualMatcher}).Build()
	if err != nil {
		core.Logger.Error("[metricService]: failed to build ap associated clients query", zap.Error(err))
		return nil, err
	}
	vectors, err := vtm.NewVtmClient().GetVector(&vtm.VectorRequest{Query: query, Step: 60}, &orgId)
	if err != nil {
		core.Logger.Error("[metricService]: failed to get ap associated clients", zap.Error(err))
		return nil, err
	}
	result := make(map[string]*schemas.ApClientItem)
	if len(vectors) == 0 {
		return result, nil
	}

	for _, v := range vectors {
		result[v.Metric["apName"]] = &schemas.ApClientItem{
			Channel:      v.Metric["channel"],
			NumOfClients: v.Value[1].(string),
			RadioType:    v.Metric["radioType"],
		}
	}
	return result, nil
}
