package vector

import (
	"github.com/imroc/req/v3"
	"github.com/wangxin688/narvis/client/pkg/nettysnmp/factory"
	"github.com/wangxin688/narvis/client/utils/logger"
	"go.uber.org/zap"
)

func PostWlanUsers(wlanUsers []*factory.WlanUser) {
	client := req.C()

	resp, err := client.R().SetBody(wlanUsers).Post("http://127.0.0.1:12056")
	if err != nil {
		logger.Logger.Error("[postWlanUsersToVector]: failed to post result", zap.Error(err))
	}
	logger.Logger.Info("[postWlanUsersToVector]: response: ", zap.String("response", resp.String()))

}
