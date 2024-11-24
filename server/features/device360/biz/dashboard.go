package device360_biz

import (
	"time"

	"github.com/wangxin688/narvis/intend/logger"
	"github.com/wangxin688/narvis/intend/metrics"
	"github.com/wangxin688/narvis/intend/model/devicerole"
	"github.com/wangxin688/narvis/server/features/device360/schemas"
	device360_utils "github.com/wangxin688/narvis/server/features/device360/utils"
	"github.com/wangxin688/narvis/server/pkg/contextvar"
	"github.com/wangxin688/narvis/server/pkg/vtm"
	"github.com/wangxin688/narvis/server/tools/errors"
	"go.uber.org/zap"
)

func getAllDeviceHealth(siteId *string, orgId string) (*schemas.HealthHeatMap, error) {

	query := vtm.NewPromQLBuilder(string(metrics.HealthScore)).WithWindow("5m")
	if siteId != nil && *siteId != "" {
		query = query.WithLabels(vtm.Label{
			Name:    "siteId",
			Value:   *siteId,
			Matcher: vtm.EqualMatcher,
		})
	}
	queryString, err := query.Build()
	if err != nil {
		logger.Logger.Error("[metricService]: failed to build all device health query", zap.Error(err))
		return nil, errors.NewError(errors.CodeQueryBuildFailed, errors.MsgQueryBuildFailed, "all device health")
	}
	vectors, err := vtm.NewVtmClient().GetVector(&vtm.VectorRequest{Query: queryString, Step: 60}, &orgId)
	if err != nil {
		logger.Logger.Error("[metricService]: failed to get all device health", zap.Error(err))
		return nil, err
	}
	return device360_utils.VectorResponseToHealthMap(vectors), nil
}

func getSwitchHealth(siteId *string, orgId string) (*schemas.HealthHeatMap, error) {
	query := vtm.NewPromQLBuilder(string(metrics.HealthScore)).
		WithLabels(vtm.Label{
			Name:    "deviceRole",
			Value:   string(devicerole.Switch),
			Matcher: vtm.LikeMatcher,
		}).WithWindow("5m")
	if siteId != nil && *siteId != "" {
		query = query.WithLabels(vtm.Label{
			Name:    "siteId",
			Value:   *siteId,
			Matcher: vtm.EqualMatcher,
		})
	}
	queryString, err := query.Build()
	if err != nil {
		logger.Logger.Error("[metricService]: failed to build switch health query", zap.Error(err))
		return nil, errors.NewError(errors.CodeQueryBuildFailed, errors.MsgQueryBuildFailed, "switch health")
	}
	vectors, err := vtm.NewVtmClient().GetVector(&vtm.VectorRequest{Query: queryString, Step: 60}, &orgId)
	if err != nil {
		logger.Logger.Error("[metricService]: failed to get switch health", zap.Error(err))
		return nil, err
	}
	return device360_utils.VectorResponseToHealthMap(vectors), nil
}

func getAccessPointsHealth(siteId *string, orgId string) (*schemas.HealthHeatMap, error) {
	query := vtm.NewPromQLBuilder(string(metrics.HealthScore)).
		WithLabels(vtm.Label{
			Name:    "deviceRole",
			Value:   string(devicerole.WlanAP),
			Matcher: vtm.EqualMatcher,
		}).WithWindow("5m")
	if siteId != nil && *siteId != "" {
		query = query.WithLabels(vtm.Label{
			Name:    "siteId",
			Value:   *siteId,
			Matcher: vtm.EqualMatcher,
		})
	}
	queryString, err := query.Build()
	if err != nil {
		logger.Logger.Error("[metricService]: failed to build access point health query", zap.Error(err))
		return nil, errors.NewError(errors.CodeQueryBuildFailed, errors.MsgQueryBuildFailed, "access point health")
	}
	vectors, err := vtm.NewVtmClient().GetVector(&vtm.VectorRequest{Query: queryString, Step: 60}, &orgId)
	if err != nil {
		logger.Logger.Error("[metricService]: failed to get access point health", zap.Error(err))
		return nil, err
	}
	return device360_utils.VectorResponseToHealthMap(vectors), nil
}

