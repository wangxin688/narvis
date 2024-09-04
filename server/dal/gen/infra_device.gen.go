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

func newDevice(db *gorm.DB, opts ...gen.DOOption) device {
	_device := device{}

	_device.deviceDo.UseDB(db, opts...)
	_device.deviceDo.UseModel(&models.Device{})

	tableName := _device.deviceDo.TableName()
	_device.ALL = field.NewAsterisk(tableName)
	_device.Id = field.NewString(tableName, "id")
	_device.CreatedAt = field.NewTime(tableName, "createdAt")
	_device.UpdatedAt = field.NewTime(tableName, "updatedAt")
	_device.Name = field.NewString(tableName, "name")
	_device.ManagementIp = field.NewString(tableName, "managementIp")
	_device.Status = field.NewString(tableName, "status")
	_device.Platform = field.NewString(tableName, "platform")
	_device.ProductFamily = field.NewString(tableName, "productFamily")
	_device.DeviceModel = field.NewString(tableName, "deviceModel")
	_device.Manufacturer = field.NewString(tableName, "manufacturer")
	_device.DeviceRole = field.NewString(tableName, "deviceRole")
	_device.Floor = field.NewString(tableName, "floor")
	_device.IsRegistered = field.NewBool(tableName, "isRegistered")
	_device.ChassisId = field.NewString(tableName, "chassisId")
	_device.SerialNumber = field.NewString(tableName, "serialNumber")
	_device.Description = field.NewString(tableName, "description")
	_device.OsVersion = field.NewString(tableName, "osVersion")
	_device.OsPatch = field.NewString(tableName, "osPatch")
	_device.RackId = field.NewString(tableName, "rackId")
	_device.RackPosition = field.NewString(tableName, "rackPosition")
	_device.MonitorId = field.NewString(tableName, "monitorId")
	_device.TemplateId = field.NewString(tableName, "templateId")
	_device.SiteId = field.NewString(tableName, "siteId")
	_device.OrganizationId = field.NewString(tableName, "organizationId")
	_device.Rack = deviceBelongsToRack{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Rack", "models.Rack"),
		Site: struct {
			field.RelationField
			Organization struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Rack.Site", "models.Site"),
			Organization: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Rack.Site.Organization", "models.Organization"),
			},
		},
		Organization: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Rack.Organization", "models.Organization"),
		},
	}

	_device.Template = deviceBelongsToTemplate{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Template", "models.Template"),
	}

	_device.Site = deviceBelongsToSite{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Site", "models.Site"),
	}

	_device.Organization = deviceBelongsToOrganization{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Organization", "models.Organization"),
	}

	_device.fillFieldMap()

	return _device
}

type device struct {
	deviceDo

	ALL            field.Asterisk
	Id             field.String
	CreatedAt      field.Time
	UpdatedAt      field.Time
	Name           field.String
	ManagementIp   field.String
	Status         field.String
	Platform       field.String
	ProductFamily  field.String
	DeviceModel    field.String
	Manufacturer   field.String
	DeviceRole     field.String
	Floor          field.String
	IsRegistered   field.Bool
	ChassisId      field.String
	SerialNumber   field.String
	Description    field.String
	OsVersion      field.String
	OsPatch        field.String
	RackId         field.String
	RackPosition   field.String
	MonitorId      field.String
	TemplateId     field.String
	SiteId         field.String
	OrganizationId field.String
	Rack           deviceBelongsToRack

	Template deviceBelongsToTemplate

	Site deviceBelongsToSite

	Organization deviceBelongsToOrganization

	fieldMap map[string]field.Expr
}

func (d device) Table(newTableName string) *device {
	d.deviceDo.UseTable(newTableName)
	return d.updateTableName(newTableName)
}

func (d device) As(alias string) *device {
	d.deviceDo.DO = *(d.deviceDo.As(alias).(*gen.DO))
	return d.updateTableName(alias)
}

