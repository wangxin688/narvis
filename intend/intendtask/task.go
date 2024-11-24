package intendtask

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

	intend_device "github.com/wangxin688/narvis/intend/model/device"
	nettyx_wlanstation "github.com/wangxin688/narvis/intend/model/wlanstation"
	"github.com/wangxin688/narvis/intend/utils"
)

const ScanDeviceBasicInfo = "ScanDeviceBasicInfo"
const ScanDevice = "ScanDevice"
const ScanIPAM = "ScanIPAM"
const ScanMacAddressTable = "ScanMacAddressTable"
const ScanAp = "ScanAp"
const WebSSH = "WebSSH"
const ConfigurationBackup = "ConfigurationBackup"
const WlanUser = "WlanUser"

const ScanDeviceBasicInfoCallback = "ScanDeviceBasicInfoCallback"
const ScanDeviceCallback = "ScanDeviceCallback"
const ScanIPAMCallback = "ScanIPAMCallback"
const ScanMacAddressTableCallback = "ScanMacAddressTableCallback"
const ScanApCallback = "ScanApCallback"
const ConfigurationBackupCallback = "ConfigurationBackupCallback"
const WlanUserCallback = "WlanUserCallback"

const DeviceBasicInfoCbUrl = "/api/v1/task/scan-device-basic"
const DeviceCbUrl = "/api/v1/task/scan-devices"
const MacAddressTableCbUrl = "/api/v1/task/scan-mac"
const ApCbUrl = "/api/v1/task/scan-aps"
const WebSocketCbUrl = "/api/v1/webssh/proxy"
const ConfigurationBackupCbUrl = "/api/v1/task/config-backup"
const WlanUserCbUrl = "/api/v1/task/wlan-users"

// 正式落库后的数据扫描schema
type BaseSnmpTask struct {
	TaskId       string            `json:"taskId"`
	SubTaskId    string            `json:"subTaskId"`
	TaskName     string            `json:"taskName"`
	SiteId       string            `json:"siteId"`
	DeviceId     string            `json:"deviceId"`
	ManagementIp string            `json:"managementIp"`
	Callback     string            `json:"callback"`
	SnmpConfig   *SnmpV2Credential `json:"snmpConfig"`
}

// 基于网段的基础设备信息扫描schema
type BaseSnmpScanTask struct {
	TaskId     string            `json:"taskId"`
	TaskName   string            `json:"taskName"`
	Callback   string            `json:"callback"`
	SnmpConfig *SnmpV2Credential `json:"snmpConfig"`
	Range      string            `json:"range"` // CIDR
}

type WebSSHTask struct {
	TaskName     string `json:"taskName"`
	SessionId    string `json:"sessionId"`
	ManagementIP string `json:"managementIp"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Port         uint16 `json:"port"`
	Rows         int    `json:"rows"`
	Cols         int    `json:"cols"`
}

type ConfigurationBackupTask struct {
	TaskId       string `json:"taskId"`
	TaskName     string `json:"taskName"`
	Callback     string `json:"callback"`
	DeviceId     string `json:"deviceId"`
	ManagementIp string `json:"managementIp"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Port         uint16 `json:"port"`
	Platform     string `json:"platform"`
}

type ConfigurationBackupTaskResult struct {
	Configuration string `json:"configuration"`
	DeviceId      string `json:"deviceId"`
	BackupTime    string `json:"backupTime"`
	HashValue     string `json:"hashValue"`
	Error         string `json:"error"`
}

type SnmpV2Credential struct {
	Community      string `json:"community"`
	MaxRepetitions uint8  `json:"maxRepetitions"`
	Timeout        uint8  `json:"timeout"`
	Port           uint16 `json:"port"`
}

type DeviceBasicInfo struct {
	Name         string `json:"name"`
	Description  string `json:"description"`
	ChassisId    string `json:"chassisId"`
	ManagementIp string `json:"managementIp"`
}

