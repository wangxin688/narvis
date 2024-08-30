package schemas

type DeviceCreate struct {
	Name          string `json:"name" binding:"required"`
	ManagementIp  string `json:"managementIp" binding:"required,ip_addr"`
	Status        string `json:"status" binding:"required,oneof=Active Inactive"`
	Platform      string `json:"platform" binding:"required"`
	ProductFamily string `json:"productFamily" binding:"required"`
}

type DeviceShort struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	ManagementIp string `json:"managementIp"`
	Status       string `json:"status"`
}

type DeviceShortList []DeviceShort
