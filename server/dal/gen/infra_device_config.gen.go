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

func newDeviceConfig(db *gorm.DB, opts ...gen.DOOption) deviceConfig {
	_deviceConfig := deviceConfig{}

	_deviceConfig.deviceConfigDo.UseDB(db, opts...)
	_deviceConfig.deviceConfigDo.UseModel(&models.DeviceConfig{})

	tableName := _deviceConfig.deviceConfigDo.TableName()
	_deviceConfig.ALL = field.NewAsterisk(tableName)
	_deviceConfig.Id = field.NewString(tableName, "id")
	_deviceConfig.CreatedAt = field.NewTime(tableName, "createdAt")
	_deviceConfig.Configuration = field.NewString(tableName, "configuration")
	_deviceConfig.TotalLines = field.NewUint32(tableName, "totalLines")
	_deviceConfig.LinesAdded = field.NewUint32(tableName, "linesAdded")
	_deviceConfig.LinesDeleted = field.NewUint32(tableName, "linesDeleted")
	_deviceConfig.Md5Checksum = field.NewString(tableName, "md5Checksum")
	_deviceConfig.DeviceId = field.NewString(tableName, "deviceId")
	_deviceConfig.Device = deviceConfigBelongsToDevice{
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

	_deviceConfig.fillFieldMap()

	return _deviceConfig
}

type deviceConfig struct {
	deviceConfigDo

	ALL           field.Asterisk
	Id            field.String
	CreatedAt     field.Time
	Configuration field.String
	TotalLines    field.Uint32
	LinesAdded    field.Uint32
	LinesDeleted  field.Uint32
	Md5Checksum   field.String
	DeviceId      field.String
	Device        deviceConfigBelongsToDevice

	fieldMap map[string]field.Expr
}

func (d deviceConfig) Table(newTableName string) *deviceConfig {
	d.deviceConfigDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d deviceConfig) As(alias string) *deviceConfig {
	d.deviceConfigDo.DO = *(d.deviceConfigDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *deviceConfig) updateTableName(table string) *deviceConfig {
	d.ALL = field.NewAsterisk(table)
	d.Id = field.NewString(table, "id")
	d.CreatedAt = field.NewTime(table, "createdAt")
	d.Configuration = field.NewString(table, "configuration")
	d.TotalLines = field.NewUint32(table, "totalLines")
	d.LinesAdded = field.NewUint32(table, "linesAdded")
	d.LinesDeleted = field.NewUint32(table, "linesDeleted")
	d.Md5Checksum = field.NewString(table, "md5Checksum")
	d.DeviceId = field.NewString(table, "deviceId")

	d.fillFieldMap()

	return d
}

func (d *deviceConfig) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *deviceConfig) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 9)
	d.fieldMap["id"] = d.Id
	d.fieldMap["createdAt"] = d.CreatedAt
	d.fieldMap["configuration"] = d.Configuration
	d.fieldMap["totalLines"] = d.TotalLines
	d.fieldMap["linesAdded"] = d.LinesAdded
	d.fieldMap["linesDeleted"] = d.LinesDeleted
	d.fieldMap["md5Checksum"] = d.Md5Checksum
	d.fieldMap["deviceId"] = d.DeviceId

}

func (d deviceConfig) clone(db *gorm.DB) deviceConfig {
	d.deviceConfigDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d deviceConfig) replaceDB(db *gorm.DB) deviceConfig {
	d.deviceConfigDo.ReplaceDB(db)
	return d
}

