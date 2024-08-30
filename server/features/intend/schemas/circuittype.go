package schemas

import "github.com/wangxin688/narvis/server/tools/schemas"

type CircuitType struct {
	CircuitType    string       `json:"circuitType"`
	Description    schemas.I18n `json:"description"`
	ConnectionType string       `json:"connectionType"`
}

type CircuitTypeQuery struct {
	CircuitType    *string `form:"circuitType" binding:"omitempty"`
	Description    *string `form:"description" binding:"omitempty"`
	ConnectionType *string `form:"connectionType" binding:"omitempty,oneof=WAN LAN"`
	Keyword        *string `form:"keyword" binding:"omitempty"`
}
