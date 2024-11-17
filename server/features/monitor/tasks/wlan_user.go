package monitor_tasks

import (
	"fmt"
	"time"

	"github.com/wangxin688/narvis/intend/intendtask"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/infra"
	"github.com/wangxin688/narvis/server/models/ckmodel"
	"go.uber.org/zap"
)

func WlanUserCallback(wlanUsers *intendtask.WlanUserTaskResult) error {
	if wlanUsers == nil {
		core.Logger.Error("[WlanUserCallback]: wlanUsers is nil")
		return fmt.Errorf("wlanUsers is nil")
	}

	numberOfUsers := len(wlanUsers.WlanUsers)
	core.Logger.Info("[WlanUserCallback]: received wlan user callback", zap.Int("number of users", numberOfUsers))

	if numberOfUsers == 0 {
		core.Logger.Warn("[WlanUserCallback]: no users in the callback")
		return nil
	}

	users := make([]ckmodel.WlanStation, numberOfUsers)
	now := time.Now()
	for index, user := range wlanUsers.WlanUsers {
		if user == nil {
			core.Logger.Warn("[WlanUserCallback]: encountered nil user, skipping")
			continue
		}

		var apMac, apName, chanBandWidth string
		var snr, maxSpeed uint64

		if user.StationApMac != nil {
			apMac = *user.StationApMac
		}
		if user.StationApName != nil {
			apName = *user.StationApName
		}
		if user.StationChanBandWidth != nil {
			chanBandWidth = *user.StationChanBandWidth
		}
		if user.StationSNR != nil {
			snr = *user.StationSNR
		}
		if user.StationMaxSpeed != nil {
			maxSpeed = *user.StationMaxSpeed
		}

		users[index] = ckmodel.WlanStation{
			Ts:                      now,
			StationUsername:         user.StationUsername,
			StationMac:              user.StationMac,
			StationIp:               user.StationIp,
			StationApMac:            apMac,
			StationApName:           apName,
			StationESSID:            user.StationESSID,
			StationChannel:          uint16(user.StationChannel),
			StationChannelBandWidth: chanBandWidth,
			StationSNR:              uint16(snr),
			StationRSSI:             int16(user.StationRSSI),
			StationRxBits:           user.StationRxBits,
			StationTxBits:           user.StationTxBits,
			StationMaxSpeed:         uint32(maxSpeed),
			StationOnlineTime:       uint32(user.StationOnlineTime),
			SiteId:                  wlanUsers.SiteId,
			OrganizationId:          wlanUsers.DeviceId,
		}
	}

	err := infra.ClickHouseDB.CreateInBatches(users, 5000).Error
	if err != nil {
		core.Logger.Error("[WlanUserCallback] failed to batch create wlan user", zap.Error(err))
		return err
	}

	return nil
}
