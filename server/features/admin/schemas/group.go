package schemas

import (
	"time"

	"github.com/wangxin688/narvis/server/tools/schemas"
)

type GroupCreate struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description"`
	RoleId      string  `json:"roleId" binding:"required,uuid"`
}

type GroupUpdate struct {
	Name        *string `json:"name" binding:"omitempty"`
	Description *string `json:"description" binding:"omitempty"`
	RoleId      *string `json:"roleId" binding:"omitempty,uuid"`
}

type GroupQuery struct {
	schemas.PageInfo
	Id     *[]string `form:"id" binding:"omitempty,list_uuid"`
	Name   *[]string `form:"name" binding:"omitempty"`
	RoleId *[]string `form:"roleId" binding:"omitempty,list_uuid"`
}

type GroupShort struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Group struct {
	GroupShort
	Description *string   `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Role        RoleShort `json:"role"`
}

type GroupList []*Group
