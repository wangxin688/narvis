package device360_utils

import (
	"github.com/wangxin688/narvis/server/features/device360/schemas"
	"github.com/wangxin688/narvis/server/pkg/vtm"
)

func VectorResponseToHealthMap(response []*vtm.VectorResponse) *schemas.HealthHeatMap {
	if len(response) == 0 {
		return nil
	}
	healthMap := &schemas.HealthHeatMap{
		Good:           0,
		NeedsAttention: 0,
		Poor:           0,
		Unhealthy:      0,
		Unknown:        0,
	}
	for _, vector := range response {
		score := ScoreToHealth(vector.Value[1].(string))
		healthMap.Add(score)
	}
	return healthMap
}
