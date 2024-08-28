package juniper

import (
	"github.com/wangxin688/narvis/client/pkg/nettyssh/connections"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/driver"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/types"
)

func NewDevice(connection connections.Connection, Platform string) (types.Device, error) {
	devDriver := driver.NewDriver(connection, "\n")
	return &JunOSDevice{
		Prompt:   "",
		Driver:   devDriver,
		Platform: Platform,
	}, nil

}
