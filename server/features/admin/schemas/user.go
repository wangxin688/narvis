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
	GroupID  string  `json:"group_id" binding:"required,uuid"`
	RoleID   string  `json:"role_id" binding:"required,uuid"`
	AuthType uint8   `json:"auth_type" binding:"required,gte=0,lte=4"`
}

type UserUpdate struct {
	Username *string `json:"username" binding:"omitempty"`
	Email    *string `json:"email" binding:"omitempty,email"`
	Password *string `json:"password" binding:"omitempty"`
	Status   *string `json:"status" binding:"omitempty,oneof=Active Inactive"`
	Avatar   *string `json:"avatar" binding:"omitempty"`
	GroupID  *string `json:"group_id" binding:"omitempty,uuid"`
	RoleID   *string `json:"role_id" binding:"omitempty,uuid"`
}

type UserQuery struct {
	schemas.PageInfo
	ID       *[]string `form:"id" binding:"omitempty,list_uuid"`
	Username *[]string `form:"username" binding:"omitempty"`
	Email    *[]string `form:"email" binding:"omitempty"`
	Status   *string   `form:"status" binding:"omitempty,oneof=Active Inactive"`
	RoleID   *[]string `form:"role_id" binding:"omitempty,list_uuid"`
	GroupID  *[]string `form:"group_id" binding:"omitempty,list_uuid"`
	AuthType *uint8    `form:"auth_type" binding:"omitempty,gte=0,lte=4"`
}

type User struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Status    string     `json:"status"`
	Avatar    *string    `json:"avatar"`
	AuthType  uint8      `json:"auth_type"`
	Group     GroupShort `json:"group"`
	Role      RoleShort  `json:"role"`
}

type UserList []*User
