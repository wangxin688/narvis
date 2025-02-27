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

func newServer(db *gorm.DB, opts ...gen.DOOption) server {
	_server := server{}

	_server.serverDo.UseDB(db, opts...)
	_server.serverDo.UseModel(&models.Server{})

	tableName := _server.serverDo.TableName()
	_server.ALL = field.NewAsterisk(tableName)
	_server.Id = field.NewString(tableName, "id")
	_server.CreatedAt = field.NewTime(tableName, "createdAt")
	_server.UpdatedAt = field.NewTime(tableName, "updatedAt")
	_server.Name = field.NewString(tableName, "name")
	_server.ManagementIp = field.NewString(tableName, "managementIp")
	_server.Manufacturer = field.NewString(tableName, "manufacturer")
	_server.Status = field.NewString(tableName, "status")
	_server.OsVersion = field.NewString(tableName, "osVersion")
	_server.RackId = field.NewString(tableName, "rackId")
	_server.RackPosition = field.NewString(tableName, "rackPosition")
	_server.Cpu = field.NewUint8(tableName, "Cpu")
	_server.Memory = field.NewUint64(tableName, "memory")
	_server.Disk = field.NewUint64(tableName, "disk")
	_server.Description = field.NewString(tableName, "description")
	_server.MonitorId = field.NewString(tableName, "monitorId")
	_server.TemplateId = field.NewString(tableName, "templateId")
	_server.SiteId = field.NewString(tableName, "siteId")
	_server.OrganizationId = field.NewString(tableName, "organizationId")
	_server.Rack = serverBelongsToRack{
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

	_server.Template = serverBelongsToTemplate{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Template", "models.Template"),
	}

	_server.Site = serverBelongsToSite{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Site", "models.Site"),
	}

	_server.Organization = serverBelongsToOrganization{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Organization", "models.Organization"),
	}

	_server.fillFieldMap()

	return _server
}

type server struct {
	serverDo

	ALL            field.Asterisk
	Id             field.String
	CreatedAt      field.Time
	UpdatedAt      field.Time
	Name           field.String
	ManagementIp   field.String
	Manufacturer   field.String
	Status         field.String
	OsVersion      field.String
	RackId         field.String
	RackPosition   field.String
	Cpu            field.Uint8
	Memory         field.Uint64
	Disk           field.Uint64
	Description    field.String
	MonitorId      field.String
	TemplateId     field.String
	SiteId         field.String
	OrganizationId field.String
	Rack           serverBelongsToRack

	Template serverBelongsToTemplate

	Site serverBelongsToSite

	Organization serverBelongsToOrganization

	fieldMap map[string]field.Expr
}

func (s server) Table(newTableName string) *server {
	s.serverDo.UseTable(newTableName)
	return s.updateTableName(newTableName)
}

func (s server) As(alias string) *server {
	s.serverDo.DO = *(s.serverDo.As(alias).(*gen.DO))
	return s.updateTableName(alias)
}

func (s *server) updateTableName(table string) *server {
	s.ALL = field.NewAsterisk(table)
	s.Id = field.NewString(table, "id")
	s.CreatedAt = field.NewTime(table, "createdAt")
	s.UpdatedAt = field.NewTime(table, "updatedAt")
	s.Name = field.NewString(table, "name")
	s.ManagementIp = field.NewString(table, "managementIp")
	s.Manufacturer = field.NewString(table, "manufacturer")
	s.Status = field.NewString(table, "status")
	s.OsVersion = field.NewString(table, "osVersion")
	s.RackId = field.NewString(table, "rackId")
	s.RackPosition = field.NewString(table, "rackPosition")
	s.Cpu = field.NewUint8(table, "Cpu")
	s.Memory = field.NewUint64(table, "memory")
	s.Disk = field.NewUint64(table, "disk")
	s.Description = field.NewString(table, "description")
	s.MonitorId = field.NewString(table, "monitorId")
	s.TemplateId = field.NewString(table, "templateId")
	s.SiteId = field.NewString(table, "siteId")
	s.OrganizationId = field.NewString(table, "organizationId")

	s.fillFieldMap()

	return s
}

func (s *server) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := s.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (s *server) fillFieldMap() {
	s.fieldMap = make(map[string]field.Expr, 22)
	s.fieldMap["id"] = s.Id
	s.fieldMap["createdAt"] = s.CreatedAt
	s.fieldMap["updatedAt"] = s.UpdatedAt
	s.fieldMap["name"] = s.Name
	s.fieldMap["managementIp"] = s.ManagementIp
	s.fieldMap["manufacturer"] = s.Manufacturer
	s.fieldMap["status"] = s.Status
	s.fieldMap["osVersion"] = s.OsVersion
	s.fieldMap["rackId"] = s.RackId
	s.fieldMap["rackPosition"] = s.RackPosition
	s.fieldMap["Cpu"] = s.Cpu
	s.fieldMap["memory"] = s.Memory
	s.fieldMap["disk"] = s.Disk
	s.fieldMap["description"] = s.Description
	s.fieldMap["monitorId"] = s.MonitorId
	s.fieldMap["templateId"] = s.TemplateId
	s.fieldMap["siteId"] = s.SiteId
	s.fieldMap["organizationId"] = s.OrganizationId

}

func (s server) clone(db *gorm.DB) server {
	s.serverDo.ReplaceConnPool(db.Statement.ConnPool)
	return s
}

func (s server) replaceDB(db *gorm.DB) server {
	s.serverDo.ReplaceDB(db)
	return s
}

type serverBelongsToRack struct {
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

func (a serverBelongsToRack) Where(conds ...field.Expr) *serverBelongsToRack {
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

func (a serverBelongsToRack) WithContext(ctx context.Context) *serverBelongsToRack {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a serverBelongsToRack) Session(session *gorm.Session) *serverBelongsToRack {
	a.db = a.db.Session(session)
	return &a
}

func (a serverBelongsToRack) Model(m *models.Server) *serverBelongsToRackTx {
	return &serverBelongsToRackTx{a.db.Model(m).Association(a.Name())}
}

type serverBelongsToRackTx struct{ tx *gorm.Association }

func (a serverBelongsToRackTx) Find() (result *models.Rack, err error) {
	return result, a.tx.Find(&result)
}

func (a serverBelongsToRackTx) Append(values ...*models.Rack) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a serverBelongsToRackTx) Replace(values ...*models.Rack) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a serverBelongsToRackTx) Delete(values ...*models.Rack) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a serverBelongsToRackTx) Clear() error {
	return a.tx.Clear()
}

