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
	"github.com/wangxin688/narvis/server/tools/errors"
	"go.uber.org/zap"
)

var SessionWMap sync.Map

func AddSession(sessionId string, ws chan *websocket.Conn) {
	if _, ok := SessionWMap.Load(sessionId); ok {
		return
	}
	SessionWMap.Store(sessionId, ws)
}
func SendSignalToProxy(sessionId, managementIP string, cred *schemas.CliCredential) error {
	signal := intendtask.WebSSHTask{
		TaskName:     intendtask.WebSSH,
		ManagementIP: managementIP,
		SessionId:    sessionId,
		Username:     cred.Username,
		Password:     cred.Password,
		Port:         cred.Port,
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
	done, ok := SessionWMap.Load(sessionId)
	if !ok {
		return nil, errors.NewError(errors.CodeSessionIdNotFound, errors.MsgSessionIdNotFound)
	}
	proxyWS, ok := done.(*websocket.Conn)
	if !ok {
		return nil, errors.NewError(errors.CodeSessionIdNotFound, errors.MsgSessionIdNotFound)
	}
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
