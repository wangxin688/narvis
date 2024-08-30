package models

import (
	"time"

	"gorm.io/datatypes"
)

var AlertTableName = "alert"
var AlertGroupTableName = "alert_group"
var AlertActionLogTableName = "alert_action_log"
var MaintenanceTableName = "alert_maintenance"
var RootCauseTableName = "alert_root_cause"
var SubscriptionTableName = "alert_subscription"
var SubscriptionRecordTableName = "alert_subscription_record"

type Label struct {
	Tag   string `json:"tag"`
	Value string `json:"value"`
}

type Condition struct {
	Item  string   `json:"item"`
	Value []string `json:"value"`
}

type ChannelConfig struct {
	WebhookUrl     *string            `json:"webhookUrl"`
	WebhookHeaders *map[string]string `json:"webhookHeaders"`
	Email          *string            `json:"email"`
}

type Alert struct {
	Id                string                     `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Status            uint8                      `gorm:"column:status;type:smallint;default:0"` // 0: firing 1: resolved
	StartedAt         time.Time                  `gorm:"column:startedAt;autoCreateTime;not null"`
	ResolvedAt        *time.Time                 `gorm:"column:resolvedAt;default:null"`
	Acknowledged      bool                       `gorm:"column:acknowledged;default:false"`
	Suppressed        bool                       `gorm:"column:suppressed;default:false"`
	Inhibited         bool                       `gorm:"column:inhibited;default:true"`
	Severity          uint8                      `gorm:"column:severity;default:0"` // P1 P2 P3 P4
	Duration          *string                    `gorm:"-"`
	AlertName         string                     `gorm:"column:alertName;not null;index"`
	Labels            datatypes.JSONSlice[Label] `gorm:"column:labels;type:json;default:null"`
	EventId           string                     `gorm:"column:eventId;type:string;index"`
	TriggerId         string                     `gorm:"column:triggerId;type:string"`
	UserId            *string                    `gorm:"column:userId;type:uuid;default:null"`
	User              User                       `gorm:"constraint:Ondelete:SET NULL"`
	SiteId            string                     `gorm:"column:siteId;type:uuid;not null;index"`
	Site              Site                       `gorm:"constraint:Ondelete:CASCADE"`
	DeviceId          *string                    `gorm:"column:deviceId;type:uuid;index"`
	Device            Device                     `gorm:"constraint:Ondelete:CASCADE"`
	ApId              *string                    `gorm:"column:apId;type:uuid;index"`
	Ap                AP                         `gorm:"constraint:Ondelete:CASCADE"`
	CircuitId         *string                    `gorm:"column:circuitId;type:uuid;index"`
	Circuit           Circuit                    `gorm:"constraint:Ondelete:CASCADE"`
	DeviceInterfaceId *string                    `gorm:"column:deviceInterfaceId;type:uuid;index"`
	DeviceInterface   DeviceInterface            `gorm:"constraint:Ondelete:SET NULL"`
	MaintenanceId     *string                    `gorm:"column:maintenanceId;type:uuid;default:null"`
	Maintenance       Maintenance                `gorm:"constraint:Ondelete:SET NULL"`
	OrganizationId    string                     `gorm:"column:organizationId;type:uuid;not null;index"`
	Organization      Organization               `gorm:"constraint:Ondelete:CASCADE"`
}

func (Alert) TableName() string {
	return AlertTableName
}

type AlertGroup struct {
	Id             string       `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Status         uint8        `gorm:"column:status;type:smallint;default:0"` // 0: firing 1: resolved
	StartedAt      time.Time    `gorm:"column:startedAt;autoCreateTime;not null"`
	ResolvedAt     *time.Time   `gorm:"column:resolvedAt;default:null"`
	Acknowledged   bool         `gorm:"column:acknowledged;default:false"`
	Suppressed     bool         `gorm:"column:suppressed;default:false"`
	Severity       uint8        `gorm:"column:severity;default:0"` // 0: info 1: warning 2: critical 3: disaster
	Duration       *string      `gorm:"-"`
	AlertName      string       `gorm:"column:alertName;not null;index"`
	GroupKey       string       `gorm:"column:groupKey;not null"`
	HashKey        string       `gorm:"column:hashKey;not null"`
	SiteId         string       `gorm:"column:siteId;type:uuid;not null;index"`
	Site           Site         `gorm:"constraint:Ondelete:CASCADE"`
	OrganizationId string       `gorm:"column:organizationId;type:uuid;not null;index"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (AlertGroup) TableName() string {
	return AlertGroupTableName
}

type AlertActionLog struct {
	BaseDbModel

	Acknowledged *bool        `gorm:"column:acknowledged;default:null"`
	Resolved     *bool        `gorm:"column:resolved;default:null"`
	Suppressed   *bool        `gorm:"column:suppressed;default:null"`
	Comment      *string      `gorm:"column:comment;default:null"`
	AssignUserId *string      `gorm:"column:assignUserId;type:uuid;default:null"`
	AssignUser   User         `gorm:"constraint:Ondelete:SET NULL;foreignKey:AssignUserId"`
	CreatedById  string       `gorm:"column:createdById;type:uuid"`
	CreatedBy    User         `gorm:"constraint:Ondelete:SET NULL;foreignKey:CreatedById"`
	Alert        []Alert      `gorm:"many2many:alert_action_logs"`
	AlertGroup   []AlertGroup `gorm:"many2many:alert_group_action_logs"`
}

func (AlertActionLog) TableName() string {
	return AlertActionLogTableName
}

type Maintenance struct {
	BaseDbModel

	Name            string                         `gorm:"not null"`
	StartedAt       time.Time                      `gorm:"column:startedAt;not null;"`
	EndedAt         time.Time                      `gorm:"column:endedAt;not null"`
	MaintenanceType string                         `gorm:"column:maintenanceType;not null"`
	Conditions      datatypes.JSONSlice[Condition] `gorm:"column:conditions;type:json;not null"`
	Description     *string                        `gorm:"column:description;default:null"`
	OrganizationId  string                         `gorm:"column:organizationId;type:uuid;not null;index"`
	Organization    Organization                   `gorm:"constraint:Ondelete:CASCADE"`
	Alert           []Alert
}

func (Maintenance) TableName() string {
	return MaintenanceTableName
}

type RootCause struct {
	BaseDbModel
	Name           string       `gorm:"column:name;not null;uniqueIndex:idx_name_organization_id"`
	Description    *string      `gorm:"column:description;default:null"`
	Category       *string      `gorm:"column:category;default:null"`
	OrganizationId string       `gorm:"column:organizationId;type:uuid;uniqueIndex:idx_name_organization_id;not null"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (RootCause) TableName() string {
	return RootCauseTableName
}

type Subscription struct {
	BaseDbModel

	Name           string                            `gorm:"column:name;not null"`
	Enabled        bool                              `gorm:"column:enabled;not null;default:true"`
	Deduplication  bool                              `gorm:"column:deduplication;not null;default:true"`
	Conditions     datatypes.JSONType[Condition]     `gorm:"column:conditions;type:json;not null"`
	SendResolved   bool                              `gorm:"column:sendResolved;not null;default:true"`
	RepeatInterval uint32                            `gorm:"column:repeatInterval;type:smallint;not null;default:0"`
	ChannelType    uint8                             `gorm:"column:channelType;type:smallint;not null;default:0"`
	ChannelConfig  datatypes.JSONType[ChannelConfig] `gorm:"column:channelConfig;type:json;not null"`
	CreatedById    string                            `gorm:"column:createdById;type:uuid;not null;index"`
	CreatedBy      User                              `gorm:"constraint:Ondelete:SET NULL;foreignKey:CreatedById"`
	OrganizationId string                            `gorm:"column:organizationId;type:uuid;not null;index"`
	Organization   Organization                      `gorm:"constraint:Ondelete:CASCADE"`
}

func (Subscription) TableName() string {
	return SubscriptionTableName
}

type SubscriptionRecord struct {
	BaseDbSingleModel
	SubscriptionId string  `gorm:"column:subscriptionId;type:uuid;not null;index"`
	AlertId        *string `gorm:"column:alertId;type:uuid;default:null;index"`
	AlertGroup     *string `gorm:"column:alertGroup;type:uuid;default:null;index"`
	Status         uint8   `gorm:"column:status;type:smallint;default:1"` // 0: failed, 1: success
	FailedReason   *string `gorm:"column:failedReason;default:null"`
}

func (SubscriptionRecord) TableName() string {
	return SubscriptionRecordTableName
}
