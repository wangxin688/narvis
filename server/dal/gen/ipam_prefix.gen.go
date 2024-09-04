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

func newPrefix(db *gorm.DB, opts ...gen.DOOption) prefix {
	_prefix := prefix{}

	_prefix.prefixDo.UseDB(db, opts...)
	_prefix.prefixDo.UseModel(&models.Prefix{})

	tableName := _prefix.prefixDo.TableName()
	_prefix.ALL = field.NewAsterisk(tableName)
	_prefix.Id = field.NewString(tableName, "id")
	_prefix.CreatedAt = field.NewTime(tableName, "createdAt")
	_prefix.UpdatedAt = field.NewTime(tableName, "updatedAt")
	_prefix.Range = field.NewString(tableName, "range")
	_prefix.Version = field.NewString(tableName, "version")
	_prefix.Type = field.NewString(tableName, "type")
	_prefix.VlanId = field.NewUint32(tableName, "vlanId")
	_prefix.VlanName = field.NewString(tableName, "vlanName")
	_prefix.SiteId = field.NewString(tableName, "siteId")
	_prefix.OrganizationId = field.NewString(tableName, "organizationId")
	_prefix.Site = prefixBelongsToSite{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Site", "models.Site"),
		Organization: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Site.Organization", "models.Organization"),
		},
	}

	_prefix.Organization = prefixBelongsToOrganization{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Organization", "models.Organization"),
	}

	_prefix.fillFieldMap()

	return _prefix
}

type prefix struct {
	prefixDo

	ALL            field.Asterisk
	Id             field.String
	CreatedAt      field.Time
	UpdatedAt      field.Time
	Range          field.String
	Version        field.String
	Type           field.String
	VlanId         field.Uint32
	VlanName       field.String
	SiteId         field.String
	OrganizationId field.String
	Site           prefixBelongsToSite

	Organization prefixBelongsToOrganization

	fieldMap map[string]field.Expr
}

func (p prefix) Table(newTableName string) *prefix {
	p.prefixDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p prefix) As(alias string) *prefix {
	p.prefixDo.DO = *(p.prefixDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *prefix) updateTableName(table string) *prefix {
	p.ALL = field.NewAsterisk(table)
	p.Id = field.NewString(table, "id")
	p.CreatedAt = field.NewTime(table, "createdAt")
	p.UpdatedAt = field.NewTime(table, "updatedAt")
	p.Range = field.NewString(table, "range")
	p.Version = field.NewString(table, "version")
	p.Type = field.NewString(table, "type")
	p.VlanId = field.NewUint32(table, "vlanId")
	p.VlanName = field.NewString(table, "vlanName")
	p.SiteId = field.NewString(table, "siteId")
	p.OrganizationId = field.NewString(table, "organizationId")

	p.fillFieldMap()

	return p
}

func (p *prefix) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *prefix) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 12)
	p.fieldMap["id"] = p.Id
	p.fieldMap["createdAt"] = p.CreatedAt
	p.fieldMap["updatedAt"] = p.UpdatedAt
	p.fieldMap["range"] = p.Range
	p.fieldMap["version"] = p.Version
	p.fieldMap["type"] = p.Type
	p.fieldMap["vlanId"] = p.VlanId
	p.fieldMap["vlanName"] = p.VlanName
	p.fieldMap["siteId"] = p.SiteId
	p.fieldMap["organizationId"] = p.OrganizationId

}

func (p prefix) clone(db *gorm.DB) prefix {
	p.prefixDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p prefix) replaceDB(db *gorm.DB) prefix {
	p.prefixDo.ReplaceDB(db)
	return p
}

type prefixBelongsToSite struct {
	db *gorm.DB

	field.RelationField

	Organization struct {
		field.RelationField
	}
}

