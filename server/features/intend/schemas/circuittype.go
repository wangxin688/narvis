package schemas

import "github.com/wangxin688/narvis/server/tools/schemas"

type CircuitType struct {
	CircuitType    string       `json:"circuit_type"`
	Description    schemas.I18n `json:"description"`
	ConnectionType string       `json:"connection_type"`
}

type CircuitTypeQuery struct {
	CircuitType    *string `form:"circuit_type" binding:"omitempty"`
	Description    *string `form:"description" binding:"omitempty"`
	ConnectionType *string `form:"connection_type" binding:"omitempty,oneof=WAN LAN"`
	Keyword        *string `form:"keyword" binding:"omitempty"`
}
