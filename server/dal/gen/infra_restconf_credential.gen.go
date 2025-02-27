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

func newRestconfCredential(db *gorm.DB, opts ...gen.DOOption) restconfCredential {
	_restconfCredential := restconfCredential{}

	_restconfCredential.restconfCredentialDo.UseDB(db, opts...)
	_restconfCredential.restconfCredentialDo.UseModel(&models.RestconfCredential{})

	tableName := _restconfCredential.restconfCredentialDo.TableName()
	_restconfCredential.ALL = field.NewAsterisk(tableName)
	_restconfCredential.Id = field.NewString(tableName, "id")
	_restconfCredential.CreatedAt = field.NewTime(tableName, "createdAt")
	_restconfCredential.UpdatedAt = field.NewTime(tableName, "updatedAt")
	_restconfCredential.Url = field.NewString(tableName, "url")
	_restconfCredential.Username = field.NewString(tableName, "username")
	_restconfCredential.Password = field.NewString(tableName, "password")
	_restconfCredential.DeviceId = field.NewString(tableName, "deviceId")
	_restconfCredential.OrganizationId = field.NewString(tableName, "organizationId")
	_restconfCredential.Device = restconfCredentialBelongsToDevice{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Device", "models.Device"),
		Rack: struct {
			field.RelationField
			Site struct {
				field.RelationField
				Organization struct {
					field.RelationField
				}
			}
			Organization struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Device.Rack", "models.Rack"),
			Site: struct {
				field.RelationField
				Organization struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("Device.Rack.Site", "models.Site"),
				Organization: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Device.Rack.Site.Organization", "models.Organization"),
				},
			},
			Organization: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Device.Rack.Organization", "models.Organization"),
			},
		},
		Template: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Device.Template", "models.Template"),
		},
		Site: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Device.Site", "models.Site"),
		},
		Organization: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Device.Organization", "models.Organization"),
		},
	}

	_restconfCredential.Organization = restconfCredentialBelongsToOrganization{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Organization", "models.Organization"),
	}

	_restconfCredential.fillFieldMap()

	return _restconfCredential
}

type restconfCredential struct {
	restconfCredentialDo

	ALL            field.Asterisk
	Id             field.String
	CreatedAt      field.Time
	UpdatedAt      field.Time
	Url            field.String
	Username       field.String
	Password       field.String
	DeviceId       field.String
	OrganizationId field.String
	Device         restconfCredentialBelongsToDevice

	Organization restconfCredentialBelongsToOrganization

	fieldMap map[string]field.Expr
}

func (r restconfCredential) Table(newTableName string) *restconfCredential {
	r.restconfCredentialDo.UseTable(newTableName)
	return r.updateTableName(newTableName)
}

func (r restconfCredential) As(alias string) *restconfCredential {
	r.restconfCredentialDo.DO = *(r.restconfCredentialDo.As(alias).(*gen.DO))
	return r.updateTableName(alias)
}

func (r *restconfCredential) updateTableName(table string) *restconfCredential {
	r.ALL = field.NewAsterisk(table)
	r.Id = field.NewString(table, "id")
	r.CreatedAt = field.NewTime(table, "createdAt")
	r.UpdatedAt = field.NewTime(table, "updatedAt")
	r.Url = field.NewString(table, "url")
	r.Username = field.NewString(table, "username")
	r.Password = field.NewString(table, "password")
	r.DeviceId = field.NewString(table, "deviceId")
	r.OrganizationId = field.NewString(table, "organizationId")

	r.fillFieldMap()

	return r
}

func (r *restconfCredential) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := r.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (r *restconfCredential) fillFieldMap() {
	r.fieldMap = make(map[string]field.Expr, 10)
	r.fieldMap["id"] = r.Id
	r.fieldMap["createdAt"] = r.CreatedAt
	r.fieldMap["updatedAt"] = r.UpdatedAt
	r.fieldMap["url"] = r.Url
	r.fieldMap["username"] = r.Username
	r.fieldMap["password"] = r.Password
	r.fieldMap["deviceId"] = r.DeviceId
	r.fieldMap["organizationId"] = r.OrganizationId

}

