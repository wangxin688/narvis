package cisco

import (
	"errors"

	"github.com/wangxin688/narvis/client/pkg/nettyssh/connections"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/driver"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/types"
	plt "github.com/wangxin688/narvis/intend/platform"
)

func NewDevice(conn connections.Connection, platform string) (types.CiscoDevice, error) {
	deviceDriver := driver.NewDriver(conn, "\n")
	base := CSCODevice{
		Driver:   deviceDriver,
		Prompt:   "",
		Platform: platform,
	}
	switch platform {
	case string(plt.CiscoASA):
		return &ASADevice{Driver: deviceDriver, Prompt: "", base: &base}, nil
	case string(plt.CiscoIosXE):
		return &XEDevice{Driver: deviceDriver, Prompt: "", base: &base}, nil
	case string(plt.CiscoIosXR):
		return &IOSXRDevice{Driver: deviceDriver, Prompt: "", base: &base}, nil
	case string(plt.CiscoNexusOS):
		return &NXOSDevice{Driver: deviceDriver, Prompt: "", base: &base}, nil
	case string(plt.CiscoIos):
		return &IOSDevice{Driver: deviceDriver, Prompt: "", base: &base}, nil
	default:
		return nil, errors.New("unsupported platform: " + platform)
	}

}
