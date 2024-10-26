package schemas

type OrmDiff struct {
	Before any `json:"before"`
	After  any `json:"after"`
}
