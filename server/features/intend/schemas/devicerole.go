package schemas

import "github.com/wangxin688/narvis/server/tools/schemas"

type DeviceRoleQuery struct {
	Name         *string `form:"name" binding:"omitempty"`
	Abbreviation *string `form:"abbreviation" binding:"omitempty"`
	Keyword      *string `form:"keyword" binding:"omitempty"`
}

type DeviceRole struct {
	DeviceRole    string       `json:"deviceRole"`
	Description   schemas.I18n `json:"description"`
	Weight        uint16       `json:"weight"`
	Abbreviation  string       `json:"abbreviation"`
	ProductFamily string       `json:"productFamily"`
}
