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

func newCircuit(db *gorm.DB, opts ...gen.DOOption) circuit {
	_circuit := circuit{}

	_circuit.circuitDo.UseDB(db, opts...)
	_circuit.circuitDo.UseModel(&models.Circuit{})

	tableName := _circuit.circuitDo.TableName()
	_circuit.ALL = field.NewAsterisk(tableName)
	_circuit.Id = field.NewString(tableName, "id")
	_circuit.CreatedAt = field.NewTime(tableName, "createdAt")
	_circuit.UpdatedAt = field.NewTime(tableName, "updatedAt")
	_circuit.Name = field.NewString(tableName, "name")
	_circuit.CId = field.NewString(tableName, "cId")
	_circuit.Status = field.NewString(tableName, "status")
	_circuit.CircuitType = field.NewString(tableName, "circuitType")
	_circuit.RxBandWidth = field.NewUint32(tableName, "rxBandWidth")
	_circuit.TxBandWidth = field.NewUint32(tableName, "txBandWidth")
	_circuit.Ipv4Address = field.NewString(tableName, "ipv4Address")
	_circuit.Ipv6Address = field.NewString(tableName, "ipv6Address")
	_circuit.Description = field.NewString(tableName, "description")
	_circuit.Provider = field.NewString(tableName, "provider")
	_circuit.SiteId = field.NewString(tableName, "siteId")
	_circuit.DeviceId = field.NewString(tableName, "deviceId")
	_circuit.InterfaceId = field.NewString(tableName, "interfaceId")
	_circuit.MonitorId = field.NewString(tableName, "monitorId")
	_circuit.OrganizationId = field.NewString(tableName, "organizationId")
	_circuit.Site = circuitBelongsToSite{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Site", "models.Site"),
		Organization: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Site.Organization", "models.Organization"),
		},
	}

	_circuit.Device = circuitBelongsToDevice{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Device", "models.Device"),
		Rack: struct {
			field.RelationField
			Site struct {
				field.RelationField
			}
			Organization struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Device.Rack", "models.Rack"),
			Site: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Device.Rack.Site", "models.Site"),
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

	_circuit.DeviceInterface = circuitBelongsToDeviceInterface{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("DeviceInterface", "models.DeviceInterface"),
		Device: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("DeviceInterface.Device", "models.Device"),
		},
		Site: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("DeviceInterface.Site", "models.Site"),
		},
	}

	_circuit.Organization = circuitBelongsToOrganization{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Organization", "models.Organization"),
	}

	_circuit.fillFieldMap()

	return _circuit
}

type circuit struct {
	circuitDo

	ALL            field.Asterisk
	Id             field.String
	CreatedAt      field.Time
	UpdatedAt      field.Time
	Name           field.String
	CId            field.String
	Status         field.String
	CircuitType    field.String
	RxBandWidth    field.Uint32
	TxBandWidth    field.Uint32
	Ipv4Address    field.String
	Ipv6Address    field.String
	Description    field.String
	Provider       field.String
	SiteId         field.String
	DeviceId       field.String
	InterfaceId    field.String
	MonitorId      field.String
	OrganizationId field.String
	Site           circuitBelongsToSite

	Device circuitBelongsToDevice

	DeviceInterface circuitBelongsToDeviceInterface

	Organization circuitBelongsToOrganization

	fieldMap map[string]field.Expr
}

