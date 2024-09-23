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

func newProxy(db *gorm.DB, opts ...gen.DOOption) proxy {
	_proxy := proxy{}

	_proxy.proxyDo.UseDB(db, opts...)
	_proxy.proxyDo.UseModel(&models.Proxy{})

	tableName := _proxy.proxyDo.TableName()
	_proxy.ALL = field.NewAsterisk(tableName)
	_proxy.Id = field.NewString(tableName, "id")
	_proxy.CreatedAt = field.NewTime(tableName, "createdAt")
	_proxy.UpdatedAt = field.NewTime(tableName, "updatedAt")
	_proxy.Name = field.NewString(tableName, "name")
	_proxy.Active = field.NewBool(tableName, "active")
	_proxy.ProxyId = field.NewString(tableName, "proxyId")
	_proxy.LastSeen = field.NewTime(tableName, "lastSeen")
	_proxy.OrganizationId = field.NewString(tableName, "organizationId")
	_proxy.Organization = proxyBelongsToOrganization{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Organization", "models.Organization"),
	}

	_proxy.fillFieldMap()

	return _proxy
}

type proxy struct {
	proxyDo

	ALL            field.Asterisk
	Id             field.String
	CreatedAt      field.Time
	UpdatedAt      field.Time
	Name           field.String
	Active         field.Bool
	ProxyId        field.String
	LastSeen       field.Time
	OrganizationId field.String
	Organization   proxyBelongsToOrganization

	fieldMap map[string]field.Expr
}

func (p proxy) Table(newTableName string) *proxy {
	p.proxyDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p proxy) As(alias string) *proxy {
	p.proxyDo.DO = *(p.proxyDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *proxy) updateTableName(table string) *proxy {
	p.ALL = field.NewAsterisk(table)
	p.Id = field.NewString(table, "id")
	p.CreatedAt = field.NewTime(table, "createdAt")
	p.UpdatedAt = field.NewTime(table, "updatedAt")
	p.Name = field.NewString(table, "name")
	p.Active = field.NewBool(table, "active")
	p.ProxyId = field.NewString(table, "proxyId")
	p.LastSeen = field.NewTime(table, "lastSeen")
	p.OrganizationId = field.NewString(table, "organizationId")

	p.fillFieldMap()

	return p
}

func (p *proxy) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *proxy) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 9)
	p.fieldMap["id"] = p.Id
	p.fieldMap["createdAt"] = p.CreatedAt
	p.fieldMap["updatedAt"] = p.UpdatedAt
	p.fieldMap["name"] = p.Name
	p.fieldMap["active"] = p.Active
	p.fieldMap["proxyId"] = p.ProxyId
	p.fieldMap["lastSeen"] = p.LastSeen
	p.fieldMap["organizationId"] = p.OrganizationId

}

func (p proxy) clone(db *gorm.DB) proxy {
	p.proxyDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p proxy) replaceDB(db *gorm.DB) proxy {
	p.proxyDo.ReplaceDB(db)
	return p
}

type proxyBelongsToOrganization struct {
	db *gorm.DB

	field.RelationField
}

