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

func newLLDPNeighbor(db *gorm.DB, opts ...gen.DOOption) lLDPNeighbor {
	_lLDPNeighbor := lLDPNeighbor{}

	_lLDPNeighbor.lLDPNeighborDo.UseDB(db, opts...)
	_lLDPNeighbor.lLDPNeighborDo.UseModel(&models.LLDPNeighbor{})

	tableName := _lLDPNeighbor.lLDPNeighborDo.TableName()
	_lLDPNeighbor.ALL = field.NewAsterisk(tableName)
	_lLDPNeighbor.Id = field.NewString(tableName, "id")
	_lLDPNeighbor.CreatedAt = field.NewTime(tableName, "created_at")
	_lLDPNeighbor.UpdatedAt = field.NewTime(tableName, "updated_at")
	_lLDPNeighbor.SourceInterfaceId = field.NewString(tableName, "source_interface_id")
	_lLDPNeighbor.SourceDeviceId = field.NewString(tableName, "source_device_id")
	_lLDPNeighbor.TargetInterfaceId = field.NewString(tableName, "target_interface_id")
	_lLDPNeighbor.TargetDeviceId = field.NewString(tableName, "target_device_id")
	_lLDPNeighbor.Active = field.NewBool(tableName, "active")
	_lLDPNeighbor.SiteId = field.NewString(tableName, "site_id")
	_lLDPNeighbor.OrganizationId = field.NewString(tableName, "organization_id")
	_lLDPNeighbor.SourceInterface = lLDPNeighborBelongsToSourceInterface{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("SourceInterface", "models.DeviceInterface"),
		Device: struct {
			field.RelationField
			Rack struct {
				field.RelationField
				Location struct {
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
				Site struct {
					field.RelationField
				}
				Organization struct {
					field.RelationField
				}
			}
			Location struct {
				field.RelationField
			}
			Site struct {
				field.RelationField
			}
			Organization struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("SourceInterface.Device", "models.Device"),
			Rack: struct {
				field.RelationField
				Location struct {
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
				Site struct {
					field.RelationField
				}
				Organization struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("SourceInterface.Device.Rack", "models.Rack"),
				Location: struct {
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
				}{
					RelationField: field.NewRelation("SourceInterface.Device.Rack.Location", "models.Location"),
					Parent: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("SourceInterface.Device.Rack.Location.Parent", "models.Location"),
					},
					Site: struct {
						field.RelationField
						Organization struct {
							field.RelationField
						}
					}{
						RelationField: field.NewRelation("SourceInterface.Device.Rack.Location.Site", "models.Site"),
						Organization: struct {
							field.RelationField
						}{
							RelationField: field.NewRelation("SourceInterface.Device.Rack.Location.Site.Organization", "models.Organization"),
						},
					},
					Organization: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("SourceInterface.Device.Rack.Location.Organization", "models.Organization"),
					},
				},
				Site: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("SourceInterface.Device.Rack.Site", "models.Site"),
				},
				Organization: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("SourceInterface.Device.Rack.Organization", "models.Organization"),
				},
			},
			Location: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("SourceInterface.Device.Location", "models.Location"),
			},
			Site: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("SourceInterface.Device.Site", "models.Site"),
			},
			Organization: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("SourceInterface.Device.Organization", "models.Organization"),
			},
		},
		Site: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("SourceInterface.Site", "models.Site"),
		},
	}

	_lLDPNeighbor.SourceDevice = lLDPNeighborBelongsToSourceDevice{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("SourceDevice", "models.Device"),
	}

	_lLDPNeighbor.TargetInterface = lLDPNeighborBelongsToTargetInterface{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("TargetInterface", "models.DeviceInterface"),
	}

	_lLDPNeighbor.TargetDevice = lLDPNeighborBelongsToTargetDevice{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("TargetDevice", "models.Device"),
	}

	_lLDPNeighbor.Site = lLDPNeighborBelongsToSite{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Site", "models.Site"),
	}

	_lLDPNeighbor.Organization = lLDPNeighborBelongsToOrganization{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Organization", "models.Organization"),
	}

	_lLDPNeighbor.fillFieldMap()

	return _lLDPNeighbor
}

