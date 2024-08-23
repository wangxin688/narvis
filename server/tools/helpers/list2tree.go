package helpers

type TreeNodeInterface[T any] interface {
	GetID() T
	GetParentID() *T
	SetChildren([]TreeNodeInterface[T])
	GetChildren() []TreeNodeInterface[T]
}

// Category 是一个 GORM 模型，表示数据库中的一个类别
/*
type Category struct {
	ID       string  `gorm:"primaryKey"`
	ParentID *string `gorm:"column:parent_id"`
	Children []*Category `gorm:"-"`
}

func (c *Category) GetID() string {
	return c.ID
}

func (c *Category) GetParentID() *string {
	return c.ParentID
}

func (c *Category) SetChildren(children []TreeNodeInterface[string]) {
	c.Children = []*Category{}
	for _, child := range children {
		c.Children = append(c.Children, child.(*Category))
	}
}

func (c *Category) GetChildren() []TreeNodeInterface[string] {
	children := []TreeNodeInterface[string]{}
	for _, child := range c.Children {
		children = append(children, child)
	}
	return children
}

var categories []Category
	db.Find(&categories)

	// 构建树结构
	categoryInterfaces := make([]TreeNodeInterface[string], len(categories))
	for i := range categories {
		categoryInterfaces[i] = &categories[i]
	}
	tree := BuildTree(categoryInterfaces)

	// 打印树结构
	result, _ := json.MarshalIndent(tree, "", "  ")
	fmt.Println(string(result))
func main() {
	// 初始化 GORM 和 SQLite 数据库（示例用）
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移数据库
	db.AutoMigrate(&Category{})

	// 插入一些示例数据
	db.Create(&Category{ID: "1", ParentID: nil})
	db.Create(&Category{ID: "2", ParentID: ptr("1")})
	db.Create(&Category{ID: "3", ParentID: ptr("1")})
	db.Create(&Category{ID: "4", ParentID: ptr("2")})
	db.Create(&Category{ID: "5", ParentID: nil})
	db.Create(&Category{ID: "6", ParentID: ptr("5")})

	// 查询所有类别
	var categories []Category
	db.Find(&categories)

	// 构建树结构
	categoryInterfaces := make([]TreeNodeInterface[string], len(categories))
	for i := range categories {
		categoryInterfaces[i] = &categories[i]
	}
	tree := BuildTree(categoryInterfaces)

	// 打印树结构
	result, _ := json.MarshalIndent(tree, "", "  ")
	fmt.Println(string(result))
}

// 辅助函数，用于获取字符串指针
func ptr(s string) *string {
	return &s
}
*/

func List2Tree[T comparable, N TreeNodeInterface[T]](nodes []N) []N {
	idMap := make(map[T]N)

	var roots []N
	for _, node := range nodes {
		idMap[node.GetID()] = node
	}
	for _, node := range nodes {
		parentID := node.GetParentID()
		if parentID != nil {
			if parentNode, ok := idMap[*parentID]; ok {
				parentNode.SetChildren(append(parentNode.GetChildren(), node))
			}
		} else {
			roots = append(roots, node)
		}
	}
	return roots
}