type deviceConfigBelongsToDevice struct {
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

func (a deviceConfigBelongsToDevice) Where(conds ...field.Expr) *deviceConfigBelongsToDevice {
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

func (a deviceConfigBelongsToDevice) WithContext(ctx context.Context) *deviceConfigBelongsToDevice {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a deviceConfigBelongsToDevice) Session(session *gorm.Session) *deviceConfigBelongsToDevice {
	a.db = a.db.Session(session)
	return &a
}

func (a deviceConfigBelongsToDevice) Model(m *models.DeviceConfig) *deviceConfigBelongsToDeviceTx {
	return &deviceConfigBelongsToDeviceTx{a.db.Model(m).Association(a.Name())}
}

type deviceConfigBelongsToDeviceTx struct{ tx *gorm.Association }

func (a deviceConfigBelongsToDeviceTx) Find() (result *models.Device, err error) {
	return result, a.tx.Find(&result)
}

func (a deviceConfigBelongsToDeviceTx) Append(values ...*models.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a deviceConfigBelongsToDeviceTx) Replace(values ...*models.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a deviceConfigBelongsToDeviceTx) Delete(values ...*models.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a deviceConfigBelongsToDeviceTx) Clear() error {
	return a.tx.Clear()
}

func (a deviceConfigBelongsToDeviceTx) Count() int64 {
	return a.tx.Count()
}

type deviceConfigDo struct{ gen.DO }

type IDeviceConfigDo interface {
	gen.SubQuery
	Debug() IDeviceConfigDo
	WithContext(ctx context.Context) IDeviceConfigDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDeviceConfigDo
	WriteDB() IDeviceConfigDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDeviceConfigDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDeviceConfigDo
	Not(conds ...gen.Condition) IDeviceConfigDo
	Or(conds ...gen.Condition) IDeviceConfigDo
	Select(conds ...field.Expr) IDeviceConfigDo
	Where(conds ...gen.Condition) IDeviceConfigDo
	Order(conds ...field.Expr) IDeviceConfigDo
	Distinct(cols ...field.Expr) IDeviceConfigDo
	Omit(cols ...field.Expr) IDeviceConfigDo
	Join(table schema.Tabler, on ...field.Expr) IDeviceConfigDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDeviceConfigDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDeviceConfigDo
	Group(cols ...field.Expr) IDeviceConfigDo
	Having(conds ...gen.Condition) IDeviceConfigDo
	Limit(limit int) IDeviceConfigDo
	Offset(offset int) IDeviceConfigDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDeviceConfigDo
	Unscoped() IDeviceConfigDo
	Create(values ...*models.DeviceConfig) error
	CreateInBatches(values []*models.DeviceConfig, batchSize int) error
	Save(values ...*models.DeviceConfig) error
	First() (*models.DeviceConfig, error)
	Take() (*models.DeviceConfig, error)
	Last() (*models.DeviceConfig, error)
	Find() ([]*models.DeviceConfig, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.DeviceConfig, err error)
	FindInBatches(result *[]*models.DeviceConfig, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.DeviceConfig) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDeviceConfigDo
	Assign(attrs ...field.AssignExpr) IDeviceConfigDo
	Joins(fields ...field.RelationField) IDeviceConfigDo
	Preload(fields ...field.RelationField) IDeviceConfigDo
	FirstOrInit() (*models.DeviceConfig, error)
	FirstOrCreate() (*models.DeviceConfig, error)
	FindByPage(offset int, limit int) (result []*models.DeviceConfig, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDeviceConfigDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d deviceConfigDo) Debug() IDeviceConfigDo {
	return d.withDO(d.DO.Debug())
}

func (d deviceConfigDo) WithContext(ctx context.Context) IDeviceConfigDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d deviceConfigDo) ReadDB() IDeviceConfigDo {
	return d.Clauses(dbresolver.Read)
}

func (d deviceConfigDo) WriteDB() IDeviceConfigDo {
	return d.Clauses(dbresolver.Write)
}

func (d deviceConfigDo) Session(config *gorm.Session) IDeviceConfigDo {
	return d.withDO(d.DO.Session(config))
}

func (d deviceConfigDo) Clauses(conds ...clause.Expression) IDeviceConfigDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d deviceConfigDo) Returning(value interface{}, columns ...string) IDeviceConfigDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d deviceConfigDo) Not(conds ...gen.Condition) IDeviceConfigDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d deviceConfigDo) Or(conds ...gen.Condition) IDeviceConfigDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d deviceConfigDo) Select(conds ...field.Expr) IDeviceConfigDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d deviceConfigDo) Where(conds ...gen.Condition) IDeviceConfigDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d deviceConfigDo) Order(conds ...field.Expr) IDeviceConfigDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d deviceConfigDo) Distinct(cols ...field.Expr) IDeviceConfigDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d deviceConfigDo) Omit(cols ...field.Expr) IDeviceConfigDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d deviceConfigDo) Join(table schema.Tabler, on ...field.Expr) IDeviceConfigDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d deviceConfigDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDeviceConfigDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d deviceConfigDo) RightJoin(table schema.Tabler, on ...field.Expr) IDeviceConfigDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d deviceConfigDo) Group(cols ...field.Expr) IDeviceConfigDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d deviceConfigDo) Having(conds ...gen.Condition) IDeviceConfigDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d deviceConfigDo) Limit(limit int) IDeviceConfigDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d deviceConfigDo) Offset(offset int) IDeviceConfigDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d deviceConfigDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDeviceConfigDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d deviceConfigDo) Unscoped() IDeviceConfigDo {
	return d.withDO(d.DO.Unscoped())
}

func (d deviceConfigDo) Create(values ...*models.DeviceConfig) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d deviceConfigDo) CreateInBatches(values []*models.DeviceConfig, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d deviceConfigDo) Save(values ...*models.DeviceConfig) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d deviceConfigDo) First() (*models.DeviceConfig, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.DeviceConfig), nil
	}
}

func (d deviceConfigDo) Take() (*models.DeviceConfig, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.DeviceConfig), nil
	}
}

func (d deviceConfigDo) Last() (*models.DeviceConfig, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.DeviceConfig), nil
	}
}

func (d deviceConfigDo) Find() ([]*models.DeviceConfig, error) {
	result, err := d.DO.Find()
	return result.([]*models.DeviceConfig), err
}

func (d deviceConfigDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.DeviceConfig, err error) {
	buf := make([]*models.DeviceConfig, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d deviceConfigDo) FindInBatches(result *[]*models.DeviceConfig, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d deviceConfigDo) Attrs(attrs ...field.AssignExpr) IDeviceConfigDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d deviceConfigDo) Assign(attrs ...field.AssignExpr) IDeviceConfigDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d deviceConfigDo) Joins(fields ...field.RelationField) IDeviceConfigDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d deviceConfigDo) Preload(fields ...field.RelationField) IDeviceConfigDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d deviceConfigDo) FirstOrInit() (*models.DeviceConfig, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.DeviceConfig), nil
	}
}

func (d deviceConfigDo) FirstOrCreate() (*models.DeviceConfig, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.DeviceConfig), nil
	}
}

func (d deviceConfigDo) FindByPage(offset int, limit int) (result []*models.DeviceConfig, count int64, err error) {
	result, err = d.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = d.Offset(-1).Limit(-1).Count()
	return
}

func (d deviceConfigDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d deviceConfigDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d deviceConfigDo) Delete(models ...*models.DeviceConfig) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *deviceConfigDo) withDO(do gen.Dao) *deviceConfigDo {
	d.DO = *do.(*gen.DO)
	return d
}
