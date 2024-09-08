package zschema

import (
	"encoding/json"
	"fmt"
)

type ZbxRequest struct {
	JsonRpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  any    `json:"params"`
	ID      uint64 `json:"id"`
}

type ZbxApiError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ZbxResponse struct {
	JsonRpc string          `json:"jsonrpc"`
	Result  json.RawMessage `json:"result"`
	Error   *ZbxApiError    `json:"error,omitempty"`
	ID      uint64          `json:"id"`
}

func (rsp *ZbxResponse) HasError() error {
	if rsp.Error != nil && rsp.Error.Code != 0 {
		return fmt.Errorf("request zbx system error: code:%d data %s, message: %s", rsp.Error.Code, rsp.Error.Data, rsp.Error.Message)
	}
	return nil
}

func (rsp *ZbxResponse) GetResult(v any) {
	json.Unmarshal(rsp.Result, &v)
}

type Tag struct {
	Tag   string
	Value string
}

type Macro struct {
	Macro string
	Value string
}

type GroupID struct {
	GroupID string `json:"groupid"`
}
type HostID struct {
	HostID string `json:"hostid"`
}
type TemplateID struct {
	TemplateID string `json:"templateid"`
}

type HostGroup struct {
	GroupID string `json:"groupid"`
	Name    string `json:"name"`
}

type HostGroupCreate struct {
	Name string `json:"name"`
}

type HostGroupCreateResult struct {
	GroupIDs []string `json:"groupids"`
}

type HostGroupUpdate struct {
	GroupID string `json:"groupid"`
	Name    string `json:"name"`
}

type HostGroupUpdateResult struct {
	GroupIDs []string `json:"groupid"`
}

type HostGroupGet struct {
	GroupIDs *[]string `json:"groupids,omitempty"`
	HostIDs  *[]string `json:"hostids,omitempty"`
}

type HostGroupDeleteResult struct {
	GroupIDs []string `json:"groupids"`
}

type Host struct {
	HostID          string `json:"hostid"`
	Host            string `json:"host"`
	MonitoredBy     uint8  `json:"monitored_by"`
	AssignedProxyID uint   `json:"proxyid"`
	Status          uint8  `json:"status"`
}

type HostCreate struct {
	Host        string          `json:"host"`
	Interfaces  []HostInterface `json:"interfaces"`
	Groups      []GroupID       `json:"groups"`
	Tags        []Tag           `json:"tags"`
	Macros      []Macro         `json:"macros"`
	MonitoredBy uint8           `json:"monitored_by"` // 0:server 1:proxy 2:proxy_group
	ProxyID     uint            `json:"proxyid"`
	Status      uint8           `json:"status"` // 0:enable 1:disabled
	Templates   []TemplateID    `json:"templates"`
}

type HostCreateResult struct {
	HostIDs []string `json:"hostids"`
}

type HostGet struct {
	GroupIDs    *[]string `json:"groupids,omitempty"`
	HostIDs     *[]string `json:"hostids,omitempty"`
	ProxyIDs    *[]string `json:"proxyids,omitempty"`
	TemplateIDs *[]string `json:"templateids,omitempty"`
	Output      string    `json:"output,omitempty"` // extend /shorten
}

type HostUpdate struct {
	HostID        string                 `json:"hostid"`
	Host          *string                `json:"host,omitempty"`
	Groups        *[]GroupID             `json:"groups,omitempty"`
	Interface     *[]HostInterfaceUpdate `json:"interfaces,omitempty"`
	TemplateClear *[]string              `json:"template_clear,omitempty"` // clear and replace templates
	Tags          *[]Tag                 `json:"tags,omitempty"`
	Macros        *[]Macro               `json:"macros,omitempty"`
	Status        *uint8                 `json:"status,omitempty"`
	ProxyID       *uint                  `json:"proxyid,omitempty"`
}

type HostUpdateResult struct {
	HostIDs []string `json:"hostids"`
}

