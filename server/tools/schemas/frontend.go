package schemas

type PieChartData struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

type Bar struct {
	Title string `json:"title"`
	Data  []int  `json:"data"`
}

type BarChartData struct {
	ChartData []Bar
	XAxis     []string `json:"xAxis"`
}
