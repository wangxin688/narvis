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
	var scanResults []*intendtask.DeviceBasicInfoScanResponse
	if err := json.Unmarshal(data, &scanResults); err != nil {
		return err
	}

	managementIPs := lo.Map(scanResults, func(v *intendtask.DeviceBasicInfoScanResponse, _ int) string {
		return v.ManagementIp
	})
	if len(managementIPs) == 0 {
		core.Logger.Info("[deviceBasicInfoScanCallback]: no scan devices found")
		return nil
	}
	dbDevices, err := infra_biz.NewScanDeviceService().GetByScanResult(managementIPs, scanResults[0].OrganizationId)
	if err != nil {
		core.Logger.Error("[deviceBasicInfoScanCallback]: get db scan devices failed", zap.Error(err))
		return err
	}

	// find new devices
	var newDevices []*models.ScanDevice
	for _, scanResult := range scanResults {
		if _, ok := (*dbDevices)[scanResult.ManagementIp]; !ok {
			newDevices = append(newDevices, &models.ScanDevice{
				OrganizationId: scanResult.OrganizationId,
				ManagementIp:   scanResult.ManagementIp,
				Platform:       scanResult.Platform,
				Manufacturer:   scanResult.Manufacturer,
				DeviceModel:    scanResult.DeviceModel,
				ChassisId:      scanResult.ChassisId,
				Name:           scanResult.Name,
				Description:    scanResult.Description,
			})
		}
	}

	if len(newDevices) > 0 {
		err = gen.ScanDevice.CreateInBatches(newDevices, len(newDevices))
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
