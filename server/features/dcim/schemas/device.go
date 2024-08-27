package schemas

type DeviceCreate struct {
	Name          string `json:"name" binding:"required"`
	ManagementIp  string `json:"management_ip" binding:"required,ip_addr"`
	Status        string `json:"status" binding:"required,oneof=Active Inactive"`
	Platform      string `json:"platform" binding:"required"`
	ProductFamily string `json:"product_family" binding:"required"`
}

type DeviceShort struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	ManagementIp string `json:"management_ip"`
	Status       string `json:"status"`
}

type DeviceShortList []DeviceShort
