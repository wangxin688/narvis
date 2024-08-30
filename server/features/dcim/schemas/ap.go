package schemas

import "time"

type ApCoordinate struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
	Z float32 `json:"z"`
}

type AP struct {
	Id           int           `json:"id"`
	CreatedAt    time.Time     `json:"createdAt"`
	UpdatedAt    time.Time     `json:"updatedAt"`
	Name         string        `json:"name"`
	Status       string        `json:"status"`
	MacAddress   *string       `json:"macAddress"`
	SerialNumber *string       `json:"serialNumber"`
	ManagementIP string        `json:"managementIp"`
	DeviceModel  string        `json:"deviceType"`
	Manufacturer string        `json:"manufacturer"`
	DeviceRole   string        `json:"deviceRole"`
	Version      *string       `json:"version"`
	GroupName    *string       `json:"groupName"`
	Coordinate   *ApCoordinate `json:"coordinate"`
	ActiveWac    DeviceShort   `json:"activeWac"`
	Location     LocationShort `json:"location"`
	Site         SiteShort     `json:"site"`
}

type APList []AP

type APShort struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	ManagementIP string `json:"managementIp"`
}

type APShortList []APShort
