package devicemodel

import (
	"github.com/wangxin688/narvis/intend/manufacturer"
	"github.com/wangxin688/narvis/intend/platform"
)

var UnknownDeviceModel = "Unknown"

type DeviceModel struct {
	Platform     platform.Platform
	Manufacturer manufacturer.Manufacturer
	DeviceModel  string
}
