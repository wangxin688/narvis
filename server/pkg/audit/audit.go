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

// AuditLogMixin provides audit log for create, update, delete
// must notice that update and delete should load object from db first and change/delete action will be recorded.
// otherwise, it will not be recorded
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
	if !a.enableAuditing(tx) || tx.RowsAffected <= 0 {
		return
	}
	target, err := getDBObjectBeforeOperation(tx)
	if err != nil {
		return
	}
	a.createAuditLog(tx, target, AuditCreateAction)
}

func (a *AuditLogMixin) afterUpdate(tx *gorm.DB) {
	if !a.enableAuditing(tx) || tx.RowsAffected <= 0 {
		return
	}
	target, err := getDBObjectBeforeOperation(tx)
	if err != nil {
		return
	}
	a.createAuditLog(tx, target, AuditUpdateAction)

}

func (a *AuditLogMixin) afterDelete(tx *gorm.DB) {
	if !a.enableAuditing(tx) || tx.RowsAffected <= 0 {
		return
	}
	target, err := getDBObjectBeforeOperation(tx)
	if err != nil {
		return
	}
	a.createAuditLog(tx, target, AuditDeleteAction)
}

func (a *AuditLogMixin) createAuditLog(tx *gorm.DB, target *snapshot, action string) {
	requestId := global.XRequestId.Get()
	if requestId == "" {
		return
	}
	userId := global.UserId.Get()
	var auditLogs = make([]*models.AuditLog, 0)
	for pkId, data := range target.data {
		if pkId == "" {
			continue
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			core.Logger.Error("[createAuditLog]: json marshal error", zap.Error(err))
			continue
		}
		if action != AuditUpdateAction {
			auditLogs = append(auditLogs, &models.AuditLog{
				ObjectType:     tx.Statement.Schema.Table,
				ObjectId:       pkId,
				RequestId:      &requestId,
				UserId:         &userId,
				Action:         action,
				Data:           jsonData,
				OrganizationId: global.OrganizationId.Get(),
			})
		} else {
			diff := global.OrmDiff.Get()[pkId]
			jsonDiff, err := json.Marshal(diff)
			if err != nil {
				core.Logger.Error("[createAuditLog]: json marshal error", zap.Error(err))
				continue
			}
			auditLogs = append(auditLogs, &models.AuditLog{
				ObjectType:     tx.Statement.Schema.Table,
				ObjectId:       pkId,
				RequestId:      &requestId,
				UserId:         &userId,
				Action:         action,
				Data:           jsonData,
				Diff:           jsonDiff,
				OrganizationId: global.OrganizationId.Get(),
			})
		}
	}
	if len(auditLogs) == 0 {
		return
	}
	auditLog := auditLogs
	if err := tx.Session(&gorm.Session{SkipHooks: true, NewDB: true}).Create(auditLog).Error; err != nil {
		core.Logger.Error("[createAuditLog]: commit create audit log error", zap.Error(err))
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

func getDBObjectBeforeOperation(tx *gorm.DB) (*snapshot, error) {
	if tx.DryRun || tx.Error != nil {
		return nil, nil //nolint:nilerr
	}
	var targetObj interface{}
	if getItemType(tx.Statement.ReflectValue.Type()) == tx.Statement.Schema.ModelType {
		targetObj = tx.Statement.ReflectValue.Interface()
	} else {
		primaryKeyValue := getPkKeyValue(tx.Statement.ReflectValue)
		if len(primaryKeyValue) == 0 {
			return nil, nil
		}
		target := reflect.New(reflect.SliceOf(tx.Statement.Schema.ModelType)).Interface()
		targetObj = target
		if err := tx.Session(&gorm.Session{}).Table(tx.Statement.Schema.Table).Where("id IN ?", primaryKeyValue).Find(target).Error; err != nil {
			core.Logger.Error("[createAuditLog]: get target error before operation", zap.Error(err))
			return nil, err
		}
	}
	s, err := getSnapshot(targetObj, tx.Statement.Schema.Fields)
	if err != nil {
		core.Logger.Error("[createAuditLog]: get snapshot error before operation", zap.Error(err))
		return nil, err
	}
	return s, nil

}

func (a *AuditLogMixin) RegisterCallbacks(tx *gorm.DB) {
	tx.Callback().Create().After("gorm:create").Register("audit:after_create", a.afterCreate)
	tx.Callback().Update().After("gorm:update").Register("audit:after_update", a.afterUpdate)
	tx.Callback().Delete().After("gorm:delete").Register("audit:after_delete", a.afterDelete)
}
