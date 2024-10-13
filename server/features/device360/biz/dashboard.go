package device360_biz

import (
	"github.com/wangxin688/narvis/intend/devicerole"
	"github.com/wangxin688/narvis/intend/metrics"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/features/device360/schemas"
	device360_utils "github.com/wangxin688/narvis/server/features/device360/utils"
	"github.com/wangxin688/narvis/server/pkg/vtm"
	"github.com/wangxin688/narvis/server/tools/errors"
	"go.uber.org/zap"
)

func GetAllDeviceHealth(siteId *string, orgId string) (*schemas.HealthHeatMap, error) {

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
		core.Logger.Error("[metricService]: failed to build all device health query", zap.Error(err))
		return nil, errors.NewError(errors.CodeQueryBuildFailed, errors.MsgQueryBuildFailed, "all device health")
	}
	vectors, err := vtm.NewVtmClient().GetVector(&vtm.VectorRequest{Query: queryString, Step: 60}, &orgId)
	if err != nil {
		core.Logger.Error("[metricService]: failed to get all device health", zap.Error(err))
		return nil, err
	}
	return device360_utils.VectorResponseToHealthMap(vectors), nil
}

func GetSwitchHealth(siteId *string, orgId string) (*schemas.HealthHeatMap, error) {
	query := vtm.NewPromQLBuilder(string(metrics.HealthScore)).
		WithLabels(vtm.Label{
			Name:    "deviceRole",
			Value:   getSwitchDeviceRoles(),
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
		core.Logger.Error("[metricService]: failed to build switch health query", zap.Error(err))
		return nil, errors.NewError(errors.CodeQueryBuildFailed, errors.MsgQueryBuildFailed, "switch health")
	}
	vectors, err := vtm.NewVtmClient().GetVector(&vtm.VectorRequest{Query: queryString, Step: 60}, &orgId)
	if err != nil {
		core.Logger.Error("[metricService]: failed to get switch health", zap.Error(err))
		return nil, err
	}
	return device360_utils.VectorResponseToHealthMap(vectors), nil
}

func GetAccessPointsHealth(siteId *string, orgId string) (*schemas.HealthHeatMap, error) {
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
		core.Logger.Error("[metricService]: failed to build access point health query", zap.Error(err))
		return nil, errors.NewError(errors.CodeQueryBuildFailed, errors.MsgQueryBuildFailed, "access point health")
	}
	vectors, err := vtm.NewVtmClient().GetVector(&vtm.VectorRequest{Query: queryString, Step: 60}, &orgId)
	if err != nil {
		core.Logger.Error("[metricService]: failed to get access point health", zap.Error(err))
		return nil, err
	}
	return device360_utils.VectorResponseToHealthMap(vectors), nil
}

func GetGatewayHealth(siteId *string, orgId string) (*schemas.HealthHeatMap, error) {
	query := vtm.NewPromQLBuilder(string(metrics.HealthScore)).
		WithLabels(vtm.Label{
			Name:    "deviceRole",
			Value:   getGatewayDeviceRoles(),
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
		core.Logger.Error("[metricService]: failed to build gateway health query", zap.Error(err))
		return nil, errors.NewError(errors.CodeQueryBuildFailed, errors.MsgQueryBuildFailed, "gateway health")
	}
	vectors, err := vtm.NewVtmClient().GetVector(&vtm.VectorRequest{Query: queryString, Step: 60}, &orgId)
	if err != nil {
		core.Logger.Error("[metricService]: failed to get gateway health", zap.Error(err))
		return nil, err
	}
	return device360_utils.VectorResponseToHealthMap(vectors), nil
}

// TODO: add circuit health in device360 task
// func getCircuitHealth(siteId *string, orgId string) (*schemas.HealthHeatMap, error) {
// 	return nil, nil
// }
