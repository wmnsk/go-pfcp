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
			description: "DLBufferingSuggestedPacketCount/UpdateBARWithinSessionReportResponse",
			structured: ie.NewUpdateBARWithinSessionReportResponse(
				ie.NewBARID(0xff),
				ie.NewDownlinkDataNotificationDelay(100*time.Millisecond),
				ie.NewDLBufferingDuration(30*time.Second),
				ie.NewDLBufferingSuggestedPacketCount(0xffff),
				ie.NewSuggestedBufferingPacketsCount(0x01),
			),
			decoded:     0xffff,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.DLBufferingSuggestedPacketCount() },
		}, {
			description: "DLDataPacketsSize",
			structured:  ie.NewDLDataPacketsSize(0xffff),
			decoded:     0xffff,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.DLDataPacketsSize() },
		}, {
			description: "DLDataPacketsSize/DownlinkDataReport",
			structured: ie.NewDownlinkDataReport(
				ie.NewPDRID(0xffff),
				ie.NewDLDataPacketsSize(0xffff),
			),
			decoded:     0xffff,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.DLDataPacketsSize() },
		}, {
			description: "Ethertype",
			structured:  ie.NewEthertype(0xffff),
			decoded:     0xffff,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.Ethertype() },
		}, {
			description: "Ethertype/PDI",
			structured: ie.NewPDI(
				ie.NewSourceInterface(ie.SrcInterfaceAccess),
				ie.NewEthertype(0xffff),
			),
			decoded:     0xffff,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.Ethertype() },
		}, {
			description: "Ethertype/EthernetPacketFilter",
			structured: ie.NewEthernetPacketFilter(
				ie.NewEthernetFilterID(0xffffffff),
				ie.NewEthertype(0xffff),
			),
			decoded:     0xffff,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.Ethertype() },
		}, {
			description: "MARID",
			structured:  ie.NewMARID(0x1111),
			decoded:     0x1111,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.MARID() },
		}, {
			description: "MARID/CreatePDR",
			structured: ie.NewCreatePDR(
				ie.NewPDRID(0xffff),
				ie.NewMARID(0x1111),
			),
			decoded:     0x1111,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.MARID() },
		}, {
			description: "MARID/CreateMAR",
			structured: ie.NewCreateMAR(
				ie.NewMARID(0x1111),
				ie.NewSteeringFunctionality(ie.SteeringFunctionalityATSSSLL),
			),
			decoded:     0x1111,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.MARID() },
		}, {
			description: "MARID/RemoveMAR",
			structured: ie.NewRemoveMAR(
				ie.NewMARID(0x1111),
			),
			decoded:     0x1111,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.MARID() },
		}, {
			description: "MARID/UpdateMAR",
			structured: ie.NewUpdateMAR(
				ie.NewMARID(0x1111),
				ie.NewSteeringFunctionality(ie.SteeringFunctionalityATSSSLL),
			),
			decoded:     0x1111,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.MARID() },
		}, {
			description: "NumberOfReports",
			structured:  ie.NewNumberOfReports(0x1111),
			decoded:     0x1111,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.NumberOfReports() },
		}, {
			description: "NumberOfReports/CreateURR",
			structured: ie.NewCreateURR(
				ie.NewURRID(0xffffffff),
				ie.NewNumberOfReports(0x1111),
			),
			decoded:     0x1111,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.NumberOfReports() },
		}, {
			description: "NumberOfReports/UpdateURR",
			structured: ie.NewUpdateURR(
				ie.NewURRID(0xffffffff),
				ie.NewNumberOfReports(0x1111),
			),
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
			description: "PDRID/CreatePDR",
			structured: ie.NewCreatePDR(
				ie.NewPDRID(0xffff),
				ie.NewPrecedence(0x11111111),
			),
			decoded:     0xffff,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.PDRID() },
		}, {
			description: "PDRID/UpdatePDR",
			structured: ie.NewUpdatePDR(
				ie.NewPDRID(0xffff),
				ie.NewPrecedence(0x11111111),
			),
			decoded:     0xffff,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.PDRID() },
		}, {
			description: "PDRID/RemovePDR",
			structured: ie.NewRemovePDR(
				ie.NewPDRID(0xffff),
			),
			decoded:     0xffff,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.PDRID() },
		}, {
			description: "PDRID/CreatedPDR",
			structured: ie.NewCreatedPDR(
				ie.NewPDRID(0xffff),
				ie.NewFTEID(0x11111111, net.ParseIP("127.0.0.1"), nil, nil),
			),
			decoded:     0xffff,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.PDRID() },
		}, {
			description: "PDRID/ApplicationDetectionInformation",
			structured: ie.NewApplicationDetectionInformation(
				ie.NewApplicationID("https://github.com/wmnsk/go-pfcp/"),
				ie.NewPDRID(0xffff),
			),
			decoded:     0xffff,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.PDRID() },
		}, {
			description: "PDRID/UsageReportWithinSessionReportRequest",
			structured: ie.NewUsageReportWithinSessionReportRequest(
				ie.NewURRID(0xffffffff),
				ie.NewApplicationDetectionInformation(
					ie.NewApplicationID("https://github.com/wmnsk/go-pfcp/"),
					ie.NewPDRID(0xffff),
				),
			),
			decoded:     0xffff,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.PDRID() },
		}, {
			description: "PDRID/DownlinkDataReport",
			structured: ie.NewDownlinkDataReport(
				ie.NewPDRID(0xffff),
				ie.NewDownlinkDataServiceInformation(true, true, 0xff, 0xff),
			),
			decoded:     0xffff,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.PDRID() },
		}, {
			description: "PDRID/UpdatedPDR",
			structured: ie.NewUpdatedPDR(
				ie.NewPDRID(0xffff),
				ie.NewFTEID(0x11111111, net.ParseIP("127.0.0.1"), nil, nil),
			),
			decoded:     0xffff,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.PDRID() },
		}, {
			description: "ReportingTriggers",
			structured:  ie.NewReportingTriggers(0x1122),
			decoded:     0x1122,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.ReportingTriggers() },
		}, {
			description: "ReportingTriggers/CreateURR",
			structured: ie.NewCreateURR(
				ie.NewURRID(0xffffffff),
				ie.NewReportingTriggers(0x1122),
			),
			decoded:     0x1122,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.ReportingTriggers() },
		}, {
			description: "ReportingTriggers/UpdateURR",
			structured: ie.NewUpdateURR(
				ie.NewURRID(0xffffffff),
				ie.NewReportingTriggers(0x1122),
			),
			decoded:     0x1122,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.ReportingTriggers() },
		}, {
			description: "TransportLevelMarking",
			structured:  ie.NewTransportLevelMarking(0x1111),
			decoded:     0x1111,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.TransportLevelMarking() },
		}, {
			description: "TransportLevelMarking/ForwardingParameters",
			structured: ie.NewForwardingParameters(
				ie.NewDestinationInterface(ie.DstInterfaceAccess),
				ie.NewTransportLevelMarking(0x1111),
			),
			decoded:     0x1111,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.TransportLevelMarking() },
		}, {
			description: "TransportLevelMarking/UpdateForwardingParameters",
			structured: ie.NewUpdateForwardingParameters(
				ie.NewDestinationInterface(ie.DstInterfaceAccess),
				ie.NewTransportLevelMarking(0x1111),
			),
			decoded:     0x1111,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.TransportLevelMarking() },
		}, {
			description: "TransportLevelMarking/DuplicatingParameters",
			structured: ie.NewDuplicatingParameters(
				ie.NewDestinationInterface(ie.DstInterfaceAccess),
				ie.NewTransportLevelMarking(0x1111),
			),
			decoded:     0x1111,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.TransportLevelMarking() },
		}, {
			description: "TransportLevelMarking/UpdateDuplicatingParameters",
			structured: ie.NewUpdateDuplicatingParameters(
				ie.NewDestinationInterface(ie.DstInterfaceAccess),
				ie.NewTransportLevelMarking(0x1111),
			),
			decoded:     0x1111,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.TransportLevelMarking() },
		}, {
			description: "TransportLevelMarking/GTPUPathQoSControlInformation",
			structured: ie.NewGTPUPathQoSControlInformation(
				ie.NewRemoteGTPUPeer(0x0e, "127.0.0.1", "", ie.DstInterfaceAccess, "some.instance.example"),
				ie.NewTransportLevelMarking(0x1111),
			),
			decoded:     0x1111,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.TransportLevelMarking() },
		}, {
			description: "TransportLevelMarking/GTPUPathQoSReport",
			structured: ie.NewGTPUPathQoSReport(
				ie.NewRemoteGTPUPeer(0x0e, "127.0.0.1", "", ie.DstInterfaceAccess, "some.instance.example"),
				ie.NewTransportLevelMarking(0x1111),
			),
			decoded:     0x1111,
			decoderFunc: func(i *ie.IE) (uint16, error) { return i.TransportLevelMarking() },
		}, {
			description: "TransportLevelMarking/QoSInformationInGTPUPathQoSReport",
			structured: ie.NewQoSInformationInGTPUPathQoSReport(
				ie.NewAveragePacketDelay(10*time.Second),
				ie.NewTransportLevelMarking(0x1111),
			),
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
