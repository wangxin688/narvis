package schemas

import "time"

type ApCoordinate struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

type AP struct {
	ID           int           `json:"id"`
	CreatedAt    time.Time     `json:"created_at"`
	UpdatedAt    time.Time     `json:"updated_at"`
	Name         string        `json:"name"`
	Status       string        `json:"status"`
	MacAddress   *string       `json:"mac_address"`
	SerialNumber *string       `json:"serial_number"`
	ManagementIP string        `json:"management_ip"`
	DeviceModel  string        `json:"device_type"`
	Manufacturer string        `json:"manufacturer"`
	DeviceRole   string        `json:"device_role"`
	Version      *string       `json:"version"`
	GroupName    *string       `json:"group_name"`
	Coordinate   *ApCoordinate `json:"coordinate"`
	ActiveWac    DeviceShort   `json:"active_wac"`
	Location     LocationShort `json:"location"`
	Site         SiteShort     `json:"site"`
}

type APList []AP

type APShort struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	ManagementIP string `json:"management_ip"`
}

type APShortList []APShort