func (a serverBelongsToRackTx) Count() int64 {
	return a.tx.Count()
}

type serverBelongsToTemplate struct {
	db *gorm.DB

	field.RelationField
}

func (a serverBelongsToTemplate) Where(conds ...field.Expr) *serverBelongsToTemplate {
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

func (a serverBelongsToTemplate) WithContext(ctx context.Context) *serverBelongsToTemplate {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a serverBelongsToTemplate) Session(session *gorm.Session) *serverBelongsToTemplate {
	a.db = a.db.Session(session)
	return &a
}

func (a serverBelongsToTemplate) Model(m *models.Server) *serverBelongsToTemplateTx {
	return &serverBelongsToTemplateTx{a.db.Model(m).Association(a.Name())}
}

type serverBelongsToTemplateTx struct{ tx *gorm.Association }

func (a serverBelongsToTemplateTx) Find() (result *models.Template, err error) {
	return result, a.tx.Find(&result)
}

func (a serverBelongsToTemplateTx) Append(values ...*models.Template) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a serverBelongsToTemplateTx) Replace(values ...*models.Template) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a serverBelongsToTemplateTx) Delete(values ...*models.Template) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a serverBelongsToTemplateTx) Clear() error {
	return a.tx.Clear()
}

func (a serverBelongsToTemplateTx) Count() int64 {
	return a.tx.Count()
}

type serverBelongsToSite struct {
	db *gorm.DB

	field.RelationField
}