type lLDPNeighbor struct {
	lLDPNeighborDo

	ALL               field.Asterisk
	Id                field.String
	CreatedAt         field.Time
	UpdatedAt         field.Time
	SourceInterfaceId field.String
	SourceDeviceId    field.String
	TargetInterfaceId field.String
	TargetDeviceId    field.String
	Active            field.Bool
	SiteId            field.String
	OrganizationId    field.String
	SourceInterface   lLDPNeighborBelongsToSourceInterface

	SourceDevice lLDPNeighborBelongsToSourceDevice

	TargetInterface lLDPNeighborBelongsToTargetInterface

	TargetDevice lLDPNeighborBelongsToTargetDevice

	Site lLDPNeighborBelongsToSite

	Organization lLDPNeighborBelongsToOrganization

	fieldMap map[string]field.Expr
}

func (l lLDPNeighbor) Table(newTableName string) *lLDPNeighbor {
	l.lLDPNeighborDo.UseTable(newTableName)
	return l.updateTableName(newTableName)
}

func (l lLDPNeighbor) As(alias string) *lLDPNeighbor {
	l.lLDPNeighborDo.DO = *(l.lLDPNeighborDo.As(alias).(*gen.DO))
	return l.updateTableName(alias)
}

func (l *lLDPNeighbor) updateTableName(table string) *lLDPNeighbor {
	l.ALL = field.NewAsterisk(table)
	l.Id = field.NewString(table, "id")
	l.CreatedAt = field.NewTime(table, "created_at")
	l.UpdatedAt = field.NewTime(table, "updated_at")
	l.SourceInterfaceId = field.NewString(table, "source_interface_id")
	l.SourceDeviceId = field.NewString(table, "source_device_id")
	l.TargetInterfaceId = field.NewString(table, "target_interface_id")
	l.TargetDeviceId = field.NewString(table, "target_device_id")
	l.Active = field.NewBool(table, "active")
	l.SiteId = field.NewString(table, "site_id")
	l.OrganizationId = field.NewString(table, "organization_id")

	l.fillFieldMap()

	return l
}

func (l *lLDPNeighbor) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := l.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (l *lLDPNeighbor) fillFieldMap() {
	l.fieldMap = make(map[string]field.Expr, 16)
	l.fieldMap["id"] = l.Id
	l.fieldMap["created_at"] = l.CreatedAt
	l.fieldMap["updated_at"] = l.UpdatedAt
	l.fieldMap["source_interface_id"] = l.SourceInterfaceId
	l.fieldMap["source_device_id"] = l.SourceDeviceId
	l.fieldMap["target_interface_id"] = l.TargetInterfaceId
	l.fieldMap["target_device_id"] = l.TargetDeviceId
	l.fieldMap["active"] = l.Active
	l.fieldMap["site_id"] = l.SiteId
	l.fieldMap["organization_id"] = l.OrganizationId

}

func (l lLDPNeighbor) clone(db *gorm.DB) lLDPNeighbor {
	l.lLDPNeighborDo.ReplaceConnPool(db.Statement.ConnPool)
	return l
}

func (l lLDPNeighbor) replaceDB(db *gorm.DB) lLDPNeighbor {
	l.lLDPNeighborDo.ReplaceDB(db)
	return l
}

type lLDPNeighborBelongsToSourceInterface struct {
	db *gorm.DB

	field.RelationField

	Device struct {
		field.RelationField
		Rack struct {
			field.RelationField
			Location struct {
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
			Site struct {
				field.RelationField
			}
			Organization struct {
				field.RelationField
			}
		}
		Location struct {
			field.RelationField
		}
		Site struct {
			field.RelationField
		}
		Organization struct {
			field.RelationField
		}
	}
	Site struct {
		field.RelationField
	}
}

func (a lLDPNeighborBelongsToSourceInterface) Where(conds ...field.Expr) *lLDPNeighborBelongsToSourceInterface {
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

func (a lLDPNeighborBelongsToSourceInterface) WithContext(ctx context.Context) *lLDPNeighborBelongsToSourceInterface {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a lLDPNeighborBelongsToSourceInterface) Session(session *gorm.Session) *lLDPNeighborBelongsToSourceInterface {
	a.db = a.db.Session(session)
	return &a
}

func (a lLDPNeighborBelongsToSourceInterface) Model(m *models.LLDPNeighbor) *lLDPNeighborBelongsToSourceInterfaceTx {
	return &lLDPNeighborBelongsToSourceInterfaceTx{a.db.Model(m).Association(a.Name())}
}

type lLDPNeighborBelongsToSourceInterfaceTx struct{ tx *gorm.Association }

func (a lLDPNeighborBelongsToSourceInterfaceTx) Find() (result *models.DeviceInterface, err error) {
	return result, a.tx.Find(&result)
}

func (a lLDPNeighborBelongsToSourceInterfaceTx) Append(values ...*models.DeviceInterface) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a lLDPNeighborBelongsToSourceInterfaceTx) Replace(values ...*models.DeviceInterface) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a lLDPNeighborBelongsToSourceInterfaceTx) Delete(values ...*models.DeviceInterface) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a lLDPNeighborBelongsToSourceInterfaceTx) Clear() error {
	return a.tx.Clear()
}

