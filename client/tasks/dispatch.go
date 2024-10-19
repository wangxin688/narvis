package tasks

import (
	"encoding/json"

	"github.com/wangxin688/narvis/client/utils/logger"
	"github.com/wangxin688/narvis/intend/intendtask"
	"go.uber.org/zap"
)

// TODO: optimize error handling return to server, add to struct rather than return error
func TaskDispatcher(data []byte) {

	task := make(map[string]any)

	err := json.Unmarshal(data, &task)
	if err != nil {
		logger.Logger.Error("taskDispatcher]: received wrong task: unmarshal to map err: ", zap.Error(err))
	}
	if _, ok := task["taskName"]; !ok {
		logger.Logger.Error("taskDispatcher]: received wrong task: no taskName")
	}
	taskName := task["taskName"].(string)
	if taskName == intendtask.WebSSH {
		err = webSSHTask(data)
		if err != nil {
			logger.Logger.Error("taskDispatcher]: webSSHTask err: ", zap.Error(err))
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
				logger.Logger.Error("taskDispatcher]: scanDeviceBasicInfo err: ", zap.Error(err))
			}
			scanDeviceBasicInfoCallback(results, taskId)
		case intendtask.ScanDevice:
			results, err := scanDevice(data)
			if err != nil {
				logger.Logger.Error("taskDispatcher]: scanDevice err: ", zap.Error(err))
			}
			scanDeviceCallback(results, taskId)
		case intendtask.ScanAp:
			results, err := scanAp(data)
			if err != nil {
				logger.Logger.Error("taskDispatcher]: scanAp err: ", zap.Error(err))
			}
			scanApCallback(results, taskId)

		case intendtask.ScanMacAddressTable:
			results, err := scanMacAddressTable(data)
			if err != nil {
				logger.Logger.Error("taskDispatcher]: scanMacAddressTable err: ", zap.Error(err))
			}
			scanMacAddressTableCallback(results, taskId)
		case intendtask.ConfigurationBackup:
			results := configurationBackupTask(data)
			if err != nil {
				logger.Logger.Error("taskDispatcher]: configurationBackupTask err: ", zap.Error(err))
			}
			configBackUpCallback(results, taskId)
		}
	}

}
