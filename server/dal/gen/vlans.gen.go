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

func newVlan(db *gorm.DB, opts ...gen.DOOption) vlan {
	_vlan := vlan{}

	_vlan.vlanDo.UseDB(db, opts...)
	_vlan.vlanDo.UseModel(&models.Vlan{})

	tableName := _vlan.vlanDo.TableName()
	_vlan.ALL = field.NewAsterisk(tableName)
	_vlan.Id = field.NewString(tableName, "id")
	_vlan.CreatedAt = field.NewTime(tableName, "created_at")
	_vlan.UpdatedAt = field.NewTime(tableName, "updated_at")
	_vlan.Name = field.NewString(tableName, "name")
	_vlan.Vid = field.NewUint32(tableName, "vid")
	_vlan.Description = field.NewString(tableName, "description")
	_vlan.Status = field.NewString(tableName, "status")
	_vlan.SiteId = field.NewString(tableName, "site_id")
	_vlan.OrganizationId = field.NewString(tableName, "organization_id")
	_vlan.Site = vlanBelongsToSite{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Site", "models.Site"),
		Organization: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Site.Organization", "models.Organization"),
		},
	}

	_vlan.Organization = vlanBelongsToOrganization{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Organization", "models.Organization"),
	}

	_vlan.fillFieldMap()

	return _vlan
}

type vlan struct {
	vlanDo

	ALL            field.Asterisk
	Id             field.String
	CreatedAt      field.Time
	UpdatedAt      field.Time
	Name           field.String
	Vid            field.Uint32
	Description    field.String
	Status         field.String
	SiteId         field.String
	OrganizationId field.String
	Site           vlanBelongsToSite

	Organization vlanBelongsToOrganization

	fieldMap map[string]field.Expr
}

func (v vlan) Table(newTableName string) *vlan {
	v.vlanDo.UseTable(newTableName)
	return v.updateTableName(newTableName)
}

func (v vlan) As(alias string) *vlan {
	v.vlanDo.DO = *(v.vlanDo.As(alias).(*gen.DO))
	return v.updateTableName(alias)
}

func (v *vlan) updateTableName(table string) *vlan {
	v.ALL = field.NewAsterisk(table)
	v.Id = field.NewString(table, "id")
	v.CreatedAt = field.NewTime(table, "created_at")
	v.UpdatedAt = field.NewTime(table, "updated_at")
	v.Name = field.NewString(table, "name")
	v.Vid = field.NewUint32(table, "vid")
	v.Description = field.NewString(table, "description")
	v.Status = field.NewString(table, "status")
	v.SiteId = field.NewString(table, "site_id")
	v.OrganizationId = field.NewString(table, "organization_id")

	v.fillFieldMap()

	return v
}

func (v *vlan) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := v.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (v *vlan) fillFieldMap() {
	v.fieldMap = make(map[string]field.Expr, 11)
	v.fieldMap["id"] = v.Id
	v.fieldMap["created_at"] = v.CreatedAt
	v.fieldMap["updated_at"] = v.UpdatedAt
	v.fieldMap["name"] = v.Name
	v.fieldMap["vid"] = v.Vid
	v.fieldMap["description"] = v.Description
	v.fieldMap["status"] = v.Status
	v.fieldMap["site_id"] = v.SiteId
	v.fieldMap["organization_id"] = v.OrganizationId

}

func (v vlan) clone(db *gorm.DB) vlan {
	v.vlanDo.ReplaceConnPool(db.Statement.ConnPool)
	return v
}

func (v vlan) replaceDB(db *gorm.DB) vlan {
	v.vlanDo.ReplaceDB(db)
	return v
}

type vlanBelongsToSite struct {
	db *gorm.DB

	field.RelationField

	Organization struct {
		field.RelationField
	}
}

