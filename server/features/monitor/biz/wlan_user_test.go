package metric_biz_test

import (
	"testing"
	"time"

	metric_biz "github.com/wangxin688/narvis/server/features/monitor/biz"
	"github.com/wangxin688/narvis/server/features/monitor/schemas"
	"github.com/wangxin688/narvis/server/tests/fixtures"
)

func TestGetWlanUserTrend(t *testing.T) {
	fixtures.FixturePrepare()

	ws := metric_biz.NewWlanUserService()
	result, err := ws.GetWlanUserTrend(&schemas.WlanUserTrendRequest{
		StartedAt: time.Now().Add(time.Hour * 30 * -24),
		EndedAt:   time.Now(),
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log(result)

}
