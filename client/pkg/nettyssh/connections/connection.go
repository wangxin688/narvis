package connections

type Connection interface {
	Connect() error
	Disconnect()
	Read() (string, error)
	Write(cmd string) int
	SetTimeout(timeout uint8)
}

func NewConnection(host, username, password string, port uint8) (Connection, error) {
	conn, err := NewSSHConn(host, username, password, port)
	if err != nil {
		return nil, err
	}
	return &conn, nil
}
