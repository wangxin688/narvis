package mock

import (
	"fmt"
	"time"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/tests/fixtures"
)

func mockDeviceAlerts(orgId, siteId, deviceId string) {
	createAlerts := make([]*models.Alert, 0)
	for i := 0; i < 20; i++ {
		status := lo.Sample([]uint8{0, 1})
		startedAt := fixtures.MockTimeBeforeNow()
		var resolvedAt *time.Time
		if status == 1 {
			now := time.Now()
			resolvedAt = &now
		}
		acked := lo.Sample([]bool{true, false})
		alertName := lo.Sample([]string{
			"high_cpu_utilization",
			"high_memory_utilization",
			"high_disk_utilization",
			"high_system_load",
			"unavailable_by_icmp_ping",
			"fan_status_abnormal",
		})
		eventId := fmt.Sprintf("mock_%s_%d", deviceId, i)
		triggerId := fmt.Sprintf("mock_%s_%d", deviceId, i)
		createAlerts = append(createAlerts, &models.Alert{
			Status:         status,
			StartedAt:      startedAt,
			ResolvedAt:     resolvedAt,
			Acknowledged:   acked,
			AlertName:      alertName,
			EventId:        eventId,
			TriggerId:      triggerId,
			SiteId:         siteId,
			DeviceId:       &deviceId,
			OrganizationId: orgId,
		})

	}
	err := gen.Alert.CreateInBatches(createAlerts, 500)
	if err != nil {
		panic(err)
	}

}

func mockApAlerts(orgId, siteId, apId string) {
	createAlerts := make([]*models.Alert, 0)
	for i := 0; i < 20; i++ {
		status := lo.Sample([]uint8{0, 1})
		startedAt := fixtures.MockTimeBeforeNow()
		var resolvedAt *time.Time
		if status == 1 {
			now := time.Now()
			resolvedAt = &now
		}
		acked := lo.Sample([]bool{true, false})
		alertName := lo.Sample([]string{
			"high_channel_utilization",
			"high_channel_interference",
			"high_channel_noise",
			"high_client_number",
			"wireless_access_point_down",
		})
		eventId := fmt.Sprintf("mock_%s_%d", apId, i)
		triggerId := fmt.Sprintf("mock_%s_%d", apId, i)
		createAlerts = append(createAlerts, &models.Alert{
			Status:         status,
			StartedAt:      startedAt,
			ResolvedAt:     resolvedAt,
			Acknowledged:   acked,
			AlertName:      alertName,
			EventId:        eventId,
			TriggerId:      triggerId,
			SiteId:         siteId,
			ApId:           &apId,
			OrganizationId: orgId,
		})

	}
	err := gen.Alert.CreateInBatches(createAlerts, 500)
	if err != nil {
		panic(err)
	}
}

func mockCircuitAlerts(orgId, siteId, circuitId string) {
	createAlerts := make([]*models.Alert, 0)
	for i := 0; i < 20; i++ {
		status := lo.Sample([]uint8{0, 1})
		startedAt := fixtures.MockTimeBeforeNow()
		var resolvedAt *time.Time
		if status == 1 {
			now := time.Now()
			resolvedAt = &now
		}
		acked := lo.Sample([]bool{true, false})
		alertName := lo.Sample([]string{
			"high_icmp_ping_response_time",
			"high_icmp_ping_loss",
			"high_bandwidth_usage",
			"high_error_rate",
		})
		eventId := fmt.Sprintf("mock_%s_%d", circuitId, i)
		triggerId := fmt.Sprintf("mock_%s_%d", circuitId, i)
		createAlerts = append(createAlerts, &models.Alert{
			Status:         status,
			StartedAt:      startedAt,
			ResolvedAt:     resolvedAt,
			Acknowledged:   acked,
			AlertName:      alertName,
			EventId:        eventId,
			TriggerId:      triggerId,
			SiteId:         siteId,
			CircuitId:      &circuitId,
			OrganizationId: orgId,
		})

	}
	err := gen.Alert.CreateInBatches(createAlerts, 500)
	if err != nil {
		panic(err)
	}
}

func mockServerAlerts(orgId, siteId, serverId string) {
	createAlerts := make([]*models.Alert, 0)
	for i := 0; i < 100; i++ {
		status := lo.Sample([]uint8{0, 1})
		startedAt := fixtures.MockTimeBeforeNow()
		var resolvedAt *time.Time
		if status == 1 {
			now := time.Now()
			resolvedAt = &now
		}
		acked := lo.Sample([]bool{true, false})
		alertName := lo.Sample([]string{
			"high_icmp_ping_response_time",
			"high_icmp_ping_loss",
			"high_inode_utilization",
			"high_system_load_average",
			"high_swap_space_utilization",
		})
		eventId := fmt.Sprintf("mock_%s_%d", serverId, i)
		triggerId := fmt.Sprintf("mock_%s_%d", serverId, i)
		createAlerts = append(createAlerts, &models.Alert{
			Status:         status,
			StartedAt:      startedAt,
			ResolvedAt:     resolvedAt,
			Acknowledged:   acked,
			AlertName:      alertName,
			EventId:        eventId,
			TriggerId:      triggerId,
			SiteId:         siteId,
			ServerId:       &serverId,
			OrganizationId: orgId,
		})

	}
	err := gen.Alert.CreateInBatches(createAlerts, 500)
	if err != nil {
		panic(err)
	}
}
