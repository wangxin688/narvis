package schemas

import (
	"time"

	"github.com/wangxin688/narvis/server/tools/schemas"
)

type ProviderCreate struct {
	Name        string  `json:"name" binding:"required"`
	Icon        *string `json:"icon" binding:"omitempty"`
	Description *string `json:"description" binding:"omitempty"`
}

type ProviderUpdate struct {
	Name        *string `json:"name" binding:"omitempty"`
	Icon        *string `json:"icon" binding:"omitempty"`
	Description *string `json:"description" binding:"omitempty"`
}

type ProviderQuery struct {
	schemas.PageInfo
	Name *[]string `form:"name" binding:"omitempty"`
}

type Provider struct {
	Id          string    `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	Name        string    `json:"name"`
	Icon        *string   `json:"icon"`
	Description *string   `json:"description"`
}

type ProviderList []Provider

type ProviderShort struct {
	Id   string  `json:"id"`
	Name string  `json:"name"`
	Icon *string `json:"icon"`
}

type ProviderShortList []ProviderShort
