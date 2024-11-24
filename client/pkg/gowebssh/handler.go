package gowebssh

import (
	"time"

	"github.com/gorilla/websocket"
	"github.com/wangxin688/narvis/intend/logger"
	"go.uber.org/zap"
)

func WsHandleError(ws *websocket.Conn, err error) bool {
	if err != nil {
		logger.Logger.Error("handler ws ERROR:", zap.Error(err))
		dt := time.Now().Add(time.Second)
		if err := ws.WriteControl(websocket.CloseMessage, []byte(err.Error()), dt); err != nil {
			logger.Logger.Error("websocket writes control message failed:", zap.Error(err))
		}
		return true
	}
	return false
}