func (a serverBelongsToSite) Where(conds ...field.Expr) *serverBelongsToSite {
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

func (a serverBelongsToSite) WithContext(ctx context.Context) *serverBelongsToSite {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a serverBelongsToSite) Session(session *gorm.Session) *serverBelongsToSite {
	a.db = a.db.Session(session)
	return &a
}

func (a serverBelongsToSite) Model(m *models.Server) *serverBelongsToSiteTx {
	return &serverBelongsToSiteTx{a.db.Model(m).Association(a.Name())}
}

type serverBelongsToSiteTx struct{ tx *gorm.Association }

func (a serverBelongsToSiteTx) Find() (result *models.Site, err error) {
	return result, a.tx.Find(&result)
}

func (a serverBelongsToSiteTx) Append(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a serverBelongsToSiteTx) Replace(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a serverBelongsToSiteTx) Delete(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a serverBelongsToSiteTx) Clear() error {
	return a.tx.Clear()
}

func (a serverBelongsToSiteTx) Count() int64 {
	return a.tx.Count()
}

type serverBelongsToOrganization struct {
	db *gorm.DB

	field.RelationField
}

func (a serverBelongsToOrganization) Where(conds ...field.Expr) *serverBelongsToOrganization {
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

func (a serverBelongsToOrganization) WithContext(ctx context.Context) *serverBelongsToOrganization {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a serverBelongsToOrganization) Session(session *gorm.Session) *serverBelongsToOrganization {
	a.db = a.db.Session(session)
	return &a
}

func (a serverBelongsToOrganization) Model(m *models.Server) *serverBelongsToOrganizationTx {
	return &serverBelongsToOrganizationTx{a.db.Model(m).Association(a.Name())}
}

type serverBelongsToOrganizationTx struct{ tx *gorm.Association }

func (a serverBelongsToOrganizationTx) Find() (result *models.Organization, err error) {
	return result, a.tx.Find(&result)
}

func (a serverBelongsToOrganizationTx) Append(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a serverBelongsToOrganizationTx) Replace(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a serverBelongsToOrganizationTx) Delete(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a serverBelongsToOrganizationTx) Clear() error {
	return a.tx.Clear()
}

func (a serverBelongsToOrganizationTx) Count() int64 {
	return a.tx.Count()
}

type serverDo struct{ gen.DO }

type IServerDo interface {
	gen.SubQuery
	Debug() IServerDo
	WithContext(ctx context.Context) IServerDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IServerDo
	WriteDB() IServerDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IServerDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IServerDo
	Not(conds ...gen.Condition) IServerDo
	Or(conds ...gen.Condition) IServerDo
	Select(conds ...field.Expr) IServerDo
	Where(conds ...gen.Condition) IServerDo
	Order(conds ...field.Expr) IServerDo
	Distinct(cols ...field.Expr) IServerDo
	Omit(cols ...field.Expr) IServerDo
	Join(table schema.Tabler, on ...field.Expr) IServerDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IServerDo
	RightJoin(table schema.Tabler, on ...field.Expr) IServerDo
	Group(cols ...field.Expr) IServerDo
	Having(conds ...gen.Condition) IServerDo
	Limit(limit int) IServerDo
	Offset(offset int) IServerDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IServerDo
	Unscoped() IServerDo
	Create(values ...*models.Server) error
	CreateInBatches(values []*models.Server, batchSize int) error
	Save(values ...*models.Server) error
	First() (*models.Server, error)
	Take() (*models.Server, error)
	Last() (*models.Server, error)
	Find() ([]*models.Server, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Server, err error)
	FindInBatches(result *[]*models.Server, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.Server) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IServerDo
	Assign(attrs ...field.AssignExpr) IServerDo
	Joins(fields ...field.RelationField) IServerDo
	Preload(fields ...field.RelationField) IServerDo
	FirstOrInit() (*models.Server, error)
	FirstOrCreate() (*models.Server, error)
	FindByPage(offset int, limit int) (result []*models.Server, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IServerDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (s serverDo) Debug() IServerDo {
	return s.withDO(s.DO.Debug())
}

func (s serverDo) WithContext(ctx context.Context) IServerDo {
	return s.withDO(s.DO.WithContext(ctx))
}

func (s serverDo) ReadDB() IServerDo {
	return s.Clauses(dbresolver.Read)
}

func (s serverDo) WriteDB() IServerDo {
	return s.Clauses(dbresolver.Write)
}

func (s serverDo) Session(config *gorm.Session) IServerDo {
	return s.withDO(s.DO.Session(config))
}

func (s serverDo) Clauses(conds ...clause.Expression) IServerDo {
	return s.withDO(s.DO.Clauses(conds...))
}

func (s serverDo) Returning(value interface{}, columns ...string) IServerDo {
	return s.withDO(s.DO.Returning(value, columns...))
}

func (s serverDo) Not(conds ...gen.Condition) IServerDo {
	return s.withDO(s.DO.Not(conds...))
}

func (s serverDo) Or(conds ...gen.Condition) IServerDo {
	return s.withDO(s.DO.Or(conds...))
}

func (s serverDo) Select(conds ...field.Expr) IServerDo {
	return s.withDO(s.DO.Select(conds...))
}

func (s serverDo) Where(conds ...gen.Condition) IServerDo {
	return s.withDO(s.DO.Where(conds...))
}

func (s serverDo) Order(conds ...field.Expr) IServerDo {
	return s.withDO(s.DO.Order(conds...))
}

func (s serverDo) Distinct(cols ...field.Expr) IServerDo {
	return s.withDO(s.DO.Distinct(cols...))
}

func (s serverDo) Omit(cols ...field.Expr) IServerDo {
	return s.withDO(s.DO.Omit(cols...))
}

func (s serverDo) Join(table schema.Tabler, on ...field.Expr) IServerDo {
	return s.withDO(s.DO.Join(table, on...))
}

func (s serverDo) LeftJoin(table schema.Tabler, on ...field.Expr) IServerDo {
	return s.withDO(s.DO.LeftJoin(table, on...))
}

func (s serverDo) RightJoin(table schema.Tabler, on ...field.Expr) IServerDo {
	return s.withDO(s.DO.RightJoin(table, on...))
}

func (s serverDo) Group(cols ...field.Expr) IServerDo {
	return s.withDO(s.DO.Group(cols...))
}

func (s serverDo) Having(conds ...gen.Condition) IServerDo {
	return s.withDO(s.DO.Having(conds...))
}

func (s serverDo) Limit(limit int) IServerDo {
	return s.withDO(s.DO.Limit(limit))
}

func (s serverDo) Offset(offset int) IServerDo {
	return s.withDO(s.DO.Offset(offset))
}

func (s serverDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IServerDo {
	return s.withDO(s.DO.Scopes(funcs...))
}

func (s serverDo) Unscoped() IServerDo {
	return s.withDO(s.DO.Unscoped())
}

func (s serverDo) Create(values ...*models.Server) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Create(values)
}

func (s serverDo) CreateInBatches(values []*models.Server, batchSize int) error {
	return s.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (s serverDo) Save(values ...*models.Server) error {
	if len(values) == 0 {
		return nil
	}
	return s.DO.Save(values)
}

func (s serverDo) First() (*models.Server, error) {
	if result, err := s.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.Server), nil
	}
}

func (s serverDo) Take() (*models.Server, error) {
	if result, err := s.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.Server), nil
	}
}

func (s serverDo) Last() (*models.Server, error) {
	if result, err := s.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.Server), nil
	}
}

