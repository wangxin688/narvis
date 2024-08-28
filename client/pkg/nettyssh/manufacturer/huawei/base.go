package huawei

import "github.com/wangxin688/narvis/client/pkg/nettyssh/driver"


type HuaweiDevice struct {
	Driver driver.IDriver
	Platform string
	Prompt string
	Secret string
}


func (d *HuaweiDevice) Connect() error {
	return nil
}
func (d *HuaweiDevice) Disconnect() {
}
func (d *HuaweiDevice) SetTimeout(timeout uint8) {
}

func (d *HuaweiDevice) SendCommand(cmd string) (string, error) {

}


func (d *HuaweiDevice) SendConfigSet(cmds []string) (string, error) {}


func (d *HuaweiDevice) sessionPreparation() {}