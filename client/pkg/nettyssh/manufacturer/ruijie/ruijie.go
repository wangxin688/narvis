package ruijie

import (
	"github.com/wangxin688/narvis/client/pkg/nettyssh/driver"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/types"
)

type RuijieDevice struct {
	Driver driver.IDriver
	Prompt string
	base   types.CiscoDevice
}

func (d *RuijieDevice) Connect() error {
	return d.base.Connect()

}

func (d *RuijieDevice) SendCommand(cmd string) (string, error) {
	return d.base.SendCommand(cmd)

}

func (d *RuijieDevice) SendConfigSet(commands []string) (string, error) {
	return d.base.SendConfigSet(commands)

}

func (d *RuijieDevice) Disconnect() {
	d.base.Disconnect()

}

func (d *RuijieDevice) SetSecret(secret string) {
	d.base.SetSecret(secret)

}

func (d *RuijieDevice) SetTimeout(timeout uint8) {
	d.base.SetTimeout(timeout)
}
