package schemas

import (
	"time"
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
	ParentId *string `json:"parentId" binding:"omitempty,uuid"`
}

type MenuUpdate struct {
	Path     *string `json:"path" binding:"omitempty"`
	Name     *string `json:"name" binding:"omitempty"`
	Redirect *string `json:"redirect" binding:"omitempty"`
	Meta     *Meta   `json:"meta" binding:"omitempty"`
	ParentId *string `json:"parentId" binding:"omitempty,uuid"`
}

type Menu struct {
	Id        string    `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Path      string    `json:"path"`
	Name      string    `json:"name"`
	Redirect  *string   `json:"redirect"`
	ParentId  *string   `json:"parentId"`
	Meta      *Meta     `json:"meta"`
	Children  []*Menu   `json:"children"`
}

func (m *Menu) GetId() string {
	return m.Id
}

func (m *Menu) GetParentId() *string {
	return m.ParentId
}

// func (m *Menu) SetChildren(children []helpers.TreeNodeInterface[string]) {
// 	m.Children = []*Menu{}
// 	for _, child := range children {
// 		m.Children = append(m.Children, child.(*Menu))
// 	}
// }

// func (m *Menu) GetChildren() []helpers.TreeNodeInterface[string] {
// 	children := []helpers.TreeNodeInterface[string]{}
// 	for _, child := range m.Children {
// 		children = append(children, child)
// 	}
// 	return children
// }
