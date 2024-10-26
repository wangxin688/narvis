package schemas

type Condition struct {
	Item  string   `json:"item" binding:"required"`
	Value []string `json:"value" binding:"required"` // ["*"] means match all
}
