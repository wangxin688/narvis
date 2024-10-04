package webssh_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/wangxin688/narvis/server/core"
	infra_biz "github.com/wangxin688/narvis/server/features/infra/biz"
	webssh_biz "github.com/wangxin688/narvis/server/features/webssh/biz"
	"go.uber.org/zap"
)

func handleWebSSHRequest(c *gin.Context) error {
	upGrader := websocket.Upgrader{
		// cross origin domain
		CheckOrigin: func(r *http.Request) bool {
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
		return err
	}
	defer wsConn.Close()
	deviceId := c.Param("deviceId")
	deviceConnectionInfo, err := infra_biz.NewCliCredentialService().GetCredentialByDeviceId(deviceId)
	if err != nil {
		return err
	}
	managementIP, err := infra_biz.NewDeviceService().GetManagementIP(deviceId)
	if err != nil {
		return err
	}
	sessionId := uuid.New().String()

	webssh_biz.AddSession(sessionId)
	webssh_biz.SendSignalToProxy(sessionId, managementIP, deviceConnectionInfo)
	proxyWSConn, err := webssh_biz.WaitForProxyWebSocket(sessionId)
	if err != nil {
		core.Logger.Error("[webssh]: failed to wait for proxy websocket", zap.Error(err))
		return err
	}
	webssh_biz.RelaySSHData(wsConn, proxyWSConn)
	return nil
}
