// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package gen

import (
	"context"

	"github.com/wangxin688/narvis/server/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"
)

func newGroup(db *gorm.DB, opts ...gen.DOOption) group {
	_group := group{}

	_group.groupDo.UseDB(db, opts...)
	_group.groupDo.UseModel(&models.Group{})

	tableName := _group.groupDo.TableName()
	_group.ALL = field.NewAsterisk(tableName)
	_group.Id = field.NewString(tableName, "id")
	_group.CreatedAt = field.NewTime(tableName, "createdAt")
	_group.UpdatedAt = field.NewTime(tableName, "updatedAt")
	_group.Name = field.NewString(tableName, "name")
	_group.Description = field.NewString(tableName, "description")
	_group.RoleId = field.NewString(tableName, "role_id")
	_group.OrganizationId = field.NewString(tableName, "organizationId")
	_group.User = groupHasManyUser{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("User", "models.User"),
		Group: struct {
			field.RelationField
			Role struct {
				field.RelationField
				Organization struct {
					field.RelationField
				}
				Menus struct {
					field.RelationField
					Parent struct {
						field.RelationField
					}
					Permission struct {
						field.RelationField
						Menu struct {
							field.RelationField
						}
					}
				}
			}
			Organization struct {
				field.RelationField
			}
			User struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("User.Group", "models.Group"),
			Role: struct {
				field.RelationField
				Organization struct {
					field.RelationField
				}
				Menus struct {
					field.RelationField
					Parent struct {
						field.RelationField
					}
					Permission struct {
						field.RelationField
						Menu struct {
							field.RelationField
						}
					}
				}
			}{
				RelationField: field.NewRelation("User.Group.Role", "models.Role"),
				Organization: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("User.Group.Role.Organization", "models.Organization"),
				},
				Menus: struct {
					field.RelationField
					Parent struct {
						field.RelationField
					}
					Permission struct {
						field.RelationField
						Menu struct {
							field.RelationField
						}
					}
				}{
					RelationField: field.NewRelation("User.Group.Role.Menus", "models.Menu"),
					Parent: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("User.Group.Role.Menus.Parent", "models.Menu"),
					},
					Permission: struct {
						field.RelationField
						Menu struct {
							field.RelationField
						}
					}{
						RelationField: field.NewRelation("User.Group.Role.Menus.Permission", "models.Permission"),
						Menu: struct {
							field.RelationField
						}{
							RelationField: field.NewRelation("User.Group.Role.Menus.Permission.Menu", "models.Menu"),
						},
					},
				},
			},
			Organization: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("User.Group.Organization", "models.Organization"),
			},
			User: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("User.Group.User", "models.User"),
			},
		},
		Role: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("User.Role", "models.Role"),
		},
		Organization: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("User.Organization", "models.Organization"),
		},
	}

	_group.Role = groupBelongsToRole{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Role", "models.Role"),
	}

	_group.Organization = groupBelongsToOrganization{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Organization", "models.Organization"),
	}

	_group.fillFieldMap()

	return _group
}

type group struct {
	groupDo

	ALL            field.Asterisk
	Id             field.String
	CreatedAt      field.Time
	UpdatedAt      field.Time
	Name           field.String
	Description    field.String
	RoleId         field.String
	OrganizationId field.String
	User           groupHasManyUser

	Role groupBelongsToRole

	Organization groupBelongsToOrganization

	fieldMap map[string]field.Expr
}

func (g group) Table(newTableName string) *group {
	g.groupDo.UseTable(newTableName)
	return g.updateTableName(newTableName)
}

func (g group) As(alias string) *group {
	g.groupDo.DO = *(g.groupDo.As(alias).(*gen.DO))
	return g.updateTableName(alias)
}

func (g *group) updateTableName(table string) *group {
	g.ALL = field.NewAsterisk(table)
	g.Id = field.NewString(table, "id")
	g.CreatedAt = field.NewTime(table, "createdAt")
	g.UpdatedAt = field.NewTime(table, "updatedAt")
	g.Name = field.NewString(table, "name")
	g.Description = field.NewString(table, "description")
	g.RoleId = field.NewString(table, "role_id")
	g.OrganizationId = field.NewString(table, "organizationId")

	g.fillFieldMap()

	return g
}

