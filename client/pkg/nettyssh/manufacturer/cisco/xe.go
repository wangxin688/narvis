package cisco

import (
	"github.com/wangxin688/narvis/client/pkg/nettyssh/driver"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/types"
)

type XEDevice struct {
	Driver driver.IDriver
	Prompt string
	base   types.CiscoDevice
}

func (d *XEDevice) Connect() error {
	return d.base.Connect()

}

func (d *XEDevice) Disconnect() {

	d.base.Disconnect()

}

func (d *XEDevice) SendCommand(cmd string) (string, error) {
	return d.base.SendCommand(cmd)

}

func (d *XEDevice) SendConfigSet(commands []string) (string, error) {
	return d.base.SendConfigSet(commands)

}
func (d *XEDevice) SetSecret(secret string) {
	d.base.SetSecret(secret)
}

func (d *XEDevice) SetTimeout(timeout uint8) {
	d.base.SetTimeout(timeout)
}
