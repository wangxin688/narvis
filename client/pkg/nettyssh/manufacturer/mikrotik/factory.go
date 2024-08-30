package mikrotik

import (
	"errors"

	"github.com/wangxin688/narvis/client/pkg/nettyssh/connections"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/driver"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/types"
	plt "github.com/wangxin688/narvis/intend/platform"
)

func NewDevice(connection connections.Connection, platform string) (types.Device, error) {
	devDriver := driver.NewDriver(connection, "\r")

	switch platform {
	case string(plt.MikroTik):
		return &MikroTikRouterOS{
			Driver:   devDriver,
			Platform: platform,
			Prompt:   "",
		}, nil
	default:
		return nil, errors.New("unsupported platform: " + platform)

	}

}
