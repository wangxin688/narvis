package schemas

import "time"

type DeviceHealthQuery struct {
	SiteId *string `form:"siteId" binding:"omitempty,uuid"`
}

type DeviceHealthTrendQuery struct {
	SiteId    *string `form:"siteId" binding:"omitempty,uuid"`
	StartedAt time.Time
	EndedAt   time.Time
}

type HealthResponse struct {
	Firewall HealthHeatMap `json:"Firewall"`
	WlanAP  HealthHeatMap `json:"WlanAP"`
	Switch  HealthHeatMap `json:"Switch"`
	Device  HealthHeatMap `json:"Device"`
	Circuit HealthHeatMap `json:"Circuit"`
}
