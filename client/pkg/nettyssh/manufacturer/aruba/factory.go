package aruba

import (
	"errors"

	"github.com/wangxin688/narvis/client/pkg/nettyssh/connections"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/driver"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/manufacturer/cisco"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/types"
)

func NewDevice(connection connections.Connection, platform string) (types.CiscoDevice, error) {
	devDriver := driver.NewDriver(connection, "\n")

	base := cisco.CSCODevice{
		Driver:   devDriver,
		Platform: "cisco_ios",
	}
	if platform != "aruba_os" {
		return nil, errors.New("unsupported Aruba device type: " + platform)

	}

	return &ArubaOSDevice{
		Driver: devDriver,
		base:   &base,
	}, nil

}
