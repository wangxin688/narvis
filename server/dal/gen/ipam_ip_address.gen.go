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

func newIpAddress(db *gorm.DB, opts ...gen.DOOption) ipAddress {
	_ipAddress := ipAddress{}

	_ipAddress.ipAddressDo.UseDB(db, opts...)
	_ipAddress.ipAddressDo.UseModel(&models.IpAddress{})

	tableName := _ipAddress.ipAddressDo.TableName()
	_ipAddress.ALL = field.NewAsterisk(tableName)
	_ipAddress.Id = field.NewString(tableName, "id")
	_ipAddress.CreatedAt = field.NewTime(tableName, "createdAt")
	_ipAddress.UpdatedAt = field.NewTime(tableName, "updatedAt")
	_ipAddress.Address = field.NewString(tableName, "address")
	_ipAddress.Network = field.NewString(tableName, "network")
	_ipAddress.Status = field.NewString(tableName, "status")
	_ipAddress.MacAddress = field.NewString(tableName, "macAddress")
	_ipAddress.Description = field.NewString(tableName, "description")
	_ipAddress.Type = field.NewString(tableName, "type")
	_ipAddress.SiteId = field.NewString(tableName, "siteId")
	_ipAddress.OrganizationId = field.NewString(tableName, "organizationId")
	_ipAddress.Site = ipAddressBelongsToSite{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Site", "models.Site"),
		Organization: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Site.Organization", "models.Organization"),
		},
	}

	_ipAddress.Organization = ipAddressBelongsToOrganization{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Organization", "models.Organization"),
	}

	_ipAddress.fillFieldMap()

	return _ipAddress
}

type ipAddress struct {
	ipAddressDo

	ALL            field.Asterisk
	Id             field.String
	CreatedAt      field.Time
	UpdatedAt      field.Time
	Address        field.String
	Network        field.String
	Status         field.String
	MacAddress     field.String
	Description    field.String
	Type           field.String
	SiteId         field.String
	OrganizationId field.String
	Site           ipAddressBelongsToSite

	Organization ipAddressBelongsToOrganization

	fieldMap map[string]field.Expr
}

func (i ipAddress) Table(newTableName string) *ipAddress {
	i.ipAddressDo.UseTable(newTableName)
	return i.updateTableName(newTableName)
}

func (i ipAddress) As(alias string) *ipAddress {
	i.ipAddressDo.DO = *(i.ipAddressDo.As(alias).(*gen.DO))
	return i.updateTableName(alias)
}

func (i *ipAddress) updateTableName(table string) *ipAddress {
	i.ALL = field.NewAsterisk(table)
	i.Id = field.NewString(table, "id")
	i.CreatedAt = field.NewTime(table, "createdAt")
	i.UpdatedAt = field.NewTime(table, "updatedAt")
	i.Address = field.NewString(table, "address")
	i.Network = field.NewString(table, "network")
	i.Status = field.NewString(table, "status")
	i.MacAddress = field.NewString(table, "macAddress")
	i.Description = field.NewString(table, "description")
	i.Type = field.NewString(table, "type")
	i.SiteId = field.NewString(table, "siteId")
	i.OrganizationId = field.NewString(table, "organizationId")

	i.fillFieldMap()

	return i
}

func (i *ipAddress) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := i.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (i *ipAddress) fillFieldMap() {
	i.fieldMap = make(map[string]field.Expr, 13)
	i.fieldMap["id"] = i.Id
	i.fieldMap["createdAt"] = i.CreatedAt
	i.fieldMap["updatedAt"] = i.UpdatedAt
	i.fieldMap["address"] = i.Address
	i.fieldMap["network"] = i.Network
	i.fieldMap["status"] = i.Status
	i.fieldMap["macAddress"] = i.MacAddress
	i.fieldMap["description"] = i.Description
	i.fieldMap["type"] = i.Type
	i.fieldMap["siteId"] = i.SiteId
	i.fieldMap["organizationId"] = i.OrganizationId

}

func (i ipAddress) clone(db *gorm.DB) ipAddress {
	i.ipAddressDo.ReplaceConnPool(db.Statement.ConnPool)
	return i
}

func (i ipAddress) replaceDB(db *gorm.DB) ipAddress {
	i.ipAddressDo.ReplaceDB(db)
	return i
}

type ipAddressBelongsToSite struct {
	db *gorm.DB

	field.RelationField

	Organization struct {
		field.RelationField
	}
}

