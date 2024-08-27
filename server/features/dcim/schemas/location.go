package schemas

type LocationCreate struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description" binding:"omitempty"`
	ParentID    *string `json:"parent_id" binding:"omitempty,uuid"`
	SiteID      *string `json:"site_id" binding:"omitempty,uuid"`
}

type LocationShort struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
