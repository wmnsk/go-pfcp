// Copyright 2019-2020 go-pfcp authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package ie_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/wmnsk/go-pfcp/ie"
)

func TestUint8IEs(t *testing.T) {
	cases := []struct {
		description string
		structured  *ie.IE
		decoded     uint8
		decoderFunc func(*ie.IE) (uint8, error)
	}{
		{
			description: "AccessAvailabilityInformation",
			structured:  ie.NewAccessAvailabilityInformation(3, 3),
			decoded:     0x0f,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.AccessAvailabilityInformation() },
		}, {
			description: "ApplyAction",
			structured:  ie.NewApplyAction(4),
			decoded:     4,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.ApplyAction() },
		}, {
			description: "ATSSSLLControlInformation",
			structured:  ie.NewATSSSLLControlInformation(1),
			decoded:     1,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.ATSSSLLControlInformation() },
		}, {
			description: "ATSSSLLInformation",
			structured:  ie.NewATSSSLLInformation(1),
			decoded:     1,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.ATSSSLLInformation() },
		}, {
			description: "BARID",
			structured:  ie.NewBARID(0xff),
			decoded:     0xff,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.BARID() },
		}, {
			description: "Cause",
			structured:  ie.NewCause(ie.CauseRequestAccepted),
			decoded:     ie.CauseRequestAccepted,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.Cause() },
		}, {
			description: "CPFunctionFeatures",
			structured:  ie.NewCPFunctionFeatures(0x3f),
			decoded:     0x3f,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.CPFunctionFeatures() },
		}, {
			description: "CreateBridgeInfoForTSC",
			structured:  ie.NewCreateBridgeInfoForTSC(1),
			decoded:     1,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.CreateBridgeInfoForTSC() },
		}, {
			description: "DestinationInterface",
			structured:  ie.NewDestinationInterface(ie.DstInterfaceAccess),
			decoded:     ie.DstInterfaceAccess,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.DestinationInterface() },
		}, {
			description: "DLBufferingSuggestedPacketCount",
			structured:  ie.NewDLBufferingSuggestedPacketCount(0xff),
			decoded:     0xff,
			decoderFunc: func(i *ie.IE) (uint8, error) {
				v, err := i.DLBufferingSuggestedPacketCount()
				return uint8(v), err
			},
		}, {
			description: "EthernetFilterProperties",
			structured:  ie.NewEthernetFilterProperties(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.EthernetFilterProperties() },
		}, {
			description: "EthernetPDUSessionInformation",
			structured:  ie.NewEthernetPDUSessionInformation(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.EthernetPDUSessionInformation() },
		}, {
			description: "GateStatus/OpenOpen",
			structured:  ie.NewGateStatus(ie.GateStatusOpen, ie.GateStatusOpen),
			decoded:     0,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.GateStatus() },
		}, {
			description: "GateStatus/OpenClosed",
			structured:  ie.NewGateStatus(ie.GateStatusOpen, ie.GateStatusClosed),
			decoded:     1,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.GateStatus() },
		}, {
			description: "GateStatus/ClosedOpen",
			structured:  ie.NewGateStatus(ie.GateStatusClosed, ie.GateStatusOpen),
			decoded:     4,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.GateStatus() },
		}, {
			description: "GateStatus/ClosedClosed",
			structured:  ie.NewGateStatus(ie.GateStatusClosed, ie.GateStatusClosed),
			decoded:     5,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.GateStatus() },
		}, {
			description: "GTPUPathInterfaceType",
			structured:  ie.NewGTPUPathInterfaceType(1, 1),
			decoded:     3,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.GTPUPathInterfaceType() },
		}, {
			description: "MeasurementInformation",
			structured:  ie.NewMeasurementInformation(0x1f),
			decoded:     0x1f,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.MeasurementInformation() },
		}, {
			description: "Metric",
			structured:  ie.NewMetric(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.Metric() },
		}, {
			description: "MPTCPControlInformation",
			structured:  ie.NewMPTCPControlInformation(1),
			decoded:     1,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.MPTCPControlInformation() },
		}, {
			description: "MTEDTControlInformation",
			structured:  ie.NewMTEDTControlInformation(1),
			decoded:     1,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.MTEDTControlInformation() },
		}, {
			description: "NodeReportType",
			structured:  ie.NewNodeReportType(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.NodeReportType() },
		}, {
			description: "OCIFlags",
			structured:  ie.NewOCIFlags(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.OCIFlags() },
		}, {
			description: "PacketReplicationAndDetectionCarryOnInformation",
			structured:  ie.NewPacketReplicationAndDetectionCarryOnInformation(0x0f),
			decoded:     0x0f,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.PacketReplicationAndDetectionCarryOnInformation() },
		}, {
			description: "PagingPolicyIndicator",
			structured:  ie.NewPagingPolicyIndicator(1),
			decoded:     1,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.PagingPolicyIndicator() },
		}, {
			description: "PDNType",
			structured:  ie.NewPDNType(ie.PDNTypeIPv4),
			decoded:     ie.PDNTypeIPv4,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.PDNType() },
		}, {
			description: "PFCPAssociationReleaseRequest",
			structured:  ie.NewPFCPAssociationReleaseRequest(1, 1),
			decoded:     3,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.PFCPAssociationReleaseRequest() },
		}, {
			description: "PFCPASRspFlags",
			structured:  ie.NewPFCPASRspFlags(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.PFCPASRspFlags() },
		}, {
			description: "PFCPAUReqFlags",
			structured:  ie.NewPFCPAUReqFlags(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.PFCPAUReqFlags() },
		}, {
			description: "PFCPSEReqFlags",
			structured:  ie.NewPFCPSEReqFlags(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.PFCPSEReqFlags() },
		}, {
			description: "PFCPSRRspFlags",
			structured:  ie.NewPFCPSRRspFlags(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.PFCPSRRspFlags() },
		}, {
			description: "PMFControlInformation",
			structured:  ie.NewPMFControlInformation(0xff),
			decoded:     0x01, // first bit should only be evaluated.
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.PMFControlInformation() },
		}, {
			description: "Priority",
			structured:  ie.NewPriority(ie.PriorityActive),
			decoded:     ie.PriorityActive,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.Priority() },
		}, {
			description: "Proxying",
			structured:  ie.NewProxying(1, 1),
			decoded:     0x03,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.Proxying() },
		}, {
			description: "QERControlIndications",
			structured:  ie.NewQERControlIndications(1, 1, 1),
			decoded:     0x07,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.QERControlIndications() },
		}, {
			description: "QFI",
			structured:  ie.NewQFI(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.QFI() },
		}, {
			description: "QoSReportTrigger",
			structured:  ie.NewQoSReportTrigger(1, 1, 1),
			decoded:     0x07,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.QoSReportTrigger() },
		}, {
			description: "ReportType",
			structured:  ie.NewReportType(1, 1, 1, 1),
			decoded:     0x0f,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.ReportType() },
		}, {
			description: "ReportingFrequency",
			structured:  ie.NewReportingFrequency(1, 1, 1),
			decoded:     0x07,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.ReportingFrequency() },
		}, {
			description: "RequestedAccessAvailabilityInformation",
			structured:  ie.NewRequestedAccessAvailabilityInformation(1),
			decoded:     1,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.RequestedAccessAvailabilityInformation() },
		}, {
			description: "RequestedClockDriftInformation",
			structured:  ie.NewRequestedClockDriftInformation(1, 1),
			decoded:     0x03,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.RequestedClockDriftInformation() },
		}, {
			description: "RequestedQoSMonitoring",
			structured:  ie.NewRequestedQoSMonitoring(1, 1, 1),
			decoded:     0x07,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.RequestedQoSMonitoring() },
		}, {
			description: "RQI",
			structured:  ie.NewRQI(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.RQI() },
		}, {
			description: "SourceInterface",
			structured:  ie.NewSourceInterface(ie.SrcInterfaceAccess),
			decoded:     ie.SrcInterfaceAccess,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.SourceInterface() },
		}, {
			description: "SRRID",
			structured:  ie.NewSRRID(255),
			decoded:     255,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.SRRID() },
		}, {
			description: "SteeringFunctionality",
			structured:  ie.NewSteeringFunctionality(ie.SteeringFunctionalityATSSSLL),
			decoded:     ie.SteeringFunctionalityATSSSLL,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.SteeringFunctionality() },
		}, {
			description: "SteeringMode",
			structured:  ie.NewSteeringMode(ie.SteeringModeActiveStandby),
			decoded:     ie.SteeringModeActiveStandby,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.SteeringMode() },
		}, {
			description: "SuggestedBufferingPacketsCount",
			structured:  ie.NewSuggestedBufferingPacketsCount(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.SuggestedBufferingPacketsCount() },
		}, {
			description: "TGPPInterfaceType",
			structured:  ie.NewTGPPInterfaceType(ie.TGPPInterfaceTypeS1U),
			decoded:     ie.TGPPInterfaceTypeS1U,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.TGPPInterfaceType() },
		}, {
			description: "TrafficEndpointID",
			structured:  ie.NewTrafficEndpointID(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.TrafficEndpointID() },
		}, {
			description: "TSNTimeDomainNumber",
			structured:  ie.NewTSNTimeDomainNumber(255),
			decoded:     255,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.TSNTimeDomainNumber() },
		}, {
			description: "UsageInformation",
			structured:  ie.NewUsageInformation(1, 1, 1, 1),
			decoded:     0x0f,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.UsageInformation() },
		}, {
			description: "Weight",
			structured:  ie.NewWeight(0x01),
			decoded:     0x01,
			decoderFunc: func(i *ie.IE) (uint8, error) { return i.Weight() },
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
