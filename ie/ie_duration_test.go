// Copyright 2019-2020 go-pfcp authors. All rights reserved.
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
			description: "DownlinkDataNotificationDelay",
			structured:  ie.NewDownlinkDataNotificationDelay(100 * time.Millisecond),
			decoded:     100 * time.Millisecond,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.DownlinkDataNotificationDelay() },
		}, {
			description: "DurationMeasurement",
			structured:  ie.NewDurationMeasurement(10 * time.Second),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.DurationMeasurement() },
		}, {
			description: "EthernetInactivityTimer",
			structured:  ie.NewEthernetInactivityTimer(10 * time.Second),
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
			description: "MinimumPacketDelay",
			structured:  ie.NewMinimumPacketDelay(10 * time.Second),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.MinimumPacketDelay() },
		}, {
			description: "MinimumWaitTime",
			structured:  ie.NewMinimumWaitTime(10 * time.Second),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.MinimumWaitTime() },
		}, {
			description: "QuotaHoldingTime",
			structured:  ie.NewQuotaHoldingTime(10 * time.Second),
			decoded:     10 * time.Second,
			decoderFunc: func(i *ie.IE) (time.Duration, error) { return i.QuotaHoldingTime() },
		}, {
			description: "SubsequentTimeQuota",
			structured:  ie.NewSubsequentTimeQuota(10 * time.Second),
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
			description: "TimeQuota",
			structured:  ie.NewTimeQuota(10 * time.Second),
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