func (a vlanBelongsToSite) Where(conds ...field.Expr) *vlanBelongsToSite {
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

func (a vlanBelongsToSite) WithContext(ctx context.Context) *vlanBelongsToSite {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a vlanBelongsToSite) Session(session *gorm.Session) *vlanBelongsToSite {
	a.db = a.db.Session(session)
	return &a
}

func (a vlanBelongsToSite) Model(m *models.Vlan) *vlanBelongsToSiteTx {
	return &vlanBelongsToSiteTx{a.db.Model(m).Association(a.Name())}
}

type vlanBelongsToSiteTx struct{ tx *gorm.Association }

func (a vlanBelongsToSiteTx) Find() (result *models.Site, err error) {
	return result, a.tx.Find(&result)
}

func (a vlanBelongsToSiteTx) Append(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a vlanBelongsToSiteTx) Replace(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a vlanBelongsToSiteTx) Delete(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a vlanBelongsToSiteTx) Clear() error {
	return a.tx.Clear()
}

func (a vlanBelongsToSiteTx) Count() int64 {
	return a.tx.Count()
}

type vlanBelongsToOrganization struct {
	db *gorm.DB

	field.RelationField
}

func (a vlanBelongsToOrganization) Where(conds ...field.Expr) *vlanBelongsToOrganization {
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

func (a vlanBelongsToOrganization) WithContext(ctx context.Context) *vlanBelongsToOrganization {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a vlanBelongsToOrganization) Session(session *gorm.Session) *vlanBelongsToOrganization {
	a.db = a.db.Session(session)
	return &a
}

func (a vlanBelongsToOrganization) Model(m *models.Vlan) *vlanBelongsToOrganizationTx {
	return &vlanBelongsToOrganizationTx{a.db.Model(m).Association(a.Name())}
}

type vlanBelongsToOrganizationTx struct{ tx *gorm.Association }

func (a vlanBelongsToOrganizationTx) Find() (result *models.Organization, err error) {
	return result, a.tx.Find(&result)
}

func (a vlanBelongsToOrganizationTx) Append(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a vlanBelongsToOrganizationTx) Replace(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a vlanBelongsToOrganizationTx) Delete(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a vlanBelongsToOrganizationTx) Clear() error {
	return a.tx.Clear()
}

func (a vlanBelongsToOrganizationTx) Count() int64 {
	return a.tx.Count()
}

type vlanDo struct{ gen.DO }

type IVlanDo interface {
	gen.SubQuery
	Debug() IVlanDo
	WithContext(ctx context.Context) IVlanDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IVlanDo
	WriteDB() IVlanDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IVlanDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IVlanDo
	Not(conds ...gen.Condition) IVlanDo
	Or(conds ...gen.Condition) IVlanDo
	Select(conds ...field.Expr) IVlanDo
	Where(conds ...gen.Condition) IVlanDo
	Order(conds ...field.Expr) IVlanDo
	Distinct(cols ...field.Expr) IVlanDo
	Omit(cols ...field.Expr) IVlanDo
	Join(table schema.Tabler, on ...field.Expr) IVlanDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IVlanDo
	RightJoin(table schema.Tabler, on ...field.Expr) IVlanDo
	Group(cols ...field.Expr) IVlanDo
	Having(conds ...gen.Condition) IVlanDo
	Limit(limit int) IVlanDo
	Offset(offset int) IVlanDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IVlanDo
	Unscoped() IVlanDo
	Create(values ...*models.Vlan) error
	CreateInBatches(values []*models.Vlan, batchSize int) error
	Save(values ...*models.Vlan) error
	First() (*models.Vlan, error)
	Take() (*models.Vlan, error)
	Last() (*models.Vlan, error)
	Find() ([]*models.Vlan, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Vlan, err error)
	FindInBatches(result *[]*models.Vlan, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Vlan) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IVlanDo
	Assign(attrs ...field.AssignExpr) IVlanDo
	Joins(fields ...field.RelationField) IVlanDo
	Preload(fields ...field.RelationField) IVlanDo
	FirstOrInit() (*models.Vlan, error)
	FirstOrCreate() (*models.Vlan, error)
	FindByPage(offset int, limit int) (result []*models.Vlan, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IVlanDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (v vlanDo) Debug() IVlanDo {
	return v.withDO(v.DO.Debug())
}

func (v vlanDo) WithContext(ctx context.Context) IVlanDo {
	return v.withDO(v.DO.WithContext(ctx))
}

func (v vlanDo) ReadDB() IVlanDo {
	return v.Clauses(dbresolver.Read)
}

func (v vlanDo) WriteDB() IVlanDo {
	return v.Clauses(dbresolver.Write)
}

func (v vlanDo) Session(config *gorm.Session) IVlanDo {
	return v.withDO(v.DO.Session(config))
}

func (v vlanDo) Clauses(conds ...clause.Expression) IVlanDo {
	return v.withDO(v.DO.Clauses(conds...))
}

func (v vlanDo) Returning(value interface{}, columns ...string) IVlanDo {
	return v.withDO(v.DO.Returning(value, columns...))
}

func (v vlanDo) Not(conds ...gen.Condition) IVlanDo {
	return v.withDO(v.DO.Not(conds...))
}

func (v vlanDo) Or(conds ...gen.Condition) IVlanDo {
	return v.withDO(v.DO.Or(conds...))
}

func (v vlanDo) Select(conds ...field.Expr) IVlanDo {
	return v.withDO(v.DO.Select(conds...))
}

func (v vlanDo) Where(conds ...gen.Condition) IVlanDo {
	return v.withDO(v.DO.Where(conds...))
}

func (v vlanDo) Order(conds ...field.Expr) IVlanDo {
	return v.withDO(v.DO.Order(conds...))
}

func (v vlanDo) Distinct(cols ...field.Expr) IVlanDo {
	return v.withDO(v.DO.Distinct(cols...))
}

func (v vlanDo) Omit(cols ...field.Expr) IVlanDo {
	return v.withDO(v.DO.Omit(cols...))
}

func (v vlanDo) Join(table schema.Tabler, on ...field.Expr) IVlanDo {
	return v.withDO(v.DO.Join(table, on...))
}

func (v vlanDo) LeftJoin(table schema.Tabler, on ...field.Expr) IVlanDo {
	return v.withDO(v.DO.LeftJoin(table, on...))
}

func (v vlanDo) RightJoin(table schema.Tabler, on ...field.Expr) IVlanDo {
	return v.withDO(v.DO.RightJoin(table, on...))
}

func (v vlanDo) Group(cols ...field.Expr) IVlanDo {
	return v.withDO(v.DO.Group(cols...))
}

func (v vlanDo) Having(conds ...gen.Condition) IVlanDo {
	return v.withDO(v.DO.Having(conds...))
}

func (v vlanDo) Limit(limit int) IVlanDo {
	return v.withDO(v.DO.Limit(limit))
}

func (v vlanDo) Offset(offset int) IVlanDo {
	return v.withDO(v.DO.Offset(offset))
}

func (v vlanDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IVlanDo {
	return v.withDO(v.DO.Scopes(funcs...))
}

func (v vlanDo) Unscoped() IVlanDo {
	return v.withDO(v.DO.Unscoped())
}

func (v vlanDo) Create(values ...*models.Vlan) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Create(values)
}

func (v vlanDo) CreateInBatches(values []*models.Vlan, batchSize int) error {
	return v.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (v vlanDo) Save(values ...*models.Vlan) error {
	if len(values) == 0 {
		return nil
	}
	return v.DO.Save(values)
}

func (v vlanDo) First() (*models.Vlan, error) {
	if result, err := v.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Vlan), nil
	}
}

func (v vlanDo) Take() (*models.Vlan, error) {
	if result, err := v.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Vlan), nil
	}
}

func (v vlanDo) Last() (*models.Vlan, error) {
	if result, err := v.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Vlan), nil
	}
}

