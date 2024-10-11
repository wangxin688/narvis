package schemas

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
