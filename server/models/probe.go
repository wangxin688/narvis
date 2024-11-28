package models

type ProbeConfig struct {
	BaseDbModel

	ProbeType string `gorm:"column:probeType;not null"` // support http/dns/tcp/udp
	Host      string `gorm:"column:host;not null"`
	Port      uint16 `gorm:"column:port;not null"`
}