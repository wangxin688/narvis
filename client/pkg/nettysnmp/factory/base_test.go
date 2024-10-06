package factory

import (
	"testing"
	"time"

	"github.com/gosnmp/gosnmp"
	"go.uber.org/zap"
)

func TestSnmpDiscovery_SysDescr(t *testing.T) {
	community := "public"
	baseConfig := BaseSnmpConfig{
		Port:           161,
		Version:        gosnmp.Version2c,
		Community:      &community,
		Timeout:        5,
		MaxRepetitions: 50,
	}
	snmpConfig := SnmpConfig{
		IpAddress:      "127.0.0.1",
		BaseSnmpConfig: baseConfig,
	}
	if !snmpConfig.validate() {
		t.Fatalf("failed to validate snmp config")
	}
	session, err := NewSnmpDiscovery(snmpConfig)
	if err != nil {
		t.Fatalf("failed to create snmp session: %s", zap.Error(err))
	}
	discovery := session.Discovery()

	t.Logf("sysDescr: %v", discovery)

}

func TestGetSysDescr(t *testing.T) {
	community := "public"
	session := gosnmp.GoSNMP{
		Target:    "127.0.0.1",
		Port:      161,
		Version:   gosnmp.Version2c,
		Community: community,
		Timeout:   time.Duration(2) * time.Second,
		Retries:   2,
		MaxOids:   50,
	}
	err := session.Connect()
	if err != nil {
		t.Fatalf("failed to connect: %s", zap.Error(err))
	}
	defer session.Conn.Close()
	result, err := session.Get([]string{SysDescr})
	if err != nil {
		t.Fatalf("failed to get sysDescr: %s", zap.Error(err))
	}
	t.Logf("sysDescr: %s", result.Variables[0].Value)
}

func TestMac(t *testing.T) {
	community := "public"
	session := gosnmp.GoSNMP{
		Target:    "127.0.0.1",
		Port:      161,
		Version:   gosnmp.Version2c,
		Community: community,
		Timeout:   time.Duration(2) * time.Second,
		Retries:   2,
		MaxOids:   50,
	}
	err := session.Connect()
	if err != nil {
		t.Fatalf("failed to connect: %s", zap.Error(err))
	}
	defer session.Conn.Close()
	result, err := session.BulkWalkAll(".1.3.6.1.2.1.2.2.1.1")
	if err != nil {
		t.Fatalf("failed to get sysDescr: %s", zap.Error(err))
	}
	for _, variable := range result {
		value := variable.Value.([]byte)
		vali := hex2mac(value)
		t.Logf("mac: %s", vali)

	}
}
