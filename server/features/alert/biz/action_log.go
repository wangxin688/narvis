package alert_biz

import (
	"github.com/samber/lo"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/features/alert/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/tools/errors"
)

type ActionLogService struct{}

func NewActionLogService() *ActionLogService {
	return &ActionLogService{}
}

func (a *ActionLogService) CreateActionLog(log *schemas.ActionLogCreate) (string, error) {
	orgId := global.OrganizationId.Get()
	if log.AssignUserId != nil {
		_, err := gen.User.Select(gen.User.Id).Where(
			gen.User.Id.Eq(*log.AssignUserId),
			gen.User.OrganizationId.Eq(orgId),
		).First()
		if err != nil {
			return "", err
		}
	}
	alerts, err := gen.Alert.Where(gen.Alert.Id.In(log.AlertId...), gen.Alert.OrganizationId.Eq(orgId)).Find()
	if err != nil {
		return "", err
	}
	alertIds := lo.Map(alerts, func(alert *models.Alert, _ int) string {
		return alert.Id
	})
	if len(alertIds) == 0 {
		return "", errors.NewError(errors.CodeNotFound, errors.MsgNotFound, models.AlertTableName, "id", log.AlertId)
	}
	log.AlertId = alertIds
	newActions := make([]*models.AlertActionLog, len(alertIds))
	for i, alertId := range alertIds {
		newActions[i] = &models.AlertActionLog{
			AlertId:      alertId,
			AssignUserId: log.AssignUserId,
			Comment:      log.Comment,
			RootCauseId:  log.RootCauseId,
			Resolved:     log.Resolved,
			Suppressed:   log.Suppressed,
			Acknowledged: log.Acknowledged,
			CreatedById:  global.UserId.Get(),
		}
	}

	err = gen.AlertActionLog.CreateInBatches(newActions, len(newActions))
	if err != nil {
		return "", err
	}
	return "", nil
}
