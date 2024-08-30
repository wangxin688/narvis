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

func newLocation(db *gorm.DB, opts ...gen.DOOption) location {
	_location := location{}

	_location.locationDo.UseDB(db, opts...)
	_location.locationDo.UseModel(&models.Location{})

	tableName := _location.locationDo.TableName()
	_location.ALL = field.NewAsterisk(tableName)
	_location.Id = field.NewString(tableName, "id")
	_location.CreatedAt = field.NewTime(tableName, "createdAt")
	_location.UpdatedAt = field.NewTime(tableName, "updatedAt")
	_location.Name = field.NewString(tableName, "name")
	_location.Description = field.NewString(tableName, "description")
	_location.ParentId = field.NewString(tableName, "parentId")
	_location.SiteId = field.NewString(tableName, "siteId")
	_location.OrganizationId = field.NewString(tableName, "organizationId")
	_location.Parent = locationBelongsToParent{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Parent", "models.Location"),
		Parent: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Parent.Parent", "models.Location"),
		},
		Site: struct {
			field.RelationField
			Organization struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Parent.Site", "models.Site"),
			Organization: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Parent.Site.Organization", "models.Organization"),
			},
		},
		Organization: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Parent.Organization", "models.Organization"),
		},
	}

	_location.Site = locationBelongsToSite{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Site", "models.Site"),
	}

	_location.Organization = locationBelongsToOrganization{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Organization", "models.Organization"),
	}

	_location.fillFieldMap()

	return _location
}

type location struct {
	locationDo

	ALL            field.Asterisk
	Id             field.String
	CreatedAt      field.Time
	UpdatedAt      field.Time
	Name           field.String
	Description    field.String
	ParentId       field.String
	SiteId         field.String
	OrganizationId field.String
	Parent         locationBelongsToParent

	Site locationBelongsToSite

	Organization locationBelongsToOrganization

	fieldMap map[string]field.Expr
}

func (l location) Table(newTableName string) *location {
	l.locationDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l location) As(alias string) *location {
	l.locationDo.DO = *(l.locationDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *location) updateTableName(table string) *location {
	l.ALL = field.NewAsterisk(table)
	l.Id = field.NewString(table, "id")
	l.CreatedAt = field.NewTime(table, "createdAt")
	l.UpdatedAt = field.NewTime(table, "updatedAt")
	l.Name = field.NewString(table, "name")
	l.Description = field.NewString(table, "description")
	l.ParentId = field.NewString(table, "parentId")
	l.SiteId = field.NewString(table, "siteId")
	l.OrganizationId = field.NewString(table, "organizationId")

	l.fillFieldMap()

	return l
}

func (l *location) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *location) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 11)
	l.fieldMap["id"] = l.Id
	l.fieldMap["createdAt"] = l.CreatedAt
	l.fieldMap["updatedAt"] = l.UpdatedAt
	l.fieldMap["name"] = l.Name
	l.fieldMap["description"] = l.Description
	l.fieldMap["parentId"] = l.ParentId
	l.fieldMap["siteId"] = l.SiteId
	l.fieldMap["organizationId"] = l.OrganizationId

}

func (l location) clone(db *gorm.DB) location {
	l.locationDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l location) replaceDB(db *gorm.DB) location {
	l.locationDo.ReplaceDB(db)
	return l
}

type locationBelongsToParent struct {
	db *gorm.DB

	field.RelationField

	Parent struct {
		field.RelationField
	}
	Site struct {
		field.RelationField
		Organization struct {
			field.RelationField
		}
	}
	Organization struct {
		field.RelationField
	}
}

