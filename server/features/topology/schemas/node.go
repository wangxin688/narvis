package schemas

type Node struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	ManagementIp string  `json:"managementIp"`
	DeviceRole   string  `json:"deviceRole"`
	ActiveAlerts string  `json:"activeAlerts"`
	Weight       uint16  `json:"weight"`
	Floor        *string `json:"location"`
}

type LineBase struct {
	Id                       string  `json:"id"`
	LineAlias                string  `json:"lineAlias"`
	LineStatus               string  `json:"lineStatus"`
	LocalDeviceId            string  `json:"localDeviceId"`
	LocalDeviceName          string  `json:"localDeviceName"`
	LocalDeviceManagementIp  string  `json:"localDeviceManagementIp"`
	LocalDeviceRole          string  `json:"localDeviceRole"`
	LocalIfName              string  `json:"localIfName"`
	RemoteDeviceId           string  `json:"remoteDeviceId"`
	RemoteDeviceName         string  `json:"remoteDeviceName"`
	RemoteDeviceManagementIp string  `json:"remoteDeviceManagementIp"`
	RemoteDeviceRole         string  `json:"remoteDeviceRole"`
	RemoteIfName             *string `json:"remoteIfName,omitempty"`
}

type Line struct {
	Source   string      `json:"source"`
	Target   string      `json:"target"`
	Type     string      `json:"type"`
	LineInfo []*LineBase `json:"lineInfo"`
}

type SiteTopology struct {
	// Floors []*Floors
	Nodes []*Node `json:"nodes"`
	Lines []*Line `json:"lines"`
}
