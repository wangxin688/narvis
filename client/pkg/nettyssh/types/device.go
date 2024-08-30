package types

type Device interface {
	Connect() error
	Disconnect()
	SendCommand(cmd string) (string, error)
	SendConfigSet(commands []string) (string, error)
	SetTimeout(timeout uint8)
}

type CiscoDevice interface {
	Connect() error
	Disconnect()
	SendCommand(cmd string) (string, error)
	SendConfigSet(commands []string) (string, error)
	SetSecret(secret string)
	SetTimeout(timeout uint8)
}
