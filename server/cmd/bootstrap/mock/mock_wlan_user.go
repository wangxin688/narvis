package mock

import (
	"time"

	"github.com/samber/lo"
	"github.com/wangxin688/narvis/server/dal/gen"
	"github.com/wangxin688/narvis/server/models"
	"github.com/wangxin688/narvis/server/tests/fixtures"
)

func mockWlanUser(orgId string, siteId string) {
	newWlanUsers := make([]*models.WlanStation, 0)
	stationMacs := []string{
		"00:00:00:00:00:00",
		"00:00:00:00:00:01",
		"00:00:00:00:00:02",
		"00:00:00:00:00:03",
		"00:00:00:00:00:04",
		"00:00:00:00:00:05",
		"00:00:00:00:00:06",
		"00:00:00:00:00:07",
		"00:00:00:00:00:08",
		"00:00:00:00:00:09",
		"01:00:00:00:00:00",
		"01:00:00:00:00:01",
		"01:00:00:00:00:02",
		"01:00:00:00:00:03",
		"01:00:00:00:00:04",
		"01:00:00:00:00:05",
		"01:00:00:00:00:06",
		"01:00:00:00:00:07",
		"01:00:00:00:00:08",
		"01:00:00:00:00:09",
	}
	stationIps := []string{
		"192.168.100.1",
		"192.168.100.2",
		"192.168.100.3",
		"192.168.100.4",
		"192.168.100.5",
		"192.168.100.6",
		"192.168.100.7",
		"192.168.100.8",
		"192.168.100.9",
		"192.168.101.10",
		"192.168.101.1",
		"192.168.101.2",
		"192.168.101.3",
		"192.168.101.4",
		"192.168.101.5",
		"192.168.101.6",
		"192.168.101.7",
		"192.168.101.8",
		"192.168.101.9",
		"192.168.102.10",
	}
	stationUsernames := []string{
		"user0@example.com",
		"user1@example.com",
		"user2@example.com",
		"user3@example.com",
		"user4@example.com",
		"user5@example.com",
		"user6@example.com",
		"user7@example.com",
		"user8@example.com",
		"user9@example.com",
		"0user@example.com",
		"1user@example.com",
		"2user@example.com",
		"3user@example.com",
		"4user@example.com",
		"5user@example.com",
		"6user@example.com",
		"7user@example.com",
		"8user@example.com",
		"9user@example.com",
	}
	aps, err := fixtures.GetSiteApNames(siteId)
	if err != nil {
		panic(err)
	}
	essids := []string{"mockSSID1", "mockSSID2", "mockSSID3", "mockSSID4", "mockSSID5"}
	vlanIds := []uint16{801, 802, 803, 804, 805}
	stationRadio := []struct {
		Channel   uint16
		RadioType string
		BandWidth string
	}{
		{1, "2.4GHz", "20MHz"},
		{6, "2.4GHz", "20MHz"},
		{11, "2.4GHz", "20MHz"},
		{36, "5GHz", "20MHz"},
		{40, "5GHz", "20MHz"},
		{44, "5GHz", "40MHz"},
		{48, "5GHz", "20MHz"},
		{52, "5GHz", "80MHz"},
		{56, "5GHz", "20MHz"},
		{149, "5GHz", "40MHz"},
		{157, "5GHz", "40MHz"},
		{165, "5GHz", "20MHz"},
	}
	snr := []uint8{33, 34, 28, 41, 35, 36, 40}
	rssi := []int8{-65, -66, -67, -68, -69, -70, -71}
	maxSpeed := []uint64{100, 200, 300, 400, 500, 600, 700}
	onlineTime := []uint64{10000, 20000, 30000, 40000, 50000, 6000000, 700000}

	for i := 0; i < 20; i++ {
		vlan := lo.Sample(vlanIds)
		snr := lo.Sample(snr)
		channel := lo.Sample(stationRadio)
		speed := lo.Sample(maxSpeed)
		online := lo.Sample(onlineTime)
		newWlanUsers = append(newWlanUsers, &models.WlanStation{
			StationMac:           stationMacs[i],
			StationIp:            stationIps[i],
			StationUsername:      stationUsernames[i],
			StationApName:        &aps[i],
			StationESSID:         lo.Sample(essids),
			StationVlan:          &vlan,
			StationChannel:       channel.Channel,
			StationChanBandWidth: &channel.BandWidth,
			StationRadioType:     channel.RadioType,
			StationSNR:           &snr,
			StationRSSI:          lo.Sample(rssi),
			StationRxBits:        100000 + uint64(time.Now().Unix()),
			StationTxBits:        200000 + uint64(time.Now().Unix()),
			StationMaxSpeed:      &speed,
			StationOnlineTime:    online,
			SiteId:               siteId,
			OrganizationId:       orgId,
		})
	}
	err = gen.WlanStation.CreateInBatches(newWlanUsers, 200)
	if err != nil {
		panic(err)
	}
}