func (d *device) updateTableName(table string) *device {
	d.ALL = field.NewAsterisk(table)
	d.Id = field.NewString(table, "id")
	d.CreatedAt = field.NewTime(table, "createdAt")
	d.UpdatedAt = field.NewTime(table, "updatedAt")
	d.Name = field.NewString(table, "name")
	d.ManagementIp = field.NewString(table, "managementIp")
	d.Status = field.NewString(table, "status")
	d.Platform = field.NewString(table, "platform")
	d.ProductFamily = field.NewString(table, "productFamily")
	d.DeviceModel = field.NewString(table, "deviceModel")
	d.Manufacturer = field.NewString(table, "manufacturer")
	d.DeviceRole = field.NewString(table, "deviceRole")
	d.Floor = field.NewString(table, "floor")
	d.IsRegistered = field.NewBool(table, "isRegistered")
	d.ChassisId = field.NewString(table, "chassisId")
	d.SerialNumber = field.NewString(table, "serialNumber")
	d.Description = field.NewString(table, "description")
	d.OsVersion = field.NewString(table, "osVersion")
	d.OsPatch = field.NewString(table, "osPatch")
	d.RackId = field.NewString(table, "rackId")
	d.RackPosition = field.NewString(table, "rackPosition")
	d.MonitorId = field.NewString(table, "monitorId")
	d.TemplateId = field.NewString(table, "templateId")
	d.SiteId = field.NewString(table, "siteId")
	d.OrganizationId = field.NewString(table, "organizationId")

	d.fillFieldMap()

	return d
}

func (d *device) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := d.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (d *device) fillFieldMap() {
	d.fieldMap = make(map[string]field.Expr, 28)
	d.fieldMap["id"] = d.Id
	d.fieldMap["createdAt"] = d.CreatedAt
	d.fieldMap["updatedAt"] = d.UpdatedAt
	d.fieldMap["name"] = d.Name
	d.fieldMap["managementIp"] = d.ManagementIp
	d.fieldMap["status"] = d.Status
	d.fieldMap["platform"] = d.Platform
	d.fieldMap["productFamily"] = d.ProductFamily
	d.fieldMap["deviceModel"] = d.DeviceModel
	d.fieldMap["manufacturer"] = d.Manufacturer
	d.fieldMap["deviceRole"] = d.DeviceRole
	d.fieldMap["floor"] = d.Floor
	d.fieldMap["isRegistered"] = d.IsRegistered
	d.fieldMap["chassisId"] = d.ChassisId
	d.fieldMap["serialNumber"] = d.SerialNumber
	d.fieldMap["description"] = d.Description
	d.fieldMap["osVersion"] = d.OsVersion
	d.fieldMap["osPatch"] = d.OsPatch
	d.fieldMap["rackId"] = d.RackId
	d.fieldMap["rackPosition"] = d.RackPosition
	d.fieldMap["monitorId"] = d.MonitorId
	d.fieldMap["templateId"] = d.TemplateId
	d.fieldMap["siteId"] = d.SiteId
	d.fieldMap["organizationId"] = d.OrganizationId

}

func (d device) clone(db *gorm.DB) device {
	d.deviceDo.ReplaceConnPool(db.Statement.ConnPool)
	return d
}

func (d device) replaceDB(db *gorm.DB) device {
	d.deviceDo.ReplaceDB(db)
	return d
}

type deviceBelongsToRack struct {
	db *gorm.DB

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

func (a deviceBelongsToRack) Where(conds ...field.Expr) *deviceBelongsToRack {
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

func (a deviceBelongsToRack) WithContext(ctx context.Context) *deviceBelongsToRack {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a deviceBelongsToRack) Session(session *gorm.Session) *deviceBelongsToRack {
	a.db = a.db.Session(session)
	return &a
}

func (a deviceBelongsToRack) Model(m *models.Device) *deviceBelongsToRackTx {
	return &deviceBelongsToRackTx{a.db.Model(m).Association(a.Name())}
}

type deviceBelongsToRackTx struct{ tx *gorm.Association }

func (a deviceBelongsToRackTx) Find() (result *models.Rack, err error) {
	return result, a.tx.Find(&result)
}

func (a deviceBelongsToRackTx) Append(values ...*models.Rack) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a deviceBelongsToRackTx) Replace(values ...*models.Rack) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a deviceBelongsToRackTx) Delete(values ...*models.Rack) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a deviceBelongsToRackTx) Clear() error {
	return a.tx.Clear()
}

