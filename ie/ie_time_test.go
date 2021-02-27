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

func TestTimeIEs(t *testing.T) {
	cases := []struct {
		description string
		structured  *ie.IE
		decoded     time.Time
		decoderFunc func(*ie.IE) (time.Time, error)
	}{
		{
			description: "ActivationTime",
			structured:  ie.NewActivationTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.ActivationTime() },
		}, {
			description: "ActivationTime/CreatePDR",
			structured: ie.NewCreatePDR(
				ie.NewPDRID(0xffff),
				ie.NewActivationTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.ActivationTime() },
		}, {
			description: "ActivationTime/UpdatePDR",
			structured: ie.NewUpdatePDR(
				ie.NewPDRID(0xffff),
				ie.NewActivationTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.ActivationTime() },
		}, {
			description: "DeactivationTime",
			structured:  ie.NewDeactivationTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.DeactivationTime() },
		}, {
			description: "DeactivationTime/CreatePDR",
			structured: ie.NewCreatePDR(
				ie.NewPDRID(0xffff),
				ie.NewDeactivationTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.DeactivationTime() },
		}, {
			description: "DeactivationTime/UpdatePDR",
			structured: ie.NewUpdatePDR(
				ie.NewPDRID(0xffff),
				ie.NewDeactivationTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.DeactivationTime() },
		}, {
			description: "EndTime",
			structured:  ie.NewEndTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.EndTime() },
		}, {
			description: "EndTime/UsageReportWithinSessionModificationResponse",
			structured: ie.NewUsageReportWithinSessionModificationResponse(
				ie.NewURRID(0xffffffff),
				ie.NewEndTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.EndTime() },
		}, {
			description: "EndTime/UsageReportWithinSessionDeletionResponse",
			structured: ie.NewUsageReportWithinSessionDeletionResponse(
				ie.NewURRID(0xffffffff),
				ie.NewEndTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.EndTime() },
		}, {
			description: "EndTime/UsageReportWithinSessionReportRequest",
			structured: ie.NewUsageReportWithinSessionReportRequest(
				ie.NewURRID(0xffffffff),
				ie.NewEndTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.EndTime() },
		}, {
			description: "EventTimeStamp",
			structured:  ie.NewEventTimeStamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.EventTimeStamp() },
		}, {
			description: "EventTimeStamp/UsageReportWithinSessionReportRequest",
			structured: ie.NewUsageReportWithinSessionReportRequest(
				ie.NewURRID(0xffffffff),
				ie.NewEventTimeStamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.EventTimeStamp() },
		}, {
			description: "EventTimeStamp/GTPUPathQoSReport",
			structured: ie.NewGTPUPathQoSReport(
				ie.NewRemoteGTPUPeer(0x0e, "127.0.0.1", "", ie.DstInterfaceAccess, "some.instance.example"),
				ie.NewEventTimeStamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.EventTimeStamp() },
		}, {
			description: "EventTimeStamp/QoSMonitoringReport",
			structured: ie.NewQoSMonitoringReport(
				ie.NewQFI(0x01),
				ie.NewEventTimeStamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.EventTimeStamp() },
		}, {
			description: "EventTimeStamp/ClockDriftReport",
			structured: ie.NewClockDriftReport(
				ie.NewTSNTimeDomainNumber(255),
				ie.NewEventTimeStamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.EventTimeStamp() },
		}, {
			description: "MonitoringTime",
			structured:  ie.NewMonitoringTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.MonitoringTime() },
		}, {
			description: "MonitoringTime/CreateURR",
			structured: ie.NewCreateURR(
				ie.NewURRID(0xffffffff),
				ie.NewMonitoringTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.MonitoringTime() },
		}, {
			description: "MonitoringTime/UpdateURR",
			structured: ie.NewUpdateURR(
				ie.NewURRID(0xffffffff),
				ie.NewMonitoringTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.MonitoringTime() },
		}, {
			description: "MonitoringTime/AdditionalMonitoringTime",
			structured: ie.NewAdditionalMonitoringTime(
				ie.NewSubsequentVolumeThreshold(0x01, 0x1111111111111111, 0, 0),
				ie.NewMonitoringTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.MonitoringTime() },
		}, {
			description: "QuotaValidityTime",
			structured:  ie.NewQuotaValidityTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.QuotaValidityTime() },
		}, {
			description: "QuotaValidityTime/CreateURR",
			structured: ie.NewCreateURR(
				ie.NewURRID(0xffffffff),
				ie.NewQuotaValidityTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.QuotaValidityTime() },
		}, {
			description: "QuotaValidityTime/UpdateURR",
			structured: ie.NewUpdateURR(
				ie.NewURRID(0xffffffff),
				ie.NewQuotaValidityTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.QuotaValidityTime() },
		}, {
			description: "RecoveryTimeStamp",
			structured:  ie.NewRecoveryTimeStamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.RecoveryTimeStamp() },
		}, {
			description: "StartTime",
			structured:  ie.NewStartTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.StartTime() },
		}, {
			description: "StartTime/UsageReportWithinSessionModificationResponse",
			structured: ie.NewUsageReportWithinSessionModificationResponse(
				ie.NewURRID(0xffffffff),
				ie.NewStartTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.StartTime() },
		}, {
			description: "StartTime/UsageReportWithinSessionDeletionResponse",
			structured: ie.NewUsageReportWithinSessionDeletionResponse(
				ie.NewURRID(0xffffffff),
				ie.NewStartTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.StartTime() },
		}, {
			description: "StartTime/UsageReportWithinSessionReportRequest",
			structured: ie.NewUsageReportWithinSessionReportRequest(
				ie.NewURRID(0xffffffff),
				ie.NewStartTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.StartTime() },
		}, {
			description: "StartTime/GTPUPathQoSReport",
			structured: ie.NewGTPUPathQoSReport(
				ie.NewRemoteGTPUPeer(0x0e, "127.0.0.1", "", ie.DstInterfaceAccess, "some.instance.example"),
				ie.NewStartTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.StartTime() },
		}, {
			description: "StartTime/QoSMonitoringReport",
			structured: ie.NewQoSMonitoringReport(
				ie.NewQFI(0x01),
				ie.NewStartTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.StartTime() },
		}, {
			description: "TimeOfFirstPacket",
			structured:  ie.NewTimeOfFirstPacket(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.TimeOfFirstPacket() },
		}, {
			description: "TimeOfFirstPacket/UsageReportWithinSessionModificationResponse",
			structured: ie.NewUsageReportWithinSessionModificationResponse(
				ie.NewURRID(0xffffffff),
				ie.NewTimeOfFirstPacket(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.TimeOfFirstPacket() },
		}, {
			description: "TimeOfFirstPacket/UsageReportWithinSessionDeletionResponse",
			structured: ie.NewUsageReportWithinSessionDeletionResponse(
				ie.NewURRID(0xffffffff),
				ie.NewTimeOfFirstPacket(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.TimeOfFirstPacket() },
		}, {
			description: "TimeOfFirstPacket/UsageReportWithinSessionReportRequest",
			structured: ie.NewUsageReportWithinSessionReportRequest(
				ie.NewURRID(0xffffffff),
				ie.NewTimeOfFirstPacket(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.TimeOfFirstPacket() },
		}, {
			description: "TimeOfLastPacket",
			structured:  ie.NewTimeOfLastPacket(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.TimeOfLastPacket() },
		}, {
			description: "TimeOfLastPacket/UsageReportWithinSessionModificationResponse",
			structured: ie.NewUsageReportWithinSessionModificationResponse(
				ie.NewURRID(0xffffffff),
				ie.NewTimeOfLastPacket(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.TimeOfLastPacket() },
		}, {
			description: "TimeOfLastPacket/UsageReportWithinSessionDeletionResponse",
			structured: ie.NewUsageReportWithinSessionDeletionResponse(
				ie.NewURRID(0xffffffff),
				ie.NewTimeOfLastPacket(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.TimeOfLastPacket() },
		}, {
			description: "TimeOfLastPacket/UsageReportWithinSessionReportRequest",
			structured: ie.NewUsageReportWithinSessionReportRequest(
				ie.NewURRID(0xffffffff),
				ie.NewTimeOfLastPacket(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.TimeOfLastPacket() },
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
