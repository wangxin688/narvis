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

type ServerResponse struct {
	TaskId string `json:"taskId"`
}

func scanDeviceBasicInfoCallback(data []*intendtask.DeviceBasicInfoScanResponse) {
	server := newServer()
	if server == nil {
		logger.Logger.Error("[scanDeviceBasicInfoCallback]: failed to create server")
		return
	}
	resp, err := server.R().SetBody(data).Post(intendtask.DeviceBasicInfoCbUrl)
	if err != nil {
		logger.Logger.Error("[scanDeviceBasicInfoCallback]: failed to post", err)
		return
	}
	if !resp.IsSuccessState() {
		logger.Logger.Error("[scanDeviceBasicInfoCallback]: failed to post result to server", resp.Status)
	}
}

func scanDeviceCallback(data []*intendtask.DeviceScanResponse) {
	server := newServer()
	if server == nil {
		logger.Logger.Error("[scanDeviceCallback]: failed to create server")
		return
	}
	resp, err := server.R().SetBody(data).Post(intendtask.DeviceCbUrl)
	if err != nil {
		logger.Logger.Error("[scanDeviceCallback]: failed to post", err)
		return
	}
	if !resp.IsSuccessState() {
		logger.Logger.Error("[scanDeviceCallback]: failed to post result to server", resp.Status)
	}
}

func scanApCallback(data []*intendtask.ApScanResponse) {
	server := newServer()
	if server == nil {
		logger.Logger.Error("[scanApCallback]: failed to create server")
		return
	}

	resp, err := server.R().SetBody(data).Post(intendtask.ApCbUrl)
	if err != nil {
		logger.Logger.Error("[scanApCallback]: failed to post", err)
		return
	}

	if !resp.IsSuccessState() {
		logger.Logger.Error("[scanApCallback]: failed to post result to server", resp.Status)
	}
}

func scanMacAddressTableCallback(data []*intendtask.MacAddressTableScanResponse) {
	server := newServer()
	if server == nil {
		logger.Logger.Error("[ScanMacAddressTableCallback]: failed to create server")
		return
	}

	resp, err := server.R().SetBody(data).Post(intendtask.MacAddressTableCbUrl)
	if err != nil {
		logger.Logger.Error("[ScanMacAddressTableCallback]: failed to post", err)
		return
	}

	if !resp.IsSuccessState() {
		logger.Logger.Error("[ScanMacAddressTableCallback]: failed to post result to server", resp.Status)
	}
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
