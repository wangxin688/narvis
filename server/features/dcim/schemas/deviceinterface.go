package schemas

import "time"

type DeviceInterface struct {
	Id            string    `json:"id"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
	IfIndex       uint64    `json:"ifIndex"`
	IfName        string    `json:"ifName"`
	IfDescr       string    `json:"ifDescr"`
	IfType        uint64    `json:"ifType"`
	IfMtu         uint64    `json:"ifMtu"`
	IfSpeed       uint64    `json:"ifSpeed"`
	IfPhysAddr    string    `json:"ifPhysAddr"`
	IfAdminStatus uint64    `json:"ifAdminStatus"`
	IfOperStatus  uint64    `json:"ifOperStatus"`
	IfLastChange  uint64    `json:"ifLastChange"`
	IfHighSpeed   uint64    `json:"ifHighSpeed"`
	IfIpAddress   string    `json:"ifIpAddress"`
}
