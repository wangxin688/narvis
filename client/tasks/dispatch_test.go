package tasks

import (
	"encoding/json"
	"testing"

	"github.com/wangxin688/narvis/client/config"
	mem_cache "github.com/wangxin688/narvis/client/utils/cache"
	"github.com/wangxin688/narvis/intend/intendtask"
)

func TestDispatcher(t *testing.T) {
	config.SetupConfig()
	mem_cache.InitCache()

	scanTask := map[string]any{
		"taskId":   "e53db778-bb10-411c-ba47-d6e422060f29",
		"taskName": intendtask.ScanDeviceBasicInfo,
		"callback": intendtask.DeviceBasicInfoCbUrl,
		"snmpConfig": map[string]any{
			"community":      "public",
			"port":           161,
			"timeout":        10,
			"maxRepetitions": 50,
		},
		"range": "10.90.200.129/28",
	}
	baseTask, _ := json.Marshal(scanTask)
	TaskDispatcher(baseTask)
}
