package juniper

import (
	"errors"
	"strings"

	"github.com/wangxin688/narvis/client/pkg/nettyssh/driver"
)

type JunOSDevice struct {
	Driver   driver.IDriver
	Platform string
	Prompt   string
}

func (d *JunOSDevice) Connect() error {
	if err := d.Driver.Connect(); err != nil {
		return err
	}
	prompt, err := d.Driver.FindDevicePrompt("(@.*)[#>%]", "%")
	if err != nil {
		return err
	}
	d.Prompt = prompt

	return d.sessionPreparation()

}

func (d *JunOSDevice) Disconnect() {

	d.Driver.Disconnect()

}

func (d *JunOSDevice) SendCommand(cmd string) (string, error) {

	result, err := d.Driver.SendCommand(cmd, d.Prompt)

	return result, err

}

func (d *JunOSDevice) SendConfigSet(commands []string) (string, error) {

	results, _ := d.Driver.SendCommand("configure", d.Prompt)

	commands = append(commands, "commit", "exit")

	out, err := d.Driver.SendCommandsSet(commands, d.Prompt)
	results += out
	return results, err

}

func (d *JunOSDevice) sessionPreparation() error {

	out, err := d.Driver.SendCommand("cli", d.Prompt)
	if err != nil {
		return errors.New("failed to send cli command" + err.Error())
	}
	if !strings.Contains(out, ">") {
		return errors.New("failed to enter cli mode, device output: " + out)
	}

	_, err = d.SendCommand("set cli screen-length 0")

	if err != nil {
		return errors.New("failed to disable pagination" + err.Error())
	}
	return nil

}

func (d *JunOSDevice) SetTimeout(timeout uint8) {
	d.Driver.SetTimeout(timeout)
}
