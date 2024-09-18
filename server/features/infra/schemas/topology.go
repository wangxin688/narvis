package schemas

import "fmt"

type Link struct {
	Id                       string `json:"id"`
	LineAlias                string `json:"lineAlias"`
	LocalDeviceId            string `json:"localDeviceId"`
	LocalDeviceManagementIp  string `json:"localDeviceManagementIp"`
	LocalDeviceName          string `json:"localDeviceName"`
	LocalDeviceRole          string `json:"localDeviceRole"`
	LocalIfName              string `json:"localIfName"`
	LocalIfDescr             string `json:"localIfDescr"`
	RemoteDeviceId           string `json:"remoteDeviceId"`
	RemoteDeviceManagementIp string `json:"remoteDeviceManagementIp"`
	RemoteDeviceName         string `json:"remoteDeviceName"`
	RemoteDeviceRole         string `json:"remoteDeviceRole"`
	RemoteIfName             string `json:"remoteIfName"`
	RemoteIfDescr            string `json:"remoteIfDescr"`
}

func (l *Link) GenerateAlias() {
	l.LineAlias = fmt.Sprintf("%s:%s<->%s:%s", l.LocalDeviceName, l.LocalIfName, l.RemoteDeviceName, l.RemoteIfName)
}

type Node struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	ManagementIp string `json:"managementIp"`
	DeviceRole   string `json:"deviceRole"`
	Manufacturer string `json:"manufacturer"`
	DeviceModel  string `json:"deviceModel"`
	NumOfAlerts  int    `json:"numOfAlerts"`
	HealthScore  int    `json:"healthScore"`
	Weight       uint16 `json:"weight"`
	PathId       string `json:"pathId"` // ID of the path between two nodes
}

type ClientNode struct {
	RadioType    string `json:"radioType"`
	Channel      string `json:"channel"`
	NumOfClients int    `json:"numOfClients"`
}

type Topology struct {
	SiteId string `json:"siteId"`
	Nodes  []Node `json:"nodes"`
	Links  []Link `json:"links"`
}
