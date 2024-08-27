package schemas

import "github.com/wangxin688/narvis/server/tools/schemas"

type RackCreate struct {
	Name         string  `json:"name" binding:"required"`
	AssetTag     *string `json:"asset_tag" binding:"omitempty"`
	SerialNumber *string `json:"serial_number" binding:"omitempty"`
	UHeight      *uint8  `json:"u_height" binding:"omitempty"`      // default 42
	Height       float32 `json:"height" binding:"omitempty"`        // default 2
	Width        float32 `json:"width" binding:"omitempty"`         // default 0.6
	Depth        float32 `json:"depth" binding:"omitempty"`         // default 0.8
	DescUnit     *bool   `json:"desc_unit" binding:"omitempty"`     // default true
	LocationID   *string `json:"location_id" binding:"omitempty,uuid"`
	SiteID       string  `json:"site_id" binding:"uuid"`
}

type RackUpdate struct {
	Name         *string  `json:"name" binding:"omitempty"`
	AssetTag     *string  `json:"asset_tag" binding:"omitempty"`
	SerialNumber *string  `json:"serial_number" binding:"omitempty"`
	UHeight      *uint8   `json:"u_height" binding:"omitempty"`
	Height       *float32 `json:"height" binding:"omitempty"`
	Width        *float32 `json:"width" binding:"omitempty"`
	Depth        *float32 `json:"depth" binding:"omitempty"`
	DescUnit     *bool    `json:"desc_unit" binding:"omitempty"`
	LocationID   *string  `json:"location_id" binding:"omitempty,uuid"`
	SiteID       *string  `json:"site_id" binding:"omitempty,uuid"`
}

type RackQuery struct {
	schemas.PageInfo
	ID           *[]string `form:"id" binding:"omitempty,list_uuid"`
	Name         *[]string `form:"name" binding:"omitempty"`
	SerialNumber *string   `form:"serial_number" binding:"omitempty"`
	AssetTag     *[]string `form:"asset_tag" binding:"omitempty"`
}

type Rack struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	AssetTag     *string       `json:"asset_tag"`
	SerialNumber *string       `json:"serial_number"`
	UHeight      uint8         `json:"u_height"`
	Height       float32       `json:"height"`
	Width        float32       `json:"width"`
	Depth        float32       `json:"depth"`
	DescUnit     bool          `json:"desc_unit"`
	Location     LocationShort `json:"location"`
	Site         SiteShort     `json:"site"`
}

type RackList []Rack

type RackShort struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type RackShortList []RackShort