func (a lLDPNeighborBelongsToSourceInterfaceTx) Count() int64 {
	return a.tx.Count()
}

type lLDPNeighborBelongsToSourceDevice struct {
	db *gorm.DB

	field.RelationField
}

func (a lLDPNeighborBelongsToSourceDevice) Where(conds ...field.Expr) *lLDPNeighborBelongsToSourceDevice {
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

func (a lLDPNeighborBelongsToSourceDevice) WithContext(ctx context.Context) *lLDPNeighborBelongsToSourceDevice {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a lLDPNeighborBelongsToSourceDevice) Session(session *gorm.Session) *lLDPNeighborBelongsToSourceDevice {
	a.db = a.db.Session(session)
	return &a
}

func (a lLDPNeighborBelongsToSourceDevice) Model(m *models.LLDPNeighbor) *lLDPNeighborBelongsToSourceDeviceTx {
	return &lLDPNeighborBelongsToSourceDeviceTx{a.db.Model(m).Association(a.Name())}
}

type lLDPNeighborBelongsToSourceDeviceTx struct{ tx *gorm.Association }

func (a lLDPNeighborBelongsToSourceDeviceTx) Find() (result *models.Device, err error) {
	return result, a.tx.Find(&result)
}

func (a lLDPNeighborBelongsToSourceDeviceTx) Append(values ...*models.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a lLDPNeighborBelongsToSourceDeviceTx) Replace(values ...*models.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a lLDPNeighborBelongsToSourceDeviceTx) Delete(values ...*models.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a lLDPNeighborBelongsToSourceDeviceTx) Clear() error {
	return a.tx.Clear()
}

func (a lLDPNeighborBelongsToSourceDeviceTx) Count() int64 {
	return a.tx.Count()
}

type lLDPNeighborBelongsToTargetInterface struct {
	db *gorm.DB

	field.RelationField
}

func (a lLDPNeighborBelongsToTargetInterface) Where(conds ...field.Expr) *lLDPNeighborBelongsToTargetInterface {
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

func (a lLDPNeighborBelongsToTargetInterface) WithContext(ctx context.Context) *lLDPNeighborBelongsToTargetInterface {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a lLDPNeighborBelongsToTargetInterface) Session(session *gorm.Session) *lLDPNeighborBelongsToTargetInterface {
	a.db = a.db.Session(session)
	return &a
}

func (a lLDPNeighborBelongsToTargetInterface) Model(m *models.LLDPNeighbor) *lLDPNeighborBelongsToTargetInterfaceTx {
	return &lLDPNeighborBelongsToTargetInterfaceTx{a.db.Model(m).Association(a.Name())}
}

type lLDPNeighborBelongsToTargetInterfaceTx struct{ tx *gorm.Association }

func (a lLDPNeighborBelongsToTargetInterfaceTx) Find() (result *models.DeviceInterface, err error) {
	return result, a.tx.Find(&result)
}

func (a lLDPNeighborBelongsToTargetInterfaceTx) Append(values ...*models.DeviceInterface) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a lLDPNeighborBelongsToTargetInterfaceTx) Replace(values ...*models.DeviceInterface) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a lLDPNeighborBelongsToTargetInterfaceTx) Delete(values ...*models.DeviceInterface) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a lLDPNeighborBelongsToTargetInterfaceTx) Clear() error {
	return a.tx.Clear()
}

func (a lLDPNeighborBelongsToTargetInterfaceTx) Count() int64 {
	return a.tx.Count()
}

type lLDPNeighborBelongsToTargetDevice struct {
	db *gorm.DB

	field.RelationField
}

