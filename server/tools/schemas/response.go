package schemas

type ListResponse struct {
	Results interface{} `json:"results"`
	Total   int64       `json:"total"`
}

type IDResponse struct {
	ID string `json:"id"`
}
