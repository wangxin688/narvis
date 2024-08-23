package am

import "time"

type Matcher struct {
	Name    string `json:"name"`
	Value   string `json:"value"`
	IsRegex bool   `json:"isRegex"`
	IsEqual bool   `json:"isEqual"`
}

type Alert struct {
	Labels       map[string]string `json:"labels"`
	Annotations  map[string]string `json:"annotations"`
	StartsAt     time.Time         `json:"startsAt"` // rfc3339
	EndsAt       time.Time         `json:"endsAt"`   // rfc3339
	GeneratorURL string            `json:"generatorURL"`
	Status       string            `json:"status"`
}

type AlertStatus struct {
	State       string   `json:"state"` // unprocessed, pending, firing, resolved
	SilencedBy  []string `json:"silencedBy"`
	InhibitedBy []string `json:"inhibitedBy"`
}

type AlertResponse struct {
	Labels       map[string]string `json:"labels"`
	Annotations  map[string]string `json:"annotations"`
	StartsAt     time.Time         `json:"startsAt"`  // rfc3339
	EndsAt       time.Time         `json:"endsAt"`    // rfc3339
	UpdatedAt    time.Time         `json:"updatedAt"` // rfc3339
	Status       AlertStatus       `json:"status"`
	Fingerprint  string            `json:"fingerprint"`
	GeneratorURL string            `json:"generatorURL"`
}

type AlertGroupResponse struct {
	Labels map[string]string `json:"labels"`
	Alerts []Alert           `json:"alerts"`
}

type AlertRequest struct {
	Active      *bool    `json:"active,omitempty"`
	Silenced    *bool    `json:"silenced,omitempty"`
	Inhibited   *bool    `json:"inhibited,omitempty"`
	Unprocessed *bool    `json:"unprocessed,omitempty"`
	Filter      []string `json:"filter,omitempty"`
}

type AlertSilenceCreate struct {
	Matchers  []Matcher `json:"matchers"`
	StartsAt  time.Time `json:"startsAt"`
	EndsAt    time.Time `json:"endsAt"`
	CreatedBy string    `json:"createdBy"`
	Comment   string    `json:"comment"`
}

type AlertSilenceUpdate struct {
	ID string `json:"id"`
	AlertSilenceCreate
}
