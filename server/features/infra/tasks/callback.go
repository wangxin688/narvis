package infra_tasks

import (
	"github.com/samber/lo"
	"github.com/wangxin688/narvis/intend/intendtask"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/dal/gen"
	infra_biz "github.com/wangxin688/narvis/server/features/infra/biz"
	"github.com/wangxin688/narvis/server/models"
	"go.uber.org/zap"
)

func DeviceBasicInfoScanCallback(scanResults []*intendtask.DeviceBasicInfoScanResponse) error {
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
		if _, ok := dbDevices[scanResult.ManagementIp]; !ok {
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

func ScanApCallback(scanResults []*intendtask.ApScanResponse) error {
	managementIPs := lo.Map(scanResults, func(v *intendtask.ApScanResponse, _ int) string {
		return v.ManagementIp
	})
	if len(managementIPs) == 0 {
		core.Logger.Info("[apScanCallback]: no scan aps found")
		return nil
	}
	apService := infra_biz.NewApService()
	dbAps, err := apService.GetByIpsAndSiteId(managementIPs, scanResults[0].SiteId, scanResults[0].OrganizationId)
	if err != nil {
		core.Logger.Error("[apScanCallback]: get db aps failed", zap.Error(err))
		return err
	}
	var newAps []*models.AP
	var updateAps []*models.AP
	for _, scanResult := range scanResults {
		ap, ok := dbAps[scanResult.ManagementIp]
		if !ok {
			newAps = append(newAps, &models.AP{
				OrganizationId:  scanResult.OrganizationId,
				Status:          "Active",
				ManagementIp:    scanResult.ManagementIp,
				Name:            scanResult.Name,
				SerialNumber:    scanResult.SerialNumber,
				DeviceModel:     scanResult.DeviceModel,
				Manufacturer:    scanResult.Manufacturer,
				GroupName:       scanResult.GroupName,
				WlanACIpAddress: scanResult.WlanACIpAddress,
				MacAddress:      scanResult.MacAddress,
				SiteId:          scanResult.SiteId,
			})
		} else if apService.CalApHash(ap) != scanResult.CalApHash() {
			ap.GroupName = scanResult.GroupName
			ap.Name = scanResult.Name
			ap.MacAddress = scanResult.MacAddress
			ap.SerialNumber = scanResult.SerialNumber
			ap.DeviceModel = scanResult.DeviceModel
			ap.Manufacturer = scanResult.Manufacturer
			ap.WlanACIpAddress = scanResult.WlanACIpAddress
			updateAps = append(updateAps, ap)
		}
	}
	if len(newAps) > 0 {
		err = gen.AP.CreateInBatches(newAps, len(newAps))
		if err != nil {
			core.Logger.Error("[apScanCallback]: failed to create new ap", zap.Error(err))
			return err
		}
	}
	if len(updateAps) > 0 {
		err = gen.AP.Save(updateAps...)
		if err != nil {
			core.Logger.Error("[apScanCallback]: failed to update existed ap", zap.Error(err))
			return err
		}
	}
	return nil
}