func (a locationBelongsToParent) Where(conds ...field.Expr) *locationBelongsToParent {
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

func (a locationBelongsToParent) WithContext(ctx context.Context) *locationBelongsToParent {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a locationBelongsToParent) Session(session *gorm.Session) *locationBelongsToParent {
	a.db = a.db.Session(session)
	return &a
}

func (a locationBelongsToParent) Model(m *models.Location) *locationBelongsToParentTx {
	return &locationBelongsToParentTx{a.db.Model(m).Association(a.Name())}
}

type locationBelongsToParentTx struct{ tx *gorm.Association }

func (a locationBelongsToParentTx) Find() (result *models.Location, err error) {
	return result, a.tx.Find(&result)
}

func (a locationBelongsToParentTx) Append(values ...*models.Location) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a locationBelongsToParentTx) Replace(values ...*models.Location) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a locationBelongsToParentTx) Delete(values ...*models.Location) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a locationBelongsToParentTx) Clear() error {
	return a.tx.Clear()
}

func (a locationBelongsToParentTx) Count() int64 {
	return a.tx.Count()
}

type locationBelongsToSite struct {
	db *gorm.DB

	field.RelationField
}

func (a locationBelongsToSite) Where(conds ...field.Expr) *locationBelongsToSite {
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

func (a locationBelongsToSite) WithContext(ctx context.Context) *locationBelongsToSite {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a locationBelongsToSite) Session(session *gorm.Session) *locationBelongsToSite {
	a.db = a.db.Session(session)
	return &a
}

func (a locationBelongsToSite) Model(m *models.Location) *locationBelongsToSiteTx {
	return &locationBelongsToSiteTx{a.db.Model(m).Association(a.Name())}
}

type locationBelongsToSiteTx struct{ tx *gorm.Association }

func (a locationBelongsToSiteTx) Find() (result *models.Site, err error) {
	return result, a.tx.Find(&result)
}

func (a locationBelongsToSiteTx) Append(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a locationBelongsToSiteTx) Replace(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a locationBelongsToSiteTx) Delete(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a locationBelongsToSiteTx) Clear() error {
	return a.tx.Clear()
}

func (a locationBelongsToSiteTx) Count() int64 {
	return a.tx.Count()
}

type locationBelongsToOrganization struct {
	db *gorm.DB

	field.RelationField
}

func (a locationBelongsToOrganization) Where(conds ...field.Expr) *locationBelongsToOrganization {
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

func (a locationBelongsToOrganization) WithContext(ctx context.Context) *locationBelongsToOrganization {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a locationBelongsToOrganization) Session(session *gorm.Session) *locationBelongsToOrganization {
	a.db = a.db.Session(session)
	return &a
}

func (a locationBelongsToOrganization) Model(m *models.Location) *locationBelongsToOrganizationTx {
	return &locationBelongsToOrganizationTx{a.db.Model(m).Association(a.Name())}
}

type locationBelongsToOrganizationTx struct{ tx *gorm.Association }

func (a locationBelongsToOrganizationTx) Find() (result *models.Organization, err error) {
	return result, a.tx.Find(&result)
}

func (a locationBelongsToOrganizationTx) Append(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a locationBelongsToOrganizationTx) Replace(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a locationBelongsToOrganizationTx) Delete(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a locationBelongsToOrganizationTx) Clear() error {
	return a.tx.Clear()
}

func (a locationBelongsToOrganizationTx) Count() int64 {
	return a.tx.Count()
}

type locationDo struct{ gen.DO }

type ILocationDo interface {
	gen.SubQuery
	Debug() ILocationDo
	WithContext(ctx context.Context) ILocationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ILocationDo
	WriteDB() ILocationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ILocationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILocationDo
	Not(conds ...gen.Condition) ILocationDo
	Or(conds ...gen.Condition) ILocationDo
	Select(conds ...field.Expr) ILocationDo
	Where(conds ...gen.Condition) ILocationDo
	Order(conds ...field.Expr) ILocationDo
	Distinct(cols ...field.Expr) ILocationDo
	Omit(cols ...field.Expr) ILocationDo
	Join(table schema.Tabler, on ...field.Expr) ILocationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILocationDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILocationDo
	Group(cols ...field.Expr) ILocationDo
	Having(conds ...gen.Condition) ILocationDo
	Limit(limit int) ILocationDo
	Offset(offset int) ILocationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILocationDo
	Unscoped() ILocationDo
	Create(values ...*models.Location) error
	CreateInBatches(values []*models.Location, batchSize int) error
	Save(values ...*models.Location) error
	First() (*models.Location, error)
	Take() (*models.Location, error)
	Last() (*models.Location, error)
	Find() ([]*models.Location, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Location, err error)
	FindInBatches(result *[]*models.Location, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Location) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILocationDo
	Assign(attrs ...field.AssignExpr) ILocationDo
	Joins(fields ...field.RelationField) ILocationDo
	Preload(fields ...field.RelationField) ILocationDo
	FirstOrInit() (*models.Location, error)
	FirstOrCreate() (*models.Location, error)
	FindByPage(offset int, limit int) (result []*models.Location, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILocationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l locationDo) Debug() ILocationDo {
	return l.withDO(l.DO.Debug())
}

func (l locationDo) WithContext(ctx context.Context) ILocationDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l locationDo) ReadDB() ILocationDo {
	return l.Clauses(dbresolver.Read)
}

func (l locationDo) WriteDB() ILocationDo {
	return l.Clauses(dbresolver.Write)
}

func (l locationDo) Session(config *gorm.Session) ILocationDo {
	return l.withDO(l.DO.Session(config))
}

func (l locationDo) Clauses(conds ...clause.Expression) ILocationDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l locationDo) Returning(value interface{}, columns ...string) ILocationDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l locationDo) Not(conds ...gen.Condition) ILocationDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l locationDo) Or(conds ...gen.Condition) ILocationDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l locationDo) Select(conds ...field.Expr) ILocationDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l locationDo) Where(conds ...gen.Condition) ILocationDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l locationDo) Order(conds ...field.Expr) ILocationDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l locationDo) Distinct(cols ...field.Expr) ILocationDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l locationDo) Omit(cols ...field.Expr) ILocationDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l locationDo) Join(table schema.Tabler, on ...field.Expr) ILocationDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l locationDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILocationDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l locationDo) RightJoin(table schema.Tabler, on ...field.Expr) ILocationDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l locationDo) Group(cols ...field.Expr) ILocationDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l locationDo) Having(conds ...gen.Condition) ILocationDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l locationDo) Limit(limit int) ILocationDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l locationDo) Offset(offset int) ILocationDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l locationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILocationDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l locationDo) Unscoped() ILocationDo {
	return l.withDO(l.DO.Unscoped())
}

