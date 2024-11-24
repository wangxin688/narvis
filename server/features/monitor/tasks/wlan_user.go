package monitor_tasks

import (
	"fmt"

	"github.com/wangxin688/narvis/intend/intendtask"
	"github.com/wangxin688/narvis/intend/logger"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/models"
	"go.uber.org/zap"
)

func WlanUserCallback(wlanUsers *intendtask.WlanUserTaskResult) error {
	if wlanUsers == nil {
		logger.Logger.Error("[WlanUserCallback]: wlanUsers is nil")
		return fmt.Errorf("wlanUsers is nil")
	}

	numberOfUsers := len(wlanUsers.WlanUsers)
	logger.Logger.Info("[WlanUserCallback]: received wlan user callback", zap.Int("number of users", numberOfUsers))

	if numberOfUsers == 0 {
		logger.Logger.Warn("[WlanUserCallback]: no users in the callback")
		return nil
	}

	users := make([]*models.WlanStation, numberOfUsers)
	for index, user := range wlanUsers.WlanUsers {
		if user == nil {
			logger.Logger.Warn("[WlanUserCallback]: encountered nil user, skipping")
			continue
		}
		users[index] = &models.WlanStation{
			StationUsername:      user.StationUsername,
			StationMac:           user.StationMac,
			StationIp:            user.StationIp,
			StationApMac:         user.StationApMac,
			StationApName:        user.StationApName,
			StationESSID:         user.StationESSID,
			StationChannel:       user.StationChannel,
			StationChanBandWidth: user.StationChanBandWidth,
			StationSNR:           user.StationSNR,
			StationRSSI:          user.StationRSSI,
			StationRxBits:        user.StationRxBits,
			StationTxBits:        user.StationTxBits,
			StationMaxSpeed:      user.StationMaxSpeed,
			StationOnlineTime:    user.StationOnlineTime,
			SiteId:               wlanUsers.SiteId,
			OrganizationId:       wlanUsers.DeviceId,
		}
	}

	err := gen.WlanStation.CreateInBatches(users, 5000)
	if err != nil {
		logger.Logger.Error("[WlanUserCallback] failed to batch create wlan user", zap.Error(err))
		return err
	}

	return nil
}
