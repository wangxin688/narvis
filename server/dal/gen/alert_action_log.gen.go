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

func newAlertActionLog(db *gorm.DB, opts ...gen.DOOption) alertActionLog {
	_alertActionLog := alertActionLog{}

	_alertActionLog.alertActionLogDo.UseDB(db, opts...)
	_alertActionLog.alertActionLogDo.UseModel(&models.AlertActionLog{})

	tableName := _alertActionLog.alertActionLogDo.TableName()
	_alertActionLog.ALL = field.NewAsterisk(tableName)
	_alertActionLog.Id = field.NewString(tableName, "id")
	_alertActionLog.CreatedAt = field.NewTime(tableName, "createdAt")
	_alertActionLog.UpdatedAt = field.NewTime(tableName, "updatedAt")
	_alertActionLog.Acknowledged = field.NewBool(tableName, "acknowledged")
	_alertActionLog.Resolved = field.NewBool(tableName, "resolved")
	_alertActionLog.Suppressed = field.NewBool(tableName, "suppressed")
	_alertActionLog.Comment = field.NewString(tableName, "comment")
	_alertActionLog.AssignUserId = field.NewString(tableName, "assignUserId")
	_alertActionLog.CreatedById = field.NewString(tableName, "createdById")
	_alertActionLog.AssignUser = alertActionLogBelongsToAssignUser{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("AssignUser", "models.User"),
		Role: struct {
			field.RelationField
			Organization struct {
				field.RelationField
			}
			Menus struct {
				field.RelationField
				Parent struct {
					field.RelationField
				}
				Permission struct {
					field.RelationField
					Menu struct {
						field.RelationField
					}
				}
			}
		}{
			RelationField: field.NewRelation("AssignUser.Role", "models.Role"),
			Organization: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("AssignUser.Role.Organization", "models.Organization"),
			},
			Menus: struct {
				field.RelationField
				Parent struct {
					field.RelationField
				}
				Permission struct {
					field.RelationField
					Menu struct {
						field.RelationField
					}
				}
			}{
				RelationField: field.NewRelation("AssignUser.Role.Menus", "models.Menu"),
				Parent: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("AssignUser.Role.Menus.Parent", "models.Menu"),
				},
				Permission: struct {
					field.RelationField
					Menu struct {
						field.RelationField
					}
				}{
					RelationField: field.NewRelation("AssignUser.Role.Menus.Permission", "models.Permission"),
					Menu: struct {
						field.RelationField
					}{
						RelationField: field.NewRelation("AssignUser.Role.Menus.Permission.Menu", "models.Menu"),
					},
				},
			},
		},
		Organization: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("AssignUser.Organization", "models.Organization"),
		},
	}

	_alertActionLog.CreatedBy = alertActionLogBelongsToCreatedBy{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("CreatedBy", "models.User"),
	}

	_alertActionLog.Alert = alertActionLogManyToManyAlert{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("Alert", "models.Alert"),
		User: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Alert.User", "models.User"),
		},
		Site: struct {
			field.RelationField
			Organization struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Alert.Site", "models.Site"),
			Organization: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Alert.Site.Organization", "models.Organization"),
			},
		},
		Device: struct {
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
		}{
			RelationField: field.NewRelation("Alert.Device", "models.Device"),
			Rack: struct {
				field.RelationField
				Site struct {
					field.RelationField
				}
				Organization struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("Alert.Device.Rack", "models.Rack"),
				Site: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Alert.Device.Rack.Site", "models.Site"),
				},
				Organization: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Alert.Device.Rack.Organization", "models.Organization"),
				},
			},
			Template: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Alert.Device.Template", "models.Template"),
			},
			Site: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Alert.Device.Site", "models.Site"),
			},
			Organization: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Alert.Device.Organization", "models.Organization"),
			},
		},
		Ap: struct {
			field.RelationField
			Site struct {
				field.RelationField
			}
			Organization struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Alert.Ap", "models.AP"),
			Site: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Alert.Ap.Site", "models.Site"),
			},
			Organization: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Alert.Ap.Organization", "models.Organization"),
			},
		},
		Circuit: struct {
			field.RelationField
			Site struct {
				field.RelationField
			}
			Device struct {
				field.RelationField
			}
			DeviceInterface struct {
				field.RelationField
				Device struct {
					field.RelationField
				}
				Site struct {
					field.RelationField
				}
			}
			Organization struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Alert.Circuit", "models.Circuit"),
			Site: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Alert.Circuit.Site", "models.Site"),
			},
			Device: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Alert.Circuit.Device", "models.Device"),
			},
			DeviceInterface: struct {
				field.RelationField
				Device struct {
					field.RelationField
				}
				Site struct {
					field.RelationField
				}
			}{
				RelationField: field.NewRelation("Alert.Circuit.DeviceInterface", "models.DeviceInterface"),
				Device: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Alert.Circuit.DeviceInterface.Device", "models.Device"),
				},
				Site: struct {
					field.RelationField
				}{
					RelationField: field.NewRelation("Alert.Circuit.DeviceInterface.Site", "models.Site"),
				},
			},
			Organization: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Alert.Circuit.Organization", "models.Organization"),
			},
		},
		DeviceInterface: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Alert.DeviceInterface", "models.DeviceInterface"),
		},
		Maintenance: struct {
			field.RelationField
			Organization struct {
				field.RelationField
			}
			Alert struct {
				field.RelationField
			}
		}{
			RelationField: field.NewRelation("Alert.Maintenance", "models.Maintenance"),
			Organization: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Alert.Maintenance.Organization", "models.Organization"),
			},
			Alert: struct {
				field.RelationField
			}{
				RelationField: field.NewRelation("Alert.Maintenance.Alert", "models.Alert"),
			},
		},
		Organization: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("Alert.Organization", "models.Organization"),
		},
	}

	_alertActionLog.AlertGroup = alertActionLogManyToManyAlertGroup{
		db: db.Session(&gorm.Session{}),

		RelationField: field.NewRelation("AlertGroup", "models.AlertGroup"),
		Site: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("AlertGroup.Site", "models.Site"),
		},
		Organization: struct {
			field.RelationField
		}{
			RelationField: field.NewRelation("AlertGroup.Organization", "models.Organization"),
		},
	}

	_alertActionLog.fillFieldMap()

	return _alertActionLog
}

