package vtm

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/imroc/req/v3"
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

func NewVtmClient(url, username, password string) *VtmClient {
	once.Do(func() {
		vtmInstance = &VtmClient{
			Client: req.C().SetBaseURL(url).SetCommonBasicAuth(username, password).
				OnAfterResponse(func(client *req.Client, resp *req.Response) error {
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
	json.Unmarshal(rsp.Data.Result, &v)
}

func (c *VtmClient) GetLabelValues(label QueryLabelRequest, OrganizationID string) (result []string, err error) {
	label_query_path := fmt.Sprintf("/api/v1/label/%s/values", label.LabelName)
	if label.Match != "" {
		c.R().SetQueryParam("match", label.Match)
	}
	if label.Start != nil {
		c.R().SetQueryParam("start", fmt.Sprintf("%d", *label.Start))
	}
	if label.End != nil {
		c.R().SetQueryParam("end", fmt.Sprintf("%d", *label.End))
	}
	if OrganizationID != "" {
		c.R().SetQueryParam("extra_label=organization_id", OrganizationID)
	}
	rsp := VtmResponse{}
	c.R().SetSuccessResult(&rsp)
	_, err = c.R().Get(label_query_path)
	if err != nil {
		return nil, err
	}
	rsp.getResult(&result)
	return
}

func (c *VtmClient) GetVector(query VectorRequest, OrganizationID string) (results []*VectorResponse, err error) {

	vector_query_path := "/api/v1/query"
	if query.Time != nil {
		c.R().SetQueryParam("time", fmt.Sprintf("%d", *query.Time))
	}
	if query.LegendFormat != nil {
		c.R().SetQueryParam("legend_format", *query.LegendFormat)
	}
	if OrganizationID != "" {
		c.R().SetQueryParam("extra_label=organization_id", OrganizationID)
	}
	rsp := VtmResponse{}
	c.R().SetSuccessResult(&rsp)
	_, err = c.R().Get(vector_query_path)
	if err != nil {
		return nil, err
	}
	rsp.getResult(&results)
	return
}

func (c *VtmClient) GetMatrix(query MatrixRequest, OrganizationID string) (results []*MatrixResponse, err error) {

	matrix_query_path := "/api/v1/query_range"

	if query.Start != 0 {
		c.R().SetQueryParam("start", fmt.Sprintf("%d", query.Start))
	}
	if query.End != 0 {
		c.R().SetQueryParam("end", fmt.Sprintf("%d", query.End))
	}
	if query.Step != 0 {
		c.R().SetQueryParam("step", fmt.Sprintf("%d", query.Step))
	}
	if query.LegendFormat != nil {
		c.R().SetQueryParam("legend_format", *query.LegendFormat)
	}
	if OrganizationID != "" {
		c.R().SetQueryParam("extra_label=organization_id", OrganizationID)
	}

	rsp := VtmResponse{}
	c.R().SetSuccessResult(&rsp)
	_, err = c.R().Get(matrix_query_path)
	if err != nil {
		return nil, err
	}
	rsp.getResult(&results)
	return
}

func (v *VtmClient) GetBulkVector(query []VectorRequest, OrganizationID string) (results []*VectorResponse, err error) {

	var wg sync.WaitGroup

	for _, q := range query {
		wg.Add(1)
		go func(q VectorRequest) {
			defer wg.Done()
			matrix, err := v.GetVector(q, OrganizationID)
			if err != nil {
				return
			}
			results = append(results, matrix...)
		}(q)
	}

	wg.Wait()
	return
}

func (v *VtmClient) GetBulkMatrix(query []MatrixRequest, OrganizationID string) (results []*MatrixResponse, err error) {

	var wg sync.WaitGroup

	for _, q := range query {
		wg.Add(1)
		go func(q MatrixRequest) {
			defer wg.Done()
			matrix, err := v.GetMatrix(q, OrganizationID)
			if err != nil {
				return
			}
			results = append(results, matrix...)
		}(q)
	}

	wg.Wait()
	return
}

func (c *VtmClient) BulkImportMetrics(metrics []Metric, OrganizationID string) {
	importPath := "/api/v1/import/prometheus"
	if OrganizationID != "" {
		c.R().SetQueryParam("extra_label=organization_id", OrganizationID)
	}

	metricJson := metricStringBuilder(metrics)
	c.R().SetBody(metricJson).Post(importPath)
}

func metricStringBuilder(metrics []Metric) string {

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
