package schemas

type DeviceRoleQuery struct {
	Name    *string `form:"name" binding:"omitempty"`
	Keyword *string `form:"keyword" binding:"omitempty"`
}