type HostDeleteResult struct {
	HostIDs []string `json:"hostids"`
}

type HostMassUpdate struct {
	Hosts   []HostID   `json:"hostids"`
	Status  *uint8     `json:"status,omitempty"` // 0:enable 1:disabled
	Groups  *[]GroupID `json:"groups,omitempty"`
	ProxyID *uint      `json:"proxyid,omitempty"`
}

type HostInterface struct {
	Type    uint8   `json:"type"`  // 1:agent 2:snmp 3:ipmi 4:jmx
	Main    uint8   `json:"main"`  // 0:no 1:yes
	UseIp   uint8   `json:"useip"` // 0:use dns 1:use ip
	IP      string  `json:"ip"`
	Port    uint32  `json:"port"` // default agent 10050, snmp 161
	Details Details `json:"details"`
}

type HostInterfaceUpdate struct {
	InterfaceID string   `json:"interfaceid"`
	Ip          *string  `json:"ip,omitempty"`
	Port        *uint32  `json:"port,omitempty"`
	Details     *Details `json:"details,omitempty"`
}

type Details struct {
	Version   uint8  `json:"version"`
	Community string `json:"community"`
	Bulk      uint8  `json:"bulk"` // 0:no 1:yes
}

type TemplateCreate struct {
	Host      string       `json:"host"`
	Groups    []GroupID    `json:"groups"`
	Templates []TemplateID `json:"templates"`
}
type TemplateCreateResult struct {
	TemplateIDs []string `json:"templateids"`
}

type TemplateGet struct {
	TemplateIDs *[]string `json:"templateids,omitempty"`
	GroupIDs    *[]string `json:"groupids,omitempty"`
	Output      *string   `json:"output,omitempty"`
}

type TemplateUpdate struct {
	TemplateID string        `json:"templateid"`
	Host       *string       `json:"host,omitempty"`
	Groups     *[]GroupID    `json:"groups,omitempty"`
	Templates  *[]TemplateID `json:"templates,omitempty"`
}

type TemplateUpdateResult struct {
	TemplateIDs []string `json:"templateids"`
}

type TemplateDeleteResult struct {
	TemplateIDs []string `json:"templateids"`
}

type TemplateGroupCreate struct {
	Name string `json:"name"`
}

type TemplateGroupCreateResult struct {
	GroupIDs []string `json:"groupids"`
}

type TemplateGroupGet struct {
	GroupIDs    *[]string          `json:"groupids,omitempty"`
	TemplateIDs *[]string          `json:"templateids,omitempty"`
	Filter      *map[string]string `json:"filter,omitempty"`
}

type ProxyCreate struct {
	Name           string `json:"name"`
	OperationMode  uint8  `json:"operation_mode"`   // 0: active 1:passive
	TlsConnect     uint8  `json:"tls_connect"`      // 1: No encryption.2: PSK.3: certificate
	TlsPskIDentity string `json:"tls_psk_identity"` // proxy config, can keep it as host
	TlsAccept      uint8  `json:"tls_accept"`       // 1: No encryption.2: PSK.3: certificate
	TlsPsk         string `json:"tls_psk"`          // pre-shared key
}

type ProxyCreateResult struct {
	ProxyIDs []string `json:"proxyids"`
}

type ProxyGet struct {
	ProxyIDs *[]string `json:"proxyids,omitempty"`
	Output   *string   `json:"output,omitempty"`
}

type ProxyDeleteResult struct {
	ProxyIDs []string `json:"proxyids"`
}

type Event struct {
	EventID string `json:"eventid"`
	Status  uint8  `json:"status"`
}

type EventGet struct {
	EventIDs []string `json:"eventids,omitempty"`
	Output   string   `json:"output,omitempty"`
}

type EventAcknowledge struct {
	EventIDs []string `json:"eventids"`
	Action   uint32   `json:"action"` // 1: close
}

type EventAcknowledgeResult struct {
	EventIDs []string `json:"eventids"`
}
