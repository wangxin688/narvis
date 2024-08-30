package arista

import (
	"github.com/wangxin688/narvis/client/pkg/nettyssh/driver"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/types"
)

type EOSDevice struct {
	Driver driver.IDriver
	Prompt string
	base   types.CiscoDevice
}

func (d *EOSDevice) Connect() error {
	return d.base.Connect()

}

func (d *EOSDevice) SendCommand(cmd string) (string, error) {
	return d.base.SendCommand(cmd)

}

func (d *EOSDevice) SendConfigSet(commands []string) (string, error) {
	return d.base.SendConfigSet(commands)

}

func (d *EOSDevice) Disconnect() {
	d.base.Disconnect()

}

func (d *EOSDevice) SetSecret(secret string) {
	d.base.SetSecret(secret)

}

func (d *EOSDevice) SetTimeout(timeout uint8) {
	d.base.SetTimeout(timeout)
}
