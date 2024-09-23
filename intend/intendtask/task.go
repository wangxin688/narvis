package intendtask

const ScanDeviceBasicInfo = "ScanDeviceBasicInfo"
const ScanDevice = "ScanDevice"
const ScanMacAddressTable = "ScanMacAddressTable"
const ScanAp = "ScanAp"

const ScanDeviceBasicInfoCallback = "ScanDeviceBasicInfoCallback"
const ScanDeviceCallback = "ScanDeviceCallback"
const ScanMacAddressTableCallback = "ScanMacAddressTableCallback"
const ScanApCallback = "ScanApCallback"

const DeviceBasicInfoCbUrl = "/api/v1/task/scan-device-basic"
const DeviceCbUrl = "/api/v1/task/scan-device"
const MacAddressTableCbUrl = "/api/v1/task/scan-mac"
const ApCbUrl = "/api/v1/task/scan-ap"

// 正式落库后的数据扫描schema
type BaseSnmpTask struct {
	TaskId       string            `json:"taskId"`
	SubTaskId    string            `json:"subTaskId"`
	TaskName     string            `json:"taskName"`
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
	IfIndex       uint64 `json:"ifIndex"`
	IfName        string `json:"ifName"`
	IfDescr       string `json:"ifDescr"`
	IfType        string `json:"ifType"`
	IfMtu         uint64 `json:"ifMtu"`
	IfSpeed       uint64 `json:"ifSpeed"`
	IfPhysAddr    string `json:"ifPhysAddr"`
	IfAdminStatus string `json:"ifAdminStatus"`
	IfOperStatus  string `json:"ifOperStatus"`
	IfLastChange  uint64 `json:"ifLastChange"`
	IfHighSpeed   uint64 `json:"ifHighSpeed"`
	IfIpAddress   string `json:"ifIpAddress"`
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

type DeviceScanResponse struct{}
type MacAddressTableScanResponse struct{}
type ApScanResponse struct{}
