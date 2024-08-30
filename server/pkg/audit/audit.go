package audit

import (
	"encoding/json"
	"reflect"

	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const AuditCreateAction = "Create"
const AuditUpdateAction = "Update"
const AuditDeleteAction = "Delete"

type AuditLogMixin struct {
	auditTableMap map[string]struct{}
}

func NewAuditLogMixin() *AuditLogMixin {
	return &AuditLogMixin{
		auditTableMap: make(map[string]struct{}),
	}
}

func (a *AuditLogMixin) AuditTableRegister(tableName []string) *AuditLogMixin {
	for _, v := range tableName {
		if v != models.AuditLogTableName {
			a.auditTableMap[v] = struct{}{}
		}
		a.auditTableMap[v] = struct{}{}
	}
	return a
}

func (a *AuditLogMixin) afterCreate(tx *gorm.DB) {
	if !a.enableAuditing(tx) || tx.RowsAffected == 0 {
		return
	}
	target, err := getDBObjectBeforeOperation(tx)
	if err != nil {
		return
	}
	a.createAuditLog(tx, target, AuditCreateAction)
}

func (a *AuditLogMixin) afterUpdate(tx *gorm.DB) {
	if !a.enableAuditing(tx) || tx.RowsAffected == 0 {
		return
	}
	target, err := getDBObjectBeforeOperation(tx)
	if err != nil {
		return
	}
	a.createAuditLog(tx, target, AuditUpdateAction)

}

func (a *AuditLogMixin) beforeDelete(tx *gorm.DB) {
	if !a.enableAuditing(tx) || tx.RowsAffected == 0 {
		return
	}
	target, err := getDBObjectBeforeOperation(tx)
	if err != nil {
		return
	}
	a.createAuditLog(tx, target, AuditDeleteAction)
}

func (a *AuditLogMixin) createAuditLog(tx *gorm.DB, target map[string]any, action string) {
	requestId := global.XRequestId.Get()
	userId := global.UserId.Get()
	auditLog := &models.AuditLog{
		ObjectType:     tx.Statement.Schema.Table,
		ObjectId:       getKeyFromMap("Id", target),
		RequestId:      &requestId,
		UserId:         &userId,
		Action:         action,
		Data:           prepareData(target),
		OrganizationId: global.OrganizationId.Get(),
	}
	if err := tx.Session(&gorm.Session{SkipHooks: true, NewDB: true}).Create(auditLog).Error; err != nil {
		core.Logger.Error("AuditLog.createAuditLog commit create audit log error", zap.Error(err))
	}

}

func (a *AuditLogMixin) enableAuditing(tx *gorm.DB) bool {
	if tx.DryRun || tx.Error != nil {
		return false
	}
	if _, ok := a.auditTableMap[tx.Statement.Table]; !ok {
		return false
	}
	if tx.Statement.Schema.ModelType == nil {
		return false
	}
	return true
}

func getDBObjectBeforeOperation(tx *gorm.DB) (map[string]any, error) {
	if tx.DryRun || tx.Error != nil {
		return nil, nil
	}
	objMap := make(map[string]any)
	objType := reflect.TypeOf(tx.Statement.ReflectValue.Interface())
	target := reflect.New(objType).Interface()

	primaryKeyValue := getPkKeyValue(tx.Statement.ReflectValue)
	if primaryKeyValue == "" {
		return nil, nil
	}

	if err := tx.Session(&gorm.Session{}).Table(tx.Statement.Schema.Table).Where("id = ?", primaryKeyValue).Find(target).Error; err != nil {
		core.Logger.Error("AuditLog.getDBObjectBeforeOperation get target error", zap.Error(err))
		return nil, err
	}
	jsonBytes, err := json.Marshal(target)
	if err != nil {
		core.Logger.Error("AuditLog.getDBObjectBeforeOperation json.Marshal error", zap.Error(err))
		return nil, err
	}
	if err := json.Unmarshal(jsonBytes, &objMap); err != nil {
		core.Logger.Error("AuditLog.getDBObjectBeforeOperation json.Unmarshal error", zap.Error(err))
		return nil, err
	}
	return objMap, nil
}

func (a *AuditLogMixin) RegisterCallbacks(tx *gorm.DB) {
	tx.Callback().Create().After("gorm:create").Register("audit:after_create", a.afterCreate)
	tx.Callback().Update().After("gorm:update").Register("audit:after_update", a.afterUpdate)
	tx.Callback().Delete().Before("gorm:delete").Register("audit:before_delete", a.beforeDelete)
}
