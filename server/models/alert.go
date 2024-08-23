package models

import (
	"time"

	"github.com/wangxin688/narvis/server/global"
	"gorm.io/datatypes"
)

type Tag struct {
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type Condition struct {
	Item  string   `json:"item"`
	Value []string `json:"value"`
}

type ChannelConfig struct {
	WebhookUrl     *string            `json:"webhook_url"`
	WebhookHeaders *map[string]string `json:"webhook_headers"`
	Email          *string            `json:"email"`
}

type Alert struct {
	ID                string                   `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Status            uint8                    `gorm:"default:0"` // 0: firing 1: resolved
	StartedAt         time.Time                `gorm:"autoCreateTime;not null"`
	ResolvedAt        *time.Time               `gorm:"default:null"`
	Acknowledged      bool                     `gorm:"default:false"`
	Suppressed        bool                     `gorm:"default:false"`
	Inhibited         bool                     `gorm:"default:true"`
	Severity          uint8                    `gorm:"default:0"` // 0: info 1: warning 2: critical 3: disaster
	Duration          *string                  `gorm:"-"`
	AlertName         string                   `gorm:"not null;index"`
	Tag               datatypes.JSONSlice[Tag] `gorm:"type:json;default:null"`
	EventId           string                   `gorm:"type:string;index"`
	TriggerId         string                   `gorm:"type:string"`
	UserId            *string                  `gorm:"type:uuid;default:null"`
	User              User                     `gorm:"constraint:Ondelete:SET NULL"`
	SiteId            string                   `gorm:"type:uuid;not null;index"`
	Site              Site                     `gorm:"constraint:Ondelete:CASCADE"`
	DeviceId          *string                  `gorm:"type:uuid;index"`
	Device            Device                   `gorm:"constraint:Ondelete:CASCADE"`
	ApId              *string                  `gorm:"type:uuid;index"`
	Ap                AP                       `gorm:"constraint:Ondelete:CASCADE"`
	CircuitId         *string                  `gorm:"type:uuid;index"`
	Circuit           Circuit                  `gorm:"constraint:Ondelete:CASCADE"`
	DeviceInterfaceId *string                  `gorm:"type:uuid;index"`
	DeviceInterface   DeviceInterface          `gorm:"constraint:Ondelete:SET NULL"`
	MaintenanceId     *string                  `gorm:"type:uuid;default:null"`
	Maintenance       Maintenance              `gorm:"constraint:Ondelete:SET NULL"`
	OrganizationId    string                   `gorm:"type:uuid;not null;index"`
	Organization      Organization             `gorm:"constraint:Ondelete:CASCADE"`
}

type AlertGroup struct {
	ID             string       `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Status         uint8        `gorm:"default:0"` // 0: firing 1: resolved
	StartedAt      time.Time    `gorm:"autoCreateTime;not null"`
	ResolvedAt     *time.Time   `gorm:"default:null"`
	Acknowledged   bool         `gorm:"default:false"`
	Suppressed     bool         `gorm:"default:false"`
	Severity       uint8        `gorm:"default:0"` // 0: info 1: warning 2: critical 3: disaster
	Duration       *string      `gorm:"-"`
	AlertName      string       `gorm:"not null;index"`
	GroupKey       string       `gorm:"not null"`
	HashKey        string       `gorm:"not null"`
	SiteId         string       `gorm:"type:uuid;not null;index"`
	Site           Site         `gorm:"constraint:Ondelete:CASCADE"`
	OrganizationId string       `gorm:"type:uuid;not null;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

type ActionLog struct {
	global.BaseDbModel

	Acknowledged *bool        `gorm:"default:null"`
	Resolved     *bool        `gorm:"default:null"`
	Suppressed   *bool        `gorm:"default:null"`
	Comment      *string      `gorm:"default:null"`
	AssignUserId *string      `gorm:"type:uuid;default:null"`
	AssignUser   User         `gorm:"constraint:Ondelete:SET NULL;foreignKey:AssignUserId"`
	CreatedById  string       `gorm:"type:uuid"`
	CreatedBy    User         `gorm:"constraint:Ondelete:SET NULL;foreignKey:CreatedById"`
	Alert        []Alert      `gorm:"many2many:alert_action_logs"`
	AlertGroup   []AlertGroup `gorm:"many2many:alert_group_action_logs"`
}

type Maintenance struct {
	global.BaseDbModel

	Name            string `gorm:"not null"`
	StartedAt       time.Time
	EndedAt         *time.Time
	MaintenanceType string                         `gorm:"not null"`
	Conditions      datatypes.JSONSlice[Condition] `gorm:"type:json;not null"`
	Description     *string
	OrganizationId  string       `gorm:"type:uuid;not null;index"`
	Organization    Organization `gorm:"constraint:Ondelete:CASCADE"`
	Alert           []Alert
}

type RootCause struct {
	global.BaseDbModel
	Name           string       `gorm:"not null;uniqueIndex:idx_name_organization_id"`
	Description    *string      `gorm:"default:null"`
	Category       *string      `gorm:"default:null"`
	OrganizationId string       `gorm:"type:uuid;uniqueIndex:idx_name_organization_id;not null"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

type Subscription struct {
	global.BaseDbModel

	Name           string                            `gorm:"not null"`
	Enabled        bool                              `gorm:"not null;default:true"`
	Deduplication  bool                              `gorm:"not null;default:true"`
	Conditions     datatypes.JSONType[Condition]     `gorm:"type:json;not null"`
	SendResolved   bool                              `gorm:"not null;default:true"`
	RepeatInterval int                               `gorm:"not null;default:0"`
	ChannelType    uint8                             `gorm:"not null;default:0"`
	ChannelConfig  datatypes.JSONType[ChannelConfig] `gorm:"type:json;not null"`
	CreatedById    string                            `gorm:"type:uuid;not null;index"`
	CreatedBy      User                              `gorm:"constraint:Ondelete:SET NULL;foreignKey:CreatedById"`
	OrganizationId string                            `gorm:"type:uuid;not null;index"`
	Organization   Organization                      `gorm:"constraint:Ondelete:CASCADE"`
}

type SubscriptionRecord struct {
	global.BaseDbSingleModel
	SubscriptionId string  `gorm:"type:uuid;not null;index"`
	AlertId        *string `gorm:"type:uuid;default:null;index"`
	AlertGroup     *string `gorm:"type:uuid;default:null;index"`
}
