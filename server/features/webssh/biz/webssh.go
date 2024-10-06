package webssh_biz

import (
	"encoding/json"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/wangxin688/narvis/intend/intendtask"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/features/infra/schemas"
	"github.com/wangxin688/narvis/server/global"
	"github.com/wangxin688/narvis/server/pkg/rmq"
	"go.uber.org/zap"
)

var SessionWMap sync.Map

func AddSession(sessionId string, ws chan *websocket.Conn) {
	if _, ok := SessionWMap.Load(sessionId); ok {
		return
	}
	SessionWMap.Store(sessionId, ws)
}
func SendSignalToProxy(sessionId, managementIP string, cred *schemas.CliCredential, cols, rows int) error {
	signal := intendtask.WebSSHTask{
		TaskName:     intendtask.WebSSH,
		ManagementIP: managementIP,
		SessionId:    sessionId,
		Username:     cred.Username,
		Password:     cred.Password,
		Port:         cred.Port,
		Cols:         cols,
		Rows:         rows,
	}
	signalByte, err := json.Marshal(signal)
	if err != nil {
		return err
	}
	err = rmq.PublishProxyMessage(signalByte, global.OrganizationId.Get())
	if err != nil {
		core.Logger.Error("[webssh.proxySignal]: failed to publish message to rabbitmq", zap.Error(err))
		return err
	}
	return nil
}

func WaitForProxyWebSocket(sessionId string) (*websocket.Conn, error) {
	done := make(chan *websocket.Conn)
	SessionWMap.Store(sessionId, done)
	proxyWS := <-done
	SessionWMap.Delete(sessionId)
	return proxyWS, nil
}

func RelaySSHData(browserWS *websocket.Conn, proxyWS *websocket.Conn) {
	go func() {
		// from browser to client(stdin)
		for {
			_, message, err := browserWS.ReadMessage()
			if err != nil {
				core.Logger.Error("[webssh]: failed to read message from browser", zap.Error(err))
				return
			}
			proxyWS.WriteMessage(websocket.TextMessage, message)
		}
	}()

	for {
		_, message, err := proxyWS.ReadMessage()
		if err != nil {
			core.Logger.Error("[webssh]: failed to read message from proxy", zap.Error(err))
			return
		}
		browserWS.WriteMessage(websocket.TextMessage, message)
	}
}