func (a proxyBelongsToOrganization) Where(conds ...field.Expr) *proxyBelongsToOrganization {
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

func (a proxyBelongsToOrganization) WithContext(ctx context.Context) *proxyBelongsToOrganization {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a proxyBelongsToOrganization) Session(session *gorm.Session) *proxyBelongsToOrganization {
	a.db = a.db.Session(session)
	return &a
}

func (a proxyBelongsToOrganization) Model(m *models.Proxy) *proxyBelongsToOrganizationTx {
	return &proxyBelongsToOrganizationTx{a.db.Model(m).Association(a.Name())}
}

type proxyBelongsToOrganizationTx struct{ tx *gorm.Association }

func (a proxyBelongsToOrganizationTx) Find() (result *models.Organization, err error) {
	return result, a.tx.Find(&result)
}

func (a proxyBelongsToOrganizationTx) Append(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a proxyBelongsToOrganizationTx) Replace(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a proxyBelongsToOrganizationTx) Delete(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a proxyBelongsToOrganizationTx) Clear() error {
	return a.tx.Clear()
}

func (a proxyBelongsToOrganizationTx) Count() int64 {
	return a.tx.Count()
}

type proxyDo struct{ gen.DO }

type IProxyDo interface {
	gen.SubQuery
	Debug() IProxyDo
	WithContext(ctx context.Context) IProxyDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IProxyDo
	WriteDB() IProxyDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IProxyDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IProxyDo
	Not(conds ...gen.Condition) IProxyDo
	Or(conds ...gen.Condition) IProxyDo
	Select(conds ...field.Expr) IProxyDo
	Where(conds ...gen.Condition) IProxyDo
	Order(conds ...field.Expr) IProxyDo
	Distinct(cols ...field.Expr) IProxyDo
	Omit(cols ...field.Expr) IProxyDo
	Join(table schema.Tabler, on ...field.Expr) IProxyDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IProxyDo
	RightJoin(table schema.Tabler, on ...field.Expr) IProxyDo
	Group(cols ...field.Expr) IProxyDo
	Having(conds ...gen.Condition) IProxyDo
	Limit(limit int) IProxyDo
	Offset(offset int) IProxyDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IProxyDo
	Unscoped() IProxyDo
	Create(values ...*models.Proxy) error
	CreateInBatches(values []*models.Proxy, batchSize int) error
	Save(values ...*models.Proxy) error
	First() (*models.Proxy, error)
	Take() (*models.Proxy, error)
	Last() (*models.Proxy, error)
	Find() ([]*models.Proxy, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Proxy, err error)
	FindInBatches(result *[]*models.Proxy, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Proxy) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IProxyDo
	Assign(attrs ...field.AssignExpr) IProxyDo
	Joins(fields ...field.RelationField) IProxyDo
	Preload(fields ...field.RelationField) IProxyDo
	FirstOrInit() (*models.Proxy, error)
	FirstOrCreate() (*models.Proxy, error)
	FindByPage(offset int, limit int) (result []*models.Proxy, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IProxyDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p proxyDo) Debug() IProxyDo {
	return p.withDO(p.DO.Debug())
}

func (p proxyDo) WithContext(ctx context.Context) IProxyDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p proxyDo) ReadDB() IProxyDo {
	return p.Clauses(dbresolver.Read)
}

func (p proxyDo) WriteDB() IProxyDo {
	return p.Clauses(dbresolver.Write)
}

func (p proxyDo) Session(config *gorm.Session) IProxyDo {
	return p.withDO(p.DO.Session(config))
}

func (p proxyDo) Clauses(conds ...clause.Expression) IProxyDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p proxyDo) Returning(value interface{}, columns ...string) IProxyDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p proxyDo) Not(conds ...gen.Condition) IProxyDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p proxyDo) Or(conds ...gen.Condition) IProxyDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p proxyDo) Select(conds ...field.Expr) IProxyDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p proxyDo) Where(conds ...gen.Condition) IProxyDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p proxyDo) Order(conds ...field.Expr) IProxyDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p proxyDo) Distinct(cols ...field.Expr) IProxyDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p proxyDo) Omit(cols ...field.Expr) IProxyDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p proxyDo) Join(table schema.Tabler, on ...field.Expr) IProxyDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p proxyDo) LeftJoin(table schema.Tabler, on ...field.Expr) IProxyDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p proxyDo) RightJoin(table schema.Tabler, on ...field.Expr) IProxyDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p proxyDo) Group(cols ...field.Expr) IProxyDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p proxyDo) Having(conds ...gen.Condition) IProxyDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p proxyDo) Limit(limit int) IProxyDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p proxyDo) Offset(offset int) IProxyDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p proxyDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IProxyDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p proxyDo) Unscoped() IProxyDo {
	return p.withDO(p.DO.Unscoped())
}

func (p proxyDo) Create(values ...*models.Proxy) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p proxyDo) CreateInBatches(values []*models.Proxy, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p proxyDo) Save(values ...*models.Proxy) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p proxyDo) First() (*models.Proxy, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Proxy), nil
	}
}

func (p proxyDo) Take() (*models.Proxy, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Proxy), nil
	}
}

func (p proxyDo) Last() (*models.Proxy, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Proxy), nil
	}
}

func (p proxyDo) Find() ([]*models.Proxy, error) {
	result, err := p.DO.Find()
	return result.([]*models.Proxy), err
}

func (p proxyDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Proxy, err error) {
	buf := make([]*models.Proxy, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p proxyDo) FindInBatches(result *[]*models.Proxy, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p proxyDo) Attrs(attrs ...field.AssignExpr) IProxyDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p proxyDo) Assign(attrs ...field.AssignExpr) IProxyDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p proxyDo) Joins(fields ...field.RelationField) IProxyDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p proxyDo) Preload(fields ...field.RelationField) IProxyDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p proxyDo) FirstOrInit() (*models.Proxy, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Proxy), nil
	}
}

func (p proxyDo) FirstOrCreate() (*models.Proxy, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Proxy), nil
	}
}

func (p proxyDo) FindByPage(offset int, limit int) (result []*models.Proxy, count int64, err error) {
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

func (p proxyDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p proxyDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p proxyDo) Delete(models ...*models.Proxy) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *proxyDo) withDO(do gen.Dao) *proxyDo {
	p.DO = *do.(*gen.DO)
	return p
}
