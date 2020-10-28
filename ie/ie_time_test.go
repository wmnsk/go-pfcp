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
			description: "DeactivationTime",
			structured:  ie.NewDeactivationTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.DeactivationTime() },
		}, {
			description: "EndTime",
			structured:  ie.NewEndTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.EndTime() },
		}, {
			description: "EventTimeStamp",
			structured:  ie.NewEventTimeStamp(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.EventTimeStamp() },
		}, {
			description: "MonitoringTime",
			structured:  ie.NewMonitoringTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.MonitoringTime() },
		}, {
			description: "QuotaValidityTime",
			structured:  ie.NewQuotaValidityTime(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
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
			description: "TimeOfFirstPacket",
			structured:  ie.NewTimeOfFirstPacket(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
			decoded:     time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC),
			decoderFunc: func(i *ie.IE) (time.Time, error) { return i.TimeOfFirstPacket() },
		}, {
			description: "TimeOfLastPacket",
			structured:  ie.NewTimeOfLastPacket(time.Date(2019, time.January, 1, 0, 0, 0, 0, time.UTC)),
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
