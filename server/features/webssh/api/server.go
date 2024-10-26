package webssh_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/wangxin688/narvis/server/core"
	infra_biz "github.com/wangxin688/narvis/server/features/infra/biz"
	webssh_biz "github.com/wangxin688/narvis/server/features/webssh/biz"
	"github.com/wangxin688/narvis/server/tools/errors"
	"go.uber.org/zap"
)

func handleWebSSHRequest(c *gin.Context) error {
	upGrader := websocket.Upgrader{
		// cross origin domain
		CheckOrigin: func(_ *http.Request) bool {
			return true
		},
		// 处理 Sec-WebSocket-Protocol Header
		//Subprotocols: []string{r.Header.Get("Sec-WebSocket-Protocol")},
		Subprotocols:    []string{"webssh"},
		ReadBufferSize:  8192,
		WriteBufferSize: 8192,
	}
	wsConn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		core.Logger.Error("[webssh]: failed to upgrade", zap.Error(err))
		return errors.NewError(errors.CodeWebSocketInitFail, errors.MsgWebSocketInitFail, err)
	}
	defer wsConn.Close()
	deviceId := c.Param("deviceId")
	var query WebSSHServer
	if err = c.ShouldBindQuery(&query); err != nil {
		return err
	}
	if query.Cols == 0 {
		query.Cols = 80
	}
	if query.Rows == 0 {
		query.Rows = 40
	}
	deviceConnectionInfo, err := infra_biz.NewCliCredentialService().GetCredentialByDeviceId(deviceId)
	if err != nil {
		return err
	}
	managementIP, err := infra_biz.NewDeviceService().GetManagementIP(deviceId)
	if err != nil {
		return err
	}
	sessionId := uuid.New().String()

	err = webssh_biz.SendSignalToProxy(sessionId, managementIP, deviceConnectionInfo, query.Cols, query.Rows)
	if err != nil {
		core.Logger.Error("[webssh]: failed to send signal to proxy", zap.Error(err))
		return err
	}
	proxyWSConn, err := webssh_biz.WaitForProxyWebSocket(sessionId)
	if err != nil {
		core.Logger.Error("[webssh]: failed to wait for proxy websocket", zap.Error(err))
		return err
	}
	webssh_biz.RelaySSHData(wsConn, proxyWSConn)
	return nil
}