func (g *group) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := g.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (g *group) fillFieldMap() {
	g.fieldMap = make(map[string]field.Expr, 10)
	g.fieldMap["id"] = g.Id
	g.fieldMap["createdAt"] = g.CreatedAt
	g.fieldMap["updatedAt"] = g.UpdatedAt
	g.fieldMap["name"] = g.Name
	g.fieldMap["description"] = g.Description
	g.fieldMap["role_id"] = g.RoleId
	g.fieldMap["organizationId"] = g.OrganizationId

}

func (g group) clone(db *gorm.DB) group {
	g.groupDo.ReplaceConnPool(db.Statement.ConnPool)
	return g
}

func (g group) replaceDB(db *gorm.DB) group {
	g.groupDo.ReplaceDB(db)
	return g
}

type groupHasManyUser struct {
	db *gorm.DB

	field.RelationField

	Group struct {
		field.RelationField
		Role struct {
			field.RelationField
			Organization struct {
				field.RelationField
			}
			Menus struct {
				field.RelationField
				Parent struct {
					field.RelationField
				}
				Permission struct {
					field.RelationField
					Menu struct {
						field.RelationField
					}
				}
			}
		}
		Organization struct {
			field.RelationField
		}
		User struct {
			field.RelationField
		}
	}
	Role struct {
		field.RelationField
	}
	Organization struct {
		field.RelationField
	}
}

func (a groupHasManyUser) Where(conds ...field.Expr) *groupHasManyUser {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a groupHasManyUser) WithContext(ctx context.Context) *groupHasManyUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a groupHasManyUser) Session(session *gorm.Session) *groupHasManyUser {
	a.db = a.db.Session(session)
	return &a
}

func (a groupHasManyUser) Model(m *models.Group) *groupHasManyUserTx {
	return &groupHasManyUserTx{a.db.Model(m).Association(a.Name())}
}

type groupHasManyUserTx struct{ tx *gorm.Association }

func (a groupHasManyUserTx) Find() (result []*models.User, err error) {
	return result, a.tx.Find(&result)
}

func (a groupHasManyUserTx) Append(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a groupHasManyUserTx) Replace(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a groupHasManyUserTx) Delete(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a groupHasManyUserTx) Clear() error {
	return a.tx.Clear()
}

func (a groupHasManyUserTx) Count() int64 {
	return a.tx.Count()
}

type groupBelongsToRole struct {
	db *gorm.DB

	field.RelationField
}

func (a groupBelongsToRole) Where(conds ...field.Expr) *groupBelongsToRole {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a groupBelongsToRole) WithContext(ctx context.Context) *groupBelongsToRole {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a groupBelongsToRole) Session(session *gorm.Session) *groupBelongsToRole {
	a.db = a.db.Session(session)
	return &a
}

func (a groupBelongsToRole) Model(m *models.Group) *groupBelongsToRoleTx {
	return &groupBelongsToRoleTx{a.db.Model(m).Association(a.Name())}
}

type groupBelongsToRoleTx struct{ tx *gorm.Association }

func (a groupBelongsToRoleTx) Find() (result *models.Role, err error) {
	return result, a.tx.Find(&result)
}

func (a groupBelongsToRoleTx) Append(values ...*models.Role) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a groupBelongsToRoleTx) Replace(values ...*models.Role) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a groupBelongsToRoleTx) Delete(values ...*models.Role) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a groupBelongsToRoleTx) Clear() error {
	return a.tx.Clear()
}

func (a groupBelongsToRoleTx) Count() int64 {
	return a.tx.Count()
}

type groupBelongsToOrganization struct {
	db *gorm.DB

	field.RelationField
}

func (a groupBelongsToOrganization) Where(conds ...field.Expr) *groupBelongsToOrganization {
	if len(conds) == 0 {
		return &a
	}

	exprs := make([]clause.Expression, 0, len(conds))
	for _, cond := range conds {
		exprs = append(exprs, cond.BeCond().(clause.Expression))
	}
	a.db = a.db.Clauses(clause.Where{Exprs: exprs})
	return &a
}

