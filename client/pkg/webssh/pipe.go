package webssh

import (
	"fmt"
	"io"
	"time"

	"github.com/gorilla/websocket"
	"github.com/wangxin688/narvis/client/utils/logger"
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
			logger.Logger.Error("Failed to read from websocket connection", err)
			return
		}

		if messageType == websocket.TextMessage && len(message) > 0 {
			// Write the message to the ssh connection's stdin
			_, err = s.stdin.Write(message)
			if err != nil {
				// If there's an error writing to the stdin, log it and return
				logger.Logger.Trace("Failed to write to stdin", err)
				return
			}
		} else {
			// Log any unknown message types
			logger.Logger.Info(fmt.Sprintf("Received unknown message type: %d", messageType))
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

	for range ticker.C {
		// Read from the stdout
		buf := make([]byte, 1024)
		n, err := s.stdout.Read(buf)
		if err != nil {
			logger.Logger.Error("Failed to read from stdout", err)
			break
		}

		// Send the data to the client
		if err = WsSendText(conn, buf[:n]); err != nil {
			logger.Logger.Error("Failed to send text message", err)
			break
		}
	}
}

func Quit(quit chan int) {
	quit <- 1
}

func WsSendText(conn *websocket.Conn, message []byte) error {
	err := conn.WriteMessage(websocket.TextMessage, message)
	return err
}
