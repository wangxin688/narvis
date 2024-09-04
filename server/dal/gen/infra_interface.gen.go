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

func newDeviceInterface(db *gorm.DB, opts ...gen.DOOption) deviceInterface {
	_deviceInterface := deviceInterface{}

	_deviceInterface.deviceInterfaceDo.UseDB(db, opts...)
	_deviceInterface.deviceInterfaceDo.UseModel(&models.DeviceInterface{})

	tableName := _deviceInterface.deviceInterfaceDo.TableName()
	_deviceInterface.ALL = field.NewAsterisk(tableName)
	_deviceInterface.Id = field.NewString(tableName, "id")
	_deviceInterface.CreatedAt = field.NewTime(tableName, "createdAt")
	_deviceInterface.UpdatedAt = field.NewTime(tableName, "updatedAt")
	_deviceInterface.IfName = field.NewString(tableName, "ifName")
	_deviceInterface.IfIndex = field.NewUint64(tableName, "ifIndex")
	_deviceInterface.IfDescr = field.NewString(tableName, "ifDescr")
	_deviceInterface.IfSpeed = field.NewUint64(tableName, "ifSpeed")
	_deviceInterface.IfType = field.NewUint64(tableName, "ifType")
	_deviceInterface.IfMtu = field.NewUint64(tableName, "ifMtu")
	_deviceInterface.IfAdminStatus = field.NewUint64(tableName, "ifAdminStatus")
	_deviceInterface.IfOperStatus = field.NewUint64(tableName, "ifOperStatus")
	_deviceInterface.IfLastChange = field.NewUint64(tableName, "ifLastChange")
	_deviceInterface.IfHighSpeed = field.NewUint64(tableName, "ifHighSpeed")
	_deviceInterface.IfPhysAddr = field.NewString(tableName, "ifPhysAddr")
	_deviceInterface.IfIpAddress = field.NewString(tableName, "ifIpAddress")
	_deviceInterface.DeviceId = field.NewString(tableName, "deviceId")
	_deviceInterface.SiteId = field.NewString(tableName, "siteId")
	_deviceInterface.Device = deviceInterfaceBelongsToDevice{
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

	_deviceInterface.Site = deviceInterfaceBelongsToSite{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Site", "models.Site"),
	}

	_deviceInterface.fillFieldMap()

	return _deviceInterface
}

type deviceInterface struct {
	deviceInterfaceDo

	ALL           field.Asterisk
	Id            field.String
	CreatedAt     field.Time
	UpdatedAt     field.Time
	IfName        field.String
	IfIndex       field.Uint64
	IfDescr       field.String
	IfSpeed       field.Uint64
	IfType        field.Uint64
	IfMtu         field.Uint64
	IfAdminStatus field.Uint64
	IfOperStatus  field.Uint64
	IfLastChange  field.Uint64
	IfHighSpeed   field.Uint64
	IfPhysAddr    field.String
	IfIpAddress   field.String
	DeviceId      field.String
	SiteId        field.String
	Device        deviceInterfaceBelongsToDevice

	Site deviceInterfaceBelongsToSite

	fieldMap map[string]field.Expr
}

func (d deviceInterface) Table(newTableName string) *deviceInterface {
	d.deviceInterfaceDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d deviceInterface) As(alias string) *deviceInterface {
	d.deviceInterfaceDo.DO = *(d.deviceInterfaceDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *deviceInterface) updateTableName(table string) *deviceInterface {
	d.ALL = field.NewAsterisk(table)
	d.Id = field.NewString(table, "id")
	d.CreatedAt = field.NewTime(table, "createdAt")
	d.UpdatedAt = field.NewTime(table, "updatedAt")
	d.IfName = field.NewString(table, "ifName")
	d.IfIndex = field.NewUint64(table, "ifIndex")
	d.IfDescr = field.NewString(table, "ifDescr")
	d.IfSpeed = field.NewUint64(table, "ifSpeed")
	d.IfType = field.NewUint64(table, "ifType")
	d.IfMtu = field.NewUint64(table, "ifMtu")
	d.IfAdminStatus = field.NewUint64(table, "ifAdminStatus")
	d.IfOperStatus = field.NewUint64(table, "ifOperStatus")
	d.IfLastChange = field.NewUint64(table, "ifLastChange")
	d.IfHighSpeed = field.NewUint64(table, "ifHighSpeed")
	d.IfPhysAddr = field.NewString(table, "ifPhysAddr")
	d.IfIpAddress = field.NewString(table, "ifIpAddress")
	d.DeviceId = field.NewString(table, "deviceId")
	d.SiteId = field.NewString(table, "siteId")

	d.fillFieldMap()

	return d
}

func (d *deviceInterface) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *deviceInterface) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 19)
	d.fieldMap["id"] = d.Id
	d.fieldMap["createdAt"] = d.CreatedAt
	d.fieldMap["updatedAt"] = d.UpdatedAt
	d.fieldMap["ifName"] = d.IfName
	d.fieldMap["ifIndex"] = d.IfIndex
	d.fieldMap["ifDescr"] = d.IfDescr
	d.fieldMap["ifSpeed"] = d.IfSpeed
	d.fieldMap["ifType"] = d.IfType
	d.fieldMap["ifMtu"] = d.IfMtu
	d.fieldMap["ifAdminStatus"] = d.IfAdminStatus
	d.fieldMap["ifOperStatus"] = d.IfOperStatus
	d.fieldMap["ifLastChange"] = d.IfLastChange
	d.fieldMap["ifHighSpeed"] = d.IfHighSpeed
	d.fieldMap["ifPhysAddr"] = d.IfPhysAddr
	d.fieldMap["ifIpAddress"] = d.IfIpAddress
	d.fieldMap["deviceId"] = d.DeviceId
	d.fieldMap["siteId"] = d.SiteId

}

