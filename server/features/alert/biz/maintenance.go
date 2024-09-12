package alert_biz

import (
	"time"

	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/models"
)

type MaintenanceService struct{}

func NewMaintenanceService() *MaintenanceService {
	return &MaintenanceService{}
}

func (m *MaintenanceService) GetActiveMaintenance() ([]*models.Maintenance, error) {
	timeNow := time.Now().UTC()
	return gen.Maintenance.Where(
		gen.Maintenance.StartedAt.Lte(timeNow),
		gen.Maintenance.EndedAt.Gt(timeNow),
	).Find()
}
