package biz

type MonitorSite struct {
	SiteID string
}

func NewMonitorSite(siteID string) *MonitorSite {
	return &MonitorSite{
		SiteID: siteID,
	}
}

func (m *MonitorSite) Create(siteID string) {
}

func (m *MonitorSite) Update(siteID string) {

}

func (m *MonitorSite) Delete(siteID string) {

}