func (a prefixBelongsToSite) Where(conds ...field.Expr) *prefixBelongsToSite {
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

func (a prefixBelongsToSite) WithContext(ctx context.Context) *prefixBelongsToSite {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a prefixBelongsToSite) Session(session *gorm.Session) *prefixBelongsToSite {
	a.db = a.db.Session(session)
	return &a
}

func (a prefixBelongsToSite) Model(m *models.Prefix) *prefixBelongsToSiteTx {
	return &prefixBelongsToSiteTx{a.db.Model(m).Association(a.Name())}
}

type prefixBelongsToSiteTx struct{ tx *gorm.Association }

func (a prefixBelongsToSiteTx) Find() (result *models.Site, err error) {
	return result, a.tx.Find(&result)
}

func (a prefixBelongsToSiteTx) Append(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a prefixBelongsToSiteTx) Replace(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a prefixBelongsToSiteTx) Delete(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a prefixBelongsToSiteTx) Clear() error {
	return a.tx.Clear()
}

func (a prefixBelongsToSiteTx) Count() int64 {
	return a.tx.Count()
}

type prefixBelongsToOrganization struct {
	db *gorm.DB

	field.RelationField
}

func (a prefixBelongsToOrganization) Where(conds ...field.Expr) *prefixBelongsToOrganization {
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

func (a prefixBelongsToOrganization) WithContext(ctx context.Context) *prefixBelongsToOrganization {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a prefixBelongsToOrganization) Session(session *gorm.Session) *prefixBelongsToOrganization {
	a.db = a.db.Session(session)
	return &a
}

func (a prefixBelongsToOrganization) Model(m *models.Prefix) *prefixBelongsToOrganizationTx {
	return &prefixBelongsToOrganizationTx{a.db.Model(m).Association(a.Name())}
}

type prefixBelongsToOrganizationTx struct{ tx *gorm.Association }

func (a prefixBelongsToOrganizationTx) Find() (result *models.Organization, err error) {
	return result, a.tx.Find(&result)
}

func (a prefixBelongsToOrganizationTx) Append(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a prefixBelongsToOrganizationTx) Replace(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a prefixBelongsToOrganizationTx) Delete(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a prefixBelongsToOrganizationTx) Clear() error {
	return a.tx.Clear()
}

func (a prefixBelongsToOrganizationTx) Count() int64 {
	return a.tx.Count()
}

type prefixDo struct{ gen.DO }

type IPrefixDo interface {
	gen.SubQuery
	Debug() IPrefixDo
	WithContext(ctx context.Context) IPrefixDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IPrefixDo
	WriteDB() IPrefixDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IPrefixDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IPrefixDo
	Not(conds ...gen.Condition) IPrefixDo
	Or(conds ...gen.Condition) IPrefixDo
	Select(conds ...field.Expr) IPrefixDo
	Where(conds ...gen.Condition) IPrefixDo
	Order(conds ...field.Expr) IPrefixDo
	Distinct(cols ...field.Expr) IPrefixDo
	Omit(cols ...field.Expr) IPrefixDo
	Join(table schema.Tabler, on ...field.Expr) IPrefixDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IPrefixDo
	RightJoin(table schema.Tabler, on ...field.Expr) IPrefixDo
	Group(cols ...field.Expr) IPrefixDo
	Having(conds ...gen.Condition) IPrefixDo
	Limit(limit int) IPrefixDo
	Offset(offset int) IPrefixDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IPrefixDo
	Unscoped() IPrefixDo
	Create(values ...*models.Prefix) error
	CreateInBatches(values []*models.Prefix, batchSize int) error
	Save(values ...*models.Prefix) error
	First() (*models.Prefix, error)
	Take() (*models.Prefix, error)
	Last() (*models.Prefix, error)
	Find() ([]*models.Prefix, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Prefix, err error)
	FindInBatches(result *[]*models.Prefix, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Prefix) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IPrefixDo
	Assign(attrs ...field.AssignExpr) IPrefixDo
	Joins(fields ...field.RelationField) IPrefixDo
	Preload(fields ...field.RelationField) IPrefixDo
	FirstOrInit() (*models.Prefix, error)
	FirstOrCreate() (*models.Prefix, error)
	FindByPage(offset int, limit int) (result []*models.Prefix, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IPrefixDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p prefixDo) Debug() IPrefixDo {
	return p.withDO(p.DO.Debug())
}

func (p prefixDo) WithContext(ctx context.Context) IPrefixDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p prefixDo) ReadDB() IPrefixDo {
	return p.Clauses(dbresolver.Read)
}

func (p prefixDo) WriteDB() IPrefixDo {
	return p.Clauses(dbresolver.Write)
}

func (p prefixDo) Session(config *gorm.Session) IPrefixDo {
	return p.withDO(p.DO.Session(config))
}

func (p prefixDo) Clauses(conds ...clause.Expression) IPrefixDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p prefixDo) Returning(value interface{}, columns ...string) IPrefixDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p prefixDo) Not(conds ...gen.Condition) IPrefixDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p prefixDo) Or(conds ...gen.Condition) IPrefixDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p prefixDo) Select(conds ...field.Expr) IPrefixDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p prefixDo) Where(conds ...gen.Condition) IPrefixDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p prefixDo) Order(conds ...field.Expr) IPrefixDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p prefixDo) Distinct(cols ...field.Expr) IPrefixDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p prefixDo) Omit(cols ...field.Expr) IPrefixDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p prefixDo) Join(table schema.Tabler, on ...field.Expr) IPrefixDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p prefixDo) LeftJoin(table schema.Tabler, on ...field.Expr) IPrefixDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p prefixDo) RightJoin(table schema.Tabler, on ...field.Expr) IPrefixDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p prefixDo) Group(cols ...field.Expr) IPrefixDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p prefixDo) Having(conds ...gen.Condition) IPrefixDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p prefixDo) Limit(limit int) IPrefixDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p prefixDo) Offset(offset int) IPrefixDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p prefixDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IPrefixDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p prefixDo) Unscoped() IPrefixDo {
	return p.withDO(p.DO.Unscoped())
}

func (p prefixDo) Create(values ...*models.Prefix) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p prefixDo) CreateInBatches(values []*models.Prefix, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p prefixDo) Save(values ...*models.Prefix) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p prefixDo) First() (*models.Prefix, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Prefix), nil
	}
}

func (p prefixDo) Take() (*models.Prefix, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Prefix), nil
	}
}

func (p prefixDo) Last() (*models.Prefix, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Prefix), nil
	}
}

func (p prefixDo) Find() ([]*models.Prefix, error) {
	result, err := p.DO.Find()
	return result.([]*models.Prefix), err
}

func (p prefixDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Prefix, err error) {
	buf := make([]*models.Prefix, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p prefixDo) FindInBatches(result *[]*models.Prefix, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p prefixDo) Attrs(attrs ...field.AssignExpr) IPrefixDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p prefixDo) Assign(attrs ...field.AssignExpr) IPrefixDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p prefixDo) Joins(fields ...field.RelationField) IPrefixDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p prefixDo) Preload(fields ...field.RelationField) IPrefixDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p prefixDo) FirstOrInit() (*models.Prefix, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Prefix), nil
	}
}

func (p prefixDo) FirstOrCreate() (*models.Prefix, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Prefix), nil
	}
}

func (p prefixDo) FindByPage(offset int, limit int) (result []*models.Prefix, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p prefixDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p prefixDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p prefixDo) Delete(models ...*models.Prefix) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *prefixDo) withDO(do gen.Dao) *prefixDo {
	p.DO = *do.(*gen.DO)
	return p
}
