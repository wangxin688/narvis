package models

import "gorm.io/datatypes"

const TaskResultTableName = "task_result"

type Result struct {
	Data   any
	Errors []string
}

type TaskResult struct {
	BaseDbModel
	Name           string                     `gorm:"column:name;not null"`
	Status         string                     `gorm:"column:status;default:InProgress"` // InProgress, Success, Failed
	SubTaskId      *string                    `gorm:"column:subTaskId;type:uuid;"`
	Result         datatypes.JSONType[Result] `gorm:"column:result;type:json"`
	ProxyId        *string                     `gorm:"column:proxyId;type:uuid;default:null"`
	Proxy          Proxy                      `gorm:"constraint:Ondelete:SET NULL"`
	OrganizationId string                     `gorm:"column:organizationId;type:uuid;index"`
	Organization   Organization               `gorm:"constraint:Ondelete:CASCADE"`
}

func (TaskResult) TableName() string {
	return TaskResultTableName
}
