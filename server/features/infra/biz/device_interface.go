package infra_biz

import (
	"fmt"

	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/tools/helpers"
)

type InterfaceService struct {
}

func NewDeviceInterfaceService() *InterfaceService {
	return &InterfaceService{}
}

// Get device interfaces with hashValue calculation, return map[HashValue]DeviceInterface
func (d *InterfaceService) GetDeviceInterfaces(deviceId string) (map[string]*models.DeviceInterface, error) {

	di, err := gen.DeviceInterface.Select(gen.DeviceInterface.Id, gen.DeviceInterface.IfName, gen.DeviceInterface.IfIndex).Where(gen.DeviceInterface.DeviceId.Eq(deviceId)).Find()
	if err != nil {
		return nil, err
	}
	res := make(map[string]*models.DeviceInterface)
	for _, item := range di {
		hashValue := d.calHashValue(item)
		res[hashValue] = item
	}
	return res, nil
}

func (d *InterfaceService) calHashValue(item *models.DeviceInterface) string {

	hashString := fmt.Sprintf(
		"%s-%s-%s-%d-%d-%d-%s-%s-%s-%d-%s",
		item.IfName,
		item.IfDescr,
		item.IfType,
		item.IfHighSpeed,
		item.IfMtu,
		item.IfSpeed,
		helpers.PtrStringToString(item.IfPhysAddr),
		item.IfAdminStatus,
		item.IfOperStatus,
		item.IfLastChange,
		helpers.PtrStringToString(item.IfIpAddress))
	return helpers.StringToMd5(hashString)
}
