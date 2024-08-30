package huawei

import "github.com/wangxin688/narvis/client/pkg/nettyssh/driver"

type HuaweiVrp struct {
	Driver   driver.IDriver
	Platform string
	Prompt   string
	Secret   string
}

func (d *HuaweiVrp) Connect() error {
	if err := d.Driver.Connect(); err != nil {
		return err
	}
	prompt, err := d.Driver.FindDevicePrompt("\\[\\]>]", ">|]")
	if err != nil {
		return err
	}
	d.Prompt = prompt
	return nil
}
func (d *HuaweiVrp) Disconnect() {
	d.Driver.Disconnect()
}
func (d *HuaweiVrp) SetTimeout(timeout uint8) {
	d.Driver.SetTimeout(timeout)
}

func (d *HuaweiVrp) SendCommand(cmd string) (string, error) {

	result, err := d.Driver.SendCommand(cmd, d.Prompt)

	return result, err
}

func (d *HuaweiVrp) SendConfigSet(commands []string) (string, error) {

	result, err := d.Driver.SendCommandsSet(commands, d.Prompt)

	return result, err
}

func (d *HuaweiVrp) sessionPreparation() error {
	_, err := d.Driver.SendCommand("", d.Prompt)

	return err

}
