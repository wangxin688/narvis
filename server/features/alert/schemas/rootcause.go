package schemas

import "time"

type RootCause struct {
	Id          string     `json:"id"`
	Name        string     `json:"name"`
	Description *string    `json:"description"`
	Category    *string    `json:"category" `
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   *time.Time `json:"updatedAt"`
}

type RootCauseCreate struct {
	Name        string  `json:"name" binding:"required"`
	Description *string `json:"description" binding:"omitempty"`
	Category    *string `json:"category" binding:"omitempty"`
}

type RootCauseUpdate struct {
	Name        *string `json:"name" binding:"omitempty"`
	Description *string `json:"description" binding:"omitempty"`
	Category    *string `json:"category" binding:"omitempty"`
}
