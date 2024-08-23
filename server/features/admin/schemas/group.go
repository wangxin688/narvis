package schemas

import "github.com/wangxin688/narvis/server/schemas"

type GroupCreate struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
	RoleId      string  `json:"role_id" binding:"required,uuid"`
}

type GroupUpdate struct {
	Name        *string `json:"name" binding:"omitempty"`
	Description *string `json:"description" binding:"omitempty"`
	RoleId      string  `json:"role_id" binding:"omitempty,uuid"`
}

type GroupQuery struct {
	schemas.PageInfo
	Id   *[]string `json:"id" binding:"omitempty,list_uuid"`
	Name *[]string `json:"name" binding:"omitempty"`
}

type GroupShort struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Group struct {
	GroupShort
	Role Role `json:"role"`
}
