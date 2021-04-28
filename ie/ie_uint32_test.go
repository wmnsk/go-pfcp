// Copyright 2019-2021 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie_test

import (
	"net"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/wmnsk/go-pfcp/ie"
)

var (
	mac1, _ = net.ParseMAC("12:34:56:78:90:01")
	mac2, _ = net.ParseMAC("12:34:56:78:90:02")
	mac3, _ = net.ParseMAC("12:34:56:78:90:03")
	mac4, _ = net.ParseMAC("12:34:56:78:90:04")
)

func TestUint32IEs(t *testing.T) {
	cases := []struct {
		description string
		structured  *ie.IE
		decoded     uint32
		decoderFunc func(*ie.IE) (uint32, error)
	}{
		{
			description: "AggregatedURRID",
			structured:  ie.NewAggregatedURRID(0xffffffff),
			decoded:     0xffffffff,
			decoderFunc: func(i *ie.IE) (uint32, error) { return i.AggregatedURRID() },
		}, {
			description: "AggregatedURRID/AggregatedURRs",
			structured: ie.NewAggregatedURRs(
				ie.NewAggregatedURRID(0xffffffff),
				ie.NewMultiplier(0xffffffffffffffff, 0x11223344),
			),
			decoded:     0xffffffff,
			decoderFunc: func(i *ie.IE) (uint32, error) { return i.AggregatedURRID() },
		}, {
			description: "AveragingWindow",
			structured:  ie.NewAveragingWindow(0xffffffff),
			decoded:     0xffffffff,
			decoderFunc: func(i *ie.IE) (uint32, error) { return i.AveragingWindow() },
		}, {
			description: "AveragingWindow/CreateQER",
			structured: ie.NewCreateQER(
				ie.NewQERID(0xffffffff),
				ie.NewQERCorrelationID(0x11111111),
				ie.NewGateStatus(ie.GateStatusOpen, ie.GateStatusOpen),
				ie.NewMBR(0x11111111, 0x22222222),
				ie.NewGBR(0x11111111, 0x22222222),
				ie.NewPacketRate(0x03, ie.TimeUnitMinute, 0x1122, ie.TimeUnitMinute, 0x3344),
				ie.NewPacketRateStatus(0x07, 0x1111, 0x2222, 0x3333, 0x4444, time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
				ie.NewDLFlowLevelMarking(0x03, 0x1122, 0x3344),
				ie.NewQFI(0x01),
				ie.NewRQI(0x01),
				ie.NewPagingPolicyIndicator(1),
				ie.NewAveragingWindow(0xffffffff),
				ie.NewQERControlIndications(1, 1, 1),
			),
			decoded:     0xffffffff,
			decoderFunc: func(i *ie.IE) (uint32, error) { return i.AveragingWindow() },
		}, {
			description: "CumulativeRateRatioMeasurement",
			structured:  ie.NewCumulativeRateRatioMeasurement(0xffffffff),
			decoded:     0xffffffff,
			decoderFunc: func(i *ie.IE) (uint32, error) { return i.CumulativeRateRatioMeasurement() },
		}, {
			description: "CumulativeRateRatioThreshold",
			structured:  ie.NewCumulativeRateRatioThreshold(0xffffffff),
			decoded:     0xffffffff,
			decoderFunc: func(i *ie.IE) (uint32, error) { return i.CumulativeRateRatioThreshold() },
		}, {
			description: "CumulativeRateRatioThreshold/ClockDriftControlInformation",
			structured: ie.NewClockDriftControlInformation(
				ie.NewRequestedClockDriftInformation(1, 1),
				ie.NewTSNTimeDomainNumber(255),
				ie.NewTimeOffsetThreshold(10*time.Second),
				ie.NewCumulativeRateRatioThreshold(0xffffffff),
			),
			decoded:     0xffffffff,
			decoderFunc: func(i *ie.IE) (uint32, error) { return i.CumulativeRateRatioThreshold() },
		}, {
			description: "CumulativeRateRatioThreshold/ClockDriftReport",
			structured: ie.NewClockDriftReport(
				ie.NewTSNTimeDomainNumber(255),
				ie.NewTimeOffsetThreshold(10*time.Second),
				ie.NewCumulativeRateRatioThreshold(0xffffffff),
				ie.NewEventTimeStamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     0xffffffff,
			decoderFunc: func(i *ie.IE) (uint32, error) { return i.CumulativeRateRatioThreshold() },
		}, {
			description: "DSTTPortNumber",
			structured:  ie.NewDSTTPortNumber(0xffffffff),
			decoded:     0xffffffff,
			decoderFunc: func(i *ie.IE) (uint32, error) { return i.DSTTPortNumber() },
		}, {
			description: "DSTTPortNumber/CreatedBridgeInfoForTSC",
			structured: ie.NewCreatedBridgeInfoForTSC(
				ie.NewDSTTPortNumber(0xffffffff),
				ie.NewNWTTPortNumber(0xffffffff),
				ie.NewTSNBridgeID(mac1),
			),
			decoded:     0xffffffff,
			decoderFunc: func(i *ie.IE) (uint32, error) { return i.DSTTPortNumber() },
		}, {
			description: "EthernetFilterID",
			structured:  ie.NewEthernetFilterID(0xffffffff),
			decoded:     0xffffffff,
			decoderFunc: func(i *ie.IE) (uint32, error) { return i.EthernetFilterID() },
		}, {
			description: "EthernetFilterID/PDI",
			structured: ie.NewPDI(
				ie.NewSourceInterface(ie.SrcInterfaceAccess),
				ie.NewFTEID(0x01, 0x11111111, net.ParseIP("127.0.0.1"), nil, 0),
				ie.NewNetworkInstance("some.instance.example"),
				ie.NewRedundantTransmissionParametersInPDI(
					ie.NewFTEID(0x01, 0x11111111, net.ParseIP("127.0.0.1"), nil, 0),
					ie.NewNetworkInstance("some.instance.example"),
				),
				ie.NewUEIPAddress(0x02, "127.0.0.1", "", 0, 0),
				ie.NewTrafficEndpointID(0x01),
				ie.NewSDFFilter("aaaaaaaa", "bb", "cccc", "ddd", 0xffffffff),
				ie.NewApplicationID("https://github.com/wmnsk/go-pfcp/"),
				ie.NewEthernetPDUSessionInformation(0x01),
				ie.NewEthernetPacketFilter(
					ie.NewEthernetFilterID(0xffffffff),
					ie.NewEthernetFilterProperties(0x01),
					ie.NewMACAddress(mac1, mac2, mac3, mac4),
					ie.NewEthertype(0xffff),
					ie.NewCTAG(0x07, 1, 1, 4095),
					ie.NewSTAG(0x07, 1, 1, 4095),
					ie.NewSDFFilter("aaaaaaaa", "bb", "cccc", "ddd", 0xffffffff),
				),
			),
			decoded:     0xffffffff,
			decoderFunc: func(i *ie.IE) (uint32, error) { return i.EthernetFilterID() },
		}, {
			description: "EthernetFilterID/EthernetPacketFilter",
			structured: ie.NewEthernetPacketFilter(
				ie.NewEthernetFilterID(0xffffffff),
				ie.NewEthernetFilterProperties(0x01),
				ie.NewMACAddress(mac1, mac2, mac3, mac4),
				ie.NewEthertype(0xffff),
				ie.NewCTAG(0x07, 1, 1, 4095),
				ie.NewSTAG(0x07, 1, 1, 4095),
				ie.NewSDFFilter("aaaaaaaa", "bb", "cccc", "ddd", 0xffffffff),
			),
			decoded:     0xffffffff,
			decoderFunc: func(i *ie.IE) (uint32, error) { return i.EthernetFilterID() },
		}, {
			description: "EventQuota",
			structured:  ie.NewEventQuota(0xffffffff),
			decoded:     0xffffffff,
			decoderFunc: func(i *ie.IE) (uint32, error) { return i.EventQuota() },
		}, {
			description: "EventQuota/CreateURR",
			structured: ie.NewCreateURR(
				ie.NewURRID(0xffffffff),
				ie.NewMeasurementMethod(1, 1, 1),
				ie.NewReportingTriggers(0x1122),
				ie.NewMeasurementPeriod(10*time.Second),
				ie.NewVolumeThreshold(0x07, 0x3333333333333333, 0x1111111111111111, 0x2222222222222222),
				ie.NewVolumeQuota(0x07, 0xffffffffffffffff, 0xffffffffffffffff, 0xffffffffffffffff),
				ie.NewEventThreshold(0xffffffff),
				ie.NewEventQuota(0xffffffff),
				ie.NewTimeThreshold(0x11111111),
				ie.NewTimeQuota(10*time.Second),
				ie.NewQuotaHoldingTime(10*time.Second),
				ie.NewDroppedDLTrafficThreshold(true, true, 0xffffffffffffffff, 0xffffffffffffffff),
				ie.NewQuotaValidityTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
				ie.NewMonitoringTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
				ie.NewSubsequentVolumeThreshold(0x07, 0x3333333333333333, 0x1111111111111111, 0x2222222222222222),
				ie.NewSubsequentTimeThreshold(0x11111111),
				ie.NewSubsequentVolumeQuota(0x07, 0x3333333333333333, 0x1111111111111111, 0x2222222222222222),
				ie.NewSubsequentTimeQuota(10*time.Second),
				ie.NewSubsequentEventThreshold(0xffffffff),
				ie.NewSubsequentEventQuota(0xffffffff),
				ie.NewInactivityDetectionTime(0x11111111),
				ie.NewLinkedURRID(0xffffffff),
				ie.NewMeasurementInformation(0x1f),
				ie.NewTimeQuotaMechanism(ie.BTITCTP, 10*time.Second),
				ie.NewAggregatedURRs(
					ie.NewAggregatedURRID(0xffffffff),
					ie.NewMultiplier(0xffffffffffffffff, 0x11223344),
				),
				ie.NewFARID(0xffffffff),
				ie.NewEthernetInactivityTimer(10*time.Second),
				ie.NewAdditionalMonitoringTime(
					ie.NewMonitoringTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewSubsequentVolumeThreshold(0x01, 0x1111111111111111, 0, 0),
					ie.NewSubsequentTimeThreshold(0x11111111),
					ie.NewSubsequentVolumeQuota(0x01, 0x1111111111111111, 0, 0),
					ie.NewSubsequentTimeQuota(10*time.Second),
					ie.NewEventThreshold(0xffffffff),
					ie.NewEventQuota(0xffffffff),
				),
				ie.NewNumberOfReports(0x1111),
			),
			decoded:     0xffffffff,
			decoderFunc: func(i *ie.IE) (uint32, error) { return i.EventQuota() },
		}, {
			description: "EventQuota/UpdateURR",
			structured: ie.NewUpdateURR(
				ie.NewURRID(0xffffffff),
				ie.NewMeasurementMethod(1, 1, 1),
				ie.NewReportingTriggers(0x1122),
				ie.NewMeasurementPeriod(10*time.Second),
				ie.NewVolumeThreshold(0x07, 0x3333333333333333, 0x1111111111111111, 0x2222222222222222),
				ie.NewVolumeQuota(0x07, 0xffffffffffffffff, 0xffffffffffffffff, 0xffffffffffffffff),
				ie.NewEventThreshold(0xffffffff),
				ie.NewEventQuota(0xffffffff),
				ie.NewTimeThreshold(0x11111111),
				ie.NewTimeQuota(10*time.Second),
				ie.NewQuotaHoldingTime(10*time.Second),
				ie.NewDroppedDLTrafficThreshold(true, true, 0xffffffffffffffff, 0xffffffffffffffff),
				ie.NewQuotaValidityTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
				ie.NewMonitoringTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
				ie.NewSubsequentVolumeThreshold(0x07, 0x3333333333333333, 0x1111111111111111, 0x2222222222222222),
				ie.NewSubsequentTimeThreshold(0x11111111),
				ie.NewSubsequentVolumeQuota(0x07, 0x3333333333333333, 0x1111111111111111, 0x2222222222222222),
				ie.NewSubsequentTimeQuota(10*time.Second),
				ie.NewSubsequentEventThreshold(0xffffffff),
				ie.NewSubsequentEventQuota(0xffffffff),
				ie.NewInactivityDetectionTime(0x11111111),
				ie.NewLinkedURRID(0xffffffff),
				ie.NewMeasurementInformation(0x1f),
				ie.NewTimeQuotaMechanism(ie.BTITCTP, 10*time.Second),
				ie.NewAggregatedURRs(
					ie.NewAggregatedURRID(0xffffffff),
					ie.NewMultiplier(0xffffffffffffffff, 0x11223344),
				),
				ie.NewFARID(0xffffffff),
				ie.NewEthernetInactivityTimer(10*time.Second),
				ie.NewAdditionalMonitoringTime(
					ie.NewMonitoringTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewSubsequentVolumeThreshold(0x01, 0x1111111111111111, 0, 0),
					ie.NewSubsequentTimeThreshold(0x11111111),
					ie.NewSubsequentVolumeQuota(0x01, 0x1111111111111111, 0, 0),
					ie.NewSubsequentTimeQuota(10*time.Second),
					ie.NewEventThreshold(0xffffffff),
					ie.NewEventQuota(0xffffffff),
				),
				ie.NewNumberOfReports(0x1111),
			),
			decoded:     0xffffffff,
			decoderFunc: func(i *ie.IE) (uint32, error) { return i.EventQuota() },
		}, {
			description: "EventQuota/AdditionalMonitoringTime",
			structured: ie.NewAdditionalMonitoringTime(
				ie.NewMonitoringTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
				ie.NewSubsequentVolumeThreshold(0x01, 0x1111111111111111, 0, 0),
				ie.NewSubsequentTimeThreshold(0x11111111),
				ie.NewSubsequentVolumeQuota(0x01, 0x1111111111111111, 0, 0),
				ie.NewSubsequentTimeQuota(10*time.Second),
				ie.NewEventThreshold(0xffffffff),
				ie.NewEventQuota(0xffffffff),
			),
			decoded:     0xffffffff,
			decoderFunc: func(i *ie.IE) (uint32, error) { return i.EventQuota() },
		}, {
			description: "EventThreshold",
			structured:  ie.NewEventThreshold(0xffffffff),
			decoded:     0xffffffff,
			decoderFunc: func(i *ie.IE) (uint32, error) { return i.EventThreshold() },
		}, {
			description: "EventThreshold/CreateURR",
			structured: ie.NewCreateURR(
				ie.NewURRID(0xffffffff),
				ie.NewMeasurementMethod(1, 1, 1),
				ie.NewReportingTriggers(0x1122),
				ie.NewMeasurementPeriod(10*time.Second),
				ie.NewVolumeThreshold(0x07, 0x3333333333333333, 0x1111111111111111, 0x2222222222222222),
				ie.NewVolumeQuota(0x07, 0xffffffffffffffff, 0xffffffffffffffff, 0xffffffffffffffff),
				ie.NewEventThreshold(0xffffffff),
				ie.NewEventQuota(0xffffffff),
				ie.NewTimeThreshold(0x11111111),
				ie.NewTimeQuota(10*time.Second),
				ie.NewQuotaHoldingTime(10*time.Second),
				ie.NewDroppedDLTrafficThreshold(true, true, 0xffffffffffffffff, 0xffffffffffffffff),
				ie.NewQuotaValidityTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
				ie.NewMonitoringTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
				ie.NewSubsequentVolumeThreshold(0x07, 0x3333333333333333, 0x1111111111111111, 0x2222222222222222),
				ie.NewSubsequentTimeThreshold(0x11111111),
				ie.NewSubsequentVolumeQuota(0x07, 0x3333333333333333, 0x1111111111111111, 0x2222222222222222),
				ie.NewSubsequentTimeQuota(10*time.Second),
				ie.NewSubsequentEventThreshold(0xffffffff),
				ie.NewSubsequentEventQuota(0xffffffff),
				ie.NewInactivityDetectionTime(0x11111111),
				ie.NewLinkedURRID(0xffffffff),
				ie.NewMeasurementInformation(0x1f),
				ie.NewTimeQuotaMechanism(ie.BTITCTP, 10*time.Second),
				ie.NewAggregatedURRs(
					ie.NewAggregatedURRID(0xffffffff),
					ie.NewMultiplier(0xffffffffffffffff, 0x11223344),
				),
				ie.NewFARID(0xffffffff),
				ie.NewEthernetInactivityTimer(10*time.Second),
				ie.NewAdditionalMonitoringTime(
					ie.NewMonitoringTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewSubsequentVolumeThreshold(0x01, 0x1111111111111111, 0, 0),
					ie.NewSubsequentTimeThreshold(0x11111111),
					ie.NewSubsequentVolumeQuota(0x01, 0x1111111111111111, 0, 0),
					ie.NewSubsequentTimeQuota(10*time.Second),
					ie.NewEventThreshold(0xffffffff),
					ie.NewEventQuota(0xffffffff),
				),
				ie.NewNumberOfReports(0x1111),
			),
			decoded:     0xffffffff,
			decoderFunc: func(i *ie.IE) (uint32, error) { return i.EventThreshold() },
		}, {
			description: "EventThreshold/UpdateURR",
			structured: ie.NewUpdateURR(
				ie.NewURRID(0xffffffff),
				ie.NewMeasurementMethod(1, 1, 1),
				ie.NewReportingTriggers(0x1122),
				ie.NewMeasurementPeriod(10*time.Second),
				ie.NewVolumeThreshold(0x07, 0x3333333333333333, 0x1111111111111111, 0x2222222222222222),
				ie.NewVolumeQuota(0x07, 0xffffffffffffffff, 0xffffffffffffffff, 0xffffffffffffffff),
				ie.NewEventThreshold(0xffffffff),
				ie.NewEventQuota(0xffffffff),
				ie.NewTimeThreshold(0x11111111),
				ie.NewTimeQuota(10*time.Second),
				ie.NewQuotaHoldingTime(10*time.Second),
				ie.NewDroppedDLTrafficThreshold(true, true, 0xffffffffffffffff, 0xffffffffffffffff),
				ie.NewQuotaValidityTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
				ie.NewMonitoringTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
				ie.NewSubsequentVolumeThreshold(0x07, 0x3333333333333333, 0x1111111111111111, 0x2222222222222222),
				ie.NewSubsequentTimeThreshold(0x11111111),
				ie.NewSubsequentVolumeQuota(0x07, 0x3333333333333333, 0x1111111111111111, 0x2222222222222222),
				ie.NewSubsequentTimeQuota(10*time.Second),
				ie.NewSubsequentEventThreshold(0xffffffff),
				ie.NewSubsequentEventQuota(0xffffffff),
				ie.NewInactivityDetectionTime(0x11111111),
				ie.NewLinkedURRID(0xffffffff),
				ie.NewMeasurementInformation(0x1f),
				ie.NewTimeQuotaMechanism(ie.BTITCTP, 10*time.Second),
				ie.NewAggregatedURRs(
					ie.NewAggregatedURRID(0xffffffff),
					ie.NewMultiplier(0xffffffffffffffff, 0x11223344),
				),
				ie.NewFARID(0xffffffff),
				ie.NewEthernetInactivityTimer(10*time.Second),
				ie.NewAdditionalMonitoringTime(
					ie.NewMonitoringTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
					ie.NewSubsequentVolumeThreshold(0x01, 0x1111111111111111, 0, 0),
					ie.NewSubsequentTimeThreshold(0x11111111),
					ie.NewSubsequentVolumeQuota(0x01, 0x1111111111111111, 0, 0),
					ie.NewSubsequentTimeQuota(10*time.Second),
					ie.NewEventThreshold(0xffffffff),
					ie.NewEventQuota(0xffffffff),
				),
				ie.NewNumberOfReports(0x1111),
			),
			decoded:     0xffffffff,
			decoderFunc: func(i *ie.IE) (uint32, error) { return i.EventThreshold() },
		}, {
			description: "EventThreshold/AdditionalMonitoringTime",
			structured: ie.NewAdditionalMonitoringTime(
				ie.NewMonitoringTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
				ie.NewSubsequentVolumeThreshold(0x01, 0x1111111111111111, 0, 0),
				ie.NewSubsequentTimeThreshold(0x11111111),
				ie.NewSubsequentVolumeQuota(0x01, 0x1111111111111111, 0, 0),
				ie.NewSubsequentTimeQuota(10*time.Second),
				ie.NewEventThreshold(0xffffffff),
				ie.NewEventQuota(0xffffffff),
			),
			decoded:     0xffffffff,
			decoderFunc: func(i *ie.IE) (uint32, error) { return i.EventThreshold() },
		}, {
			description: "Multiplier/Exponent",
			structured:  ie.NewMultiplier(0xffffffffffffffff, 0x11223344),
			decoded:     0x11223344,
			decoderFunc: func(i *ie.IE) (uint32, error) { return i.Exponent() },
		}, {
			description: "SNSSAI/SD",
			structured:  ie.NewSNSSAI(0x11, 0x223344),
			decoded:     0x223344,
			decoderFunc: func(i *ie.IE) (uint32, error) { return i.SD() },
		}, {
			description: "SNSSAI/UEIPAddressPoolInformation/SD",
			structured: ie.NewUEIPAddressPoolInformation(
				ie.NewUEIPAddressPoolIdentity("go-pfcp"),
				ie.NewNetworkInstance("some.instance.example"),
				ie.NewSNSSAI(0x11, 0x223344),
				ie.NewIPVersion(true, false),
			),
			decoded:     0x223344,
			decoderFunc: func(i *ie.IE) (uint32, error) { return i.SD() },
		},
	}

	for _, c := range cases {
		t.Run(c.description, func(t *testing.T) {
			got, err := c.decoderFunc(c.structured)
			if err != nil {
				t.Fatal(err)
			}

			if diff := cmp.Diff(got, c.decoded); diff != "" {
				t.Error(diff)
			}
		})
	}
}
