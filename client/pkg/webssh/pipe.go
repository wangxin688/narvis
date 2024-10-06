package webssh

import (
	"fmt"
	"io"
	"time"

	"github.com/gorilla/websocket"
	"github.com/wangxin688/narvis/client/utils/logger"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
)

type SSHConnection struct {
	session *ssh.Session
	stdin   io.WriteCloser
	stdout  io.Reader
	stderr  io.Reader
}

// Recv reads from the websocket connection and writes to the ssh connection's stdin
// When the quit channel is closed, the function will exit.
func (s *SSHConnection) Recv(conn *websocket.Conn, quit chan int) {
	defer Quit(quit)

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			// If there's an error reading from the websocket connection, log it and return
			logger.Logger.Error("[websshStdinPipe]: failed to read from websocket connection", zap.Error(err))
			return
		}

		if messageType == websocket.TextMessage && len(message) > 0 {
			// Write the message to the ssh connection's stdin
			_, err = s.stdin.Write(message)
			if err != nil {
				// If there's an error writing to the stdin, log it and return
				logger.Logger.Error("[websshStdinPipe]: failed to write to stdin", zap.Error(err))
				return
			}
		} else {
			// Log any unknown message types
			logger.Logger.Info(fmt.Sprintf("[websshStdinPipe]: received unknown message type: %d", messageType))
		}
	}
}

// Send reads from the ssh connection's stdout and writes to the websocket connection
// This function uses a ticker to check for available data on the stdout every 60us.
// This is needed because the ssh library does not provide a way to wait for available data on the stdout.
func (s *SSHConnection) Send(conn *websocket.Conn, quit chan int) {
	defer Quit(quit)

	ticker := time.NewTicker(60 * time.Microsecond)
	defer ticker.Stop()

	buf := make([]byte, 4096) // add to buffer size to 4096 for network devices
	for range ticker.C {
		n, err := s.stdout.Read(buf)
		if err != nil {
			logger.Logger.Error("[websshSendPipe]: failed to read from stdout", zap.Error(err))
			break
		}

		// If the buffer is full, it's likely that there's more data to read.
		// We'll keep reading until we've read all the available data.
		for n == len(buf) {
			logger.Logger.Info((""))
			if err := WsSendText(conn, buf); err != nil {
				logger.Logger.Error("[websshSendPipe]: failed to send text message", zap.Error(err))
				break
			}
			n, err = s.stdout.Read(buf)
			if err != nil {
				logger.Logger.Error("[]websshSendPipe]: failed to read from stdout", zap.Error(err))
				break
			}
		}

		// Send any remaining data
		if n > 0 {
			if err := WsSendText(conn, buf[:n]); err != nil {
				logger.Logger.Error("[]websshSendPipe]: failed to send text message", zap.Error(err))
				break
			}
		}
	}
}

func Quit(quit chan int) {
	quit <- 1
}

func WsSendText(conn *websocket.Conn, message []byte) error {
	logger.Logger.Info("[websshSendPipe]: sending text message", zap.String("message", string(message)))
	err := conn.WriteMessage(websocket.TextMessage, message)
	return err
}