func (d deviceInterface) clone(db *gorm.DB) deviceInterface {
	d.deviceInterfaceDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d deviceInterface) replaceDB(db *gorm.DB) deviceInterface {
	d.deviceInterfaceDo.ReplaceDB(db)
	return d
}

type deviceInterfaceBelongsToDevice struct {
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

func (a deviceInterfaceBelongsToDevice) Where(conds ...field.Expr) *deviceInterfaceBelongsToDevice {
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

func (a deviceInterfaceBelongsToDevice) WithContext(ctx context.Context) *deviceInterfaceBelongsToDevice {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a deviceInterfaceBelongsToDevice) Session(session *gorm.Session) *deviceInterfaceBelongsToDevice {
	a.db = a.db.Session(session)
	return &a
}

func (a deviceInterfaceBelongsToDevice) Model(m *models.DeviceInterface) *deviceInterfaceBelongsToDeviceTx {
	return &deviceInterfaceBelongsToDeviceTx{a.db.Model(m).Association(a.Name())}
}

type deviceInterfaceBelongsToDeviceTx struct{ tx *gorm.Association }

func (a deviceInterfaceBelongsToDeviceTx) Find() (result *models.Device, err error) {
	return result, a.tx.Find(&result)
}

func (a deviceInterfaceBelongsToDeviceTx) Append(values ...*models.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a deviceInterfaceBelongsToDeviceTx) Replace(values ...*models.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a deviceInterfaceBelongsToDeviceTx) Delete(values ...*models.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a deviceInterfaceBelongsToDeviceTx) Clear() error {
	return a.tx.Clear()
}

func (a deviceInterfaceBelongsToDeviceTx) Count() int64 {
	return a.tx.Count()
}

type deviceInterfaceBelongsToSite struct {
	db *gorm.DB

	field.RelationField
}

func (a deviceInterfaceBelongsToSite) Where(conds ...field.Expr) *deviceInterfaceBelongsToSite {
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

func (a deviceInterfaceBelongsToSite) WithContext(ctx context.Context) *deviceInterfaceBelongsToSite {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a deviceInterfaceBelongsToSite) Session(session *gorm.Session) *deviceInterfaceBelongsToSite {
	a.db = a.db.Session(session)
	return &a
}

func (a deviceInterfaceBelongsToSite) Model(m *models.DeviceInterface) *deviceInterfaceBelongsToSiteTx {
	return &deviceInterfaceBelongsToSiteTx{a.db.Model(m).Association(a.Name())}
}

type deviceInterfaceBelongsToSiteTx struct{ tx *gorm.Association }

func (a deviceInterfaceBelongsToSiteTx) Find() (result *models.Site, err error) {
	return result, a.tx.Find(&result)
}

func (a deviceInterfaceBelongsToSiteTx) Append(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a deviceInterfaceBelongsToSiteTx) Replace(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a deviceInterfaceBelongsToSiteTx) Delete(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a deviceInterfaceBelongsToSiteTx) Clear() error {
	return a.tx.Clear()
}

func (a deviceInterfaceBelongsToSiteTx) Count() int64 {
	return a.tx.Count()
}

type deviceInterfaceDo struct{ gen.DO }

type IDeviceInterfaceDo interface {
	gen.SubQuery
	Debug() IDeviceInterfaceDo
	WithContext(ctx context.Context) IDeviceInterfaceDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDeviceInterfaceDo
	WriteDB() IDeviceInterfaceDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDeviceInterfaceDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDeviceInterfaceDo
	Not(conds ...gen.Condition) IDeviceInterfaceDo
	Or(conds ...gen.Condition) IDeviceInterfaceDo
	Select(conds ...field.Expr) IDeviceInterfaceDo
	Where(conds ...gen.Condition) IDeviceInterfaceDo
	Order(conds ...field.Expr) IDeviceInterfaceDo
	Distinct(cols ...field.Expr) IDeviceInterfaceDo
	Omit(cols ...field.Expr) IDeviceInterfaceDo
	Join(table schema.Tabler, on ...field.Expr) IDeviceInterfaceDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDeviceInterfaceDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDeviceInterfaceDo
	Group(cols ...field.Expr) IDeviceInterfaceDo
	Having(conds ...gen.Condition) IDeviceInterfaceDo
	Limit(limit int) IDeviceInterfaceDo
	Offset(offset int) IDeviceInterfaceDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDeviceInterfaceDo
	Unscoped() IDeviceInterfaceDo
	Create(values ...*models.DeviceInterface) error
	CreateInBatches(values []*models.DeviceInterface, batchSize int) error
	Save(values ...*models.DeviceInterface) error
	First() (*models.DeviceInterface, error)
	Take() (*models.DeviceInterface, error)
	Last() (*models.DeviceInterface, error)
	Find() ([]*models.DeviceInterface, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.DeviceInterface, err error)
	FindInBatches(result *[]*models.DeviceInterface, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.DeviceInterface) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDeviceInterfaceDo
	Assign(attrs ...field.AssignExpr) IDeviceInterfaceDo
	Joins(fields ...field.RelationField) IDeviceInterfaceDo
	Preload(fields ...field.RelationField) IDeviceInterfaceDo
	FirstOrInit() (*models.DeviceInterface, error)
	FirstOrCreate() (*models.DeviceInterface, error)
	FindByPage(offset int, limit int) (result []*models.DeviceInterface, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDeviceInterfaceDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d deviceInterfaceDo) Debug() IDeviceInterfaceDo {
	return d.withDO(d.DO.Debug())
}

func (d deviceInterfaceDo) WithContext(ctx context.Context) IDeviceInterfaceDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d deviceInterfaceDo) ReadDB() IDeviceInterfaceDo {
	return d.Clauses(dbresolver.Read)
}

func (d deviceInterfaceDo) WriteDB() IDeviceInterfaceDo {
	return d.Clauses(dbresolver.Write)
}

func (d deviceInterfaceDo) Session(config *gorm.Session) IDeviceInterfaceDo {
	return d.withDO(d.DO.Session(config))
}

func (d deviceInterfaceDo) Clauses(conds ...clause.Expression) IDeviceInterfaceDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d deviceInterfaceDo) Returning(value interface{}, columns ...string) IDeviceInterfaceDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d deviceInterfaceDo) Not(conds ...gen.Condition) IDeviceInterfaceDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d deviceInterfaceDo) Or(conds ...gen.Condition) IDeviceInterfaceDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d deviceInterfaceDo) Select(conds ...field.Expr) IDeviceInterfaceDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d deviceInterfaceDo) Where(conds ...gen.Condition) IDeviceInterfaceDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d deviceInterfaceDo) Order(conds ...field.Expr) IDeviceInterfaceDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d deviceInterfaceDo) Distinct(cols ...field.Expr) IDeviceInterfaceDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d deviceInterfaceDo) Omit(cols ...field.Expr) IDeviceInterfaceDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d deviceInterfaceDo) Join(table schema.Tabler, on ...field.Expr) IDeviceInterfaceDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d deviceInterfaceDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDeviceInterfaceDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d deviceInterfaceDo) RightJoin(table schema.Tabler, on ...field.Expr) IDeviceInterfaceDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d deviceInterfaceDo) Group(cols ...field.Expr) IDeviceInterfaceDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d deviceInterfaceDo) Having(conds ...gen.Condition) IDeviceInterfaceDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d deviceInterfaceDo) Limit(limit int) IDeviceInterfaceDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d deviceInterfaceDo) Offset(offset int) IDeviceInterfaceDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d deviceInterfaceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDeviceInterfaceDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d deviceInterfaceDo) Unscoped() IDeviceInterfaceDo {
	return d.withDO(d.DO.Unscoped())
}