func (a deviceBelongsToRackTx) Count() int64 {
	return a.tx.Count()
}

type deviceBelongsToTemplate struct {
	db *gorm.DB

	field.RelationField
}

func (a deviceBelongsToTemplate) Where(conds ...field.Expr) *deviceBelongsToTemplate {
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

func (a deviceBelongsToTemplate) WithContext(ctx context.Context) *deviceBelongsToTemplate {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a deviceBelongsToTemplate) Session(session *gorm.Session) *deviceBelongsToTemplate {
	a.db = a.db.Session(session)
	return &a
}

func (a deviceBelongsToTemplate) Model(m *models.Device) *deviceBelongsToTemplateTx {
	return &deviceBelongsToTemplateTx{a.db.Model(m).Association(a.Name())}
}

type deviceBelongsToTemplateTx struct{ tx *gorm.Association }

func (a deviceBelongsToTemplateTx) Find() (result *models.Template, err error) {
	return result, a.tx.Find(&result)
}

func (a deviceBelongsToTemplateTx) Append(values ...*models.Template) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a deviceBelongsToTemplateTx) Replace(values ...*models.Template) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a deviceBelongsToTemplateTx) Delete(values ...*models.Template) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a deviceBelongsToTemplateTx) Clear() error {
	return a.tx.Clear()
}

func (a deviceBelongsToTemplateTx) Count() int64 {
	return a.tx.Count()
}

type deviceBelongsToSite struct {
	db *gorm.DB

	field.RelationField
}

