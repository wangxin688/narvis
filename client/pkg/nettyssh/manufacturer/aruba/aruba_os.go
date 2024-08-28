package aruba

import (
	"github.com/wangxin688/narvis/client/pkg/nettyssh/driver"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/types"
)

type ArubaOSDevice struct {
	Driver driver.IDriver
	Prompt string
	base   types.CiscoDevice
}

func (d *ArubaOSDevice) Connect() error {
	return d.base.Connect()

}

func (d *ArubaOSDevice) SendCommand(cmd string) (string, error) {
	return d.base.SendCommand(cmd)

}

func (d *ArubaOSDevice) SendConfigSet(cmds []string) (string, error) {
	return d.base.SendConfigSet(cmds)

}

func (d *ArubaOSDevice) Disconnect() {
	d.base.Disconnect()

}

func (d *ArubaOSDevice) SetSecret(secret string) {
	d.base.SetSecret(secret)

}

func (d *ArubaOSDevice) SetTimeout(timeout uint8) {
	d.base.SetTimeout(timeout)
}
