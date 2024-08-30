package nettyssh

import "github.com/wangxin688/narvis/client/pkg/nettyssh/types"

type DeviceOption func(interface{}) error

func SecretOption(secret string) func(device interface{}) error {
	return func(device interface{}) error {
		device.(types.CiscoDevice).SetSecret(secret)
		return nil
	}
}

func TimeoutOption(timeout uint8) func(device interface{}) error {
	return func(device interface{}) error {
		device.(types.CiscoDevice).SetTimeout(timeout)
		return nil
	}
}
