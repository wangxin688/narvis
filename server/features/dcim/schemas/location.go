package schemas

type LocationCreate struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description" binding:"omitempty"`
	ParentId    *string `json:"parentId" binding:"omitempty,uuid"`
	SiteId      *string `json:"siteId" binding:"omitempty,uuid"`
}

type LocationShort struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
