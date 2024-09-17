package infra_tasks

import "github.com/wangxin688/narvis/server/features/infra/schemas"

type BaseSnmpTask struct {
	TaskId       string                   `json:"taskId"`
	TaskName     string                   `json:"taskName"`
	SnmpConfig   schemas.SnmpV2Credential `json:"snmpConfig"`
	DeviceId     string                   `json:"deviceId"`
	ManagementIP string                   `json:"managementIp"`
	Callback     string                   `json:"callback"`
}

