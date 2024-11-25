// Copyright 2024 wangxin.jeffry@gmail.com
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package netdisco

import (
	"fmt"
	"net"
	"time"

	"github.com/go-ping/ping"
	"github.com/gosnmp/gosnmp"
	"github.com/wangxin688/narvis/intend/logger"
	"github.com/wangxin688/narvis/intend/model/snmp"
	"github.com/wangxin688/narvis/intend/netdisco/factory"
	"go.uber.org/zap"
)

type Dispatcher struct {
}

// SnmpReachable returns true if the target is reachable via SNMP.
func (d *Dispatcher) SnmpReachable(session *gosnmp.GoSNMP) bool {
	// Try to get the sysName from the target, if it fails, it's not reachable.
	result, err := session.Get([]string{factory.SysName})
	if err != nil || result == nil {
		logger.Logger.Info("[dispatcher]: SnmpReachable failed", zap.Error(err), zap.String("target", session.Target))
		return false
	}
	// If the result is not empty, the target is reachable.
	return len(result.Variables) > 0
}

// IcmpReachable returns true if the target is reachable via ICMP.
// linux need privilege for udp
func (d *Dispatcher) IcmpReachable(address string) bool {
	pinger, err := ping.NewPinger(address)
	if err != nil {
		logger.Logger.Info("[dispatcher]: IcmpReachable failed", zap.String("target", address), zap.Error(err))
		return false
	}
	pinger.Count = 2
	pinger.Interval = time.Duration(100) * time.Millisecond
	pinger.Timeout = time.Second
	err = pinger.Run()
	if err != nil {
		return false
	}
	return pinger.Statistics().PacketsRecv > 0
}

// SshReachable returns true if the target is reachable via SSH. default port is 22
func (d *Dispatcher) SshReachable(address string) bool {
	timeout := time.Second
	conn, err := net.DialTimeout("tcp", address+":22", timeout)
	if err != nil {
		logger.Logger.Info("[dispatcher]: SshReachable failed", zap.String("target", address), zap.Error(err))
		return false
	}
	if conn == nil {
		logger.Logger.Info("[dispatcher]: SshReachable failed", zap.String("target", address), zap.String("reason", "connection is nil"))
		return false
	}
	defer func() {
		if conn != nil {
			_ = conn.Close()
		}
	}()
	return true
}

// Session returns a new GoSNMP session.
func (d *Dispatcher) Session(config *snmp.SnmpConfig) (*gosnmp.GoSNMP, error) {
	session, err := factory.NewSnmpSession(config)
	if err != nil {
		return nil, err
	}
	return session, nil
}

// SysObjectID returns the sysObjectID from RFC-1213MIB
func (d *Dispatcher) SysObjectID(session *gosnmp.GoSNMP) string {
	oid := []string{factory.SysObjectID}
	result, err := session.Get(oid)
	if err != nil {
		return ""
	}

	for _, variable := range result.Variables {
		if variable.Type == gosnmp.ObjectIdentifier {
			return fmt.Sprintf("%s", variable.Value)
		}
	}
	return ""
}

// Driver returns a new SnmpDriver.
func (d *Dispatcher) Driver(config *snmp.SnmpConfig) (SnmpDriver, error) {
	snmpReachable := true
	session, err := d.Session(config)
	if err != nil || session == nil {
		snmpReachable = false
	} else {
		snmpReachable = d.SnmpReachable(session)
	}
	result := map[string]string{
		"snmpReachable": fmt.Sprintf("%t", snmpReachable),
	}

	if !snmpReachable {
		return nil, fmt.Errorf("target %s snmp unreachable: %v", config.IpAddress, result)
	}
	if config.Platform != nil && *config.Platform != "" {
		return getFactory(*config.Platform, config)
	}
	sysObjectId := d.SysObjectID(session)
	if sysObjectId == "" {
		return nil, fmt.Errorf("Unable to get sysObjectId, target manufacturer is not supported: %s", config.IpAddress)
	}
	deviceType := getDeviceModel(sysObjectId)
	return getFactory(deviceType.Platform, config)
}

func NewNetDisco() *Dispatcher {
	return &Dispatcher{}
}
