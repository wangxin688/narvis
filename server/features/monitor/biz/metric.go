package metric_biz

import (
	"strings"
	"time"

	"github.com/wangxin688/narvis/intend/logger"
	"github.com/wangxin688/narvis/intend/metrics"
	"github.com/wangxin688/narvis/server/pkg/contextvar"
	"github.com/wangxin688/narvis/server/pkg/vtm"
	"github.com/wangxin688/narvis/server/tools/errors"
	"go.uber.org/zap"
)

type MetricService struct {
	StartedAt  time.Time `form:"startedAt" binding:"required,datetime"`
	EndedAt    time.Time `form:"endedAt" binding:"required,datetime"`
	DataPoints uint      `form:"dataPoints"`
	SiteId     *[]string `form:"siteId" binding:"omitempty,list_uuid"`
	DeviceId   *[]string `form:"deviceId" binding:"omitempty,list_uuid"`
	ApName     *[]string `form:"apName" binding:"omitempty"`
	IfName     *[]string `form:"ifName" binding:"omitempty"`
	RadioType  *[]string `form:"radioType" binding:"omitempty;oneof=2.4GHz 5GHz 6GHz"`
	FuncName   *string   `form:"funcName" binding:"omitempty"`
	MetricName []string  `form:"metricName" binding:"required"`
}

func (m *MetricService) QueryMatrix() ([]*vtm.MatrixResponse, error) {

	results, err := m.queryMatrix()
	if err != nil {
		return nil, err
	}
	vtm.AdjustMatrixResponse(results, m.StartedAt.Unix())

	return results, nil
}

func (m *MetricService) buildBasicQuery() ([]*vtm.MatrixRequest, error) {

	results := make([]*vtm.MatrixRequest, len(m.MetricName))

	for i, name := range m.MetricName {
		query := vtm.NewPromQLBuilder(name)
		meta := metrics.GetMetric(name)
		if meta == nil {
			return nil, errors.NewError(errors.CodeMetricNotDefined, errors.MsgMetricNotDefined, name)
		}
		if m.FuncName != nil {
			query = query.WithFuncName(*m.FuncName)
		}

		if m.SiteId != nil {
			query = query.WithLabels(vtm.Label{
				Name:    "siteId",
				Value:   listToReString(*m.SiteId),
				Matcher: vtm.LikeMatcher,
			})
		}
		if m.DeviceId != nil {
			query = query.WithLabels(vtm.Label{
				Name:    "deviceId",
				Value:   listToReString(*m.DeviceId),
				Matcher: vtm.LikeMatcher,
			})
		}
		if m.ApName != nil {
			query = query.WithLabels(vtm.Label{
				Name:    "apName",
				Value:   listToReString(*m.ApName),
				Matcher: vtm.LikeMatcher,
			})
		}

		if m.IfName != nil {
			query = query.WithLabels(vtm.Label{
				Name:    "ifName",
				Value:   listToReString(*m.IfName),
				Matcher: vtm.LikeMatcher,
			})
		}

		if m.RadioType != nil {
			query = query.WithLabels(vtm.Label{
				Name:    "radioType",
				Value:   listToReString(*m.RadioType),
				Matcher: vtm.LikeMatcher,
			})
		}
		query = query.WithWindow("5m")
		queryString, err := query.Build()
		if err != nil {
			logger.Logger.Error("[metricService]: failed to build query", zap.String("metric", name), zap.Error(err))
			return nil, errors.NewError(errors.CodeQueryBuildFailed, errors.MsgQueryBuildFailed, name)
		}
		if m.DataPoints == 0 {
			m.DataPoints = 200
		}

		results[i] = &vtm.MatrixRequest{
			Query:        queryString,
			Step:         vtm.CalculateInterval(m.StartedAt.Unix(), m.EndedAt.Unix(), m.DataPoints),
			LegendFormat: &meta.Legend,
			Start:        m.StartedAt.Unix(),
			End:          m.EndedAt.Unix(),
		}
	}
	return results, nil
}

func (m *MetricService) queryMatrix() ([]*vtm.MatrixResponse, error) {
	if len(m.MetricName) == 0 {
		return nil, nil
	}
	queries, err := m.buildBasicQuery()
	if err != nil {
		return nil, err
	}
	orgId := contextvar.OrganizationId.Get()
	return vtm.NewVtmClient().GetBulkMatrix(queries, &orgId)
}

func listToReString(values []string) string {
	if len(values) == 0 {
		return ""
	}
	return strings.Join(values, "|")
}