type alertActionLog struct {
	alertActionLogDo

	ALL          field.Asterisk
	Id           field.String
	CreatedAt    field.Time
	UpdatedAt    field.Time
	Acknowledged field.Bool
	Resolved     field.Bool
	Suppressed   field.Bool
	Comment      field.String
	AssignUserId field.String
	CreatedById  field.String
	AssignUser   alertActionLogBelongsToAssignUser

	CreatedBy alertActionLogBelongsToCreatedBy

	Alert alertActionLogManyToManyAlert

	AlertGroup alertActionLogManyToManyAlertGroup

	fieldMap map[string]field.Expr
}

func (a alertActionLog) Table(newTableName string) *alertActionLog {
	a.alertActionLogDo.UseTable(newTableName)
	return a.updateTableName(newTableName)
}

func (a alertActionLog) As(alias string) *alertActionLog {
	a.alertActionLogDo.DO = *(a.alertActionLogDo.As(alias).(*gen.DO))
	return a.updateTableName(alias)
}

func (a *alertActionLog) updateTableName(table string) *alertActionLog {
	a.ALL = field.NewAsterisk(table)
	a.Id = field.NewString(table, "id")
	a.CreatedAt = field.NewTime(table, "createdAt")
	a.UpdatedAt = field.NewTime(table, "updatedAt")
	a.Acknowledged = field.NewBool(table, "acknowledged")
	a.Resolved = field.NewBool(table, "resolved")
	a.Suppressed = field.NewBool(table, "suppressed")
	a.Comment = field.NewString(table, "comment")
	a.AssignUserId = field.NewString(table, "assignUserId")
	a.CreatedById = field.NewString(table, "createdById")

	a.fillFieldMap()

	return a
}

func (a *alertActionLog) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := a.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (a *alertActionLog) fillFieldMap() {
	a.fieldMap = make(map[string]field.Expr, 13)
	a.fieldMap["id"] = a.Id
	a.fieldMap["createdAt"] = a.CreatedAt
	a.fieldMap["updatedAt"] = a.UpdatedAt
	a.fieldMap["acknowledged"] = a.Acknowledged
	a.fieldMap["resolved"] = a.Resolved
	a.fieldMap["suppressed"] = a.Suppressed
	a.fieldMap["comment"] = a.Comment
	a.fieldMap["assignUserId"] = a.AssignUserId
	a.fieldMap["createdById"] = a.CreatedById

}

