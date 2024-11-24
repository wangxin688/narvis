package vtm

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/imroc/req/v3"
	"github.com/wangxin688/narvis/server/config"
)

var once sync.Once
var vtmInstance *VtmClient

type VtmClient struct {
	*req.Client
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func (e *ErrorMessage) Error() string {
	return fmt.Sprintf("request metrics system error: %s", e.Message)
}

func NewVtmClient() *VtmClient {
	url := config.Settings.Vtm.Url
	username := config.Settings.Vtm.Username
	password := config.Settings.Vtm.Password
	once.Do(func() {
		vtmInstance = &VtmClient{
			Client: req.C().SetBaseURL(url).SetCommonBasicAuth(username, password).
				OnAfterResponse(func(_ *req.Client, resp *req.Response) error {
					if resp.Err != nil {
						return resp.Err
					}
					if errMsg, ok := resp.ErrorResult().(*ErrorMessage); ok {
						resp.Err = errMsg
						return nil
					}
					if !resp.IsSuccessState() {
						resp.Err = fmt.Errorf("request metrics system error: %s, status code: %s", resp.Dump(), resp.Status)
					}
					return nil
				}),
		}
	})
	return vtmInstance
}

func (rsp *VtmResponse) getResult(v any) {
	json.Unmarshal(rsp.Data.Result, &v) //nolint: errcheck
}

func (c *VtmClient) GetLabelValues(label *QueryLabelRequest, OrganizationId *string) (result []string, err error) {
	label_query_path := fmt.Sprintf("/api/v1/label/%s/values", label.LabelName)
	request := c.R()
	if label.Match != "" {
		request = request.SetQueryParam("match", label.Match)
	}
	if label.Start != nil {
		request = request.SetQueryParam("start", fmt.Sprintf("%d", *label.Start))
	}
	if label.End != nil {
		request = request.SetQueryParam("end", fmt.Sprintf("%d", *label.End))
	}
	if OrganizationId != nil {
		request = request.SetQueryParam("extra_label=organizationId", *OrganizationId)
	}
	rsp := VtmResponse{}
	request = request.SetSuccessResult(&rsp)
	_, err = request.Get(label_query_path)
	if err != nil {
		return nil, err
	}
	rsp.getResult(&result)
	return
}

func (c *VtmClient) GetVector(query *VectorRequest, OrganizationId *string) (results []*VectorResponse, err error) {

	vector_query_path := "/api/v1/query"
	request := c.R()
	request = request.SetQueryParam("query", query.Query)
	if query.Time != nil {
		request = request.SetQueryParam("time", fmt.Sprintf("%d", *query.Time))
	}
	if query.LegendFormat != nil {
		request = request.SetQueryParam("legend_format", *query.LegendFormat)
	}
	if OrganizationId != nil {
		request = request.SetQueryParam("extra_label=organizationId", *OrganizationId)
	}
	rsp := VtmResponse{}
	request = request.SetSuccessResult(&rsp)
	_, err = request.Get(vector_query_path)
	if err != nil {
		return nil, err
	}
	rsp.getResult(&results)
	return
}

func (c *VtmClient) GetMatrix(query *MatrixRequest, OrganizationId *string) (results []*MatrixResponse, err error) {

	matrix_query_path := "/api/v1/query_range"
	request := c.R()
	request = request.SetQueryParam("query", query.Query)
	if query.Start != 0 {
		request = request.SetQueryParam("start", fmt.Sprintf("%d", query.Start))
	}
	if query.End != 0 {
		request = request.SetQueryParam("end", fmt.Sprintf("%d", query.End))
	}
	if query.Step != 0 {
		request = request.SetQueryParam("step", fmt.Sprintf("%d", query.Step))
	}
	if query.LegendFormat != nil {
		request = request.SetQueryParam("legend_format", *query.LegendFormat)
	}
	if OrganizationId != nil {
		request = request.SetQueryParam("extra_label=organizationId", *OrganizationId)
	}

	rsp := VtmResponse{}
	request = request.SetSuccessResult(&rsp)
	_, err = request.Get(matrix_query_path)
	if err != nil {
		return nil, err
	}
	rsp.getResult(&results)
	return
}

func (c *VtmClient) GetBulkVector(queries []*VectorRequest, organizationID *string) ([]*VectorResponse, error) {
	results := make([]*VectorResponse, 0, len(queries))
	resultChan := make(chan []*VectorResponse, len(queries))
	var wg sync.WaitGroup

	for _, query := range queries {
		wg.Add(1)
		go func(q *VectorRequest) {
			defer wg.Done()
			vectors, err := c.GetVector(q, organizationID)
			if err != nil {
				resultChan <- nil
				return
			}
			resultChan <- vectors
		}(query)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		if result == nil {
			return nil, fmt.Errorf("error occurred while getting vector response")
		}
		results = append(results, result...)
	}

	return results, nil
}

func (c *VtmClient) GetBulkMatrix(requests []*MatrixRequest, organizationID *string) ([]*MatrixResponse, error) {
	results := make([]*MatrixResponse, 0)
	resultChan := make(chan []*MatrixResponse, len(requests))
	var wg sync.WaitGroup

	for _, request := range requests {
		wg.Add(1)
		go func(req *MatrixRequest) {
			defer wg.Done()
			matrix, err := c.GetMatrix(req, organizationID)
			if err != nil {
				resultChan <- nil
				return
			}
			resultChan <- matrix
		}(request)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	for result := range resultChan {
		if result == nil {
			return results, fmt.Errorf("error occurred while getting matrix response")
		}
		results = append(results, result...)
	}

	return results, nil
}

func (c *VtmClient) BulkImportMetrics(metrics []*Metric, OrganizationId *string) (err error) {
	importPath := "/api/v1/import/prometheus"
	if OrganizationId != nil {
		c.R().SetQueryParam("extra_label=organizationId", *OrganizationId)
	}

	metricJson := metricStringBuilder(metrics)
	_, err = c.R().SetBody(metricJson).Post(importPath)
	if err != nil {
		return err
	}
	return nil
}

func metricStringBuilder(metrics []*Metric) string {

	var builder strings.Builder
	for _, metric := range metrics {
		builder.WriteString(metric.Metric)
		builder.WriteString("{")
		for key, value := range metric.Labels {
			builder.WriteString(fmt.Sprintf(`%s="%s",`, key, value))
		}
		builder.WriteString(fmt.Sprintf("} %v %v\n", metric.Value, metric.Timestamp))
	}

	return builder.String()
}

func CalculateInterval(start, end int64, dataPoints uint) int {
	intervals := []int{
		60,
		3 * 60,
		5 * 60,
		10 * 60,
		15 * 60,
		20 * 60,
		30 * 60,
		60 * 60,
		2 * 60 * 60,
		3 * 60 * 60,
		4 * 60 * 60,
		5 * 60 * 60,
		6 * 60 * 60,
		7 * 60 * 60,
		8 * 60 * 60,
		9 * 60 * 60,
		10 * 60 * 60,
		11 * 60 * 60,
		12 * 60 * 60,
		13 * 60 * 60,
		14 * 60 * 60,
		15 * 60 * 60,
		16 * 60 * 60,
		17 * 60 * 60,
		18 * 60 * 60,
		19 * 60 * 60,
		20 * 60 * 60,
		21 * 60 * 60,
		22 * 60 * 60,
		23 * 60 * 60,
		24 * 60 * 60,
	}
	timeDelta := end - start
	if timeDelta < 0 {
		return 0
	}
	for _, interval := range intervals {
		if timeDelta/int64(interval) < int64(dataPoints) {
			return interval
		}
	}
	return 60
}

// AdjustMatrixResponse adjusts matrix response based on startedAt,
// if startedAt is greater than the first value of the matrix response, it will be replaced by startedAt
func AdjustMatrixResponse(matrix []*MatrixResponse, startedAt int64) {

	for _, response := range matrix {
		vStartedAt := response.Values[0][0].(int64)
		if vStartedAt < startedAt {
			response.Values[0][0] = startedAt
		}
	}
}
