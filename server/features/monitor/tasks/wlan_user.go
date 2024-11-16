package monitor_tasks

import (
	"github.com/wangxin688/narvis/intend/intendtask"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/infra"
	"github.com/wangxin688/narvis/server/models/ckmodel"
)

func WlanUserCallback(wlanUsers *intendtask.WlanUserTaskResult) error {
	users := make([]ckmodel.WlanStation, len(wlanUsers.WlanUsers))
	for index, user := range wlanUsers.WlanUsers {
		users[index] = ckmodel.WlanStation{
			StationUsername:         user.StationUsername,
			StationMac:              user.StationMac,
			StationIp:               user.StationIp,
			StationApMac:            *user.StationApMac,
			StationApName:           *user.StationApName,
			StationESSID:            user.StationESSID,
			StationChannel:          uint16(user.StationChannel),
			StationChannelBandWidth: *user.StationChanBandWidth,
			StationSNR:              uint16(*user.StationSNR),
			StationRSSI:             int16(user.StationRSSI),
			StationRxBits:           user.StationRxBits,
			StationTxBits:           user.StationTxBits,
			StationMaxSpeed:         uint32(*user.StationMaxSpeed),
			StationOnlineTime:       uint32(user.StationOnlineTime),
			SiteId:                  wlanUsers.SiteId,
			OrganizationId:          wlanUsers.DeviceId,
		}
	}
	err := infra.ClickHouseDB.CreateInBatches(users, 5000).Error
	if err != nil {
		core.Logger.Error("[WlanUserCallback] failed to batch create wlan user, error: %s[]")
		return err
	}
	return nil
}
