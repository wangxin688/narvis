package schemas

import (
	"time"

	"github.com/wangxin688/narvis/server/tools/schemas"
)

type UserCreate struct {
	Username string  `json:"username" binding:"required"`
	Email    string  `json:"email" binding:"required;email"`
	Password string  `json:"password" `
	Status   string  `json:"status" binding:"required,oneof=Active Inactive"`
	Avatar   *string `json:"avatar"`
	GroupId  string  `json:"groupId" binding:"required,uuid"`
	RoleId   string  `json:"roleId" binding:"required,uuid"`
	AuthType uint8   `json:"authType" binding:"required,gte=0,lte=4"`
}

type UserUpdate struct {
	Username *string `json:"username" binding:"omitempty"`
	Email    *string `json:"email" binding:"omitempty,email"`
	Password *string `json:"password" binding:"omitempty"`
	Status   *string `json:"status" binding:"omitempty,oneof=Active Inactive"`
	Avatar   *string `json:"avatar" binding:"omitempty"`
	GroupId  *string `json:"groupId" binding:"omitempty,uuid"`
	RoleId   *string `json:"roleId" binding:"omitempty,uuid"`
}

type UserQuery struct {
	schemas.PageInfo
	Id       *[]string `form:"id" binding:"omitempty,list_uuid"`
	Username *[]string `form:"username" binding:"omitempty"`
	Email    *[]string `form:"email" binding:"omitempty"`
	Status   *string   `form:"status" binding:"omitempty,oneof=Active Inactive"`
	RoleId   *[]string `form:"roleId" binding:"omitempty,list_uuid"`
	GroupId  *[]string `form:"groupId" binding:"omitempty,list_uuid"`
	AuthType *uint8    `form:"authType" binding:"omitempty,gte=0,lte=4"`
}

type User struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Status    string     `json:"status"`
	Avatar    *string    `json:"avatar"`
	AuthType  uint8      `json:"authType"`
	Group     GroupShort `json:"group"`
	Role      RoleShort  `json:"role"`
}

type UserList []*User

type UserMe struct {
	Id        string     `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Status    string     `json:"status"`
	Avatar    *string    `json:"avatar"`
	AuthType  uint8      `json:"authType"`
	Group     GroupShort `json:"group"`
	Role      RoleShort  `json:"role"`
	// Menus     MenuList   `json:"menus"`
}

type UserUpdateMe struct {
	Username *string `json:"username" binding:"omitempty"`
	Email    *string `json:"email" binding:"omitempty,email"`
	Password *string `json:"password" binding:"omitempty"`
	Avatar   *string `json:"avatar" binding:"omitempty"`
}
