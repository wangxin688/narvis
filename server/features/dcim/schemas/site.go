package schemas

import (
	"time"

	"github.com/wangxin688/narvis/server/tools/schemas"
)

type SiteCreate struct {
	Name        string  `json:"name" binding:"required"`
	SiteCode    string  `json:"site_code" binding:"required"`
	Status      string  `json:"status" binding:"required,oneof=Active Inactive"`
	Region      string  `json:"region" binding:"required"`
	TimeZone    string  `json:"time_zone" binding:"required"`
	Latitude    float32 `json:"latitude" binding:"required"`
	Longitude   float32 `json:"longitude" binding:"required"`
	Address     string  `json:"address" binding:"required"`
	Description *string `json:"description"`
}

type SiteUpdate struct {
	Name        *string  `json:"name" binding:"omitempty"`
	SiteCode    *string  `json:"site_code" binding:"omitempty"`
	Region      *string  `json:"region" binding:"omitempty"`
	TimeZone    *string  `json:"time_zone" binding:"omitempty"`
	Latitude    *float32 `json:"latitude" binding:"omitempty"`
	Longitude   *float32 `json:"longitude" binding:"omitempty"`
	Address     *string  `json:"address" binding:"omitempty"`
	Description *string  `json:"description" binding:"omitempty"`
}

type SiteQuery struct {
	schemas.PageInfo
	ID       *[]string `form:"id" binding:"omitempty,list_uuid"`
	Name     *[]string `form:"name" binding:"omitempty"`
	SiteCode *[]string `form:"site_code" binding:"omitempty"`
	Region   *[]string `form:"region" binding:"omitempty"`
	Status   *string   `form:"status" binding:"omitempty,oneof=Active Inactive"`
}

type Site struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	SiteCode    string    `json:"site_code"`
	Status      string    `json:"status"`
	Region      string    `json:"region"`
	TimeZone    string    `json:"time_zone"`
	Latitude    float32   `json:"latitude"`
	Longitude   float32   `json:"longitude"`
	Address     string    `json:"address"`
	Description *string   `json:"description"`
}

type SiteList []Site

type SiteShort struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	SiteCode string `json:"site_code"`
}

type SiteShortList []SiteShort