func (a deviceBelongsToSite) Where(conds ...field.Expr) *deviceBelongsToSite {
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

func (a deviceBelongsToSite) WithContext(ctx context.Context) *deviceBelongsToSite {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a deviceBelongsToSite) Session(session *gorm.Session) *deviceBelongsToSite {
	a.db = a.db.Session(session)
	return &a
}

func (a deviceBelongsToSite) Model(m *models.Device) *deviceBelongsToSiteTx {
	return &deviceBelongsToSiteTx{a.db.Model(m).Association(a.Name())}
}

type deviceBelongsToSiteTx struct{ tx *gorm.Association }

func (a deviceBelongsToSiteTx) Find() (result *models.Site, err error) {
	return result, a.tx.Find(&result)
}

func (a deviceBelongsToSiteTx) Append(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a deviceBelongsToSiteTx) Replace(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a deviceBelongsToSiteTx) Delete(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a deviceBelongsToSiteTx) Clear() error {
	return a.tx.Clear()
}

func (a deviceBelongsToSiteTx) Count() int64 {
	return a.tx.Count()
}

type deviceBelongsToOrganization struct {
	db *gorm.DB

	field.RelationField
}

func (a deviceBelongsToOrganization) Where(conds ...field.Expr) *deviceBelongsToOrganization {
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

func (a deviceBelongsToOrganization) WithContext(ctx context.Context) *deviceBelongsToOrganization {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a deviceBelongsToOrganization) Session(session *gorm.Session) *deviceBelongsToOrganization {
	a.db = a.db.Session(session)
	return &a
}

func (a deviceBelongsToOrganization) Model(m *models.Device) *deviceBelongsToOrganizationTx {
	return &deviceBelongsToOrganizationTx{a.db.Model(m).Association(a.Name())}
}

type deviceBelongsToOrganizationTx struct{ tx *gorm.Association }

func (a deviceBelongsToOrganizationTx) Find() (result *models.Organization, err error) {
	return result, a.tx.Find(&result)
}

func (a deviceBelongsToOrganizationTx) Append(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a deviceBelongsToOrganizationTx) Replace(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a deviceBelongsToOrganizationTx) Delete(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a deviceBelongsToOrganizationTx) Clear() error {
	return a.tx.Clear()
}

func (a deviceBelongsToOrganizationTx) Count() int64 {
	return a.tx.Count()
}

type deviceDo struct{ gen.DO }

type IDeviceDo interface {
	gen.SubQuery
	Debug() IDeviceDo
	WithContext(ctx context.Context) IDeviceDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IDeviceDo
	WriteDB() IDeviceDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IDeviceDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IDeviceDo
	Not(conds ...gen.Condition) IDeviceDo
	Or(conds ...gen.Condition) IDeviceDo
	Select(conds ...field.Expr) IDeviceDo
	Where(conds ...gen.Condition) IDeviceDo
	Order(conds ...field.Expr) IDeviceDo
	Distinct(cols ...field.Expr) IDeviceDo
	Omit(cols ...field.Expr) IDeviceDo
	Join(table schema.Tabler, on ...field.Expr) IDeviceDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IDeviceDo
	RightJoin(table schema.Tabler, on ...field.Expr) IDeviceDo
	Group(cols ...field.Expr) IDeviceDo
	Having(conds ...gen.Condition) IDeviceDo
	Limit(limit int) IDeviceDo
	Offset(offset int) IDeviceDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IDeviceDo
	Unscoped() IDeviceDo
	Create(values ...*models.Device) error
	CreateInBatches(values []*models.Device, batchSize int) error
	Save(values ...*models.Device) error
	First() (*models.Device, error)
	Take() (*models.Device, error)
	Last() (*models.Device, error)
	Find() ([]*models.Device, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Device, err error)
	FindInBatches(result *[]*models.Device, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Device) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IDeviceDo
	Assign(attrs ...field.AssignExpr) IDeviceDo
	Joins(fields ...field.RelationField) IDeviceDo
	Preload(fields ...field.RelationField) IDeviceDo
	FirstOrInit() (*models.Device, error)
	FirstOrCreate() (*models.Device, error)
	FindByPage(offset int, limit int) (result []*models.Device, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IDeviceDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (d deviceDo) Debug() IDeviceDo {
	return d.withDO(d.DO.Debug())
}

func (d deviceDo) WithContext(ctx context.Context) IDeviceDo {
	return d.withDO(d.DO.WithContext(ctx))
}

func (d deviceDo) ReadDB() IDeviceDo {
	return d.Clauses(dbresolver.Read)
}

func (d deviceDo) WriteDB() IDeviceDo {
	return d.Clauses(dbresolver.Write)
}

func (d deviceDo) Session(config *gorm.Session) IDeviceDo {
	return d.withDO(d.DO.Session(config))
}

func (d deviceDo) Clauses(conds ...clause.Expression) IDeviceDo {
	return d.withDO(d.DO.Clauses(conds...))
}

func (d deviceDo) Returning(value interface{}, columns ...string) IDeviceDo {
	return d.withDO(d.DO.Returning(value, columns...))
}

func (d deviceDo) Not(conds ...gen.Condition) IDeviceDo {
	return d.withDO(d.DO.Not(conds...))
}

func (d deviceDo) Or(conds ...gen.Condition) IDeviceDo {
	return d.withDO(d.DO.Or(conds...))
}

func (d deviceDo) Select(conds ...field.Expr) IDeviceDo {
	return d.withDO(d.DO.Select(conds...))
}

func (d deviceDo) Where(conds ...gen.Condition) IDeviceDo {
	return d.withDO(d.DO.Where(conds...))
}

func (d deviceDo) Order(conds ...field.Expr) IDeviceDo {
	return d.withDO(d.DO.Order(conds...))
}

func (d deviceDo) Distinct(cols ...field.Expr) IDeviceDo {
	return d.withDO(d.DO.Distinct(cols...))
}

func (d deviceDo) Omit(cols ...field.Expr) IDeviceDo {
	return d.withDO(d.DO.Omit(cols...))
}

func (d deviceDo) Join(table schema.Tabler, on ...field.Expr) IDeviceDo {
	return d.withDO(d.DO.Join(table, on...))
}

func (d deviceDo) LeftJoin(table schema.Tabler, on ...field.Expr) IDeviceDo {
	return d.withDO(d.DO.LeftJoin(table, on...))
}

func (d deviceDo) RightJoin(table schema.Tabler, on ...field.Expr) IDeviceDo {
	return d.withDO(d.DO.RightJoin(table, on...))
}

func (d deviceDo) Group(cols ...field.Expr) IDeviceDo {
	return d.withDO(d.DO.Group(cols...))
}

func (d deviceDo) Having(conds ...gen.Condition) IDeviceDo {
	return d.withDO(d.DO.Having(conds...))
}

func (d deviceDo) Limit(limit int) IDeviceDo {
	return d.withDO(d.DO.Limit(limit))
}

func (d deviceDo) Offset(offset int) IDeviceDo {
	return d.withDO(d.DO.Offset(offset))
}

func (d deviceDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IDeviceDo {
	return d.withDO(d.DO.Scopes(funcs...))
}

func (d deviceDo) Unscoped() IDeviceDo {
	return d.withDO(d.DO.Unscoped())
}

func (d deviceDo) Create(values ...*models.Device) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Create(values)
}

func (d deviceDo) CreateInBatches(values []*models.Device, batchSize int) error {
	return d.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (d deviceDo) Save(values ...*models.Device) error {
	if len(values) == 0 {
		return nil
	}
	return d.DO.Save(values)
}

func (d deviceDo) First() (*models.Device, error) {
	if result, err := d.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Device), nil
	}
}

func (d deviceDo) Take() (*models.Device, error) {
	if result, err := d.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Device), nil
	}
}

func (d deviceDo) Last() (*models.Device, error) {
	if result, err := d.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Device), nil
	}
}

