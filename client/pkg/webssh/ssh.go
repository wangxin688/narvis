package webssh

import (
	"fmt"
	"net"
	"time"

	"github.com/wangxin688/narvis/client/utils/logger"
	"golang.org/x/crypto/ssh"
)

func CreateSSHClient(username, password, host string, port uint16) (*ssh.Client, error) {
	config := &ssh.ClientConfig{
		User:            username,
		Auth:            []ssh.AuthMethod{ssh.Password(password)},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		Timeout:         time.Duration(6) * time.Second,
		Config: ssh.Config{
			Ciphers: []string{
				"aes128-ctr", "aes192-ctr",
				"aes256-ctr", "aes128-gcm@openssh.com",
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

func NewTerminal(client *ssh.Client, cols, rows int) (*SSHConnection, error) {

	session, err := client.NewSession()
	if err != nil {
		return nil, err
	}
	modes := ssh.TerminalModes{
		ssh.ECHO:          1,
		ssh.TTY_OP_ISPEED: 14400,
		ssh.TTY_OP_OSPEED: 14400,
		ssh.IEXTEN:        0,
	}
	err = session.RequestPty("xterm", rows, cols, modes)
	if err != nil {
		return nil, err
	}
	stdout, err := session.StdoutPipe()
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[newSSHTerminal]: failed to get stdout from ip %s, err: %s", client.RemoteAddr().String(), err))
		return nil, err
	}
	stderr, err := session.StderrPipe()
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[newSSHTerminal]: failed to get stderr from ip %s, err: %s", client.RemoteAddr().String(), err))
		return nil, err
	}
	stdin, err := session.StdinPipe()
	if err != nil {
		logger.Logger.Error(fmt.Sprintf("[newSSHTerminal]: failed to get stdin from ip %s, err: %s", client.RemoteAddr().String(), err))
		return nil, err
	}

	return &SSHConnection{
		session: session,
		stdin:   stdin,
		stdout:  stdout,
		stderr:  stderr,
	}, nil
}
