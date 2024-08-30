package schemas

type ListResponse struct {
	Results interface{} `json:"results"`
	Total   int64       `json:"total"`
}

type IdResponse struct {
	Id string `json:"id"`
}

type ErrorResponse struct {
	Code    uint64 `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}
