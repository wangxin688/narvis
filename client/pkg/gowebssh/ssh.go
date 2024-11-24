package gowebssh

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/wangxin688/narvis/intend/logger"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh"
)

func NewSshClient(host, user, password string, port uint16) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User:            user,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Duration(6) * time.Second,
		Config: ssh.Config{
			Ciphers: []string{
				"aes128-ctr", "aes192-ctr",
				"aes256-ctr", "aes128-gcm@openssh.com",
				"aes256-gcm@openssh.com", "chacha20-poly1305@openssh.com",
				"arcfour256", "arcfour128", "aes128-cbc",
				"3des-cbc", "aes192-cbc", "aes256-cbc",
			},
		},
	}
	targetHost := fmt.Sprintf("%s:%d", host, port)
	networkType := "tcp"
	addr := net.ParseIP(host)
	if addr.To4() == nil {
		targetHost = fmt.Sprintf("[%s]:%d", host, port)
		networkType = "tcp6"
	}
	client, err := ssh.Dial(networkType, targetHost, config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

type wsBufferWriter struct {
	buffer bytes.Buffer
	mu     sync.Mutex
}

func (w *wsBufferWriter) Write(p []byte) (int, error) {
	w.mu.Lock()
	defer w.mu.Unlock()
	return w.buffer.Write(p)
}

const (
	wsMsgCmd    = "cmd"
	wsMsgResize = "resize"
)

type wsMsg struct {
	Type string `json:"type"`
	Cmd  string `json:"cmd"`
	Cols int    `json:"cols"`
	Rows int    `json:"rows"`
}

type SshConn struct {
	StdinPipe   io.WriteCloser
	ComboOutput *wsBufferWriter
	Session     *ssh.Session
}

func flushComboOutput(w *wsBufferWriter, wsConn *websocket.Conn) error {
	if w.buffer.Len() != 0 {
		err := wsConn.WriteMessage(websocket.TextMessage, w.buffer.Bytes())
		if err != nil {
			return err
		}
		w.buffer.Reset()
	}
	return nil
}

func NewSshConn(cols, rows int, sshClient *ssh.Client) (*SshConn, error) {
	sshSession, err := sshClient.NewSession()
	if err != nil {
		return nil, err
	}

	stdinP, err := sshSession.StdinPipe()
	if err != nil {
		return nil, err
	}

	comboWriter := new(wsBufferWriter)
	sshSession.Stdout = comboWriter
	sshSession.Stderr = comboWriter

	modes := ssh.TerminalModes{
		ssh.ECHO:          1,     // disable echo
		ssh.TTY_OP_ISPEED: 14400, // input speed = 14.4kbaud
		ssh.TTY_OP_OSPEED: 14400, // output speed = 14.4kbaud
	}
	if err := sshSession.RequestPty("xterm", rows, cols, modes); err != nil {
		return nil, err
	}
	// Start remote shell
	if err := sshSession.Shell(); err != nil {
		return nil, err
	}
	return &SshConn{StdinPipe: stdinP, ComboOutput: comboWriter, Session: sshSession}, nil
}

func (s *SshConn) Close() {
	if s.Session != nil {
		s.Session.Close()
	}

}

// ReceiveWsMsg  receive websocket msg do some handling then write into ssh.session.stdin
func (s *SshConn) ReceiveWsMsg(wsConn *websocket.Conn, logBuff *bytes.Buffer, exitCh chan bool) {
	//tells other go routine quit
	defer setQuit(exitCh)
	for {
		select {
		case <-exitCh:
			return
		default:
			//read websocket msg
			_, wsData, err := wsConn.ReadMessage()
			if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
				return
			}
			if err != nil {
				logger.Logger.Error("reading webSocket message failed", zap.Error(err))
				return
			}
			//unmashal bytes into struct
			msgObj := wsMsg{
				Type: "cmd",
				Cmd:  "",
				Rows: 50,
				Cols: 180,
			}
			if err := json.Unmarshal(wsData, &msgObj); err != nil {
				logger.Logger.Error("unmarshal websocket message failed", zap.Error(err))
			}
			switch msgObj.Type {
			case wsMsgResize:
				//handle xterm.js size change
				if msgObj.Cols > 0 && msgObj.Rows > 0 {
					if err := s.Session.WindowChange(msgObj.Rows, msgObj.Cols); err != nil {
						logger.Logger.Error("ssh pty change windows size failed", zap.Error(err))
					}
				}
			case wsMsgCmd:
				decodeBytes := []byte(msgObj.Cmd)
				if _, err := s.StdinPipe.Write(decodeBytes); err != nil {
					logger.Logger.Error("ws cmd bytes write to ssh.stdin pipe failed", zap.Error(err))
				}
				if _, err := logBuff.Write(decodeBytes); err != nil {
					logger.Logger.Error("write received cmd into log buffer failed")
				}
			}
		}
	}
}
func (s *SshConn) SendComboOutput(wsConn *websocket.Conn, exitCh chan bool) {
	defer setQuit(exitCh)

	tick := time.NewTicker(time.Millisecond * time.Duration(20))
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			if err := flushComboOutput(s.ComboOutput, wsConn); err != nil {
				logger.Logger.Error("ssh sending combo output to webSocket failed", zap.Error(err))
				return
			}
		case <-exitCh:
			return
		}
	}
}

func (s *SshConn) SessionWait(quitChan chan bool) {
	if err := s.Session.Wait(); err != nil {
		if ExitError, ok := err.(*ssh.ExitError); ok {
			logger.Logger.Error("[websshConn]: Remote command exited with status", zap.Int("exit code", ExitError.ExitStatus()), zap.Error(err))
		} else {
			logger.Logger.Error("[websshConn]: Failed to run remote command", zap.Error(err))

		}
		setQuit(quitChan)
	}
}

func setQuit(ch chan bool) {
	ch <- true
}
