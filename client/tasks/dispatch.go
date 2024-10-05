package tasks

import (
	"encoding/json"

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
	taskName := task["taskName"].(string)
	if taskName == intendtask.WebSSH {
		err = webSSHTask(data)
		if err != nil {
			logger.Logger.Error("taskDispatcher]: webSSHTask err: ", err)
		}
	} else {
		if _, ok := task["callback"]; !ok {
			logger.Logger.Error("taskDispatcher]: received wrong task: no callback")
		}
		taskId := task["taskId"].(string)
		switch taskName {
		case intendtask.ScanDeviceBasicInfo:
			results, err := scanDeviceBasicInfo(data)
			if err != nil {
				logger.Logger.Error("taskDispatcher]: scanDeviceBasicInfo err: ", err)
			}
			scanDeviceBasicInfoCallback(results, taskId)
		case intendtask.ScanDevice:
			results, err := scanDevice(data)
			if err != nil {
				logger.Logger.Error("taskDispatcher]: scanDevice err: ", err)
			}
			scanDeviceCallback(results, taskId)
		case intendtask.ScanAp:
			results, err := scanAp(data)
			if err != nil {
				logger.Logger.Error("taskDispatcher]: scanAp err: ", err)
			}
			scanApCallback(results, taskId)

		case intendtask.ScanMacAddressTable:
			results, err := scanMacAddressTable(data)
			if err != nil {
				logger.Logger.Error("taskDispatcher]: scanMacAddressTable err: ", err)
			}
			scanMacAddressTableCallback(results, taskId)
		}
	}

}
