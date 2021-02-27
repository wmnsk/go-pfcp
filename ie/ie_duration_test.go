// Copyright 2019-2021 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie_test

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/wmnsk/go-pfcp/ie"
)

func TestDurationIEs(t *testing.T) {
	cases := []struct {
		description string
		structured  *ie.IE
		decoded     time.Duration
		decoderFunc func(*ie.IE) (time.Duration, error)
	}{
		{
			description: "AveragePacketDelay",
			structured:  ie.NewAveragePacketDelay(10 * time.Second),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.AveragePacketDelay() },
		}, {
			description: "AveragePacketDelay/GTPUPathQoSControlInformation",
			structured: ie.NewGTPUPathQoSControlInformation(
				ie.NewRemoteGTPUPeer(0x0e, "127.0.0.1", "", ie.DstInterfaceAccess, "some.instance.example"),
				ie.NewAveragePacketDelay(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.AveragePacketDelay() },
		}, {
			description: "AveragePacketDelay/GTPUPathQoSReport",
			structured: ie.NewGTPUPathQoSReport(
				ie.NewRemoteGTPUPeer(0x0e, "127.0.0.1", "", ie.DstInterfaceAccess, "some.instance.example"),
				ie.NewAveragePacketDelay(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.AveragePacketDelay() },
		}, {
			description: "AveragePacketDelay/QoSInformationInGTPUPathQoSReport",
			structured: ie.NewQoSInformationInGTPUPathQoSReport(
				ie.NewTransportLevelMarking(0x1111),
				ie.NewAveragePacketDelay(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.AveragePacketDelay() },
		}, {
			description: "DLBufferingDuration/20hr",
			structured:  ie.NewDLBufferingDuration(20 * time.Hour),
			decoded:     20 * time.Hour,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.DLBufferingDuration() },
		}, {
			description: "DLBufferingDuration/15min",
			structured:  ie.NewDLBufferingDuration(15 * time.Minute),
			decoded:     15 * time.Minute,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.DLBufferingDuration() },
		}, {
			description: "DLBufferingDuration/30sec",
			structured:  ie.NewDLBufferingDuration(30 * time.Second),
			decoded:     30 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.DLBufferingDuration() },
		}, {
			description: "DLBufferingDuration/UpdateBARWithinSessionReportResponse",
			structured: ie.NewUpdateBARWithinSessionReportResponse(
				ie.NewBARID(0xff),
				ie.NewDLBufferingDuration(30*time.Second),
			),
			decoded:     30 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.DLBufferingDuration() },
		}, {
			description: "DownlinkDataNotificationDelay",
			structured:  ie.NewDownlinkDataNotificationDelay(100 * time.Millisecond),
			decoded:     100 * time.Millisecond,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.DownlinkDataNotificationDelay() },
		}, {
			description: "DownlinkDataNotificationDelay/CreateBAR",
			structured: ie.NewCreateBAR(
				ie.NewBARID(0xff),
				ie.NewDownlinkDataNotificationDelay(100*time.Millisecond),
			),
			decoded:     100 * time.Millisecond,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.DownlinkDataNotificationDelay() },
		}, {
			description: "DownlinkDataNotificationDelay/UpdateBARWithinSessionReportResponse",
			structured: ie.NewUpdateBARWithinSessionReportResponse(
				ie.NewBARID(0xff),
				ie.NewDownlinkDataNotificationDelay(100*time.Millisecond),
			),
			decoded:     100 * time.Millisecond,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.DownlinkDataNotificationDelay() },
		}, {
			description: "DownlinkDataNotificationDelay/UpdateBARWithinSessionModificationRequest",
			structured: ie.NewUpdateBARWithinSessionModificationRequest(
				ie.NewBARID(0xff),
				ie.NewDownlinkDataNotificationDelay(100*time.Millisecond),
			),
			decoded:     100 * time.Millisecond,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.DownlinkDataNotificationDelay() },
		}, {
			description: "DurationMeasurement",
			structured:  ie.NewDurationMeasurement(10 * time.Second),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.DurationMeasurement() },
		}, {
			description: "DurationMeasurement/UsageReportWithinSessionModificationResponse",
			structured: ie.NewUsageReportWithinSessionModificationResponse(
				ie.NewURRID(0xffffffff),
				ie.NewDurationMeasurement(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.DurationMeasurement() },
		}, {
			description: "DurationMeasurement/UsageReportWithinSessionDeletionResponse",
			structured: ie.NewUsageReportWithinSessionDeletionResponse(
				ie.NewURRID(0xffffffff),
				ie.NewDurationMeasurement(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.DurationMeasurement() },
		}, {
			description: "DurationMeasurement/UsageReportWithinSessionReportRequest",
			structured: ie.NewUsageReportWithinSessionReportRequest(
				ie.NewURRID(0xffffffff),
				ie.NewDurationMeasurement(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.DurationMeasurement() },
		}, {
			description: "EthernetInactivityTimer",
			structured:  ie.NewEthernetInactivityTimer(10 * time.Second),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.EthernetInactivityTimer() },
		}, {
			description: "EthernetInactivityTimer/CreateURR",
			structured: ie.NewCreateURR(
				ie.NewURRID(0xffffffff),
				ie.NewEthernetInactivityTimer(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.EthernetInactivityTimer() },
		}, {
			description: "EthernetInactivityTimer/UpdateURR",
			structured: ie.NewUpdateURR(
				ie.NewURRID(0xffffffff),
				ie.NewEthernetInactivityTimer(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.EthernetInactivityTimer() },
		}, {
			description: "GracefulReleasePeriod/20hr",
			structured:  ie.NewGracefulReleasePeriod(20 * time.Hour),
			decoded:     20 * time.Hour,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.GracefulReleasePeriod() },
		}, {
			description: "GracefulReleasePeriod/15min",
			structured:  ie.NewGracefulReleasePeriod(15 * time.Minute),
			decoded:     15 * time.Minute,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.GracefulReleasePeriod() },
		}, {
			description: "GracefulReleasePeriod/30sec",
			structured:  ie.NewGracefulReleasePeriod(30 * time.Second),
			decoded:     30 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.GracefulReleasePeriod() },
		}, {
			description: "MaximumPacketDelay",
			structured:  ie.NewMaximumPacketDelay(10 * time.Second),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.MaximumPacketDelay() },
		}, {
			description: "MaximumPacketDelay/GTPUPathQoSControlInformation",
			structured: ie.NewGTPUPathQoSControlInformation(
				ie.NewRemoteGTPUPeer(0x0e, "127.0.0.1", "", ie.DstInterfaceAccess, "some.instance.example"),
				ie.NewMaximumPacketDelay(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.MaximumPacketDelay() },
		}, {
			description: "MaximumPacketDelay/GTPUPathQoSReport",
			structured: ie.NewGTPUPathQoSReport(
				ie.NewRemoteGTPUPeer(0x0e, "127.0.0.1", "", ie.DstInterfaceAccess, "some.instance.example"),
				ie.NewMaximumPacketDelay(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.MaximumPacketDelay() },
		}, {
			description: "MaximumPacketDelay/QoSInformationInGTPUPathQoSReport",
			structured: ie.NewQoSInformationInGTPUPathQoSReport(
				ie.NewTransportLevelMarking(0x1111),
				ie.NewMaximumPacketDelay(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.MaximumPacketDelay() },
		}, {
			description: "MeasurementPeriod",
			structured:  ie.NewMeasurementPeriod(10 * time.Second),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.MeasurementPeriod() },
		}, {
			description: "MeasurementPeriod/CreateURR",
			structured: ie.NewCreateURR(
				ie.NewURRID(0xffffffff),
				ie.NewMeasurementPeriod(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.MeasurementPeriod() },
		}, {
			description: "MeasurementPeriod/UpdateURR",
			structured: ie.NewUpdateURR(
				ie.NewURRID(0xffffffff),
				ie.NewMeasurementPeriod(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.MeasurementPeriod() },
		}, {
			description: "MeasurementPeriod/QoSMonitoringPerQoSFlowControlInformation",
			structured: ie.NewQoSMonitoringPerQoSFlowControlInformation(
				ie.NewQFI(0x01),
				ie.NewMeasurementPeriod(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.MeasurementPeriod() },
		}, {
			description: "MinimumPacketDelay",
			structured:  ie.NewMinimumPacketDelay(10 * time.Second),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.MinimumPacketDelay() },
		}, {
			description: "MinimumPacketDelay/GTPUPathQoSControlInformation",
			structured: ie.NewGTPUPathQoSControlInformation(
				ie.NewRemoteGTPUPeer(0x0e, "127.0.0.1", "", ie.DstInterfaceAccess, "some.instance.example"),
				ie.NewMinimumPacketDelay(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.MinimumPacketDelay() },
		}, {
			description: "MinimumPacketDelay/GTPUPathQoSReport",
			structured: ie.NewGTPUPathQoSReport(
				ie.NewRemoteGTPUPeer(0x0e, "127.0.0.1", "", ie.DstInterfaceAccess, "some.instance.example"),
				ie.NewMinimumPacketDelay(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.MinimumPacketDelay() },
		}, {
			description: "MinimumPacketDelay/QoSInformationInGTPUPathQoSReport",
			structured: ie.NewQoSInformationInGTPUPathQoSReport(
				ie.NewTransportLevelMarking(0x1111),
				ie.NewMinimumPacketDelay(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.MinimumPacketDelay() },
		}, {
			description: "MinimumWaitTime",
			structured:  ie.NewMinimumWaitTime(10 * time.Second),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.MinimumWaitTime() },
		}, {
			description: "MinimumWaitTime/QoSMonitoringPerQoSFlowControlInformation",
			structured: ie.NewQoSMonitoringPerQoSFlowControlInformation(
				ie.NewQFI(0x01),
				ie.NewMinimumWaitTime(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.MinimumWaitTime() },
		}, {
			description: "QuotaHoldingTime",
			structured:  ie.NewQuotaHoldingTime(10 * time.Second),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.QuotaHoldingTime() },
		}, {
			description: "QuotaHoldingTime/CreateURR",
			structured: ie.NewCreateURR(
				ie.NewURRID(0xffffffff),
				ie.NewQuotaHoldingTime(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.QuotaHoldingTime() },
		}, {
			description: "QuotaHoldingTime/UpdateURR",
			structured: ie.NewUpdateURR(
				ie.NewURRID(0xffffffff),
				ie.NewQuotaHoldingTime(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.QuotaHoldingTime() },
		}, {
			description: "SubsequentTimeQuota",
			structured:  ie.NewSubsequentTimeQuota(10 * time.Second),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.SubsequentTimeQuota() },
		}, {
			description: "SubsequentTimeQuota/CreateURR",
			structured: ie.NewCreateURR(
				ie.NewURRID(0xffffffff),
				ie.NewSubsequentTimeQuota(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.SubsequentTimeQuota() },
		}, {
			description: "SubsequentTimeQuota/UpdateURR",
			structured: ie.NewUpdateURR(
				ie.NewURRID(0xffffffff),
				ie.NewSubsequentTimeQuota(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.SubsequentTimeQuota() },
		}, {
			description: "SubsequentTimeQuota/AdditionalMonitoringTime",
			structured: ie.NewAdditionalMonitoringTime(
				ie.NewMonitoringTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
				ie.NewSubsequentTimeQuota(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.SubsequentTimeQuota() },
		}, {
			description: "TimeOffsetMeasurement",
			structured:  ie.NewTimeOffsetMeasurement(10 * time.Second),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.TimeOffsetMeasurement() },
		}, {
			description: "TimeOffsetThreshold",
			structured:  ie.NewTimeOffsetThreshold(10 * time.Second),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.TimeOffsetThreshold() },
		}, {
			description: "TimeOffsetThreshold/ClockDriftControlInformation",
			structured: ie.NewClockDriftControlInformation(
				ie.NewRequestedClockDriftInformation(1, 1),
				ie.NewTimeOffsetThreshold(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.TimeOffsetThreshold() },
		}, {
			description: "TimeOffsetThreshold/ClockDriftReport",
			structured: ie.NewClockDriftReport(
				ie.NewTSNTimeDomainNumber(255),
				ie.NewTimeOffsetThreshold(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.TimeOffsetThreshold() },
		}, {
			description: "TimeQuota",
			structured:  ie.NewTimeQuota(10 * time.Second),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.TimeQuota() },
		}, {
			description: "TimeQuota/CreateURR",
			structured: ie.NewCreateURR(
				ie.NewURRID(0xffffffff),
				ie.NewTimeQuota(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.TimeQuota() },
		}, {
			description: "TimeQuota/UpdateURR",
			structured: ie.NewUpdateURR(
				ie.NewURRID(0xffffffff),
				ie.NewTimeQuota(10*time.Second),
			),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.TimeQuota() },
		}, {
			description: "Timer/20hr",
			structured:  ie.NewTimer(20 * time.Hour),
			decoded:     20 * time.Hour,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.Timer() },
		}, {
			description: "Timer/15min",
			structured:  ie.NewTimer(15 * time.Minute),
			decoded:     15 * time.Minute,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.Timer() },
		}, {
			description: "Timer/30sec",
			structured:  ie.NewTimer(30 * time.Second),
			decoded:     30 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.Timer() },
		}, {
			description: "Timer/OverloadControlInformation",
			structured: ie.NewOverloadControlInformation(
				ie.NewSequenceNumber(0xffffffff),
				ie.NewTimer(30*time.Second),
			),
			decoded:     30 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.Timer() },
		}, {
			description: "Timer/GTPUPathQoSControlInformation",
			structured: ie.NewGTPUPathQoSControlInformation(
				ie.NewRemoteGTPUPeer(0x0e, "127.0.0.1", "", ie.DstInterfaceAccess, "some.instance.example"),
				ie.NewTimer(30*time.Second),
			),
			decoded:     30 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.Timer() },
		}, {
			description: "UserPlaneInactivityTimer",
			structured:  ie.NewUserPlaneInactivityTimer(10 * time.Second),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.UserPlaneInactivityTimer() },
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
