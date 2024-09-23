package tasks

import (
	"sync"
	"time"

	"github.com/imroc/req/v3"
	"github.com/wangxin688/narvis/client/config"
	"github.com/wangxin688/narvis/client/utils/logger"
	"github.com/wangxin688/narvis/client/utils/security"
	"github.com/wangxin688/narvis/intend/intendtask"
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
		logger.Logger.Error("[scanDeviceBasicInfoCallback]: failed to create server")
		return
	}
	resp, err := server.R().SetBody(data).SetHeader(xTaskID, taskId).Post(intendtask.DeviceBasicInfoCbUrl)
	if err != nil {
		logger.Logger.Error("[scanDeviceBasicInfoCallback]: failed to post", err)
		return
	}
	if !resp.IsSuccessState() {
		logger.Logger.Error("[scanDeviceBasicInfoCallback]: failed to post result to server", resp.Status)
	}
	logger.Logger.Info("[scanDeviceBasicInfoCallback] post result to server success")
}

func scanDeviceCallback(data []*intendtask.DeviceScanResponse, taskId string) {
	server := newServer()
	if server == nil {
		logger.Logger.Error("[scanDeviceCallback]: failed to create server")
		return
	}
	resp, err := server.R().SetBody(data).SetHeader(xTaskID, taskId).Post(intendtask.DeviceCbUrl)
	if err != nil {
		logger.Logger.Error("[scanDeviceCallback]: failed to post", err)
		return
	}
	if !resp.IsSuccessState() {
		logger.Logger.Error("[scanDeviceCallback]: failed to post result to server", resp.Status)
	}

	logger.Logger.Info("[scanDeviceCallback] post result to server success")
}

func scanApCallback(data []*intendtask.ApScanResponse, taskId string) {
	server := newServer()
	if server == nil {
		logger.Logger.Error("[scanApCallback]: failed to create server")
		return
	}

	resp, err := server.R().SetBody(data).SetHeader(xTaskID, taskId).Post(intendtask.ApCbUrl)
	if err != nil {
		logger.Logger.Error("[scanApCallback]: failed to post", err)
		return
	}

	if !resp.IsSuccessState() {
		logger.Logger.Error("[scanApCallback]: failed to post result to server", resp.Status)
	}

	logger.Logger.Info("[scanApCallback] post result to server success")
}

func scanMacAddressTableCallback(data []*intendtask.MacAddressTableScanResponse, taskId string) {
	server := newServer()
	if server == nil {
		logger.Logger.Error("[ScanMacAddressTableCallback]: failed to create server")
		return
	}

	resp, err := server.R().SetBody(data).SetHeader(xTaskID, taskId).Post(intendtask.MacAddressTableCbUrl)
	if err != nil {
		logger.Logger.Error("[ScanMacAddressTableCallback]: failed to post", err)
		return
	}

	if !resp.IsSuccessState() {
		logger.Logger.Error("[ScanMacAddressTableCallback]: failed to post result to server", resp.Status)
	}

	logger.Logger.Info("[ScanMacAddressTableCallback] post result to server success")
}

func newServer() *req.Client {
	once.Do(func() {
		token, err := security.ProxyToken(config.Settings.PROXY_ID, config.Settings.SECRET_KEY)
		if err != nil {
			logger.Logger.Error("[taskCallback]: failed to generate proxy token", err)
		}
		server = req.C().SetBaseURL(config.Settings.SERVER_URL).
			SetCommonContentType("application/json").
			SetCommonBearerAuthToken(token).
			SetCommonRetryCount(2).
			SetTimeout(time.Duration(5) * time.Second)
	})

	return server
}