func (d deviceInterfaceDo) Create(values ...*models.DeviceInterface) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d deviceInterfaceDo) CreateInBatches(values []*models.DeviceInterface, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d deviceInterfaceDo) Save(values ...*models.DeviceInterface) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d deviceInterfaceDo) First() (*models.DeviceInterface, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.DeviceInterface), nil
	}
}

func (d deviceInterfaceDo) Take() (*models.DeviceInterface, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.DeviceInterface), nil
	}
}

func (d deviceInterfaceDo) Last() (*models.DeviceInterface, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.DeviceInterface), nil
	}
}

func (d deviceInterfaceDo) Find() ([]*models.DeviceInterface, error) {
	result, err := d.DO.Find()
	return result.([]*models.DeviceInterface), err
}

func (d deviceInterfaceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.DeviceInterface, err error) {
	buf := make([]*models.DeviceInterface, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d deviceInterfaceDo) FindInBatches(result *[]*models.DeviceInterface, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d deviceInterfaceDo) Attrs(attrs ...field.AssignExpr) IDeviceInterfaceDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d deviceInterfaceDo) Assign(attrs ...field.AssignExpr) IDeviceInterfaceDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d deviceInterfaceDo) Joins(fields ...field.RelationField) IDeviceInterfaceDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d deviceInterfaceDo) Preload(fields ...field.RelationField) IDeviceInterfaceDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d deviceInterfaceDo) FirstOrInit() (*models.DeviceInterface, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.DeviceInterface), nil
	}
}

func (d deviceInterfaceDo) FirstOrCreate() (*models.DeviceInterface, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.DeviceInterface), nil
	}
}

func (d deviceInterfaceDo) FindByPage(offset int, limit int) (result []*models.DeviceInterface, count int64, err error) {
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

func (d deviceInterfaceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d deviceInterfaceDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d deviceInterfaceDo) Delete(models ...*models.DeviceInterface) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *deviceInterfaceDo) withDO(do gen.Dao) *deviceInterfaceDo {
	d.DO = *do.(*gen.DO)
	return d
}
