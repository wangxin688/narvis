package nettyssh

import (
	"errors"
	"strings"

	"github.com/wangxin688/narvis/client/pkg/nettyssh/connections"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/manufacturer/arista"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/manufacturer/cisco"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/manufacturer/h3c"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/manufacturer/huawei"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/manufacturer/juniper"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/manufacturer/mikrotik"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/manufacturer/ruijie"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/manufacturer/tplink"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/types"
)

func NewDevice(Host string, username string, password string, platform string, port uint8, options ...DeviceOption) (types.Device, error) {
	var device types.Device

	//for Mikrotik you need to append +ct200w to username
	if strings.Contains(platform, "mikrotik") {
		username += "+ct200w"
	}

	//create connection
	connection, err := connections.NewConnection(Host, username, password, port)
	if err != nil {
		return nil, err
	}

	//create the Device
	if strings.Contains(platform, "cisco") {
		device, err = cisco.NewDevice(connection, platform)
	} else if strings.Contains(platform, "arista") {
		device, err = arista.NewDevice(connection, platform)
	} else if strings.Contains(platform, "juniper") {
		device, err = juniper.NewDevice(connection, platform)
	} else if strings.Contains(platform, "mikrotik") {
		device, err = mikrotik.NewDevice(connection, platform)
	} else if strings.Contains(platform, "huawei") {
		device, err = huawei.NewDevice(connection, platform)
	} else if strings.Contains(platform, "h3c") {
		device, err = h3c.NewDevice(connection, platform)
	} else if strings.Contains(platform, "ruijie") {
		device, err = ruijie.NewDevice(connection, platform)
	} else if strings.Contains(platform, "tp_link") {
		device, err = tplink.NewDevice(connection, platform)
	} else {
		return nil, errors.New("platform not supported: " + platform)
	}
	if err != nil {
		return nil, err
	}

	// running options Functions.
	for _, option := range options {
		err := option(device)
		if err != nil {
			return nil, err
		}
	}

	return device, nil

}
