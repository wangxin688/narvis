package mikrotik

import "github.com/wangxin688/narvis/client/pkg/nettyssh/driver"

type MikroTikRouterOS struct {
	Driver   driver.IDriver
	Platform string
	Prompt   string
}

func (d *MikroTikRouterOS) Connect() error {

	if err := d.Driver.Connect(); err != nil {
		return err
	}
	prompt, err := d.Driver.FindDevicePrompt("\\[.*(@.*\\] >)", "] >")
	if err != nil {
		return err
	}
	d.Prompt = prompt
	return d.sessionPreparation()

}

func (d *MikroTikRouterOS) Disconnect() {

	d.Driver.Disconnect()

}

func (d *MikroTikRouterOS) SendCommand(cmd string) (string, error) {

	result, err := d.Driver.SendCommand(cmd, d.Prompt)

	return result, err

}

func (d *MikroTikRouterOS) SendConfigSet(commands []string) (string, error) {

	results, err := d.Driver.SendCommandsSet(commands, d.Prompt)

	return results, err

}

func (d *MikroTikRouterOS) sessionPreparation() error {
	_, err := d.Driver.SendCommand("", d.Prompt)
	return err

}

func (d *MikroTikRouterOS) SetTimeout(timeout uint8) {
	d.Driver.SetTimeout(timeout)
}
