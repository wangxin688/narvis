package nettyssh

import (
	"errors"
	"strings"

	"github.com/wangxin688/narvis/client/pkg/nettyssh/connections"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/manufacturer/arista"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/manufacturer/cisco"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/manufacturer/huawei"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/manufacturer/juniper"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/manufacturer/mikrotik"
	"github.com/wangxin688/narvis/client/pkg/nettyssh/types"
)

func NewDevice(Host string, Username string, Password string, DeviceType string, Port uint8, Options ...DeviceOption) (types.Device, error) {
	var device types.Device

	//for Mikrotik you need to append +ct200w to username
	if strings.Contains(DeviceType, "mikrotik") {
		Username += "+ct200w"
	}

	//create connection
	connection, err := connections.NewConnection(Host, Username, Password, Port)
	if err != nil {
		return nil, err
	}

	//create the Device
	if strings.Contains(DeviceType, "cisco") {
		device, err = cisco.NewDevice(connection, DeviceType)
	} else if strings.Contains(DeviceType, "arista") {
		device, err = arista.NewDevice(connection, DeviceType)
	} else if strings.Contains(DeviceType, "juniper") {
		device, err = juniper.NewDevice(connection, DeviceType)
	} else if strings.Contains(DeviceType, "mikrotik") {
		device, err = mikrotik.NewDevice(connection, DeviceType)
	} else if strings.Contains(DeviceType, "huawei") {
		device, err = huawei.NewDevice(connection, DeviceType)
	} else {
		return nil, errors.New("DeviceType not supported: " + DeviceType)
	}
	if err != nil {
		return nil, err
	}

	// running Options Functions.
	for _, option := range Options {
		err := option(device)
		if err != nil {
			return nil, err
		}
	}

	return device, nil

}