func (a alertActionLog) clone(db *gorm.DB) alertActionLog {
	a.alertActionLogDo.ReplaceConnPool(db.Statement.ConnPool)
	return a
}

func (a alertActionLog) replaceDB(db *gorm.DB) alertActionLog {
	a.alertActionLogDo.ReplaceDB(db)
	return a
}

type alertActionLogBelongsToAssignUser struct {
	db *gorm.DB

	field.RelationField

	Role struct {
		field.RelationField
		Organization struct {
			field.RelationField
		}
		Menus struct {
			field.RelationField
			Parent struct {
				field.RelationField
			}
			Permission struct {
				field.RelationField
				Menu struct {
					field.RelationField
				}
			}
		}
	}
	Organization struct {
		field.RelationField
	}
}

func (a alertActionLogBelongsToAssignUser) Where(conds ...field.Expr) *alertActionLogBelongsToAssignUser {
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

func (a alertActionLogBelongsToAssignUser) WithContext(ctx context.Context) *alertActionLogBelongsToAssignUser {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a alertActionLogBelongsToAssignUser) Session(session *gorm.Session) *alertActionLogBelongsToAssignUser {
	a.db = a.db.Session(session)
	return &a
}

func (a alertActionLogBelongsToAssignUser) Model(m *models.AlertActionLog) *alertActionLogBelongsToAssignUserTx {
	return &alertActionLogBelongsToAssignUserTx{a.db.Model(m).Association(a.Name())}
}

type alertActionLogBelongsToAssignUserTx struct{ tx *gorm.Association }

func (a alertActionLogBelongsToAssignUserTx) Find() (result *models.User, err error) {
	return result, a.tx.Find(&result)
}

func (a alertActionLogBelongsToAssignUserTx) Append(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a alertActionLogBelongsToAssignUserTx) Replace(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a alertActionLogBelongsToAssignUserTx) Delete(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a alertActionLogBelongsToAssignUserTx) Clear() error {
	return a.tx.Clear()
}

func (a alertActionLogBelongsToAssignUserTx) Count() int64 {
	return a.tx.Count()
}

type alertActionLogBelongsToCreatedBy struct {
	db *gorm.DB

	field.RelationField
}

func (a alertActionLogBelongsToCreatedBy) Where(conds ...field.Expr) *alertActionLogBelongsToCreatedBy {
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

func (a alertActionLogBelongsToCreatedBy) WithContext(ctx context.Context) *alertActionLogBelongsToCreatedBy {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a alertActionLogBelongsToCreatedBy) Session(session *gorm.Session) *alertActionLogBelongsToCreatedBy {
	a.db = a.db.Session(session)
	return &a
}

func (a alertActionLogBelongsToCreatedBy) Model(m *models.AlertActionLog) *alertActionLogBelongsToCreatedByTx {
	return &alertActionLogBelongsToCreatedByTx{a.db.Model(m).Association(a.Name())}
}

type alertActionLogBelongsToCreatedByTx struct{ tx *gorm.Association }

func (a alertActionLogBelongsToCreatedByTx) Find() (result *models.User, err error) {
	return result, a.tx.Find(&result)
}

func (a alertActionLogBelongsToCreatedByTx) Append(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a alertActionLogBelongsToCreatedByTx) Replace(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a alertActionLogBelongsToCreatedByTx) Delete(values ...*models.User) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a alertActionLogBelongsToCreatedByTx) Clear() error {
	return a.tx.Clear()
}

func (a alertActionLogBelongsToCreatedByTx) Count() int64 {
	return a.tx.Count()
}

type alertActionLogManyToManyAlert struct {
	db *gorm.DB

	field.RelationField

	User struct {
		field.RelationField
	}
	Site struct {
		field.RelationField
		Organization struct {
			field.RelationField
		}
	}
	Device struct {
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
	Ap struct {
		field.RelationField
		Site struct {
			field.RelationField
		}
		Organization struct {
			field.RelationField
		}
	}
	Circuit struct {
		field.RelationField
		Site struct {
			field.RelationField
		}
		Device struct {
			field.RelationField
		}
		DeviceInterface struct {
			field.RelationField
			Device struct {
				field.RelationField
			}
			Site struct {
				field.RelationField
			}
		}
		Organization struct {
			field.RelationField
		}
	}
	DeviceInterface struct {
		field.RelationField
	}
	Maintenance struct {
		field.RelationField
		Organization struct {
			field.RelationField
		}
		Alert struct {
			field.RelationField
		}
	}
	Organization struct {
		field.RelationField
	}
}

func (a alertActionLogManyToManyAlert) Where(conds ...field.Expr) *alertActionLogManyToManyAlert {
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

func (a alertActionLogManyToManyAlert) WithContext(ctx context.Context) *alertActionLogManyToManyAlert {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a alertActionLogManyToManyAlert) Session(session *gorm.Session) *alertActionLogManyToManyAlert {
	a.db = a.db.Session(session)
	return &a
}

func (a alertActionLogManyToManyAlert) Model(m *models.AlertActionLog) *alertActionLogManyToManyAlertTx {
	return &alertActionLogManyToManyAlertTx{a.db.Model(m).Association(a.Name())}
}

type alertActionLogManyToManyAlertTx struct{ tx *gorm.Association }

func (a alertActionLogManyToManyAlertTx) Find() (result []*models.Alert, err error) {
	return result, a.tx.Find(&result)
}

func (a alertActionLogManyToManyAlertTx) Append(values ...*models.Alert) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a alertActionLogManyToManyAlertTx) Replace(values ...*models.Alert) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a alertActionLogManyToManyAlertTx) Delete(values ...*models.Alert) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a alertActionLogManyToManyAlertTx) Clear() error {
	return a.tx.Clear()
}

