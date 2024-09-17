package alert_biz

import "github.com/wangxin688/narvis/server/features/alert/schemas"

type ActionLogService struct{}

func NewActionLogService() *ActionLogService {
	return &ActionLogService{}
}

func (a *ActionLogService) CreateActionLog(log *schemas.ActionLogCreate) (string, error) {

}