func (c circuit) Table(newTableName string) *circuit {
	c.circuitDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c circuit) As(alias string) *circuit {
	c.circuitDo.DO = *(c.circuitDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *circuit) updateTableName(table string) *circuit {
	c.ALL = field.NewAsterisk(table)
	c.Id = field.NewString(table, "id")
	c.CreatedAt = field.NewTime(table, "createdAt")
	c.UpdatedAt = field.NewTime(table, "updatedAt")
	c.Name = field.NewString(table, "name")
	c.CId = field.NewString(table, "cId")
	c.Status = field.NewString(table, "status")
	c.CircuitType = field.NewString(table, "circuitType")
	c.RxBandWidth = field.NewUint32(table, "rxBandWidth")
	c.TxBandWidth = field.NewUint32(table, "txBandWidth")
	c.Ipv4Address = field.NewString(table, "ipv4Address")
	c.Ipv6Address = field.NewString(table, "ipv6Address")
	c.Description = field.NewString(table, "description")
	c.Provider = field.NewString(table, "provider")
	c.SiteId = field.NewString(table, "siteId")
	c.DeviceId = field.NewString(table, "deviceId")
	c.InterfaceId = field.NewString(table, "interfaceId")
	c.MonitorId = field.NewString(table, "monitorId")
	c.OrganizationId = field.NewString(table, "organizationId")

	c.fillFieldMap()

	return c
}

func (c *circuit) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *circuit) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 22)
	c.fieldMap["id"] = c.Id
	c.fieldMap["createdAt"] = c.CreatedAt
	c.fieldMap["updatedAt"] = c.UpdatedAt
	c.fieldMap["name"] = c.Name
	c.fieldMap["cId"] = c.CId
	c.fieldMap["status"] = c.Status
	c.fieldMap["circuitType"] = c.CircuitType
	c.fieldMap["rxBandWidth"] = c.RxBandWidth
	c.fieldMap["txBandWidth"] = c.TxBandWidth
	c.fieldMap["ipv4Address"] = c.Ipv4Address
	c.fieldMap["ipv6Address"] = c.Ipv6Address
	c.fieldMap["description"] = c.Description
	c.fieldMap["provider"] = c.Provider
	c.fieldMap["siteId"] = c.SiteId
	c.fieldMap["deviceId"] = c.DeviceId
	c.fieldMap["interfaceId"] = c.InterfaceId
	c.fieldMap["monitorId"] = c.MonitorId
	c.fieldMap["organizationId"] = c.OrganizationId

}

func (c circuit) clone(db *gorm.DB) circuit {
	c.circuitDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c circuit) replaceDB(db *gorm.DB) circuit {
	c.circuitDo.ReplaceDB(db)
	return c
}

type circuitBelongsToSite struct {
	db *gorm.DB

	field.RelationField

	Organization struct {
		field.RelationField
	}
}