func (a groupBelongsToOrganization) WithContext(ctx context.Context) *groupBelongsToOrganization {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a groupBelongsToOrganization) Session(session *gorm.Session) *groupBelongsToOrganization {
	a.db = a.db.Session(session)
	return &a
}

func (a groupBelongsToOrganization) Model(m *models.Group) *groupBelongsToOrganizationTx {
	return &groupBelongsToOrganizationTx{a.db.Model(m).Association(a.Name())}
}

type groupBelongsToOrganizationTx struct{ tx *gorm.Association }

func (a groupBelongsToOrganizationTx) Find() (result *models.Organization, err error) {
	return result, a.tx.Find(&result)
}

func (a groupBelongsToOrganizationTx) Append(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a groupBelongsToOrganizationTx) Replace(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a groupBelongsToOrganizationTx) Delete(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a groupBelongsToOrganizationTx) Clear() error {
	return a.tx.Clear()
}

func (a groupBelongsToOrganizationTx) Count() int64 {
	return a.tx.Count()
}

type groupDo struct{ gen.DO }

type IGroupDo interface {
	gen.SubQuery
	Debug() IGroupDo
	WithContext(ctx context.Context) IGroupDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IGroupDo
	WriteDB() IGroupDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IGroupDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IGroupDo
	Not(conds ...gen.Condition) IGroupDo
	Or(conds ...gen.Condition) IGroupDo
	Select(conds ...field.Expr) IGroupDo
	Where(conds ...gen.Condition) IGroupDo
	Order(conds ...field.Expr) IGroupDo
	Distinct(cols ...field.Expr) IGroupDo
	Omit(cols ...field.Expr) IGroupDo
	Join(table schema.Tabler, on ...field.Expr) IGroupDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IGroupDo
	RightJoin(table schema.Tabler, on ...field.Expr) IGroupDo
	Group(cols ...field.Expr) IGroupDo
	Having(conds ...gen.Condition) IGroupDo
	Limit(limit int) IGroupDo
	Offset(offset int) IGroupDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IGroupDo
	Unscoped() IGroupDo
	Create(values ...*models.Group) error
	CreateInBatches(values []*models.Group, batchSize int) error
	Save(values ...*models.Group) error
	First() (*models.Group, error)
	Take() (*models.Group, error)
	Last() (*models.Group, error)
	Find() ([]*models.Group, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Group, err error)
	FindInBatches(result *[]*models.Group, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Group) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IGroupDo
	Assign(attrs ...field.AssignExpr) IGroupDo
	Joins(fields ...field.RelationField) IGroupDo
	Preload(fields ...field.RelationField) IGroupDo
	FirstOrInit() (*models.Group, error)
	FirstOrCreate() (*models.Group, error)
	FindByPage(offset int, limit int) (result []*models.Group, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IGroupDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (g groupDo) Debug() IGroupDo {
	return g.withDO(g.DO.Debug())
}

func (g groupDo) WithContext(ctx context.Context) IGroupDo {
	return g.withDO(g.DO.WithContext(ctx))
}

func (g groupDo) ReadDB() IGroupDo {
	return g.Clauses(dbresolver.Read)
}

func (g groupDo) WriteDB() IGroupDo {
	return g.Clauses(dbresolver.Write)
}

func (g groupDo) Session(config *gorm.Session) IGroupDo {
	return g.withDO(g.DO.Session(config))
}

func (g groupDo) Clauses(conds ...clause.Expression) IGroupDo {
	return g.withDO(g.DO.Clauses(conds...))
}

func (g groupDo) Returning(value interface{}, columns ...string) IGroupDo {
	return g.withDO(g.DO.Returning(value, columns...))
}

func (g groupDo) Not(conds ...gen.Condition) IGroupDo {
	return g.withDO(g.DO.Not(conds...))
}

func (g groupDo) Or(conds ...gen.Condition) IGroupDo {
	return g.withDO(g.DO.Or(conds...))
}

func (g groupDo) Select(conds ...field.Expr) IGroupDo {
	return g.withDO(g.DO.Select(conds...))
}

func (g groupDo) Where(conds ...gen.Condition) IGroupDo {
	return g.withDO(g.DO.Where(conds...))
}

func (g groupDo) Order(conds ...field.Expr) IGroupDo {
	return g.withDO(g.DO.Order(conds...))
}

func (g groupDo) Distinct(cols ...field.Expr) IGroupDo {
	return g.withDO(g.DO.Distinct(cols...))
}

func (g groupDo) Omit(cols ...field.Expr) IGroupDo {
	return g.withDO(g.DO.Omit(cols...))
}

func (g groupDo) Join(table schema.Tabler, on ...field.Expr) IGroupDo {
	return g.withDO(g.DO.Join(table, on...))
}

func (g groupDo) LeftJoin(table schema.Tabler, on ...field.Expr) IGroupDo {
	return g.withDO(g.DO.LeftJoin(table, on...))
}

func (g groupDo) RightJoin(table schema.Tabler, on ...field.Expr) IGroupDo {
	return g.withDO(g.DO.RightJoin(table, on...))
}

func (g groupDo) Group(cols ...field.Expr) IGroupDo {
	return g.withDO(g.DO.Group(cols...))
}

func (g groupDo) Having(conds ...gen.Condition) IGroupDo {
	return g.withDO(g.DO.Having(conds...))
}

func (g groupDo) Limit(limit int) IGroupDo {
	return g.withDO(g.DO.Limit(limit))
}

func (g groupDo) Offset(offset int) IGroupDo {
	return g.withDO(g.DO.Offset(offset))
}

func (g groupDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IGroupDo {
	return g.withDO(g.DO.Scopes(funcs...))
}

func (g groupDo) Unscoped() IGroupDo {
	return g.withDO(g.DO.Unscoped())
}

func (g groupDo) Create(values ...*models.Group) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Create(values)
}

func (g groupDo) CreateInBatches(values []*models.Group, batchSize int) error {
	return g.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (g groupDo) Save(values ...*models.Group) error {
	if len(values) == 0 {
		return nil
	}
	return g.DO.Save(values)
}

func (g groupDo) First() (*models.Group, error) {
	if result, err := g.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Group), nil
	}
}

func (g groupDo) Take() (*models.Group, error) {
	if result, err := g.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Group), nil
	}
}

