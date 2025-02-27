package infra_tasks

import (
	"strings"
	"time"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/intend/helpers/bgtask"
	"github.com/wangxin688/narvis/intend/intendtask"
	"github.com/wangxin688/narvis/intend/logger"
	intend_device "github.com/wangxin688/narvis/intend/model/device"
	"github.com/wangxin688/narvis/server/dal/gen"
	infra_biz "github.com/wangxin688/narvis/server/features/infra/biz"
	ipam_utils "github.com/wangxin688/narvis/server/features/ipam/utils"
	"github.com/wangxin688/narvis/server/infra"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/tools/errors"
	"github.com/wangxin688/narvis/server/tools/helpers"
	"go.uber.org/zap"
	"golang.org/x/exp/rand"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func DeviceBasicInfoScanCallback(scanResults []*intendtask.DeviceBasicInfoScanResponse) error {
	managementIPs := lo.Map(scanResults, func(v *intendtask.DeviceBasicInfoScanResponse, _ int) string {
		return v.ManagementIp
	})
	if len(managementIPs) == 0 {
		logger.Logger.Info("[deviceBasicInfoScanCallback]: no scan devices found")
		return nil
	}
	dbDevices, err := infra_biz.NewScanDeviceService().GetByScanResult(managementIPs, scanResults[0].OrganizationId)
	if err != nil {
		logger.Logger.Error("[deviceBasicInfoScanCallback]: get db scan devices failed", zap.Error(err))
		return err
	}
	infraDevices, err := infra_biz.NewDeviceService().GetDeviceByManagementIp(managementIPs, scanResults[0].OrganizationId)
	if err != nil {
		logger.Logger.Error("[deviceBasicInfoScanCallback]: get infra scan devices failed", zap.Error(err))
		return err
	}

	// find new devices
	var newDevices []*models.ScanDevice
	for _, scanResult := range scanResults {
		if _, ok := dbDevices[scanResult.ManagementIp]; !ok {
			// if device is already registered in devices table, ignore it.
			if _, ok := infraDevices[scanResult.ManagementIp]; ok {
				continue
			}
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
			logger.Logger.Error("[deviceBasicInfoScanCallback]: create scan devices failed", zap.Error(err))
			return err
		}
	}

	return nil
}

func DeviceScanCallback(data *intendtask.DeviceScanResponse) error {
	device, err := gen.Device.Where(
		gen.Device.Id.Eq(data.DeviceId),
		gen.Device.OrganizationId.Eq(data.OrganizationId),
	).First()
	if err != nil {
		logger.Logger.Error("[deviceScanCallback]: get db device failed", zap.Error(err))
		return err
	}
	err = deviceUpdateHandler(data, device)
	if err != nil {
		logger.Logger.Error("[deviceScanCallback.deviceUpdate]: update db device failed", zap.Error(err))
		return err
	}
	if len(data.Interfaces) > 0 {
		err = interfacesCallbackHandler(data.DeviceId, data.SiteId, data.Interfaces)
		if err != nil {
			logger.Logger.Error("[deviceScanCallback.interfaces]: update db device failed", zap.Error(err))
		}
	}
	if len(data.LldpNeighbors) > 0 {
		err = lldpCallbackHandler(data.DeviceId, data.SiteId, data.OrganizationId, data.LldpNeighbors)
		if err != nil {
			logger.Logger.Error("[deviceScanCallback.lldp]: update db device failed", zap.Error(err))
		}
	}
	if len(data.Vlans) > 0 {
		err = vlanCallbackHandler(data.DeviceId, data.SiteId, data.OrganizationId, data.Vlans)
		if err != nil {
			logger.Logger.Error("[deviceScanCallback.vlans]: update db device failed", zap.Error(err))
		}
	}

	err = lldpCallbackHandler(data.DeviceId, data.SiteId, data.OrganizationId, data.LldpNeighbors)
	if err != nil {
		logger.Logger.Error("[deviceScanCallback.lldp]: update db device failed", zap.Error(err))
	}
	if len(data.ArpTable) > 0 {
		// due to arp table may have huge size, we do it in background task
		bgtask.BackgroundTask(func() {
			arpTableCallbackHandler(data.DeviceId, data.SiteId, data.OrganizationId, data.ArpTable) //nolint: errcheck
		})
	}

	return nil
}

func lldpCallbackHandler(deviceId, siteId, orgId string, data []*intend_device.LldpNeighbor) error {
	if len(data) == 0 {
		logger.Logger.Info("[deviceScanCallback.lldp]: no lldp neighbors found", zap.String("deviceId", deviceId))
		return nil
	}
	remoteHostNames := make([]string, 0)
	remoteChassisIds := lo.Map(data, func(v *intend_device.LldpNeighbor, _ int) string {
		if v.RemoteChassisId == "" {
			remoteHostNames = append(remoteHostNames, v.RemoteHostname)
		}
		return v.RemoteChassisId
	})
	remoteChassisIds = lo.Filter(remoteChassisIds, func(v string, _ int) bool {
		return v != ""
	})
	remoteDevices, err := infra_biz.NewDeviceService().GetDeviceByChassisIds(remoteChassisIds, orgId)
	if err != nil {
		logger.Logger.Error("[deviceScanCallback.lldp]: get devices failed by chassis ids", zap.Error(err))
		return err
	}
	remoteAps, err := infra_biz.NewApService().GetApByMacAddresses(remoteChassisIds, orgId)
	if err != nil {
		logger.Logger.Error("[deviceScanCallback.lldp]: get aps failed by chassis ids", zap.Error(err))
		return err
	}
	lldpService := infra_biz.NewLldpNeighborService()
	lldpDeviceNeighbors, err := lldpService.GetDeviceLldpNeighbors(deviceId)
	if err != nil {
		logger.Logger.Error("[deviceScanCallback.lldp]: get device lldp neighbors failed", zap.Error(err))
		return err
	}
	lldpApNeighbors, err := lldpService.GetApLldpNeighbors(deviceId)
	if err != nil {
		logger.Logger.Error("[deviceScanCallback.lldp]: get ap lldp neighbors failed", zap.Error(err))
		return err
	}
	createDeviceLldp := make([]*models.LLDPNeighbor, 0)
	createApLldp := make([]*models.ApLLDPNeighbor, 0)
	for _, lldp := range data {
		remoteChassisId := lldp.RemoteChassisId
		if remoteChassisId == "" {
			logger.Logger.Info("[deviceScanCallback.lldp]: remote chassis id is empty", zap.Any("lldpInfo", lldp))
		}
		if _, ok := remoteDevices[remoteChassisId]; !ok {
			if _, ok := remoteAps[remoteChassisId]; !ok {
				logger.Logger.Warn("[deviceScanCallback.lldp]: remote device not found in device/ap table", zap.Any("lldp", lldp))
				continue
			}
		}
		if _, ok := remoteDevices[lldp.RemoteChassisId]; ok {
			lldp.HashValue = lldp.CalHashValue()
			if _, ok := lldpDeviceNeighbors[lldp.HashValue]; !ok {
				createDeviceLldp = append(createDeviceLldp, &models.LLDPNeighbor{
					LocalDeviceId:  deviceId,
					LocalIfName:    lldp.LocalIfName,
					LocalIfDescr:   lldp.LocalIfDescr,
					RemoteDeviceId: remoteDevices[remoteChassisId].Id,
					RemoteIfName:   lldp.RemoteIfName,
					RemoteIfDescr:  lldp.RemoteIfDescr,
					SiteId:         siteId,
					OrganizationId: orgId,
					HashValue:      lldp.HashValue,
				})
			}
		} else if _, ok := remoteAps[remoteChassisId]; ok {
			lldp.HashValue = lldp.CalHashValue()
			if _, ok := lldpApNeighbors[lldp.HashValue]; !ok {
				createApLldp = append(createApLldp, &models.ApLLDPNeighbor{
					LocalDeviceId:  deviceId,
					LocalIfName:    lldp.LocalIfName,
					LocalIfDescr:   lldp.LocalIfDescr,
					RemoteApId:     remoteAps[remoteChassisId].Id,
					HashValue:      lldp.HashValue,
					SiteId:         siteId,
					OrganizationId: orgId,
				})
			}
		} else {
			logger.Logger.Warn("[deviceScanCallback.lldp]: lldp neighbor not found in device/ap table", zap.Any("lldp", lldp))
			continue
		}
	}
	if dLen := len(createDeviceLldp); dLen > 0 {
		err = gen.LLDPNeighbor.UnderlyingDB().Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "localDeviceId"}, {Name: "localIfName"}},
			UpdateAll: true,
		}).CreateInBatches(createDeviceLldp, dLen).Error
		if err != nil {
			logger.Logger.Error("[deviceScanCallback.lldp]: failed to insert device lldp neighbors", zap.Error(err))
			return err
		}
	}
	if aLen := len(createApLldp); aLen > 0 {
		err = gen.ApLLDPNeighbor.UnderlyingDB().Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "localDeviceId"}, {Name: "localIfName"}},
			UpdateAll: true,
		}).CreateInBatches(createApLldp, aLen).Error
		if err != nil {
			logger.Logger.Error("[deviceScanCallback.lldp]: failed to insert ap lldp neighbors", zap.Error(err))
			return err
		}
	}
	return nil
}

