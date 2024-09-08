package tplink

import (
	"github.com/wangxin688/narvis/client/pkg/nettyssh/driver"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/types"
)

type TpLinkDevice struct {
	Driver driver.IDriver
	Prompt string
	base   types.CiscoDevice
}

func (d *TpLinkDevice) Connect() error {
	return d.base.Connect()

}

func (d *TpLinkDevice) SendCommand(cmd string) (string, error) {
	return d.base.SendCommand(cmd)

}

func (d *TpLinkDevice) SendConfigSet(commands []string) (string, error) {
	return d.base.SendConfigSet(commands)

}

func (d *TpLinkDevice) Disconnect() {
	d.base.Disconnect()

}

func (d *TpLinkDevice) SetSecret(secret string) {
	d.base.SetSecret(secret)

}

func (d *TpLinkDevice) SetTimeout(timeout uint8) {
	d.base.SetTimeout(timeout)
}
