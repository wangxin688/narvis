package schemas

import (
	"github.com/wangxin688/narvis/server/pkg/am"
)

type AlertGroupCreate struct {
	Status            string            `json:"status" binding:"required,oneof=firing resolved"`
	Alerts            []*am.Alert       `json:"alerts" binding:"required"`
	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`
	ExternalURL       string            `json:"externalURL"`
	GroupKey          string            `json:"groupKey"`
}

func (a *AlertGroupCreate) GetAlertLabels() []map[string]string {
	results := make([]map[string]string, 0)
	for _, alert := range a.Alerts {
		results = append(results, alert.Labels)
	}
	return results
}

type Event struct {
	EventId   string `json:"eventId"`
	IsActive  bool   `json:"isActive"`
	Inhibited bool   `json:"inhibited"`
	Severity  string `json:"severity"`
}