func (a alertActionLogManyToManyAlertTx) Count() int64 {
	return a.tx.Count()
}

type alertActionLogManyToManyAlertGroup struct {
	db *gorm.DB

	field.RelationField

	Site struct {
		field.RelationField
	}
	Organization struct {
		field.RelationField
	}
}

func (a alertActionLogManyToManyAlertGroup) Where(conds ...field.Expr) *alertActionLogManyToManyAlertGroup {
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

func (a alertActionLogManyToManyAlertGroup) WithContext(ctx context.Context) *alertActionLogManyToManyAlertGroup {
	a.db = a.db.WithContext(ctx)
	return &a
}

func (a alertActionLogManyToManyAlertGroup) Session(session *gorm.Session) *alertActionLogManyToManyAlertGroup {
	a.db = a.db.Session(session)
	return &a
}

func (a alertActionLogManyToManyAlertGroup) Model(m *models.AlertActionLog) *alertActionLogManyToManyAlertGroupTx {
	return &alertActionLogManyToManyAlertGroupTx{a.db.Model(m).Association(a.Name())}
}

type alertActionLogManyToManyAlertGroupTx struct{ tx *gorm.Association }

func (a alertActionLogManyToManyAlertGroupTx) Find() (result []*models.AlertGroup, err error) {
	return result, a.tx.Find(&result)
}

func (a alertActionLogManyToManyAlertGroupTx) Append(values ...*models.AlertGroup) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Append(targetValues...)
}

func (a alertActionLogManyToManyAlertGroupTx) Replace(values ...*models.AlertGroup) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Replace(targetValues...)
}

func (a alertActionLogManyToManyAlertGroupTx) Delete(values ...*models.AlertGroup) (err error) {
	targetValues := make([]interface{}, len(values))
	for i, v := range values {
		targetValues[i] = v
	}
	return a.tx.Delete(targetValues...)
}

func (a alertActionLogManyToManyAlertGroupTx) Clear() error {
	return a.tx.Clear()
}

func (a alertActionLogManyToManyAlertGroupTx) Count() int64 {
	return a.tx.Count()
}

type alertActionLogDo struct{ gen.DO }