func (r restconfCredential) clone(db *gorm.DB) restconfCredential {
	r.restconfCredentialDo.ReplaceConnPool(db.Statement.ConnPool)
	return r
}

func (r restconfCredential) replaceDB(db *gorm.DB) restconfCredential {
	r.restconfCredentialDo.ReplaceDB(db)
	return r
}

type restconfCredentialBelongsToDevice struct {
	db *gorm.DB

	field.RelationField

	Rack struct {
		field.RelationField
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
	Template struct {
		field.RelationField
	}
	Site struct {
		field.RelationField
	}
	Organization struct {
		field.RelationField
	}
}

func (a restconfCredentialBelongsToDevice) Where(conds ...field.Expr) *restconfCredentialBelongsToDevice {
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

func (a restconfCredentialBelongsToDevice) WithContext(ctx context.Context) *restconfCredentialBelongsToDevice {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a restconfCredentialBelongsToDevice) Session(session *gorm.Session) *restconfCredentialBelongsToDevice {
	a.db = a.db.Session(session)
	return &a
}

func (a restconfCredentialBelongsToDevice) Model(m *models.RestconfCredential) *restconfCredentialBelongsToDeviceTx {
	return &restconfCredentialBelongsToDeviceTx{a.db.Model(m).Association(a.Name())}
}

type restconfCredentialBelongsToDeviceTx struct{ tx *gorm.Association }

func (a restconfCredentialBelongsToDeviceTx) Find() (result *models.Device, err error) {
	return result, a.tx.Find(&result)
}

func (a restconfCredentialBelongsToDeviceTx) Append(values ...*models.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a restconfCredentialBelongsToDeviceTx) Replace(values ...*models.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a restconfCredentialBelongsToDeviceTx) Delete(values ...*models.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a restconfCredentialBelongsToDeviceTx) Clear() error {
	return a.tx.Clear()
}

func (a restconfCredentialBelongsToDeviceTx) Count() int64 {
	return a.tx.Count()
}

type restconfCredentialBelongsToOrganization struct {
	db *gorm.DB

	field.RelationField
}

func (a restconfCredentialBelongsToOrganization) Where(conds ...field.Expr) *restconfCredentialBelongsToOrganization {
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

func (a restconfCredentialBelongsToOrganization) WithContext(ctx context.Context) *restconfCredentialBelongsToOrganization {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a restconfCredentialBelongsToOrganization) Session(session *gorm.Session) *restconfCredentialBelongsToOrganization {
	a.db = a.db.Session(session)
	return &a
}

func (a restconfCredentialBelongsToOrganization) Model(m *models.RestconfCredential) *restconfCredentialBelongsToOrganizationTx {
	return &restconfCredentialBelongsToOrganizationTx{a.db.Model(m).Association(a.Name())}
}

type restconfCredentialBelongsToOrganizationTx struct{ tx *gorm.Association }

func (a restconfCredentialBelongsToOrganizationTx) Find() (result *models.Organization, err error) {
	return result, a.tx.Find(&result)
}

func (a restconfCredentialBelongsToOrganizationTx) Append(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a restconfCredentialBelongsToOrganizationTx) Replace(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a restconfCredentialBelongsToOrganizationTx) Delete(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a restconfCredentialBelongsToOrganizationTx) Clear() error {
	return a.tx.Clear()
}

func (a restconfCredentialBelongsToOrganizationTx) Count() int64 {
	return a.tx.Count()
}

type restconfCredentialDo struct{ gen.DO }

type IRestconfCredentialDo interface {
	gen.SubQuery
	Debug() IRestconfCredentialDo
	WithContext(ctx context.Context) IRestconfCredentialDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IRestconfCredentialDo
	WriteDB() IRestconfCredentialDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IRestconfCredentialDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IRestconfCredentialDo
	Not(conds ...gen.Condition) IRestconfCredentialDo
	Or(conds ...gen.Condition) IRestconfCredentialDo
	Select(conds ...field.Expr) IRestconfCredentialDo
	Where(conds ...gen.Condition) IRestconfCredentialDo
	Order(conds ...field.Expr) IRestconfCredentialDo
	Distinct(cols ...field.Expr) IRestconfCredentialDo
	Omit(cols ...field.Expr) IRestconfCredentialDo
	Join(table schema.Tabler, on ...field.Expr) IRestconfCredentialDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IRestconfCredentialDo
	RightJoin(table schema.Tabler, on ...field.Expr) IRestconfCredentialDo
	Group(cols ...field.Expr) IRestconfCredentialDo
	Having(conds ...gen.Condition) IRestconfCredentialDo
	Limit(limit int) IRestconfCredentialDo
	Offset(offset int) IRestconfCredentialDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IRestconfCredentialDo
	Unscoped() IRestconfCredentialDo
	Create(values ...*models.RestconfCredential) error
	CreateInBatches(values []*models.RestconfCredential, batchSize int) error
	Save(values ...*models.RestconfCredential) error
	First() (*models.RestconfCredential, error)
	Take() (*models.RestconfCredential, error)
	Last() (*models.RestconfCredential, error)
	Find() ([]*models.RestconfCredential, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.RestconfCredential, err error)
	FindInBatches(result *[]*models.RestconfCredential, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.RestconfCredential) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IRestconfCredentialDo
	Assign(attrs ...field.AssignExpr) IRestconfCredentialDo
	Joins(fields ...field.RelationField) IRestconfCredentialDo
	Preload(fields ...field.RelationField) IRestconfCredentialDo
	FirstOrInit() (*models.RestconfCredential, error)
	FirstOrCreate() (*models.RestconfCredential, error)
	FindByPage(offset int, limit int) (result []*models.RestconfCredential, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IRestconfCredentialDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (r restconfCredentialDo) Debug() IRestconfCredentialDo {
	return r.withDO(r.DO.Debug())
}

func (r restconfCredentialDo) WithContext(ctx context.Context) IRestconfCredentialDo {
	return r.withDO(r.DO.WithContext(ctx))
}

func (r restconfCredentialDo) ReadDB() IRestconfCredentialDo {
	return r.Clauses(dbresolver.Read)
}

func (r restconfCredentialDo) WriteDB() IRestconfCredentialDo {
	return r.Clauses(dbresolver.Write)
}

func (r restconfCredentialDo) Session(config *gorm.Session) IRestconfCredentialDo {
	return r.withDO(r.DO.Session(config))
}

func (r restconfCredentialDo) Clauses(conds ...clause.Expression) IRestconfCredentialDo {
	return r.withDO(r.DO.Clauses(conds...))
}

func (r restconfCredentialDo) Returning(value interface{}, columns ...string) IRestconfCredentialDo {
	return r.withDO(r.DO.Returning(value, columns...))
}

func (r restconfCredentialDo) Not(conds ...gen.Condition) IRestconfCredentialDo {
	return r.withDO(r.DO.Not(conds...))
}

func (r restconfCredentialDo) Or(conds ...gen.Condition) IRestconfCredentialDo {
	return r.withDO(r.DO.Or(conds...))
}

func (r restconfCredentialDo) Select(conds ...field.Expr) IRestconfCredentialDo {
	return r.withDO(r.DO.Select(conds...))
}

func (r restconfCredentialDo) Where(conds ...gen.Condition) IRestconfCredentialDo {
	return r.withDO(r.DO.Where(conds...))
}

func (r restconfCredentialDo) Order(conds ...field.Expr) IRestconfCredentialDo {
	return r.withDO(r.DO.Order(conds...))
}

func (r restconfCredentialDo) Distinct(cols ...field.Expr) IRestconfCredentialDo {
	return r.withDO(r.DO.Distinct(cols...))
}

func (r restconfCredentialDo) Omit(cols ...field.Expr) IRestconfCredentialDo {
	return r.withDO(r.DO.Omit(cols...))
}

func (r restconfCredentialDo) Join(table schema.Tabler, on ...field.Expr) IRestconfCredentialDo {
	return r.withDO(r.DO.Join(table, on...))
}

func (r restconfCredentialDo) LeftJoin(table schema.Tabler, on ...field.Expr) IRestconfCredentialDo {
	return r.withDO(r.DO.LeftJoin(table, on...))
}

func (r restconfCredentialDo) RightJoin(table schema.Tabler, on ...field.Expr) IRestconfCredentialDo {
	return r.withDO(r.DO.RightJoin(table, on...))
}

func (r restconfCredentialDo) Group(cols ...field.Expr) IRestconfCredentialDo {
	return r.withDO(r.DO.Group(cols...))
}

func (r restconfCredentialDo) Having(conds ...gen.Condition) IRestconfCredentialDo {
	return r.withDO(r.DO.Having(conds...))
}

func (r restconfCredentialDo) Limit(limit int) IRestconfCredentialDo {
	return r.withDO(r.DO.Limit(limit))
}

func (r restconfCredentialDo) Offset(offset int) IRestconfCredentialDo {
	return r.withDO(r.DO.Offset(offset))
}

func (r restconfCredentialDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IRestconfCredentialDo {
	return r.withDO(r.DO.Scopes(funcs...))
}

func (r restconfCredentialDo) Unscoped() IRestconfCredentialDo {
	return r.withDO(r.DO.Unscoped())
}

func (r restconfCredentialDo) Create(values ...*models.RestconfCredential) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Create(values)
}

func (r restconfCredentialDo) CreateInBatches(values []*models.RestconfCredential, batchSize int) error {
	return r.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (r restconfCredentialDo) Save(values ...*models.RestconfCredential) error {
	if len(values) == 0 {
		return nil
	}
	return r.DO.Save(values)
}

func (r restconfCredentialDo) First() (*models.RestconfCredential, error) {
	if result, err := r.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.RestconfCredential), nil
	}
}

func (r restconfCredentialDo) Take() (*models.RestconfCredential, error) {
	if result, err := r.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.RestconfCredential), nil
	}
}

func (r restconfCredentialDo) Last() (*models.RestconfCredential, error) {
	if result, err := r.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.RestconfCredential), nil
	}
}

