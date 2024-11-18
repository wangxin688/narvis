package intendtask

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"

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

type DeviceInterface struct {
	IfIndex       uint64  `json:"ifIndex"`
	IfName        string  `json:"ifName"`
	IfDescr       string  `json:"ifDescr"`
	IfType        string  `json:"ifType"`
	IfMtu         uint64  `json:"ifMtu"`
	IfSpeed       uint64  `json:"ifSpeed"`
	IfPhysAddr    *string `json:"ifPhysAddr"`
	IfAdminStatus string  `json:"ifAdminStatus"`
	IfOperStatus  string  `json:"ifOperStatus"`
	IfLastChange  uint64  `json:"ifLastChange"`
	IfHighSpeed   uint64  `json:"ifHighSpeed"`
	IfIpAddress   *string `json:"ifIpAddress"`
	HashValue     string  `json:"hashValue"`
}

func (d *DeviceInterface) CalHashValue() string {

	hashString := fmt.Sprintf(
		"%s-%s-%s-%d-%d-%d-%s-%s-%s-%d-%s",
		d.IfName,
		d.IfDescr,
		d.IfType,
		d.IfHighSpeed,
		d.IfMtu,
		d.IfSpeed,
		utils.PtrStringToString(d.IfPhysAddr),
		d.IfAdminStatus,
		d.IfOperStatus,
		d.IfLastChange,
		utils.PtrStringToString(d.IfIpAddress))
	hash := md5.New()
	_, _ = hash.Write([]byte(hashString))
	return hex.EncodeToString(hash.Sum(nil))
}

type LldpNeighbor struct {
	LocalChassisId  string `json:"localChassisId"`
	LocalHostname   string `json:"localHostname"`
	LocalIfName     string `json:"localIfName"`
	LocalIfDescr    string `json:"localIfDescr"`
	RemoteChassisId string `json:"remoteChassisId"`
	RemoteHostname  string `json:"remoteHostname"`
	RemoteIfName    string `json:"remoteIfName"`
	RemoteIfDescr   string `json:"remoteIfDescr"`
	HashValue       string `json:"hashValue"`
}

func (l *LldpNeighbor) CalHashValue() string {
	hashString := fmt.Sprintf(
		"%s-%s-%s-%s-%s-%s",
		l.LocalChassisId,
		l.LocalIfName,
		l.LocalIfDescr,
		l.RemoteChassisId,
		l.RemoteIfName,
		l.RemoteIfDescr)
	hash := md5.New()
	_, _ = hash.Write([]byte(hashString))
	return hex.EncodeToString(hash.Sum(nil))
}

func (l *LldpNeighbor) CalApHashValue() string {
	hashString := fmt.Sprintf(
		"%s-%s-%s-%s",
		l.LocalChassisId,
		l.LocalIfName,
		l.LocalIfDescr,
		l.RemoteChassisId)
	hash := md5.New()
	_, _ = hash.Write([]byte(hashString))
	return hex.EncodeToString(hash.Sum(nil))
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

type VlanItem struct {
	VlanId   uint32 `json:"vlanId"`
	VlanName string `json:"vlanName"`
	IfIndex  uint64 `json:"ifIndex"`
	Range    string `json:"range"`
	Gateway  string `json:"gateway"`
}

type ArpItem struct {
	IpAddress  string `json:"ipAddress"`
	MacAddress string `json:"macAddress"`
	Type       string `json:"type"`
	IfIndex    uint64 `json:"ifIndex"`
	VlanId     uint32 `json:"vlanId"`
	Range      string `json:"range"`
	HashValue  string `json:"hashValue"`
}

type DeviceScanResponse struct {
	DeviceId       string             `json:"deviceId"`
	SiteId         string             `json:"siteId"`
	Name           string             `json:"name"`
	Description    string             `json:"description"`
	ChassisId      *string            `json:"chassisId"`
	ManagementIp   string             `json:"managementIp"`
	Manufacturer   string             `json:"manufacturer"`
	DeviceModel    string             `json:"deviceModel"`
	Platform       string             `json:"platform"`
	OrganizationId string             `json:"organizationId"`
	Interfaces     []*DeviceInterface `json:"interfaces"`
	LldpNeighbors  []*LldpNeighbor    `json:"lldpNeighbors"`
	Entities       []*Entity          `json:"entities"`
	Stacks         []*Stack           `json:"stacks"`
	Vlans          []*VlanItem        `json:"vlans"`
	ArpTable       []*ArpItem         `json:"arpTable"`
	Errors         []string           `json:"errors"`
	SnmpReachable  bool               `json:"snmpReachable"`
	SshReachable   bool               `json:"sshReachable"`
	IcmpReachable  bool               `json:"icmpReachable"`
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
	Errors         []string        `json:"errors"`
	WlanUsers      []*WlanUserItem `json:"wlanUsers"`
	SiteId         string          `json:"siteId"`
	DeviceId       string          `json:"deviceId"`
	OrganizationId string          `json:"organizationId"`
}

type WlanUserItem struct {
	StationMac           string  `json:"stationMac"`                     // 终端MAC
	StationIp            string  `json:"stationIp"`                      // 终端IP
	StationUsername      string  `json:"stationUsername"`                // 终端用户名
	StationApMac         *string `json:"stationApMac,omitempty"`         // AP MAC
	StationApName        *string `json:"stationApName,omitempty"`        // AP 名称
	StationESSID         string  `json:"stationESSID"`                   // ESSID
	StationVlan          *uint64 `json:"stationVlan,omitempty"`          // VLAN
	StationChannel       uint64  `json:"stationChannel,omitempty"`       // 信道
	StationChanBandWidth *string `json:"stationChanBandWidth,omitempty"` // 信道带宽
	StationRadioType     string  `json:"stationRadioType"`               // radio类型
	StationSNR           *uint64 `json:"stationSNR,omitempty"`           // 终端NR
	StationRSSI          uint64  `json:"stationRSSI"`                    // 终端RSSI
	StationRxBits        uint64  `json:"stationRxBits"`                  // 终端下行流量
	StationTxBits        uint64  `json:"stationTxBits"`                  // 终端上行流量
	StationMaxSpeed      *uint64 `json:"stationMaxSpeed,omitempty"`      // 终端协商速率
	StationOnlineTime    uint64  `json:"stationOnlineTime"`              // 终端在线时间
}
