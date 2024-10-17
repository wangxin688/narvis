package schemas

import ts "github.com/wangxin688/narvis/server/tools/schemas"

// TODO: Discuss
// Healthy Degraded Unhealthy Unavailable, Unknown for SRE
// Good, NeedsAttention, Critical, Unhealthy, Unknown
type HealthHeatMap struct {
	Good           int `json:"good"`
	NeedsAttention int `json:"needsAttention"`
	Poor           int `json:"poor"`
	Unhealthy      int `json:"unhealthy"`
	Unknown        int `json:"unknown"`
}

func (h *HealthHeatMap) Add(score string) {
	switch score {
	case "good":
		h.Good++
	case "needsAttention":
		h.NeedsAttention++
	case "poor":
		h.Poor++
	case "unhealthy":
		h.Unhealthy++
	case "unknown":
		h.Unknown++
	}
}

func (h *HealthHeatMap) ToPieChat() []*ts.PieChartData {
	return []*ts.PieChartData{
		{Name: "good", Value: h.Good},
		{Name: "needsAttention", Value: h.NeedsAttention},
		{Name: "poor", Value: h.Poor},
		{Name: "unhealthy", Value: h.Unhealthy},
		{Name: "unknown", Value: h.Unknown},
	}
}
