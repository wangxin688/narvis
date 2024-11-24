package devicemodel

import (
	manufacturer "github.com/wangxin688/narvis/intend/model/manufacturer"
	platform "github.com/wangxin688/narvis/intend/model/platform"
)

var UnknownDeviceModel = "Unknown"

type DeviceModel struct {
	Platform     platform.Platform
	Manufacturer manufacturer.Manufacturer
	DeviceModel  string
}
