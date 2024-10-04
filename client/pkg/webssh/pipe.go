package webssh

import (
	"fmt"
	"io"
	"time"

	"github.com/gorilla/websocket"
	"golang.org/x/crypto/ssh"
)

type SSHConnection struct {
	session *ssh.Session
	stdin   io.WriteCloser
	stdout  io.Reader
}

func (s *SSHConnection) Recv(conn *websocket.Conn, quit chan int) {
	defer Quit(quit)
	var (
		bytes []byte
		err   error
	)
	for {
		if _, bytes, err = conn.ReadMessage(); err != nil {
			return
		}
		if len(bytes) > 0 {
			if _, e := s.stdin.Write(bytes); e != nil {
				return
			}

		}
	}
}

func (s *SSHConnection) Send(conn *websocket.Conn, quit chan int) {
	defer Quit(quit)
	var (
		read int
		err  error
	)
	tick := time.NewTicker(60 * time.Microsecond)
	defer tick.Stop()
Loop:
	for range tick.C {
		i := make([]byte, 1024)
		if read, err = s.stdout.Read(i); err != nil {
			fmt.Println(err)
			break Loop
		}
		if err = WsSendText(conn, i[:read]); err != nil {
			fmt.Println(err)
			break Loop
		}
	}

}

func Quit(quit chan int) {
	quit <- 1
}

func WsSendText(conn *websocket.Conn, msg []byte) error {
	if err := conn.WriteMessage(websocket.TextMessage, msg); err != nil {
		return err
	}
	return nil
}
