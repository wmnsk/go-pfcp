// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wmnsk/go-pfcp/ie"
)

func TestUint16IEs(t *testing.T) {
	cases := []struct {
		description string
		structured  *ie.IE
		decoded     uint16
		decoderFunc func(*ie.IE) (uint16, error)
	}{
		{
			description: "AdditionalUsageReportsInformation",
			structured:  ie.NewAdditionalUsageReportsInformation(0x00ff),
			decoded:     0x00ff,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.AdditionalUsageReportsInformation() },
		}, {
			description: "DLBufferingSuggestedPacketCount",
			structured:  ie.NewDLBufferingSuggestedPacketCount(0xffff),
			decoded:     0xffff,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.DLBufferingSuggestedPacketCount() },
		}, {
			description: "DLDataPacketsSize",
			structured:  ie.NewDLDataPacketsSize(0xffff),
			decoded:     0xffff,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.DLDataPacketsSize() },
		}, {
			description: "Ethertype",
			structured:  ie.NewEthertype(0xffff),
			decoded:     0xffff,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.Ethertype() },
		}, {
			description: "MARID",
			structured:  ie.NewMARID(0x1111),
			decoded:     0x1111,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.MARID() },
		}, {
			description: "NumberOfReports",
			structured:  ie.NewNumberOfReports(0x1111),
			decoded:     0x1111,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.NumberOfReports() },
		}, {
			description: "OffendingIE",
			structured:  ie.NewOffendingIE(ie.Cause),
			decoded:     ie.Cause,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.OffendingIE() },
		}, {
			description: "PDRID",
			structured:  ie.NewPDRID(0xffff),
			decoded:     0xffff,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.PDRID() },
		}, {
			description: "ReportingTriggers",
			structured:  ie.NewReportingTriggers(0x1122),
			decoded:     0x1122,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.ReportingTriggers() },
		}, {
			description: "TransportLevelMarking",
			structured:  ie.NewTransportLevelMarking(0x1111),
			decoded:     0x1111,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.TransportLevelMarking() },
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
