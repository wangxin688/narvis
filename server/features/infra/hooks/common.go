package hooks

import (
	"fmt"
	"strings"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/server/core"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/pkg/zbx/zschema"
	"go.uber.org/zap"
)

func proxySelect(orgId string) string {
	proxies, err := gen.Proxy.Select(gen.Proxy.ProxyId).Where(gen.Proxy.OrganizationId.Eq(orgId)).Find()
	if err != nil {
		core.Logger.Error(fmt.Sprintf("[proxyChoice]: proxyChoice for organization failed with orgId %s", orgId), zap.Error(err))
		return ""
	}
	if len(proxies) == 0 {
		return ""
	}
	proxyIds := make([]string, 0, len(proxies))
	for _, proxy := range proxies {
		proxyIds = append(proxyIds, *proxy.ProxyId)
	}
	return lo.Sample(proxyIds)
}

func deviceTemplateSelect(device *models.Device) (string, error) {

	template, err := gen.Template.Where(
		gen.Template.Platform.Eq(device.Platform),
		gen.Template.DeviceRole.Eq(device.DeviceRole),
	).First()
	if err != nil {
		core.Logger.Error(fmt.Sprintf("[templateChoice]: templateChoice for device failed with device %s, manufacturer %s, deviceRole %s", device.Id, device.Manufacturer, device.DeviceRole), zap.Error(err))
		return "", err
	}

	return template.TemplateId, nil
}

func circuitTemplateSelect() (string, error) {
	template, err := gen.Template.Where(
		gen.Template.TemplateName.Eq("template_icmp_ping_circuit"),
	).First()
	if err != nil {
		core.Logger.Error("templateChoice failed", zap.Error(err))
		return "", err
	}
	return template.Id, nil
}

func circuitHostTemplateSelect() (string, error) {
	template, err := gen.Template.Where(
		gen.Template.TemplateName.Eq("template_interface_circuit"),
	).First()
	if err != nil {
		core.Logger.Error("templateChoice failed", zap.Error(err))
		return "", err
	}
	return template.Id, nil
}

func genCircuitMacros(circuit *models.Circuit, deviceInterface *models.DeviceInterface) *[]zschema.Macro {
	macros := make([]zschema.Macro, 0)
	macros = append(macros, zschema.Macro{
		Macro: "{$IF.MAX.RX.BAND:" + deviceInterface.IfName + "}",
		Value: fmt.Sprintf("%dM", circuit.RxBandWidth),
	})
	macros = append(macros, zschema.Macro{
		Macro: "{$IF.MAX.TX.BAND:" + deviceInterface.IfName + "}",
		Value: fmt.Sprintf("%dM", circuit.TxBandWidth),
	})
	macros = append(macros, zschema.Macro{
		Macro: "{$NET.IF.IFNAME.MATCHES}",
		Value: deviceInterface.IfName,
	})
	return &macros
}

func getGlobalCommunityMacroName(enterpriseCode string) string {
	return "{$" + strings.ToUpper(enterpriseCode) + "}"
}

func snmpV2CommunitySelect(deviceId string, orgId string) (community string, port uint16) {
	cred, err := gen.SnmpV2Credential.Where(
		gen.SnmpV2Credential.DeviceId.Eq(deviceId)).Find()
	if err != nil {
		core.Logger.Error(fmt.Sprintf("[snmpV2CommunitySelect]: snmpV2CommunitySelect for device failed with device %s", deviceId), zap.Error(err))
		return "", 161
	}
	if len(cred) == 0 {
		enterpriseCode, err := getOrgEnterpriseCode(orgId)
		if err != nil {
			return "{$SNMP_COMMUNITY}", 161
		}
		return getGlobalCommunityMacroName(enterpriseCode), 161
	}
	return cred[0].Community, cred[0].Port
}

