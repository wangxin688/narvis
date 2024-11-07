package schemas

import "github.com/wangxin688/narvis/server/tools/schemas"

type RackCreate struct {
	Name         string  `json:"name" binding:"required"`
	SerialNumber *string `json:"serialNumber" binding:"omitempty"`
	UHeight      *uint8  `json:"uHeight" binding:"omitempty,gt=1,lte=50"` // default 42
	SiteId       string  `json:"siteId" binding:"uuid"`
}

type RackUpdate struct {
	Name         *string `json:"name" binding:"omitempty"`
	SerialNumber *string `json:"serialNumber" binding:"omitempty"`
	UHeight      *uint8  `json:"uHeight" binding:"omitempty,gt=1,lte=50"`
}

type RackQuery struct {
	schemas.PageInfo
	Id           *[]string `form:"id" binding:"omitempty,list_uuid"`
	Name         *[]string `form:"name" binding:"omitempty"`
	SerialNumber *[]string `form:"serialNumber" binding:"omitempty"`
	SiteId       *string   `form:"siteId" binding:"omitempty,uuid"`
}

type Rack struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	SerialNumber *string `json:"serialNumber"`
	UHeight      uint8   `json:"uHeight"`
	SiteId       string  `json:"siteId"`
}

type RackList []Rack

type RackShort struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type RackShortList []RackShort

type RackElevation struct {
	Id                 string               `json:"id"`
	Name               string               `json:"name"`
	SerialNumber       *string              `json:"serialNumber"`
	UHeight            uint8                `json:"uHeight"`
	SiteId             string               `json:"siteId"`
	Items              []*RackElevationItem `json:"items"`
	AvailablePositions []uint8              `json:"availablePositions"`
}

type RackElevationItem struct {
	Id              string  `json:"deviceId"`
	Name            string  `json:"name"`
	ManagementIp    string  `json:"managementIp"`
	DeviceRole      string  `json:"deviceRole"`
	OperatingStatus string  `json:"operatingStatus"`
	Position        []uint8 `json:"position"`
}
