package gossh

import (
	"fmt"
	"time"

	"github.com/scrapli/scrapligo/driver/network"
	"github.com/scrapli/scrapligo/driver/options"
	"github.com/scrapli/scrapligo/platform"
	"github.com/scrapli/scrapligo/transport"
	"github.com/wangxin688/narvis/client/pkg/gossh/commands"
)

type ConnectionInfo struct {
	ManagementIp string `json:"managementIp"`
	Port         int    `json:"port"`
	Platform     string `json:"platform"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Timeout      int    `json:"timeout"`
}

type Connection struct {
	*ConnectionInfo
	*network.Driver
}

func NewConnection(connectionInfo *ConnectionInfo) (*Connection, error) {
	driverPlatform, err := getScrapliPlatform(connectionInfo.Platform)
	if err != nil {
		return nil, fmt.Errorf("[sshConnection]: not supported platform %s: %w", connectionInfo.Platform, err)
	}
	timeout := time.Duration(connectionInfo.Timeout) * time.Second
	p, err := platform.NewPlatform(
		driverPlatform,
		connectionInfo.ManagementIp,
		options.WithAuthNoStrictKey(),
		options.WithAuthPassword(connectionInfo.Password),
		options.WithAuthUsername(connectionInfo.Username),
		options.WithPort(connectionInfo.Port),
		options.WithTimeoutOps(timeout),
		options.WithStandardTransportExtraCiphers([]string{
			"aes128-ctr", "aes192-ctr",
			"aes256-ctr", "aes128-gcm@openssh.com",
			"aes256-gcm@openssh.com", "chacha20-poly1305@openssh.com",
			"arcfour256", "arcfour128", "aes128-cbc",
			"3des-cbc", "aes192-cbc", "aes256-cbc",
		}),
		options.WithAuthNoStrictKey(),
		options.WithTransportType(transport.StandardTransport),
		options.WithSystemTransportOpenArgs(
			[]string{
				"-o",
				"KexAlgorithms=+diffie-hellman-group-exchange-sha1,diffie-hellman-group14-sha1",
				"-o",
				// note that PubkeyAcceptedKeyTypes works on older versions of openssh, whereas
				// PubKeyAcceptedAlgorithms is the option on >=8.5, runners in actions use older
				// version so we'll roll with th older type here.
				"PubkeyAcceptedKeyTypes=+ssh-rsa",
				"-o",
				"HostKeyAlgorithms=+ssh-dss,ssh-rsa,rsa-sha2-512,rsa-sha2-256,ssh-rsa,ssh-ed25519",
			},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("[sshConnection]: not supported platform %s: %w", connectionInfo.Platform, err)
	}
	driver, err := p.GetNetworkDriver()
	if err != nil {
		return nil, fmt.Errorf("[sshConnection]: failed to get network driver: %w", err)
	}
	err = driver.Open()
	if err != nil {
		return nil, fmt.Errorf("[sshConnection]: failed to open ssh connection: %w", err)
	}
	return &Connection{connectionInfo, driver}, nil
}

func (c *Connection) Close() error {
	return c.Driver.Close()
}

func (c *Connection) ShowRunningConfig() (string, error) {
	cmd, err := commands.ShowConfigurationCmd(c.Platform)
	if err != nil {
		return "", err
	}
	result, err := c.Driver.SendCommand(cmd)
	if err != nil {
		return "", err
	}
	return result.Result, nil
}
