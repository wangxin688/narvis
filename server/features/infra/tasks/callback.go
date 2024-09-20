package infra_tasks

import (
	"encoding/json"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/intend/intendtask"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/dal/gen"
	infra_biz "github.com/wangxin688/narvis/server/features/infra/biz"
	"github.com/wangxin688/narvis/server/models"
	"go.uber.org/zap"
)

func deviceBasicInfoScanCallback(data []byte) error {
	basicInfo := make([]*intendtask.DeviceBasicInfoScanResponse, 0)
	err := json.Unmarshal(data, &basicInfo)
	if err != nil {
		return err
	}
	ips := lo.Map(basicInfo, func(v *intendtask.DeviceBasicInfoScanResponse, _ int) string {
		return v.ManagementIp
	})
	if len(ips) == 0 {
		core.Logger.Info("[deviceBasicInfoScanCallback]: no scan devices found")
		return nil
	}
	dbDevices, err := infra_biz.NewScanDeviceService().GetByScanResult(ips, basicInfo[0].OrganizationId)
	if err != nil {
		core.Logger.Error("[deviceBasicInfoScanCallback]: get db scan devices failed", zap.Error(err))
		return err
	}
	addedDevices := make([]*models.ScanDevice, 0)
	for _, basicInfo := range basicInfo {
		if _, ok := (*dbDevices)[basicInfo.ManagementIp]; !ok {
			addedDevices = append(addedDevices, &models.ScanDevice{
				OrganizationId: basicInfo.OrganizationId,
				ManagementIp:   basicInfo.ManagementIp,
				Platform:       basicInfo.Platform,
				Manufacturer:   basicInfo.Manufacturer,
				DeviceModel:    basicInfo.DeviceModel,
				ChassisId:      basicInfo.ChassisId,
				Name:           basicInfo.Name,
				Description:    basicInfo.Description,
			})
		}
	}

	if len(addedDevices) > 0 {
		err = gen.ScanDevice.CreateInBatches(addedDevices, len(addedDevices))
		if err != nil {
			core.Logger.Error("[deviceBasicInfoScanCallback]: create scan devices failed", zap.Error(err))
			return err
		}
	}
	return nil
}

func deviceScanCallback(data []byte) error {
	return nil
}

func apScanCallback(data []byte) error {
	return nil
}
