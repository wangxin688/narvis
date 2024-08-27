package schemas

import "time"

type DeviceInterface struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	IfName      string    `json:"if_name"`
	IfDescr     string    `json:"if_descr"`
	Speed       uint64    `json:"speed"`
	Mode        uint64    `json:"mode"`
	Mtu         uint64    `json:"mtu"`
	AdminStatus uint64    `json:"admin_status"`
	OperStatus  uint64    `json:"oper_status"`
	LastChange  time.Time `json:"last_change"`
}


