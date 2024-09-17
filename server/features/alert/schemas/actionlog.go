package schemas

import (
	"time"

	as "github.com/wangxin688/narvis/server/features/admin/schemas"
)

type ActionLog struct {
	Id           string        `json:"id"`
	Resolved     *bool         `json:"resolved"`
	CreatedAt    time.Time     `json:"createdAt"`
	Acknowledged *bool         `json:"acknowledged"`
	Suppressed   *bool         `json:"suppressed"`
	AssignUser   *as.UserShort `json:"assignUser"`
	Comment      *string       `json:"comment"`
	RootCause    *RootCause    `json:"rootCause"`
	CreatedBy    *as.UserShort `json:"createdBy"`
	Actions      []string      `json:"actions"`
}

func (a *ActionLog) GenerateActions() {

	actions := make([]string, 0)
	if a.Acknowledged != nil {
		actions = append(actions, "acknowledged")
	}
	if a.Suppressed != nil {
		actions = append(actions, "suppressed")
	}
	if a.RootCause != nil {
		actions = append(actions, "rootCause")
	}
	if a.Resolved != nil {
		actions = append(actions, "resolved")
	}
	if a.Comment != nil {
		actions = append(actions, "comment")
	}
	if a.AssignUser != nil {
		actions = append(actions, "assignUser")
	}
	a.Actions = actions
}

type ActionLogCreate struct {
	AlertId      []string `json:"alertIds" binding:"required,list_uuid"`
	Acknowledged *bool    `json:"acknowledged" binding:"omitempty"`
	Suppressed   *bool    `json:"suppressed" binding:"omitempty"`
	Resolved     *bool    `json:"resolved" binding:"omitempty"`
	Comment      *string  `json:"comment" binding:"omitempty"`
	AssignUserId *string  `json:"assignUserId" binding:"omitempty,uuid"`
	RootCauseId  *string  `json:"rootCauseId" binding:"omitempty,uuid"`
}
