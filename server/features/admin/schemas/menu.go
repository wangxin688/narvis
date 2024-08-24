package schemas

import "time"

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

type MenuCreate struct {
	Path     string  `json:"path" binding:"required"`
	Name     string  `json:"name" binding:"required"`
	Redirect *string `json:"redirect"`
	Meta     *Meta   `json:"meta"`
	ParentID *string `json:"parent_id" binding:"omitempty,uuid"`
}

type MenuUpdate struct {
	Path     *string `json:"path" binding:"omitempty"`
	Name     *string `json:"name" binding:"omitempty"`
	Redirect *string `json:"redirect" binding:"omitempty"`
	Meta     *Meta   `json:"meta" binding:"omitempty"`
	ParentID *string `json:"parent_id" binding:"omitempty,uuid"`
}

type MenuTree struct {
	ID        string      `json:"id"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Time   `json:"updated_at"`
	Path      string      `json:"path"`
	Name      string      `json:"name"`
	Redirect  *string     `json:"redirect"`
	ParentID  *string     `json:"parent_id"`
	Meta      *Meta       `json:"meta"`
	Children  []*MenuTree `json:"children"`
}
