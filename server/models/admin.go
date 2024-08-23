package models

import (
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/schemas"
	"gorm.io/datatypes"
)

var UserSearchFields = []string{"username", "email"}
var RoleSearchFields = []string{"name"}
var GroupSearchFields = []string{"name"}

var UserTableName = "sys_user"
var RoleTableName = "sys_role"
var GroupTableName = "sys_group"
var MenuTableName = "sys_menu"
var PermissionTableName = "sys_permission"

type Transition struct {
	Name            string `json:"name"`
	EnterTransition string `json:"enterTransition"`
	LeaveTransition string `json:"leaveTransition"`
}

type Meta struct {
	Title        string     `json:"title"`
	Icon         string     `json:"icon"`
	ExtraIcon    *string    `json:"extraIcon"`
	ShowLink     bool       `json:"showLink"`
	ShowParent   bool       `json:"showParent"`
	Rank         uint16     `json:"rank"`
	Roles        []string   `json:"roles"`
	Auths        []string   `json:"auths"`
	KeepAlive    bool       `json:"keepAlive"`
	FrameSrc     *string    `json:"frameSrc"`
	FrameLoading bool       `json:"frameLoading"`
	Transition   Transition `json:"transition"`
	HiddenTag    bool       `json:"hiddenTag"`
	DynamicLevel uint16     `json:"dynamicLevel"`
	ActivePath   string     `json:"activePath"`
}

type User struct {
	global.BaseDbModel
	Username       string       `gorm:"not null"`
	Email          string       `gorm:"uniqueIndex:idx_user_email_organization_id;not null;index"`
	Password       string       `gorm:"not null"`
	Avatar         *string      `gorm:"default:null"`
	GroupID        string       `gorm:"type:uuid;not null"`
	Group          Group        `gorm:"constraint:Ondelete:RESTRICT"`
	RoleID         string       `gorm:"type:uuid,not null"`
	Role           Role         `gorm:"constraint:Ondelete:RESTRICT"`
	AuthType       uint8        `gorm:"type:smallint;default:0"`
	OrganizationID string       `gorm:"type:uuid;uniqueIndex:idx_user_email_organization_id;not null"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (User) TableName() string {
	return UserTableName
}

type Role struct {
	global.BaseDbModel
	Name           string       `gorm:"uniqueIndex:idx_role_name_organization_id;not null"`
	Description    *string      `gorm:"default:null"`
	Menus          Menu         `gorm:"many2many:role_menus"`
	OrganizationID string       `gorm:"type:uuid;uniqueIndex:idx_role_name_organization_id;not null"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (Role) TableName() string {
	return RoleTableName
}

type Group struct {
	global.BaseDbModel
	Name           string       `gorm:"uniqueIndex:idx_group_name_organization_id;not null"`
	Description    *string      `gorm:"default:null"`
	RoleID         string       `gorm:"type:uuid;not null"`
	Role           Role         `gorm:"constraint:Ondelete:RESTRICT"`
	OrganizationID string       `gorm:"tye:uuid;uniqueIndex:idx_group_name_organization_id;not null"`
	Organization   Organization `gorm:"constraint:Ondelete:CASCADE"`
}

func (Group) TableName() string {
	return GroupTableName
}

type Permission struct {
	global.BaseDbModel
	Name        string                            `gorm:"unique;not null"`
	Path        string                            `gorm:"unique;not null"`
	Method      string                            ``
	Tag         *string                           `` // need update from api
	Description *datatypes.JSONType[schemas.I18n] `` // need update from api
	Menu        []*Menu                           `gorm:"many2many:menu_permissions"`
}

func (Permission) TableName() string {
	return PermissionTableName
}

type Menu struct {
	global.BaseDbModel
	Path       string                   `gorm:"unique;not null"`
	Name       string                   `gorm:"not null"`
	Redirect   *string                  `gorm:"default:null"`
	ParentID   *string                  `gorm:"default:null;type:uuid"`
	Parent     *Menu                    `gorm:"constraint:Ondelete:RESTRICT;references:ID"`
	Meta       datatypes.JSONType[Meta] `gorm:"type:json"`
	Permission []*Permission            `gorm:"many2many:menu_permissions"`
}

// TODO: confirm the menu design
// 1. Menu的priority仅给前端使用，让用户在选择页面权限时能够展示页面具有哪些权限
// 2. Menu和Permission关联时，移除过多的约束，由超级管理员自行管理保证mapping的准确性，以此降低整个数据模型和开发的复杂度
// 3. 获取Role的Permission时，通过Role关联的Menu所关联的Permission来获取权限列表，不通过Role关联的Permission来获取权限列表（role join menu join permission），牺牲较少量的查询性能来降低开发和数据维护的复杂度
// 4. 如果需要单独的API授权能力，增加长期Token表直接和Permission管理
// 5. AuditLog的钩子函数暂时不集成
// 6. Admin白名单和和黑名单路径由于gin框架的设计，统一使用route.FullPath来做
