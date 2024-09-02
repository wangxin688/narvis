package devicerole

import (
	"github.com/samber/lo"
	"github.com/wangxin688/narvis/intend/common"
)

type DeviceRoleEnum string
type ProductFamilyEnum string

const (
	Gateway              ProductFamilyEnum = "Gateway"
	Switching            ProductFamilyEnum = "Switching"
	Wireless             ProductFamilyEnum = "Wireless"
	Computing            ProductFamilyEnum = "Computing"
	UnknownProductFamily ProductFamilyEnum = "Unknown"
)

const (
	// WAN Layer
	WanRouter      DeviceRoleEnum = "WanRouter"
	Firewall       DeviceRoleEnum = "Firewall"
	InternetSwitch DeviceRoleEnum = "InternetSwitch"
	// Campus Lan Layer
	CoreSwitch         DeviceRoleEnum = "CoreSwitch"
	DistributionSwitch DeviceRoleEnum = "DistributionSwitch"
	AccessSwitch       DeviceRoleEnum = "AccessSwitch"
	WlanAC             DeviceRoleEnum = "WlanAC"
	WlanAP             DeviceRoleEnum = "WlanAP"
	Server             DeviceRoleEnum = "Server"
	UnknownDeviceRole  DeviceRoleEnum = "Unknown"
)

type DeviceRole struct {
	DeviceRole    DeviceRoleEnum    `json:"device_role"`
	Description   common.I18n       `json:"description"`
	Weight        uint16            `json:"weight"`
	ProductFamily ProductFamilyEnum `json:"product_family"`
}

func (d DeviceRole) ToMap() map[string]any {
	result := make(map[string]any)
	result["deviceRole"] = d.DeviceRole
	result["description"] = map[string]string{"en": d.Description.En, "zh": d.Description.Zh}
	result["weight"] = d.Weight
	result["productFamily"] = d.ProductFamily
	return result
}

func getDeviceRoleMeta() map[DeviceRoleEnum]DeviceRole {
	// do not change the value to pointer since data cannot be modified
	deviceRoleMeta := map[DeviceRoleEnum]DeviceRole{
		WanRouter: {
			DeviceRole:    WanRouter,
			Description:   common.I18n{En: "WanRouter", Zh: "出口路由器"},
			Weight:        10,
			ProductFamily: Gateway},
		Firewall: {
			DeviceRole:    Firewall,
			Description:   common.I18n{En: "Firewall", Zh: "防火墙"},
			Weight:        10,
			ProductFamily: Gateway},
		InternetSwitch: {
			DeviceRole:    InternetSwitch,
			Description:   common.I18n{En: "InternetSwitch", Zh: "互联网交换机"},
			Weight:        10,
			ProductFamily: Switching},
		CoreSwitch: {
			DeviceRole:    CoreSwitch,
			Description:   common.I18n{En: "CoreSwitch", Zh: "核心交换机"},
			Weight:        100,
			ProductFamily: Switching},
		DistributionSwitch: {
			DeviceRole:    DistributionSwitch,
			Description:   common.I18n{En: "DistributionSwitch", Zh: "汇聚交换机"},
			Weight:        150,
			ProductFamily: Switching},
		AccessSwitch: {
			DeviceRole:    AccessSwitch,
			Description:   common.I18n{En: "AccessSwitch", Zh: "接入交换机"},
			Weight:        200,
			ProductFamily: Switching},
		WlanAC: {
			DeviceRole:    WlanAC,
			Description:   common.I18n{En: "WirelessController", Zh: "无线控制器"},
			Weight:        120,
			ProductFamily: Wireless},
		WlanAP: {
			DeviceRole:    WlanAP,
			Description:   common.I18n{En: "WlanAP", Zh: "无线AP"},
			Weight:        300,
			ProductFamily: Wireless},
		Server: {
			DeviceRole:    Server,
			Description:   common.I18n{En: "Server", Zh: "服务器"},
			Weight:        300,
			ProductFamily: Computing},
		UnknownDeviceRole: {
			DeviceRole:    UnknownDeviceRole,
			Description:   common.I18n{En: "Unknown", Zh: "未知"},
			Weight:        300,
			ProductFamily: UnknownProductFamily,
		},
	}
	return deviceRoleMeta

}

func GetListDeviceRole() (int, []DeviceRole) {
	result := getDeviceRoleMeta()
	return len(result), lo.Values(result)
}

func GetDeviceRole(deviceRole string) DeviceRole {

	deviceRoleMeta := getDeviceRoleMeta()

	if role, ok := deviceRoleMeta[DeviceRoleEnum(deviceRole)]; ok {
		return role
	}
	return deviceRoleMeta[UnknownDeviceRole]
}

func GetSwitchDeviceRole() *[]DeviceRoleEnum {
	result := make([]DeviceRoleEnum, 0)
	for k, v := range getDeviceRoleMeta() {
		if v.ProductFamily == Switching {
			result = append(result, k)
		}
	}
	return &result
}