func (v vlanDo) Find() ([]*models.Vlan, error) {
	result, err := v.DO.Find()
	return result.([]*models.Vlan), err
}

func (v vlanDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Vlan, err error) {
	buf := make([]*models.Vlan, 0, batchSize)
	err = v.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (v vlanDo) FindInBatches(result *[]*models.Vlan, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return v.DO.FindInBatches(result, batchSize, fc)
}

func (v vlanDo) Attrs(attrs ...field.AssignExpr) IVlanDo {
	return v.withDO(v.DO.Attrs(attrs...))
}

func (v vlanDo) Assign(attrs ...field.AssignExpr) IVlanDo {
	return v.withDO(v.DO.Assign(attrs...))
}

func (v vlanDo) Joins(fields ...field.RelationField) IVlanDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Joins(_f))
	}
	return &v
}

func (v vlanDo) Preload(fields ...field.RelationField) IVlanDo {
	for _, _f := range fields {
		v = *v.withDO(v.DO.Preload(_f))
	}
	return &v
}

func (v vlanDo) FirstOrInit() (*models.Vlan, error) {
	if result, err := v.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Vlan), nil
	}
}

func (v vlanDo) FirstOrCreate() (*models.Vlan, error) {
	if result, err := v.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Vlan), nil
	}
}

func (v vlanDo) FindByPage(offset int, limit int) (result []*models.Vlan, count int64, err error) {
	result, err = v.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = v.Offset(-1).Limit(-1).Count()
	return
}

func (v vlanDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = v.Count()
	if err != nil {
		return
	}

	err = v.Offset(offset).Limit(limit).Scan(result)
	return
}

func (v vlanDo) Scan(result interface{}) (err error) {
	return v.DO.Scan(result)
}

func (v vlanDo) Delete(models ...*models.Vlan) (result gen.ResultInfo, err error) {
	return v.DO.Delete(models)
}

func (v *vlanDo) withDO(do gen.Dao) *vlanDo {
	v.DO = *do.(*gen.DO)
	return v
}
