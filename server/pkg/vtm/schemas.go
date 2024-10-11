package vtm

import (
	"encoding/json"
)

type QueryLabelRequest struct {
	Start     *int64 `json:"start,omitempty"`
	End       *int64 `json:"end,omitempty"`
	Match     string `json:"match"` // query string
	LabelName string `json:"label_name"`
}
type VectorRequest struct {
	Step         int     `json:"step"`
	Query        string  `json:"query"`
	LegendFormat *string `json:"legend_format,omitempty"`
	Time         *int64  `json:"time,omitempty"`
}

type MatrixRequest struct {
	Step         int     `json:"step"`
	Query        string  `json:"query"`
	LegendFormat *string `json:"legend_format,omitempty"`
	Start        int64   `json:"start"`
	End          int64   `json:"end"`
}

type VectorResponse struct {
	Metric map[string]string `json:"metric"`
	Value  [2]any            `json:"value"`
	Legend *string           `json:"legend"`
}

type MatrixResponse struct {
	Metric map[string]string `json:"metric"`
	Values [][2]any          `json:"value"`
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
	Timestamp int64             `json:"timestamp"`
}