func (r restconfCredentialDo) Find() ([]*models.RestconfCredential, error) {
	result, err := r.DO.Find()
	return result.([]*models.RestconfCredential), err
}

func (r restconfCredentialDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.RestconfCredential, err error) {
	buf := make([]*models.RestconfCredential, 0, batchSize)
	err = r.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (r restconfCredentialDo) FindInBatches(result *[]*models.RestconfCredential, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return r.DO.FindInBatches(result, batchSize, fc)
}

func (r restconfCredentialDo) Attrs(attrs ...field.AssignExpr) IRestconfCredentialDo {
	return r.withDO(r.DO.Attrs(attrs...))
}

func (r restconfCredentialDo) Assign(attrs ...field.AssignExpr) IRestconfCredentialDo {
	return r.withDO(r.DO.Assign(attrs...))
}

func (r restconfCredentialDo) Joins(fields ...field.RelationField) IRestconfCredentialDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Joins(_f))
	}
	return &r
}

func (r restconfCredentialDo) Preload(fields ...field.RelationField) IRestconfCredentialDo {
	for _, _f := range fields {
		r = *r.withDO(r.DO.Preload(_f))
	}
	return &r
}

func (r restconfCredentialDo) FirstOrInit() (*models.RestconfCredential, error) {
	if result, err := r.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.RestconfCredential), nil
	}
}

func (r restconfCredentialDo) FirstOrCreate() (*models.RestconfCredential, error) {
	if result, err := r.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.RestconfCredential), nil
	}
}

func (r restconfCredentialDo) FindByPage(offset int, limit int) (result []*models.RestconfCredential, count int64, err error) {
	result, err = r.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = r.Offset(-1).Limit(-1).Count()
	return
}

func (r restconfCredentialDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = r.Count()
	if err != nil {
		return
	}

	err = r.Offset(offset).Limit(limit).Scan(result)
	return
}

func (r restconfCredentialDo) Scan(result interface{}) (err error) {
	return r.DO.Scan(result)
}

func (r restconfCredentialDo) Delete(models ...*models.RestconfCredential) (result gen.ResultInfo, err error) {
	return r.DO.Delete(models)
}

func (r *restconfCredentialDo) withDO(do gen.Dao) *restconfCredentialDo {
	r.DO = *do.(*gen.DO)
	return r
}
