package schemas

import (
	"time"

	"github.com/wangxin688/narvis/server/tools/schemas"
)

type GroupCreate struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
	RoleID      string  `json:"role_id" binding:"required,uuid"`
}

type GroupUpdate struct {
	Name        *string `json:"name" binding:"omitempty"`
	Description *string `json:"description" binding:"omitempty"`
	RoleID      *string `json:"role_id" binding:"omitempty,uuid"`
}

type GroupQuery struct {
	schemas.PageInfo
	ID     *[]string `form:"id" binding:"omitempty,list_uuid"`
	Name   *[]string `form:"name" binding:"omitempty"`
	RoleID *[]string `form:"role_id" binding:"omitempty,list_uuid"`
}

type GroupShort struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type Group struct {
	GroupShort
	Description *string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Role        RoleShort `json:"role"`
}

type GroupList []*Group
