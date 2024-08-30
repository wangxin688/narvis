package schemas

type LocationCreate struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description" binding:"omitempty"`
	ParentId    *string `json:"parent_id" binding:"omitempty,uuid"`
	SiteId      *string `json:"site_id" binding:"omitempty,uuid"`
}

type LocationShort struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