func (a lLDPNeighborBelongsToTargetDevice) Where(conds ...field.Expr) *lLDPNeighborBelongsToTargetDevice {
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

func (a lLDPNeighborBelongsToTargetDevice) WithContext(ctx context.Context) *lLDPNeighborBelongsToTargetDevice {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a lLDPNeighborBelongsToTargetDevice) Session(session *gorm.Session) *lLDPNeighborBelongsToTargetDevice {
	a.db = a.db.Session(session)
	return &a
}

func (a lLDPNeighborBelongsToTargetDevice) Model(m *models.LLDPNeighbor) *lLDPNeighborBelongsToTargetDeviceTx {
	return &lLDPNeighborBelongsToTargetDeviceTx{a.db.Model(m).Association(a.Name())}
}

type lLDPNeighborBelongsToTargetDeviceTx struct{ tx *gorm.Association }

func (a lLDPNeighborBelongsToTargetDeviceTx) Find() (result *models.Device, err error) {
	return result, a.tx.Find(&result)
}

func (a lLDPNeighborBelongsToTargetDeviceTx) Append(values ...*models.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a lLDPNeighborBelongsToTargetDeviceTx) Replace(values ...*models.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a lLDPNeighborBelongsToTargetDeviceTx) Delete(values ...*models.Device) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a lLDPNeighborBelongsToTargetDeviceTx) Clear() error {
	return a.tx.Clear()
}

func (a lLDPNeighborBelongsToTargetDeviceTx) Count() int64 {
	return a.tx.Count()
}

type lLDPNeighborBelongsToSite struct {
	db *gorm.DB

	field.RelationField
}

func (a lLDPNeighborBelongsToSite) Where(conds ...field.Expr) *lLDPNeighborBelongsToSite {
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

func (a lLDPNeighborBelongsToSite) WithContext(ctx context.Context) *lLDPNeighborBelongsToSite {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a lLDPNeighborBelongsToSite) Session(session *gorm.Session) *lLDPNeighborBelongsToSite {
	a.db = a.db.Session(session)
	return &a
}

func (a lLDPNeighborBelongsToSite) Model(m *models.LLDPNeighbor) *lLDPNeighborBelongsToSiteTx {
	return &lLDPNeighborBelongsToSiteTx{a.db.Model(m).Association(a.Name())}
}

type lLDPNeighborBelongsToSiteTx struct{ tx *gorm.Association }

func (a lLDPNeighborBelongsToSiteTx) Find() (result *models.Site, err error) {
	return result, a.tx.Find(&result)
}

func (a lLDPNeighborBelongsToSiteTx) Append(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a lLDPNeighborBelongsToSiteTx) Replace(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a lLDPNeighborBelongsToSiteTx) Delete(values ...*models.Site) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a lLDPNeighborBelongsToSiteTx) Clear() error {
	return a.tx.Clear()
}

func (a lLDPNeighborBelongsToSiteTx) Count() int64 {
	return a.tx.Count()
}

type lLDPNeighborBelongsToOrganization struct {
	db *gorm.DB

	field.RelationField
}

func (a lLDPNeighborBelongsToOrganization) Where(conds ...field.Expr) *lLDPNeighborBelongsToOrganization {
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

func (a lLDPNeighborBelongsToOrganization) WithContext(ctx context.Context) *lLDPNeighborBelongsToOrganization {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a lLDPNeighborBelongsToOrganization) Session(session *gorm.Session) *lLDPNeighborBelongsToOrganization {
	a.db = a.db.Session(session)
	return &a
}

func (a lLDPNeighborBelongsToOrganization) Model(m *models.LLDPNeighbor) *lLDPNeighborBelongsToOrganizationTx {
	return &lLDPNeighborBelongsToOrganizationTx{a.db.Model(m).Association(a.Name())}
}

type lLDPNeighborBelongsToOrganizationTx struct{ tx *gorm.Association }

func (a lLDPNeighborBelongsToOrganizationTx) Find() (result *models.Organization, err error) {
	return result, a.tx.Find(&result)
}

func (a lLDPNeighborBelongsToOrganizationTx) Append(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a lLDPNeighborBelongsToOrganizationTx) Replace(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a lLDPNeighborBelongsToOrganizationTx) Delete(values ...*models.Organization) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a lLDPNeighborBelongsToOrganizationTx) Clear() error {
	return a.tx.Clear()
}

func (a lLDPNeighborBelongsToOrganizationTx) Count() int64 {
	return a.tx.Count()
}

type lLDPNeighborDo struct{ gen.DO }

