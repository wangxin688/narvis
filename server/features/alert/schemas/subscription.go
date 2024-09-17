package schemas

import (
	"time"

	"github.com/wangxin688/narvis/server/features/admin/schemas"
	ts "github.com/wangxin688/narvis/server/tools/schemas"
)

type ChannelConfig struct {
	WebhookUrl     *string            `json:"webhook" binding:"omitempty"`
	WebhookHeaders *map[string]string `json:"webhook_headers" binding:"omitempty"`
	Email          *string            `json:"email" binding:"omitempty,email"`
}

type SubscriptionCreate struct {
	Name           string        `json:"name" binding:"required"`
	Enabled        bool          `json:"enabled" binding:"required"`
	Deduplication  bool          `json:"deduplication" binding:"required"`
	Conditions     []Condition   `json:"conditions" binding:"required"`
	SendResolved   bool          `json:"sendResolved" binding:"required"`
	RepeatInterval uint32        `json:"repeatInterval" binding:"required"`
	ChannelType    string        `json:"channelType" binding:"required, oneof=Webhook Email"`
	ChannelConfig  ChannelConfig `json:"channelConfig" binding:"required"`
}

type SubscriptionUpdate struct {
	Name           *string        `json:"name" binding:"omitempty"`
	Enabled        *bool          `json:"enabled" binding:"omitempty"`
	Deduplication  *bool          `json:"deduplication" binding:"omitempty"`
	Conditions     *[]Condition   `json:"conditions" binding:"omitempty"`
	SendResolved   *bool          `json:"sendResolved" binding:"omitempty"`
	RepeatInterval *uint32        `json:"repeatInterval" binding:"omitempty"`
	ChannelType    *string        `json:"channelType" binding:"omitempty,oneof=Webhook Email"`
	ChannelConfig  *ChannelConfig `json:"channelConfig" binding:"omitempty"`
}

type Subscription struct {
	Id             string             `json:"id"`
	CreatedAt      time.Time          `json:"createdAt"`
	UpdatedAt      *time.Time         `json:"updatedAt"`
	Name           string             `json:"name"`
	Enabled        bool               `json:"enabled"`
	Deduplication  bool               `json:"deduplication"`
	Conditions     []Condition        `json:"conditions"`
	SendResolved   bool               `json:"sendResolved"`
	RepeatInterval uint32             `json:"repeatInterval"`
	ChannelType    string             `json:"channelType"`
	ChannelConfig  ChannelConfig      `json:"channelConfig"`
	CreatedBy      schemas.UserShort  `json:"createdBy"`
	UpdatedBy      *schemas.UserShort `json:"updatedBy"`
}

type SubscriptionQuery struct {
	ts.PageInfo
	Id          *[]string `json:"id" binding:"omitempty,list_uuid"`
	Name        *string   `json:"name" binding:"omitempty"`
	Enabled     *bool     `json:"enabled" binding:"omitempty"`
	ChannelType *string   `json:"channelType" binding:"omitempty,oneof=Webhook Email"`
}
