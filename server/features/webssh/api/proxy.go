package webssh_api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/wangxin688/narvis/server/core"
	webssh_biz "github.com/wangxin688/narvis/server/features/webssh/biz"
	"github.com/wangxin688/narvis/server/tools/errors"
	"go.uber.org/zap"
)

func handleProxyWebSocket(c *gin.Context) error {
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
	sessionId := c.Param("sessionId")
	if sessionId == "" {
		core.Logger.Warn("[webssh]:received unknown empty sessionId")
		wsConn.Close()
		return errors.NewError(errors.CodeSessionIdEmpty, errors.MsgSessionIdEmpty)
	}
	if ch, ok := webssh_biz.SessionWMap.Load(sessionId); ok {
		core.Logger.Info("[webssh]: received session from webssh socket", zap.String("sessionId", sessionId))
		done := ch.(chan *websocket.Conn)
		done <- wsConn
		return nil
	}
	wsConn.Close()
	return errors.NewError(errors.CodeSessionIdNotFound, errors.MsgSessionIdNotFound)
}
