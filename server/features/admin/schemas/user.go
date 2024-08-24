package schemas

import "time"

type UserCreate struct {
	Username string  `json:"username" binding:"required"`
	Email    string  `json:"email" binding:"required;email"`
	Password string  `json:"password" `
	Avatar   *string `json:"avatar"`
	GroupID  string  `json:"group_id" binding:"required,uuid"`
	RoleID   string  `json:"role_id" binding:"required,uuid"`
	AuthType uint8   `json:"auth_type" binding:"required,gte=0,lte=4"`
}

type UserUpdate struct {
	Username *string `json:"username" binding:"omitempty"`
	Email    *string `json:"email" binding:"omitempty,email"`
	Password *string `json:"password" binding:"omitempty"`
	Avatar   *string `json:"avatar" binding:"omitempty"`
	GroupID  *string `json:"group_id" binding:"omitempty,uuid"`
	RoleID   *string `json:"role_id" binding:"omitempty,uuid"`
}

type User struct {
	ID        string     `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Avatar    *string    `json:"avatar"`
	AuthType  uint8      `json:"auth_type"`
	Group     GroupShort `json:"group"`
	Role      RoleShort  `json:"role"`
}
