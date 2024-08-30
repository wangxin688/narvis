package aruba

import (
	"errors"

	"github.com/wangxin688/narvis/client/pkg/nettyssh/connections"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/driver"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/manufacturer/cisco"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/types"
	plt "github.com/wangxin688/narvis/intend/platform"
)

func NewDevice(connection connections.Connection, platform string) (types.CiscoDevice, error) {
	devDriver := driver.NewDriver(connection, "\n")

	base := cisco.CSCODevice{
		Driver:   devDriver,
		Platform: string(plt.CiscoIos),
	}
	if platform != string(plt.Aruba) {
		return nil, errors.New("unsupported Aruba platform: " + platform)
	
	}

	return &ArubaOSDevice{
		Driver: devDriver,
		base:   &base,
	}, nil

}
