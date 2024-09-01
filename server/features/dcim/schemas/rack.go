package schemas

import "github.com/wangxin688/narvis/server/tools/schemas"

type RackCreate struct {
	Name         string  `json:"name" binding:"required"`
	SerialNumber *string `json:"serialNumber" binding:"omitempty"`
	UHeight      *uint8  `json:"uHeight" binding:"omitempty"`  // default 42
	DescUnit     *bool   `json:"descUnit" binding:"omitempty"` // default true
	SiteId       string  `json:"siteId" binding:"uuid"`
}

type RackUpdate struct {
	Name *string `json:"name" binding:"omitempty"`
}

type RackQuery struct {
	schemas.PageInfo
	Id           *[]string `form:"id" binding:"omitempty,list_uuid"`
	Name         *[]string `form:"name" binding:"omitempty"`
	SerialNumber *string   `form:"serialNumber" binding:"omitempty"`
}

type Rack struct {
	Id           string        `json:"id"`
	Name         string        `json:"name"`
	SerialNumber *string       `json:"serialNumber"`
	UHeight      uint8         `json:"uHeight"`
	Site         SiteShort     `json:"site"`
}

type RackList []Rack

type RackShort struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type RackShortList []RackShort
