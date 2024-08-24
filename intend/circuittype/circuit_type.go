package circuittype

import (
	"github.com/samber/lo"
	"github.com/wangxin688/narvis/intend/common"
)

type CircuitTypeEnum string
type ConnectionTypeEum string

const (
	P2P       CircuitTypeEnum = "P2P"
	Internet  CircuitTypeEnum = "Internet"
	MPLS      CircuitTypeEnum = "MPLS"
	DarkFiber CircuitTypeEnum = "DarkFiber"
	ADSL      CircuitTypeEnum = "ADSL"
	Unknown   CircuitTypeEnum = "Unknown"
)

const (
	WAN ConnectionTypeEum = "WAN"
	LAN ConnectionTypeEum = "LAN"
)

type CircuitType struct {
	CircuitType    CircuitTypeEnum
	Description    common.I18n
	ConnectionType ConnectionTypeEum
}

func getCircuitTypeMeta() map[CircuitTypeEnum]CircuitType {
	circuitTypeMeta := map[CircuitTypeEnum]CircuitType{
		P2P:       {CircuitType: P2P, Description: common.I18n{En: "Point-to-point", Zh: "点对点专线"}, ConnectionType: WAN},
		Internet:  {CircuitType: Internet, Description: common.I18n{En: "Internet", Zh: "互联网专线"}, ConnectionType: WAN},
		MPLS:      {CircuitType: MPLS, Description: common.I18n{En: "MPLS", Zh: "MPLS专线"}, ConnectionType: WAN},
		DarkFiber: {CircuitType: DarkFiber, Description: common.I18n{En: "DarkFiber", Zh: "裸光专线"}, ConnectionType: WAN},
		ADSL:      {CircuitType: ADSL, Description: common.I18n{En: "ADSL", Zh: "ADSL专线"}, ConnectionType: WAN},
		Unknown:   {CircuitType: Unknown, Description: common.I18n{En: "Unknown", Zh: "其他专线"}, ConnectionType: WAN},
	}
	return circuitTypeMeta
}

// Get List of CircuitType
func GetListCircuitType() []CircuitType {
	return lo.Values(getCircuitTypeMeta())
}

// Get CircuitType by name
func GetCircuitType(circuitType string) CircuitType {
	circuitTypeMeta := getCircuitTypeMeta()
	circuitTypeMetaValue, ok := circuitTypeMeta[CircuitTypeEnum(circuitType)]
	if !ok {
		return circuitTypeMeta[CircuitTypeEnum(Unknown)]
	}
	return circuitTypeMetaValue
}
