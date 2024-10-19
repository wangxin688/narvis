package infra_biz

import (
	"bufio"
	"strings"

	"github.com/andreyvit/diff"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/models"
)

type DeviceConfigService struct {
}

func NewDeviceConfigService() *DeviceConfigService {
	return &DeviceConfigService{}
}

func (s *DeviceConfigService) GetLatestDeviceConfigByDeviceId(deviceId string) (*models.DeviceConfig, error) {

	deviceConfig, err := gen.DeviceConfig.Where(
		gen.DeviceConfig.DeviceId.Eq(deviceId)).Order(gen.DeviceConfig.Id.Desc()).Limit(1).Find()

	if err != nil {
		return nil, err
	}
	if len(deviceConfig) == 0 {
		return nil, nil
	}
	return deviceConfig[0], nil
}

func GetConfigTotalLines(config string) int {
	scanner := bufio.NewScanner(strings.NewReader(config))
	lineCount := 0
	for scanner.Scan() {
		lineCount++
	}
	return lineCount
}

func GetLineChanges(config1 string, config2 string) (added, deleted int) {
	diffValue := diff.LineDiffAsLines(config1, config2)
	for _, d := range diffValue {
		if strings.HasPrefix(d, "+") {
			added++
		} else if strings.HasPrefix(d, "-") {
			deleted++
		}
	}
	return
}