func vlanCallbackHandler(deviceId, siteId, orgId string, data []*intend_device.VlanItem) error {
	if len(data) <= 0 {
		logger.Logger.Info("[deviceScanCallback.vlans]: no vlans found in device", zap.String("deviceId", deviceId))
		return nil
	}
	createPrefixes := make([]*models.Prefix, 0)
	setOfRange := make([]string, 0)
	for _, vlan := range data {
		if vlan.Network == "" || lo.Contains(setOfRange, vlan.Network) {
			continue
		}
		setOfRange = append(setOfRange, vlan.Network)
		gateway := ipam_utils.TrimGatewayMask(vlan.Gateway)
		createPrefixes = append(createPrefixes, &models.Prefix{
			SiteId:         siteId,
			OrganizationId: orgId,
			Range:          vlan.Network,
			Version:        ipam_utils.CidrVersion(vlan.Network),
			Gateway:        &gateway,
			VlanId:         &vlan.VlanId,
			VlanName:       &vlan.VlanName,
		})
	}
	if len(createPrefixes) > 0 {
		err := gen.Prefix.UnderlyingDB().Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "vlanId"}, {Name: "siteId"}},
			UpdateAll: true,
		}).CreateInBatches(createPrefixes, len(createPrefixes)).Error
		if err != nil {
			logger.Logger.Error("[deviceScanCallback.vlans]: failed to insert prefixes", zap.Error(err))
			return err
		}
	}
	return nil
}