type IAlertActionLogDo interface {
	gen.SubQuery
	Debug() IAlertActionLogDo
	WithContext(ctx context.Context) IAlertActionLogDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IAlertActionLogDo
	WriteDB() IAlertActionLogDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IAlertActionLogDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IAlertActionLogDo
	Not(conds ...gen.Condition) IAlertActionLogDo
	Or(conds ...gen.Condition) IAlertActionLogDo
	Select(conds ...field.Expr) IAlertActionLogDo
	Where(conds ...gen.Condition) IAlertActionLogDo
	Order(conds ...field.Expr) IAlertActionLogDo
	Distinct(cols ...field.Expr) IAlertActionLogDo
	Omit(cols ...field.Expr) IAlertActionLogDo
	Join(table schema.Tabler, on ...field.Expr) IAlertActionLogDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IAlertActionLogDo
	RightJoin(table schema.Tabler, on ...field.Expr) IAlertActionLogDo
	Group(cols ...field.Expr) IAlertActionLogDo
	Having(conds ...gen.Condition) IAlertActionLogDo
	Limit(limit int) IAlertActionLogDo
	Offset(offset int) IAlertActionLogDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IAlertActionLogDo
	Unscoped() IAlertActionLogDo
	Create(values ...*models.AlertActionLog) error
	CreateInBatches(values []*models.AlertActionLog, batchSize int) error
	Save(values ...*models.AlertActionLog) error
	First() (*models.AlertActionLog, error)
	Take() (*models.AlertActionLog, error)
	Last() (*models.AlertActionLog, error)
	Find() ([]*models.AlertActionLog, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.AlertActionLog, err error)
	FindInBatches(result *[]*models.AlertActionLog, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*models.AlertActionLog) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IAlertActionLogDo
	Assign(attrs ...field.AssignExpr) IAlertActionLogDo
	Joins(fields ...field.RelationField) IAlertActionLogDo
	Preload(fields ...field.RelationField) IAlertActionLogDo
	FirstOrInit() (*models.AlertActionLog, error)
	FirstOrCreate() (*models.AlertActionLog, error)
	FindByPage(offset int, limit int) (result []*models.AlertActionLog, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IAlertActionLogDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (a alertActionLogDo) Debug() IAlertActionLogDo {
	return a.withDO(a.DO.Debug())
}

func (a alertActionLogDo) WithContext(ctx context.Context) IAlertActionLogDo {
	return a.withDO(a.DO.WithContext(ctx))
}

func (a alertActionLogDo) ReadDB() IAlertActionLogDo {
	return a.Clauses(dbresolver.Read)
}

func (a alertActionLogDo) WriteDB() IAlertActionLogDo {
	return a.Clauses(dbresolver.Write)
}

func (a alertActionLogDo) Session(config *gorm.Session) IAlertActionLogDo {
	return a.withDO(a.DO.Session(config))
}

func (a alertActionLogDo) Clauses(conds ...clause.Expression) IAlertActionLogDo {
	return a.withDO(a.DO.Clauses(conds...))
}

func (a alertActionLogDo) Returning(value interface{}, columns ...string) IAlertActionLogDo {
	return a.withDO(a.DO.Returning(value, columns...))
}

func (a alertActionLogDo) Not(conds ...gen.Condition) IAlertActionLogDo {
	return a.withDO(a.DO.Not(conds...))
}

func (a alertActionLogDo) Or(conds ...gen.Condition) IAlertActionLogDo {
	return a.withDO(a.DO.Or(conds...))
}

func (a alertActionLogDo) Select(conds ...field.Expr) IAlertActionLogDo {
	return a.withDO(a.DO.Select(conds...))
}

func (a alertActionLogDo) Where(conds ...gen.Condition) IAlertActionLogDo {
	return a.withDO(a.DO.Where(conds...))
}

func (a alertActionLogDo) Order(conds ...field.Expr) IAlertActionLogDo {
	return a.withDO(a.DO.Order(conds...))
}

func (a alertActionLogDo) Distinct(cols ...field.Expr) IAlertActionLogDo {
	return a.withDO(a.DO.Distinct(cols...))
}

func (a alertActionLogDo) Omit(cols ...field.Expr) IAlertActionLogDo {
	return a.withDO(a.DO.Omit(cols...))
}

func (a alertActionLogDo) Join(table schema.Tabler, on ...field.Expr) IAlertActionLogDo {
	return a.withDO(a.DO.Join(table, on...))
}

func (a alertActionLogDo) LeftJoin(table schema.Tabler, on ...field.Expr) IAlertActionLogDo {
	return a.withDO(a.DO.LeftJoin(table, on...))
}

func (a alertActionLogDo) RightJoin(table schema.Tabler, on ...field.Expr) IAlertActionLogDo {
	return a.withDO(a.DO.RightJoin(table, on...))
}

func (a alertActionLogDo) Group(cols ...field.Expr) IAlertActionLogDo {
	return a.withDO(a.DO.Group(cols...))
}

func (a alertActionLogDo) Having(conds ...gen.Condition) IAlertActionLogDo {
	return a.withDO(a.DO.Having(conds...))
}

func (a alertActionLogDo) Limit(limit int) IAlertActionLogDo {
	return a.withDO(a.DO.Limit(limit))
}

func (a alertActionLogDo) Offset(offset int) IAlertActionLogDo {
	return a.withDO(a.DO.Offset(offset))
}

func (a alertActionLogDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IAlertActionLogDo {
	return a.withDO(a.DO.Scopes(funcs...))
}

func (a alertActionLogDo) Unscoped() IAlertActionLogDo {
	return a.withDO(a.DO.Unscoped())
}

func (a alertActionLogDo) Create(values ...*models.AlertActionLog) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Create(values)
}

func (a alertActionLogDo) CreateInBatches(values []*models.AlertActionLog, batchSize int) error {
	return a.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (a alertActionLogDo) Save(values ...*models.AlertActionLog) error {
	if len(values) == 0 {
		return nil
	}
	return a.DO.Save(values)
}

func (a alertActionLogDo) First() (*models.AlertActionLog, error) {
	if result, err := a.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*models.AlertActionLog), nil
	}
}

func (a alertActionLogDo) Take() (*models.AlertActionLog, error) {
	if result, err := a.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*models.AlertActionLog), nil
	}
}

