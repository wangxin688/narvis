package h3c

import (
	"errors"

	"github.com/wangxin688/narvis/client/pkg/nettyssh/driver"
)

type H3C struct {
	Driver   driver.IDriver
	Platform string
	Prompt   string
	Secret   string
}

func (d *H3C) Connect() error {
	if err := d.Driver.Connect(); err != nil {
		return err
	}
	prompt, err := d.Driver.FindDevicePrompt("\r\n?(\\S+)[\\]>]", "[\\]|>]")
	if err != nil {
		return err
	}
	d.Prompt = prompt
	return d.sessionPreparation()
}
func (d *H3C) Disconnect() {
	d.Driver.Disconnect()
}
func (d *H3C) SetTimeout(timeout uint8) {
	d.Driver.SetTimeout(timeout)
}

func (d *H3C) SendCommand(cmd string) (string, error) {

	result, err := d.Driver.SendCommand(cmd, d.Prompt)

	return result, err
}

func (d *H3C) SendConfigSet(commands []string) (string, error) {
	result, _ := d.Driver.SendCommand("system-view", d.Prompt)
	commands = append(commands, "quit")
	out, err := d.Driver.SendCommandsSet(commands, d.Prompt)

	result += out
	return result, err
}

func (d *H3C) sessionPreparation() error {
	_, err := d.Driver.SendCommand("", d.Prompt)
	if err != nil {
		return err
	}
	_, err = d.Driver.SendCommand("screen-length 0 temporary", d.Prompt)

	if err != nil {
		return errors.New("failed to disable pagination" + err.Error())
	}
	return nil

}
