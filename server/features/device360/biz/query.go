package device360_biz

import (
	"strings"

	"github.com/wangxin688/narvis/intend/devicerole"
)

func getSwitchDeviceRoles() string {
	deviceRoles := []string{
		string(devicerole.CoreSwitch),
		string(devicerole.DistributionSwitch),
		string(devicerole.AccessSwitch),
		string(devicerole.InternetSwitch),
	}
	return strings.Join(deviceRoles, "|")
}

func getGatewayDeviceRoles() string {
	deviceRoles := []string{
		string(devicerole.WanRouter),
		string(devicerole.Firewall),
	}
	return strings.Join(deviceRoles, "|")
}
