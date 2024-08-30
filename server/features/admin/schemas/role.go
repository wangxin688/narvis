package schemas

import (
	"time"

	"github.com/wangxin688/narvis/server/tools/schemas"
)

type RoleCreate struct {
	Name        string   `json:"name" binding:"required"`
	Description *string  `json:"description"`
	Menus       []string `json:"menus" binding:"required,list_uuid"`
}

type RoleUpdate struct {
	Name        *string   `json:"name" binding:"omitempty"`
	Description *string   `json:"description" binding:"omitempty"`
	Menus       *[]string `json:"menus" binding:"omitempty,list_uuid"`
}

type RoleQuery struct {
	schemas.PageInfo
	Id   *[]string `json:"id" binding:"omitempty,list_uuid"`
	Name *[]string `json:"name" binding:"omitempty"`
}

type RoleDetail struct {
	Id          string    `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	Menus       *[]Menu   `json:"menus"`
}

type Role struct {
	Id          string    `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
}

type RoleList []Role

type RoleShort struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type RoleShorts []RoleShort