func getOrgEnterpriseCode(orgId string) (string, error) {
	org, err := gen.Organization.Select(gen.Organization.EnterpriseCode).Where(gen.Organization.Id.Eq(orgId)).First()
	if err != nil {
		core.Logger.Error(fmt.Sprintf("getOrgEnterpriseCode for organization failed with orgId %s", orgId), zap.Error(err))
		return "", err
	}
	return org.EnterpriseCode, nil
}

func getHostGroupId(siteId string) (*string, error) {

	site, err := gen.Site.Select(gen.Site.MonitorId).Where(gen.Site.Id.Eq(siteId)).First()
	if err != nil {
		core.Logger.Error(fmt.Sprintf("getHostGroupId for site failed with siteId %s", siteId), zap.Error(err))
		return nil, err
	}
	return site.MonitorId, nil
}

func getHostMonitorStatus(hostStatus string) uint8 {
	if hostStatus == "Inactive" {
		return 1
	}
	if hostStatus == "Active" {
		return 0
	}
	return 1
}

func genGroups(hgId string) []zschema.GroupID {
	return []zschema.GroupID{{GroupID: hgId}}
}

func genTemplates(templateId string) []zschema.TemplateID {
	return []zschema.TemplateID{{TemplateID: templateId}}
}

func genDeviceTags(device *models.Device) *[]zschema.Tag {

	tags := make([]zschema.Tag, 0)
	tags = append(tags, zschema.Tag{
		Tag:   "deviceRole",
		Value: device.DeviceRole,
	})
	tags = append(tags, zschema.Tag{
		Tag:   "manufacturer",
		Value: device.Manufacturer,
	})
	tags = append(tags, zschema.Tag{
		Tag:   "name",
		Value: device.Name,
	})
	tags = append(tags, zschema.Tag{
		Tag:   "organizationId",
		Value: device.OrganizationId,
	})
	return &tags
}

func genCircuitTags(circuit *models.Circuit) *[]zschema.Tag {

	tags := make([]zschema.Tag, 0)
	tags = append(tags, zschema.Tag{
		Tag:   "name",
		Value: circuit.Name,
	})
	tags = append(tags, zschema.Tag{
		Tag:   "organizationId",
		Value: circuit.OrganizationId,
	})
	return &tags
}

func genHostInterfaces(managementIp string, community string, port uint16) []zschema.HostInterfaceCreate {

	hostInterface := make([]zschema.HostInterfaceCreate, 0)
	bulk := uint8(1)
	maxRepetitions := uint8(50)
	hostInterface = append(hostInterface, zschema.HostInterfaceCreate{
		Type:    2,
		Main:    1,
		UseIp:   1,
		IP:      getPureIp(managementIp),
		Port:    uint32(port),
		Details: zschema.Details{Version: 2, Community: community, Bulk: &bulk, MaxRepetitions: &maxRepetitions},
	})

	return hostInterface
}

func genCircuitInterfaces(circuit *models.Circuit) []zschema.HostInterfaceCreate {
	var interfaces = make([]zschema.HostInterfaceCreate, 0)
	if circuit.Ipv4Address != nil && *circuit.Ipv4Address != "" {
		interfaces = genHostInterfaces(*circuit.Ipv4Address, "{$SNMP_COMMUNITY}", 161)
	} else if circuit.Ipv6Address != nil && *circuit.Ipv6Address != "" {
		interfaces = genHostInterfaces(*circuit.Ipv6Address, "{$SNMP_COMMUNITY}", 161)
	}
	return interfaces
}

// getPureIp : remove ipv4 netmask or ipv6 suffix
// example: 1.1.1.1/24 -> 1.1.1.1
// example: 5be8:dde9:7f0b:d5a7:bd01:b3be:9c69:573b/64 -> 5be8:dde9:7f0b:d5a7:bd01:b3be:9c69:573b
func getPureIp(ip string) string {
	if !strings.Contains("/", ip) || !strings.Contains("::", ip) {
		return ip
	}
	if strings.Contains("::", ip) {
		return strings.Split(ip, "::")[0]
	}

	return strings.Split(ip, "/")[0]
}
