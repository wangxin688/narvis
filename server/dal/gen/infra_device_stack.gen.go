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

func newDeviceStack(db *gorm.DB, opts ...gen.DOOption) deviceStack {
	_deviceStack := deviceStack{}

	_deviceStack.deviceStackDo.UseDB(db, opts...)
	_deviceStack.deviceStackDo.UseModel(&models.DeviceStack{})

	tableName := _deviceStack.deviceStackDo.TableName()
	_deviceStack.ALL = field.NewAsterisk(tableName)
	_deviceStack.Id = field.NewString(tableName, "id")
	_deviceStack.CreatedAt = field.NewTime(tableName, "createdAt")
	_deviceStack.UpdatedAt = field.NewTime(tableName, "updatedAt")
	_deviceStack.Priority = field.NewUint8(tableName, "priority")
	_deviceStack.SerialNumber = field.NewString(tableName, "serialNumber")
	_deviceStack.MacAddress = field.NewString(tableName, "macAddress")
	_deviceStack.DeviceId = field.NewString(tableName, "deviceId")
	_deviceStack.Device = deviceStackBelongsToDevice{
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

	_deviceStack.fillFieldMap()

	return _deviceStack
}

type deviceStack struct {
	deviceStackDo

	ALL          field.Asterisk
	Id           field.String
	CreatedAt    field.Time
	UpdatedAt    field.Time
	Priority     field.Uint8
	SerialNumber field.String
	MacAddress   field.String
	DeviceId     field.String
	Device       deviceStackBelongsToDevice

	fieldMap map[string]field.Expr
}

func (d deviceStack) Table(newTableName string) *deviceStack {
	d.deviceStackDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d deviceStack) As(alias string) *deviceStack {
	d.deviceStackDo.DO = *(d.deviceStackDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *deviceStack) updateTableName(table string) *deviceStack {
	d.ALL = field.NewAsterisk(table)
	d.Id = field.NewString(table, "id")
	d.CreatedAt = field.NewTime(table, "createdAt")
	d.UpdatedAt = field.NewTime(table, "updatedAt")
	d.Priority = field.NewUint8(table, "priority")
	d.SerialNumber = field.NewString(table, "serialNumber")
	d.MacAddress = field.NewString(table, "macAddress")
	d.DeviceId = field.NewString(table, "deviceId")

	d.fillFieldMap()

	return d
}

func (d *deviceStack) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *deviceStack) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 8)
	d.fieldMap["id"] = d.Id
	d.fieldMap["createdAt"] = d.CreatedAt
	d.fieldMap["updatedAt"] = d.UpdatedAt
	d.fieldMap["priority"] = d.Priority
	d.fieldMap["serialNumber"] = d.SerialNumber
	d.fieldMap["macAddress"] = d.MacAddress
	d.fieldMap["deviceId"] = d.DeviceId

}

func (d deviceStack) clone(db *gorm.DB) deviceStack {
	d.deviceStackDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d deviceStack) replaceDB(db *gorm.DB) deviceStack {
	d.deviceStackDo.ReplaceDB(db)
	return d
}

