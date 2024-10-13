package schemas

type DevicePanel struct {
	Slots      []*Slot            `json:"slots"`
	Interfaces []*PanelInterface `json:"interfaces"`
}

type Slot struct {
	Id            string `json:"id"`
	Label         string `json:"label"`
	NumberOfPorts int    `json:"numberOfPorts"`
}

type PanelInterface struct {
	Id            string  `json:"id"`
	IfIndex       uint64  `json:"ifIndex"`
	IfName        string  `json:"ifName"`
	IfDescr       string  `json:"ifDescr"`
	IfType        string  `json:"ifType"`
	IfMtu         uint64  `json:"ifMtu"`
	IfSpeed       uint64  `json:"ifSpeed"`
	IfPhysAddr    *string `json:"ifPhysAddr"`
	IfAdminStatus string  `json:"ifAdminStatus"`
	IfOperStatus  string  `json:"ifOperStatus"`
	IfLastChange  string  `json:"ifLastChange"`
	IfHighSpeed   uint64  `json:"ifHighSpeed"`
	IfIpAddress   *string `json:"ifIpAddress"`
	ComboId       *string `json:"comboId"`
	Label         *string `json:"label"`
	Type          string  `json:"type"`
}
