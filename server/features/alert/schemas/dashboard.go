package schemas

import "time"

type TopX struct {
	Value int    `json:"value"`
	Name  string `json:"name"`
}

type Trend struct {
	Name  time.Time `json:"name"`
	Value int       `json:"value"`
}
