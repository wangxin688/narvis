package vtm

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"

	"github.com/imroc/req/v3"
	"github.com/wangxin688/narvis/server/core"
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
	url := core.Settings.Vtm.Url
	username := core.Settings.Vtm.Username
	password := core.Settings.Vtm.Password
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
	json.Unmarshal(rsp.Data.Result, &v)
}

func (c *VtmClient) GetLabelValues(label QueryLabelRequest, OrganizationId *string) (result []string, err error) {
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
	if OrganizationId != nil {
		c.R().SetQueryParam("extra_label=organizationId", *OrganizationId)
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

func (c *VtmClient) GetVector(query VectorRequest, OrganizationId *string) (results []*VectorResponse, err error) {

	vector_query_path := "/api/v1/query"
	if query.Time != nil {
		c.R().SetQueryParam("time", fmt.Sprintf("%d", *query.Time))
	}
	if query.LegendFormat != nil {
		c.R().SetQueryParam("legend_format", *query.LegendFormat)
	}
	if OrganizationId != nil {
		c.R().SetQueryParam("extra_label=organizationId", *OrganizationId)
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

func (c *VtmClient) GetMatrix(query MatrixRequest, OrganizationId *string) (results []*MatrixResponse, err error) {

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
	if OrganizationId != nil {
		c.R().SetQueryParam("extra_label=organizationId", *OrganizationId)
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

func (v *VtmClient) GetBulkVector(queries []VectorRequest, organizationID *string) ([]*VectorResponse, error) {
	results := make([]*VectorResponse, 0, len(queries))
	resultChan := make(chan []*VectorResponse, len(queries))
	var wg sync.WaitGroup

	for _, query := range queries {
		wg.Add(1)
		go func(q VectorRequest) {
			defer wg.Done()
			vectors, err := v.GetVector(q, organizationID)
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

func (v *VtmClient) GetBulkMatrix(requests []MatrixRequest, organizationID *string) ([]*MatrixResponse, error) {
	results := make([]*MatrixResponse, 0)
	resultChan := make(chan []*MatrixResponse, len(requests))
	var wg sync.WaitGroup

	for _, request := range requests {
		wg.Add(1)
		go func(req MatrixRequest) {
			defer wg.Done()
			matrix, err := v.GetMatrix(req, organizationID)
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