func (a circuitBelongsToSite) Where(conds ...field.Expr) *circuitBelongsToSite {
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

func (a circuitBelongsToSite) WithContext(ctx context.Context) *circuitBelongsToSite {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a circuitBelongsToSite) Session(session *gorm.Session) *circuitBelongsToSite {
	a.db = a.db.Session(session)
	return &a
}

func (a circuitBelongsToSite) Model(m *models.Circuit) *circuitBelongsToSiteTx {
	return &circuitBelongsToSiteTx{a.db.Model(m).Association(a.Name())}
}

type circuitBelongsToSiteTx struct{ tx *gorm.Association }

func (a circuitBelongsToSiteTx) Find() (result *models.Site, err error) {
	return result, a.tx.Find(&result)
}

func (a circuitBelongsToSiteTx) Append(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a circuitBelongsToSiteTx) Replace(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a circuitBelongsToSiteTx) Delete(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a circuitBelongsToSiteTx) Clear() error {
	return a.tx.Clear()
}

func (a circuitBelongsToSiteTx) Count() int64 {
	return a.tx.Count()
}

type circuitBelongsToDevice struct {
	db *gorm.DB

	field.RelationField

	Rack struct {
		field.RelationField
		Site struct {
			field.RelationField
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

func (a circuitBelongsToDevice) Where(conds ...field.Expr) *circuitBelongsToDevice {
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

func (a circuitBelongsToDevice) WithContext(ctx context.Context) *circuitBelongsToDevice {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a circuitBelongsToDevice) Session(session *gorm.Session) *circuitBelongsToDevice {
	a.db = a.db.Session(session)
	return &a
}

func (a circuitBelongsToDevice) Model(m *models.Circuit) *circuitBelongsToDeviceTx {
	return &circuitBelongsToDeviceTx{a.db.Model(m).Association(a.Name())}
}

type circuitBelongsToDeviceTx struct{ tx *gorm.Association }

func (a circuitBelongsToDeviceTx) Find() (result *models.Device, err error) {
	return result, a.tx.Find(&result)
}

func (a circuitBelongsToDeviceTx) Append(values ...*models.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a circuitBelongsToDeviceTx) Replace(values ...*models.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a circuitBelongsToDeviceTx) Delete(values ...*models.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a circuitBelongsToDeviceTx) Clear() error {
	return a.tx.Clear()
}

func (a circuitBelongsToDeviceTx) Count() int64 {
	return a.tx.Count()
}

type circuitBelongsToDeviceInterface struct {
	db *gorm.DB

	field.RelationField

	Device struct {
		field.RelationField
	}
	Site struct {
		field.RelationField
	}
}

func (a circuitBelongsToDeviceInterface) Where(conds ...field.Expr) *circuitBelongsToDeviceInterface {
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

func (a circuitBelongsToDeviceInterface) WithContext(ctx context.Context) *circuitBelongsToDeviceInterface {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a circuitBelongsToDeviceInterface) Session(session *gorm.Session) *circuitBelongsToDeviceInterface {
	a.db = a.db.Session(session)
	return &a
}

func (a circuitBelongsToDeviceInterface) Model(m *models.Circuit) *circuitBelongsToDeviceInterfaceTx {
	return &circuitBelongsToDeviceInterfaceTx{a.db.Model(m).Association(a.Name())}
}

type circuitBelongsToDeviceInterfaceTx struct{ tx *gorm.Association }

func (a circuitBelongsToDeviceInterfaceTx) Find() (result *models.DeviceInterface, err error) {
	return result, a.tx.Find(&result)
}

func (a circuitBelongsToDeviceInterfaceTx) Append(values ...*models.DeviceInterface) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a circuitBelongsToDeviceInterfaceTx) Replace(values ...*models.DeviceInterface) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a circuitBelongsToDeviceInterfaceTx) Delete(values ...*models.DeviceInterface) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a circuitBelongsToDeviceInterfaceTx) Clear() error {
	return a.tx.Clear()
}

func (a circuitBelongsToDeviceInterfaceTx) Count() int64 {
	return a.tx.Count()
}

type circuitBelongsToOrganization struct {
	db *gorm.DB

	field.RelationField
}

func (a circuitBelongsToOrganization) Where(conds ...field.Expr) *circuitBelongsToOrganization {
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

func (a circuitBelongsToOrganization) WithContext(ctx context.Context) *circuitBelongsToOrganization {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a circuitBelongsToOrganization) Session(session *gorm.Session) *circuitBelongsToOrganization {
	a.db = a.db.Session(session)
	return &a
}

func (a circuitBelongsToOrganization) Model(m *models.Circuit) *circuitBelongsToOrganizationTx {
	return &circuitBelongsToOrganizationTx{a.db.Model(m).Association(a.Name())}
}

type circuitBelongsToOrganizationTx struct{ tx *gorm.Association }

func (a circuitBelongsToOrganizationTx) Find() (result *models.Organization, err error) {
	return result, a.tx.Find(&result)
}

func (a circuitBelongsToOrganizationTx) Append(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a circuitBelongsToOrganizationTx) Replace(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a circuitBelongsToOrganizationTx) Delete(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a circuitBelongsToOrganizationTx) Clear() error {
	return a.tx.Clear()
}

func (a circuitBelongsToOrganizationTx) Count() int64 {
	return a.tx.Count()
}

type circuitDo struct{ gen.DO }

type ICircuitDo interface {
	gen.SubQuery
	Debug() ICircuitDo
	WithContext(ctx context.Context) ICircuitDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICircuitDo
	WriteDB() ICircuitDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICircuitDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICircuitDo
	Not(conds ...gen.Condition) ICircuitDo
	Or(conds ...gen.Condition) ICircuitDo
	Select(conds ...field.Expr) ICircuitDo
	Where(conds ...gen.Condition) ICircuitDo
	Order(conds ...field.Expr) ICircuitDo
	Distinct(cols ...field.Expr) ICircuitDo
	Omit(cols ...field.Expr) ICircuitDo
	Join(table schema.Tabler, on ...field.Expr) ICircuitDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICircuitDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICircuitDo
	Group(cols ...field.Expr) ICircuitDo
	Having(conds ...gen.Condition) ICircuitDo
	Limit(limit int) ICircuitDo
	Offset(offset int) ICircuitDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICircuitDo
	Unscoped() ICircuitDo
	Create(values ...*models.Circuit) error
	CreateInBatches(values []*models.Circuit, batchSize int) error
	Save(values ...*models.Circuit) error
	First() (*models.Circuit, error)
	Take() (*models.Circuit, error)
	Last() (*models.Circuit, error)
	Find() ([]*models.Circuit, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Circuit, err error)
	FindInBatches(result *[]*models.Circuit, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Circuit) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICircuitDo
	Assign(attrs ...field.AssignExpr) ICircuitDo
	Joins(fields ...field.RelationField) ICircuitDo
	Preload(fields ...field.RelationField) ICircuitDo
	FirstOrInit() (*models.Circuit, error)
	FirstOrCreate() (*models.Circuit, error)
	FindByPage(offset int, limit int) (result []*models.Circuit, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICircuitDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c circuitDo) Debug() ICircuitDo {
	return c.withDO(c.DO.Debug())
}

func (c circuitDo) WithContext(ctx context.Context) ICircuitDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c circuitDo) ReadDB() ICircuitDo {
	return c.Clauses(dbresolver.Read)
}

func (c circuitDo) WriteDB() ICircuitDo {
	return c.Clauses(dbresolver.Write)
}

func (c circuitDo) Session(config *gorm.Session) ICircuitDo {
	return c.withDO(c.DO.Session(config))
}

func (c circuitDo) Clauses(conds ...clause.Expression) ICircuitDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c circuitDo) Returning(value interface{}, columns ...string) ICircuitDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c circuitDo) Not(conds ...gen.Condition) ICircuitDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c circuitDo) Or(conds ...gen.Condition) ICircuitDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c circuitDo) Select(conds ...field.Expr) ICircuitDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c circuitDo) Where(conds ...gen.Condition) ICircuitDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c circuitDo) Order(conds ...field.Expr) ICircuitDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c circuitDo) Distinct(cols ...field.Expr) ICircuitDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c circuitDo) Omit(cols ...field.Expr) ICircuitDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c circuitDo) Join(table schema.Tabler, on ...field.Expr) ICircuitDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c circuitDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICircuitDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c circuitDo) RightJoin(table schema.Tabler, on ...field.Expr) ICircuitDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c circuitDo) Group(cols ...field.Expr) ICircuitDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c circuitDo) Having(conds ...gen.Condition) ICircuitDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c circuitDo) Limit(limit int) ICircuitDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c circuitDo) Offset(offset int) ICircuitDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c circuitDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICircuitDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c circuitDo) Unscoped() ICircuitDo {
	return c.withDO(c.DO.Unscoped())
}