func getFirewallHealth(siteId *string, orgId string) (*schemas.HealthHeatMap, error) {
	query := vtm.NewPromQLBuilder(string(metrics.HealthScore)).
		WithLabels(vtm.Label{
			Name:    "deviceRole",
			Value:   string(devicerole.FireWall),
			Matcher: vtm.LikeMatcher,
		}).WithWindow("5m")
	if siteId != nil && *siteId != "" {
		query = query.WithLabels(vtm.Label{
			Name:    "siteId",
			Value:   *siteId,
			Matcher: vtm.EqualMatcher,
		})
	}
	queryString, err := query.Build()
	if err != nil {
		logger.Logger.Error("[metricService]: failed to build gateway health query", zap.Error(err))
		return nil, errors.NewError(errors.CodeQueryBuildFailed, errors.MsgQueryBuildFailed, "gateway health")
	}
	vectors, err := vtm.NewVtmClient().GetVector(&vtm.VectorRequest{Query: queryString, Step: 60}, &orgId)
	if err != nil {
		logger.Logger.Error("[metricService]: failed to get gateway health", zap.Error(err))
		return nil, err
	}
	return device360_utils.VectorResponseToHealthMap(vectors), nil
}

func getSLAQuery(siteId *string, slaType string) (string, error) {
	matcher := vtm.EqualMatcher
	if slaType == "device" {
		matcher = vtm.NotEqualMatcher
	}

	query := vtm.NewPromQLBuilder(string(metrics.HealthScore)).WithFuncName(
		"avg_over_time").WithLabels(vtm.Label{
		Name:    "deviceRole",
		Value:   string(devicerole.WlanAP),
		Matcher: matcher,
	}).WithWindow("3m")
	if siteId != nil && *siteId != "" {
		query = query.WithLabels(vtm.Label{
			Name:    "siteId",
			Value:   *siteId,
			Matcher: vtm.EqualMatcher,
		}).WithAgg(vtm.Aggregation{
			Op:     vtm.Avg,
			AggWay: vtm.GroupBy,
			By:     []string{"__name__", "siteId"},
		})
	} else {
		query = query.WithAgg(vtm.Aggregation{
			Op:     vtm.Avg,
			AggWay: vtm.GroupBy,
			By:     []string{"__name__", "organizationId"},
		})
	}
	return query.Build()
}

// TODO: add circuit sla support
func GetSLA(siteId *string, orgId string, startedAtGte time.Time, startedAtLte time.Time) ([]*vtm.MatrixResponse, error) {
	wlanSLAQueryString, err := getSLAQuery(siteId, "wlan")
	if err != nil {
		logger.Logger.Error("[metricService]: failed to build wlan sla query", zap.Error(err))
		return nil, errors.NewError(errors.CodeQueryBuildFailed, errors.MsgQueryBuildFailed, "wlan sla")
	}
	deviceSLAQueryString, err := getSLAQuery(siteId, "device")
	if err != nil {
		logger.Logger.Error("[metricService]: failed to build device sla query", zap.Error(err))
		return nil, errors.NewError(errors.CodeQueryBuildFailed, errors.MsgQueryBuildFailed, "device sla")
	}
	wlanSLAQuery := vtm.MatrixRequest{
		Query: wlanSLAQueryString, Step: 60, Start: startedAtGte.Unix(), End: startedAtLte.Unix()}
	deviceSLAQuery := vtm.MatrixRequest{
		Query: deviceSLAQueryString, Step: 60, Start: startedAtGte.Unix(), End: startedAtLte.Unix()}
	vectors, err := vtm.NewVtmClient().GetBulkMatrix([]*vtm.MatrixRequest{&wlanSLAQuery, &deviceSLAQuery}, &orgId)
	if err != nil {
		logger.Logger.Error("[metricService]: failed to get wlan sla", zap.Error(err))
		return nil, err
	}
	return vectors, nil

}

func GetHealth(query *schemas.DeviceHealthQuery) (*schemas.HealthResponse, error) {
	orgId := contextvar.OrganizationId.Get()
	switchHealth, err := getSwitchHealth(query.SiteId, orgId)
	if err != nil {
		logger.Logger.Error("[metricService]: failed to get switch health", zap.Error(err))
		return nil, err
	}
	allDeviceHealth, err := getAllDeviceHealth(query.SiteId, orgId)
	if err != nil {
		logger.Logger.Error("[metricService]: failed to get all device health", zap.Error(err))
		return nil, err
	}
	firewallHealth, err := getFirewallHealth(query.SiteId, orgId)
	if err != nil {
		logger.Logger.Error("[metricService]: failed to get gateway health", zap.Error(err))
		return nil, err
	}
	accessPointHealth, err := getAccessPointsHealth(query.SiteId, orgId)
	if err != nil {
		logger.Logger.Error("[metricService]: failed to get access point health", zap.Error(err))
		return nil, err
	}
	return &schemas.HealthResponse{
		Firewall: *firewallHealth,
		WlanAP:   *accessPointHealth,
		Switch:   *switchHealth,
		Device:   *allDeviceHealth,
	}, nil
}
