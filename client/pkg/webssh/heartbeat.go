package webssh

import (
	"time"

	"github.com/gorilla/websocket"
)

const (
	websocketHeartbeatInterval = 30 * time.Second
	websocketHeartbeatTimeout  = 180 * time.Second
)

func StartWebSocketHeartbeat(wsConn *websocket.Conn) {
	readDeadline := time.Now().Add(websocketHeartbeatTimeout)
	wsConn.SetReadDeadline(readDeadline)

	// Reset the read deadline when a pong is received
	wsConn.SetPongHandler(func(string) error {
		wsConn.SetReadDeadline(readDeadline)
		return nil
	})

	// Send a ping every 30 seconds
	ticker := time.NewTicker(websocketHeartbeatInterval)
	defer ticker.Stop()

	for range ticker.C {
		if err := wsConn.WriteMessage(websocket.PingMessage, nil); err != nil {
			wsConn.Close()
			return
		}
	}
}