func (c circuitDo) Create(values ...*models.Circuit) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c circuitDo) CreateInBatches(values []*models.Circuit, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c circuitDo) Save(values ...*models.Circuit) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c circuitDo) First() (*models.Circuit, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Circuit), nil
	}
}

func (c circuitDo) Take() (*models.Circuit, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Circuit), nil
	}
}

func (c circuitDo) Last() (*models.Circuit, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Circuit), nil
	}
}

func (c circuitDo) Find() ([]*models.Circuit, error) {
	result, err := c.DO.Find()
	return result.([]*models.Circuit), err
}

func (c circuitDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Circuit, err error) {
	buf := make([]*models.Circuit, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c circuitDo) FindInBatches(result *[]*models.Circuit, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c circuitDo) Attrs(attrs ...field.AssignExpr) ICircuitDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c circuitDo) Assign(attrs ...field.AssignExpr) ICircuitDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c circuitDo) Joins(fields ...field.RelationField) ICircuitDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c circuitDo) Preload(fields ...field.RelationField) ICircuitDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c circuitDo) FirstOrInit() (*models.Circuit, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Circuit), nil
	}
}

func (c circuitDo) FirstOrCreate() (*models.Circuit, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Circuit), nil
	}
}

func (c circuitDo) FindByPage(offset int, limit int) (result []*models.Circuit, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c circuitDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c circuitDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c circuitDo) Delete(models ...*models.Circuit) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *circuitDo) withDO(do gen.Dao) *circuitDo {
	c.DO = *do.(*gen.DO)
	return c
}