func (a ipAddressBelongsToSite) Where(conds ...field.Expr) *ipAddressBelongsToSite {
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

func (a ipAddressBelongsToSite) WithContext(ctx context.Context) *ipAddressBelongsToSite {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a ipAddressBelongsToSite) Session(session *gorm.Session) *ipAddressBelongsToSite {
	a.db = a.db.Session(session)
	return &a
}

func (a ipAddressBelongsToSite) Model(m *models.IpAddress) *ipAddressBelongsToSiteTx {
	return &ipAddressBelongsToSiteTx{a.db.Model(m).Association(a.Name())}
}

type ipAddressBelongsToSiteTx struct{ tx *gorm.Association }

func (a ipAddressBelongsToSiteTx) Find() (result *models.Site, err error) {
	return result, a.tx.Find(&result)
}

func (a ipAddressBelongsToSiteTx) Append(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a ipAddressBelongsToSiteTx) Replace(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a ipAddressBelongsToSiteTx) Delete(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a ipAddressBelongsToSiteTx) Clear() error {
	return a.tx.Clear()
}

func (a ipAddressBelongsToSiteTx) Count() int64 {
	return a.tx.Count()
}

type ipAddressBelongsToOrganization struct {
	db *gorm.DB

	field.RelationField
}

func (a ipAddressBelongsToOrganization) Where(conds ...field.Expr) *ipAddressBelongsToOrganization {
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

func (a ipAddressBelongsToOrganization) WithContext(ctx context.Context) *ipAddressBelongsToOrganization {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a ipAddressBelongsToOrganization) Session(session *gorm.Session) *ipAddressBelongsToOrganization {
	a.db = a.db.Session(session)
	return &a
}

func (a ipAddressBelongsToOrganization) Model(m *models.IpAddress) *ipAddressBelongsToOrganizationTx {
	return &ipAddressBelongsToOrganizationTx{a.db.Model(m).Association(a.Name())}
}

type ipAddressBelongsToOrganizationTx struct{ tx *gorm.Association }

func (a ipAddressBelongsToOrganizationTx) Find() (result *models.Organization, err error) {
	return result, a.tx.Find(&result)
}

func (a ipAddressBelongsToOrganizationTx) Append(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a ipAddressBelongsToOrganizationTx) Replace(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a ipAddressBelongsToOrganizationTx) Delete(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a ipAddressBelongsToOrganizationTx) Clear() error {
	return a.tx.Clear()
}

func (a ipAddressBelongsToOrganizationTx) Count() int64 {
	return a.tx.Count()
}

type ipAddressDo struct{ gen.DO }

type IIpAddressDo interface {
	gen.SubQuery
	Debug() IIpAddressDo
	WithContext(ctx context.Context) IIpAddressDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IIpAddressDo
	WriteDB() IIpAddressDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IIpAddressDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IIpAddressDo
	Not(conds ...gen.Condition) IIpAddressDo
	Or(conds ...gen.Condition) IIpAddressDo
	Select(conds ...field.Expr) IIpAddressDo
	Where(conds ...gen.Condition) IIpAddressDo
	Order(conds ...field.Expr) IIpAddressDo
	Distinct(cols ...field.Expr) IIpAddressDo
	Omit(cols ...field.Expr) IIpAddressDo
	Join(table schema.Tabler, on ...field.Expr) IIpAddressDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IIpAddressDo
	RightJoin(table schema.Tabler, on ...field.Expr) IIpAddressDo
	Group(cols ...field.Expr) IIpAddressDo
	Having(conds ...gen.Condition) IIpAddressDo
	Limit(limit int) IIpAddressDo
	Offset(offset int) IIpAddressDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IIpAddressDo
	Unscoped() IIpAddressDo
	Create(values ...*models.IpAddress) error
	CreateInBatches(values []*models.IpAddress, batchSize int) error
	Save(values ...*models.IpAddress) error
	First() (*models.IpAddress, error)
	Take() (*models.IpAddress, error)
	Last() (*models.IpAddress, error)
	Find() ([]*models.IpAddress, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.IpAddress, err error)
	FindInBatches(result *[]*models.IpAddress, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.IpAddress) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IIpAddressDo
	Assign(attrs ...field.AssignExpr) IIpAddressDo
	Joins(fields ...field.RelationField) IIpAddressDo
	Preload(fields ...field.RelationField) IIpAddressDo
	FirstOrInit() (*models.IpAddress, error)
	FirstOrCreate() (*models.IpAddress, error)
	FindByPage(offset int, limit int) (result []*models.IpAddress, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IIpAddressDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (i ipAddressDo) Debug() IIpAddressDo {
	return i.withDO(i.DO.Debug())
}

func (i ipAddressDo) WithContext(ctx context.Context) IIpAddressDo {
	return i.withDO(i.DO.WithContext(ctx))
}

func (i ipAddressDo) ReadDB() IIpAddressDo {
	return i.Clauses(dbresolver.Read)
}

func (i ipAddressDo) WriteDB() IIpAddressDo {
	return i.Clauses(dbresolver.Write)
}

func (i ipAddressDo) Session(config *gorm.Session) IIpAddressDo {
	return i.withDO(i.DO.Session(config))
}

func (i ipAddressDo) Clauses(conds ...clause.Expression) IIpAddressDo {
	return i.withDO(i.DO.Clauses(conds...))
}

func (i ipAddressDo) Returning(value interface{}, columns ...string) IIpAddressDo {
	return i.withDO(i.DO.Returning(value, columns...))
}

func (i ipAddressDo) Not(conds ...gen.Condition) IIpAddressDo {
	return i.withDO(i.DO.Not(conds...))
}

func (i ipAddressDo) Or(conds ...gen.Condition) IIpAddressDo {
	return i.withDO(i.DO.Or(conds...))
}

func (i ipAddressDo) Select(conds ...field.Expr) IIpAddressDo {
	return i.withDO(i.DO.Select(conds...))
}

func (i ipAddressDo) Where(conds ...gen.Condition) IIpAddressDo {
	return i.withDO(i.DO.Where(conds...))
}

func (i ipAddressDo) Order(conds ...field.Expr) IIpAddressDo {
	return i.withDO(i.DO.Order(conds...))
}

func (i ipAddressDo) Distinct(cols ...field.Expr) IIpAddressDo {
	return i.withDO(i.DO.Distinct(cols...))
}

func (i ipAddressDo) Omit(cols ...field.Expr) IIpAddressDo {
	return i.withDO(i.DO.Omit(cols...))
}

func (i ipAddressDo) Join(table schema.Tabler, on ...field.Expr) IIpAddressDo {
	return i.withDO(i.DO.Join(table, on...))
}

func (i ipAddressDo) LeftJoin(table schema.Tabler, on ...field.Expr) IIpAddressDo {
	return i.withDO(i.DO.LeftJoin(table, on...))
}

func (i ipAddressDo) RightJoin(table schema.Tabler, on ...field.Expr) IIpAddressDo {
	return i.withDO(i.DO.RightJoin(table, on...))
}

func (i ipAddressDo) Group(cols ...field.Expr) IIpAddressDo {
	return i.withDO(i.DO.Group(cols...))
}

func (i ipAddressDo) Having(conds ...gen.Condition) IIpAddressDo {
	return i.withDO(i.DO.Having(conds...))
}

func (i ipAddressDo) Limit(limit int) IIpAddressDo {
	return i.withDO(i.DO.Limit(limit))
}

func (i ipAddressDo) Offset(offset int) IIpAddressDo {
	return i.withDO(i.DO.Offset(offset))
}

func (i ipAddressDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IIpAddressDo {
	return i.withDO(i.DO.Scopes(funcs...))
}

func (i ipAddressDo) Unscoped() IIpAddressDo {
	return i.withDO(i.DO.Unscoped())
}

func (i ipAddressDo) Create(values ...*models.IpAddress) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Create(values)
}

func (i ipAddressDo) CreateInBatches(values []*models.IpAddress, batchSize int) error {
	return i.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (i ipAddressDo) Save(values ...*models.IpAddress) error {
	if len(values) == 0 {
		return nil
	}
	return i.DO.Save(values)
}

func (i ipAddressDo) First() (*models.IpAddress, error) {
	if result, err := i.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.IpAddress), nil
	}
}

func (i ipAddressDo) Take() (*models.IpAddress, error) {
	if result, err := i.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.IpAddress), nil
	}
}

func (i ipAddressDo) Last() (*models.IpAddress, error) {
	if result, err := i.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.IpAddress), nil
	}
}