type ILLDPNeighborDo interface {
	gen.SubQuery
	Debug() ILLDPNeighborDo
	WithContext(ctx context.Context) ILLDPNeighborDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ILLDPNeighborDo
	WriteDB() ILLDPNeighborDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ILLDPNeighborDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ILLDPNeighborDo
	Not(conds ...gen.Condition) ILLDPNeighborDo
	Or(conds ...gen.Condition) ILLDPNeighborDo
	Select(conds ...field.Expr) ILLDPNeighborDo
	Where(conds ...gen.Condition) ILLDPNeighborDo
	Order(conds ...field.Expr) ILLDPNeighborDo
	Distinct(cols ...field.Expr) ILLDPNeighborDo
	Omit(cols ...field.Expr) ILLDPNeighborDo
	Join(table schema.Tabler, on ...field.Expr) ILLDPNeighborDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ILLDPNeighborDo
	RightJoin(table schema.Tabler, on ...field.Expr) ILLDPNeighborDo
	Group(cols ...field.Expr) ILLDPNeighborDo
	Having(conds ...gen.Condition) ILLDPNeighborDo
	Limit(limit int) ILLDPNeighborDo
	Offset(offset int) ILLDPNeighborDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ILLDPNeighborDo
	Unscoped() ILLDPNeighborDo
	Create(values ...*models.LLDPNeighbor) error
	CreateInBatches(values []*models.LLDPNeighbor, batchSize int) error
	Save(values ...*models.LLDPNeighbor) error
	First() (*models.LLDPNeighbor, error)
	Take() (*models.LLDPNeighbor, error)
	Last() (*models.LLDPNeighbor, error)
	Find() ([]*models.LLDPNeighbor, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.LLDPNeighbor, err error)
	FindInBatches(result *[]*models.LLDPNeighbor, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.LLDPNeighbor) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ILLDPNeighborDo
	Assign(attrs ...field.AssignExpr) ILLDPNeighborDo
	Joins(fields ...field.RelationField) ILLDPNeighborDo
	Preload(fields ...field.RelationField) ILLDPNeighborDo
	FirstOrInit() (*models.LLDPNeighbor, error)
	FirstOrCreate() (*models.LLDPNeighbor, error)
	FindByPage(offset int, limit int) (result []*models.LLDPNeighbor, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ILLDPNeighborDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (l lLDPNeighborDo) Debug() ILLDPNeighborDo {
	return l.withDO(l.DO.Debug())
}

func (l lLDPNeighborDo) WithContext(ctx context.Context) ILLDPNeighborDo {
	return l.withDO(l.DO.WithContext(ctx))
}

func (l lLDPNeighborDo) ReadDB() ILLDPNeighborDo {
	return l.Clauses(dbresolver.Read)
}

func (l lLDPNeighborDo) WriteDB() ILLDPNeighborDo {
	return l.Clauses(dbresolver.Write)
}

func (l lLDPNeighborDo) Session(config *gorm.Session) ILLDPNeighborDo {
	return l.withDO(l.DO.Session(config))
}

func (l lLDPNeighborDo) Clauses(conds ...clause.Expression) ILLDPNeighborDo {
	return l.withDO(l.DO.Clauses(conds...))
}

func (l lLDPNeighborDo) Returning(value interface{}, columns ...string) ILLDPNeighborDo {
	return l.withDO(l.DO.Returning(value, columns...))
}

func (l lLDPNeighborDo) Not(conds ...gen.Condition) ILLDPNeighborDo {
	return l.withDO(l.DO.Not(conds...))
}

func (l lLDPNeighborDo) Or(conds ...gen.Condition) ILLDPNeighborDo {
	return l.withDO(l.DO.Or(conds...))
}

func (l lLDPNeighborDo) Select(conds ...field.Expr) ILLDPNeighborDo {
	return l.withDO(l.DO.Select(conds...))
}

func (l lLDPNeighborDo) Where(conds ...gen.Condition) ILLDPNeighborDo {
	return l.withDO(l.DO.Where(conds...))
}

func (l lLDPNeighborDo) Order(conds ...field.Expr) ILLDPNeighborDo {
	return l.withDO(l.DO.Order(conds...))
}

func (l lLDPNeighborDo) Distinct(cols ...field.Expr) ILLDPNeighborDo {
	return l.withDO(l.DO.Distinct(cols...))
}

func (l lLDPNeighborDo) Omit(cols ...field.Expr) ILLDPNeighborDo {
	return l.withDO(l.DO.Omit(cols...))
}

func (l lLDPNeighborDo) Join(table schema.Tabler, on ...field.Expr) ILLDPNeighborDo {
	return l.withDO(l.DO.Join(table, on...))
}

func (l lLDPNeighborDo) LeftJoin(table schema.Tabler, on ...field.Expr) ILLDPNeighborDo {
	return l.withDO(l.DO.LeftJoin(table, on...))
}

func (l lLDPNeighborDo) RightJoin(table schema.Tabler, on ...field.Expr) ILLDPNeighborDo {
	return l.withDO(l.DO.RightJoin(table, on...))
}

func (l lLDPNeighborDo) Group(cols ...field.Expr) ILLDPNeighborDo {
	return l.withDO(l.DO.Group(cols...))
}

func (l lLDPNeighborDo) Having(conds ...gen.Condition) ILLDPNeighborDo {
	return l.withDO(l.DO.Having(conds...))
}

func (l lLDPNeighborDo) Limit(limit int) ILLDPNeighborDo {
	return l.withDO(l.DO.Limit(limit))
}

func (l lLDPNeighborDo) Offset(offset int) ILLDPNeighborDo {
	return l.withDO(l.DO.Offset(offset))
}

func (l lLDPNeighborDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ILLDPNeighborDo {
	return l.withDO(l.DO.Scopes(funcs...))
}

func (l lLDPNeighborDo) Unscoped() ILLDPNeighborDo {
	return l.withDO(l.DO.Unscoped())
}

func (l lLDPNeighborDo) Create(values ...*models.LLDPNeighbor) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Create(values)
}

func (l lLDPNeighborDo) CreateInBatches(values []*models.LLDPNeighbor, batchSize int) error {
	return l.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (l lLDPNeighborDo) Save(values ...*models.LLDPNeighbor) error {
	if len(values) == 0 {
		return nil
	}
	return l.DO.Save(values)
}

func (l lLDPNeighborDo) First() (*models.LLDPNeighbor, error) {
	if result, err := l.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.LLDPNeighbor), nil
	}
}

func (l lLDPNeighborDo) Take() (*models.LLDPNeighbor, error) {
	if result, err := l.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.LLDPNeighbor), nil
	}
}