func (d deviceDo) Find() ([]*models.Device, error) {
	result, err := d.DO.Find()
	return result.([]*models.Device), err
}

func (d deviceDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Device, err error) {
	buf := make([]*models.Device, 0, batchSize)
	err = d.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (d deviceDo) FindInBatches(result *[]*models.Device, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return d.DO.FindInBatches(result, batchSize, fc)
}

func (d deviceDo) Attrs(attrs ...field.AssignExpr) IDeviceDo {
	return d.withDO(d.DO.Attrs(attrs...))
}

func (d deviceDo) Assign(attrs ...field.AssignExpr) IDeviceDo {
	return d.withDO(d.DO.Assign(attrs...))
}

func (d deviceDo) Joins(fields ...field.RelationField) IDeviceDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Joins(_f))
	}
	return &d
}

func (d deviceDo) Preload(fields ...field.RelationField) IDeviceDo {
	for _, _f := range fields {
		d = *d.withDO(d.DO.Preload(_f))
	}
	return &d
}

func (d deviceDo) FirstOrInit() (*models.Device, error) {
	if result, err := d.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Device), nil
	}
}

func (d deviceDo) FirstOrCreate() (*models.Device, error) {
	if result, err := d.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Device), nil
	}
}

func (d deviceDo) FindByPage(offset int, limit int) (result []*models.Device, count int64, err error) {
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

func (d deviceDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = d.Count()
	if err != nil {
		return
	}

	err = d.Offset(offset).Limit(limit).Scan(result)
	return
}

func (d deviceDo) Scan(result interface{}) (err error) {
	return d.DO.Scan(result)
}

func (d deviceDo) Delete(models ...*models.Device) (result gen.ResultInfo, err error) {
	return d.DO.Delete(models)
}

func (d *deviceDo) withDO(do gen.Dao) *deviceDo {
	d.DO = *do.(*gen.DO)
	return d
}
