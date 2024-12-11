package infra_biz

import (
	"github.com/wangxin688/narvis/intend/logger"
	"github.com/wangxin688/narvis/server/dal/gen"
	"go.uber.org/zap"
)

type IsolationService struct{}

func NewIsolationService() *IsolationService {
	return &IsolationService{}
}

func (i *IsolationService) CheckSiteNotFound(siteId, orgId string) error {
	_, err := gen.Site.Select(gen.Site.Id).Where(gen.Site.Id.Eq(siteId), gen.Site.OrganizationId.Eq(orgId)).First()
	if err != nil {
		logger.Logger.Warn(
			"[isolationChecking]: attacking may happened, siteId not found under Org",
			zap.String("siteId", siteId),
			zap.String("orgId", orgId))
		return err
	}
	return nil
}

func (i *IsolationService) CheckDeviceNotFound(deviceId, orgId string) error {
	_, err := gen.Device.Select(gen.Device.Id).Where(gen.Device.Id.Eq(deviceId), gen.Device.OrganizationId.Eq(orgId)).First()
	if err != nil {
		logger.Logger.Warn(
			"[isolationChecking]: attacking may happened, deviceId not found under Org",
			zap.String("deviceId", deviceId),
			zap.String("orgId", orgId))
		return err
	}
	return nil
}

func (i *IsolationService) CheckRackNotFound(rackId, orgId string) error {
	_, err := gen.Rack.Select(gen.Rack.Id).Where(gen.Rack.Id.Eq(rackId), gen.Rack.OrganizationId.Eq(orgId)).First()
	if err != nil {
		logger.Logger.Warn(
			"[isolationChecking]: attacking may happened, rackId not found under Org",
			zap.String("rackId", rackId),
			zap.String("orgId", orgId))
		return err
	}
	return nil
}

func (i *IsolationService) CheckDeviceInterfaceNotFound(interfaceId, orgId string) error {
	iface, err := gen.DeviceInterface.Select(gen.DeviceInterface.Id, gen.DeviceInterface.SiteId).Where(gen.DeviceInterface.Id.Eq(interfaceId)).First()
	if err != nil {
		return err
	}
	err = i.CheckSiteNotFound(iface.SiteId, orgId)
	if err != nil {
		logger.Logger.Warn(
			"[isolationChecking]: attacking may happened, interfaceId not found under Org",
			zap.String("interfaceId", interfaceId),
			zap.String("orgId", orgId))
		return err
	}
	return nil
}

func (i *IsolationService) CheckServerNotFound(serverId, orgId string) error {
	_, err := gen.Server.Select(gen.Server.Id).Where(gen.Server.Id.Eq(serverId), gen.Server.OrganizationId.Eq(orgId)).First()
	if err != nil {
		logger.Logger.Warn(
			"[isolationChecking]: attacking may happened, serverId not found under Org",
			zap.String("serverId", serverId),
			zap.String("orgId", orgId))
		return err
	}
	return nil
}