func (g groupDo) Last() (*models.Group, error) {
	if result, err := g.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Group), nil
	}
}

func (g groupDo) Find() ([]*models.Group, error) {
	result, err := g.DO.Find()
	return result.([]*models.Group), err
}

func (g groupDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Group, err error) {
	buf := make([]*models.Group, 0, batchSize)
	err = g.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (g groupDo) FindInBatches(result *[]*models.Group, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return g.DO.FindInBatches(result, batchSize, fc)
}

func (g groupDo) Attrs(attrs ...field.AssignExpr) IGroupDo {
	return g.withDO(g.DO.Attrs(attrs...))
}

func (g groupDo) Assign(attrs ...field.AssignExpr) IGroupDo {
	return g.withDO(g.DO.Assign(attrs...))
}

func (g groupDo) Joins(fields ...field.RelationField) IGroupDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Joins(_f))
	}
	return &g
}

func (g groupDo) Preload(fields ...field.RelationField) IGroupDo {
	for _, _f := range fields {
		g = *g.withDO(g.DO.Preload(_f))
	}
	return &g
}

func (g groupDo) FirstOrInit() (*models.Group, error) {
	if result, err := g.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Group), nil
	}
}

func (g groupDo) FirstOrCreate() (*models.Group, error) {
	if result, err := g.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Group), nil
	}
}

func (g groupDo) FindByPage(offset int, limit int) (result []*models.Group, count int64, err error) {
	result, err = g.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = g.Offset(-1).Limit(-1).Count()
	return
}

func (g groupDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = g.Count()
	if err != nil {
		return
	}

	err = g.Offset(offset).Limit(limit).Scan(result)
	return
}

func (g groupDo) Scan(result interface{}) (err error) {
	return g.DO.Scan(result)
}

func (g groupDo) Delete(models ...*models.Group) (result gen.ResultInfo, err error) {
	return g.DO.Delete(models)
}

func (g *groupDo) withDO(do gen.Dao) *groupDo {
	g.DO = *do.(*gen.DO)
	return g
}
