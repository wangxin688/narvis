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

func newSyslog(db *gorm.DB, opts ...gen.DOOption) syslog {
	_syslog := syslog{}

	_syslog.syslogDo.UseDB(db, opts...)
	_syslog.syslogDo.UseModel(&models.Syslog{})

	tableName := _syslog.syslogDo.TableName()
	_syslog.ALL = field.NewAsterisk(tableName)
	_syslog.Time = field.NewTime(tableName, "time")
	_syslog.Facility = field.NewString(tableName, "facility")
	_syslog.Severity = field.NewString(tableName, "severity")
	_syslog.Message = field.NewString(tableName, "message")
	_syslog.DeviceId = field.NewString(tableName, "deviceId")
	_syslog.SiteId = field.NewString(tableName, "siteId")
	_syslog.OrganizationId = field.NewString(tableName, "organizationId")

	_syslog.fillFieldMap()

	return _syslog
}

type syslog struct {
	syslogDo

	ALL            field.Asterisk
	Time           field.Time
	Facility       field.String
	Severity       field.String
	Message        field.String
	DeviceId       field.String
	SiteId         field.String
	OrganizationId field.String

	fieldMap map[string]field.Expr
}

func (s syslog) Table(newTableName string) *syslog {
	s.syslogDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s syslog) As(alias string) *syslog {
	s.syslogDo.DO = *(s.syslogDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *syslog) updateTableName(table string) *syslog {
	s.ALL = field.NewAsterisk(table)
	s.Time = field.NewTime(table, "time")
	s.Facility = field.NewString(table, "facility")
	s.Severity = field.NewString(table, "severity")
	s.Message = field.NewString(table, "message")
	s.DeviceId = field.NewString(table, "deviceId")
	s.SiteId = field.NewString(table, "siteId")
	s.OrganizationId = field.NewString(table, "organizationId")

	s.fillFieldMap()

	return s
}

func (s *syslog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *syslog) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 7)
	s.fieldMap["time"] = s.Time
	s.fieldMap["facility"] = s.Facility
	s.fieldMap["severity"] = s.Severity
	s.fieldMap["message"] = s.Message
	s.fieldMap["deviceId"] = s.DeviceId
	s.fieldMap["siteId"] = s.SiteId
	s.fieldMap["organizationId"] = s.OrganizationId
}

func (s syslog) clone(db *gorm.DB) syslog {
	s.syslogDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s syslog) replaceDB(db *gorm.DB) syslog {
	s.syslogDo.ReplaceDB(db)
	return s
}

type syslogDo struct{ gen.DO }