func (i ipAddressDo) Find() ([]*models.IpAddress, error) {
	result, err := i.DO.Find()
	return result.([]*models.IpAddress), err
}

func (i ipAddressDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.IpAddress, err error) {
	buf := make([]*models.IpAddress, 0, batchSize)
	err = i.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (i ipAddressDo) FindInBatches(result *[]*models.IpAddress, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return i.DO.FindInBatches(result, batchSize, fc)
}

func (i ipAddressDo) Attrs(attrs ...field.AssignExpr) IIpAddressDo {
	return i.withDO(i.DO.Attrs(attrs...))
}

func (i ipAddressDo) Assign(attrs ...field.AssignExpr) IIpAddressDo {
	return i.withDO(i.DO.Assign(attrs...))
}

func (i ipAddressDo) Joins(fields ...field.RelationField) IIpAddressDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Joins(_f))
	}
	return &i
}

func (i ipAddressDo) Preload(fields ...field.RelationField) IIpAddressDo {
	for _, _f := range fields {
		i = *i.withDO(i.DO.Preload(_f))
	}
	return &i
}

func (i ipAddressDo) FirstOrInit() (*models.IpAddress, error) {
	if result, err := i.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.IpAddress), nil
	}
}

func (i ipAddressDo) FirstOrCreate() (*models.IpAddress, error) {
	if result, err := i.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.IpAddress), nil
	}
}

func (i ipAddressDo) FindByPage(offset int, limit int) (result []*models.IpAddress, count int64, err error) {
	result, err = i.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = i.Offset(-1).Limit(-1).Count()
	return
}

func (i ipAddressDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = i.Count()
	if err != nil {
		return
	}

	err = i.Offset(offset).Limit(limit).Scan(result)
	return
}

func (i ipAddressDo) Scan(result interface{}) (err error) {
	return i.DO.Scan(result)
}

func (i ipAddressDo) Delete(models ...*models.IpAddress) (result gen.ResultInfo, err error) {
	return i.DO.Delete(models)
}

func (i *ipAddressDo) withDO(do gen.Dao) *ipAddressDo {
	i.DO = *do.(*gen.DO)
	return i
}