func arpTableCallbackHandler(deviceId, siteId, orgId string, data []*intend_device.ArpItem) error {
	if len(data) <= 0 {
		logger.Logger.Info("[deviceScanCallback.arp]: no arp found in device", zap.String("deviceId", deviceId))
		return nil
	}
	createArps := make([]*models.IpAddress, 0)
	for _, arp := range data {
		if arp.Range == "" {
			continue
		}
		createArps = append(createArps, &models.IpAddress{
			SiteId:         siteId,
			OrganizationId: orgId,
			Address:        arp.IpAddress,
			MacAddress:     &arp.MacAddress,
			Vlan:           &arp.VlanId,
			Range:          &arp.Range,
			Type:           arp.Type,
		})
	}
	if len(createArps) > 0 {
		// add jitter to avoid race condition
		jitter := time.Duration(rand.Intn(500)) * time.Millisecond
		time.Sleep(jitter)
		err := gen.IpAddress.UnderlyingDB().Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "address"}, {Name: "siteId"}},
			UpdateAll: true,
		}).CreateInBatches(createArps, 50).Error
		if err != nil {
			logger.Logger.Error("[deviceScanCallback.arp]: failed to insert arp to ipAddress table", zap.Error(err))
			return err
		}
	}
	return nil
}
func interfacesCallbackHandler(deviceId string, siteId string, data []*intend_device.DeviceInterface) error {
	interfaces, err := infra_biz.NewDeviceInterfaceService().GetDeviceInterfaces(deviceId)
	if err != nil {
		logger.Logger.Error("[deviceScanCallback.interfaces]: get device interfaces failed", zap.String("deviceId", deviceId), zap.Error(err))
	}
	createdInterfaces := make([]*models.DeviceInterface, 0)
	for _, df := range data {
		if _, ok := interfaces[df.HashValue]; ok {
			continue
		}
		createdInterfaces = append(createdInterfaces, &models.DeviceInterface{
			DeviceId:      deviceId,
			SiteId:        siteId,
			IfIndex:       df.IfIndex,
			IfName:        df.IfName,
			IfDescr:       df.IfDescr,
			IfSpeed:       df.IfSpeed,
			IfType:        df.IfType,
			IfMtu:         df.IfMtu,
			IfAdminStatus: df.IfAdminStatus,
			IfOperStatus:  df.IfOperStatus,
			IfLastChange:  df.IfLastChange,
			IfHighSpeed:   df.IfHighSpeed,
			IfPhysAddr:    df.IfPhysAddr,
			IfIpAddress:   df.IfIpAddress,
		})
	}

	if len(createdInterfaces) > 0 {
		err = gen.DeviceInterface.UnderlyingDB().Clauses(clause.OnConflict{
			Columns:   []clause.Column{{Name: "ifIndex"}, {Name: "deviceId"}},
			UpdateAll: true,
		}).CreateInBatches(createdInterfaces, len(createdInterfaces)).Error
		if err != nil {
			logger.Logger.Error("[deviceScanCallback.interfaces]: failed to insert device interfaces", zap.Error(err))
			return err
		}
	}

	return nil
}

