package cisco

import (
	"github.com/wangxin688/narvis/client/pkg/nettyssh/driver"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/types"
)

type ASADevice struct {
	Driver driver.IDriver
	Prompt string
	base   types.CiscoDevice
}

func (d *ASADevice) Connect() error {
	return d.base.Connect()

}

func (d *ASADevice) Disconnect() {
	d.base.Disconnect()

}

func (d *ASADevice) SendCommand(cmd string) (string, error) {
	return d.base.SendCommand(cmd)

}

func (d *ASADevice) SendConfigSet(commands []string) (string, error) {
	return d.base.SendConfigSet(commands)
}

func (d *ASADevice) SetSecret(secret string) {
	d.base.SetSecret(secret)
}

func (d *ASADevice) SetTimeout(timeout uint8) {
	d.base.SetTimeout(timeout)
}
