package schemas

import (
	"time"

	"github.com/wangxin688/narvis/server/tools/helpers"
)

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

type Menu struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Path      string    `json:"path"`
	Name      string    `json:"name"`
	Redirect  *string   `json:"redirect"`
	ParentID  *string   `json:"parent_id"`
	Meta      *Meta     `json:"meta"`
	Children  []*Menu   `json:"children"`
}

func (m *Menu) GetID() string {
	return m.ID
}

func (m *Menu) GetParentID() *string {
	return m.ParentID
}

func (m *Menu) SetChildren(children []helpers.TreeNodeInterface[string]) {
	m.Children = []*Menu{}
	for _, child := range children {
		m.Children = append(m.Children, child.(*Menu))
	}
}

func (m *Menu) GetChildren() []helpers.TreeNodeInterface[string] {
	children := []helpers.TreeNodeInterface[string]{}
	for _, child := range m.Children {
		children = append(children, child)
	}
	return children
}
