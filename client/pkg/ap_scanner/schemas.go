package ap_scanner

type AP struct {
	Name         string `json:"name"`
	OperStatus   string `json:"operStatus"`
	MacAddress   string `json:"macAddress"`
	SerialNumber string `json:"serialNumber"`
	ManagementIp string `json:"managementIp"`
	ActiveWacIp  string `json:"activeWacIp"`
	GroupName    string `json:"groupName"`
}