func (s serverDo) Find() ([]*models.Server, error) {
	result, err := s.DO.Find()
	return result.([]*models.Server), err
}

func (s serverDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.Server, err error) {
	buf := make([]*models.Server, 0, batchSize)
	err = s.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (s serverDo) FindInBatches(result *[]*models.Server, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return s.DO.FindInBatches(result, batchSize, fc)
}

func (s serverDo) Attrs(attrs ...field.AssignExpr) IServerDo {
	return s.withDO(s.DO.Attrs(attrs...))
}

func (s serverDo) Assign(attrs ...field.AssignExpr) IServerDo {
	return s.withDO(s.DO.Assign(attrs...))
}

func (s serverDo) Joins(fields ...field.RelationField) IServerDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Joins(_f))
	}
	return &s
}

func (s serverDo) Preload(fields ...field.RelationField) IServerDo {
	for _, _f := range fields {
		s = *s.withDO(s.DO.Preload(_f))
	}
	return &s
}

func (s serverDo) FirstOrInit() (*models.Server, error) {
	if result, err := s.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.Server), nil
	}
}

func (s serverDo) FirstOrCreate() (*models.Server, error) {
	if result, err := s.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.Server), nil
	}
}

func (s serverDo) FindByPage(offset int, limit int) (result []*models.Server, count int64, err error) {
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

func (s serverDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = s.Count()
	if err != nil {
		return
	}

	err = s.Offset(offset).Limit(limit).Scan(result)
	return
}

func (s serverDo) Scan(result interface{}) (err error) {
	return s.DO.Scan(result)
}

func (s serverDo) Delete(models ...*models.Server) (result gen.ResultInfo, err error) {
	return s.DO.Delete(models)
}

func (s *serverDo) withDO(do gen.Dao) *serverDo {
	s.DO = *do.(*gen.DO)
	return s
}
