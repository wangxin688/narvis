package schemas

import (
	"time"

	"github.com/wangxin688/narvis/server/tools/schemas"
)

type SiteCreate struct {
	Name        string  `json:"name" binding:"required"`
	SiteCode    string  `json:"siteCode" binding:"required"`
	Status      string  `json:"status" binding:"required,oneof=Active Inactive"`
	Region      string  `json:"region" binding:"required"`
	TimeZone    string  `json:"timeZone" binding:"required"`
	Latitude    float32 `json:"latitude" binding:"required,latitude"`
	Longitude   float32 `json:"longitude" binding:"required,longitude"`
	Address     string  `json:"address" binding:"required"`
	Description *string `json:"description"`
}

type SiteUpdate struct {
	Name        *string  `json:"name" binding:"omitempty"`
	SiteCode    *string  `json:"siteCode" binding:"omitempty"`
	Region      *string  `json:"region" binding:"omitempty"`
	TimeZone    *string  `json:"timeZone" binding:"omitempty"`
	Latitude    *float32 `json:"latitude" binding:"omitempty,latitude"`
	Longitude   *float32 `json:"longitude" binding:"omitempty,longitude"`
	Address     *string  `json:"address" binding:"omitempty"`
	Description *string  `json:"description" binding:"omitempty"`
}

type SiteQuery struct {
	schemas.PageInfo
	Id       *[]string `form:"id" binding:"omitempty,list_uuid"`
	Name     *[]string `form:"name" binding:"omitempty"`
	SiteCode *[]string `form:"siteCode" binding:"omitempty"`
	Region   *[]string `form:"region" binding:"omitempty"`
	Status   *string   `form:"status" binding:"omitempty,oneof=Active Inactive"`
}

type Site struct {
	Id          string    `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Name        string    `json:"name"`
	SiteCode    string    `json:"siteCode"`
	Status      string    `json:"status"`
	Region      string    `json:"region"`
	TimeZone    string    `json:"timeZone"`
	Latitude    float32   `json:"latitude"`
	Longitude   float32   `json:"longitude"`
	Address     string    `json:"address"`
	Description *string   `json:"description"`
}

type SiteDetail struct {
	Site
	SwitchCount  int64 `json:"switchCount"`
	ApCount      int64 `json:"apCount"`
	RackCount    int64 `json:"rackCount"`
	CircuitCount int64 `json:"circuitCount"`
	VlanCount    int64 `json:"vlanCount"`
	GatewayCount int64 `json:"gatewayCount"`
}

type SiteList []Site

type SiteShort struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	SiteCode string `json:"siteCode"`
}

type SiteShortList []SiteShort
