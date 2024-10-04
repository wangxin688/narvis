package webssh

import (
	"fmt"
	"net"
	"time"

	"golang.org/x/crypto/ssh"
)

func CreateSSHClient(username, password, host string, port uint16) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User:            username,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Duration(6) * time.Second,
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

func NewTerminal(client *ssh.Client, cols, rows int) (*SSHConnection, error) {

	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
	}
	err = session.RequestPty("xterm", rows, cols, modes)
	if err != nil {
		return nil, err
	}
	pipe, _ := session.StdinPipe()
	stdout, _ := session.StdoutPipe()
	return &SSHConnection{
		session: session,
		stdin:   pipe,
		stdout:  stdout,
	}, nil
}
