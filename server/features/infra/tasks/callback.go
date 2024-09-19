package infra_tasks

import (
	"encoding/json"

	"github.com/wangxin688/narvis/intend/intendtask"
)

func deviceBasicInfoScanCallback(data []byte) error {
	basicInfo := make([]*intendtask.DeviceBasicInfoScanResponse, 0)
	err := json.Unmarshal(data, &basicInfo)
	if err != nil {
		return err
	}
	return nil
}

func deviceScanCallback(data []byte) error {
	return nil
}

func apScanCallback(data []byte) error {
	return nil
}

