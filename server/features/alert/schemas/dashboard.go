package schemas

type TopX struct {
	Value int    `json:"value"`
	Name  string `json:"name"`
}

type TrendItem struct {
	Value int
	Severity string
	Date string
}

type Trend struct {
	Date  string `json:"name"`
	Value int    `json:"value"`
}

type AlertTrend struct {
	Severity string   `json:"severity"`
	Trend    []*Trend `json:"trend"`
}
