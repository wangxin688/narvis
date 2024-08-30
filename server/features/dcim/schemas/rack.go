package schemas

import "github.com/wangxin688/narvis/server/tools/schemas"

type RackCreate struct {
	Name         string  `json:"name" binding:"required"`
	AssetTag     *string `json:"assetTag" binding:"omitempty"`
	SerialNumber *string `json:"serialNumber" binding:"omitempty"`
	UHeight      *uint8  `json:"uHeight" binding:"omitempty"`  // default 42
	Height       float32 `json:"height" binding:"omitempty"`   // default 2
	Width        float32 `json:"width" binding:"omitempty"`    // default 0.6
	Depth        float32 `json:"depth" binding:"omitempty"`    // default 0.8
	DescUnit     *bool   `json:"descUnit" binding:"omitempty"` // default true
	LocationId   *string `json:"location_id" binding:"omitempty,uuid"`
	SiteId       string  `json:"site_id" binding:"uuid"`
}

type RackUpdate struct {
	Name         *string  `json:"name" binding:"omitempty"`
	AssetTag     *string  `json:"assetTag" binding:"omitempty"`
	SerialNumber *string  `json:"serialNumber" binding:"omitempty"`
	UHeight      *uint8   `json:"uHeight" binding:"omitempty"`
	Height       *float32 `json:"height" binding:"omitempty"`
	Width        *float32 `json:"width" binding:"omitempty"`
	Depth        *float32 `json:"depth" binding:"omitempty"`
	DescUnit     *bool    `json:"descUnit" binding:"omitempty"`
	LocationId   *string  `json:"locationId" binding:"omitempty,uuid"`
	SiteId       *string  `json:"siteId" binding:"omitempty,uuid"`
}

type RackQuery struct {
	schemas.PageInfo
	Id           *[]string `form:"id" binding:"omitempty,list_uuid"`
	Name         *[]string `form:"name" binding:"omitempty"`
	SerialNumber *string   `form:"serialNumber" binding:"omitempty"`
	AssetTag     *[]string `form:"assetTag" binding:"omitempty"`
}

type Rack struct {
	Id           string        `json:"id"`
	Name         string        `json:"name"`
	AssetTag     *string       `json:"assetTag"`
	SerialNumber *string       `json:"serialNumber"`
	UHeight      uint8         `json:"uHeight"`
	Height       float32       `json:"height"`
	Width        float32       `json:"width"`
	Depth        float32       `json:"depth"`
	DescUnit     bool          `json:"descUnit"`
	Location     LocationShort `json:"location"`
	Site         SiteShort     `json:"site"`
}

type RackList []Rack

type RackShort struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type RackShortList []RackShort
