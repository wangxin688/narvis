package tasks

import (
	"encoding/json"

	"github.com/wangxin688/narvis/client/utils/helpers"
	"github.com/wangxin688/narvis/client/utils/logger"
	"github.com/wangxin688/narvis/intend/intendtask"
)

func TaskDispatcher(data []byte) {

	task := make(map[string]any)

	err := json.Unmarshal(data, &task)
	if err != nil {
		logger.Logger.Error("taskDispatcher]: received wrong task: unmarshal to map err: ", err)
	}
	if _, ok := task["taskName"]; !ok {
		logger.Logger.Error("taskDispatcher]: received wrong task: no taskName")
	}
	if _, ok := task["callback"]; !ok {
		logger.Logger.Error("taskDispatcher]: received wrong task: no callback")
	}
	taskName := task["taskName"].(string)
	switch taskName {
	case intendtask.ScanDeviceBasicInfo:
		helpers.BackgroundTask(func() {
			results, err := scanDeviceBasicInfo(data)
			if err != nil {
				logger.Logger.Error("taskDispatcher]: scanDeviceBasicInfo err: ", err)
			}
			scanDeviceBasicInfoCallback(results)
		})
	case intendtask.ScanDevice:
		helpers.BackgroundTask(func() {
			results, err := scanDevice(data)
			if err != nil {
				logger.Logger.Error("taskDispatcher]: scanDevice err: ", err)
			}
			scanDeviceCallback(results)
		})
	case intendtask.ScanAp:
		helpers.BackgroundTask(func() {
			results, err := scanAp(data)
			if err != nil {
				logger.Logger.Error("taskDispatcher]: scanAp err: ", err)
			}
			scanApCallback(results)
		})
	case intendtask.ScanMacAddressTable:
		helpers.BackgroundTask(func() {
			results, err := scanMacAddressTable(data)
			if err != nil {
				logger.Logger.Error("taskDispatcher]: scanMacAddressTable err: ", err)
			}
			scanMacAddressTableCallback(results)
		})

	}

}
