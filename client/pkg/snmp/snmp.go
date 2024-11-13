package snmp

import (
	"fmt"
	"time"

	"github.com/gosnmp/gosnmp"
	"github.com/wangxin688/narvis/client/utils/logger"
	"go.uber.org/zap"
)

type SnmpV3Params struct {
	ContextEngineId *string
	SecurityName    *string
	SecurityLevel   gosnmp.SnmpV3SecurityModel
	AuthProtocol    gosnmp.SnmpV3AuthProtocol
	AuthPassword    *string
	PrivProtocol    gosnmp.SnmpV3PrivProtocol
	PrivPassword    *string
}

type SnmpDiscovery struct {
	Session   *gosnmp.GoSNMP
	IpAddress string
}

type BaseSnmpConfig struct {
	Port           uint16
	Version        gosnmp.SnmpVersion
	Timeout        uint8
	Community      *string
	V3Params       *SnmpV3Params
	MaxRepetitions int
}

type SnmpConfig struct {
	IpAddress string
	BaseSnmpConfig
}

func NewSnmpSession(config SnmpConfig) (*gosnmp.GoSNMP, error) {
	var snmpSession *gosnmp.GoSNMP
	if !config.validate() {
		return nil, fmt.Errorf("invalid snmp config parameters for %s", config.IpAddress)
	}
	snmpSession = &gosnmp.GoSNMP{
		Target:   config.IpAddress,
		Port:     config.Port,
		Timeout:  time.Duration(config.Timeout) * time.Second,
		Retries:  2,
		MaxOids:  int(config.MaxRepetitions),
		Version:  config.Version,
		MsgFlags: gosnmp.AuthPriv,
	}

	switch config.Version {
	case gosnmp.Version2c:
		snmpSession.Community = *config.Community
	case gosnmp.Version3:
		snmpSession.SecurityParameters = &gosnmp.UsmSecurityParameters{
			UserName:                 *config.V3Params.SecurityName,
			AuthenticationProtocol:   config.V3Params.AuthProtocol,
			AuthenticationPassphrase: *config.V3Params.AuthPassword,
			PrivacyProtocol:          config.V3Params.PrivProtocol,
			PrivacyPassphrase:        *config.V3Params.PrivPassword,
		}
	default:
		return nil, fmt.Errorf("unsupported snmp version: %d", config.Version)
	}

	err := snmpSession.Connect()
	if err != nil {
		snmpSession.Conn.Close()
		logger.Logger.Info("snmp connect error", zap.String("ip", config.IpAddress), zap.Error(err))
		return nil, err
	}
	return snmpSession, nil
}

func NewSnmpDiscovery(sc SnmpConfig) (*SnmpDiscovery, error) {
	session, err := NewSnmpSession(sc)
	if err != nil {
		return nil, err
	}
	return &SnmpDiscovery{
		Session:   session,
		IpAddress: session.Target,
	}, nil
}

// validate checks if the snmp config is valid
func (c *SnmpConfig) validate() bool {
	switch c.Version {
	case gosnmp.Version2c:
		return c.Community != nil
	case gosnmp.Version3:
		return c.V3Params != nil
	default:
		return false
	}
}
