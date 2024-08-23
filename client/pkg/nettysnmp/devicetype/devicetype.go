package devicetype

import (
	"github.com/wangxin688/narvis/intend/manufacturer"
	"github.com/wangxin688/narvis/intend/platform"
)

var UnknownDeviceType = "Unknown"

type DeviceType struct {
	Platform     platform.Platform
	Manufacturer manufacturer.Manufacturer
	DeviceType   string
}
