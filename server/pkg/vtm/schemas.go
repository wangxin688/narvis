package vtm

import (
	"encoding/json"
)

type QueryLabelRequest struct {
	Start     *uint64 `json:"start"`
	End       *uint64 `json:"end"`
	Match     string  `json:"match"` // query string
	LabelName string  `json:"label_name"`
}

type BaseRequest struct {
	Step         uint64  `json:"step"`
	Query        string  `json:"query"`
	LegendFormat *string `json:"legend_format"`
}

type VectorRequest struct {
	BaseRequest
	Time *uint `json:"time"`
}

type MatrixRequest struct {
	BaseRequest
	Start uint64 `json:"start"`
	End   uint64 `json:"end"`
}

type VectorResponse struct {
	Metric map[string]string `json:"metric"`
	Value  [2]string         `json:"value"`
	Legend *string           `json:"legend"`
}

type MatrixResponse struct {
	Metric map[string]string `json:"metric"`
	Values [][][2]string     `json:"value"`
	Legend *string           `json:"legend"`
}

type VtmResponse struct {
	Status string  `json:"status"`
	Data   VtmData `json:"data"`
}

type VtmData struct {
	ResultType string          `json:"resultType"`
	Result     json.RawMessage `json:"result"`
}

type Metric struct {
	Metric    string            `json:"metric"`
	Labels    map[string]string `json:"labels"`
	Value     float64           `json:"value"`
	Timestamp uint64            `json:"timestamp"`
}