type deviceStackBelongsToDevice struct {
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

func (a deviceStackBelongsToDevice) Where(conds ...field.Expr) *deviceStackBelongsToDevice {
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

func (a deviceStackBelongsToDevice) WithContext(ctx context.Context) *deviceStackBelongsToDevice {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a deviceStackBelongsToDevice) Session(session *gorm.Session) *deviceStackBelongsToDevice {
	a.db = a.db.Session(session)
	return &a
}

func (a deviceStackBelongsToDevice) Model(m *models.DeviceStack) *deviceStackBelongsToDeviceTx {
	return &deviceStackBelongsToDeviceTx{a.db.Model(m).Association(a.Name())}
}

type deviceStackBelongsToDeviceTx struct{ tx *gorm.Association }

func (a deviceStackBelongsToDeviceTx) Find() (result *models.Device, err error) {
	return result, a.tx.Find(&result)
}

func (a deviceStackBelongsToDeviceTx) Append(values ...*models.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a deviceStackBelongsToDeviceTx) Replace(values ...*models.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a deviceStackBelongsToDeviceTx) Delete(values ...*models.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a deviceStackBelongsToDeviceTx) Clear() error {
	return a.tx.Clear()
}

func (a deviceStackBelongsToDeviceTx) Count() int64 {
	return a.tx.Count()
}

type deviceStackDo struct{ gen.DO }

type IDeviceStackDo interface {
	gen.SubQuery
	Debug() IDeviceStackDo
	WithContext(ctx context.Context) IDeviceStackDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDeviceStackDo
	WriteDB() IDeviceStackDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDeviceStackDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDeviceStackDo
	Not(conds ...gen.Condition) IDeviceStackDo
	Or(conds ...gen.Condition) IDeviceStackDo
	Select(conds ...field.Expr) IDeviceStackDo
	Where(conds ...gen.Condition) IDeviceStackDo
	Order(conds ...field.Expr) IDeviceStackDo
	Distinct(cols ...field.Expr) IDeviceStackDo
	Omit(cols ...field.Expr) IDeviceStackDo
	Join(table schema.Tabler, on ...field.Expr) IDeviceStackDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDeviceStackDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDeviceStackDo
	Group(cols ...field.Expr) IDeviceStackDo
	Having(conds ...gen.Condition) IDeviceStackDo
	Limit(limit int) IDeviceStackDo
	Offset(offset int) IDeviceStackDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDeviceStackDo
	Unscoped() IDeviceStackDo
	Create(values ...*models.DeviceStack) error
	CreateInBatches(values []*models.DeviceStack, batchSize int) error
	Save(values ...*models.DeviceStack) error
	First() (*models.DeviceStack, error)
	Take() (*models.DeviceStack, error)
	Last() (*models.DeviceStack, error)
	Find() ([]*models.DeviceStack, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.DeviceStack, err error)
	FindInBatches(result *[]*models.DeviceStack, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.DeviceStack) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDeviceStackDo
	Assign(attrs ...field.AssignExpr) IDeviceStackDo
	Joins(fields ...field.RelationField) IDeviceStackDo
	Preload(fields ...field.RelationField) IDeviceStackDo
	FirstOrInit() (*models.DeviceStack, error)
	FirstOrCreate() (*models.DeviceStack, error)
	FindByPage(offset int, limit int) (result []*models.DeviceStack, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDeviceStackDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d deviceStackDo) Debug() IDeviceStackDo {
	return d.withDO(d.DO.Debug())
}

func (d deviceStackDo) WithContext(ctx context.Context) IDeviceStackDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d deviceStackDo) ReadDB() IDeviceStackDo {
	return d.Clauses(dbresolver.Read)
}

func (d deviceStackDo) WriteDB() IDeviceStackDo {
	return d.Clauses(dbresolver.Write)
}

func (d deviceStackDo) Session(config *gorm.Session) IDeviceStackDo {
	return d.withDO(d.DO.Session(config))
}

func (d deviceStackDo) Clauses(conds ...clause.Expression) IDeviceStackDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d deviceStackDo) Returning(value interface{}, columns ...string) IDeviceStackDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d deviceStackDo) Not(conds ...gen.Condition) IDeviceStackDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d deviceStackDo) Or(conds ...gen.Condition) IDeviceStackDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d deviceStackDo) Select(conds ...field.Expr) IDeviceStackDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d deviceStackDo) Where(conds ...gen.Condition) IDeviceStackDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d deviceStackDo) Order(conds ...field.Expr) IDeviceStackDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d deviceStackDo) Distinct(cols ...field.Expr) IDeviceStackDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d deviceStackDo) Omit(cols ...field.Expr) IDeviceStackDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d deviceStackDo) Join(table schema.Tabler, on ...field.Expr) IDeviceStackDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d deviceStackDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDeviceStackDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d deviceStackDo) RightJoin(table schema.Tabler, on ...field.Expr) IDeviceStackDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d deviceStackDo) Group(cols ...field.Expr) IDeviceStackDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d deviceStackDo) Having(conds ...gen.Condition) IDeviceStackDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d deviceStackDo) Limit(limit int) IDeviceStackDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d deviceStackDo) Offset(offset int) IDeviceStackDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d deviceStackDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDeviceStackDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d deviceStackDo) Unscoped() IDeviceStackDo {
	return d.withDO(d.DO.Unscoped())
}

func (d deviceStackDo) Create(values ...*models.DeviceStack) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d deviceStackDo) CreateInBatches(values []*models.DeviceStack, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d deviceStackDo) Save(values ...*models.DeviceStack) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d deviceStackDo) First() (*models.DeviceStack, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.DeviceStack), nil
	}
}

func (d deviceStackDo) Take() (*models.DeviceStack, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.DeviceStack), nil
	}
}

func (d deviceStackDo) Last() (*models.DeviceStack, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.DeviceStack), nil
	}
}

func (d deviceStackDo) Find() ([]*models.DeviceStack, error) {
	result, err := d.DO.Find()
	return result.([]*models.DeviceStack), err
}

func (d deviceStackDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.DeviceStack, err error) {
	buf := make([]*models.DeviceStack, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d deviceStackDo) FindInBatches(result *[]*models.DeviceStack, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d deviceStackDo) Attrs(attrs ...field.AssignExpr) IDeviceStackDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d deviceStackDo) Assign(attrs ...field.AssignExpr) IDeviceStackDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d deviceStackDo) Joins(fields ...field.RelationField) IDeviceStackDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d deviceStackDo) Preload(fields ...field.RelationField) IDeviceStackDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d deviceStackDo) FirstOrInit() (*models.DeviceStack, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.DeviceStack), nil
	}
}

func (d deviceStackDo) FirstOrCreate() (*models.DeviceStack, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.DeviceStack), nil
	}
}

func (d deviceStackDo) FindByPage(offset int, limit int) (result []*models.DeviceStack, count int64, err error) {
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

func (d deviceStackDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d deviceStackDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d deviceStackDo) Delete(models ...*models.DeviceStack) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *deviceStackDo) withDO(do gen.Dao) *deviceStackDo {
	d.DO = *do.(*gen.DO)
	return d
}
