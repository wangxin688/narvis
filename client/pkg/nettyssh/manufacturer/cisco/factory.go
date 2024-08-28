package cisco

import (
	"errors"

	"github.com/wangxin688/narvis/client/pkg/nettyssh/connections"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/driver"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/types"
)

func NewDevice(conn connections.Connection, platform string) (types.CiscoDevice, error) {
	deviceDriver := driver.NewDriver(conn, "\n")
	base := CSCODevice{
		Driver:   deviceDriver,
		Prompt:   "",
		Platform: platform,
	}
	switch platform {
	case "cisco_asa":
		return &ASADevice{Driver: deviceDriver, Prompt: "", base: &base}, nil
	case "cisco_xe":
		return &XEDevice{Driver: deviceDriver, Prompt: "", base: &base}, nil
	case "cisco_xr":
		return &IOSXRDevice{Driver: deviceDriver, Prompt: "", base: &base}, nil
	case "cisco_nxos":
		return &NXOSDevice{Driver: deviceDriver, Prompt: "", base: &base}, nil
	case "cisco_ios":
		return &IOSDevice{Driver: deviceDriver, Prompt: "", base: &base}, nil
	default:
		return nil, errors.New("unsupported platform: " + platform)
	}

}
