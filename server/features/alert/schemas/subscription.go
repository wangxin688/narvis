package schemas

type ChannelConfig struct {
	WebhookUrl     *string            `json:"webhook" binding:"omitempty"`
	WebhookHeaders *map[string]string `json:"webhook_headers" binding:"omitempty"`
}

type SubscriptionCreate struct {
	Name           string        `json:"name" binding:"required"`
	Enabled        bool          `json:"enabled" binding:"required"`
	Deduplication  bool          `json:"deduplication" binding:"required"`
	Conditions     []Condition   `json:"conditions" binding:"required"`
	SendResolved   bool          `json:"sendResolved" binding:"required"`
	RepeatInterval int           `json:"repeatInterval" binding:"required"`
	ChannelType    uint8         `json:"channelType" binding:"required"`
	ChannelConfig  ChannelConfig `json:"channelConfig" binding:"required"`
}

type SubscriptionUpdate struct {
	Name           *string        `json:"name" binding:"omitempty"`
	Enabled        *bool          `json:"enabled" binding:"omitempty"`
	Deduplication  *bool          `json:"deduplication" binding:"omitempty"`
	Conditions     *[]Condition   `json:"conditions" binding:"omitempty"`
	SendResolved   *bool          `json:"sendResolved" binding:"omitempty"`
	RepeatInterval *int           `json:"repeatInterval" binding:"omitempty"`
	ChannelType    *uint8         `json:"channelType" binding:"omitempty"`
	ChannelConfig  *ChannelConfig `json:"channelConfig" binding:"omitempty"`
}
