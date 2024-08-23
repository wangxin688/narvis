package schemas

import "github.com/wangxin688/narvis/server/schemas"

type UserCreate struct {
	Username string  `json:"username" binding:"required"`
	Email    string  `json:"email" binding:"required;email"`
	Password string  `json:"password" `
	Avatar   *string `json:"avatar"`
	GroupId  string  `json:"group_id" binding:"required,uuid"`
	RoleId   string  `json:"role_id" binding:"required,uuid"`
	AuthType uint8   `json:"auth_type" binding:"required,gte=0,lte=4"`
}

type UserUpdate struct {
	Username *string `json:"username" binding:"omitempty"`
	Email    *string `json:"email" binding:"omitempty,email"`
	Password *string `json:"password" binding:"omitempty"`
	Avatar   *string `json:"avatar" binding:"omitempty"`
	GroupId  *string `json:"group_id" binding:"omitempty,uuid"`
	RoleId   *string `json:"role_id" binding:"omitempty,uuid"`
}

type User struct {
	schemas.BaseResponse
	Username string     `json:"username"`
	Email    string     `json:"email"`
	Avatar   *string    `json:"avatar"`
	AuthType uint8      `json:"auth_type"`
	Group    GroupShort `json:"group"`
	Role     RoleShort  `json:"role"`
}
