package tasks

import (
	"fmt"
	"sync"
	"time"

	"github.com/imroc/req/v3"
	"github.com/wangxin688/narvis/client/config"
	"github.com/wangxin688/narvis/client/utils/logger"
	"github.com/wangxin688/narvis/client/utils/security"
	"github.com/wangxin688/narvis/intend/intendtask"
	"go.uber.org/zap"
)

var once sync.Once
var server *req.Client
var xTaskID = "X-Task-ID"

type ServerResponse struct {
	TaskId string `json:"taskId"`
}

func scanDeviceBasicInfoCallback(data []*intendtask.DeviceBasicInfoScanResponse, taskId string) {
	server := newServer()
	if server == nil {
		logger.Logger.Error(fmt.Sprintf("[scanDeviceBasicInfoCallback] [%s]: failed to create server", taskId))
		return
	}
	resp, err := server.R().SetBody(data).SetHeader(xTaskID, taskId).Post(intendtask.DeviceBasicInfoCbUrl)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[scanDeviceBasicInfoCallback] [%s] : failed to post", taskId), zap.Error(err))
		logger.Logger.Error(fmt.Sprintf("[scanDeviceBasicInfoCallback] [%s] : server response", taskId), zap.Any("response", resp.String()))
		return
	}
	if !resp.IsSuccessState() {
		logger.Logger.Error(fmt.Sprintf("[scanDeviceBasicInfoCallback] [%s] : failed to post result to server", taskId), zap.String("response", resp.String()))
	}
	logger.Logger.Info(fmt.Sprintf("[scanDeviceBasicInfoCallback] [%s] : post result to server success", taskId))
}

func scanDeviceCallback(data *intendtask.DeviceScanResponse, taskId string) {
	server := newServer()
	if server == nil {
		logger.Logger.Error(fmt.Sprintf("[scanDeviceCallback] [%s]: failed to create server", taskId))
		return
	}
	resp, err := server.R().SetBody(data).SetHeader(xTaskID, taskId).Post(intendtask.DeviceCbUrl)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[scanDeviceCallback] [%s]: failed to post", taskId), zap.Error(err))
		logger.Logger.Error(fmt.Sprintf("[scanDeviceCallback] [%s]: server response", taskId), zap.Any("response", resp.String()))
		return
	}
	if !resp.IsSuccessState() {
		logger.Logger.Error(fmt.Sprintf("[scanDeviceCallback] [%s]: failed to post result to server", taskId), zap.String("response", resp.String()))
		logger.Logger.Info(fmt.Sprintf("[scanDeviceCallback] [%s]: ", taskId), zap.Any("response", resp.String()))
	}

	logger.Logger.Info(fmt.Sprintf("[scanDeviceCallback] [%s]: post result to server success", taskId))
}

func scanApCallback(data []*intendtask.ApScanResponse, taskId string) {
	server := newServer()
	if server == nil {
		logger.Logger.Error(fmt.Sprintf("[scanApCallback] [%s]: failed to create server", taskId))
		return
	}

	resp, err := server.R().SetBody(data).SetHeader(xTaskID, taskId).Post(intendtask.ApCbUrl)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[scanApCallback] [%s]: failed to post", taskId), zap.Error(err))
		logger.Logger.Error(fmt.Sprintf("[scanApCallback] [%s]: server response", taskId), zap.Any("response", resp.String()))
		return
	}

	if !resp.IsSuccessState() {
		logger.Logger.Error(fmt.Sprintf("[scanApCallback] [%s]: failed to post result to server", taskId), zap.String("response", resp.Status))
		logger.Logger.Info(fmt.Sprintf("[scanApCallback] [%s]: ", taskId), zap.Any("response", resp.String()))
	}

	logger.Logger.Info(fmt.Sprintf("[scanApCallback] [%s]: post result to server success", taskId))
}

func scanMacAddressTableCallback(data []*intendtask.MacAddressTableScanResponse, taskId string) {
	server := newServer()
	if server == nil {
		logger.Logger.Error(fmt.Sprintf("[ScanMacAddressTableCallback] [%s]: failed to create server", taskId))
		return
	}

	resp, err := server.R().SetBody(data).SetHeader(xTaskID, taskId).Post(intendtask.MacAddressTableCbUrl)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[ScanMacAddressTableCallback] [%s]: failed to post", taskId), zap.Error(err))
		return
	}

	if !resp.IsSuccessState() {
		logger.Logger.Error(fmt.Sprintf("[ScanMacAddressTableCallback] [%s]: failed to post result to server", taskId), zap.String("response", resp.Status))
		logger.Logger.Info(fmt.Sprintf("[ScanMacAddressTableCallback] [%s]: ", taskId), zap.Any("response", resp.String()))
	}

	logger.Logger.Info(fmt.Sprintf("[ScanMacAddressTableCallback] [%s]: post result to server success", taskId))
}

func configBackUpCallback(data *intendtask.ConfigurationBackupTaskResult, taskId string) {
	server := newServer()
	if server == nil {
		logger.Logger.Error(fmt.Sprintf("[configBackUpCallback] [%s]: failed to create server", taskId))
		return
	}
	resp, err := server.R().SetBody(data).SetHeader(xTaskID, taskId).Post(intendtask.ConfigurationBackupCbUrl)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[configBackUpCallback] [%s]: failed to post", taskId), zap.Error(err))
		logger.Logger.Error(fmt.Sprintf("[configBackUpCallback] [%s]: server response", taskId), zap.Any("response", resp.String()))
		return
	}
	if !resp.IsSuccessState() {
		logger.Logger.Error(fmt.Sprintf("[configBackUpCallback] [%s]: failed to post result to server", taskId), zap.String("response", resp.String()))
	}
	logger.Logger.Info(fmt.Sprintf("[configBackUpCallback] [%s]: post result to server success", taskId))
}

func wlanUsersCallback(data *intendtask.WlanUserTaskResult, taskId string) {
	server := newServer()
	if server == nil {
		logger.Logger.Error(fmt.Sprintf("[wlanUsersCallback] [%s]: failed to create server", taskId))
		return
	}
	resp, err := server.R().SetBody(data).SetHeader(xTaskID, taskId).Post(intendtask.WlanUserCbUrl)
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[wlanUsersCallback] [%s]: failed to post", taskId), zap.Error(err))
		logger.Logger.Error(fmt.Sprintf("[wlanUsersCallback] [%s]: server response", taskId), zap.Any("response", resp.String()))
		return
	}
	if !resp.IsSuccessState() {
		logger.Logger.Error(fmt.Sprintf("[wlanUsersCallback] [%s]: failed to post result to server", taskId), zap.String("response", resp.String()))
	}
	logger.Logger.Info(fmt.Sprintf("[wlanUsersCallback] [%s]: post result to server success", taskId))
}

func newServer() *req.Client {
	once.Do(func() {
		token, err := security.ProxyToken(config.Settings.PROXY_ID, config.Settings.SECRET_KEY)
		if err != nil {
			logger.Logger.Error("[taskCallback]: failed to generate proxy token", zap.Error(err))
		}
		server = req.C().SetBaseURL(config.Settings.SERVER_URL).
			SetCommonContentType("application/json").
			SetCommonBearerAuthToken(token).
			SetCommonRetryCount(2).
			SetTimeout(time.Duration(5) * time.Second)
	})

	return server
}