type ISyslogDo interface {
	gen.SubQuery
	Debug() ISyslogDo
	WithContext(ctx context.Context) ISyslogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ISyslogDo
	WriteDB() ISyslogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ISyslogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ISyslogDo
	Not(conds ...gen.Condition) ISyslogDo
	Or(conds ...gen.Condition) ISyslogDo
	Select(conds ...field.Expr) ISyslogDo
	Where(conds ...gen.Condition) ISyslogDo
	Order(conds ...field.Expr) ISyslogDo
	Distinct(cols ...field.Expr) ISyslogDo
	Omit(cols ...field.Expr) ISyslogDo
	Join(table schema.Tabler, on ...field.Expr) ISyslogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ISyslogDo
	RightJoin(table schema.Tabler, on ...field.Expr) ISyslogDo
	Group(cols ...field.Expr) ISyslogDo
	Having(conds ...gen.Condition) ISyslogDo
	Limit(limit int) ISyslogDo
	Offset(offset int) ISyslogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ISyslogDo
	Unscoped() ISyslogDo
	Create(values ...*models.Syslog) error
	CreateInBatches(values []*models.Syslog, batchSize int) error
	Save(values ...*models.Syslog) error
	First() (*models.Syslog, error)
	Take() (*models.Syslog, error)
	Last() (*models.Syslog, error)
	Find() ([]*models.Syslog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Syslog, err error)
	FindInBatches(result *[]*models.Syslog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Syslog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ISyslogDo
	Assign(attrs ...field.AssignExpr) ISyslogDo
	Joins(fields ...field.RelationField) ISyslogDo
	Preload(fields ...field.RelationField) ISyslogDo
	FirstOrInit() (*models.Syslog, error)
	FirstOrCreate() (*models.Syslog, error)
	FindByPage(offset int, limit int) (result []*models.Syslog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ISyslogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s syslogDo) Debug() ISyslogDo {
	return s.withDO(s.DO.Debug())
}

func (s syslogDo) WithContext(ctx context.Context) ISyslogDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s syslogDo) ReadDB() ISyslogDo {
	return s.Clauses(dbresolver.Read)
}

func (s syslogDo) WriteDB() ISyslogDo {
	return s.Clauses(dbresolver.Write)
}

func (s syslogDo) Session(config *gorm.Session) ISyslogDo {
	return s.withDO(s.DO.Session(config))
}

func (s syslogDo) Clauses(conds ...clause.Expression) ISyslogDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s syslogDo) Returning(value interface{}, columns ...string) ISyslogDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s syslogDo) Not(conds ...gen.Condition) ISyslogDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s syslogDo) Or(conds ...gen.Condition) ISyslogDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s syslogDo) Select(conds ...field.Expr) ISyslogDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s syslogDo) Where(conds ...gen.Condition) ISyslogDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s syslogDo) Order(conds ...field.Expr) ISyslogDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s syslogDo) Distinct(cols ...field.Expr) ISyslogDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s syslogDo) Omit(cols ...field.Expr) ISyslogDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s syslogDo) Join(table schema.Tabler, on ...field.Expr) ISyslogDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s syslogDo) LeftJoin(table schema.Tabler, on ...field.Expr) ISyslogDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s syslogDo) RightJoin(table schema.Tabler, on ...field.Expr) ISyslogDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s syslogDo) Group(cols ...field.Expr) ISyslogDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s syslogDo) Having(conds ...gen.Condition) ISyslogDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s syslogDo) Limit(limit int) ISyslogDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s syslogDo) Offset(offset int) ISyslogDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s syslogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ISyslogDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s syslogDo) Unscoped() ISyslogDo {
	return s.withDO(s.DO.Unscoped())
}

func (s syslogDo) Create(values ...*models.Syslog) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s syslogDo) CreateInBatches(values []*models.Syslog, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s syslogDo) Save(values ...*models.Syslog) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s syslogDo) First() (*models.Syslog, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Syslog), nil
	}
}

func (s syslogDo) Take() (*models.Syslog, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Syslog), nil
	}
}

func (s syslogDo) Last() (*models.Syslog, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Syslog), nil
	}
}

func (s syslogDo) Find() ([]*models.Syslog, error) {
	result, err := s.DO.Find()
	return result.([]*models.Syslog), err
}

func (s syslogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Syslog, err error) {
	buf := make([]*models.Syslog, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s syslogDo) FindInBatches(result *[]*models.Syslog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s syslogDo) Attrs(attrs ...field.AssignExpr) ISyslogDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s syslogDo) Assign(attrs ...field.AssignExpr) ISyslogDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s syslogDo) Joins(fields ...field.RelationField) ISyslogDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s syslogDo) Preload(fields ...field.RelationField) ISyslogDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s syslogDo) FirstOrInit() (*models.Syslog, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Syslog), nil
	}
}

func (s syslogDo) FirstOrCreate() (*models.Syslog, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Syslog), nil
	}
}

func (s syslogDo) FindByPage(offset int, limit int) (result []*models.Syslog, count int64, err error) {
	result, err = s.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = s.Offset(-1).Limit(-1).Count()
	return
}

func (s syslogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s syslogDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s syslogDo) Delete(models ...*models.Syslog) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *syslogDo) withDO(do gen.Dao) *syslogDo {
	s.DO = *do.(*gen.DO)
	return s
}
