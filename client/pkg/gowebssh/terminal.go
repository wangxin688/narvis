package gowebssh

import (
	"bytes"

	"github.com/gorilla/websocket"
	"github.com/wangxin688/narvis/client/utils/logger"
	"golang.org/x/crypto/ssh"
)

type Options struct {
	Addr     string `json:"addr"`
	Port     uint16 `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Cols     int    `json:"cols"`
	Rows     int    `json:"rows"`
}

type Terminal struct {
	opts Options
	// ip       string
	ws *websocket.Conn
	// stdin    *os.File
	// stdout   *os.File
	session *SshConn
	conn    *ssh.Client
	// inited   int32
	// cancelFn context.CancelFunc
}

func NewTerminal(ws *websocket.Conn, opts Options) *Terminal {
	return &Terminal{opts: opts, ws: ws}
}

func (t *Terminal) Run() {
	var err error
	t.conn, err = NewSshClient(t.opts.Addr, t.opts.User, t.opts.Password, t.opts.Port)
	if WsHandleError(t.ws, err) {
		return
	}
	defer func() {
		t.conn.Close()
	}()
	//startTime := time.Now()
	t.session, err = NewSshConn(t.opts.Cols, t.opts.Rows, t.conn)

	if WsHandleError(t.ws, err) {
		return
	}
	defer func() {
		t.session.Close()
	}()

	quitChan := make(chan bool, 3)

	var logBuff = new(bytes.Buffer)

	go t.session.ReceiveWsMsg(t.ws, logBuff, quitChan)
	go t.session.SendComboOutput(t.ws, quitChan)
	go t.session.SessionWait(quitChan)

	<-quitChan
	logger.Logger.Info("[webssh]: websocket finished")
}