func (a alertActionLogDo) Last() (*models.AlertActionLog, error) {
	if result, err := a.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*models.AlertActionLog), nil
	}
}

func (a alertActionLogDo) Find() ([]*models.AlertActionLog, error) {
	result, err := a.DO.Find()
	return result.([]*models.AlertActionLog), err
}

func (a alertActionLogDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*models.AlertActionLog, err error) {
	buf := make([]*models.AlertActionLog, 0, batchSize)
	err = a.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (a alertActionLogDo) FindInBatches(result *[]*models.AlertActionLog, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return a.DO.FindInBatches(result, batchSize, fc)
}

func (a alertActionLogDo) Attrs(attrs ...field.AssignExpr) IAlertActionLogDo {
	return a.withDO(a.DO.Attrs(attrs...))
}

func (a alertActionLogDo) Assign(attrs ...field.AssignExpr) IAlertActionLogDo {
	return a.withDO(a.DO.Assign(attrs...))
}

func (a alertActionLogDo) Joins(fields ...field.RelationField) IAlertActionLogDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Joins(_f))
	}
	return &a
}

func (a alertActionLogDo) Preload(fields ...field.RelationField) IAlertActionLogDo {
	for _, _f := range fields {
		a = *a.withDO(a.DO.Preload(_f))
	}
	return &a
}

func (a alertActionLogDo) FirstOrInit() (*models.AlertActionLog, error) {
	if result, err := a.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*models.AlertActionLog), nil
	}
}

func (a alertActionLogDo) FirstOrCreate() (*models.AlertActionLog, error) {
	if result, err := a.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*models.AlertActionLog), nil
	}
}

func (a alertActionLogDo) FindByPage(offset int, limit int) (result []*models.AlertActionLog, count int64, err error) {
	result, err = a.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = a.Offset(-1).Limit(-1).Count()
	return
}

func (a alertActionLogDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = a.Count()
	if err != nil {
		return
	}

	err = a.Offset(offset).Limit(limit).Scan(result)
	return
}

func (a alertActionLogDo) Scan(result interface{}) (err error) {
	return a.DO.Scan(result)
}

func (a alertActionLogDo) Delete(models ...*models.AlertActionLog) (result gen.ResultInfo, err error) {
	return a.DO.Delete(models)
}

func (a *alertActionLogDo) withDO(do gen.Dao) *alertActionLogDo {
	a.DO = *do.(*gen.DO)
	return a
}