func deviceUpdateHandler(data *intendtask.DeviceScanResponse, device *models.Device) error {
	if device.Name != data.Name {
		device.Name = data.Name
	}
	if device.ChassisId != data.ChassisId {
		device.ChassisId = data.ChassisId
	}
	if device.Manufacturer != data.Manufacturer {
		device.Manufacturer = data.Manufacturer
	}
	if device.DeviceModel != data.DeviceModel {
		device.DeviceModel = data.DeviceModel
	}
	if device.Platform != data.Platform {
		device.Platform = data.Platform
	}
	if len(data.Entities) > 0 {
		device.OsVersion = &data.Entities[0].EntityPhysicalSoftwareRev
		var serNum *string
		serNums := lo.Map(data.Entities, func(v *intend_device.Entity, _ int) string {
			return v.EntityPhysicalSerialNum
		})
		if len(serNums) == 1 {
			serNum = &serNums[0]
		} else if len(serNums) > 1 {
			sr := strings.Join(serNums, ",")
			serNum = &sr
		}
		device.SerialNumber = serNum
	}
	return gen.Device.Save(device)
}

func ScanApCallback(scanResults []*intendtask.ApScanResponse) error {
	if len(scanResults) == 0 {
		logger.Logger.Info("[apScanCallback]: no scan aps found")
		return nil
	}
	ok, err := infra_biz.LicenseUsageDepends(1, scanResults[0].OrganizationId)
	if err != nil {
		return err
	}
	if !ok {
		return errors.NewError(errors.CodeLicenseCountExceeded, errors.MsgLicenseCountExceeded)
	}
	managementIPs := lo.Map(scanResults, func(v *intendtask.ApScanResponse, _ int) string {
		return v.ManagementIp
	})
	if len(managementIPs) == 0 {
		logger.Logger.Info("[apScanCallback]: no scan aps found")
		return nil
	}
	apService := infra_biz.NewApService()
	dbAps, err := apService.GetByIpsAndSiteId(managementIPs, scanResults[0].SiteId, scanResults[0].OrganizationId)
	if err != nil {
		logger.Logger.Error("[apScanCallback]: get db aps failed", zap.Error(err))
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
				OsVersion:       scanResult.OsVersion,
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
	tx := infra.DB.Session(&gorm.Session{SkipHooks: true})
	if len(newAps) > 0 {
		err = tx.Clauses(
			clause.OnConflict{
				Columns:   []clause.Column{{Name: "managementIp"}, {Name: "siteId"}},
				UpdateAll: true,
			},
		).CreateInBatches(newAps, len(newAps)).Error
		if err != nil {
			logger.Logger.Error("[apScanCallback]: failed to create new ap", zap.Error(err))
			return err
		}
	}
	if len(updateAps) > 0 {
		for _, ap := range updateAps {
			if err = tx.Save(ap).Error; err != nil {
				logger.Logger.Error("[apScanCallback]: failed to update existed ap", zap.Error(err))
				return err
			}
		}
	}
	return nil
}

func ConfigBackUpCallback(data *intendtask.ConfigurationBackupTaskResult) error {
	configService := infra_biz.NewDeviceConfigService()
	lastConfig, err := configService.GetLatestDeviceConfigByDeviceId(data.DeviceId)
	if err != nil {
		return err
	}
	totalLines := uint32(infra_biz.GetConfigTotalLines(data.Configuration))
	hashValue := helpers.StringToMd5(data.Configuration)
	if lastConfig == nil {
		newBackupConfig := &models.DeviceConfig{
			Configuration: data.Configuration,
			DeviceId:      data.DeviceId,
			TotalLines:    totalLines,
			LinesAdded:    totalLines,
			LinesDeleted:  0,
			Md5Checksum:   hashValue,
		}
		err = gen.DeviceConfig.Create(newBackupConfig)
		if err != nil {
			return err
		}
	} else {

		if lastConfig.Md5Checksum == hashValue {
			return nil
		}
		added, deleted := infra_biz.GetLineChanges(lastConfig.Configuration, data.Configuration)
		newBackupConfig := &models.DeviceConfig{
			Configuration: data.Configuration,
			DeviceId:      data.DeviceId,
			TotalLines:    totalLines,
			LinesAdded:    uint32(added),
			LinesDeleted:  uint32(deleted),
			Md5Checksum:   hashValue,
		}
		err = gen.DeviceConfig.Create(newBackupConfig)
		if err != nil {
			return err
		}
	}
	return configService.HouseKeeping(data.DeviceId)
}