type Entity struct {
	EntityPhysicalClass       string `json:"entityPhysicalClass"`
	EntityPhysicalDescr       string `json:"entityPhysicalDescr"`
	EntityPhysicalName        string `json:"entityPhysicalName"`
	EntityPhysicalSoftwareRev string `json:"entityPhysicalSoftwareRev"`
	EntityPhysicalSerialNum   string `json:"entityPhysicalSerialNum"`
}

type Stack struct {
	Id         string `json:"id"`
	Priority   uint32 `json:"priority"`
	Role       string `json:"role"`
	MacAddress string `json:"macAddress"`
}

type DeviceScanResponse struct {
	DeviceId       string                           `json:"deviceId"`
	SiteId         string                           `json:"siteId"`
	Name           string                           `json:"name"`
	Description    string                           `json:"description"`
	ChassisId      *string                          `json:"chassisId"`
	ManagementIp   string                           `json:"managementIp"`
	Manufacturer   string                           `json:"manufacturer"`
	DeviceModel    string                           `json:"deviceModel"`
	Platform       string                           `json:"platform"`
	OrganizationId string                           `json:"organizationId"`
	Interfaces     []*intend_device.DeviceInterface `json:"interfaces"`
	LldpNeighbors  []*intend_device.LldpNeighbor    `json:"lldpNeighbors"`
	Entities       []*intend_device.Entity          `json:"entities"`
	Stacks         []*intend_device.Stack           `json:"stacks"`
	Vlans          []*intend_device.VlanItem        `json:"vlans"`
	ArpTable       []*intend_device.ArpItem         `json:"arpTable"`
	Errors         []string                         `json:"errors"`
	SnmpReachable  bool                             `json:"snmpReachable"`
	SshReachable   bool                             `json:"sshReachable"`
	IcmpReachable  bool                             `json:"icmpReachable"`
}

type DeviceBasicInfoScanResponse struct {
	Name           string   `json:"name"`
	Description    string   `json:"description"`
	ChassisId      string   `json:"chassisId"`
	ManagementIp   string   `json:"managementIp"`
	Manufacturer   string   `json:"manufacturer"`
	DeviceModel    string   `json:"deviceModel"`
	Platform       string   `json:"platform"`
	OrganizationId string   `json:"organizationId"`
	Errors         []string `json:"errors"`
}

type MacAddressTableScanResponse struct{}

type ApScanResponse struct {
	Name            string  `json:"name"`
	MacAddress      *string `json:"macAddress"`
	SerialNumber    *string `json:"serialNumber"`
	ManagementIp    string  `json:"managementIp"`
	GroupName       *string `json:"groupName"`
	DeviceModel     string  `json:"deviceModel"`
	Manufacturer    string  `json:"manufacturer"`
	WlanACIpAddress *string `json:"wlanACIpAddress"`
	OsVersion       *string `json:"osVersion"`
	SiteId          string  `json:"siteId"`
	OrganizationId  string  `json:"organizationId"`
}

func (a *ApScanResponse) CalApHash() string {
	hashString := fmt.Sprintf(
		"%s-%s-%s-%s-%s-%s",
		a.Name,
		a.ManagementIp,
		utils.PtrStringToString(a.MacAddress),
		utils.PtrStringToString(a.GroupName),
		utils.PtrStringToString(a.WlanACIpAddress),
		utils.PtrStringToString(a.SerialNumber),
	)
	hash := md5.New()
	_, _ = hash.Write([]byte(hashString))
	return hex.EncodeToString(hash.Sum(nil))
}

type WlanUserTaskResult struct {
	Errors         []string                       `json:"errors"`
	WlanUsers      []*nettyx_wlanstation.WlanUser `json:"wlanUsers"`
	SiteId         string                         `json:"siteId"`
	DeviceId       string                         `json:"deviceId"`
	OrganizationId string                         `json:"organizationId"`
}