func (l locationDo) Create(values ...*models.Location) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l locationDo) CreateInBatches(values []*models.Location, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l locationDo) Save(values ...*models.Location) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l locationDo) First() (*models.Location, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Location), nil
	}
}

func (l locationDo) Take() (*models.Location, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Location), nil
	}
}

func (l locationDo) Last() (*models.Location, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Location), nil
	}
}

func (l locationDo) Find() ([]*models.Location, error) {
	result, err := l.DO.Find()
	return result.([]*models.Location), err
}

func (l locationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Location, err error) {
	buf := make([]*models.Location, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l locationDo) FindInBatches(result *[]*models.Location, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l locationDo) Attrs(attrs ...field.AssignExpr) ILocationDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l locationDo) Assign(attrs ...field.AssignExpr) ILocationDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l locationDo) Joins(fields ...field.RelationField) ILocationDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l locationDo) Preload(fields ...field.RelationField) ILocationDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l locationDo) FirstOrInit() (*models.Location, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Location), nil
	}
}

func (l locationDo) FirstOrCreate() (*models.Location, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Location), nil
	}
}

func (l locationDo) FindByPage(offset int, limit int) (result []*models.Location, count int64, err error) {
	result, err = l.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = l.Offset(-1).Limit(-1).Count()
	return
}

func (l locationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l locationDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l locationDo) Delete(models ...*models.Location) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *locationDo) withDO(do gen.Dao) *locationDo {
	l.DO = *do.(*gen.DO)
	return l
}