func (l lLDPNeighborDo) Last() (*models.LLDPNeighbor, error) {
	if result, err := l.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.LLDPNeighbor), nil
	}
}

func (l lLDPNeighborDo) Find() ([]*models.LLDPNeighbor, error) {
	result, err := l.DO.Find()
	return result.([]*models.LLDPNeighbor), err
}

func (l lLDPNeighborDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.LLDPNeighbor, err error) {
	buf := make([]*models.LLDPNeighbor, 0, batchSize)
	err = l.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (l lLDPNeighborDo) FindInBatches(result *[]*models.LLDPNeighbor, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return l.DO.FindInBatches(result, batchSize, fc)
}

func (l lLDPNeighborDo) Attrs(attrs ...field.AssignExpr) ILLDPNeighborDo {
	return l.withDO(l.DO.Attrs(attrs...))
}

func (l lLDPNeighborDo) Assign(attrs ...field.AssignExpr) ILLDPNeighborDo {
	return l.withDO(l.DO.Assign(attrs...))
}

func (l lLDPNeighborDo) Joins(fields ...field.RelationField) ILLDPNeighborDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Joins(_f))
	}
	return &l
}

func (l lLDPNeighborDo) Preload(fields ...field.RelationField) ILLDPNeighborDo {
	for _, _f := range fields {
		l = *l.withDO(l.DO.Preload(_f))
	}
	return &l
}

func (l lLDPNeighborDo) FirstOrInit() (*models.LLDPNeighbor, error) {
	if result, err := l.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.LLDPNeighbor), nil
	}
}

func (l lLDPNeighborDo) FirstOrCreate() (*models.LLDPNeighbor, error) {
	if result, err := l.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.LLDPNeighbor), nil
	}
}

func (l lLDPNeighborDo) FindByPage(offset int, limit int) (result []*models.LLDPNeighbor, count int64, err error) {
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

func (l lLDPNeighborDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = l.Count()
	if err != nil {
		return
	}

	err = l.Offset(offset).Limit(limit).Scan(result)
	return
}

func (l lLDPNeighborDo) Scan(result interface{}) (err error) {
	return l.DO.Scan(result)
}

func (l lLDPNeighborDo) Delete(models ...*models.LLDPNeighbor) (result gen.ResultInfo, err error) {
	return l.DO.Delete(models)
}

func (l *lLDPNeighborDo) withDO(do gen.Dao) *lLDPNeighborDo {
	l.DO = *do.(*gen.DO)
	return l
}
