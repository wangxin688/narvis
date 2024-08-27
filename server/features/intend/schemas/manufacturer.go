package schemas

type ManufacturerQuery struct {
	Manufacturer     *string   `form:"name" binding:"omitempty"`
	Keyword  *string   `form:"keyword" binding:"omitempty"`
}


